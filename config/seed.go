package config

import (
	"fmt"
	"log"

	"github.com/ChristianBell1995/datingapp/api/models"
	"gorm.io/gorm"
)

var users = []models.User{
	{
		Name:     "Christian Bell",
		Email:    "christiantcbell@gmail.com",
		Password: "password",
		Age:      20,
		Gender:   "Male",
		ImageUrl: "https://media.licdn.com/dms/image/C4D03AQH8Yz5J-Olb8g/profile-displayphoto-shrink_800_800/0/1550912458410?e=1677715200&v=beta&t=ueBQieYpEtpurU5kFTrCeidr2Ace6JbWkQ7VikeiQ3A",
	},
	{
		Name:     "Lady Gaga",
		Email:    "ladyGaga@gmail.com",
		Password: "password",
		Age:      20,
		Gender:   "Female",
		ImageUrl: "https://phantom-marca.unidadeditorial.es/3a39c2e8f7af966e1e826d22c3d54ea4/resize/1320/f/jpg/assets/multimedia/imagenes/2022/08/10/16600904080224.jpg",
	},
	{
		Name:     "Ziggy Stardust",
		Email:    "ziggy@gmail.com",
		Password: "password",
		Age:      20,
		Gender:   "Male",
		ImageUrl: "https://i1.sndcdn.com/artworks-000085309918-lmxzvq-t500x500.jpg",
	},
	{
		Name:     "ELton John",
		Email:    "elton@gmail.com",
		Password: "password",
		Age:      20,
		Gender:   "Male",
		ImageUrl: "https://upload.wikimedia.org/wikipedia/commons/8/80/Elton_John_2011_Shankbone_2_%28cropped%29.JPG",
	},
}

func seed(db *gorm.DB) {

	err := db.Debug().Migrator().DropTable(&models.User{}, &models.Swipe{})
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err.Error())
	}

	err = db.Debug().AutoMigrate(&models.User{}, &models.Swipe{})
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err.Error())
	}

	for i := range users {
		createUserError := db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", createUserError)
		}
	}
	fmt.Println("Successfully seeded users table")
}
