package logrus

import (
	"demo-app-go/logger"
	"log/syslog"
	"strings"

	"github.com/sirupsen/logrus"
	logrus_syslog "github.com/sirupsen/logrus/hooks/syslog"
	"github.com/spf13/viper"
)

// RegisterLogger register logger.
//
// Parameters:
// - N/A
//
// Returns:
// - N/A
func RegisterLogger() error {
	//standard configuration
	log := logrus.New()

	fileName := viper.GetString("logger.file-name")
	if fileName != "" {
		file, err := logger.OpenFile(fileName)
		if err == nil {
			log.Out = file
		} else {
			log.Info("Failed to log to file, using default stderr")
		}
	}

	if strings.ToUpper(viper.GetString("logger.log-format")) == "TEXT" {
		log.SetFormatter(&logrus.TextFormatter{})
	} else {
		log.SetFormatter(&logrus.JSONFormatter{})
	}

	syslogHost := viper.GetString("logger.syslog-host")
	if syslogHost != "" {
		hook, err := logrus_syslog.NewSyslogHook("udp", syslogHost, syslog.LOG_INFO, "")
		if err != nil {
			log.Error("Unable to connect to local syslog daemon")
		} else {
			log.AddHook(hook)
		}
	}

	mylogger := log.WithFields(logrus.Fields{
		"application": viper.GetString("logger.app-name"),
	})

	if strings.ToUpper(viper.GetString("logger.log-method")) == "Y" {
		log.SetReportCaller(true)
	}

	if viper.GetString("logger.log-level") == "Trace" {
		log.SetLevel(logrus.TraceLevel)
	} else if viper.GetString("logger.log-level") == "Debug" {
		log.SetLevel(logrus.DebugLevel)
	} else if viper.GetString("logger.log-level") == "Info" {
		log.SetLevel(logrus.InfoLevel)
	} else if viper.GetString("logger.log-level") == "Warning" {
		log.SetLevel(logrus.WarnLevel)
	} else if viper.GetString("logger.log-level") == "Error" {
		log.SetLevel(logrus.ErrorLevel)
	}

	logger.SetLogger(mylogger)
	return nil
}
