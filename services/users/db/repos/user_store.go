package repos

import (
	"errors"
	"github.com/fardinabir/auth-guard/database"
	"github.com/fardinabir/auth-guard/model"
	"gorm.io/gorm"
	"log"
)

type UserStore struct {
	DB *gorm.DB
}

func NewUserStore() *UserStore {
	d := database.GetDBConnection()
	return &UserStore{DB: d}
}

func (s *UserStore) Create(u *model.User) error {
	res := s.DB.Create(u)
	if res.Error != nil {
		log.Println("Error while creating user in db", res.Error)
		return res.Error
	}
	return nil
}

// TODO: complete update func
func (s *UserStore) Update(u *model.User) error {
	return nil
}

// TODO: 1. user can update password, separate it, use it only for profile update. 2. if id is not found it does not returns error rather success
func (s *UserStore) UpdateById(id int, u *model.User) (*model.User, error) {
	res := s.DB.Where("id = ?", id).Updates(&u)
	if res.Error != nil {
		log.Println("Error while updating user in db", id, res.Error)
		return nil, res.Error
	}
	return u, nil
}

func (s *UserStore) Delete(id int) (*model.User, error) {
	usr := &model.User{}
	res := s.DB.Delete(&usr, id)
	if res.Error != nil {
		log.Println("Error while deleting user in db", id, res.Error)
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		log.Println("Record with ", id, " not found.")
		return nil, errors.New("not Found")
	}
	return usr, nil
}

func (s *UserStore) GetUsers(q map[string]interface{}) ([]model.UserDetails, error) {
	var users []model.UserDetails
	res := s.DB.Model(&model.User{}).Where(q).Find(&users)
	if res.Error != nil {
		log.Println("Fetching Users list, ", q, res.Error)
		return nil, res.Error
	}
	return users, nil
}

func (s *UserStore) GetUserByUserName(name string) (*model.User, error) {
	usr := &model.User{}
	res := s.DB.Where("user_name = ?", name).First(usr)
	if res.Error != nil {
		log.Println("Error while getting user in db", name, res.Error)
		return nil, res.Error
	}
	return usr, nil
}

func (s *UserStore) GetUserById(id int) (*model.UserDetails, error) {
	usr := &model.UserDetails{}
	res := s.DB.Model(&model.User{}).Where("id = ?", id).First(usr)
	if res.Error != nil {
		log.Println("Error while getting user in db", id, res.Error)
		return nil, res.Error
	}
	return usr, nil
}
