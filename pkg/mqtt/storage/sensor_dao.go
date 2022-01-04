package storage

import (
	"Project/pkg/mqtt/structs"

	"github.com/go-redis/redis/v8"
)

// Interface de DAO pour écrire des données envoyées par des capteurs (CSV, Redis ou mocking)
type SensorDao interface {
	Write(structs.SensorData) error
}

// Permet de vérifier simplement que l'interface est bien applicable à RedisSensorDao
var _ SensorDao = (*RedisSensorDao)(nil)

// Implémentation du DAO pour une base de données Redis
type RedisSensorDao struct {
	client *redis.Client
}

func (r RedisSensorDao) Write(data structs.SensorData) error {
	return nil
}
