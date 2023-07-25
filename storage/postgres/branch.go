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

type BranchRepo struct {
	db *pgxpool.Pool
}

func NewBranchRepo(db *pgxpool.Pool) *BranchRepo {
	return &BranchRepo{
		db: db,
	}
}

func (r *BranchRepo) Create(ctx context.Context, req *models.CreateBranch) (string, error) {

	var (
		id    = uuid.New().String()
		query string
	)

	query = `
		INSERT INTO branch(id, name, address, phone_number, updated_at)
		VALUES ($1, $2, $3, $4,NOW())
	`

	_, err := r.db.Exec(ctx, query,
		id,
		req.Name,
		req.Address,
		req.PhoneNumber,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *BranchRepo) GetByID(ctx context.Context, req *models.BranchPrimaryKey) (*models.Branch, error) {

	var (
		query string

		id           sql.NullString
		name         sql.NullString
		address      sql.NullString
		phone_number sql.NullString
		createdAt    sql.NullString
		updatedAt    sql.NullString
	)

	query = `
		SELECT
			id,
			name,
			address,
			phone_number,
			created_at,
			updated_at
		FROM branch
		WHERE id = $1
	`

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&name,
		&address,
		&phone_number,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &models.Branch{
		Id:          id.String,
		Name:        name.String,
		Address:     address.String,
		PhoneNumber: phone_number.String,
		CreatedAt:   createdAt.String,
		UpdatedAt:   updatedAt.String,
	}, nil
}

func (r *BranchRepo) GetList(ctx context.Context, req *models.BranchGetListRequest) (*models.BranchGetListResponse, error) {

	var (
		resp   = &models.BranchGetListResponse{}
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
			address,
			phone_number,
			created_at,
			updated_at
		FROM branch
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
			id           sql.NullString
			name         sql.NullString
			address      sql.NullString
			phone_number sql.NullString
			createdAt    sql.NullString
			updatedAt    sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&name,
			&address,
			&phone_number,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return nil, err
		}

		resp.Branches = append(resp.Branches, &models.Branch{
			Id:          id.String,
			Name:        name.String,
			Address:     address.String,
			PhoneNumber: phone_number.String,
			CreatedAt:   createdAt.String,
			UpdatedAt:   updatedAt.String,
		})
	}

	return resp, nil
}

func (r *BranchRepo) Update(ctx context.Context, req *models.UpdateBranch) (int64, error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
			branch
		SET
			name = :name,
			address = :address,
			phone_number = :phone_number,
			updated_at = NOW()
		WHERE id = :id
	`

	params = map[string]interface{}{
		"id":           req.Id,
		"name":         req.Name,
		"address":      req.Address,
		"phone_number": req.PhoneNumber,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *BranchRepo) Delete(ctx context.Context, req *models.BranchPrimaryKey) error {

	_, err := r.db.Exec(ctx, "DELETE FROM branch WHERE id = $1", req.Id)
	if err != nil {
		return err
	}

	return nil
}
