package dto

type Response[T any] struct {
	Name    string
	Records []T
	Record  T
}
