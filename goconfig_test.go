package goconfig

import
(
	"testing"
)

type Example struct{
	Data int
	Name string
}

func TestLoadJSON(t *testing.T) {
	exp := Example{}
	res := Load("./testdata/test1.json", &exp)
	if res != nil {
		t.Error("File not found")
	}

	if exp.Data != 10 {
		t.Errorf("Expected %d, found %d", 10, exp.Data)
	}

	if exp.Name != "doom" {
		t.Errorf("Expected %s, found %s", "doom", exp.Name)
	}
}

func TestLoadYAML(t *testing.T) {
	exp := Example{}
	res := Load("./testdata/test1.yaml", &exp)
	if res != nil {
		t.Error("File not found")
	}

	if exp.Data != 10 {
		t.Errorf("Expected %d, found %d", 10, exp.Data)
	}

	if exp.Name != "Doom" {
		t.Errorf("Expected %s, found %s", "Doom", exp.Name)
	}
}