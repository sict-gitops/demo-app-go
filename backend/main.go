package main

import (
	"demo-app-go/api"
	"demo-app-go/logger"
	"demo-app-go/logger/loggerfactory"
	"flag"
	"log"
	"os"

	"github.com/spf13/viper"
)

const (
	version = "1.0"
)

func main() {

	configFile := flag.String("config-file", "config.yml", "Configuation file name.")
	flag.Parse()

	viper.SetConfigName(*configFile)
	viper.SetConfigType("yml")
	viper.AddConfigPath("./conf/")
	viper.AddConfigPath("./")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Cannot read configuration file: %s", err)
	}

	err := loggerfactory.InitLogging()
	if err != nil {
		log.Fatalf("Cannot initialize logger: %s", err)
	}
	defer logger.Close()

	logger.Log.Infof("Our demo app is starting. Version %s", version)

	pid := os.Getpid()
	logger.Log.Debugf("Process ID: %d", pid)

	router := api.SetupRouter()
	port := viper.GetString("api.port")
	if port == "" {
		port = "80"
	}

	err = router.Run(":" + port)
	if err != nil {
		logger.Log.Fatal("Error starting the HTTP server")
	}
}
