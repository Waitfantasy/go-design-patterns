package factory

import "testing"

func TestFactory(t *testing.T)  {
	factory := &AbstractDatabaseFactory{}
	d1 := factory.CreateDriver(MysqlDriverName)
	d1.Connection()
	d1.Close()
	d2 := factory.CreateDriver(SqlLiteDriverName)
	d2.Connection()
	d2.Close()
}
