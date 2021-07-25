package data_migration

import (
	"context"
	"io"
)

type DataMigration struct {
	sds DBService
}

func Create(sds DBService) *DataMigration {
	return &DataMigration{sds: sds}
}

type DBService interface {
	FetchSchema(context.Context) error
	RetrieveData(context.Context) error
	CreateSchema(context.Context) error
	Insert(context.Context) error
}

func (w *DataMigration) Run(ctx context.Context, r io.Reader) error {
	err := w.sds.FetchSchema(ctx)
	if err != nil {
		return err
	}
	return nil
}
