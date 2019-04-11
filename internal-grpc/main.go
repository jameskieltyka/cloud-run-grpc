package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	internal "github.com/jkieltyka/cloud-run-grpc/proto/inter"
	"google.golang.org/grpc"
)

var (
	port, _    = strconv.ParseInt(os.Getenv("PORT"), 10, 64)
	serverAddr = flag.String("serv_addr", "internal.svc", "The internal service address")
)

type internalService struct{}

func (i *internalService) GetInternalData(c context.Context, requestData *internal.InternalDataParameters) (*internal.InternalData, error) {
	return &internal.InternalData{Message: "Hello " + requestData.Name + " from the Internal Data Service!"}, nil
}

func createAndRegisterRPC() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	internal.RegisterInternalServiceServer(grpcServer, &internalService{})
	grpcServer.Serve(lis)
}

func main() {
	createAndRegisterRPC()
}
