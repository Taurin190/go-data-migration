package data_migration

import (
	"context"
	"io"
)

type DataMigration struct {
	sds DBService
	sdb DatabaseInfo
}

type DatabaseInfo struct {
	host     string
	port     int
	database string
	username string
	password string
}

func GetDatabaseInfo(host string, port int, db string, username string, password string) *DatabaseInfo {
	return &DatabaseInfo{
		host:     host,
		port:     port,
		database: db,
		username: username,
		password: password,
	}
}

func Create(sds DBService, sdb DatabaseInfo) *DataMigration {
	return &DataMigration{sds: sds, sdb: sdb}
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
