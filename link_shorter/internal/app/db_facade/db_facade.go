package link_shorter

import (
	"database/sql"
	"errors"
	"github.com/matamyn/tech_assignment_GO/link_shorter/internal/app/model"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type DbFacade struct {
	Db_ *sql.DB
}

func (f *DbFacade) CreateShortLink(link string) (*model.ShortLink, error) {
	u, err := f.checkLink(link)
	if err != nil {
		return nil, err
	}
	if u != nil {
		return u, nil
	}
	f.Db_.QueryRow(
		"INSERT INTO short_links (short_link_key,link) VALUES ($1, $2)",
		generateShortLinkKey(),
		link)
	return f.checkLink(link)
}

func (f *DbFacade) GetLink(short_link string) (*model.ShortLink, error) {
	if !strings.HasPrefix(short_link, model.DefaultLink) {
		return nil, errors.New("Incorrect short_link format")
	}

	u := &model.ShortLink{}
	key := strings.Replace(short_link, model.DefaultLink, "", 1)
	err := f.Db_.QueryRow("SELECT id, short_link_key, link FROM short_links WHERE short_link_key =$1",
		key,
	).Scan(&u.ID_, &u.ShortLinkKey_, &u.Link_)
	return u, err
}

func (f *DbFacade) checkLink(link string) (*model.ShortLink, error) {
	u := &model.ShortLink{}
	err := f.Db_.QueryRow("SELECT id, short_link_key,link FROM short_links WHERE link =$1", link).Scan(&u.ID_, &u.ShortLinkKey_, &u.Link_)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return u, nil
}

func generateShortLinkKey() string {
	rand.Seed(time.Now().UnixNano())
	short_link := strconv.FormatInt(rand.Int63n(77), 32)
	return short_link
}
