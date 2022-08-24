package main

import (
	"bytes"
	"encoding/json"
	"strconv"
)

func jsonFormat(val string) string {
	var output bytes.Buffer
	err := json.Indent(&output, []byte(val), "", "  ")
	if err != nil {
		return err.Error()
	}
	return output.String()
}

func jsonCompress(val string) string {
	var data interface{}
	_ = json.Unmarshal([]byte(val), &data)
	str, err := json.Marshal(data)
	if err != nil {
		return err.Error()
	}
	return string(str)
}

func removeEscaping(val string) string {
	s, err := strconv.Unquote(`"` + val + `"`)
	if err != nil {
		return "解析失败：" + err.Error()
	}
	return s
}
