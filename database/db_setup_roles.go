package database

import (
	"c2nofficialsitebackend/models"
	"database/sql"
)

// SetUserRole is only called once when user either signs up or logs in initially.
func SetUserRole(db *sql.DB, user *models.User) error {
	if user != nil && user.IsAdmin {
		_, err := db.Exec(`SET ROLE TO c2n_admin`)
		if err != nil {
			return err
		}
	} else {
		_, err := db.Exec(`SET ROLE TO c2n_user`)
		if err != nil {
			return err
		}
	}
	return nil
}
