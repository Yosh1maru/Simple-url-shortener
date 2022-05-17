package postgresql

import (
	"context"
	"database/sql"
)

type UrlSource struct {
	Db *sql.DB
}

func NewUrlSource(db *sql.DB) *UrlSource {
	return &UrlSource{
		Db: db,
	}
}

type UrlRow struct {
	Token     string `json:"token"`
	FullUrl   string `json:"full_url"`
	CreatedAt string `json:"created_at"`
	ExpireAt  string `json:"expire_at"`
}

func (u *UrlSource) GetToken(url string) (string, error) {
	var row UrlRow
	err := u.Db.
		QueryRowContext(context.Background(), "SELECT token FROM url WHERE full_url=$1 AND expire_at >= now()", url).
		Scan(&row.Token)
	if err == sql.ErrNoRows {
		return "", nil
	}

	return row.Token, nil
}

func (u *UrlSource) GetUrl(token string) (string, error) {
	var row UrlRow
	err := u.Db.
		QueryRowContext(context.Background(), "SELECT full_url FROM url WHERE token=$1 AND expire_at >= now()", token).
		Scan(&row.FullUrl)
	if err == sql.ErrNoRows {
		return "", nil
	}

	return row.FullUrl, nil
}

func (u *UrlSource) Save(token, url string) error {
	_, err := u.Db.Query("INSERT INTO url (token, full_url) VALUES ($1, $2)", token, url)
	if err != nil {
		return err
	}

	return nil
}
