package recipe

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	logging "app/internal/log"
)

func Salad(c *gin.Context) {
	name := "salad"
	ctx := c.Request.Context()
	logger := logging.GetLoggerWithTraceID(ctx)

	func() {
		ctx, span := tracer.Start(ctx, "野菜を切る")
		defer span.End()
		logger := logging.WithTrace(ctx, logger)
		useMaterials(logger, "レタス", "", "適量", "手で千切れ")
		useMaterials(logger, "パプリカ", "", "適量", "あると本格的")
		useMaterials(logger, "トマト", "", "適量", "種は抜いておくと水っぽくならない")
		useMaterials(logger, "玉ねぎ", "", "適量", "辛いのが苦手なやつは水に晒しておけ! BBは苦手である")

		time.Sleep(time.Duration(100) * time.Millisecond)
		transitionTo(logger, "食材の準備OK", "")
	}()

	func() {
		ctx, span := tracer.Start(ctx, "ドレッシングを作る")
		defer span.End()
		logger := logging.WithTrace(ctx, logger)
		useMaterials(logger, "白ワインビネガー", "cc", "15", "")
		useMaterials(logger, "オリーブオイル", "cc", "15", "")
		useMaterials(logger, "塩", "cc", "15", "")
		useMaterials(logger, "レモン汁", "cc", "5", "")

		doCookAction(logger, "撹拌", "", "頑張って乳化させる")

		time.Sleep(time.Duration(100) * time.Millisecond)
		transitionTo(logger, "ドレッシングが完成する", "")
	}()

	func() {
		ctx, span := tracer.Start(ctx, "仕上げ")
		defer span.End()
		logger := logging.WithTrace(ctx, logger)
		time.Sleep(time.Duration(10) * time.Millisecond)
		doCookAction(logger, "野菜とドレッシングを和える", "", "和えたらとっとと食べる")
	}()

	c.String(http.StatusOK, name+"完成")
}
