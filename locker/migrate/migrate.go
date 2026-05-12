package migrate

import (
	"log"

	"github.com/fernoe1/WATEC/locker/internal/domain"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(&domain.Locker{})
	if err != nil {
		log.Fatal(err)
	}
}
