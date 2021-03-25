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

func (u *User) UpdateUser() error {
	if err := db.Model(&User{}).Where("id = ?", u.ID).Updates(u).Error; err != nil {
		return err
	}

	return nil
}

func (u *User) DeleteUser() error {
	if err := db.Delete(&User{}, u.ID).Error; err != nil {
		return err
	}

	return nil
}

func (u *User) QueryUser(pageNum int, pageSize int) ([]User, error) {
	var (
		users []User
		err   error
	)

	if pageSize > 0 && pageNum > 0 {
		err = db.Where(u.getMaps()).Find(&users).Offset(pageNum).Limit(pageSize).Error
	} else {
		err = db.Where(u.getMaps()).Find(&users).Error
	}

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u *User) QueryUserById(id int) {
	db.First(u, id)
}

func (u *User) ExistUser() (bool, error) {
	if err := db.Select("id").Where("id = ?", u.ID).First(u).Error; err != nil {
		return false, err
	}

	return true, nil
}

func (u *User) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	if u.ID > 0 {
		maps["ID"] = u.ID
	}

	if u.Name != "" {
		maps["Name"] = u.Name
	}
	if u.Password != "" {
		maps["Password"] = u.Password
	}

	return maps
}
