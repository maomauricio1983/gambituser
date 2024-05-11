package main

import (
	"context"
	"errors"
	"fmt"
	"gambituser/awsgo"
	"gambituser/bd"
	"gambituser/models"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"os"
)

func main() {

	lambda.Start(EjecutoLambda)
}

func EjecutoLambda(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {

	awsgo.InicializoAWS()

	if !ValidoParametros() {
		fmt.Println("Erro no validos parametros")
		err := errors.New("no validos parametros")
		return event, err
	}

	var datos models.SignUp

	for row, att := range event.Request.UserAttributes {
		switch row {
		case "email":
			datos.UserEmail = att
			fmt.Println("Email address: " + datos.UserEmail)

		case "sub":
			datos.UserUUID = att
			fmt.Println("UUID address: " + datos.UserUUID)

		}
	}

	err := bd.ReadSecret()
	if err != nil {
		fmt.Println("Error al leer la secret: ", err.Error())
		return event, err
	}

	err = bd.SignUp(datos)

	return event, err
}

func ValidoParametros() bool {
	var traeParametro bool
	_, traeParametro = os.LookupEnv("SecretName")
	return traeParametro
}
