package play

import (
	"testing"
	"time"

	"github.com/jdxj/yuque/modules"

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

func TestFollowerExceed10(t *testing.T) {
	c, err := NewCounter()
	if err != nil {
		t.Fatalf("new counter err: %s\n", err)
	}
	defer c.Stop()

	res, err := c.followerExceed10()
	if err != nil {
		t.Fatalf("get user count fail: %s", err)
	}
	t.Logf("user count: %d\n", len(res))
}

func TestCounter_InsertBook(t *testing.T) {
	c, err := NewCounter()
	if err != nil {
		t.Fatalf("new counter err: %s\n", err)
	}
	defer c.Stop()

	book := &models.BookSerializer{
		ID:           1,
		Type:         "test_type",
		Slug:         "test_slug",
		Name:         "test_name",
		Namespace:    "",
		UserID:       0,
		User:         nil,
		Description:  "",
		CreatorID:    0,
		Public:       0,
		LikesCount:   123,
		WatchesCount: 456,
		CreatedAt:    "",
		UpdatedAt:    "",
	}
	if err := c.InsertBook(1, book); err != nil {
		t.Fatalf("insert book fail: %s", err)
	}
}
