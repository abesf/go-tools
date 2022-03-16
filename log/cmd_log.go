package log

import (
	"bytes"
	"encoding/json"
	"fmt"
)
//fmt shell print
func FmtLog(info ...interface{}) {
	for i:=range info{
		m, _ := json.Marshal(info[i])
		var str bytes.Buffer
		_ = json.Indent(&str, m, "", "    ")
		fmt.Println("formated: ", str.String())
	}
}