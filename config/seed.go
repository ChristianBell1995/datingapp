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
		Name:     "Lael Hines",
		Email:    "lael@hines.com",
		Password: "password",
		Age:      20,
		Gender:   "Female",
		ImageUrl: "https://media.istockphoto.com/id/1357723739/photo/studio-portrait-of-a-smiling-young-latin-woman.jpg?b=1&s=170667a&w=0&k=20&c=RIMvJI9S1mZytKJydukxUF4hRoyVbR1W3ix6gsdo72I=",
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
		Name:     "Archie Thomson",
		Email:    "elton@gmail.com",
		Password: "password",
		Age:      20,
		Gender:   "Male",
		ImageUrl: "https://images.squarespace-cdn.com/content/v1/51efe10de4b025019c37bb06/1427820821815-BZET5XTPRG0F2HYBE95M/white-background-3373-Edit.jpg?format=1000w",
	},
	{
		Name:     "Harriet Ling",
		Email:    "harriet@ling.com",
		Password: "password",
		Age:      23,
		Gender:   "Female",
		ImageUrl: "https://static.independent.co.uk/s3fs-public/thumbnails/image/2011/02/10/23/553733.jpg?quality=75&width=1200&auto=webp",
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
		_, err := users[i].SaveUser(db)
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}
	fmt.Println("Successfully seeded users table")
}
