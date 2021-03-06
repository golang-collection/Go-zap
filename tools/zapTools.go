package tools

import (
	"go.uber.org/zap"
)

/**
* @Author: super
* @Date: 2020-08-21 10:19
* @Description:
**/

var logger *zap.Logger

func init(){
	logger, _ = zap.NewProduction()
	//config := zap.NewDevelopmentConfig()
	//config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	//logger, _ = config.Build()
	defer logger.Sync()
}

func GetLogger() *zap.Logger {
	return logger
}