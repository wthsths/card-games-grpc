package mysql

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"github.com/wthsths/minigames/pb"
)

type mysqlRepository struct {
	db *sql.DB
}

func NewMySQLRepository(db *sql.DB) (*mysqlRepository, error) {
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &mysqlRepository{
		db: db,
	}, nil
}

func (repo *mysqlRepository) CreateGame(ctx context.Context, game *pb.Game) error {
	if err := repo.db.PingContext(ctx); err != nil {
		return err
	}

	stmt, err := repo.db.PrepareContext(ctx, "INSERT INTO minigames(uuid, player, type, status, data) VALUES(?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, game.Uuid, game.Player, game.Type, game.Status, string(game.Data))

	return err
}

func (repo *mysqlRepository) GetGameByUuid(ctx context.Context, uuid string) (*pb.Game, error) {
	if err := repo.db.PingContext(ctx); err != nil {
		return nil, err
	}

	sql := `SELECT uuid, player, type, status, data FROM minigames WHERE uuid = ?`
	row := repo.db.QueryRowContext(ctx, sql, uuid)

	var game pb.Game
	switch err := row.Scan(&game.Uuid, &game.Player, &game.Type, &game.Status, &game.Data); err {
	case nil:
		return &game, nil
	default:
		return nil, err
	}
}

func (repo *mysqlRepository) UpdateByUuid(ctx context.Context, uuid string, game *pb.Game) (err error) {
	if err := repo.db.PingContext(ctx); err != nil {
		return err
	}

	sql := `UPDATE minigames SET uuid = ?, player = ?, type = ?, status = ?, data = ? WHERE uuid = ?`
	_, err = repo.db.ExecContext(ctx, sql, &game.Uuid, &game.Player, &game.Type, &game.Status, &game.Data, uuid)

	return err
}

func (repo *mysqlRepository) GetGamesByPlayer(ctx context.Context, player uint64) (games []*pb.Game, err error) {
	if err := repo.db.PingContext(ctx); err != nil {
		return nil, err
	}

	sql := `SELECT uuid, player, type, status, data FROM minigames WHERE player = ?`
	rows, err := repo.db.QueryContext(ctx, sql, player)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var game pb.Game
		err := rows.Scan(&game.Uuid, &game.Player, &game.Type, &game.Status, &game.Data)
		if err != nil {
			return nil, err
		}

		games = append(games, &game)
	}

	return games, nil
}
