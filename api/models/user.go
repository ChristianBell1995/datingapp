package models

import (
	"errors"
	"html"
	"log"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"size:255;not null" json:"name"`
	Email    string `gorm:"size:100;not null;unique" json:"email"`
	Password string `gorm:"size:100;not null;" json:"password,omit"`
	Gender   string `gorm:"size:100;not null;" json:"gender"`
	Age      uint8  `gorm:"size:100;not null;" json:"age"`
	ImageUrl string `json:"imageUrl"`
}

type UserResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Gender    string    `json:"gender"`
	Age       uint8     `json:"age"`
	ImageUrl  string    `json:"imageUrl"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *User) BeforeSave() error {
	hashedPassword, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) Prepare() {
	u.ID = 0
	u.Name = html.EscapeString(strings.TrimSpace(u.Name))
	u.Gender = html.EscapeString(strings.TrimSpace(u.Gender))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "login":
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil

	default:
		if u.Name == "" {
			return errors.New("Required Name")
		}
		if u.Gender == "" {
			return errors.New("Required Gender")
		}
		if u.Age == 0 {
			return errors.New("Required Age")
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil
	}
}

func (u *User) SaveUser(db *gorm.DB) (*UserResponse, error) {

	var err error
	err = u.BeforeSave()
	if err != nil {
		log.Fatal(err)
	}

	err = db.Debug().Create(&u).Error
	if err != nil {
		return &UserResponse{}, err
	}

	return &UserResponse{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		Age:       u.Age,
		Gender:    u.Gender,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}, nil
}

func (u *User) FindAllOtherUsers(db *gorm.DB, userID uint) (*[]UserResponse, error) {
	var err error
	users := []UserResponse{}
	err = db.Debug().Model(&User{}).Limit(100).Not("id = ?", userID).Find(&users).Error
	if err != nil {
		return &[]UserResponse{}, err
	}
	return &users, err
}

func (u *User) ValidateUserPassword(db *gorm.DB, password string) error {
	var err error
	err = db.Debug().Model(User{}).Where("email = ?", u.Email).Take(&u).Error
	if err != nil {
		return err
	}

	if err = VerifyPassword(u.Password, password); err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return err
	}
	return nil
}

func (u *User) FindUserByID(db *gorm.DB, uid uint) (*UserResponse, error) {
	var err error
	err = db.Debug().Model(User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &UserResponse{}, err
	}

	return &UserResponse{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
		Age:   u.Age,
	}, nil
}
