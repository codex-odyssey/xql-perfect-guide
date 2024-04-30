package recipe

import (
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
)

var tracer = otel.GetTracerProvider().Tracer("")

func useMaterials(logger *zap.SugaredLogger, material, unit, amount, tips string) {
	logger.With(
		"material", material,
		"unit", unit,
		"amount", amount,
		"tips", tips,
	).Info("材料を使用")
}

func transitionTo(logger *zap.SugaredLogger, state, tips string) {
	logger.With(
		"state", state,
		"tips", tips,
	).Info("状態が遷移")
}

func doCookAction(logger *zap.SugaredLogger, material, action, tips string) {
	logger.With(
		"material", material,
		"action", action,
		"tips", tips,
	).Info("調理")
}
