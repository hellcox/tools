package main

import (
	"bytes"
	"encoding/json"
)

func jsonFormat(in string) string {
	var output bytes.Buffer
	err := json.Indent(&output, []byte(in), "", "  ")
	if err != nil {
		return err.Error()
	}
	return output.String()
}

func jsonCompress(in string) string {
	var data interface{}
	_ = json.Unmarshal([]byte(in), &data)
	str, err := json.Marshal(data)
	if err != nil {
		return err.Error()
	}
	return string(str)
}
