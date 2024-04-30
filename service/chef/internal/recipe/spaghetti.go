package recipe

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	logging "app/internal/log"
)

func Spaghetti(c *gin.Context) {
	name := "spaghetti"
	ctx := c.Request.Context()
	logger := logging.GetLoggerWithTraceID(ctx)

	func() {
		ctx, span := tracer.Start(ctx, "食材の準備")
		defer span.End()
		logger := logging.WithTrace(ctx, logger)
		useMaterials(logger, "にんにく", "5", "g", "国産がいいぞ")
		useMaterials(logger, "オリーブオイル", "15", "ml", "")
		useMaterials(logger, "鷹の爪", "2", "g", "")
		useMaterials(logger, "パセリ", "5", "g", "普通のパセリでもイタリアンパセリでもOK! 乾燥パセリ使うくらいなら使わない方がマシ")
		useMaterials(logger, "パスタ", "100", "g", "")
		useMaterials(logger, "水", "2", "l", "")
		useMaterials(logger, "塩", "20", "g", "1%くらいにすると美味しく茹でれる")

		time.Sleep(time.Duration(100) * time.Millisecond)
		transitionTo(logger, "食材の準備OK", "")
	}()

	func() {
		ctx, span := tracer.Start(ctx, "ソースを作る")
		defer span.End()
		logger := logging.WithTrace(ctx, logger)
		doCookAction(logger, "にんにく", "炒める", "超弱火")
		transitionTo(logger, "にんにくの香りをオイルに移す", "")
		doCookAction(logger, "鷹の爪", "炒める", "2秒くらいでいいぞ")
		transitionTo(logger, "鷹の爪の香りをオイルに移す", "")
		doCookAction(logger, "パセリ", "半量いれる", "")
		transitionTo(logger, "パセリが入れ終わる", "")
		doCookAction(logger, "茹で汁", "茹で汁を入れて乳化する", "オイルと同量くらいがおすすめ")

		time.Sleep(time.Duration(480) * time.Millisecond)
		transitionTo(logger, "ソースが作り終わる", "")
	}()

	func() {
		ctx, span := tracer.Start(ctx, "パスタを茹でる")
		defer span.End()
		logger := logging.WithTrace(ctx, logger)
		doCookAction(logger, "パスタ", "茹でる", "何も考えるな！")

		time.Sleep(time.Duration(480) * time.Millisecond)
		transitionTo(logger, "茹で終わる", "")
	}()

	func() {
		ctx, span := tracer.Start(ctx, "ソースとパスタを和える")
		defer span.End()
		logger := logging.WithTrace(ctx, logger)
		doCookAction(logger, "", "ソースとパスタを和える", "")
		time.Sleep(time.Duration(30) * time.Millisecond)
	}()

	c.String(http.StatusOK, name+"完成")
}
