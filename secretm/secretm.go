package secretm

import (
	"encoding/json"
	"fmt"
	"gambituser/awsgo"
	"gambituser/models"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

// GetSecret obtiene la secret que contiene las credenciales d ela base de datos en AWS
func GetSecret(nombreSecret string) (models.SecretRDSJson, error) {

	var datosSecret models.SecretRDSJson
	fmt.Println(" > Pido Secreto" + nombreSecret)

	svc := secretsmanager.NewFromConfig(awsgo.Cfg)

	clave, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(nombreSecret),
	})
	if err != nil {
		fmt.Println(err.Error())
		return datosSecret, err
	}

	err = json.Unmarshal([]byte(*clave.SecretString), &datosSecret)
	if err != nil {
		fmt.Println(err.Error())
		return datosSecret, err
	}

	return datosSecret, nil

}
