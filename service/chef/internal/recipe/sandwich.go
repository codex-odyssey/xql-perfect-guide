package recipe

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	logging "app/internal/log"
)

func Sandwich(c *gin.Context) {
	name := "sandwich"
	ctx := c.Request.Context()
	logger := logging.GetLoggerWithTraceID(ctx)

	func() {
		ctx, span := tracer.Start(ctx, "食材の準備")
		defer span.End()
		logger := logging.WithTrace(ctx, logger)
		useMaterials(logger, "パン", "", "適量", "")
		useMaterials(logger, "たまご", "", "適量", "")
		useMaterials(logger, "マヨネーズ", "", "適量", "")
		useMaterials(logger, "マスタード", "", "適量", "")

		time.Sleep(time.Duration(100) * time.Millisecond)
		transitionTo(logger, "食材の準備OK", "")
	}()

	func() {
		ctx, span := tracer.Start(ctx, "ゆでたまごを作る")
		defer span.End()
		logger := logging.WithTrace(ctx, logger)
		doCookAction(logger, "たまご", "茹でる", "8分くらいがベスト")

		time.Sleep(time.Duration(480) * time.Millisecond)
		transitionTo(logger, "ゆでたまごが完成", "")
	}()

	func() {
		ctx, span := tracer.Start(ctx, "パンを準備する")
		defer span.End()
		logger := logging.WithTrace(ctx, logger)
		doCookAction(logger, "パン", "焼く", "生でもうまい")
		transitionTo(logger, "パンが焼ける", "")

		time.Sleep(time.Duration(100) * time.Millisecond)
		doCookAction(logger, "マスタード", "塗る", "うっすら塗る")
		transitionTo(logger, "マスタードが塗り終わる", "")
	}()

	func() {
		ctx, span := tracer.Start(ctx, "仕上げ")
		defer span.End()
		logger := logging.WithTrace(ctx, logger)
		doCookAction(logger, "", "組み立てる", "")
		time.Sleep(time.Duration(30) * time.Millisecond)
	}()

	c.String(http.StatusOK, name+"完成")
}
