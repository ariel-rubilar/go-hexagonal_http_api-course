package memdb

import "github.com/hyperioxx/memsql"

var (
	CourseTableName = "courses"
)

func NewMemDB() (*memsql.Database, error) {
	db := memsql.NewDatabase()

	db.CreateTable(CourseTableName, []*memsql.Column{
		{
			Name: "id",
			Kind: "uuid",
		},
		{
			Name: "name",
			Kind: "string",
		},
		{
			Name: "duration",
			Kind: "string",
		},
	})
	return db, nil
}
