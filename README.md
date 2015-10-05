# goconfig [![Build Status](https://travis-ci.org/saromanov/goconfig.svg?branch=master)](https://travis-ci.org/saromanov/goconfig)

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
