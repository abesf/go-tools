package log

import (
	"bytes"
	"encoding/json"
	"fmt"
)
//fmt shell print
func FmtLog(info interface{}) {
	m, _ := json.Marshal(info)
	var str bytes.Buffer
	_ = json.Indent(&str, m, "", "    ")
	fmt.Println("formated: ", str.String())
}