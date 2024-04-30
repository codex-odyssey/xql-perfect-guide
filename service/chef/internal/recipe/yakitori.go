package recipe

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	logging "app/internal/log"
)

func Yakitori(c *gin.Context) {
	name := "yakitori"
	ctx := c.Request.Context()
	logger := logging.GetLoggerWithTraceID(ctx)

	func() {
		ctx, span := tracer.Start(ctx, "食材の準備")
		defer span.End()
		logger := logging.WithTrace(ctx, logger)
		useMaterials(logger, "鶏肉", "", "適量", "一口サイズ")
		useMaterials(logger, "ねぎ", "", "適量", "一口サイズ")

		time.Sleep(time.Duration(100) * time.Millisecond)
		transitionTo(logger, "食材の準備OK", "")
	}()

	func() {
		ctx, span := tracer.Start(ctx, "串に打つ")
		defer span.End()
		logger := logging.WithTrace(ctx, logger)
		doCookAction(logger, "鶏肉、ネギ", "串を打つ", "無心で串うち")

		time.Sleep(time.Duration(480) * time.Millisecond)
		transitionTo(logger, "串打ちが終わる", "")
	}()

	func() {
		ctx, span := tracer.Start(ctx, "焼き鳥を焼く")
		defer span.End()
		logger := logging.WithTrace(ctx, logger)
		doCookAction(logger, "", "塩を振る", "思った3倍振っていこう")
		doCookAction(logger, "", "じっくり焼く", "強火の遠火がコツ")

		time.Sleep(time.Duration(100) * time.Millisecond)
		transitionTo(logger, "肉に火が通る", "")
	}()

	c.String(http.StatusOK, name+"完成")
}
