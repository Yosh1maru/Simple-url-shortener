package controller

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
	"url-shortener/internal"
	"url-shortener/internal/api/response"
)

type UrlApiController struct {
	Encryptor *internal.Encryptor
	Response  response.Response
	Source    internal.Source
}

func NewUrlApiController(encryptor *internal.Encryptor, response response.Response, source internal.Source) *UrlApiController {
	return &UrlApiController{
		Encryptor: encryptor,
		Response:  response,
		Source:    source,
	}
}

func (u *UrlApiController) CreateShortUrlLink(w http.ResponseWriter, req *http.Request) {
	var requestedUrl internal.Url
	err := json.NewDecoder(req.Body).Decode(&requestedUrl)
	if err != nil {
		u.Response.Error(w, "Link isn't valid", http.StatusNotFound)
		return
	}

	_, err = url.ParseRequestURI(requestedUrl.Url())
	if err != nil {
		u.Response.Error(w, "Link isn't valid", http.StatusNotFound)
		return
	}

	token, _ := u.Source.GetToken(requestedUrl.Url())
	if token == "" {
		encryptedUrl := u.Encryptor.Encrypt(requestedUrl.Url())
		u.Source.Save(encryptedUrl, requestedUrl.Url())
		token = encryptedUrl
	}

	u.Response.Success(w, map[string]string{"shortURL": "http://localhost/" + token})
}

func (u *UrlApiController) RedirectShortUrlLink(w http.ResponseWriter, req *http.Request) {
	shortURLPayload := req.URL.Query()["shortURL"]

	if len(shortURLPayload) == 0 || shortURLPayload[0] == "" {
		u.Response.Error(w, "Query param shortURL doesn't exist", http.StatusNotFound)
		return
	}

	shortURL := shortURLPayload[0]
	token := strings.Split(shortURL, "/")[3]
	fullUrl, _ := u.Source.GetUrl(token)
	if fullUrl == "" {
		u.Response.Error(w, "Url not found", http.StatusNotFound)
	}

	u.Response.Redirect(w, req, fullUrl)
}
