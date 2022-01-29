package retry

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/go-toolkit/slog"
	"github.com/go-toolkit/utils/implement"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

var logger *zap.Logger

func TestMain(t *testing.T) {
	cfg := slog.Conf{}.DefaultConf()
	logger = slog.NewLogger(&cfg, `test`)
}

func TestQuickRetry(t *testing.T) {
	i := 0
	j := 0
	fun := func() error {
		i++
		return errors.New("测试")
	}
	retryExtra := func() {
		j++
	}

	require.Error(t, QuickRetry(fun, func(err error) bool {
		return true
	}, retryExtra, time.Second*10, time.Second))
	t.Log(i, j)
}

func TestRetry(t *testing.T) {
	times := 0
	err := Retry(time.Second, time.Millisecond*100, logger, func() error {
		times++
		return fmt.Errorf("%d", times)
	})

	require.Error(t, err)
	require.EqualValues(t, 10, times)
}

func TestRetryAllTheTime(t *testing.T) {
	times := 0

	implement.EnsureGo(logger, func() {
		require.NoError(t, RetryAllTheTime(time.Millisecond*100, logger, func() error {
			times++
			t.Log(times)
			return fmt.Errorf("%d", times)
		}))
	})

	time.Sleep(time.Second * 2)
	require.EqualValues(t, 20, times)
}

// TestRetryGeneric 测试RetryGeneric,实际只测试了次数限制
func TestRetryGeneric(t *testing.T) {
	times := uint(10)

	err := RetryGeneric(NewTimesJudge(times), time.Millisecond*100, logger, func() error {
		times--
		return fmt.Errorf("%d", times)
	})

	require.Error(t, err)
	require.EqualValues(t, 0, times)
}

func TestLogicCompositeJudge_FinishedAnd(t *testing.T) {
	times := uint(10)

	timesJudge := NewTimesJudge(times)

	timeout := time.Millisecond * 20
	timeoutJudge := NewTimeoutJudge(timeout)

	judge := NewLogicCompositeJudge(LogicAnd, []RetryJudge{timeoutJudge, timesJudge})

	require.False(t, judge.Finished())

	time.Sleep(timeout)

	require.False(t, judge.Finished())

	for i := 0; i < int(times)-2; i++ {
		require.False(t, judge.Finished())
	}

	require.True(t, judge.Finished())
}

// TestLogicCompositeJudge_FinishedOr 测试逻辑组合or
func TestLogicCompositeJudge_FinishedOr(t *testing.T) {
	times := uint(10)

	timesJudge := NewTimesJudge(times)

	timeout := time.Millisecond * 20
	timeoutJudge := NewTimeoutJudge(timeout)

	judge := NewLogicCompositeJudge(LogicOr, []RetryJudge{timeoutJudge, timesJudge})

	require.False(t, judge.Finished())

	time.Sleep(timeout)

	require.True(t, judge.Finished())

	// 到这里测试了超时满足的情况，继续测试次数满足

	start := time.Now()
	timesJudge = NewTimesJudge(times)

	timeout = time.Millisecond * 20
	timeoutJudge = NewTimeoutJudge(timeout)

	judge = NewLogicCompositeJudge(LogicOr, []RetryJudge{timeoutJudge, timesJudge})
	require.False(t, judge.Finished())

	for i := 0; i < int(times)-1; i++ {
		require.False(t, judge.Finished())
	}

	require.True(t, judge.Finished())
	// 确定时间不满足
	require.True(t, time.Now().Before(start.Add(timeout)))
}
