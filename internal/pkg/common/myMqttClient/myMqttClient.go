package myMqttClient

import (
	"log"
	"time"

	myLog "github.com/IMT-Atlantique-FIL-2020-2023/NADA-extended/internal/pkg/common/myLog"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func createClientOptions(brokerURI string, clientId string, username string, pswrd string, connectHandler mqtt.OnConnectHandler) *mqtt.ClientOptions {

	opts := mqtt.NewClientOptions()
	// AddBroker adds a broker URI to the list of brokers to be used.
	// The format should be "scheme://host:port"
	opts.AddBroker(brokerURI)
	// opts.SetPassword(password)
	opts.SetClientID(clientId)
	opts.SetUsername(username)
	opts.SetPassword(pswrd)

	opts.SetAutoReconnect(true)
	opts.SetMaxReconnectInterval(1 * time.Second)
	opts.OnConnect = connectHandler
	/*opts.SetConnectionLostHandler(func(c mqtt.Client, err error) {
		myLog.MyLog(myLog.Get_level_WARNING(), "myMqttClient(!!!!!! mqtt connection lost error: "+err.Error())
	})*/
	opts.SetReconnectingHandler(func(c mqtt.Client, options *mqtt.ClientOptions) {
		myLog.MyLog(myLog.Get_level_INFO(), "myMqttClient(...... mqtt reconnecting ......)")
	})
	return opts

}

func Connect(brokerURI string, clientId string, username string, pswrd string, connectHandler mqtt.OnConnectHandler) mqtt.Client {
	opts := createClientOptions(brokerURI, clientId, username, pswrd, connectHandler)
	client := mqtt.NewClient(opts)
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
		myLog.MyLog(myLog.Get_level_INFO(), "myMqttClient(Trying to connect "+brokerURI+", "+clientId+"...)")
	}
	if err := token.Error(); err != nil {
		log.Fatal(err)
	}
	myLog.MyLog(myLog.Get_level_INFO(), "myMqttClient(connected to "+brokerURI+" as "+clientId+")")
	return client
}
