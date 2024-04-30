package recipe

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	logging "app/internal/log"
)

func Yakiniku(c *gin.Context) {
	name := "yakiniku"
	ctx := c.Request.Context()
	logger := logging.GetLoggerWithTraceID(ctx)

	func() {
		ctx, span := tracer.Start(ctx, "食材の準備")
		defer span.End()
		logger := logging.WithTrace(ctx, logger)
		useMaterials(logger, "牛肉", "", "適量", "ホルモンとかあるとうまい")
		useMaterials(logger, "豚肉", "", "適量", "")
		useMaterials(logger, "鶏肉", "", "適量", "")

		time.Sleep(time.Duration(50) * time.Millisecond)
		transitionTo(logger, "食材の準備完了", "")
	}()

	func() {
		ctx, span := tracer.Start(ctx, "焼く")
		defer span.End()
		logger := logging.WithTrace(ctx, logger)
		doCookAction(logger, "肉", "焼く", "お腹壊さないように焼いてから食えよ")

		time.Sleep(time.Duration(300) * time.Millisecond)
	}()

	c.String(http.StatusOK, name+"完成")
}
