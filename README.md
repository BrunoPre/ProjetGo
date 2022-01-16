# Airport Project

Project réalisé pour un cours d'architectures distribuées par Justine F, Bruno P, Paul V et Aymeric L

## Ressources

[Site de ressources du professeur](https://www.laurent-guerin.fr/golang)

## Données enregistrées en base

```json
{
    "idAirport": "string",  // IATA code
    "idSensor": "integer",
    "value": "float",
    "measure": "string",    // Temperature, Pressure, Wind
    "date": "string"        // timestamp
}
```

## Lancer le projet
Chaque commande doit être lancée à la racine du projet, dans un terminal indépendant des autres commandes. 
1. Initialiser la base de données REDIS :
```shell
docker-compose up
```
2. Lancer le subscriber connecté à la base de données :
```shell
go run cmd/mqtt_sub/subscriber.go configs/config_redis.json
```
3. Lancer celui qui écrit les mesures dans des fichiers CSV :
```shell
go run cmd/mqtt_sub/subscriber.go configs/config_csv.json
```
4. Lancer un premier publisher (capteur de température de l'aéroport de Nantes) :
```shell
go run cmd/mqtt_pub/publisher.go configs/capteur_NTE_temp.json
```
5. Lancer un second publisher (capteur de pression de l'aéroport de Pointe-à-Pitre) :
```shell
go run cmd/mqtt_pub/publisher.go configs/capteur_PTP_pres.json
```
6. Lancer l'API :
```shell
go run pkg/api/main.go
```
7. Découverte de l'API : ouvrir le fichier `pkg/api/openapi.yaml`