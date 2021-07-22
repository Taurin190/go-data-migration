package data_migration

import "context"

type DataMigration struct {
}

func Create() *DataMigration {
	return &DataMigration{}
}

type DBService interface {
	FetchSchema(context.Context) error
	RetrieveData(context.Context) error
	CreateSchema(context.Context) error
	Insert(context.Context) error
}

func (w *DataMigration) Run() error {
	return nil
}
