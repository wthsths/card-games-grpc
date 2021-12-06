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

	// create game
	req := &pb.CreateGameRequest{
		Player: 1,
		Type:   pb.GameType_GAME_TYPE_TEXAS_HOLDEM_BONUS,
	}

	resp, err := client.CreateGame(context.Background(), req)
	if err != nil {
		log.Fatalf("error resp: %v", err)
	}
	log.Println(resp)

	getGameReq := &pb.GetGameByUuidRequest{
		Uuid: resp.Game.Uuid,
	}
	getGameResp, err := client.GetGameByUuid(context.Background(), getGameReq)
	if err != nil {
		log.Fatalf("error get game %v", err)
	}
	log.Printf("get game resp %v\n", getGameResp)

	playGameResp, err := client.PlayGame(context.Background(), &pb.PlayGameRequest{
		Action: pb.PlayGameRequest_ACTION_DEAL,
		Uuid:   resp.GetGame().Uuid,
	})
	if err != nil {
		log.Fatalf("play game err: %v\n", err)
	}
	log.Printf("play game resp %v", playGameResp)

	resp3, err := client.PlayGame(context.Background(), &pb.PlayGameRequest{
		Action: pb.PlayGameRequest_ACTION_PLAY,
		Uuid:   resp.GetGame().Uuid,
	})
	if err != nil {
		log.Fatalf("play game err: %v\n", err)
	}
	log.Printf("resp3 %v", resp3)

	resp4, err := client.PlayGame(context.Background(), &pb.PlayGameRequest{
		Action: pb.PlayGameRequest_ACTION_CHECK,
		Uuid:   resp.GetGame().Uuid,
	})
	if err != nil {
		log.Fatalf("play2 game err: %v", err)
	}
	log.Printf("resp4 %v", resp4)

	resp5, err := client.PlayGame(context.Background(), &pb.PlayGameRequest{
		Action: pb.PlayGameRequest_ACTION_CHECK,
		Uuid:   resp.GetGame().Uuid,
	})
	if err != nil {
		log.Fatalf("%v", err)
	}
	log.Println(resp5)
}
