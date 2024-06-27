package model

import "time"

type User struct {
	ID           int64     `xorm:"'id' pk autoincr"` // 主键，自增
	Username     string    `xorm:"unique"`           // 用户名，唯一, 可选
	Email        string    `xorm:"unique"`           // 邮箱，唯一
	PasswordHash string    // 存储经过哈希处理后的密码，包含盐
	CreatedAt    time.Time `xorm:"created"` // 创建时间
	UpdatedAt    time.Time `xorm:"updated"` // 更新时间
}

func InsertOneUser(user *User) (int64, error) {
	id, err := dbEngine.Insert(user)
	return id, err
}

func FindFirstUserByUsername(username string) (*User, error) {
	user := &User{Username: username}
	has, err := dbEngine.Get(user)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return user, nil
}
