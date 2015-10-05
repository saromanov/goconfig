package goconfig

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/hcl"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"strings"
)

const (
	jsontype    = "json"
	yamltype    = "yaml"
	hcltype     = "hcl"
	unknowntype = "unknow"
)

//Load provides loading of configuration file
func Load(path string, item interface{}) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return fmt.Errorf("Failed to load: %v", err)
	}

	switch checkPath(path) {
	case jsontype:
		err := json.Unmarshal([]byte(data), item)
		if err != nil {
			return err
		}

		return nil

	case yamltype:
		err := yaml.Unmarshal([]byte(data), item)
		if err != nil {
			return err
		}

		return nil

	case hcltype:
		err := hcl.Decode(item, string(data))
		if err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("Failed to load: This type is unknown")
	}
}

func checkPath(path string) string {
	if strings.HasSuffix(path, jsontype) {
		return jsontype
	}

	if strings.HasSuffix(path, yamltype) {
		return yamltype
	}

	return unknowntype
}
