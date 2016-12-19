package main

import (
	"fmt"
	"reflect"
)

type Repository struct {
	CacheMaxSeconds int64 `json:"cacheMaxSeconds"`
	CurrentTime     int64 `json:"currentTime"`
	Doc             struct {
		TropData struct {
			Two016 []struct {
				TestStruct struct {
					TestField string `json:"testString"`
				} `json:"testStruct"`
				Active   bool   `json:"Active"`
				Category string `json:"Category"`
				Status   string `json:"Status"`
				TropID   string `json:"TropId"`
				TropName string `json:"TropName"`
			} `json:"2016"`
		} `json:"TropData"`
		TropHdr struct {
			TNum int64 `json:"TNum"`
		} `json:"TropHdr"`
	} `json:"doc"`
	GeneratedTime int64  `json:"generatedTime"`
	ID            string `json:"id"`
	Status        int64  `json:"status"`
}

// type TestStruct struct {
// 	Filed1 string
// 	Field2 struct {
// 		SubFiled  string
// 		SubFiled1 string
// 		S         struct {
// 			A string `json:"a"`
// 		} `json:"s"`
// 	}
// }

func DescribeStruct(s interface{}) {
	if reflect.ValueOf(s).Kind() == reflect.Slice {
		switch reflect.ValueOf(s).Type().Elem().Kind() {
		case reflect.Struct:

			iValue := reflect.ValueOf(s).Type().Elem()
			iType := reflect.TypeOf(s).Elem()
			describeSlice(iType, iValue)

		default:
			log("!!! Shit happens !!!")
		}

		return
	}

	iValue := reflect.ValueOf(s)
	iType := reflect.TypeOf(s)

	for i := 0; i < iType.NumField(); i++ {
		v := iValue.Field(i)

		switch v.Kind() {
		case reflect.Struct:
			log(iType.Field(i).Name, "struct", iType.Field(i).Tag)
			DescribeStruct(v.Interface())
		case reflect.Slice:
			log(iType.Field(i).Name, "slice", iType.Field(i).Tag)
			DescribeStruct(v.Interface())
		default:
			log("->", iType.Field(i).Name, iType.Field(i).Type, iType.Field(i).Tag)
		}
	}
}

func describeSlice(iType reflect.Type, iValue reflect.Type) {
	// log("!!!", iType.Field(0).Type)
	for i := 0; i < iType.NumField(); i++ {

		switch iType.Field(i).Type.Kind() {
		case reflect.Struct:
			log("struct -> ", iType.Field(i).Name, "struct", iType.Field(i).Tag)
			describeSlice(iType.Field(i).Type, nil)
		// case reflect.Slice:
		// 	log(iType.Field(i).Name, "slice", iType.Field(i).Tag)
		default:
			log("describeSlice->", iType.Field(i).Name, iType.Field(i).Type, iType.Field(i).Tag)
		}
	}
}

func reflectToString(r reflect.Type) string {
	return reflect.ValueOf(r).String()
}

// func log(text ...interface{}) {
// 	var logList []interface{} = make([]interface{}, len(text))
// 	for i := range text {
// 		// logMap[i] = text[i]
// 		logList[i] = text[i]
// 	}
// 	fmt.Println(logList)
// }

func log(text ...interface{}) {
	fmt.Println(text...)
}

func main() {
	// var testStruct Repository
	DescribeStruct(Repository{})
}
