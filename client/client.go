package main

import (
	"context"
	"flag"
	"log"
	"time"

	entry "github.com/jkieltyka/cloud-run-grpc/proto/entry"
	"google.golang.org/grpc"
)

var (
	serverAddr = flag.String("serv_addr", "127.0.0.1:8080", "The internal service address")
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
	client := entry.NewEntryServiceClient(conn)
	entryMessage, err := client.CallEntry(ctx, &entry.EntryDataParameter{
		Id:   "10",
		Name: "James",
	})

	log.Println(entryMessage.Message)
}
