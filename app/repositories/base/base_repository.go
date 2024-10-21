package base

import "gorm.io/gorm"

// next step. need generic
// struct doesn't have specific data type, so you can't compare struct{A}=struct{}
// M in actual is struct related model
type AddBase[M any] interface {
	Add(data *M) (any, error)
}

type UpdateBase[I int, M any] interface {
	Update(id *I, data *M) (any, error)
}

type DetailBase[I int, M any] interface {
	Detail(id *I, data *M) (*gorm.DB, error)
}

type DeleteBase[I int, M any] interface {
	Delete(id *I, model *M) error
}

type SearchBase[D map[string]any, M any, R any] interface {
	Search(data *D, model *M) (*R, error)
}
