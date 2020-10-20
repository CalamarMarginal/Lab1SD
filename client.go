package main

import (
	"fmt"
	"log"

	"google.golang.org/grpc"
)

func main() {
	addr := fmt.Sprintf("%s:%s", host, port)
	conn, err := googlegrpc.Dial(addr, googlegrpc.WithInsecure())
	if err != nil {
		log.Fatalf("No es posible conectarse: %v", err)
	}
	client := wishgrpc.NewServerLogisticaClient(conn)

	// uso cliente
	w := &grpc.Compras{
		Name:   name,
		Status: grpc.Compras_ACTIVE,
	}

	res, err := client.CPyme(ctx, &grpc.ComprarEnPyme{Compras: C})
	fmt.Println(res)
	fmt.Println(err)
}
