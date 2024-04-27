package recipe

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	logging "app/internal/log"
	tracing "app/internal/trace"
	"app/internal/utils"
)

func Yakiniku(c *gin.Context) {
	name := "yakiniku"
	ctx := c.Request.Context()
	logger := logging.GetLoggerWithTraceID(ctx)

	tracing.CreateTrace(ctx, 30, "食材の準備", logger)
	tracing.CreateTraceWithAttribute(ctx, 30, "肉を切る", logger, map[string]interface{}{
		"deployment.environment": "staging",
	})
	tracing.CreateTraceWithAttribute(ctx, 50, "肉を焼く", logger, map[string]interface{}{
		"deployment.environment": "prod",
	})

	now := time.Now()

	bbbResponse := utils.SendRequest(ctx, utils.GetBBBProdURL(name))

	delta := time.Since(now).Microseconds()
	logger.With(
		"調理時間", fmt.Sprintf("%dm", delta),
		"BBB流評価", bbbResponse,
	).Info("情報収集完了")

	c.String(http.StatusOK, "完成します。BB流評価は星"+string(bbbResponse)+"です。")
}
