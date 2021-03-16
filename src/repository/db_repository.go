package db

type DbRepository interface {
}

type dbRepository struct {
}

func NewRepository() DbRepository {
	return &dbRepository{}
}
