package db

import (
	"fmt"
	"github.com/google/uuid"
	"reflect"
	"time"
)

func PopulateWith(ent any, values map[string]any) {
	var err error

	v := reflect.Indirect(reflect.ValueOf(ent).Elem())
	tags := NewTableSchema(ent).Fields()
	if err != nil {
		panic(err)
	}

	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		if f.Kind() == reflect.Ptr {
			f = f.Elem()
		}

		mf := tags[i]
		val := values[mf]
		if val == nil {
			continue
		}

		if !f.CanSet() {
			fmt.Printf("Unsetable field %s of type %s\n", mf, f.Kind())
			continue
		}

		if f.IsValid() {
			switch f.Kind() {
			case reflect.Array:
				if f.Type() != reflect.TypeOf(uuid.UUID{}) {
					_ = fmt.Errorf("тип змінної має бути uuid.UUID")
					continue
				}
				f.Set(reflect.ValueOf(val))
			case reflect.String:
				f.SetString(val.(string))
			case reflect.Int:
				if tmp, ok := val.(int); ok {
					f.SetInt(int64(tmp))
				}
			case reflect.Bool:
				switch val.(type) {
				case string:
					if val == `true` {
						f.SetBool(true)
					} else {
						f.SetBool(false)
					}
				default:
					f.SetBool(val.(bool))
				}
			case reflect.Struct:
				if val.(string) != "" && val.(string) != "0000-00-00 00:00:00" {
					var t time.Time
					loc := &time.Location{}
					t, err = time.ParseInLocation("2006-01-02 15:04:05", val.(string), loc)
					if err != nil {
						fmt.Printf("Datetime don't parsed correctly for field \"%s\", err is %v\n", mf, err)
					}

					f.Set(reflect.ValueOf(t))
				}
			default:
				fmt.Printf("Unpopulated field %s of type %s\n", mf, f.Kind())
			}
		}
	}
}
