package db

import (
	"fmt"
	"reflect"
)

type PrimaryKey []string

type Schema interface {
	TableName() string
	Fields() []string
}

type pkParts struct {
	field string
	order int64
}

type TableSchema struct {
	entity any
	fields []Field

	escapeColumns bool

	_fields []string
}

func NewTableSchema(entity any) *TableSchema {
	fields := buildFieldData(entity)

	res := &TableSchema{
		entity:        entity,
		fields:        fields,
		escapeColumns: true,
	}

	return res
}

func (s *TableSchema) Fields() []string {
	var tags []string

	// Отримуємо тип переданої змінної
	t := reflect.TypeOf(s.entity)

	// Перевіряємо, чи val є покажчиком, і отримуємо тип, на який він вказує
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	// Переконуємося, що ми працюємо зі структурою
	if t.Kind() != reflect.Struct {
		fmt.Println("Provided value is not a struct!")
		return tags
	}

	// Проходимо по всіх полях структури
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		// Отримуємо значення JSON тега
		tag := field.Tag.Get("db")
		if tag != "" && tag != "-" {
			tags = append(tags, tag)
		}
	}

	return tags
}

// buildFieldData gather data for all fields for building conditions for passed entity.
func buildFieldData(e any) []Field {
	t := reflect.TypeOf(e)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	numField := t.NumField()
	fields := make([]Field, 0, numField)

	for i := 0; i < numField; i++ {
		var field Field
		currentField := t.Field(i)

		jsonTag, ok := currentField.Tag.Lookup("db")
		if !ok {
			continue
		}

		field.name = jsonTag
		field.kind = currentField.Type.Kind()

		if field.kind == reflect.Ptr {
			field.nullable = true
			field.kind = currentField.Type.Elem().Kind()
		}

		fields = append(fields, field)
	}

	return fields
}

func (s *TableSchema) FieldData() []Field {
	return s.fields
}
