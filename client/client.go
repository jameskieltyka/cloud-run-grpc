package main

import (
	"context"
	"flag"
	"log"
	"time"

	internal "github.com/jkieltyka/cloud-run-grpc/proto/inter"
	"google.golang.org/grpc"
)

var (
	serverAddr = flag.String("serv_addr", "internal-inbshrhksa-uc.a.run.app:443", "The internal service address")
)

func main() {
	opts := grpc.WithInsecure()
	conn, err := grpc.Dial(*serverAddr, opts)
	defer conn.Close()
	if err != nil {
		log.Printf("unable to dial GRPC")
		return
	}

	//Create Context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//Create client and call
	client := internal.NewInternalServiceClient(conn)
	entryMessage, err := client.GetInternalData(ctx, &internal.InternalDataParameters{
		Id:   "10",
		Name: "James",
	})

	log.Println(err, entryMessage)
}
