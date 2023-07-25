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

type StorageComingRepo struct {
	db *pgxpool.Pool
}

func NewStorageComingRepo(db *pgxpool.Pool) *StorageComingRepo {
	return &StorageComingRepo{
		db: db,
	}
}

func (r *StorageComingRepo) Create(ctx context.Context, req *models.CreateStorageComing) (string, error) {

	var (
		id    = uuid.New().String()
		query string
	)

	query = `
		INSERT INTO storage_coming(id, coming_id, branch_id, date_time, updated_at)
		VALUES ($1, $2, $3, NOW(), NOW())
	`

	_, err := r.db.Exec(ctx, query,
		id,
		req.ComingId,
		helper.NewNullString(req.BranchId),
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *StorageComingRepo) GetByID(ctx context.Context, req *models.StorageComingPrimaryKey) (*models.StorageComing, error) {

	var (
		query string

		id        sql.NullString
		comingId  sql.NullString
		branchId  sql.NullString
		status    sql.NullString
		datetime  sql.NullString
		createdAt sql.NullString
		updatedAt sql.NullString
	)

	query = `
		SELECT
			id,
			coming_id,
			branch_id,
			status,
			date_time,
			created_at,
			updated_at
		FROM storage_coming
		WHERE id = $1
	`

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&comingId,
		&branchId,
		&status,
		&datetime,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &models.StorageComing{
		Id:        id.String,
		ComingId:  comingId.String,
		BranchId:  branchId.String,
		Status:    status.String,
		DateTime:  datetime.String,
		CreatedAt: createdAt.String,
		UpdatedAt: updatedAt.String,
	}, nil
}

func (r *StorageComingRepo) GetList(ctx context.Context, req *models.StorageComingGetListRequest) (*models.StorageComingGetListResponse, error) {

	var (
		resp   = &models.StorageComingGetListResponse{}
		query  string
		where  = " WHERE TRUE"
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			id,
			coming_id,
			branch_id,
			status,
			date_time,
			created_at,
			updated_at
		FROM storage_coming
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
			id        sql.NullString
			comingId  sql.NullString
			branchId  sql.NullString
			status    sql.NullString
			datetime  sql.NullString
			createdAt sql.NullString
			updatedAt sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&comingId,
			&branchId,
			&status,
			&datetime,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return nil, err
		}

		resp.StorageComings = append(resp.StorageComings, &models.StorageComing{
			Id:        id.String,
			ComingId:  comingId.String,
			BranchId:  branchId.String,
			Status:    status.String,
			DateTime:  datetime.String,
			CreatedAt: createdAt.String,
			UpdatedAt: updatedAt.String,
		})
	}

	return resp, nil
}

func (r *StorageComingRepo) Update(ctx context.Context, req *models.UpdateStorageComing) (int64, error) {

	var (
		query  string
		params map[string]interface{}
	)

	if req.Status == "in process" {
		query = `
			UPDATE
				storage_coming
			SET
				coming_id = :coming_id,
				branch_id = :branch_id,
				status = :status,
				updated_at = NOW()
			WHERE id = :id
		`
	} else if req.Status == "fineshed" {
		query = `
			UPDATE
				storage_coming
			SET
				coming_id = :coming_id,
				branch_id = :branch_id,
				status = :status,
				date_time = NOW(),
				updated_at = NOW()
			WHERE id = :id
	`
	} else {
		return 0, errors.New("There is no such status set status in process or finished !")
	}
	params = map[string]interface{}{
		"id":        req.Id,
		"coming_id": req.ComingId,
		"status":    req.Status,
		"branch_id": helper.NewNullString(req.BranchId),
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *StorageComingRepo) Delete(ctx context.Context, req *models.StorageComingPrimaryKey) error {

	_, err := r.db.Exec(ctx, "DELETE FROM storage_coming WHERE id = $1", req.Id)
	if err != nil {
		return err
	}

	return nil
}
