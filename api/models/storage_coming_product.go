package models

type StorageComingProductPrimaryKey struct {
	Id string `json:"id"`
}

type CreateStorageComingProduct struct {
	Name            string `json:"name"`
	Quantity        int32  `json:"status"`
	Price           int32  `json:"date_time"`
	CategoryId      string `json:"category_id"`
	StorageComingId string `json:"storage_coming_id"`
}

type StorageComingProduct struct {
	Id              string `json:"id"`
	Name            string `json:"name"`
	Quantity        int32  `json:"status"`
	Price           int32  `json:"date_time"`
	TotalPrice      int32  `json:"total_price"`
	CategoryId      string `json:"category_id"`
	StorageComingId string `json:"storage_coming_id"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
}

type UpdateStorageComingProduct struct {
	Id              string `json:"id"`
	Name            string `json:"name"`
	Quantity        int32  `json:"status"`
	Price           int32  `json:"date_time"`
	CategoryId      string `json:"category_id"`
	StorageComingId string `json:"storage_coming_id"`
}

type StorageComingProductGetListRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type StorageComingProductGetListResponse struct {
	Count                 int                     `json:"count"`
	StorageComingProducts []*StorageComingProduct `json:"storagecomingproducts"`
}
