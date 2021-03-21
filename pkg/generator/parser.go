package generator

import (
	"encoding/json"
	"io/ioutil"
)

func getSchema(file string) (*schema, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	parsedSchema := &schema{}
	err = json.Unmarshal(data, parsedSchema)
	if err != nil {
		return nil, err
	}

	return parsedSchema, nil
}
