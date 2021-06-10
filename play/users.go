package play

import (
	"math"
	"strconv"
	"sync"
	"time"

	"github.com/jdxj/yuque/client"
	"github.com/jdxj/yuque/db"
	"github.com/jdxj/yuque/modules"

	"github.com/astaxie/beego/logs"
	_ "github.com/go-sql-driver/mysql"
)

//func DSN() string {
//	return fmt.Sprintf("%s:%s@tcp(%s)/%s?loc=Local&parseTime=true",
//		"root", "", "127.0.0.1", "yuque")
//}
//
//func Token() string {
//	return ""
//}

const limit = math.MaxUint64

func NewCounter() (*Counter, error) {
	cli, err := client.NewClientToken(Token())
	if err != nil {
		return nil, err
	}
	stop := make(chan int)

	ds, err := db.NewDataSource(DSN())
	if err != nil {
		return nil, err
	}

	c := &Counter{
		ds:   ds,
		cli:  cli,
		stop: stop,
	}
	return c, nil
}

type Counter struct {
	wg   sync.WaitGroup
	ds   *db.DataSource
	cli  *client.Client
	stop chan int
}

func (c *Counter) Users(start uint64) {
	c.wg.Add(1)
	go func() {
		defer c.wg.Done()
		ticker := time.NewTicker(750 * time.Millisecond)
		defer ticker.Stop()

		for i := start; i < limit; i++ {
			select {
			case <-c.stop:
				logs.Debug("stop users, id: %d", i)
				return

			case <-ticker.C:
			}

			id := strconv.FormatUint(i, 10)
			user, err := c.cli.Users(id)
			if err != nil {
				if err == client.ErrNoFound {
					continue
				}
				logs.Error("id: %d, get user error: %s", i, err)
			} else {
				if err := c.InsertUser(user); err != nil {
					logs.Error("write to db err: %s", err)
				}
			}
		}
	}()
}

func (c *Counter) InsertUser(user *models.UserSerializer) error {
	ds := c.ds

	_, err := ds.Exec("INSERT INTO user (id,type,login,name,followers_count) VALUES (?,?,?,?,?)",
		user.ID, user.Type, user.Login, user.Name, user.FollowersCount)
	return err
}

func (c *Counter) Stop() error {
	close(c.stop)
	err := c.ds.Stop()
	c.wg.Wait()

	logs.Debug("stop counter")
	logs.GetBeeLogger().Flush()
	return err
}

func (c *Counter) Remaining() {
	c.wg.Add(1)
	go func() {
		defer c.wg.Done()

		ticker := time.NewTicker(2 * time.Minute)
		defer ticker.Stop()

		for {
			select {
			case <-c.stop:
				logs.Debug("stop remaining")
				return

			case <-ticker.C:
				logs.Debug("X-RateLimit-Remaining: %s", c.cli.XRateLimitRemaining())
			}
		}
	}()
}
