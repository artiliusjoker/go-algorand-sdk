package algod

import (
	"context"

	"github.com/algorand/go-algorand-sdk/client/v2/common"
)

// Shutdown the node with REST API
type Shutdown struct {
	c *Client
}

func (s *Shutdown) Do(ctx context.Context, headers ...*common.Header) (err error) {
	var response byte = 1
	err = s.c.post(ctx, &response, "/v2/shutdown", nil, headers)
	return
}
