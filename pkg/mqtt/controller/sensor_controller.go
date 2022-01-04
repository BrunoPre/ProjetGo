package controller

import "Project/pkg/mqtt/storage"

// Structure d'un contrôleur utilisant un DAO pour y écrire les données.
// Le contrôleur à pour but de vérifier la validité des données
type SensorController struct {
	dao storage.SensorDao
}
