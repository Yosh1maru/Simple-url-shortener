package internal

type Url struct {
	ShortURL string
}

type Source interface {
	Save(token, url string) error
	GetToken(url string) (string, error)
	GetUrl(token string) (string, error)
}

func (u *Url) IsValidate() bool {
	return true
}

func (u *Url) Url() string {
	return u.ShortURL
}
