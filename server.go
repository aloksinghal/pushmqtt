package main

import (
	"bitbucket.org/shadowfaxtech/negroni"
	"github.com/gorilla/mux"
	"net/http"
	//"github.com/jingweno/negroni-gorelic"
	//raven "github.com/getsentry/raven-go"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"pushmqtt/pushapi"
)

type Configuration struct {
	LicenseKey string
	AppName    string
	Dsn        string
	LogFile    string
	Broker	   string
	Client     string
}

var Config Configuration

var ServerLogger *log.Logger

var LogFileObj io.Writer



func init() {
	configFile, _ := os.Open("config.json")
	err := json.NewDecoder(configFile).Decode(&Config)
	if err != nil {
		fmt.Println("error:", err)
	}
	file, err := os.OpenFile(Config.LogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file", file, ":", err)
	}
	LogFileObj = file
	logger := log.New(file,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)
	ServerLogger = logger
}

func main() {
	router := mux.NewRouter()
	n := negroni.New()
	/*recovery := negroni.NewSentryRecovery()
	n.Use(recovery)
	recovery.PrintStack = false*/
	n.Use(negroni.NewLogger(LogFileObj))
	//n.Use(negronigorelic.New(ReportingConfig.LicenseKey, ReportingConfig.AppName, false))
	router.HandleFunc("/", pushapi.HandleIndex)
	router.HandleFunc("/publish", pushapi.PublishMessage).Methods("POST")
	n.UseHandler(router)
	ServerLogger.Println("Starting the Server")
	http.ListenAndServe("0.0.0.0:8081", n)

}
