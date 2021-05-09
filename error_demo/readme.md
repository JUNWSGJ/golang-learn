error 使用

### 作业

**问题：** 在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

**回答：**

这个问题可以从两种角度去看，
### 1. sql.ErrNoRows到底应不应该抛异常，如果不抛异常，那可以直接返回一个nil值，
这个返回不是很好，必须要求调用方去检查返回值是不是nil，调用方有可能会忘了处理，
我觉得应该包装一层，类似于sql.NullString这种struct，这样就会显示让调用方检查Valid属性，为true时才会去取值，为false就代表未找到记录，做相应的业务处理
但是这样做也有点麻烦，必须给每个返回的类型都定义一个相应的包装类型。像java里就可以用Optional<T>这种泛型来达到一样的效果，而不必针对每个类型都定义一个包装类型。


1）.如果查询的结果集是一个数组，发生sql.ErrNoRows， 可以返回一个空数组。

```golang
func FindUsersByAge(age uint8) (*[]User, error) {
    users := make([]User, 0)
    err := db.Model(&User{}).Where("age = ?", age).Find(&users).Error
    if err == gorm.ErrRecordNotFound {
        return &users, nil
    } else if err!=nil {
        return &users, errors.Wrap(err, "find users error")
    }
    return &users, err
}
```

2） 如果查询的结果集是单条记录， 比如说按照id查找用户， 
我觉得可以 用 struct 包装 返回的User，类似sql.NullString
```golang
type NullString struct {
    String string
    Valid  bool
}
```

```golang

type NullUser struct {
    User *User
    Valid  bool
}

func FindUserById(id uint) (domain.NullUser, error) {
    user := &domain.User{}
    err := db.Model(&domain.User{}).Where("id = ?", id).First(user).Error
    if err == gorm.ErrRecordNotFound {
        return domain.NullUser{ Valid: false, User: user}, nil
    }
    if  err!= nil {
        return domain.NullUser{}, err
    }
    return domain.NullUser{ Valid: true, User: user}, err
}
```

### 2. 如果把记录未找到当做异常，应该往上层抛什么异常，是将sql.ErrNoRows原样抛出，还是转化成自定义的异常
如果直接wrap sql.ErrNoRows 然后抛出，则上层业务必须得依赖sql.ErrNoRows才能判断是否是记录未找到异常，
这样会导致业务代码依赖底层的存储引擎，之后想要更换存储的引擎会很难。
所以建议在通用库或框架中自定义记录未找到的异常，将sql.ErrNoRows 转换成该异常，然后wrap往上抛。

