package postgres

import "github.com/robowealth-mutual-fund/blueprint-emitter-service/internals/infrastructure/database"

type PostgresRepository struct {
	db *database.DB
}

func NewRepository(db *database.DB) Repository {
	return &PostgresRepository{db: db}
}
