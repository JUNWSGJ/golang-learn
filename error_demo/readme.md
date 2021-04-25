error 使用

### 作业

**问题：** 在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

**回答：**
个人认为， sql.ErrNoRows，代表记录未找到，并不能算是真正意义上的异常。

1. 如果查询的结果集是一个数组，发生sql.ErrNoRows， 可以返回一个空数组。

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

2. 如果查询的结果集是单条记录， 比如说按照id查找用户， 
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


