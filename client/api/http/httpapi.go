package http

import (
	"go-sample-grpc-client/internal/data/model"
	"go-sample-grpc-client/internal/pkg/core/adapter/pbclientadapter"
	"net/http"

	"github.com/labstack/echo"
)

type APIIntegrator struct {
	pbclient pbclientadapter.OutboundAdapter
}

func NewAPIIntegrator(pbclient pbclientadapter.OutboundAdapter) APIIntegrator {
	return APIIntegrator{
		pbclient: pbclient,
	}
}

func API(e *echo.Echo, pbclient pbclientadapter.OutboundAdapter) {
	api := e.Group("")
	integrator := NewAPIIntegrator(pbclient)

	api.POST("/request", integrator.Request)
}

func (a APIIntegrator) Request(c echo.Context) error {
	var req model.RequestAPIReq
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.RequestAPIResp{
			Error: ErrInternalServer.Error(),
		})
	}

	resp, err := a.pbclient.Request(pbclientadapter.RequestRequest{
		RequestString: req.RequestString,
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.RequestAPIResp{
			Error: ErrInternalServer.Error(),
		})
	}

	return c.JSON(http.StatusOK, model.RequestAPIResp{
		ResponseString: resp.ResponseString,
	})
}
