package play

import (
	"testing"
	"time"

	"github.com/astaxie/beego/logs"
)

func TestNewCounter(t *testing.T) {
	c, err := NewCounter()
	if err != nil {
		t.Fatalf("new counter err: %s\n", err)
	}

	logs.Debug("test new counter")
	_ = c
	time.Sleep(3 * time.Minute)
}
