package database

import (
	"context"
	"database/sql"
)

type MockDB struct {
	MockOpen                  func() error
	MockClose                 func() error
	MockDatabase              func() (string, error)
	MockDatabases             func() ([]string, error)
	MockDatabaseTables        func() (map[string][]string, error)
	MockTables                func() ([]string, error)
	MockDescribeTable         func(string) ([]*ColumnDesc, error)
	MockDescribeDatabaseTable func() ([]*ColumnDesc, error)
	MockExec                  func(context.Context, string) (sql.Result, error)
	MockQuery                 func(context.Context, string) (*sql.Rows, error)
}

func (m *MockDB) Open() error {
	return m.MockOpen()
}

func (m *MockDB) Close() error {
	return m.MockClose()
}

func (m *MockDB) Database() (string, error) {
	return m.MockDatabase()
}

func (m *MockDB) Databases() ([]string, error) {
	return m.MockDatabases()
}

func (m *MockDB) Schema() (string, error) {
	return m.MockDatabase()
}

func (m *MockDB) Schemas() ([]string, error) {
	return m.MockDatabases()
}

func (m *MockDB) DatabaseTables() (map[string][]string, error) {
	return m.MockDatabaseTables()
}

func (m *MockDB) Tables() ([]string, error) {
	return m.MockTables()
}

func (m *MockDB) DescribeTable(tableName string) ([]*ColumnDesc, error) {
	return m.MockDescribeTable(tableName)
}

func (m *MockDB) DescribeDatabaseTable() ([]*ColumnDesc, error) {
	return m.MockDescribeDatabaseTable()
}

func (m *MockDB) Exec(ctx context.Context, query string) (sql.Result, error) {
	return m.MockExec(ctx, query)
}

func (m *MockDB) Query(ctx context.Context, query string) (*sql.Rows, error) {
	return m.MockQuery(ctx, query)
}

func (m *MockDB) SwitchDB(dbName string) error {
	return ErrNotImplementation
}

