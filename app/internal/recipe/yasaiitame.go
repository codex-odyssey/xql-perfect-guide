package recipe

import (
	"net/http"

	"github.com/gin-gonic/gin"

	logging "app/internal/log"
	tracing "app/internal/trace"
	"app/internal/utils"
)

func Yasaiitame(c *gin.Context) {
	name := "yasaiitame"
	ctx := c.Request.Context()
	logger := logging.GetLoggerWithTraceID(ctx)

	tracing.CreateTrace(ctx, 30, "食材の準備", logger)
	tracing.CreateTrace(ctx, 50, "野菜を切る", logger)
	tracing.CreateTrace(ctx, 30, "野菜を炒める", logger)

	chefResponse := utils.SendRequest(ctx, utils.ChefServiceURL, name)
	bbbResponse := utils.SendRequest(ctx, utils.BBBServiceURL, name)

	c.String(http.StatusOK, "約 "+string(chefResponse)+" 分で完成します。BBB流評価は星"+string(bbbResponse)+"です。")
}
