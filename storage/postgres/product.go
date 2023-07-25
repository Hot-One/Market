package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	uuid "github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"

	"market/api/models"
	"market/pkg/helper"
)

type ProductRepo struct {
	db *pgxpool.Pool
}

func NewProductRepo(db *pgxpool.Pool) *ProductRepo {
	return &ProductRepo{
		db: db,
	}
}

func (r *ProductRepo) Create(ctx context.Context, req *models.CreateProduct) (string, error) {

	var (
		id    = uuid.New().String()
		query string
	)

	query = `
		INSERT INTO product(id, name, barcode, price, category_id, updated_at)
		VALUES ($1, $2, $3, $4, $5, NOW())
	`

	_, err := r.db.Exec(ctx, query,
		id,
		req.Name,
		req.Barcode,
		req.Price,
		helper.NewNullString(req.CategoryId),
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *ProductRepo) GetByID(ctx context.Context, req *models.ProductPrimaryKey) (*models.Product, error) {

	var (
		query string

		id         sql.NullString
		name       sql.NullString
		barcode    sql.NullString
		price      sql.NullInt32
		categoryId sql.NullString
		createdAt  sql.NullString
		updatedAt  sql.NullString
	)

	query = `
		SELECT
			id,
			name,
			barcode,
			price,
			category_id,
			created_at,
			updated_at
		FROM product
		WHERE id = $1
	`

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&name,
		&barcode,
		&price,
		&categoryId,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &models.Product{
		Id:         id.String,
		Name:       name.String,
		Barcode:    barcode.String,
		Price:      price.Int32,
		CategoryId: categoryId.String,
		CreatedAt:  createdAt.String,
		UpdatedAt:  updatedAt.String,
	}, nil
}

func (r *ProductRepo) GetList(ctx context.Context, req *models.ProductGetListRequest) (*models.ProductGetListResponse, error) {

	var (
		resp   = &models.ProductGetListResponse{}
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
			barcode,
			price,
			category_id,
			created_at,
			updated_at
		FROM product
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
			id         sql.NullString
			name       sql.NullString
			barcode    sql.NullString
			price      sql.NullInt32
			categoryId sql.NullString
			createdAt  sql.NullString
			updatedAt  sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&name,
			&barcode,
			&price,
			&categoryId,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return nil, err
		}

		resp.Products = append(resp.Products, &models.Product{
			Id:         id.String,
			Name:       name.String,
			Barcode:    barcode.String,
			Price:      price.Int32,
			CategoryId: categoryId.String,
			CreatedAt:  createdAt.String,
			UpdatedAt:  updatedAt.String,
		})
	}

	return resp, nil
}

func (r *ProductRepo) Update(ctx context.Context, req *models.UpdateProduct) (int64, error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
			product
		SET
			name = :name,
			barcode = :barcode,
			price = :price,
			category_id = :category_id,
			updated_at = NOW()
		WHERE id = :id
	`

	params = map[string]interface{}{
		"id":          req.Id,
		"name":        req.Name,
		"barcode":     req.Barcode,
		"price":       req.Price,
		"category_id": helper.NewNullString(req.CategoryId),
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *ProductRepo) Patch(ctx context.Context, req *models.PatchRequest) (int64, error) {

	var (
		query string
		set   string
	)

	if len(req.Fields) <= 0 {
		return 0, errors.New("no fields")
	}

	for key := range req.Fields {
		set += fmt.Sprintf(" %s = :%s, ", key, key)
	}

	query = `
		UPDATE
			product
		SET ` + set + ` updated_at = now()
		WHERE id = :id
	`

	req.Fields["id"] = req.ID

	fmt.Println(query)

	query, args := helper.ReplaceQueryParams(query, req.Fields)
	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *ProductRepo) Delete(ctx context.Context, req *models.ProductPrimaryKey) error {

	_, err := r.db.Exec(ctx, "DELETE FROM product WHERE id = $1", req.Id)
	if err != nil {
		return err
	}

	return nil
}
