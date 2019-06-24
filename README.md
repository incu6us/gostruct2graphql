# gostruct2graphql

Library to convert Golang's struct to Graphql type.

## Example of usage
```
import github.com/incu6us/gostruct2graphql

type ExampleTestStruct struct {
	Stringer string
}

func main() {
    result, err := gostruct2graphql.GetType(ExampleTestStruct{})
    if err !=nil{
        log.Println("Failed to convert struct to Graphql type", err)
    }
    
    log.Println("Result", result)
}
```

## Response
```
RepositoryGqlType := graphql.NewList(graphql.NewObject(graphql.ObjectConfig{
		  Name: "Repository",
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

// Generated by gostruct2graphql
TestGqlType := graphql.NewObject(graphql.ObjectConfig{
		  Name: "Test",
		  Fields: graphql.Fields{"Stringer": &graphql.Field{
        Type: graphql.String,
      },},
})
```