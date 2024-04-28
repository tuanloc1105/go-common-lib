package utils

import (
	"math/big"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tuanloc1105/go-common-lib/constant"
)

func CheckAndSetTraceId(c *gin.Context) {
	if traceId, _ := c.Get(string(constant.TraceIdLogKey)); traceId == nil || traceId == "" {
		c.Set(string(constant.TraceIdLogKey), uuid.New().String())
	}
}

func GetTraceId(c *gin.Context) string {
	if traceId, _ := c.Get(string(constant.TraceIdLogKey)); traceId == nil || traceId == "" {
		return ""
	} else {
		return traceId.(string)
	}
}

func RoundHalfUpBigFloat(input *big.Float) {
	delta := constant.DeltaPositive

	if input.Sign() < 0 {
		delta = constant.DeltaNegative
	}
	input.Add(input, new(big.Float).SetFloat64(delta))
}

func GetPointerOfAnyValue[T any](a T) *T {
	return &a
}
