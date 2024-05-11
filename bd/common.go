package bd

import (
	"database/sql"
	"fmt"
	"gambituser/models"
	"gambituser/secretm"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

var SecretModel models.SecretRDSJson
var err error
var Db *sql.DB

// ReadSecret obtains the secret that contains the credentials of the database in AWS.
// It calls the GetSecret function from the secretm package to get the secret using the provided secret name.
// If there is an error, it returns that error.
func ReadSecret() error {
	SecretModel, err = secretm.GetSecret(os.Getenv("SecretName"))
	return err
}

// DbConnect establishes a connection to the MySQL database using the provided connection string.
func DbConnect() error {

	Db, err = sql.Open("mysql", ConnStr(SecretModel))
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = Db.Ping()
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Successfully connected to DB!")
	return nil
}

// ConnStr creates a connection string for the MySQL database using the provided credentials.
// It takes a `claves` parameter of type `models.SecretRDSJson` which contains the necessary
// information for establishing the connection.
// The function constructs the connection string by concatenating the username, password,
// host, and database name in the format:
// "{username}:{password}@tcp({host})/{dbname}?allowCleartextPasswords=true"
// It then returns the constructed connection string.
func ConnStr(claves models.SecretRDSJson) string {
	var dbUser, authToken, dbEndpoint, dbName string
	dbUser = claves.Username
	authToken = claves.Password
	dbEndpoint = claves.Host
	dbName = "gambit" // nombre de la base de datos que está en aws RDS
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?allowCleartextPasswords=true", dbUser, authToken, dbEndpoint, dbName)
	return dsn
}
