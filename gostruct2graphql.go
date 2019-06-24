package main

import (
	"bytes"
	"fmt"
	"reflect"
)

func describeStruct(buffer *bytes.Buffer, strct interface{}) error {
	if reflect.ValueOf(strct).Kind() == reflect.Slice {
		switch reflect.ValueOf(strct).Type().Elem().Kind() {
		case reflect.Struct:
			iType := reflect.TypeOf(strct).Elem()
			if err := describeSlice(buffer, iType); err != nil {
				return err
			}
		default:
			if err := describeSimpleType(buffer, strct); err != nil {
				return err
			}
		}

		return nil
	}

	iValue := reflect.ValueOf(strct)
	iType := reflect.TypeOf(strct)

	for i := 0; i < iType.NumField(); i++ {
		v := iValue.Field(i)

		switch v.Kind() {
		case reflect.Struct:
			buffer.WriteString(`"` + iType.Field(i).Name + `": &graphql.Field{
				Type: graphql.NewObject(graphql.ObjectConfig{
					Name: "` + iType.Field(i).Name + `",
					Fields: graphql.Fields{`)

			if err := describeStruct(buffer, v.Interface()); err != nil {
				return err
			}

			buffer.WriteString(`	},
      }),
    },`)
		case reflect.Slice:
			buffer.WriteString(`"` + iType.Field(i).Name + `": &graphql.Field{
				Type: graphql.NewList(graphql.NewObject(graphql.ObjectConfig{
					Name: "` + iType.Field(i).Name + `",
					Fields: graphql.Fields{`)

			if err := describeStruct(buffer, v.Interface()); err != nil {
				return err
			}

			buffer.WriteString(`	},
      })),
    },`)
		default:
			if err := describeSimpleType(buffer, iType.Field(i).Name, iType.Field(i).Type.String(), string(iType.Field(i).Tag)); err != nil {
				return err
			}
		}
	}

	return nil
}

func describeSlice(buffer *bytes.Buffer, iType reflect.Type) error {
	for i := 0; i < iType.NumField(); i++ {
		switch iType.Field(i).Type.Kind() {
		case reflect.Struct:
			buffer.WriteString(`"` + iType.Field(i).Name + `": &graphql.Field{
				Type: graphql.NewObject(graphql.ObjectConfig{
					Name: "` + iType.Field(i).Name + `",
					Fields: graphql.Fields{`)

			if err := describeSimpleType(buffer, iType.Field(i).Name, "struct", string(iType.Field(i).Tag)); err != nil {
				return err
			}
			if err := describeSlice(buffer, iType.Field(i).Type); err != nil {
				return err
			}
			buffer.WriteString(`},
      }),
    },`)
		case reflect.Slice:
			buffer.WriteString(`"` + iType.Field(i).Name + `": &graphql.Field{
				Type: graphql.NewList(graphql.NewObject(graphql.ObjectConfig{
					Name: "` + iType.Field(i).Name + `",
					Fields: graphql.Fields{`)
			if err := describeSimpleType(buffer, iType.Field(i).Name, "slice", string(iType.Field(i).Tag)); err != nil {
				return err
			}
			if err := describeSlice(buffer, iType.Field(i).Type.Elem()); err != nil {
				return err
			}
			buffer.WriteString(`},
      })),
    },`)
		default:
			if err := describeSimpleType(buffer, iType.Field(i).Name, iType.Field(i).Type.String(), string(iType.Field(i).Tag)); err != nil {
				return err
			}
		}
	}

	return nil
}

func describeSimpleType(buffer *bytes.Buffer, text ...interface{}) error {
	var fieldName = text[0].(string)
	var out = ""
	switch text[1] {
	case "string":
		out = `"` + fieldName + `": &graphql.Field{
        Type: graphql.String,
      },`
	case "bool":
		out = `"` + fieldName + `": &graphql.Field{
        Type: graphql.Boolean,
      },`
	case "int64":
		out = `"` + fieldName + `": &graphql.Field{
        Type: graphql.Int,
      },`
	case "int32":
		out = `"` + fieldName + `": &graphql.Field{
          Type: graphql.Int,
        },`
	case "int":
		out = `"` + fieldName + `": &graphql.Field{
            Type: graphql.Int,
          },`
	case "struct", "slice":
	default:
		return fmt.Errorf("unknown type: %s", fieldName)
	}
	buffer.WriteString(out)

	return nil
}

func GetType(strct interface{}) (string, error) {
	var buffer bytes.Buffer

	if reflect.TypeOf(strct).Kind() == reflect.Struct {
		buffer.WriteString(reflect.TypeOf(strct).Name() + `GqlType := graphql.NewObject(graphql.ObjectConfig{
		  Name: "` + reflect.TypeOf(strct).Name() + `",
		  Fields: graphql.Fields{`)

		if err := describeStruct(&buffer, strct); err != nil {
			return "", err
		}

		buffer.WriteString(`},
})`)
	}

	if reflect.TypeOf(strct).Kind() == reflect.Slice {
		buffer.WriteString(reflect.TypeOf(strct).Name() + `GqlType := graphql.NewList(graphql.NewObject(graphql.ObjectConfig{
		  Name: "` + reflect.TypeOf(strct).Name() + `",
		  Fields: graphql.Fields{`)

		if err := describeStruct(&buffer, strct); err != nil {
			return "", err
		}

		buffer.WriteString(`},
}))`)
	}
	buffer.WriteString("\n")
	defer buffer.Reset()

	return buffer.String(), nil
}
