package restapi

import (
	"context"
	"mongodb-trx/domain_mytrx/usecase/createorder"
	"mongodb-trx/shared/gogen"
	"mongodb-trx/shared/infrastructure/logger"
	"mongodb-trx/shared/model/payload"
	"mongodb-trx/shared/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *ginController) createOrderHandler() gin.HandlerFunc {

	type InportRequest = createorder.InportRequest
	type InportResponse = createorder.InportResponse

	inport := gogen.GetInport[InportRequest, InportResponse](r.GetUsecase(InportRequest{}))

	type request struct {
	}

	type response struct {
	}

	return func(c *gin.Context) {

		traceID := util.GenerateID(16)

		ctx := logger.SetTraceID(context.Background(), traceID)

		//var jsonReq request
		//if err := c.BindJSON(&jsonReq); err != nil {
		//	r.log.Error(ctx, err.Error())
		//	c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
		//	return
		//}

		var req InportRequest
		req.RandomIDForOrder = util.GenerateID(16)
		req.RandomIDForPerson = util.GenerateID(16)

		r.log.Info(ctx, util.MustJSON(req))

		res, err := inport.Execute(ctx, req)
		if err != nil {
			r.log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}

		var jsonRes response
		_ = res

		r.log.Info(ctx, util.MustJSON(jsonRes))
		c.JSON(http.StatusOK, payload.NewSuccessResponse(jsonRes, traceID))

	}
}
