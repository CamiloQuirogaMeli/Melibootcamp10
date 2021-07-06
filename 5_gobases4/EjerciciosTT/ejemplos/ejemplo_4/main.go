package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	fmt.Println("context.Background(): ", ctx)
	ctx = context.WithValue(ctx, "saludo", "hola digital house!! Quiero una milanesa")
	fmt.Println("context.WithValue(..., ..., ...): ", ctx)
	saludoWrapper(ctx)
}

func saludoWrapper(ctx context.Context) {
	saludo(ctx)
}

func saludo(ctx context.Context) {
	fmt.Println(ctx.Value("saludo"))
}
