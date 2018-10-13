package factory

type DriverNameType string

const (
	MysqlDriverName  DriverNameType = "mysql"
	SqlLiteDriverName DriverNameType = "sqllite"
)

type DatabaseFactory interface {
	CreateDriver(name DriverNameType) DatabaseDriver
}

type DatabaseDriver interface {
	Connection()
	Close()
}

type AbstractDatabaseFactory struct {
	driverMap map[string]DatabaseDriver
}

func (a *AbstractDatabaseFactory) CreateDriver(name DriverNameType) DatabaseDriver {
	switch name {
	case MysqlDriverName:
		return NewMysqlDriver()
	case SqlLiteDriverName:
		return NewMysqlDriver()
	default:
		return NewMysqlDriver()
	}
}

type MysqlDriver struct {
}

func NewMysqlDriver() *MysqlDriver {
	return &MysqlDriver{}
}

func (m *MysqlDriver) Connection() {
	println("mysql driver connection")
}

func (m *MysqlDriver) Close() {
	println("mysql driver close")
}

type SqlLiteDriver struct {

}

func (s *SqlLiteDriver) Connection() {
	println("sqllite driver connection")
}

func (s *SqlLiteDriver) Close() {
	println("sqllite driver close")
}
