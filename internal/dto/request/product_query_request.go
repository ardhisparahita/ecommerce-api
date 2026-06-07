package request

type ProductQueryRequest struct {
	Page       int    `query:"page"`
	Limit      int    `query:"limit"`
	Search     string `query:"search"`
	CategoryID uint64 `query:"category_id"`

	SortBy string `query:"sort_by"`
	Order  string `query:"order"`
}
