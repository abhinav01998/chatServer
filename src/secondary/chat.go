package secondary

import (
	"log"

	"github.com/jinzhu/gorm"

	//mysql dialect to be used at time of need
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//User is a struct
type User struct {
	uname   string `gorm:"primary_key"`
	message string
}

//Chatupdate function sends and updates the users/clients chat
func Chatupdate() {
	db, err := gorm.Open("mysql", "user:password@tcp(127.0.0.1:3306)/demoproject?charset=utf8&parseTime=True")
	db.LogMode(true)
	if err != nil {
		log.Println("Database connection unsuccessful !")
		panic(err)
	} else {
		log.Println("Database connection successful !")
		db.Debug().DropTableIfExists(&User{})
		db.Debug().AutoMigrate(&User{})
	}
	db.Close()
}
