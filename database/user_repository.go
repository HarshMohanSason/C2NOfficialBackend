package database

import(
	"c2nofficialsitebackend/models"
	"database/sql"
	"c2nofficialsitebackend/utils"
	"errors"
	"github.com/lib/pq" 
)

type UserRepository interface {
    CreateUser(user *models.User) error
}

type PostgresUserRepository struct {
    DB *sql.DB
}

//PostgresUserRepository implements the UserRepository defined interface 
func (p *PostgresUserRepository) CreateUser(user *models.User) error{

	query := "INSERT INTO users (name, email, password, auth_type) VALUES ($1, $2, $3, $4)"

	_, err := p.DB.Exec(query, user.Name, user.Email, user.Password, user.AuthType)

	if err != nil{	
		//Checking for postgres given errors 
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				return errors.New("Email already exists, please choose another email")
			case "not_null_violation":
				return errors.New("All fields are required.")
			default:
				utils.LogError(pqErr)
				return errors.New("Something went wrong. Please try again.")
			}
		}
		utils.LogError(err)
		return errors.New("An error occured creating the account. Please try again later")
	}
	return nil
}