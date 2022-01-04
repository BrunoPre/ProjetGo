package storage

import (
	"fmt"

	"github.com/go-redis/redis/v8"
)

var Dao SensorDao

func Init() {
	Dao = RedisSensorDao{redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})}
	fmt.Println(Dao)
}
