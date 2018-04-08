package postgres

import (
	"github.com/jinzhu/gorm"
	"github.com/rpoletaev/tchat"
	"golang.org/x/crypto/bcrypt"
)

// AuthService implementation of AuthService interface
type AuthService struct {
	DB *gorm.DB
}

// Register implementation of AuthService interface
func (s *AuthService) Register(acc *tchat.Account) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(acc.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	acc.Password = string(hash)
	return s.DB.Create(acc).Error
}

// Auth implementation of AuthService interface
func (s *AuthService) Auth(acc *tchat.Account) error {
	original := tchat.Account{}
	if err := s.DB.First(&original, "login = ?", acc.Login).Error; err != nil {
		return err
	}

	return bcrypt.CompareHashAndPassword([]byte(original.Password), []byte(acc.Password))
}

// Refresh implementation of AuthService interface
func (s *AuthService) Refresh(acc *tchat.Account) error {

	return nil
}
