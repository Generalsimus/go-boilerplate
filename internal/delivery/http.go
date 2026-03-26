package delivery

import (
	"context"
	"encoding/json"

	"github.com/danielgtaylor/huma/v2"
)

type BaseEndpoint[Input any, Output any] struct {
	OperationID string
	Method      string
	Path        string
	Summary     string
	Tags        []string
	Websocket   bool
	Handler     func(ctx context.Context, req *Input) (*Output, error)
}

func (b *BaseEndpoint[I, O]) Register(api huma.API) {
	httpOperation := huma.Operation{
		OperationID: b.OperationID,
		Method:      b.Method,
		Path:        b.Path,
		Summary:     b.Summary,
		Tags:        b.Tags,
	}
	huma.Register(api, httpOperation, b.Handler)

	if b.Websocket == false {
		return
	}

	wsMethodGroup := WsRouter[b.Method]
	if wsMethodGroup == nil {
		wsMethodGroup = make(map[string]func(context.Context, json.RawMessage) (any, error))
		WsRouter[b.Method] = wsMethodGroup
	}

	wsHandler := func(ctx context.Context, rawData json.RawMessage) (any, error) {
		var input I

		if len(rawData) > 0 && string(rawData) != "null" {
			if err := json.Unmarshal(rawData, &input); err != nil {
				return nil, huma.Error400BadRequest(err.Error())
			}
		}

		output, err := b.Handler(ctx, &input)
		if err != nil {
			return nil, err
		}

		return output, nil
	}

	if grp, ok := api.(*huma.Group); ok {
		grp.ModifyOperation(&httpOperation, func(modifiedOp *huma.Operation) {
			wsMethodGroup[modifiedOp.Path] = wsHandler
		})
	} else {
		wsMethodGroup[b.Path] = wsHandler
	}
}
