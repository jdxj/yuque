package modules

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Tmp struct {
	Name string `json:"name"`
}

func TestJson(t *testing.T) {
	tmp := &Tmp{Name: "1234"}
	data, _ := json.Marshal(tmp)
	fmt.Printf("%s\n", data)

	tmp2 := &Tmp{}
	data, _ = json.Marshal(tmp2)
	fmt.Printf("%s\n", data)

	tmp3 := &Tmp{}
	str := `{"name":null}`
	if err := json.Unmarshal([]byte(str), tmp3); err != nil {
		t.Fatalf("%s\n", err)
	} else {
		fmt.Printf("->:%#v\n", tmp3)
	}
}
