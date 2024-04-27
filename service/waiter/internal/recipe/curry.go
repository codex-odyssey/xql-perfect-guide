package recipe

import (
	"net/http"

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

	chefResponse := utils.SendRequest(ctx, utils.ChefServiceURL, name)
	bbbResponse := utils.SendRequest(ctx, utils.BBBServiceURL, name)

	c.String(http.StatusOK, "約 "+string(chefResponse)+" 分で完成します。BBB流評価は星"+string(bbbResponse)+"です。")
}