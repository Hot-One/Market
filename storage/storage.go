package storage

import (
	"context"

	"market/api/models"
)

type StorageI interface {
	Close()
	Branch() BranchRepoI
	Category() CategoryRepoI
	Product() ProductRepoI
	StorageComing() StorageComingRepoI
	StorageComingProduct() StorageComingProductRepoI
}

type BranchRepoI interface {
	Create(context.Context, *models.CreateBranch) (string, error)
	GetByID(context.Context, *models.BranchPrimaryKey) (*models.Branch, error)
	GetList(context.Context, *models.BranchGetListRequest) (*models.BranchGetListResponse, error)
	Update(context.Context, *models.UpdateBranch) (int64, error)
	Delete(context.Context, *models.BranchPrimaryKey) error
}

type CategoryRepoI interface {
	Create(context.Context, *models.CreateCategory) (string, error)
	GetByID(context.Context, *models.CategoryPrimaryKey) (*models.Category, error)
	GetList(context.Context, *models.CategoryGetListRequest) (*models.CategoryGetListResponse, error)
	Update(context.Context, *models.UpdateCategory) (int64, error)
	Delete(context.Context, *models.CategoryPrimaryKey) error
}

type ProductRepoI interface {
	Create(context.Context, *models.CreateProduct) (string, error)
	GetByID(context.Context, *models.ProductPrimaryKey) (*models.Product, error)
	GetList(context.Context, *models.ProductGetListRequest) (*models.ProductGetListResponse, error)
	Update(context.Context, *models.UpdateProduct) (int64, error)
	Patch(context.Context, *models.PatchRequest) (int64, error)
	Delete(context.Context, *models.ProductPrimaryKey) error
}

type StorageComingRepoI interface {
	Create(context.Context, *models.CreateStorageComing) (string, error)
	GetByID(context.Context, *models.StorageComingPrimaryKey) (*models.StorageComing, error)
	GetList(context.Context, *models.StorageComingGetListRequest) (*models.StorageComingGetListResponse, error)
	Update(context.Context, *models.UpdateStorageComing) (int64, error)
	Delete(context.Context, *models.StorageComingPrimaryKey) error
}

type StorageComingProductRepoI interface {
}
