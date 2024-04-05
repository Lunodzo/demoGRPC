package main

import (
	"context"
	"demo/grpc/invoice"
	"log"
	"net"

	"google.golang.org/grpc"
)

type myInvoiceServer struct {
	invoice.UnimplementedInvoiceServer
}

func (s myInvoiceServer ) Create(ctx context.Context,req *invoice.CreateRequest) (*invoice.CreateResponse, error){
	return &invoice.CreateResponse{
		Pdf: []byte(req.From),
		Docx: []byte("DOCX"),
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":8089")
	if err != nil {
		log.Fatal(err)
	}

	serverRegister := grpc.NewServer()
	service := &myInvoiceServer{}

	invoice.RegisterInvoiceServer(serverRegister, service)

	err = serverRegister.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
}