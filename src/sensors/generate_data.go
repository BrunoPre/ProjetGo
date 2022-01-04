package sensors

import "math/rand"

func GenerateRandomData(sensor string, allRanges map[string][2]int) (int, map[string][2]int) {
	ranges := allRanges[sensor]
	min := ranges[0]
	max := ranges[1]
	res := rand.Intn(max-min+1) + min
	// res <= max --> `res` is only kept
	ranges[1] = res
	allRanges[sensor] = ranges
	return res, allRanges
}

func InitSensorData() map[string][2]int {
	// map sensor -> init data
	initData := make(map[string][2]int) // sensor -> [min, max]
	initData["temp"] = [2]int{0, 30}
	initData["wind"] = [2]int{0, 50}         //kph wind
	initData["pressure"] = [2]int{970, 1030} // avg range of atm pressure (hPa)
	return initData
}
