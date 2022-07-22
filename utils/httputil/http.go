package httputil

import (
	"bytes"
	"equity/core/consts"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Post post请求
// url 请求地址
// params string 或 []byte
func Post[T string | []byte](url string, params T) ([]byte, error) {
	req, err := http.NewRequest(consts.HttpPost, url, strings.NewReader(string(params)))
	if err != nil {
		return nil, err
	}
	req.Header.Add(consts.ContentTypeKey, consts.ContentTypeJson)
	client := http.Client{Timeout: 8 * time.Second}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if res == nil {
		return nil, errors.New("nil response")
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("http response status code error,res.StatusCode = %d", res.StatusCode))
	}

	defer func() {
		_ = res.Body.Close()
	}()

	return ioutil.ReadAll(res.Body)
}

func PostForm(uri string, data url.Values) ([]byte, error) {
	u, err := url.ParseRequestURI(uri)
	if err != nil {
		return nil, err
	}
	urlStr := u.String()
	resp, err := http.PostForm(urlStr, data)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)
	if resp == nil {
		return nil, errors.New("nil response")
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("http response status code error,res.StatusCode = %d", resp.StatusCode))
	}
	return ioutil.ReadAll(resp.Body)
}

func PostXML(uri string, data []byte) ([]byte, error) {
	req, err := http.NewRequest(consts.HttpPost, uri, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	req.Header.Add(consts.ContentTypeKey, consts.ContentTypeXML)
	client := http.Client{Timeout: 8 * time.Second}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if res == nil {
		return nil, errors.New("nil response")
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("http response status code error,res.StatusCode = %d", res.StatusCode))
	}

	defer func() {
		_ = res.Body.Close()
	}()

	return ioutil.ReadAll(res.Body)
}
