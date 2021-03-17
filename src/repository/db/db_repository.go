package db

type DbRepository interface {
}

type dbRepository struct {
}

func NewDbRepository() DbRepository {
	return &dbRepository{}
}
