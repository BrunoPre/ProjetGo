package storage

import (
	"Project/pkg/mqtt/structs"
	"context"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/go-redis/redis/v8"
)

// Interface de DAO pour écrire des données envoyées par des capteurs (CSV, Redis ou mocking)
type SensorDao interface {
	Write(structs.SensorData) error
	//WriteCSV(structs.SensorData) error
}

// Permet de vérifier simplement que l'interface est bien applicable à RedisSensorDao
var _ SensorDao = (*RedisSensorDao)(nil)

// Implémentation du DAO pour une base de données Redis
type RedisSensorDao struct {
	client *redis.Client
}

func exist(pathFile string) (existe bool) {

	if _, err := os.Stat(pathFile); err == nil {
		existe = true

	} else if errors.Is(err, os.ErrNotExist) {
		// path/to/whatever does *not* exist
		existe = false
	}
	return
}

func WriteCSV(data structs.SensorData) error {

	AirportId := strconv.Itoa(data.AirportId)
	date := data.Timestamp.Format("2006-01-02")

	entry := []string{strconv.Itoa(data.Id), AirportId, string(data.Measure), fmt.Sprintf("%f", data.Value), date}

	path := "/ressources/" + AirportId + "_" + date + string(data.Measure) + ".csv"
	fields := []string{"id", "AirportId", "Measure", "Value", "date"}

	wrapped_entry := [][]string{entry}
	wrapped_fields := [][]string{fields}
	println("PATH : ", exist(path))

	if exist(path) {
		// Open the existing file
		csvFile, err := os.Open(path)
		if err != nil {
			log.Fatalf("failed opening file: %s", err)
		}
		// Read the existing file
		r := csv.NewReader(csvFile)
		lines, err := r.ReadAll()
		if err != nil {
			log.Fatalf("failed reading file: %s", err)
		}
		if err = csvFile.Close(); err != nil {
			log.Fatalf("failed closing file: %s", err)
		}
		wrapped_entry = append(lines, wrapped_entry...)
	} else if !exist(path) {
		wrapped_entry = append(wrapped_fields, wrapped_entry...)
	}

	csvFileWritter, err := os.Create(path)
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	csvwriter := csv.NewWriter(csvFileWritter)
	for _, empRow := range wrapped_entry {
		//fmt.Println(empRow, "\n")
		_ = csvwriter.Write(empRow)
	}

	csvwriter.Flush()

	// Close the file
	if err = csvFileWritter.Close(); err != nil {
		log.Fatalf("failed closing file: %s", err)
	}

	return nil
}

func (r RedisSensorDao) Write(data structs.SensorData) error {
	ctx := context.Background()
	json, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("Error while unmarshalling the data, %w", err)
	}
	WriteCSV(data)
	fmt.Println(r.client.Set(ctx, strconv.Itoa(data.Id), json, 0))
	return nil
}
