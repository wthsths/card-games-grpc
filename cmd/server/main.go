package main

import (
	"log"
	"net"

	minigames "github.com/wthsths/minigames/internal"
	"github.com/wthsths/minigames/internal/repository/inmem"
	"github.com/wthsths/minigames/pb"
	"google.golang.org/grpc"
)

func main() {
	repo, _ := inmem.NewInmemRepository()
	srv, _ := minigames.NewServer(repo)

	server := grpc.NewServer()

	listen, err := net.Listen("tcp", ":3001")
	if err != nil {
		log.Panic("listen error")
	}

	pb.RegisterGameServiceServer(server, srv)
	if err := server.Serve(listen); err != nil {
		log.Panic(err)
	}
}
