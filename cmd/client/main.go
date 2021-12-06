package main

import (
	"context"
	"log"

	"github.com/wthsths/minigames/pb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":3001", grpc.WithInsecure())
	if err != nil {
		log.Panic(err)
	}
	defer conn.Close()

	client := pb.NewGameServiceClient(conn)

	ctx := context.Background()

	createGame, _ := client.CreateGame(ctx, &pb.CreateGameRequest{})
	log.Printf("%+v\n", createGame)

	dealResp, _ := client.PlayGame(ctx, &pb.PlayGameRequest{
		Action: pb.PlayGameRequest_ACTION_DEAL,
		Uuid:   createGame.Game.Uuid,
	})
	log.Printf("%+v\n", dealResp)

	foldResp, _ := client.PlayGame(ctx, &pb.PlayGameRequest{
		Action: pb.PlayGameRequest_ACTION_FOLD,
		Uuid:   createGame.Game.Uuid,
	})
	log.Printf("%+v\n", foldResp)
}
