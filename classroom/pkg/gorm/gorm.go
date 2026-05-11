package gorm

import (
	"fmt"
	"log"

	"github.com/fernoe1/WATEC/classroom/config"
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewGormDB(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=postgres port=%s sslmode=%s TimeZone=%s",
		cfg.PostgreSQL.Host, cfg.PostgreSQL.User,
		cfg.PostgreSQL.Password, cfg.PostgreSQL.Port,
		cfg.PostgreSQL.SslMode, cfg.PostgreSQL.Timezone,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Use(otelgorm.NewPlugin()); err != nil {
		log.Fatal(err)
	}

	return db
}
