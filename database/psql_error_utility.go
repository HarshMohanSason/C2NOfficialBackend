package database

import (
	"errors"
	"github.com/lib/pq"
)

func ReturnPsqlError(pqErr *pq.Error) error {
	switch pqErr.Code.Name() {
	case "unique_violation":
		switch pqErr.Constraint {
		case "users_email_key":
			return errors.New("an account with this email already exists, please choose another email")
		case "categories_name_key":
			return errors.New("category name already exists, please choose another name")
		case "products_slug_key":
			return errors.New("slug already exists, please choose another one")
		default:
			return errors.New("duplicate value, please try again with different data")
		}
	case "not_null_violation":
		return errors.New("all fields are required")
	case "check_violation":
		return errors.New("invalid data, please check the input values")
	case "too_many_connections":
		return errors.New("busy server, please try again later")
	case "string_data_right_truncation":
		return errors.New("input too long for a field, try shorter input")
	default:
		return errors.New("database error. Please try again later")
	}
}
