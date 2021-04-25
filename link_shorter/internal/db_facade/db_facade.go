package db_facade

import (
	"database/sql"
	"errors"
	"github.com/matamyn/tech_assignment_GO/link_shorter/internal/common"
	"github.com/matamyn/tech_assignment_GO/link_shorter/internal/model"
	"strings"
)

type DbFacade struct {
	Db_ *sql.DB
}

var (
	SELECT_LINK       = "SELECT id, short_link_key,link FROM short_links WHERE link =$1"
	SELECT_SHORT_LINK = "SELECT id, short_link_key, link FROM short_links WHERE short_link_key =$1"
	INSERT_SHORT_LINK = "INSERT INTO short_links (short_link_key,link) VALUES ($1, $2)"
)

func InitDbFacade(conf common.MySqlDb) (*DbFacade, error) {
	full_url := conf.User + ":" +
		conf.Password + "@/" +
		conf.DbName

	db, err := sql.Open("mysql", full_url)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &DbFacade{db}, nil
}
func (f *DbFacade) CreateShortLink(link string) (*model.ShortLink, error) {
	u, err := f.checkLink(link)
	if err != nil {
		return nil, err
	}
	if u != nil {
		return u, nil
	}
	f.Db_.QueryRow(INSERT_SHORT_LINK,
		common.GenerateShortLinkKey(),
		link)
	//todo: поискать возможность возвращать значения после INSERT для MySql
	return f.checkLink(link)
}

func (f *DbFacade) GetLink(short_link string) (*model.ShortLink, error) {
	if !strings.HasPrefix(short_link, model.DefaultLink) {
		return nil, errors.New("Incorrect short_link format")
	}
	u := &model.ShortLink{}
	key := strings.Replace(short_link, model.DefaultLink, "", 1)
	err := f.Db_.QueryRow(SELECT_SHORT_LINK, key).Scan(&u.ID_, &u.ShortLinkKey_, &u.Link_)
	return u, err
}

func (f *DbFacade) checkLink(link string) (*model.ShortLink, error) {
	u := &model.ShortLink{}
	err := f.Db_.QueryRow(SELECT_LINK, link).Scan(&u.ID_, &u.ShortLinkKey_, &u.Link_)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return u, nil
}
