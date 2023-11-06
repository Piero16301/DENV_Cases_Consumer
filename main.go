package main

import (
	"DENV_Cases_Consumer/configs"
	"DENV_Cases_Consumer/controllers"
	"DENV_Cases_Consumer/models"
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
)

func main() {
	message, err := configs.Consumer.ConsumePartition("register-vector-record", 0, sarama.OffsetOldest)
	if err != nil {
		panic(err)
	}

	for {
		var vectorRecord models.VectorRecord
		messages := <-message.Messages()
		err = json.Unmarshal(messages.Value, &vectorRecord)
		if err != nil {
			panic(err)
		}

		// Enviar Inspección de Vivienda a Backend
		controllers.SendVectorRecordToBackend(&vectorRecord)

		fmt.Println("Registro de Vector registrado en la Base de Datos")
	}
}
