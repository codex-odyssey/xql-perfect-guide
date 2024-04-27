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

func Curry(c *gin.Context) {
	name := "curry"
	ctx := c.Request.Context()
	logger := logging.GetLoggerWithTraceID(ctx)

	tracing.CreateTrace(ctx, 100, "食材の準備", logger)
	tracing.CreateTrace(ctx, 30, "玉ねぎを炒める", logger)
	tracing.CreateTrace(ctx, 50, "肉を加えて炒める", logger)
	tracing.CreateTrace(ctx, 200, "水を加えて煮る", logger)
	tracing.CreateTrace(ctx, 50, "カレールーを加える", logger)
	tracing.CreateTrace(ctx, 2000, "煮込む", logger)

	now := time.Now()

	bbbResponse := utils.SendRequest(ctx, utils.GetBBBProdURL(name))

	delta := time.Since(now).Microseconds()
	logger.With(
		"調理時間", fmt.Sprintf("%dm", delta),
		"BBB流評価", bbbResponse,
	).Info("情報収集完了")

	c.String(http.StatusOK, "完成します。BB流評価は星"+string(bbbResponse)+"です。")
}
