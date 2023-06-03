package services

import (
    "fmt"

    "gorm.io/gorm"
)

type UserService struct {
    db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
    return &UserService{
        db: db,
    }
}

type User struct {
    ID       uint   `gorm:"primaryKey"`
    Name     string `gorm:"not null"`
    Email    string `gorm:"unique;not null"`
    Password string `gorm:"not null"`
}

func (u *UserService) CreateUser(name, email, password string) error {
    user := User{
        Name:     name,
        Email:    email,
        Password: password,
    }

    result := u.db.Create(&user)
    if result.Error != nil {
        return result.Error
    }

    fmt.Println("User created successfully!")
    return nil
}

