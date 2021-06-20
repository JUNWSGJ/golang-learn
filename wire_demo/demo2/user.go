package main

import (
	"errors"
	"github.com/google/wire"
)

type User struct {
	Id       int
	Username string
}

// UserService
type UserService struct {
	userRepo UserRepository // <-- UserService依赖UserRepository接口
}

// UserExist 判断指定ID的用户是否存在
func (u *UserService) UserExist(id int) bool {
	_, err := u.userRepo.GetUserByID(id)
	return err == nil
}

// NewUserService *UserService构造函数
func NewUserService(userRepo UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

// UserRepository 存放User对象的数据仓库接口,比如可以是mysql,restful api ....
type UserRepository interface {
	// GetUserByID 根据ID获取User, 如果找不到User返回对应错误信息
	GetUserByID(id int) (*User, error)
}

// mockUserRepo 模拟一个UserRepository实现
type mockUserRepo struct {
	foo string
	bar int
}

// GetUserByID UserRepository接口实现
func (u *mockUserRepo) GetUserByID(id int) (*User, error) {
	if id == u.bar {
		return &User{Id: u.bar, Username: u.foo}, nil
	}
	return nil, errors.New("user not found")
}

// NewMockUserRepo *mockUserRepo构造函数
func NewMockUserRepo(foo string, bar int) *mockUserRepo {
	return &mockUserRepo{
		foo: foo,
		bar: bar,
	}
}

// MockUserRepoSet 将 *mockUserRepo与UserRepository绑定
var MockUserRepoSet = wire.NewSet(NewMockUserRepo, wire.Bind(new(UserRepository), new(*mockUserRepo)))
