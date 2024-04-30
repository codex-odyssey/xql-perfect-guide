package recipe

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	logging "app/internal/log"
)

func Yasaiitame(c *gin.Context) {
	name := "yasaiitame"
	ctx := c.Request.Context()
	logger := logging.GetLoggerWithTraceID(ctx)

	func() {
		ctx, span := tracer.Start(ctx, "食材の準備")
		defer span.End()
		logger := logging.WithTrace(ctx, logger)
		useMaterials(logger, "豚肉", "", "適量", "一口サイズ")
		useMaterials(logger, "玉ねぎ", "", "適量", "くし切り")
		useMaterials(logger, "人参", "", "適量", "銀杏切り")
		useMaterials(logger, "キャベツ", "", "適量", "銀杏切り")
		useMaterials(logger, "オイスターソース", "", "適量", "")
		useMaterials(logger, "塩", "", "適量", "銀杏切り")
		useMaterials(logger, "胡椒", "", "適量", "銀杏切り")

		time.Sleep(time.Duration(100) * time.Millisecond)
		transitionTo(logger, "食材の準備OK", "")
	}()

	func() {
		ctx, span := tracer.Start(ctx, "肉を炒める")
		defer span.End()
		logger := logging.WithTrace(ctx, logger)
		doCookAction(logger, "豚肉", "炒める", "ちょっと焦げてるくらいがいい")

		time.Sleep(time.Duration(300) * time.Millisecond)
		transitionTo(logger, "豚肉が炒め終わる", "")
	}()

	func() {
		ctx, span := tracer.Start(ctx, "野菜を炒める")
		defer span.End()
		logger := logging.WithTrace(ctx, logger)
		doCookAction(logger, "玉ねぎ", "炒める", "ひと塩で旨みUP")
		doCookAction(logger, "人参", "炒める", "ひと塩で旨みUP")
		doCookAction(logger, "キャベツ", "炒める", "ひと塩で旨みUP")

		time.Sleep(time.Duration(100) * time.Millisecond)
		transitionTo(logger, "野菜に火が通る", "")
	}()

	func() {
		ctx, span := tracer.Start(ctx, "味付け")
		defer span.End()
		logger := logging.WithTrace(ctx, logger)
		doCookAction(logger, "オイスターソース", "絡める", "ひと塩で旨みUP")

		time.Sleep(time.Duration(100) * time.Millisecond)
		transitionTo(logger, "全体に味が回る", "")
	}()

	c.String(http.StatusOK, name+"完成")
}
