package loggerfactory

import (
	"demo-app-go/logger/loggerfactory/logrus"
	"fmt"

	"github.com/spf13/viper"
)

func InitLogging() error {
	if viper.GetString("logger.logger") == "logrus" {
		err := logrus.RegisterLogger()
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("invalid or no logger.logger set in config file")
	}
	return nil
}
