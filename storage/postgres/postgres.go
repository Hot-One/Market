package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"

	"market/config"
	"market/storage"
)

type store struct {
	db                     *pgxpool.Pool
	category               *CategoryRepo
	branch                 *BranchRepo
	product                *ProductRepo
	storage_coming         *StorageComingRepo
	storage_coming_product *StorageComingProductRepo
}

func NewConnectionPostgres(cfg *config.Config) (storage.StorageI, error) {

	connect, err := pgxpool.ParseConfig(fmt.Sprintf(
		"host=%s user=%s dbname=%s password=%s port=%d sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresUser,
		cfg.PostgresDatabase,
		cfg.PostgresPassword,
		cfg.PostgresPort,
	))

	if err != nil {
		return nil, err
	}
	connect.MaxConns = cfg.PostgresMaxConnection

	pgxpool, err := pgxpool.ConnectConfig(context.Background(), connect)
	if err != nil {
		return nil, err
	}

	return &store{
		db: pgxpool,
	}, nil
}

func (s *store) Close() {
	s.db.Close()
}

func (s *store) Branch() storage.BranchRepoI {

	if s.branch == nil {
		s.branch = NewBranchRepo(s.db)
	}

	return s.branch
}

func (s *store) Category() storage.CategoryRepoI {

	if s.category == nil {
		s.category = NewCategoryRepo(s.db)
	}

	return s.category
}

func (s *store) Product() storage.ProductRepoI {

	if s.product == nil {
		s.product = NewProductRepo(s.db)
	}

	return s.product
}

func (s *store) StorageComing() storage.StorageComingRepoI {

	if s.storage_coming == nil {
		s.storage_coming = NewStorageComingRepo(s.db)
	}

	return s.storage_coming
}

func (s *store) StorageComingProduct() storage.StorageComingProductRepoI {

	if s.storage_coming_product == nil {
		s.storage_coming_product = NewStorageComingProductRepo(s.db)
	}

	return s.storage_coming_product
}
