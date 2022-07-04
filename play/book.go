package play

import (
	"strconv"

	"github.com/astaxie/beego/logs"
)

func (c *Counter) followerExceed10() ([]int, error) {
	rows, err := c.ds.Query("SELECT id FROM user WHERE followers_count>10")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []int
	var id int
	for rows.Next() {
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		res = append(res, id)
	}
	return res, nil
}

func (c *Counter) Books() {
	userIDs, err := c.followerExceed10()
	if err != nil {
		logs.Error("get user count fail: %s", err)
		return
	}

	for _, id := range userIDs {
		books, err := c.cli.ListUserRepositories(strconv.Itoa(id))
		if err != nil {
			logs.Error("list books fail, user id: %d, err: %s", id, err)
			continue
		}

		for _, book := range books {
			if err := c.InsertBook(id, book); err != nil {
				logs.Error("insert book fail, book info: %#v", book)
				continue
			}
		}
	}
}

func (c *Counter) InsertBook(userID int, book *models.BookSerializer) error {
	_, err := c.ds.Exec(`INSERT INTO book (user_id,book_id,book_slug,book_name,book_type,likes_count,watches_count)
			VALUES (?,?,?,?,?,?,?)`,
		userID, book.ID, book.Slug, book.Name, book.Type, book.LikesCount, book.WatchesCount)
	return err
}
