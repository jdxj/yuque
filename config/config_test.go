package config

import (
	"fmt"
	"testing"
)

func TestToken(t *testing.T) {
	if err := Init("./config.yaml"); err != nil {
		t.Fatalf("%s\n", err)
	}

	fmt.Printf("token: %s\n", Token())
}
