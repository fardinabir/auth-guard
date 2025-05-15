package users

import "github.com/fardinabir/auth-guard/model"

type UserStore interface {
	Create(u *model.User) error
	Update(u *model.User) error
	UpdateById(id int, u *model.User) (*model.User, error)
	Delete(id int) (*model.User, error)
	GetUsers(q map[string]interface{}) ([]model.UserDetails, error)
	GetUserById(id int) (*model.UserDetails, error)
	GetUserByUserName(name string) (*model.User, error)
}
