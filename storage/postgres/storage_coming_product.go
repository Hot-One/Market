package postgres

import (
	"context"
	"database/sql"
	"fmt"

	uuid "github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"

	"market/api/models"
	"market/pkg/helper"
)

type StorageComingProductRepo struct {
	db *pgxpool.Pool
}

func NewStorageComingProductRepo(db *pgxpool.Pool) *StorageComingProductRepo {
	return &StorageComingProductRepo{
		db: db,
	}
}

func (r *StorageComingProductRepo) Create(ctx context.Context, req *models.CreateStorageComingProduct) (string, error) {

	var (
		id         = uuid.New().String()
		query      string
		totalprice = req.Price * req.Quantity
	)

	query = `
		INSERT INTO income_products(id, name, quantity, price, total_price, category_id, storage_coming_id, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, NOW())
	`

	_, err := r.db.Exec(ctx, query,
		id,
		req.Name,
		req.Quantity,
		req.Price,
		totalprice,
		helper.NewNullString(req.CategoryId),
		helper.NewNullString(req.StorageComingId),
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *StorageComingProductRepo) GetByID(ctx context.Context, req *models.StorageComingProductPrimaryKey) (*models.StorageComingProduct, error) {

	var (
		query string

		Id              sql.NullString
		Name            sql.NullString
		Quantity        sql.NullInt32
		Price           sql.NullInt32
		TotalPrice      sql.NullInt32
		CategoryId      sql.NullString
		StorageComingId sql.NullString
		CreatedAt       sql.NullString
		UpdatedAt       sql.NullString
	)

	query = `
		SELECT
			id,
			name,
			status,
			date_time,
			total_price,
			category_id,
			storage_coming_id,
			created_at,
			updated_at
		FROM income_products
		WHERE id = $1
	`

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&Id,
		&Name,
		&Quantity,
		&Price,
		&TotalPrice,
		&CategoryId,
		&StorageComingId,
		&CreatedAt,
		&UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &models.StorageComingProduct{
		Id:              Id.String,
		Name:            Name.String,
		Quantity:        Quantity.Int32,
		Price:           Price.Int32,
		TotalPrice:      TotalPrice.Int32,
		CategoryId:      CategoryId.String,
		StorageComingId: StorageComingId.String,
		CreatedAt:       CreatedAt.String,
		UpdatedAt:       UpdatedAt.String,
	}, nil
}

func (r *StorageComingProductRepo) GetList(ctx context.Context, req *models.StorageComingProductGetListRequest) (*models.StorageComingProductGetListResponse, error) {

	var (
		resp   = &models.StorageComingProductGetListResponse{}
		query  string
		where  = " WHERE TRUE"
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			id,
			name,
			status,
			date_time,
			total_price,
			category_id,
			storage_coming_id,
			created_at,
			updated_at
		FROM income_products
	`

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	if req.Search != "" {
		where += ` AND title ILIKE '%' || '` + req.Search + `' || '%'`
	}

	query += where + offset + limit

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			Id              sql.NullString
			Name            sql.NullString
			Quantity        sql.NullInt32
			Price           sql.NullInt32
			TotalPrice      sql.NullInt32
			CategoryId      sql.NullString
			StorageComingId sql.NullString
			CreatedAt       sql.NullString
			UpdatedAt       sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&Id,
			&Name,
			&Quantity,
			&Price,
			&TotalPrice,
			&CategoryId,
			&StorageComingId,
			&CreatedAt,
			&UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		resp.StorageComingProducts = append(resp.StorageComingProducts, &models.StorageComingProduct{
			Id:              Id.String,
			Name:            Name.String,
			Quantity:        Quantity.Int32,
			Price:           Price.Int32,
			TotalPrice:      TotalPrice.Int32,
			CategoryId:      CategoryId.String,
			StorageComingId: StorageComingId.String,
			CreatedAt:       CreatedAt.String,
			UpdatedAt:       UpdatedAt.String,
		})
	}

	return resp, nil
}

func (r *StorageComingProductRepo) Update(ctx context.Context, req *models.UpdateStorageComingProduct) (int64, error) {

	var (
		query      string
		params     map[string]interface{}
		totalprice = req.Price * req.Quantity
	)

	query = `
		UPDATE
			income_products
		SET
			name = :name,
			quantity = :quantity,
			price = :price,
			total_price = :total_price
			category_id = :category_id
			storage_coming_id = :storage_coming_id,
			updated_at = NOW()
		WHERE id = :id
	`

	params = map[string]interface{}{
		"id":                req.Id,
		"name":              req.Name,
		"quantity":          req.Quantity,
		"price":             req.Price,
		"total_price":       totalprice,
		"category_id":       helper.NewNullString(req.CategoryId),
		"storage_coming_id": helper.NewNullString(req.StorageComingId),
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *StorageComingProductRepo) Delete(ctx context.Context, req *models.StorageComingProductPrimaryKey) error {

	_, err := r.db.Exec(ctx, "DELETE FROM income_products WHERE id = $1", req.Id)
	if err != nil {
		return err
	}

	return nil
}
