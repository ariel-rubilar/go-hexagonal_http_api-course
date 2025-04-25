package memdb

import (
	"fmt"
)

var (
	CourseTableName = "courses"
)

type MemDB struct {
	store map[string][]map[string]any
}

func (d *MemDB) List(table string) ([]map[string]any, error) {
	rows := []map[string]any{}
	if tableData, ok := d.store[table]; ok {
		for _, row := range tableData {
			rows = append(rows, row)
		}
	} else {
		return nil, fmt.Errorf("table %s not found", table)
	}
	return rows, nil
}

func (d *MemDB) InsertRow(table string, values map[string]any) error {
	if _, ok := d.store[table]; !ok {
		return fmt.Errorf("table %s not found", table)
	}
	d.store[table] = append(d.store[table], values)
	return nil
}

func NewMemDB() (*MemDB, error) {
	store := make(map[string][]map[string]any)
	store[CourseTableName] = []map[string]any{}
	db := &MemDB{
		store: store,
	}

	return db, nil
}