var dummyDatabases = []string{
	"information_schema",
	"mysql",
	"performance_schema",
	"sys",
	"world",
}
var dummyDatabaseTables = map[string][]string{
	"world": []string{
		"city",
		"country",
		"countrylanguage",
	},
}
var dummyTables = []string{
	"city",
	"country",
	"countrylanguage",
}
var dummyCityColumns = []*ColumnDesc{
	{
		Schema: "world",
		Table:  "city",
		Name:   "ID",
		Type:   "int(11)",
		Null:   "NO",
		Key:    "PRI",
		Default: sql.NullString{
			String: "<null>",
			Valid:  false,
		},
		Extra: "auto_increment",
	},
	{
		Schema: "world",
		Table:  "city",
		Name:   "Name",
		Type:   "char(35)",
		Null:   "NO",
		Key:    "",
		Default: sql.NullString{
			String: "",
			Valid:  false,
		},
		Extra: "",
	},
	{
		Schema: "world",
		Table:  "city",
		Name:   "CountryCode",
		Type:   "char(3)",
		Null:   "NO",
		Key:    "MUL",
		Default: sql.NullString{
			String: "",
			Valid:  false,
		},
		Extra: "",
	},
	{
		Schema: "world",
		Table:  "city",
		Name:   "District",
		Type:   "char(20)",
		Null:   "NO",
		Key:    "",
		Default: sql.NullString{
			String: "",
			Valid:  false,
		},
		Extra: "",
	},
	{
		Schema: "world",
		Table:  "city",
		Name:   "Population",
		Type:   "int(11)",
		Null:   "NO",
		Key:    "",
		Default: sql.NullString{
			String: "",
			Valid:  false,
		},
		Extra: "",
	},
}
var dummyCountryColumns = []*ColumnDesc{
	{
		Schema: "world",
		Table:  "country",
		Name:   "Code",
		Type:   "char(3)",
		Null:   "NO",
		Key:    "PRI",
		Default: sql.NullString{
			String: "<null>",
			Valid:  false,
		},
		Extra: "auto_increment",
	},
	{
		Schema: "world",
		Table:  "country",
		Name:   "Name",
		Type:   "char(52)",
		Null:   "NO",
		Key:    "",
		Default: sql.NullString{
			String: "",
			Valid:  false,
		},
		Extra: "",
	},
	{
		Schema: "world",
		Table:  "country",
		Name:   "CountryCode",
		Type:   "char(3)",
		Null:   "NO",
		Key:    "",
		Default: sql.NullString{
			String: "",
			Valid:  false,
		},
		Extra: "",
	},
	{
		Schema: "world",
		Table:  "country",
		Name:   "Continent",
		Type:   "enum('Asia','Europe','North America','Africa','Oceania','Antarctica','South America')",
		Null:   "NO",
		Key:    "",
		Default: sql.NullString{
			String: "Asia",
			Valid:  false,
		},
		Extra: "",
	},
	{
		Schema: "world",
		Table:  "country",
		Name:   "Region",
		Type:   "char(26)",
		Null:   "NO",
		Key:    "",
		Default: sql.NullString{
			String: "",
			Valid:  false,
		},
		Extra: "",
	},
	{
		Schema: "world",
		Table:  "country",
		Name:   "SurfaceArea",
		Type:   "decimal(10,2)",
		Null:   "NO",
		Key:    "",
		Default: sql.NullString{
			String: "0.00",
			Valid:  false,
		},
		Extra: "auto_increment",
	},
	{
		Schema: "world",
		Table:  "country",
		Name:   "IndepYear",
		Type:   "smallint(6)",
		Null:   "YES",
		Key:    "",
		Default: sql.NullString{
			String: "0",
			Valid:  false,
		},
		Extra: "",
	},
	{
		Schema: "world",
		Table:  "country",
		Name:   "LifeExpectancy",
		Type:   "decimal(3,1)",
		Null:   "YES",
		Key:    "",
		Default: sql.NullString{
			String: "<null>",
			Valid:  false,
		},
		Extra: "",
	},
	{
		Schema: "world",
		Table:  "country",
		Name:   "GNP",
		Type:   "decimal(10,2)",
		Null:   "YES",
		Key:    "",
		Default: sql.NullString{
			String: "<null>",
			Valid:  false,
		},
		Extra: "",
	},
	{
		Schema: "world",
		Table:  "country",
		Name:   "GNPOld",
		Type:   "decimal(10,2)",
		Null:   "YES",
		Key:    "",
		Default: sql.NullString{
			String: "<null>",
			Valid:  false,
		},
		Extra: "",
	},
	{
		Schema: "world",
		Table:  "country",
		Name:   "LocalName",
		Type:   "char(45)",
		Null:   "NO",
		Key:    "",
		Default: sql.NullString{
			String: "",
			Valid:  false,
		},
		Extra: "",
	},
	{
		Schema: "world",
		Table:  "country",
		Name:   "GovernmentForm",
		Type:   "char(45)",
		Null:   "NO",
		Key:    "",
		Default: sql.NullString{
			String: "",
			Valid:  false,
		},
		Extra: "",
	},
	{
		Schema: "world",
		Table:  "country",
		Name:   "HeadOfState",
		Type:   "char(60)",
		Null:   "YES",
		Key:    "",
		Default: sql.NullString{
			String: "<null>",
			Valid:  false,
		},
		Extra: "",
	},
	{
		Schema: "world",
		Table:  "country",
		Name:   "Capital",
		Type:   "int(11)",
		Null:   "YES",
		Key:    "",
		Default: sql.NullString{
			String: "<null>",
			Valid:  false,
		},
		Extra: "",
	},
	{
		Schema: "world",
		Table:  "country",
		Name:   "Code2",
		Type:   "char(2)",
		Null:   "NO",
		Key:    "",
		Default: sql.NullString{
			String: "",
			Valid:  false,
		},
		Extra: "",
	},
}
var dummyCountryLanguageColumns = []*ColumnDesc{
	{
		Schema: "world",
		Table:  "countrylanguage",
		Name:   "CountryCode",
		Type:   "char(3)",
		Null:   "NO",
		Key:    "PRI",
		Default: sql.NullString{
			String: "",
			Valid:  false,
		},
		Extra: "",
	},
	{
		Schema: "world",
		Table:  "countrylanguage",
		Name:   "Language",
		Type:   "char(30)",
		Null:   "NO",
		Key:    "PRI",
		Default: sql.NullString{
			String: "",
			Valid:  false,
		},
		Extra: "",
	},
	{
		Schema: "world",
		Table:  "countrylanguage",
		Name:   "IsOfficial",
		Type:   "enum('T','F')",
		Null:   "NO",
		Key:    "F",
		Default: sql.NullString{
			String: "",
			Valid:  false,
		},
		Extra: "",
	},
	{
		Schema: "world",
		Table:  "countrylanguage",
		Name:   "Percentage",
		Type:   "decimal(4,1)",
		Null:   "NO",
		Key:    "",
		Default: sql.NullString{
			String: "0.0",
			Valid:  false,
		},
		Extra: "",
	},
}

type MockResult struct {
	MockLastInsertID func() (int64, error)
	MockRowsAffected func() (int64, error)
}

func (m *MockResult) LastInsertId() (int64, error) {
	return m.MockLastInsertID()
}
func (m *MockResult) RowsAffected() (int64, error) {
	return m.MockRowsAffected()
}

func init() {
	Register("mock", func(cfg *Config) Database {
		return &MockDB{
			MockOpen:           func() error { return nil },
			MockClose:          func() error { return nil },
			MockDatabase:       func() (string, error) { return "world", nil },
			MockDatabases:      func() ([]string, error) { return dummyDatabases, nil },
			MockDatabaseTables: func() (map[string][]string, error) { return dummyDatabaseTables, nil },
			MockTables:         func() ([]string, error) { return dummyTables, nil },
			MockDescribeTable: func(tableName string) ([]*ColumnDesc, error) {
				switch tableName {
				case "city":
					return dummyCityColumns, nil
				case "country":
					return dummyCountryColumns, nil
				case "countrylanguage":
					return dummyCountryLanguageColumns, nil
				}
				return nil, nil
			},
			MockDescribeDatabaseTable: func() ([]*ColumnDesc, error) {
				res := []*ColumnDesc{}
				res = append(res, dummyCityColumns...)
				res = append(res, dummyCountryColumns...)
				res = append(res, dummyCountryLanguageColumns...)
				return res, nil

			},
			MockExec: func(ctx context.Context, query string) (sql.Result, error) {
				return &MockResult{
					MockLastInsertID: func() (int64, error) { return 11, nil },
					MockRowsAffected: func() (int64, error) { return 22, nil },
				}, nil
			},
			MockQuery: func(ctx context.Context, query string) (*sql.Rows, error) {
				return &sql.Rows{}, nil
			},
		}
	})
}
