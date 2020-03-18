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

func TestCounter_InsertUser(t *testing.T) {
	c, err := NewCounter()
	if err != nil {
		t.Fatalf("new counter err: %s\n", err)
	}

	if err := c.InsertUser(3445, 4, "User", "jdxj", "jdxj"); err != nil {
		t.Fatalf("%s", err)
	}

	c.db.Close()
}

func TestCounter_Users(t *testing.T) {
	c, err := NewCounter()
	if err != nil {
		t.Fatalf("new counter err: %s\n", err)
	}

	c.Users()
	c.db.Close()
}
