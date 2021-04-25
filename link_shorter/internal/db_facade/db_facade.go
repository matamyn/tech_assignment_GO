package db_facade

import (
	"context"
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/matamyn/tech_assignment_GO/link_shorter/internal/common"
	"github.com/matamyn/tech_assignment_GO/link_shorter/internal/model"
	"strings"
)

type DbFacade struct {
	Db_ *sql.DB
}

func newDB(dbUrl string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dbUrl)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
func InitDbFacade(conf *common.MySqlDb) (*DbFacade, error) {
	full_url := conf.User + ":" +
		conf.Password + "@/" +
		conf.DbName
	db, err := newDB(full_url)
	if err != nil {
		return nil, err
	}
	return &DbFacade{Db_: db}, nil
}

var (
	SELECT_LINK       = "SELECT id, short_link_key, link FROM short_links WHERE link = ?"
	SELECT_SHORT_LINK = "SELECT id, short_link_key, link FROM short_links WHERE short_link_key = ?"
	INSERT_SHORT_LINK = "INSERT INTO short_links (short_link_key,link) VALUES (?, ?) ;"
)

func (f *DbFacade) CreateShortLink(link string) (*model.LinkDbRow, error) {
	ctx := context.Background()
	tx, _ := f.Db_.BeginTx(ctx, nil)

	u, err := getLink(link, tx)
	if err != nil {
		return nil, err
	}
	if u != nil {
		return u, nil
	}
	_, _ = tx.Exec(INSERT_SHORT_LINK,
		strings.ToUpper(common.GenerateShortLinkKey()),
		link)
	u, _ = getLink(link, tx)
	_ = tx.Commit()
	return u, nil
}

func (f *DbFacade) GetLink(short_link string) (*model.LinkDbRow, error) {
	if !strings.HasPrefix(short_link, model.DefaultLink) {
		return nil, errors.New("Incorrect short_link format")
	}
	u := &model.LinkDbRow{}
	key := strings.Replace(short_link, model.DefaultLink, "", 1)
	err := f.Db_.QueryRow(SELECT_SHORT_LINK, key).Scan(&u.ID_, &u.ShortLinkKey_, &u.Link_)
	return u, err
}

func getLink(link string, tx *sql.Tx) (*model.LinkDbRow, error) {
	u := &model.LinkDbRow{}
	rows := tx.QueryRow(SELECT_LINK, link)
	err := rows.Scan(&u.ID_, &u.ShortLinkKey_, &u.Link_)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return u, nil
}
