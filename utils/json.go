package utils

import (
	"encoding/json"
	"fmt"
)

func MarshalToString(v interface{}) string {
	data, _ := json.MarshalIndent(v, "", "    ")
	return fmt.Sprintf("%s", data)
}
