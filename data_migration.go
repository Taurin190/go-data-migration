package go_data_migration

type DataMigration struct {
}

func Create() *DataMigration {
	return &DataMigration{}
}

func (w *DataMigration) Run() error {
	return nil
}
