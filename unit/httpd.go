package unit

import (
	"bytes"
	"net/http"

	"github.com/miiniper/loges"
	"go.uber.org/zap"
)

func PostWithBodyString(url, body string, headers http.Header) *http.Response {
	var data []byte = []byte(body)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		loges.Loges.Error("new request error ", zap.Error(err))
		return nil
	}
	req.Header = headers
	cli := &http.Client{}
	resp, err := cli.Do(req)
	if err != nil {
		loges.Loges.Error("new response error ", zap.Error(err))
		return nil
	}
	return resp
}

func PostWithBody(url string, headers http.Header, data []byte) *http.Response {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		loges.Loges.Error("new request error ", zap.Error(err))
		return nil
	}
	req.Header = headers
	cli := &http.Client{}
	resp, err := cli.Do(req)
	if err != nil {
		loges.Loges.Error("new response error ", zap.Error(err))
		return nil
	}
	return resp
}

func GetWithHeader(url string, headers http.Header) *http.Response {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		loges.Loges.Error("new request error ", zap.Error(err))
		return nil
	}
	req.Header = headers
	cli := &http.Client{}
	resp, err := cli.Do(req)
	if err != nil {
		loges.Loges.Error("new response error ", zap.Error(err))
		return nil
	}
	return resp
}

func GetUrl(url string) *http.Response {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		loges.Loges.Error("new request error ", zap.Error(err))
		return nil
	}
	cli := &http.Client{}
	resp, err := cli.Do(req)
	if err != nil {
		loges.Loges.Error("new response error ", zap.Error(err))
		return nil
	}
	return resp
}
