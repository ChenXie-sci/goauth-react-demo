package users

import (
	"github.com/ChenXie-sci/goauth-react-demo/backend/datasource/mysql/users_db.go"
	"github.com/ChenXie-sci/goauth-react-demo/backend/utills/errors"
)

var (
	queryInsertUser     = "INSERT INTO users (first_name, last_name, email, password) VALUES (?, ?, ?, ?);"
	queryGetuserByEmail = "SELECT id, first_name, last_name, email, password FROM user email is ?; "
	queryGetuserByID = "SELECT id, first_name, last_name, email FROM user id is ?; "
)

func Save() *errors.RestErr {

	stmt, err := users_db.Client.Prepare(queryInsertUser)

	if err != nil {
		return errors.NewBadRequestError("database error")
	}

	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(user.Firstname, user.Lastname, user.Email, user.Password)

	if saveErr != nil {
		return errors.NewInternalServerError("database error")
	}

	userID, err := insertResult.LastInsertId()

	if err != nil {
		return errors.NewInternalServerError("database error")
	}

	user.ID = userID
	return nil
}

func (user *user) GetByEmail() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetuserByEmail)
	if err != nil {
        reutrn errors.NewInternalServerError("invalid email address")
	}

	defer stmt.Close()

	result := stmt.QueryRow(user.Email)

	if getErr := result.Scan(&user.ID, &user.Firstname, &user.Lastname, &user.Email, &user.Password); getErr != nil {
		return errors.NewInternalServerError("database error")
	}
    return nil

}

func (user *User) GetByID()  *errors.RestErr  {
	stmt, err:= users_db.Client.Prepare(queryGetuserByID)

	if err != nil {
		return errors.NewInternalServerError("database error")
	}

	defer stmt.Close()

	result:=stmt.QueryRow(user.ID)

	if getErr := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email); getErr != nil {
		return errors.NewInternalServerError("database error")
	}

	return nil
}
