package recovery

import (
	"github.com/getsentry/sentry-go"

	"github.com/DeAI-Artist/MintAI/core/logger"
)

func ReportPanics(fn func()) {
	defer func() {
		if err := recover(); err != nil {
			sentry.CurrentHub().Recover(err)
			sentry.Flush(logger.SentryFlushDeadline)

			panic(err)
		}
	}()
	fn()
}

func WrapRecover(lggr logger.Logger, fn func()) {
	defer func() {
		if err := recover(); err != nil {
			lggr.Recover(err)
		}
	}()
	fn()
}

func WrapRecoverHandle(lggr logger.Logger, fn func(), onPanic func(interface{})) {
	defer func() {
		if err := recover(); err != nil {
			lggr.Recover(err)

			if onPanic != nil {
				onPanic(err)
			}
		}
	}()
	fn()
}
