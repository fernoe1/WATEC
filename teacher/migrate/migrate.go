package migrate

import (
	"log"

	"github.com/fernoe1/WATEC/teacher/internal/domain"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(&domain.Teacher{}, &domain.Free{})
	if err != nil {
		log.Fatal(err)
	}
}
