package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	entry "github.com/jkieltyka/cloud-run-grpc/proto/entry"
	internal "github.com/jkieltyka/cloud-run-grpc/proto/inter"
	"google.golang.org/grpc"
)

var (
	port       = flag.Int("port", 8080, "The server port")
	serverAddr = flag.String("serv_addr", "127.0.0.1:8081", "The internal service address")
)

type entryService struct {
	prefix string
}

func (e *entryService) CallEntry(c context.Context, entryRequest *entry.EntryDataParameter) (*entry.EntryResult, error) {
	// Initialise Return
	result := &entry.EntryResult{}

	//Dial GRPC
	opts := grpc.WithInsecure()
	conn, err := grpc.Dial(*serverAddr, opts)
	defer conn.Close()
	if err != nil {
		log.Printf("unable to dial GRPC")
		return nil, nil
	}

	//Create Context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//Create client and call
	client := internal.NewInternalServiceClient(conn)
	internalMessage, err := client.GetInternalData(ctx, &internal.InternalDataParameters{
		Id:   entryRequest.Id,
		Name: entryRequest.Name,
	})

	result.Message = e.prefix + internalMessage.Message

	return result, nil
}

func createAndRegisterRPC() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	entry.RegisterEntryServiceServer(grpcServer, &entryService{prefix: "Called from the entry service -> "})
	grpcServer.Serve(lis)
}

func main() {
	createAndRegisterRPC()
}
