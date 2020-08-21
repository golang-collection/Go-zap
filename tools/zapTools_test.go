package tools

import (
	"errors"
	"go.uber.org/zap"
	"testing"
)

/**
* @Author: super
* @Date: 2020-08-21 10:19
* @Description:
**/

func TestGetLogger(t *testing.T) {
	logger := GetLogger()
	logger.Info("hello")
	logger.Error("hello", zap.Error(errors.New("hello error")))
	logger.Debug("hello", zap.Error(errors.New("hello error")))
	logger.Warn("hello")
}