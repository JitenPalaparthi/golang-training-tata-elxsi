package models

import "errors"

type User struct {
	//gorm.Model // Promoted field
	CommonModel

	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Mobile   string `json:"mobile"`
}

// type UserCreated struct{

// }

type CommonModel struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Status      string `json:"status"`
	LastUpdated uint64 `json:"last_updated" gorm:"column:last_updated"`
}

func (u *User) Validate() error {
	if u.Name == "" {
		return errors.New("Invalid user name")
	}
	if u.Email == "" {
		return errors.New("Invalid user email")
	}
	// can implement proper email validation
	// can implement proper mobile number validation etc..
	if u.Mobile == "" {
		return errors.New("Invalid user mobile numbder")
	}
	if u.Password == "" {
		return errors.New("Invalid user password")
	}
	return nil
}
