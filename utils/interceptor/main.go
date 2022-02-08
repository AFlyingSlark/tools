package main

import (
	"context"
	"fmt"
)

// grpc 的经典实现 c.cc.Invoke()

type interceptor2 func(ctx context.Context, h handler, ivk invoker) error
type handler func(ctx context.Context)
type invoker func(ctx context.Context, interceptors []interceptor2, h handler) error

func main() {
	var ctx context.Context
	var ceps []interceptor2
	var h = func(ctx context.Context) {
		fmt.Println("do something ...")
	}
	var inter1 = func(ctx context.Context, h handler, ivk invoker) error {
		h(ctx)
		return ivk(ctx, ceps, h)
	}
	var inter2 = func(ctx context.Context, h handler, ivk invoker) error {
		h(ctx)
		return ivk(ctx, ceps, h)
	}
	var inter3 = func(ctx context.Context, h handler, ivk invoker) error {
		h(ctx)
		return ivk(ctx, ceps, h)
	}
	ceps = append(ceps, inter1, inter2, inter3)
	var ivk = func(ctx context.Context, interceptors []interceptor2, h handler) error {
		fmt.Println("invoker start")
		return nil
	}
	cep := getChainInterceptor(ctx, ceps, ivk)
	cep(ctx, h, ivk)
}
func getChainInterceptor(ctx context.Context, interceptors []interceptor2, ivk invoker) interceptor2 {
	if len(interceptors) == 0 {
		return nil
	}
	if len(interceptors) == 1 {
		return interceptors[0]
	}
	return func(ctx context.Context, h handler, ivk invoker) error {
		return interceptors[0](ctx, h, getInvoker(ctx, interceptors, 0, ivk))
	}
}
func getInvoker(ctx context.Context, interceptors []interceptor2, cur int, ivk invoker) invoker {
	if cur == len(interceptors)-1 {
		return ivk
	}
	return func(ctx context.Context, interceptors []interceptor2, h handler) error {
		return interceptors[cur+1](ctx, h, getInvoker(ctx, interceptors, cur+1, ivk))
	}
}
