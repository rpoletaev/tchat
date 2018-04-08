package postgres

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/rpoletaev/tchat"
)

// UserService represents a PostgreSQL implementation of myapp.UserService.
type UserService struct {
	DB *gorm.DB
}

// User returns a user for a given id.
func (s *UserService) User(id uint) (*tchat.User, error) {
	var u *tchat.User
	err := s.DB.Model(u).First(u, id).Error
	return u, err
}

// Users implements method of UserService interface
func (s *UserService) Users() ([]*tchat.User, error) {
	var users []*tchat.User
	err := s.DB.Model(&tchat.User{}).Find(&users).Error
	return users, err
}

// Save implements method of UserService interface
func (s *UserService) Save(u *tchat.User) error {
	return s.DB.Save(u).Error
}

// Create implements method of UserService interface
func (s *UserService) Create(u *tchat.User) error {
	return s.DB.Model(&tchat.User{}).Create(u).Error
}

// Delete implements method of UserService interface
func (s *UserService) Delete(u *tchat.User) error {
	return s.DB.Delete(u).Error
}
