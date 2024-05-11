package bd

import (
	"fmt"
	"gambituser/models"
	"gambituser/tools"
)

func SignUp(sig models.SignUp) error {
	fmt.Println("Comienza registro")

	err := DbConnect()
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer Db.Close()

	/*	sentencia := fmt.Sprintf(`INSERT INTO users(User_Email,User_UUID,User_DateAdd) VALUES ('%v','%v','%v')`,
		sig.UserEmail,
		sig.UserUUID,
		tools.FechaMySQL())
	fmt.Println(sentencia)*/

	//sentencia := "INSERT INTO users (User_Email, User_UUID, User_DateAdd) VALUES ('" + sig.UserEmail + "', '" + sig.UserUUID + "', '" + tools.FechaMySQL() + "')"
	//_, err = Db.Exec(sentencia)

	query := `INSERT INTO users(User_Email,User_UUID,User_DateAdd) VALUES (?,?,?)`
	_, err = Db.Exec(query, sig.UserEmail, sig.UserUUID, tools.FechaMySQL())

	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("User registered successfully")
	return nil

}
