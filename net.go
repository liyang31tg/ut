package ut

import (
	"io/ioutil"
	"strings"
	"net/http"
	"encoding/json"
	"errors"
)

func httpGet(url string) (error,[]byte) {
	resp, err := http.Get(url)
	if err != nil {
		return errors.New("访问失败："+url),nil
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return nil,body
}

func HttpGetJSON(url string) (error,map[string]interface{}) {
	err ,body := httpGet(url)
	if err != nil {
		return err,nil
	}
	var f map[string]interface{}
	err = json.Unmarshal(body, &f)
	if err != nil {
		return errors.New("JSON解析失败："+err.Error()) ,nil
	}
	return nil,f
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


