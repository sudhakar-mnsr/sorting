package stream

import (
	"context"
	"errors"
	"io"
	"os"
	"reflect"

	"github.com/vladimirvivien/automi/api"
	autoctx "github.com/vladimirvivien/automi/api/context"
	"github.com/vladimirvivien/automi/collectors"
	"github.com/vladimirvivien/automi/emitters"
	streamop "github.com/vladimirvivien/automi/operators/stream"
	"github.com/vladimirvivien/automi/util"
)

// Stream represents a stream unto  which executor nodes can be
// attached to operate on the streamed data
type Stream struct {
	srcParam interface{}
	snkParam interface{}
	source   api.Source
	sink     api.Sink
	drain    chan error
	ops      []api.Operator
	ctx      context.Context
	logf     api.LogFunc
	errf     api.ErrorFunc
}

// New creates a new *Stream value
func New(src interface{}) *Stream {
	s := &Stream{
		srcParam: src,
		ops:      make([]api.Operator, 0),
		drain:    make(chan error),
	}

	return s
}

// WithContext sets a context.Context to use.
func (s *Stream) WithContext(ctx context.Context) *Stream {
	s.ctx = ctx
	return s
}

// WithLogFunc sets a function that will receive internal log events
// at runtime.  Supported log function type: func(interface{})
func (s *Stream) WithLogFunc(fn api.LogFunc) *Stream {
	s.logf = fn
	return s
}
