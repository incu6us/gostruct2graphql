# gostruct2graphql

Put your ```struct``` to ```structs/structs.go``` and choose it in the main method of ```main.go```, like:
```
...
func main() {

	getRootDescription(structs.Repository{})
	getRootDescription(structs.Test{})

}
```

then, run the app:

```
go run main.go
```