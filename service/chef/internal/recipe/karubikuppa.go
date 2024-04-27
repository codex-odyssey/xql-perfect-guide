package recipe

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	logging "app/internal/log"
)

func Karubikuppa(c *gin.Context) {
	name := "karubikuppa"
	ctx := c.Request.Context()
	logger := logging.GetLoggerWithTraceID(ctx)

	func() {
		ctx, span := tracer.Start(ctx, "カルビを炒める")
		defer span.End()
		logger := logging.WithTrace(ctx, logger)
		useMaterials(logger, "ごま油", "", "適量", "")
		useMaterials(logger, "ニンニク", "", "適量", "")

		time.Sleep(time.Duration(50) * time.Millisecond)
		transitionTo(logger, "香りが出る", "")
		useMaterials(logger, "牛肉", "", "適量", "多ければ多いほどうまい")
	}()

	func() {
		ctx, span := tracer.Start(ctx, "カルビクッパを煮込む")
		defer span.End()
		logger := logging.WithTrace(ctx, logger)
		useMaterials(logger, "水", "cc", "800", "")
		useMaterials(logger, "コチュジャン", "大さじ", "2", "多ければ多いほどうまい")
		useMaterials(logger, "ウェイパー", "大さじ", "2", "")
		useMaterials(logger, "ごま油", "大さじ", "2", "多ければ多いほどうまい")
		useMaterials(logger, "酒", "大さじ", "2", "")
		useMaterials(logger, "醤油", "大さじ", "2", "")
		useMaterials(logger, "砂糖", "小さじ", "1", "")
		useMaterials(logger, "塩コショウ", "", "適量", "")
		useMaterials(logger, "大根", "", "適量", "")
		useMaterials(logger, "人参", "", "適量", "")
		useMaterials(logger, "シイタケ", "", "適量", "")

		time.Sleep(time.Duration(300) * time.Millisecond)
		transitionTo(logger, "一煮立ちする", "")
	}()

	func() {
		ctx, span := tracer.Start(ctx, "仕上げ")
		defer span.End()
		logger := logging.WithTrace(ctx, logger)
		useMaterials(logger, "ニラ", "", "適量", "")
		useMaterials(logger, "もやし", "", "適量", "")
		time.Sleep(time.Duration(30) * time.Millisecond)
		transitionTo(logger, "極限まで沸騰する", "")
		useMaterials(logger, "卵", "個", "2", "")
	}()

	func() {
		ctx, span := tracer.Start(ctx, "盛り付け")
		defer span.End()
		logger := logging.WithTrace(ctx, logger)
		useMaterials(logger, "出来立ての米", "", "適量", "")
		time.Sleep(time.Duration(10) * time.Millisecond)
		transitionTo(logger, "米にカルビクッパをかける", "")
	}()

	c.String(http.StatusOK, name+"完成")
}
