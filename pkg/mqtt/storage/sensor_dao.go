package storage

import (
	"Project/pkg/mqtt/configuration"
	"Project/pkg/mqtt/structs"
	"context"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

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

type CsvSensorDAO struct {
	dirPath string
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
func (c CsvSensorDAO) Write(data structs.SensorData) error {
	if !exist(c.dirPath) {
		err := os.MkdirAll(c.dirPath, os.ModePerm)
		if err != nil {
			return err
		}
	}

	AirportId := data.AirportId
	date := data.Timestamp.Format("2006-01-02")

	// Convert data to a List
	entry := []string{strconv.Itoa(data.Id), AirportId, string(data.Measure), fmt.Sprintf("%f", data.Value), date}

	// Path's creation for the .csv
	path := "./" + c.dirPath + "/" + AirportId + "_" + date + string(data.Measure) + ".csv"
	fields := []string{"id", "AirportId", "Measure", "Value", "date"}

	// Make a list of list
	wrapped_entry := [][]string{entry}
	wrapped_fields := [][]string{fields}

	// Get the previous data if the file exist
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
		//Add the fields' to the .csv
		wrapped_entry = append(lines, wrapped_entry...)
	} else if !exist(path) {
		//Add the fields' to the .csv
		wrapped_entry = append(wrapped_fields, wrapped_entry...)
	}

	// Create the .csv
	csvFileWriter, err := os.Create(path)
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	// Write in the .csv with all the data
	csvWriter := csv.NewWriter(csvFileWriter)
	for _, empRow := range wrapped_entry {
		_ = csvWriter.Write(empRow)
	}

	csvWriter.Flush()

	// Close the file
	if err = csvFileWriter.Close(); err != nil {
		log.Fatalf("failed closing file: %s", err)
	}
	return nil
}

func (r RedisSensorDao) Write(data structs.SensorData) error {
	ctx := context.Background()
	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("Error while unmarshalling the data, %w", err)
	}
	fmt.Println(r.client.Set(ctx, fmt.Sprintf("%d_%s_%s", data.Id, data.AirportId, strings.Replace(data.Timestamp.Format(time.Stamp), " ", "_", -1)), jsonData, 0))
	return nil
}

func FactorySensorDao(config configuration.DaoConfig) (SensorDao, error) {
	switch config.DaoType {
	case "csv":
		if len(config.FilePath) == 0 {
			return nil, fmt.Errorf("%w : empty filePath", ErrInvalidConfig)
		}
		return CsvSensorDAO{config.FilePath}, nil
	case "redis":
		if config.RedisDb < 0 {
			return nil, fmt.Errorf("%w : negative RedisDb '%d'", ErrInvalidConfig, config.RedisDb)
		}
		if len(config.RedisAddr) == 0 {
			return nil, fmt.Errorf("%w : empty RedisAdrr", ErrInvalidConfig)
		}
		return RedisSensorDao{redis.NewClient(&redis.Options{
			Addr:     config.RedisAddr,
			Password: config.RedisPwd, // no password set
			DB:       config.RedisDb,  // use default DB
		})}, nil
	default:
		return nil, fmt.Errorf("%w : invalid daoType '%s'", ErrInvalidConfig, config.DaoType)
	}
}
