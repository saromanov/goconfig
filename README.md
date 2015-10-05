# goconfig

Simple loading of configuration from several formats

### Usage

```go
type Example struct{
	Data int
	Name string
}

func main() {
	exp := Example{}
	res := Load("./testdata/test1.json", &exp)
	if res != nil {
	    return
    }
}

```

### Licence
MIT
