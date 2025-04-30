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

	query := `INSERT INTO public.users (name, email, password, auth_type) VALUES ($1, $2, $3, $4)`
	_, err := p.DB.Exec(query, user.Name, user.Email, user.Password, user.AuthType)

	if err != nil {
		//psql errors
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			config.LogError(pqErr)
			return ReturnPsqlError(pqErr)
		}
		//default error
		config.LogError(err)
		return errors.New("an error occurred creating the account. Please try again later")
	}
	return nil
}

func (p *PostgresUserRepository) SearchUser(user *models.User) (*models.User, error) {

	//Searching the user where the auth type and the email matches the passed user
	query := `SELECT id, name, email, password, auth_type, is_admin, created_at, updated_at FROM users WHERE email = $1 AND auth_type = $2`

	var foundUser models.User
	//Adding values to the foundUser a user if a user is found
	err := p.DB.QueryRow(query, user.Email, user.AuthType).Scan(
		&foundUser.ID,
		&foundUser.Name,
		&foundUser.Email,
		&foundUser.Password,
		&foundUser.AuthType,
		&foundUser.IsAdmin,
		&foundUser.CreatedAt,
		&foundUser.UpdatedAt)

	if err != nil {
		// No user found with the given email
		if errors.Is(err, sql.ErrNoRows) {
			config.LogError(err)
			return nil, errors.New("no user found with this email")
		}
		//psql errors
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			config.LogError(pqErr)
			return nil, ReturnPsqlError(pqErr)
		}
		//default error
		config.LogError(pqErr)
		return nil, errors.New("an error occurred querying the database,try again later")
	}

	return &foundUser, nil
}
