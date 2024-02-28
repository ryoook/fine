package middleware

import (
	"FINE/util"
	"github.com/gin-gonic/gin"
	"time"
)

type InputIn struct {
	//LogInstance      *logs.Logger // logger
	ServiceName string // 服务名称
}

// Base set some values to ctx
func Base(in InputIn) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now().UnixNano()
		ctx.Set(util.RequestStartTime, start)

		requestID, traceID := util.GetRequestIDAndTraceID(ctx)

		ctx.Set(util.RequestID, requestID)
		ctx.Set(util.TraceID, traceID)
		ctx.Set(util.ServiceName, in.ServiceName)

		ctx.Set(util.ClientIP, ctx.ClientIP())
		ctx.Set(util.ServerIP, util.HostIP())
		fineContext := &util.FineContext{
			RequestID: requestID,
			TraceID:   traceID,
		}
		ctx.Set(util.Context, fineContext)
	}
}

// Response ...
type Response struct {
	RequestID string      `json:"request_id"`
	TraceID   string      `json:"trace_id"`
	Code      int         `json:"code"`
	Status    string      `json:"status"`
	Cost      int64       `json:"cost"`
	Data      interface{} `json:"data"`
}

func ResponseJSON(c *gin.Context) {
	c.Next()
	fineContext := util.GetContext(c)
	if requestStartTime, exist := c.Get(util.RequestStartTime); exist {
		if startTime, ok := requestStartTime.(time.Time); ok {
			fineContext.Cost = time.Since(startTime).Nanoseconds() / 1000 / 1000
		}
	}
	resErr := util.CastError(fineContext.Error)
	res := Response{
		RequestID: fineContext.RequestID,
		TraceID:   fineContext.TraceID,
		Code:      resErr.Errno,
		Status:    resErr.ErrMsg,
		Cost:      fineContext.Cost,
		Data:      fineContext.Data,
	}
	fineContext.ResponseBody = res
	c.JSON(resErr.Status, res)
}
