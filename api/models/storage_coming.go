package models

type StorageComingPrimaryKey struct {
	Id string `json:"id"`
}

type CreateStorageComing struct {
	ComingId string `json:"coming_id"`
	BranchId string `json:"branch_id"`
}

type StorageComing struct {
	Id        string `json:"id"`
	ComingId  string `json:"coming_id"`
	BranchId  string `json:"branch_id"`
	Status    string `json:"status"`
	DateTime  string `json:"date_time"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UpdateStorageComing struct {
	Id       string `json:"id"`
	ComingId string `json:"coming_id"`
	BranchId string `json:"branch_id"`
	Status   string `json:"status"`
}

type StorageComingGetListRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type StorageComingGetListResponse struct {
	Count          int              `json:"count"`
	StorageComings []*StorageComing `json:"storagecomings"`
}
