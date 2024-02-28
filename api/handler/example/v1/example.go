package v1

import (
	"FINE/internal/model"
	"FINE/internal/service"
	"FINE/util"
	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	fineContext := util.GetContext(c)
	var in model.HelloIn
	if err := c.ShouldBind(&in); err != nil {
		fineContext.ResponseError(util.ErrInvalidParam.More(err.Error()))
		return
	}
	fineContext.Response(service.Example().SayHello(c, &in))
}
