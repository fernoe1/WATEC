package migrate

import (
	"log"

	"github.com/fernoe1/WATEC/classroom/internal/domain"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(&domain.Classroom{}, &domain.Free{})
	if err != nil {
		log.Fatal(err)
	}
}
