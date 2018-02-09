// consignment-service/main.go
package main

import (
	"fmt"
	"log"
	"os"

	// Import generated from protobuf code
	micro "github.com/micro/go-micro"
	pb "github.com/vbanthia/shippy/consignment-service/proto/consignment"
	vesselProto "github.com/vbanthia/shippy/vessel-service/proto/vessel"
)

const (
	defaultHost = "localhost:27017"
)

func main() {
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = defaultHost
	}

	session, err := CreateSession(host)
	defer session.Close()

	if err != nil {
		log.Panicf("Could not connect to datastore with host %s - %v", host, err)
	}

	srv := micro.NewService(
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
	)

	vesselClient := vesselProto.NewVesselServiceClient("go.micro.srv.vessel", srv.Client())

	srv.Init()

	pb.RegisterShippingServiceHandler(srv.Server(), &service{session, vesselClient})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
