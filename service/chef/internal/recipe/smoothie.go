package recipe

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	logging "app/internal/log"
)

func Smoothie(c *gin.Context) {
	name := "smoothie"
	ctx := c.Request.Context()
	logger := logging.GetLoggerWithTraceID(ctx)

	func() {
		ctx, span := tracer.Start(ctx, "食材の準備")
		defer span.End()
		logger := logging.WithTrace(ctx, logger)
		useMaterials(logger, "ブルーベリー", "", "適量", "")
		useMaterials(logger, "牛乳", "", "適量", "")
		useMaterials(logger, "ヨーグルト", "", "適量", "")

		time.Sleep(time.Duration(100) * time.Millisecond)
		transitionTo(logger, "食材の準備OK", "")
	}()

	func() {
		ctx, span := tracer.Start(ctx, "ミキサーで混ぜる")
		defer span.End()
		logger := logging.WithTrace(ctx, logger)
		doCookAction(logger, "", "混ぜる", "何も考えるな！")

		time.Sleep(time.Duration(480) * time.Millisecond)
		transitionTo(logger, "混ぜおわる", "")
	}()

	func() {
		ctx, span := tracer.Start(ctx, "仕上げ")
		defer span.End()
		logger := logging.WithTrace(ctx, logger)
		doCookAction(logger, "", "コップに移す", "")
		time.Sleep(time.Duration(30) * time.Millisecond)
	}()

	c.String(http.StatusOK, name+"完成")
}
