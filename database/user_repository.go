package database

import (
	"c2nofficialsitebackend/config"
	"c2nofficialsitebackend/models"
	"database/sql"
	"errors"
	"github.com/lib/pq"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	SearchUser(user *models.User) (*models.User, error)
}

type PostgresUserRepository struct {
	DB *sql.DB
}

func (p *PostgresUserRepository) CreateUser(user *models.User) error {

	setRoleError := SetUserRole(p.DB, user)
	if setRoleError != nil {
		return setRoleError
	}
	query := `INSERT INTO users (name, email, password, auth_type) VALUES ($1, $2, $3, $4)`
	_, err := p.DB.Exec(query, user.Name, user.Email, user.Password, user.AuthType)

	if err != nil {
		//Checking for postgres given errors
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			switch pqErr.Code.Name() {
			case "unique_violation":
				return errors.New("email already exists, please choose another email")
			case "not_null_violation":
				return errors.New("all fields are required")
			default:
				config.LogError(pqErr)
				return errors.New("something went wrong. Please try again")
			}
		}
		config.LogError(err)
		return errors.New("an error occurred creating the account. Please try again later")
	}
	return nil
}

func (p *PostgresUserRepository) SearchUser(user *models.User) (*models.User, error) {

	setRoleError := SetUserRole(p.DB, user)
	if setRoleError != nil {
		return nil, setRoleError
	}
	//Searching the user where the auth type and the email matches the passed user
	query := `SELECT id, name, email, password, auth_type, created_at, updated_at FROM users WHERE email = $1 AND auth_type = $2`

	var foundUser models.User
	//Adding values to the foundUser a user if a user is found
	err := p.DB.QueryRow(query, user.Email, user.AuthType).Scan(
		&foundUser.ID,
		&foundUser.Name,
		&foundUser.Email,
		&foundUser.Password,
		&foundUser.AuthType,
		&foundUser.CreatedAt,
		&foundUser.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// No user found with the given email
			config.LogError(err)
			return nil, errors.New("no user found with this email")
		}
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			switch pqErr.Code.Name() {
			case "not_null_violation":
				return nil, errors.New("all fields are required")
			default:
				config.LogError(pqErr)
				return nil, errors.New("something went wrong. Please try again")
			}
		}
		return nil, err
	}
	return &foundUser, nil
}
