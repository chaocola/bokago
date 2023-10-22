package NetWork

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

/*
*网络请求封装
 */

func BaseRequest(url string, params map[string]interface{}, headers map[string]interface{}, body interface{}, method string) []byte {

	bodyByte, _ := json.Marshal(body)

	req, _ := http.NewRequest(method, url, bytes.NewBuffer(bodyByte))

	// 处理params
	UrlValue := req.URL.Query()
	if params != nil {
		for key, val := range params {
			// val 有课能是数字，所以需要转换成字符串

			UrlValue.Add(key, fmt.Sprintf("%v", val))
		}
		req.URL.RawQuery = UrlValue.Encode()
	}

	if req.Header.Get("user-agent") == "" {
		req.Header.Set("user-agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 16_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 MicroMessenger/8.0.40(0x1800282b) NetType/WIFI Language/zh_CN")
	}
	//
	//req.Header.Set("Connection", "close")
	//

	if headers != nil {
		for key, val := range headers {
			req.Header.Set(key, fmt.Sprintf("%v", val))
		}
	}

	//http client
	client := &http.Client{}
	//log.Printf("Go GET URL : %s \n", req.URL.String())

	//发送请求
	res, err := client.Do(req)
	if err != nil {
		return nil
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(res.Body) //一定要关闭res.Body
	//读取body
	resBody, err := io.ReadAll(res.Body) //把  body 内容读入字符串 s
	if err != nil {
		return nil
	}

	return resBody

}

type ClientMain struct{}

// GET http Get method
func (ClientMain) GET(url string, params map[string]interface{}, headers map[string]interface{}) []byte {
	resBody := BaseRequest(url, params, headers, nil, "GET")
	return resBody
}

// POST http post method
func (ClientMain) POST(url string, params map[string]interface{}, headers map[string]interface{}, body interface{}) []byte {

	if _, ok := headers["Content-Type"]; !ok {
		headers["Content-Type"] = "application/json;charset=UTF-8"
	}

	resBody := BaseRequest(url, params, headers, body, "POST")
	return resBody
}

var Client = ClientMain{}
