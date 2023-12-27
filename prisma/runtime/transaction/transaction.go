package transaction

import (
	"context"
	"fmt"

	"github.com/hosseinmirzapur/goravel-cli/prisma/engine"
	"github.com/hosseinmirzapur/goravel-cli/prisma/runtime/builder"
)

type TX struct {
	Engine engine.Engine
}

type Param interface {
	IsTx()
	ExtractQuery() builder.Query
}

func (r TX) Transaction(queries ...Param) Exec {
	return Exec{
		engine:  r.Engine,
		queries: queries,
	}
}

type Exec struct {
	queries  []Param
	engine   engine.Engine
	requests []engine.GQLRequest
}

func (r Exec) Exec(ctx context.Context) error {
	r.requests = make([]engine.GQLRequest, len(r.queries))
	for i, query := range r.queries {
		str, err := query.ExtractQuery().Build()
		if err != nil {
			return err
		}
		r.requests[i] = engine.GQLRequest{
			Query:     str,
			Variables: map[string]interface{}{},
		}
	}

	for _, q := range r.queries {
		//goland:noinspection GoDeferInLoop
		defer close(q.ExtractQuery().TxResult)
	}

	var result engine.GQLBatchResponse
	payload := engine.GQLBatchRequest{
		Batch:       r.requests,
		Transaction: true,
	}
	if err := r.engine.Batch(ctx, payload, &result); err != nil {
		return fmt.Errorf("could not send raw query: %w", err)
	}
	if len(result.Errors) > 0 {
		first := result.Errors[0]
		return fmt.Errorf("pql error: %s", first.RawMessage())
	}
	for i, inner := range result.Result {
		if len(inner.Errors) > 0 {
			first := result.Errors[0]
			return fmt.Errorf("pql error: %s", first.RawMessage())
		}

		r.queries[i].ExtractQuery().TxResult <- inner.Data.Result
	}
	return nil
}
