package configs

import (
	"DENV_Cases_Consumer/models"
)

func GetKafkaProperties() (models.KafkaProperties, error) {
	kfProperties := models.KafkaProperties{
		Host: "192.168.1.13",
		Port: "9092",
	}

	return kfProperties, nil
}
