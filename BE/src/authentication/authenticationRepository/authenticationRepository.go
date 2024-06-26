package authenticationRepository

import (
	"BE-shop/models/dto/authenticationDto"
	"BE-shop/src/authentication"
	"database/sql"
	"errors"
)

type authenticationRepository struct {
	db *sql.DB
}

func NewAuthenticationRepository(db *sql.DB) authentication.AuthenticationRepository {
	return authenticationRepository{db}
}

func (a authenticationRepository) RegisterUsers(req authenticationDto.Register) error {
	query := `INSERT INTO
	  users (
		users_id,
		email,
		password,
		roles,
		created_at,
		updated_at
	  )
	VALUES
	  ($1, $2, $3, $4, $5, $6)`

	_, err := a.db.Exec(query, req.UsersID, req.Email, req.Password, req.Roles, req.CreatedAt, req.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (a authenticationRepository) CheckEmailExists(usrEmail string) (bool, error) {
	query := `SELECT
	  EXISTS(
		SELECT
		  1
		FROM
		  users
		WHERE
		  email = $1
	  )`

	var exists bool
	err := a.db.QueryRow(query, usrEmail).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (a authenticationRepository) RetrieveUsers(usrEmail string) (usr authenticationDto.Register, err error) {
	querty := `SELECT
	  users_id,
	  email,
	  password,
	  roles,
	  created_at,
	  updated_at
	FROM
	  users WHERE email = $1`

	var usrData authenticationDto.Register
	row := a.db.QueryRow(querty, usrEmail)
	err = row.Scan(&usrData.UsersID, &usrData.Email, &usrData.Password, &usrData.Roles, &usrData.CreatedAt, &usrData.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return authenticationDto.Register{}, errors.New("01")
		}
		return authenticationDto.Register{}, err
	}

	return usrData, err
}
