package dal

import "gorm.io/gorm"

var DB *gorm.DB

type Database struct {
	*gorm.DB
}

// NewDatabase
func NewDatabase() *Database {
	return &Database{DB: DB}
}

// Close
func (d *Database) Close() error {
	db, err := d.DB.DB()
	if err != nil {
		return err
	}
	return db.Close()
}
