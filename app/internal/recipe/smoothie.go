package recipe

import (
	"net/http"

	"github.com/gin-gonic/gin"

	logging "app/internal/log"
	tracing "app/internal/trace"
	"app/internal/utils"
)

func Smoothie(c *gin.Context) {
	name := "smoothie"
	ctx := c.Request.Context()
	logger := logging.GetLoggerWithTraceID(ctx)

	tracing.CreateTrace(ctx, 100, "食材の準備", logger)
	tracing.CreateTrace(ctx, 200, "果物を潰す", logger)
	tracing.CreateTrace(ctx, 50, "牛乳に入れる", logger)

	chefResponse := utils.SendRequest(ctx, utils.ChefServiceURL, name)
	bbbResponse := utils.SendRequest(ctx, utils.BBBServiceURL, name)

	c.String(http.StatusOK, "約 "+string(chefResponse)+" 分で完成します。BBB流評価は星"+string(bbbResponse)+"です。")
}
