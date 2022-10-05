package bootstrap

import (
	"fmt"
	"krisboy-web-api/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDatabase(driver driverSupabase) (*gorm.DB, error) {
	if db == nil {
		dsn := fmt.Sprintf("user=%s "+
			"password=%s "+
			"host=%s "+
			"TimeZone=Asia/Singapore "+
			"port=%s "+
			"dbname=%s",
			driver.User, driver.Password, driver.Host, driver.Port, driver.DbName)
		db_lokal, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			return nil, err
		}
		if err := db_lokal.AutoMigrate(new(domain.Admin)); err != nil {
			return nil, err
		}
		db = db_lokal
	}
	return db, nil
}
