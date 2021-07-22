package service

import (
	"context"
	"github.com/Taurin190/go-data-migration/data_migration"
)

type MySQLService struct {
	service data_migration.DBService
}

func (m *MySQLService) FetchSchema(context.Context) error {
	return nil
}

func (m *MySQLService) RetrieveData(context.Context) error {
	return nil
}

func (m *MySQLService) CreateSchema(context.Context) error {
	return nil
}

func (m *MySQLService) Insert(context.Context) error {
	return nil
}

func CreateMySQLService() data_migration.DBService {
	return &MySQLService{}
}
