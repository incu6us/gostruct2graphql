package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDocument(t *testing.T) {
	type args struct {
		strct interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Simple struct conversion",
			args: struct{ strct interface{} }{strct: ExampleTestStruct{}},
			want: `ExampleTestStructGqlType := graphql.NewObject(graphql.ObjectConfig{
		  Name: "ExampleTestStruct",
		  Fields: graphql.Fields{"Stringer": &graphql.Field{
        Type: graphql.String,
      },},
})`,
		},
		{
			name: "Advanced struct conversion",
			args: struct{ strct interface{} }{strct: ExampleRepositoryStruct{}},
			want: `ExampleRepositoryStructGqlType := graphql.NewList(graphql.NewObject(graphql.ObjectConfig{
		  Name: "ExampleRepositoryStruct",
		  Fields: graphql.Fields{"CacheMaxSeconds": &graphql.Field{
        Type: graphql.Int,
      },"CurrentTime": &graphql.Field{
        Type: graphql.Int,
      },"Doc": &graphql.Field{
				Type: graphql.NewObject(graphql.ObjectConfig{
					Name: "Doc",
					Fields: graphql.Fields{"TropData": &graphql.Field{
				Type: graphql.NewObject(graphql.ObjectConfig{
					Name: "TropData",
					Fields: graphql.Fields{"Two016": &graphql.Field{
				Type: graphql.NewList(graphql.NewObject(graphql.ObjectConfig{
					Name: "Two016",
					Fields: graphql.Fields{"Active": &graphql.Field{
        Type: graphql.Boolean,
      },"Category": &graphql.Field{
        Type: graphql.String,
      },"Status": &graphql.Field{
        Type: graphql.String,
      },"TropID": &graphql.Field{
        Type: graphql.String,
      },"TropName": &graphql.Field{
        Type: graphql.String,
      },},
      })),
    },},
      }),
    },"TropHdr": &graphql.Field{
				Type: graphql.NewObject(graphql.ObjectConfig{
					Name: "TropHdr",
					Fields: graphql.Fields{"TNum": &graphql.Field{
        Type: graphql.Int,
      },},
      }),
    },},
      }),
    },"GeneratedTime": &graphql.Field{
        Type: graphql.Int,
      },"ID": &graphql.Field{
        Type: graphql.String,
      },"Status": &graphql.Field{
        Type: graphql.Int,
      },},
}))
`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := GetType(tt.args.strct)
			assert.Nil(t, err)
			assert.Equal(t, strings.TrimSpace(tt.want), strings.TrimSpace(result))
		})
	}
}

type ExampleTestStruct struct {
	Stringer string
}

type ExampleRepositoryStruct []struct {
	CacheMaxSeconds int64
	CurrentTime     int64
	Doc             struct {
		TropData struct {
			Two016 []struct {
				Active   bool
				Category string
				Status   string
				TropID   string
				TropName string
			}
		}
		TropHdr struct {
			TNum int64
		}
	}
	GeneratedTime int64
	ID            string
	Status        int64
}
