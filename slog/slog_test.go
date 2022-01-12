package slog

import (
	"testing"

	"go.uber.org/zap"
)

var (
	logger *zap.Logger
)

func Test_ValidateSLog(t *testing.T) {
	cfg := Conf{}.DefaultConf()
	logger = NewLogger(&cfg, `app`)
	logger = logger.With(zap.Any(`service`, "服务a"))

	logger.Info("读取配置参数测试", zap.Bool("是否开启调试模式", cfg.Debug))

}
