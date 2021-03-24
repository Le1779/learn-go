package database

import "fmt"

type User struct {
	ID       int `gorm:"primary_key;auto_increase'"`
	Name     string
	Password string
}

func (u *User) CreateUser() error {
	fmt.Println(db == nil)
	if err := db.Create(u).Error; err != nil {
		return err
	}

	return nil
}

func (u *User) QueryUserById(id int) {
	db.First(u, id)
}
