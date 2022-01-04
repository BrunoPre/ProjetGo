module subs

go 1.17

replace ClientPaho v1.0.0 => ./../ClientPaho

require ClientPaho v1.0.0

replace sensors v1.0.0 => ./../sensors

require sensors v1.0.0

require (
	github.com/eclipse/paho.mqtt.golang v1.3.5 // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	golang.org/x/net v0.0.0-20200425230154-ff2c4b7c35a0 // indirect
)
