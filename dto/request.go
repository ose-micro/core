package dto

type Pagination struct {
	Page  int
	Limit int
}

type Request struct {
	Pagination *Pagination
	Filter     []Filter
	Sort       []Sort
}
