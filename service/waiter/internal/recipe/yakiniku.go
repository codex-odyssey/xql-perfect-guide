package recipe

import (
	"fmt"
	"net/http"

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
	chefResponse := utils.SendRequest(ctx, utils.ChefServiceURL, name)
	bbbResponse := utils.SendRequest(ctx, utils.BBCorpURL, name)

	logger.With(
		"cooking_time", fmt.Sprintf("%sm", chefResponse),
		"BBs_rating", bbbResponse,
	).Info("情報収集完了")

	c.String(http.StatusOK, "約 "+string(chefResponse)+" 分で完成します。BB流評価は星"+string(bbbResponse)+"です。")
}
