package teleBot

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"socks5"
	"strings"
)

type RequestHandler struct {
	Token string
}

var TelegramBaseURL = "https://api.telegram.org"

func buildURL(method string, token string, params map[string][]string) string {
	urlQuery, err := url.Parse(TelegramBaseURL)
	if err != nil {
		panic(err)
	}
	urlQuery.Path += "bot" + token + "/" + method
	urlPars := url.Values{}
	for key, values := range params {
		for _, val := range values {
			urlPars.Add(key, val)
		}
	}
	urlQuery.RawQuery = urlPars.Encode()
	return urlQuery.String()
}

func (handler RequestHandler) Call(typ string, method string, params map[string][]string, target interface{}) error {
	if !IsValidPointer(target) {
		return errors.New("target must be a non-nil pointer")
	}
	query := buildURL(method, handler.Token, params)
	req, err := http.NewRequest(typ, query, nil)
	if err != nil {
		return err
	}
	resp, err := socks5.SendRequest(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return errors.New("Status " + resp.Status + " returned from " + resp.Request.URL.String())
	}
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	reader := strings.NewReader(string(buf))
	return json.NewDecoder(reader).Decode(target)
}

func (handler RequestHandler) GetMe(target *GetMeResponse) error {
	return handler.Call("GET", "getMe", map[string][]string{}, target)
}

func (handler RequestHandler) GetUpdates(params map[string][]string, target *GetUpdatesResponse) error {
	return handler.Call("GET", "getUpdates", params, target)
}

func (handler RequestHandler) SendMessage(params map[string][]string, target *ResponseUpdate) error {
	return handler.Call("GET", "sendMessage", params, target)
}
