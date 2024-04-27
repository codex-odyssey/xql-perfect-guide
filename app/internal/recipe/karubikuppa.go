package recipe

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	logging "app/internal/log"
	tracing "app/internal/trace"
	"app/internal/utils"
)

func Karubikuppa(c *gin.Context) {
	name := "karubikuppa"
	ctx := c.Request.Context()
	logger := logging.GetLoggerWithTraceID(ctx)

	tracing.CreateTrace(ctx, 100, "食材の準備", logger)
	tracing.CreateTrace(ctx, 50, "カルビを炒める", logger)
	tracing.CreateTrace(ctx, 300, "カルビクッパを煮込む", logger)
	tracing.CreateTrace(ctx, 30, "ごま油を入れる", logger)

	chefResponse := utils.SendRequest(ctx, utils.ChefServiceURL, name)
	bbbResponse := utils.SendRequest(ctx, utils.BBBServiceURL, name)

	logger.With(
		"調理時間", fmt.Sprintf("%sm", chefResponse),
		"BBB流評価", bbbResponse,
	).Info("情報収集完了")

	c.String(http.StatusOK, "約 "+string(chefResponse)+" 分で完成します。BBB流評価は星"+string(bbbResponse)+"です。")
}
