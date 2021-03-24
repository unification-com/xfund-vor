package postgresql

import "github.com/jinzhu/gorm"

type RandomnessRequestStore struct {
	db *DB
	tx *gorm.DB
}
