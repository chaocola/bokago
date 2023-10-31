package bokago

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/net/context"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

/*
*网络请求封装
 */

const (
	UserAgent       = "Mozilla/5.0 (iPhone; CPU iPhone OS 16_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 MicroMessenger/8.0.40(0x1800282b) NetType/WIFI Language/zh_CN"
	PostContentType = "application/json;charset=UTF-8"
	DefaultTimeOut  = 10 * time.Second
)

type RequestParams struct {
	Method  string
	Url     string
	Params  map[string]interface{}
	Headers map[string]interface{}
	Body    interface{}
	TimeOut time.Duration
}

func (R *RequestParams) ValidateRequestParams() error {

	R.Method = strings.ToUpper(R.Method)

	if R.Method == "" {
		R.Method = "GET"
	}

	if R.Url == "" {
		return errors.New("url is empty")
	}

	if R.TimeOut == 0 {
		R.TimeOut = DefaultTimeOut
	}

	if R.Method == "POST" {
		if _, ok := R.Headers["Content-Type"]; !ok {
			R.Headers["Content-Type"] = PostContentType
		}
	}
	if _, ok := R.Headers["user-agent"]; !ok {
		R.Headers["user-agent"] = UserAgent
	}

	return nil
}

/*
*网络请求封装
 */

func BaseRequest(params RequestParams) []byte {

	if err := params.ValidateRequestParams(); err != nil {
		log.Println(err)
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), params.TimeOut)

	defer cancel()

	bodyByte, err := json.Marshal(params.Body)

	if err != nil {
		log.Println(err)
		return nil
	}

	req, err := http.NewRequest(params.Method, params.Url, bytes.NewBuffer(bodyByte))

	if err != nil {
		log.Println(err)
		return nil
	}

	req = req.WithContext(ctx)

	// 处理params
	UrlValue := req.URL.Query()
	if params.Params != nil {
		for key, val := range params.Params {
			UrlValue.Add(key, fmt.Sprintf("%v", val))
		}
		req.URL.RawQuery = UrlValue.Encode()
	}

	// 处理headers
	if params.Headers != nil {
		for key, val := range params.Headers {
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

	//一定要关闭res.Body
	defer func() {
		if err = res.Body.Close(); err != nil {
			log.Println(err)
		}
	}()

	//读取body
	resBody, err := io.ReadAll(res.Body) //把  body 内容读入字符串 s
	if err != nil {
		return nil
	}

	return resBody

}

func handlerDefaultTimeOut(other ...time.Duration) time.Duration {
	var defaultTimeOut time.Duration
	if len(other) > 0 {
		defaultTimeOut = other[0]
	} else {
		defaultTimeOut = DefaultTimeOut
	}
	return defaultTimeOut
}

type ClientMain struct{}

// GET http Get method
func (ClientMain) GET(url string, params map[string]interface{}, headers map[string]interface{}, other ...time.Duration) []byte {
	resBody := BaseRequest(RequestParams{
		Method:  "GET",
		Url:     url,
		Params:  params,
		Headers: headers,
		TimeOut: handlerDefaultTimeOut(other...),
	})
	return resBody
}

// POST http post method
func (ClientMain) POST(url string, params map[string]interface{}, headers map[string]interface{}, body interface{}, other ...time.Duration) []byte {

	resBody := BaseRequest(RequestParams{
		Method:  "POST",
		Url:     url,
		Params:  params,
		Headers: headers,
		Body:    body,
		TimeOut: handlerDefaultTimeOut(other...),
	})
	return resBody
}

var Client = ClientMain{}
