package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Swipe struct {
	gorm.Model
	SwiperID   uint   `gorm:"size:255;not null" json:"swiperId"`
	Swiper     User   `gorm:"foreignKey:SwiperID"`
	ProfileID  uint   `gorm:"size:255;not null" json:"profileId"`
	Profile    User   `gorm:"foreignKey:ProfileID"`
	Preference string `gorm:"size:100;not null;" json:"preference"`
}

func (s *Swipe) Prepare() {
	s.ID = 0
	s.Preference = html.EscapeString(strings.TrimSpace(s.Preference))
	s.CreatedAt = time.Now()
	s.UpdatedAt = time.Now()
}

func (s *Swipe) Validate(db *gorm.DB) error {
	// check to see if swiper exists
	var swiper = User{}
	if _, err := swiper.FindUserByID(db, s.SwiperID); err != nil {
		return errors.New("Cannot find Swiper user record by ID")
	}
	// check to see if profile exists
	var profile = User{}
	if _, err := profile.FindUserByID(db, s.ProfileID); err != nil {
		return errors.New("Cannot find Profile user record by ID")
	}

	if s.Preference == "YES" || s.Preference == "NO" {
		return nil
	}
	return errors.New("Required Preference (YES/NO)")
}

func (s *Swipe) SaveSwipe(db *gorm.DB) (*Swipe, error) {
	var err error
	err = db.Debug().Model(&Swipe{}).Create(&s).Error
	if err != nil {
		return &Swipe{}, err
	}
	return s, nil
}

func (s *Swipe) IsAMatch(db *gorm.DB) (bool, error) {
	var err error
	var existingSwipe = Swipe{}
	err = db.Debug().Model(&Swipe{}).Where("swiper_id = ? AND profile_id >= ?", s.ProfileID, s.SwiperID).Find(&existingSwipe).Error
	if err != nil {
		return false, err
	}
	return existingSwipe.Preference == "YES", nil
}
