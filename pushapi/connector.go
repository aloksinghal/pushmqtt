package pushapi
import (
	MQTT "git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git"
	"log"
	"encoding/json"
	"os"
	"fmt"
	)

type Configuration struct {
	LogFile         string
	Broker          string
	Client          string
}

var LogConfig Configuration

var ServerLogger *log.Logger

var client *MQTT.Client 

func init() {
	configFile, _ := os.Open("config.json")
	err := json.NewDecoder(configFile).Decode(&LogConfig)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(LogConfig)
	file, err := os.OpenFile(LogConfig.LogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
    	fmt.Println("Failed to open log file", file, ":", err)
	}

	logger := log.New(file,
    	"INFO: ",
    	log.Ldate|log.Ltime|log.Lshortfile)
	ServerLogger = logger

	 c := CreateClient(LogConfig.Broker, LogConfig.Client)
	 c.Connect()
	 client = c
}


func CreateClient(brokerAddress string, clientName string) *MQTT.Client {
	opts := MQTT.NewClientOptions().AddBroker(brokerAddress).SetClientID(clientName)
	opts.SetCleanSession(false)
	opts.SetOnConnectHandler(onconnecthandler)
	opts.SetConnectionLostHandler(onconnectionlosthandler)
	c := MQTT.NewClient(opts)
	return c
}

func onconnectionlosthandler(c *MQTT.Client, err error) {
	ServerLogger.Print("lost connection, trying to reconnect.....")
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		ServerLogger.Fatal(token.Error())
	}
	ServerLogger.Print("Reconnected successfully")
}

func onconnecthandler(c *MQTT.Client) {
	ServerLogger.Println("on connnect called")
	ServerLogger.Println("client subscribed")
}

