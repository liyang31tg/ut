package ut

import (
	"io/ioutil"
	"strings"
	"net/http"
	"encoding/json"
	"errors"
)

func httpGet(url string) ([]byte,error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil,errors.New("访问失败："+url)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return body,err
}

func HttpGetJSON(url string) (map[string]interface{},error) {
	body,err  := httpGet(url)
	if err != nil {
		return nil,err
	}
	var f map[string]interface{}
	err = json.Unmarshal(body, &f)
	if err != nil {
		return nil,errors.New("JSON解析失败："+err.Error())
	}
	return f,nil
}

func HttpPost(url string, xml string) string {
	client := &http.Client{}

	req, err := http.NewRequest("POST", url, strings.NewReader(xml))
	if err != nil {
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
	}
	return string(body)

}


