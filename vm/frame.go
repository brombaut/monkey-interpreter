package vm

import (
	"github.com/brombaut/monkey-interpreter/code"
	"github.com/brombaut/monkey-interpreter/object"
)

type Frame struct {
	fn          *object.CompiledFunction
	ip          int
	basePointer int
}

func NewFrame(fn *object.CompiledFunction, basePointer int) *Frame {
	f := &Frame{
		fn:          fn,
		ip:          -1,
		basePointer: basePointer,
	}
	return f
}

func (f *Frame) Instructions() code.Instructions {
	return f.fn.Instructions
}
