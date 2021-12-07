package main

import (
	"flag"
	"log"
	"net"

	minigames "github.com/wthsths/minigames/internal"
	"github.com/wthsths/minigames/internal/repository"
	"github.com/wthsths/minigames/internal/repository/inmem"
	"github.com/wthsths/minigames/pb"
	"google.golang.org/grpc"
)

func main() {
	var (
		addr = flag.String("addr", ":3001", "grpc listen addr")
		repo = flag.String("repo", "inmem", "service repository")
	)
	flag.Parse()

	var r repository.Repository

	switch *repo {
	case "inmem":
		r, _ = inmem.NewInmemRepository()
	}

	srv, _ := minigames.NewServer(r)

	server := grpc.NewServer()

	listen, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Panic(err)
	}

	pb.RegisterGameServiceServer(server, srv)
	if err := server.Serve(listen); err != nil {
		log.Panic(err)
	}
}
