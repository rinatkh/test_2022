package main

import (
	"github.com/rinatkh/test_2022/config"
	"github.com/rinatkh/test_2022/internal/httpServer"
	"github.com/rinatkh/test_2022/internal/httpServer/httpErrorHandler"
	"github.com/sirupsen/logrus"
	"log"
	"time"
)

func main() {

	v, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Cannot load config: ", err.Error())
	}
	cfg, err := config.ParseConfig(v)
	if err != nil {
		log.Fatal("Config parse error", err.Error())
	}

	log := logrus.New()
	formatter := logrus.JSONFormatter{
		TimestampFormat: time.RFC3339,
	}
	switch cfg.Logger.Level {
	case "warning":
		log.SetLevel(logrus.WarnLevel)
	case "notice":
		log.SetLevel(logrus.InfoLevel)
	case "debug":
		log.SetLevel(logrus.DebugLevel)
		formatter.PrettyPrint = true
	default:
		log.SetLevel(logrus.InfoLevel)
	}
	log.SetFormatter(&formatter)
	log.Infof("log level: %s", log.Level.String())

	log.Println("Config loaded")

	errorHandler := httpErrorHandler.NewErrorHandler(cfg)
	s := httpServer.NewServer(cfg, errorHandler, logrus.NewEntry(log))
	if err = s.Run(); err != nil {
		log.Fatalln(err)
	}
}
