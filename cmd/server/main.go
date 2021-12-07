package main

import (
	"database/sql"
	"flag"
	"log"
	"net"
	"os"

	minigames "github.com/wthsths/minigames/internal"
	"github.com/wthsths/minigames/internal/repository"
	"github.com/wthsths/minigames/internal/repository/inmem"
	"github.com/wthsths/minigames/internal/repository/mysql"
	"github.com/wthsths/minigames/pb"
	"google.golang.org/grpc"
)

var config struct {
	addr    string
	repo    string
	connStr string
}

func init() {
	config.addr = os.Getenv("ADDR")
	config.repo = os.Getenv("REPO")
	config.connStr = os.Getenv("CONN_STR")

	// set defaults
	if config.addr == "" {
		config.addr = ":3001"
	}

	if config.repo == "" {
		config.repo = "inmem"
	}

	if config.connStr == "" {
		config.connStr = "root:@tcp(127.0.0.1:3306)/test"
	}
}

func main() {
	flag.StringVar(&config.addr, "addr", config.addr, "grpc listen addr")
	flag.StringVar(&config.repo, "repo", config.repo, "service repository")
	flag.StringVar(&config.connStr, "connStr", config.connStr, "repository connection string")

	flag.Parse()

	var r repository.Repository

	switch config.repo {
	case "inmem":
		r, _ = inmem.NewInmemRepository()
	case "mysql":
		conn, err := sql.Open(config.repo, config.connStr)
		if err != nil {
			panic(err)
		}
		defer conn.Close()

		r, err = mysql.NewMySQLRepository(conn)
		if err != nil {
			panic(err)
		}
	}

	srv, _ := minigames.NewServer(r)

	server := grpc.NewServer()

	listen, err := net.Listen("tcp", config.addr)
	if err != nil {
		log.Panic(err)
	}

	pb.RegisterGameServiceServer(server, srv)
	if err := server.Serve(listen); err != nil {
		log.Panic(err)
	}
}
