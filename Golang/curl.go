package curl

import (
"bytes"
"fmt"
"net"
"net/http"
"net/url"
"strings"
"time"
)

type Option struct {
Data    map[string]string
Headers map[string]string
Timeout int64
Method  string
Cacert string
Capath string
}

/**
 * non ssl curl
 */
func Curl(url string, option *Option) (data *http.Response, err error) {
httpClient, err := buildClient(option)
if err != nil {
return nil, err
}
req, err := buildRequest(url, option)
if err != nil {
return nil, err
}
response, err := httpClient.Do(req)

if err != nil {
return nil, err
}

return response, nil
}

func buildClient(option *Option) (c *http.Client, err error) {
if option != nil && option.Timeout > 0 {
timeout, _ := time.ParseDuration(fmt.Sprint(option.Timeout) + "s")
deadline := time.Now().Add(timeout)

transport := &http.Transport{
Dial: func(netw, addr string) (net.Conn, error) {
c, err := net.DialTimeout(netw, addr, timeout)
if err != nil {
return nil, err
}
c.SetDeadline(deadline)
return c, nil
},
}
c = &http.Client{Transport : transport}
} else {
c = &http.Client{}
}
return c, nil
}

func buildRequest(urlstr string, option *Option) (req *http.Request, err error) {
if option == nil {
return http.NewRequest("GET", urlstr, nil)
}

data := url.Values{}
if option.Data != nil {
for key, value := range option.Data {
data.Add(key, value)
}
}
dataStr := data.Encode()

/// method detect
var method string
switch strings.ToUpper(option.Method) {
case "HEAD":
method = "HEAD"
case "GET":
method = "GET"
case "POST":
method = "POST"
default:
method = "GET"
}
if method != "POST" {
if strings.Contains(urlstr, "?") {
if len(dataStr) > 0 {
urlstr = urlstr + "&" + dataStr
}
} else {
urlstr = urlstr + "?" + dataStr
}
req, err = http.NewRequest(method, urlstr, nil)
} else {
reader := bytes.NewReader([]byte(dataStr))
req, err = http.NewRequest(method, urlstr, reader)
}

if err != nil {
return nil, err
}

if option.Headers != nil {
for key, value := range option.Headers {
req.Header.Add(key, value)
}
}

return req, nil
}

/*
func ProcessCurlInternal(url string, option *curl.Option) ([]byte, error) {
	retry := 0
	RETRY:
	start := time.Now()
	retry ++
	response, err := Curl(url, option)
	if err != nil {
		if retry < 3 {
			goto RETRY
		}
		return nil, error
	}
	if response.StatusCode == 504 && retry < 3 {
		goto RETRY
	}

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("ProcessCurlInternal response.StatusCode !=200")
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("ProcessCurlInternal  %s", err)
	}

	return body, nil
}
*/
