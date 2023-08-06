package base

import "context"

type Func func()

type KVFuncB[K any, V any] func(K, V) bool

type FuncE func() error

type CtxFunc func(context.Context)

type EFunc func(error)

type CtxEFunc func(context.Context, error)

type CtxAFunc func(context.Context, any)

type CtxTFunc[T any] func(context.Context, T)

type FuncCtxE func() (context.Context, error)

type CtxFuncE func(context.Context) error

type CtxFuncCtxE func(context.Context) (context.Context, error)

type CtxAFuncAE func(context.Context, any) (any, error)

type CtxPFuncRE[P any, R any] func(context.Context, P) (R, error)

type CtxFuncSE func(context.Context) (string, error)

type CtxFuncIE func(context.Context) (int64, error)

type CtxFuncFE func(context.Context) (float64, error)

type BuiltinT interface {
	string | int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64 |
		[]string | []int | []int8 | []int16 | []int32 | []int64 | []uint | []uint8 | []uint16 | []uint32 | []uint64 | []float32 | []float64
}
