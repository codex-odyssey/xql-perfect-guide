package recipe

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	logging "app/internal/log"
)

func Curry(c *gin.Context) {
	name := "curry"
	ctx := c.Request.Context()
	logger := logging.GetLoggerWithTraceID(ctx)

	func() {
		ctx, span := tracer.Start(ctx, "食材の準備")
		defer span.End()
		logger := logging.WithTrace(ctx, logger)
		useMaterials(logger, "肉", "200", "g", "牛でも豚でも鶏でも羊でもなんでもうまい")
		useMaterials(logger, "玉ねぎ", "150", "g", "くし切り")
		useMaterials(logger, "人参", "100", "g", "乱切り")
		useMaterials(logger, "じゃがいも", "100", "g", "一口サイズに")
		useMaterials(logger, "水", "800", "ml", "")
		useMaterials(logger, "カレールー", "120", "g", "")
		time.Sleep(time.Duration(300) * time.Millisecond)
		transitionTo(logger, "食材の準備OK", "")
	}()

	func() {
		ctx, span := tracer.Start(ctx, "肉を炒める")
		defer span.End()
		logger := logging.WithTrace(ctx, logger)
		doCookAction(logger, "肉", "焼く", "ちょっと焦げるくらいがうまい！")

		time.Sleep(time.Duration(30) * time.Millisecond)
		transitionTo(logger, "肉が焼ける", "")
	}()

	func() {
		ctx, span := tracer.Start(ctx, "野菜を炒める")
		defer span.End()
		logger := logging.WithTrace(ctx, logger)
		doCookAction(logger, "玉ねぎ", "焼く", "ひと塩すると旨みが凝縮")
		doCookAction(logger, "人参", "焼く", "何となくでいい")
		doCookAction(logger, "じゃがいも", "焼く", "何となくでいい")
		time.Sleep(time.Duration(50) * time.Millisecond)
		transitionTo(logger, "いい感じに炒め終わる", "")
	}()

	func() {
		ctx, span := tracer.Start(ctx, "水を加えて煮る")
		defer span.End()
		logger := logging.WithTrace(ctx, logger)
		doCookAction(logger, "水", "煮る", "")
		time.Sleep(time.Duration(200) * time.Millisecond)
		transitionTo(logger, "人参とじゃがいもに火が通る", "")
	}()

	func() {
		ctx, span := tracer.Start(ctx, "カレールーを加える")
		defer span.End()
		logger := logging.WithTrace(ctx, logger)
		doCookAction(logger, "カレールー", "溶かす", "入れる時は火を消すこと")
		time.Sleep(time.Duration(50) * time.Millisecond)
		transitionTo(logger, "カレールーが溶ける", "")
	}()

	func() {
		ctx, span := tracer.Start(ctx, "煮込む")
		defer span.End()
		logger := logging.WithTrace(ctx, logger)
		doCookAction(logger, "", "煮込む", "適度なとろみがつけばOK! 煮込みすぎるとじゃがいもが消滅するぞ")
		time.Sleep(time.Duration(50) * time.Millisecond)
	}()

	c.String(http.StatusOK, name+"完成")
}
