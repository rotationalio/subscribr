package db

type DB struct {
	name string
}

func New(name string) *DB {
	return &DB{name: name}
}

type SQLDB struct {
	DB
}

func NewSQL(name string) *SQLDB {
	return &SQLDB{}
}

func (d *DB) Connect() (err error) {
	return nil
}

func (d *SQLDB) Connect() (err error) {
	return nil
}
