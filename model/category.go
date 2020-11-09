package model

type Category struct {
	ID        int
	ParentID  int
	Name      string
	Status    byte
	SortOrder int8
	IsParent  byte
	Created   string
	Updated   string
}
