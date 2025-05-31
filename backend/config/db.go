package config

import (
	"fmt"
	"log"

	"github.com/skor/agentreferal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DB() *gorm.DB {
	dsn := "host=localhost user=dpkrn password=root dbname=SC_AGENT_CUSTOMER_MAPPING"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	db.AutoMigrate(&models.SC_AGENT{})
	db.AutoMigrate(&models.SC_AGENT_CUSTOMER_MAPPING{})

	fmt.Println("Database connected successfully")

	return db
}
