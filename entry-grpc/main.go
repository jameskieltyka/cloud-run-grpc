package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	entry "github.com/jkieltyka/cloud-run-grpc/proto/entry"
	internal "github.com/jkieltyka/cloud-run-grpc/proto/inter"
	"google.golang.org/grpc"
)

var (
	port       = os.Getenv("PORT")
	serverAddr = os.Getenv("INTERNAL_SERVICE")
)

// type entryService struct {
// 	prefix string
// }

// func (e *entryService) CallEntry(c context.Context, entryRequest *entry.EntryDataParameter) (*entry.EntryResult, error) {
// 	// Initialise Return
// 	result := &entry.EntryResult{}

// 	//Dial GRPC
// 	opts := grpc.WithInsecure()
// 	conn, err := grpc.Dial(*serverAddr, opts)
// 	defer conn.Close()
// 	if err != nil {
// 		log.Printf("unable to dial GRPC")
// 		return nil, nil
// 	}

// 	//Create Context
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	//Create client and call
// 	client := internal.NewInternalServiceClient(conn)
// 	internalMessage, err := client.GetInternalData(ctx, &internal.InternalDataParameters{
// 		Id:   entryRequest.Id,
// 		Name: entryRequest.Name,
// 	})

// 	result.Message = e.prefix + internalMessage.Message

// 	return result, nil
// }

func callService(w http.ResponseWriter, r *http.Request) {
	result := &entry.EntryResult{}
	opts := grpc.WithInsecure()
	conn, err := grpc.Dial(serverAddr, opts)
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
	internalMessage, err := client.GetInternalData(ctx, &internal.InternalDataParameters{
		Id:   "test",
		Name: "test",
	})

	result.Message = "from the REST handler -->" + internalMessage.Message

	fmt.Fprintf(w, result.Message)
}

// func createAndRegisterRPC() {
// 	flag.Parse()
// 	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

// 	if err != nil {
// 		log.Fatalf("failed to listen: %v", err)
// 	}

// 	grpcServer := grpc.NewServer()
// 	entry.RegisterEntryServiceServer(grpcServer, &entryService{prefix: "Called from the entry service -> "})
// 	grpcServer.Serve(lis)
// }

func main() {
	// go createAndRegisterRPC()
	http.HandleFunc("/", callService)
	address := ":" + port
	http.ListenAndServe(address, nil)
}
