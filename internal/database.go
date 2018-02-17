package internal

import (
	"log"
	"time"

	"context"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type QRCode struct {
	gorm.Model
	ID       string
	Metadata string
}

func InsertState(ctx context.Context, metadata string) (string, error) {
	db := DBFromContext(ctx)

	qr := &QRCode{ID: "TODO", Metadata: metadata}
	if err := db.Create(qr).Error; err != nil {
		return "", err
	}

	return "TODO(ID)", nil
}

func InitDatabase(dbDsn string) (*gorm.DB, error) {
	var err error
	var db *gorm.DB

	for i := 1; i < 10; i++ {
		db, err = gorm.Open("postgres", dbDsn)
		if err == nil || i == 10 {
			break
		}
		sleep := (2 << uint(i)) * time.Second
		log.Printf("Could not connect to DB: %v", err)
		log.Printf("Waiting %v before retry", sleep)
		time.Sleep(sleep)
	}
	if err != nil {
		return nil, err
	}
	if err = db.AutoMigrate(&QRCode{}).Error; err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}
