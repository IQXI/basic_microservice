package main

import (
	//cfg "basic_microservice/internal/config"
	//lg "basic_microservice/internal/logger"
	cfd "github.com/IQXI/basic_microservice/internal/config"
	lg "github.com/IQXI/basic_microservice/internal/logger"
)

func main() {
	logger := lg.GetLogger(cfg.GetConfig())

	logger.Info("Info msg")
	logger.Fatal("Fatal msg")
}
