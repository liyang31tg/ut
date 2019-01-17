package ut

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

func HttpGet(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, errors.New("访问失败：" + url)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return body, err
}

func HttpGetJSON(url string) (map[string]interface{}, error) {
	body, err := HttpGet(url)
	if err != nil {
		return nil, err
	}
	var f map[string]interface{}
	err = json.Unmarshal(body, &f)
	if err != nil {
		return nil, errors.New("JSON解析失败：" + err.Error())
	}
	return f, nil
}

func HttpPostJSON(url string, p interface{}) (interface{}, error) {
	client := &http.Client{}
	data, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return body, err
}

func HttpPost(url string, xml string) string {
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, strings.NewReader(xml))
	if err != nil {
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
	}
	return string(body)
}
