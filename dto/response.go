package dto

type Response[T any] struct {
	Records []T
	Record  T
}
