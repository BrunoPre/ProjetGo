# Airport Project

Project réalisé pour un cours d'architectures distribuées par Justine F, Bruno P, Paul V et Aymeric L

------

## Ressources

[Site de ressources du professeur](https://www.laurent-guerin.fr/golang)

------

## Données enregistrées en base

```json
{
    "idAirport": "string",
    "idSensor": "integer",
    "value": "float",
    "measure": "string",
    "date": "string"
}
```

------

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

------

## Pipeline en détail

### Sensor publishers `cmd/mqtt_pub/publisher.go`
1. JSON config file : configuration du sensor & configuration de la connexion au MQTT Broker 
2. Connexion via MQTT client au topic `airport/{airportID}` --> MQTT broker
#### Data stream
1. Génération d'une donnée via des fonctions trigonométriques qui dépendent du temps actuel (`time.Now()`)
2. Une valeur réelle `value` est créée
3. Un SensorData est créé, de champs `{sensorID, airportID, measureType, value, timestamp}` 
4. Celui-ci est _marshalled_, puis
5. il est envoyé au MQTT broker
6. Le processus est réitéré toutes les 10 minutes

### Subscribing
1. JSON config file : configuration du sub (communiquant avec CSV ou Redis) & configuration de la connexion au MQTT Broker 
2. `FactoryControllerDAO` --> `SensorController` REDIS ou CSV
3. connexion via MQTT client à tous les topics `airport/#` --> MQTT broker
#### Sensor Controller DAO
Interface permettant d'écrire les données reçues (les `SensorData`) via la méthode `Write`. Implémentation faite selon le stockage adopté : CSV ou BDD REDIS.

### Sub REDIS -> DB REDIS
Le DAO est doté d'un client REDIS qui se connecte à la BDD via l'adresse et le mot de passe fourni par le fichier de configuration.
Chaque appel à l'implémentation de la méthode `Write` écrit le `SensorData` (valeur) dans la BDD avec un ID unique (clé) de la forme `{sensorID}_{airportID}_{currentTime}`.
Ainsi la fin de la pipeline est :
MQTT broker ---> sub REDIS -- `REDIS_client.Set` `(key, value)` --> BDD REDIS

### DB REDIS -> API
1. Connexion à la BDD
2. Requête HTTP sur l'API REST ---> route ---> handler ---> appel de la BDD
3. Réponse de la BDD ---> marshalling ---> réponse à la requête en JSON

### Pipeline résumée
- JSON config par capteur --> Pub_capteur_aéroport -- random data --> MQTT broker <-- sub REDIS ou CSV <-- JSON config par sub
- sub REDIS --> BDD --> API
- sub CSV --> CSV datalake