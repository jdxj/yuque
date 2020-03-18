package play

import (
	"database/sql"
	"strconv"
	"time"

	"github.com/jdxj/yuque/client"

	_ "github.com/go-sql-driver/mysql"

	"github.com/astaxie/beego/logs"
)

const limit = 1

func NewCounter() (*Counter, error) {
	dsn := DSN()
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	cli, err := client.NewClientToken(Token())
	if err != nil {
		return nil, err
	}

	stop := make(chan int)

	go func() {
		tickerPing := time.NewTicker(time.Second)
		defer tickerPing.Stop()
		tickerLog := time.NewTicker(time.Minute)
		defer tickerLog.Stop()

		for {
			select {
			case <-stop:
				logs.Debug("stop db")
				return

			case <-tickerPing.C:
				if err := db.Ping(); err != nil {
					logs.Error("db ping error: %s", err)
					return
				}

			case <-tickerLog.C:
				logs.Debug("db ping in normal")
			}
		}
	}()

	c := &Counter{
		db:   db,
		cli:  cli,
		stop: stop,
	}
	return c, nil
}

type Counter struct {
	start uint64
	db    *sql.DB
	stop  chan int

	cli *client.Client
}

func (c *Counter) Users() {
	ticker := time.NewTicker(750 * time.Millisecond)
	defer ticker.Stop()

	for i := uint64(0); i < limit; i++ {
		<-ticker.C

		id := strconv.FormatUint(i, 10)
		user, err := c.cli.Users(id)
		if err != nil {
			logs.Error("id: %d, get user error: %s", i, err)
			continue
		}

		// todo: write to db
		if err := c.InsertUser(user.ID, user.FollowersCount, user.Type, user.Login, user.Name); err != nil {
			logs.Error("write to db err: %s", err)
		}
	}
}

func (c *Counter) InsertUser(id, followersCount int, typ, login, name string) error {
	db := c.db

	_, err := db.Exec("INSERT INTO user (id,type,login,name,followers_count) VALUES (?,?,?,?,?) ",
		id, typ, login, name, followersCount)
	return err
}

func (c *Counter) Stop() {
	close(c.stop)
	c.db.Close()
	logs.Debug("stop counter")
}
