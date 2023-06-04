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
    Deleted  bool   `gorm:"default:false"`
}

func (u *UserService) CreateUser(name, email, password string) (*User, error) {
    user := User{
        Name:     name,
        Email:    email,
        Password: password,
    }

    result := u.db.Create(&user)
    if result.Error != nil {
        return nil, result.Error
    }

    fmt.Println("User created successfully!")
    return &user, nil
}

func (u *UserService) EditUser(id uint, name, email, password string) (*User, error) {
    var user User
    result := u.db.First(&user, id)
    if result.Error != nil {
        return nil, result.Error
    }

    user.Name = name
    user.Email = email
    user.Password = password

    result = u.db.Save(&user)
    if result.Error != nil {
        return nil, result.Error
    }

    fmt.Println("User updated successfully!")
    return &user, nil
}

func (u *UserService) GetUserByID(id uint) (*User, error) {
    var user User
    result := u.db.First(&user, id)
    if result.Error != nil {
        return nil, result.Error
    }

    return &user, nil
}

func (u *UserService) GetUserByEmail(email string) (*User, error) {
    var user User
    result := u.db.Where("email = ?", email).First(&user)
    if result.Error != nil {
        return nil, result.Error
    }

    return &user, nil
}

func (u *UserService) SoftDeleteUser(id uint) error {
    var user User
    result := u.db.First(&user, id)
    if result.Error != nil {
        return result.Error
    }

    user.Deleted = true

    result = u.db.Save(&user)
    if result.Error != nil {
        return result.Error
    }

    fmt.Println("User soft deleted successfully!")
    return nil
}

func (u *UserService) ListUsers() ([]User, error) {
    var users []User
    result := u.db.Find(&users)
    if result.Error != nil {
        return nil, result.Error
    }

    return users, nil
}

