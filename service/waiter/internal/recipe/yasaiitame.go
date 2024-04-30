package recipe

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	logging "app/internal/log"
	"app/internal/utils"
)

func Yasaiitame(c *gin.Context) {
	name := "yasaiitame"
	ctx := c.Request.Context()
	logger := logging.GetLoggerWithTraceID(ctx)

	now := time.Now()

	_ = utils.SendRequest(ctx, utils.GetChefServiceURL(name))
	bbbResponse := utils.SendRequest(ctx, utils.GetBBBProdURL(name))

	delta := time.Since(now).Microseconds()
	logger.With(
		"調理時間", fmt.Sprintf("%dm", delta),
		"BBB流評価", bbbResponse,
	).Info("情報収集完了")

	c.String(http.StatusOK, "完成します。BB流評価は星"+string(bbbResponse)+"です。")
}
