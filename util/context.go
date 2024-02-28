package util

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"strings"
	"time"
)

const (
	HeaderTraceId = "Fine-Trace-Id"
)
const (
	Context          = "fine_context"
	RequestStartTime = "request_start_time"
	RequestID        = "request_id"
	TraceID          = "trace_id"
	ServiceName      = "service_name"
	ClientIP         = "client_ip"
	ServerIP         = "server_ip"
)

type FineContext struct {
	RequestID    string
	TraceID      string
	Code         int
	Error        error
	Data         interface{}
	Cost         int64
	ResponseBody interface{}
}

func GetContext(ctx *gin.Context) *FineContext {
	// get from gin context
	if rawContext, exists := ctx.Get(Context); exists {
		if c, ok := rawContext.(*FineContext); ok {
			return c
		}
	}

	requestID, traceID := GetRequestIDAndTraceID(ctx)
	fineContext := &FineContext{
		RequestID: requestID,
		TraceID:   traceID,
	}
	start := time.Now().UnixNano()
	ctx.Set(RequestStartTime, start)
	ctx.Set(RequestID, requestID)
	ctx.Set(TraceID, traceID)
	ctx.Set(ClientIP, ctx.ClientIP())
	ctx.Set(ServerIP, HostIP())

	ctx.Set(Context, fineContext)

	return fineContext
}

func (t *FineContext) Response(data interface{}, err error) {
	t.Data = data
	t.Error = err
}

// ResponseSuccess ...
func (t *FineContext) ResponseSuccess(data interface{}) {
	t.Response(data, nil)
}

// ResponseError ...
func (t *FineContext) ResponseError(err error) {
	t.Response(nil, err)
}

// GetRequestIDAndTraceID ...
func GetRequestIDAndTraceID(ctx *gin.Context) (string, string) {
	// requestID: 'time'-'hostIP(no dot)'-'random string'
	// traceID: random string
	var (
		hIP       = HostIP()
		requestID = GenerateHostIPTraceID(hIP)
		traceID   string
	)
	if traceID = ctx.GetHeader(HeaderTraceId); len(traceID) > 0 {
		return requestID, traceID
	}
	traceID = GenRequestID()
	return requestID, traceID
}

func GenerateHostIPTraceID(hostIP string) string {
	formatTime := time.Now().Format("20060102150405")
	return fmt.Sprintf("%s-%s-%s", formatTime, strings.ReplaceAll(hostIP, ".", ""), GenRequestID())
}
func GenRequestID() string {

	chars := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()+-")

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomByte := []byte{
		chars[r.Intn(74)], chars[r.Intn(74)], chars[r.Intn(74)], chars[r.Intn(74)], chars[r.Intn(74)],
	}

	return MD5(fmt.Sprintf("%s-%s", Unique(""), string(randomByte)))
}

// MD5 ...
func MD5(input string) string {
	sum := md5.Sum([]byte(input))
	return hex.EncodeToString(sum[:])
}

func Unique(prefix string) string {
	now := time.Now()
	return fmt.Sprintf("%s%08x%05x", prefix, now.Unix(), now.UnixNano()%0x100000)
}
