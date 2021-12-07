package main

import (
	"context"
	"encoding/json"
	"log"

	texasholdembonus "github.com/wthsths/minigames/internal/game/texas-holdem-bonus"
	"github.com/wthsths/minigames/pb"
	"google.golang.org/grpc"
)

type dataStruct struct {
	PlayerHancRankString string `json:"player_hand_rank_string"`
}

func main() {
	conn, err := grpc.Dial(":3001", grpc.WithInsecure())
	if err != nil {
		log.Panic(err)
	}
	defer conn.Close()

	client := pb.NewGameServiceClient(conn)

	ctx := context.Background()

	//FoldAfterDeal(ctx, client)
	PlayAndFoldAfterDeal(ctx, client)
}

func PlayAndFoldAfterDeal(ctx context.Context, client pb.GameServiceClient) {
	createGame, _ := client.CreateGame(ctx, &pb.CreateGameRequest{})

	_, _ = client.PlayGame(ctx, &pb.PlayGameRequest{
		Action: pb.PlayGameRequest_ACTION_DEAL,
		Uuid:   createGame.Game.Uuid,
	})

	_, _ = client.PlayGame(ctx, &pb.PlayGameRequest{
		Action: pb.PlayGameRequest_ACTION_PLAY,
		Uuid:   createGame.Game.Uuid,
	})

	foldResp, err := client.PlayGame(ctx, &pb.PlayGameRequest{
		Action: pb.PlayGameRequest_ACTION_FOLD,
		Uuid:   createGame.Game.Uuid,
	})

	if err != nil {
		log.Printf("err %v", err)
	}

	log.Println(foldResp)
}

func FoldAfterDeal(ctx context.Context, client pb.GameServiceClient) {
	createGame, _ := client.CreateGame(ctx, &pb.CreateGameRequest{})

	_, _ = client.PlayGame(ctx, &pb.PlayGameRequest{
		Action: pb.PlayGameRequest_ACTION_DEAL,
		Uuid:   createGame.Game.Uuid,
	})

	foldResp, _ := client.PlayGame(ctx, &pb.PlayGameRequest{
		Action: pb.PlayGameRequest_ACTION_FOLD,
		Uuid:   createGame.Game.Uuid,
	})

	var dataStruct texasholdembonus.EntityData
	err := json.Unmarshal(foldResp.Game.Data, &dataStruct)
	if err != nil {
		log.Println(err)
	}

	//if dataStruct.PlayerHancRankString != "Straight" &&
	//	dataStruct.PlayerHancRankString != "Flush" &&
	//	dataStruct.PlayerHancRankString != "Straight Flush" &&
	//	dataStruct.PlayerHancRankString != "Four of A Kind" {
	//	time.Sleep(100 * time.Millisecond)
	//
	//		FoldAfterDeal(ctx, client)
	//		return
	//	}

	log.Printf("%+v\n", dataStruct)
}
