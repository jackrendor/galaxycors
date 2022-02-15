package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

func checkArbitrary(client *http.Client, method string, target string, headers http.Header, body io.Reader, cookies []http.Cookie) {
	req, reqErr := http.NewRequest(method, target, body)
	if reqErr != nil {
		log.Fatalf("Error forming request: %s", reqErr.Error())
	}
	payload := "galaxycors.net"
	headers.Set("Origin", payload)
	req.Header = headers
	for _, data := range cookies {
		req.AddCookie(&data)
	}

	resp, respErr := client.Do(req)
	if respErr != nil {
		log.Fatalf("Couldn't perform request: %s", reqErr.Error())
	}
	if originHeader := resp.Header.Get("Access-Control-Allow-Origin"); originHeader == payload {
		if resp.Header.Get("Access-Control-Allow-Credentials") == "true" {
			fmt.Printf("[OK] Arbitrary Origin with credentials: %s\n", originHeader)
		} else {
			fmt.Printf("[OK] Arbitrary Origin: %s\n", originHeader)
		}
	}
}

func checkPrefix(client *http.Client, method string, target string, headers http.Header, body io.Reader, cookies []http.Cookie) {
	req, reqErr := http.NewRequest(method, target, body)
	if reqErr != nil {
		log.Fatalf("Error forming request: %s", reqErr.Error())
	}
	payload := req.Host + ".galaxycors.net"
	headers.Set("Origin", payload)
	req.Header = headers
	for _, data := range cookies {
		req.AddCookie(&data)
	}

	resp, respErr := client.Do(req)
	if respErr != nil {
		log.Fatalf("Couldn't perform request: %s", reqErr.Error())
	}
	if originHeader := resp.Header.Get("Access-Control-Allow-Origin"); originHeader == payload {
		if resp.Header.Get("Access-Control-Allow-Credentials") == "true" {
			fmt.Printf("[OK] Prefix allowed with credentials: %s\n", originHeader)
		} else {
			fmt.Printf("[OK] Prefix allowed: %s\n", originHeader)
		}
	}

}
func checkSuffix(client *http.Client, method string, target string, headers http.Header, body io.Reader, cookies []http.Cookie) {
	req, reqErr := http.NewRequest(method, target, body)
	if reqErr != nil {
		log.Fatalf("Error forming request: %s", reqErr.Error())
	}
	tmp := strings.Split(req.Host, ".")
	payload := "galaxycors" + strings.Join(tmp[len(tmp)-2:], ".")
	headers.Set("Origin", payload)
	req.Header = headers
	for _, data := range cookies {
		req.AddCookie(&data)
	}

	resp, respErr := client.Do(req)
	if respErr != nil {
		log.Fatalf("Couldn't perform request: %s", reqErr.Error())
	}
	if originHeader := resp.Header.Get("Access-Control-Allow-Origin"); originHeader == payload {
		if resp.Header.Get("Access-Control-Allow-Credentials") == "true" {
			fmt.Printf("[OK] Suffix allowed with credentials: %s\n", originHeader)
		} else {
			fmt.Printf("[OK] Suffix allowed: %s\n", originHeader)
		}
	}
}

func checkNull(client *http.Client, method string, target string, headers http.Header, body io.Reader, cookies []http.Cookie) {
	req, reqErr := http.NewRequest(method, target, body)
	if reqErr != nil {
		log.Fatalf("Error forming request: %s", reqErr.Error())
	}
	payload := "null"
	headers.Set("Origin", payload)
	req.Header = headers
	for _, data := range cookies {
		req.AddCookie(&data)
	}

	resp, respErr := client.Do(req)
	if respErr != nil {
		log.Fatalf("Couldn't perform request: %s", reqErr.Error())
	}
	if originHeader := resp.Header.Get("Access-Control-Allow-Origin"); originHeader == payload {
		if resp.Header.Get("Access-Control-Allow-Credentials") == "true" {
			fmt.Printf("[OK] Null allowed (by iframe) with credentials.\n")
		}
		fmt.Printf("[OK] Null allowed (by iframe).\n")
	}

}

func main() {
	var target string
	var method string
	var headerstring string
	var timeout int
	var cookiestring string
	var bodystring string
	fmt.Println("-=GalaxyCors=-")
	flag.StringVar(&method, "method", "GET", "Method to use during requests.")
	flag.StringVar(&target, "url", "", "Url to the target.")
	flag.StringVar(&bodystring, "data", "", "Data to be sent in the body.")
	flag.StringVar(&headerstring, "header", "User-Agent: GalaxyCors 0.1", "Headers to use during requests. head1: data; head2: data2.")
	flag.StringVar(&cookiestring, "cookie", "", "Cookies to use during requests. Example: co1=val;id=000.")
	flag.IntVar(&timeout, "timeout", 10, "Timeout for connection.")
	flag.Parse()

	if target == "" {
		panic("Missing url parameter")
	}

	headers := make(http.Header)
	var cookies []http.Cookie
	var body io.Reader
	for _, data := range strings.Split(headerstring, ";") {
		tmp := strings.Split(data, ":")
		key := strings.Trim(tmp[0], " ")
		value := strings.Trim(tmp[1], " ")
		headers[key] = []string{value}
	}
	for _, data := range strings.Split(cookiestring, ";") {
		tmp := strings.Split(data, "=")
		key := strings.Trim(tmp[0], " ")
		value := strings.Trim(tmp[1], " ")
		cookies = append(cookies, http.Cookie{Name: key, Value: value, MaxAge: 300})
	}

	client := &http.Client{Timeout: time.Duration(timeout) * time.Second}
	if bodystring != "" {
		body = strings.NewReader(bodystring)
	} else {
		body = nil
	}
	req, reqErr := http.NewRequest(method, target, body)
	if reqErr != nil {
		log.Fatalf("Error forming request: %s", reqErr.Error())
	}
	req.Header = headers

	for _, data := range cookies {
		req.AddCookie(&data)
	}
	resp, respErr := client.Do(req)
	if respErr != nil {
		log.Fatalf("Coudln't send request: %s", respErr.Error())
	}

	if resp.Header.Get("Access-Control-Allow-Origin") == "*" {
		fmt.Printf("[OK] Target allows any domain (Access-Control-Allow-Origin: *)\n")
	}
	if resp.Header.Get("Access-Control-Allow-Credentials") == "true" {
		fmt.Printf("[OK] Target allows credentials (Access-Control-Allow-Credentials: true)\n")
	}
	checkArbitrary(client, method, target, headers, body, cookies)
	checkPrefix(client, method, target, headers, body, cookies)
	checkSuffix(client, method, target, headers, body, cookies)
	checkNull(client, method, target, headers, body, cookies)
}
