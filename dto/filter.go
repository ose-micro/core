package dto

const (
	EQUAL              Operator = "EQUAL"
	LIKE               Operator = "LIKE"
	BETWEEN            Operator = "BETWEEN"
	BEFORE             Operator = "BEFORE"
	AFTER              Operator = "AFTER"
	GREATER_THAN       Operator = "GREATER_THAN"
	LESS_THAN          Operator = "LESS_THAN"
	GREATER_THAN_EQUAL Operator = "GREATER_THAN_EQUAL"
	LESS_THAN_EQUAL    Operator = "LESS_THAN_EQUAL"
	DATE_EQUAL         Operator = "DATE_EQUAL"
	DATE_BETWEEN       Operator = "DATE_BETWEEN"
)

type Operator string

type Filter struct {
	Field    string
	Operator Operator
	Value    interface{}
}
