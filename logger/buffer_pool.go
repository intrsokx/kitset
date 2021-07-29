package logger

import (
	"bytes"
	"sync"
)

type BufferPool interface {
	Get() *bytes.Buffer
	Put(buffer *bytes.Buffer)
}

type bufferPool struct {
	pool *sync.Pool
}

func (b *bufferPool) Get() *bytes.Buffer {
	return b.pool.Get().(*bytes.Buffer)
}

func (b *bufferPool) Put(buffer *bytes.Buffer) {
	buffer.Reset()
	b.pool.Put(buffer)
}

var defaultPool BufferPool

func init() {
	defaultPool = &bufferPool{
		pool: &sync.Pool{
			New: func() interface{} {
				return &bytes.Buffer{}
			},
		},
	}
}

func NewBufferPool() BufferPool {
	return &bufferPool{
		pool: &sync.Pool{
			New: func() interface{} {
				return &bytes.Buffer{}
			},
		},
	}
}
