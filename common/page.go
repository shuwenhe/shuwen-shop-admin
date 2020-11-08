package common

type Page struct {
	Rows  interface{} `json:"rows,omitempty"`
	Total int         `json:"total,omitempty"`
}
