package main

import (
	"encoding/json"
	"fmt"

	database "github.com/IMT-Atlantique-FIL-2020-2023/NADA-extended/internal/app/nada-transform/database"
	env "github.com/IMT-Atlantique-FIL-2020-2023/NADA-extended/internal/pkg/common/env"
	model "github.com/IMT-Atlantique-FIL-2020-2023/NADA-extended/internal/pkg/common/model"
	myLog "github.com/IMT-Atlantique-FIL-2020-2023/NADA-extended/internal/pkg/common/myLog"
	myMqttClient "github.com/IMT-Atlantique-FIL-2020-2023/NADA-extended/internal/pkg/common/myMqttClient"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func messageHandlerFactory(db influxdb2.Client) mqtt.MessageHandler {
	return func(client mqtt.Client, message mqtt.Message) {
		myLog.MyLog(myLog.Get_level_INFO(), "received message: "+string(message.Payload()))

		// decoding JSON
		var measurement model.Measure
		err := json.Unmarshal(message.Payload(), &measurement)
		if err != nil {
			myLog.MyLog(myLog.Get_level_ERROR(), err.Error())
		}

		database.Insert(db, measurement)
	}

}

func init() {
	env.Init("./", ".nada-transform.env")
	myLog.Init(env.GetEnv("NADA_TRANSFORM_LOG_LEVEL"))
}

func main() {
	myLog.MyLog(myLog.Get_level_INFO(), "main(start)")

	var mqtt_port string = env.GetEnv("NADA_TRANSFORM_MQTT_PORT")
	var mqtt_host string = env.GetEnv("NADA_TRANSFORM_MQTT_HOST")
	var mqtt_client_id string = env.GetEnv("NADA_TRANSFORM_MQTT_CLIENT_ID")
	var topic string = env.GetEnv("NADA_TRANSFORM_MQTT_TOPIC")
	var mqtt_client_name string = env.GetEnv("NADA_TRANSFORM_MQTT_CLIENT_NAME")
	var mqtt_client_paswrd string = env.GetEnv("NADA_TRANSFORM_MQTT_PSWRD")
	db := database.ConnectToDb()
	myMqttClient.Connect(mqtt_host+":"+mqtt_port, mqtt_client_id, mqtt_client_name, mqtt_client_paswrd, func(c mqtt.Client) {
		if token := c.Subscribe(topic, 1, messageHandlerFactory(db)); token.Wait() && token.Error() != nil {
			myLog.MyLog(myLog.Get_level_ERROR(), token.Error().Error())
		}
		myLog.MyLog(myLog.Get_level_INFO(), "main(subscribed to topic \""+topic+"\")")
		myLog.MyLog(myLog.Get_level_INFO(), "main(waiting for pubs... press Enter to stop programm properly)")
	})

	fmt.Scanln()

	myLog.MyLog(myLog.Get_level_INFO(), "main(end)")
}
