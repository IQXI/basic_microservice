package main

import (
	cfg "basic_microservice/internal/config"
	lg "basic_microservice/internal/logger"
	"go.uber.org/zap"
)

var logger *zap.Logger
var config map[string]interface{}

func init() {
	config = cfg.GetConfig()
	logger = lg.GetLogger(config)

}

func main() {
	logger.Info("Info msg")
	logger.Fatal("Fatal msg")
}
