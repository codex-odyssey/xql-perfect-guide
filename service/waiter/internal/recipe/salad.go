package recipe

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	logging "app/internal/log"
	tracing "app/internal/trace"
	"app/internal/utils"
)

func Salad(c *gin.Context) {
	name := "salad"
	ctx := c.Request.Context()
	logger := logging.GetLoggerWithTraceID(ctx)

	tracing.CreateTrace(ctx, 100, "食材の準備", logger)
	tracing.CreateTrace(ctx, 100, "食材を切る", logger)
	tracing.CreateTrace(ctx, 10, "ドレッシングをかける", logger)

	chefResponse := utils.SendRequest(ctx, utils.ChefServiceURL, name)
	bbbResponse := utils.SendRequest(ctx, utils.BBCorpURL, name)

	logger.With(
		"cooking_time", fmt.Sprintf("%sm", chefResponse),
		"BBs_rating", bbbResponse,
	).Info("情報収集完了")

	c.String(http.StatusOK, "約 "+string(chefResponse)+" 分で完成します。BB流評価は星"+string(bbbResponse)+"です。")
}
