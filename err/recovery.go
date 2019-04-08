package err

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/errors"
	"github.com/txvier/base/logger"
	"net/http"
	"runtime/debug"
)

func Recovery() gin.HandlerFunc {
	return func(cxt *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				var errMsg string
				var httpCode int = http.StatusInternalServerError
				if oe, ok := err.(error); ok {
					errMsg = errors.ErrorStack(oe)
					if ae, ok := err.(*AppError); ok {
						errMsg = errors.ErrorStack(ae.GetError())
						httpCode = ae.HttpStatusCode
					}
					logger.Logger.Error(errMsg)
					debug.PrintStack()
					cxt.JSON(httpCode, map[string]string{
						"rsp_msg": oe.Error(),
					})
				}
			}
		}()
		cxt.Next()
	}
}