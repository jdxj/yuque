package db

import (
	"database/sql"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/astaxie/beego/logs"
)

//func dSN() string {
//	return fmt.Sprintf("%s:%s@tcp(%s)/%s?loc=Local&parseTime=true",
//		"", "@", "127.0.0.1", "yuque")
//}

func NewDataSource(dsn string) (*DataSource, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	ds := &DataSource{
		stop: make(chan int),
	}
	ds.DB = db
	return ds, ds.ping()
}

type DataSource struct {
	*sql.DB

	stop chan int
	wg   sync.WaitGroup
}

func (ds *DataSource) ping() error {
	if err := ds.Ping(); err != nil {
		return err
	}

	ds.wg.Add(1)
	go func() {
		defer ds.wg.Done()
		ticker := time.NewTicker(time.Minute)
		defer ticker.Stop()

		for {
			select {
			case <-ds.stop:
				logs.Debug("stop db ping")
				return

			case <-ticker.C:
			}

			if err := ds.Ping(); err != nil {
				logs.Error("ping db failed: %s", err)
				return
			}
			logs.Info("ping db success")
		}
	}()

	return nil
}

func (ds *DataSource) Stop() error {
	close(ds.stop)
	ds.wg.Wait()

	err := ds.Close()
	logs.Debug("stop db finish")
	return err
}
