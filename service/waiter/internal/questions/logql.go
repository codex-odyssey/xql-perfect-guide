package questions

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	logging "app/internal/log"
)

func HandlerLogQ(c *gin.Context) {
	ctx := c.Request.Context()
	logger := logging.GetLoggerFromCtx(ctx)

	l1 := logger.With(
		"question", "1",
	)

	l1.Errorln("foo")
	l1.Infoln("foo")
	l1.Warn("foo")
	l1.Infoln("var")
	l1.Error("var")
	l1.Infoln("foo")

	l2 := logger.With(
		"question", "2",
	)
	l2.Infoln("foofoo")
	l2.Infoln("foovar")
	l2.Infoln("varvar")
	l2.Infoln("varfoo")

	outputLabelFilterLog(c.Request.Context(), "aaa", "4s", "2mb", 1)
	outputLabelFilterLog(c.Request.Context(), "bbb", "1ms", "1mb", 2)
	outputLabelFilterLog(c.Request.Context(), "aaa", "1s", "3kb", 10)

	l4 := logger.With(
		"question", "4",
	)
	l4.Infoln("containerラベルをappラベルにしたい！")
	l4.Infoln("時代はappラベル")

	l5 := logger.With(
		"question", "5",
	)
	l5.Infoln("")
	l5.Infoln("")
	l5.Infoln("")

	c.String(http.StatusOK, "ok")
}

func outputLabelFilterLog(ctx context.Context, s, d, b string, n int) {
	logger := logging.GetLoggerFromCtx(ctx)
	logger.With(
		"string", s,
		"duration", d,
		"bytes", b,
		"number", n,
	).Infoln("label filter")
}
