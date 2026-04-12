package user

import (
	"context"
	"net/http"

	"github.com/Generalsimus/go-monolith-boilerplate/internal/delivery"
	"github.com/danielgtaylor/huma/v2"
)

type GetUserInfoRequest struct {
	Id int `query:"id" json:"id" required:"true" minimum:"0"`
}

type GetUserInfoResponseBody struct {
	ID   int    `json:"id"`
	Name string `json:"name,omitempty"`
}
type GetUserInfoResponse struct {
	Body GetUserInfoResponseBody
}

func (h *Handler) RegisterUserInfoHandler(api huma.API) {
	endpoint := &delivery.BaseEndpoint[GetUserInfoRequest, GetUserInfoResponse]{
		OperationID: "get-user-info",
		Method:      http.MethodGet,
		Path:        "/",
		Summary:     "Get user info",
		Tags:        []string{"User"},
		Websocket:   true,
		Handler:     h.GetUserInfoHandler,
	}
	endpoint.Register(api)
}

func (h *Handler) GetUserInfoHandler(ctx context.Context, req *GetUserInfoRequest) (*GetUserInfoResponse, error) {

	return &GetUserInfoResponse{Body: GetUserInfoResponseBody{ID: req.Id, Name: "name"}}, nil
}
