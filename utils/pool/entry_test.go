package pool

import (
	"os"
	"testing"

	"github.com/go-toolkit/slog"
	"go.uber.org/zap"
)

var testLogger *zap.Logger // 测试的自定义日志器

func TestMain(m *testing.M) {
	cfg := slog.Conf{}.DefaultConf()
	testLogger = slog.NewLogger(&cfg, `test-db`)

	os.Exit(m.Run())
}
