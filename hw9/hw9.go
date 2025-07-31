package hw9

import (
	"fmt"
	"reflect"
	"strings"
)

const (
	PropertyNameTag = "properties"
)

var (
	ErrPropertyTagNotFound = fmt.Errorf("tag `%s` not found", PropertyNameTag)
	ErrEmptyPropertyTag    = fmt.Errorf("tag `%s` is empty", PropertyNameTag)
)

type Person struct {
	Name    string `properties:"name"`
	Address string `properties:"address,omitempty"`
	Age     int    `properties:"age"`
	Married bool   `properties:"married"`
}

type FieldMetaInfo struct {
	name      string
	omitEmpty bool
}

func Serialize[T any](data T) (string, error) {
	sb := new(strings.Builder)
	dataType := reflect.TypeOf(data)
	dataValue := reflect.ValueOf(data)
	fieldsCount := dataType.NumField()

	for i := 0; i < fieldsCount; i++ {
		meta, err := parseFieldMeta(dataType.Field(i))
		if err != nil {
			return "", err
		}

		fieldValue := parseFieldValue(dataValue.Field(i))

		if len(fieldValue) > 0 || !meta.omitEmpty {
			sb.WriteString(meta.name)
			sb.WriteString("=")
			sb.WriteString(fieldValue)
			if i < fieldsCount-1 {
				sb.WriteString("\n")
			}
		}
	}

	return sb.String(), nil
}

func parseFieldMeta(field reflect.StructField) (*FieldMetaInfo, error) {
	props, propExists := field.Tag.Lookup(PropertyNameTag)

	if !propExists {
		return nil, ErrPropertyTagNotFound
	}

	if len(props) == 0 {
		return nil, ErrEmptyPropertyTag
	}

	parts := strings.Split(props, ",")
	return &FieldMetaInfo{
		name:      parts[0],
		omitEmpty: len(parts) > 1,
	}, nil
}

func parseFieldValue(fieldValue reflect.Value) string {
	switch fieldValue.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return fmt.Sprintf("%v", fieldValue.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return fmt.Sprintf("%v", fieldValue.Uint())
	case reflect.Float64, reflect.Float32:
		return fmt.Sprintf("%v", fieldValue.Float())
	case reflect.Bool:
		return fmt.Sprintf("%v", fieldValue.Bool())
	default:
		return fieldValue.String()
	}
}
