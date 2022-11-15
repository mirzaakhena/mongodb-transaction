package restapi

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"mongodb-trx/domain_mytrx/usecase/createorder"
	"mongodb-trx/shared/infrastructure/logger"
	"mongodb-trx/shared/infrastructure/util"
	"mongodb-trx/shared/model/payload"
)

// createOrderHandler ...
func (r *Controller) createOrderHandler(inputPort createorder.Inport) gin.HandlerFunc {

	type request struct {
	}

	type response struct {
	}

	return func(c *gin.Context) {

		traceID := util.GenerateID(16)

		ctx := logger.SetTraceID(context.Background(), traceID)

		//var jsonReq request
		//if err := c.BindJSON(&jsonReq); err != nil {
		//	r.Log.Error(ctx, err.Error())
		//	c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
		//	return
		//}

		var req createorder.InportRequest

		r.Log.Info(ctx, util.MustJSON(req))

		res, err := inputPort.Execute(ctx, req)
		if err != nil {
			r.Log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}

		var jsonRes response
		_ = res

		r.Log.Info(ctx, util.MustJSON(jsonRes))
		c.JSON(http.StatusOK, payload.NewSuccessResponse(jsonRes, traceID))

	}
}
