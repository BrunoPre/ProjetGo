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

## Pipeline en détail
### Sensor publishers `cmd/mqtt_pub/publisher.go`
JSON config file (configuration du sensor & configuration de la connexion au MQTT Broker) -- connexion via MQTT client at a topic `airport/{airportID}` --> MQTT broker
#### Data stream
Data generation (math trigonometrics respect of current time) --> `value` (float) --> SensorData = {sensor ID, airport ID, Measure type, `value`, timestamp} -- marshalled and sent every 10s --> MQTT broker

### Subscribing
JSON config file (configuration du sub (vers CSV ou Redis) & configuration de la connexion au MQTT Broker) --> FactoryControllerDAO --> sub REDIS ou CSV
-- connexion via MQTT client à tous les topics `airport/#` --> MQTT broker
#### Sensor Controller DAO
Interface permettant d'écrire les données reçues (les `SensorData`) via la méthode `Write`. Implémentation faite selon le stockage adopté : CSV ou BDD REDIS.
### Sub REDIS -> DB REDIS
Le DAO est doté d'un client REDIS qui se connecte à la BD via l'adresse et le mot de passe fourni par le fichier de config.
Chaque appel à l'implémentation de la méthode `Write` écrit le `SensorData` (valeur) dans la BDD avec un ID unique (clé) de la forme {sensorID}\_{airportID}\_{currentTime}.
Ainsi la suite de la pipeline devient :
MQTT broker ---> sub REDIS -- `REDIS_client.Set` (key, value) --> BDD REDIS

### DB REDIS -> API
Connexion à la DB.
Requête HTTP sur l'API REST ---> route ---> handler ---> appel de la BDD
Réponse de la BDD ---> marshalling ---> réponse à la requête en JSON

### Pipeline résumée
JSON config par capteur --> Pub_capteur_aéroport -- random data --> MQTT broker <-- sub REDIS ou CSV <-- JSON config par sub

sub REDIS --> BDD --> API
sub CSV --> CSV datalake

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