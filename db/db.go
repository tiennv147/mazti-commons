package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //mysql dialects
	"github.com/pkg/errors"
	"github.com/tiennv147/mazti-commons/config"
)

func NewDB(dbCfg *config.Database) (*gorm.DB, error) {
	db, err := gorm.Open("mysql", dbCfg.URL)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to open db URL")
	}
	db.DB().SetMaxOpenConns(dbCfg.MaxActive)
	db.DB().SetMaxIdleConns(dbCfg.MaxIdle)

	db.LogMode(dbCfg.LogMode)

	return db, nil
}
