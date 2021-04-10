package caller

import (
	"context"
)

type ApiCaller interface {
	Call(ctx context.Context, action string, data map[string]interface{}) (map[string]interface{}, error)
}
