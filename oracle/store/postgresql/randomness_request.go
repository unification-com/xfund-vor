package postgresql

import "github.com/jinzhu/gorm"

type RandomnessRequestStore struct {
	db *DB
	tx *gorm.DB
}

func (d RandomnessRequestStore) Migrate() {
	d.db.Exec("")
}

func (d RandomnessRequestStore) TableExists() {
	d.db.HasTable("randomness_request")
}

func (d RandomnessRequestStore) LastTransaction() {
	d.db.Exec("")
}

func (d *RandomnessRequestStore) Insert() (err error) {
	err = d.db.Create("").Error
	return
}
