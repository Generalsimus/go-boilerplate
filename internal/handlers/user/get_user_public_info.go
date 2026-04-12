package user

import (
	"context"
	"net/http"

	"github.com/Generalsimus/go-monolith-boilerplate/internal/delivery"
	"github.com/danielgtaylor/huma/v2"
)

type GetUserPublicInfoRequest struct {
	Id int `query:"id" json:"id" required:"true" minimum:"0"`
}

type GetUserPublicInfoResponseBody struct {
	ID   int    `json:"id"`
	Name string `json:"name,omitempty"`
}

type GetUserPublicInfoResponse struct {
	Body GetUserPublicInfoResponseBody
}

func (h *Handler) RegisterUserPublicInfoHandler(api huma.API) {
	endpoint := &delivery.BaseEndpoint[GetUserPublicInfoRequest, GetUserPublicInfoResponse]{
		OperationID: "get-user-public-info",
		Method:      http.MethodGet,
		Path:        "/user-public-info/",
		Summary:     "Get user public info",
		Tags:        []string{"User"},
		Websocket:   true,
		Handler:     h.GetUserPublicInfoHandler,
	}
	endpoint.Register(api)
}

func (h *Handler) GetUserPublicInfoHandler(ctx context.Context, req *GetUserPublicInfoRequest) (*GetUserPublicInfoResponse, error) {

	return &GetUserPublicInfoResponse{Body: GetUserPublicInfoResponseBody{ID: req.Id, Name: "public name"}}, nil
}
