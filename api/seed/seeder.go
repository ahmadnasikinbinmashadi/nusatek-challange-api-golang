package seed

import (
	b64 "encoding/base64"
	"log"

	"github.com/ahmadnasikinbinmashadi/login-app/api/models"
	"github.com/jinzhu/gorm"
)

var image string = "profile.jpg"
var sEnc string = b64.StdEncoding.EncodeToString([]byte(image))
var users = []models.User{
	models.User{
		Nickname:     "Steven victor",
		PhotoProfile: sEnc,
		Email:        "steven@gmail.com",
		Password:     "password",
	},
	models.User{
		Nickname:     "Martin Luther",
		PhotoProfile: sEnc,
		Email:        "luther@gmail.com",
		Password:     "password",
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}
}
