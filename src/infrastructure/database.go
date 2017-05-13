package infrastructure

type DbHandler interface {
	Execute(statement string) (int64, error)
	Query(statement string) Row
}

type Row interface {
	Scan(params ...interface{})
	Next() bool
	Close() error
}

type DbRepository struct {
	DbHandlers map[string]DbHandler
	DbHandler  DbHandler
}
