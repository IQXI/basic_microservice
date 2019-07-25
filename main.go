package main

import (
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
)

var logger *zap.Logger

func init() {

	viper.SetConfigName("config")    // name of config file (without extension)
	viper.AddConfigPath("./Configs") // path to look for the config file in
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	var outputs []string
	logg := viper.Get("logger")
	log_map := logg.(map[string]interface{})
	for _, t := range log_map["outputs"].([]interface{}) {
		outputs = append(outputs, t.(string))
	}

	var level zapcore.Level
	switch log_map["level"] {
	case "DEBUG":
		level = zapcore.DebugLevel
	case "INFO":
		level = zapcore.InfoLevel
	case "ERROR":
		level = zapcore.ErrorLevel
	default:
		level = zapcore.DebugLevel
	}

	fmt.Printf("Outputs: %v\nLevel: %v\nAll: %v\n", outputs, level, logg)

	cfg := zap.Config{
		Encoding:    "json",
		Level:       zap.NewAtomicLevelAt(level),
		OutputPaths: outputs,
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "message",

			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,

			TimeKey:    "time",
			EncodeTime: zapcore.ISO8601TimeEncoder,

			CallerKey:    "caller",
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	logger, err = cfg.Build()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	logger.Info("Info msg")
	logger.Fatal("Fatal msg")
}
