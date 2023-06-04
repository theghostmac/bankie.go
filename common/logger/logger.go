package logger

import (
	"context"
)

type Logger struct {
	ctx    context.Context
	fields map[string]interface{}
}
