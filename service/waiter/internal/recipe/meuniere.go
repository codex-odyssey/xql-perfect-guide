package recipe

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	logging "app/internal/log"
	tracing "app/internal/trace"
	"app/internal/utils"
)

func Meuniere(c *gin.Context) {
	name := "meuniere"
	ctx := c.Request.Context()
	logger := logging.GetLoggerWithTraceID(ctx)

	tracing.CreateTrace(ctx, 100, "食材の準備", logger)
	tracing.CreateTrace(ctx, 10, "サーモンに塩コショウする", logger)
	tracing.CreateTrace(ctx, 20, "小麦粉をまぶす", logger)
	tracing.CreateTrace(ctx, 80, "バターで両面を焼く", logger)
	tracing.CreateTrace(ctx, 30, "レモン汁をかける", logger)

	chefResponse := utils.SendRequest(ctx, utils.ChefServiceURL, name)
	bbbResponse := utils.SendRequest(ctx, utils.BBBServiceURL, name)

	logger.With(
		"cooking_time", fmt.Sprintf("%sm", chefResponse),
		"BBs_rating", bbbResponse,
	).Info("情報収集完了")

	c.String(http.StatusOK, "約 "+string(chefResponse)+" 分で完成します。BBB流評価は星"+string(bbbResponse)+"です。")
}
