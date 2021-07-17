package httpClient

import (
	"client/domain"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"log"
	"os/exec"
	"runtime"
	"strconv"
)

type httpClient struct {
	client *resty.Client
}

func NewHttpClient() *httpClient {
	return &httpClient{
		resty.New(),
	}
}

func (h *httpClient) ReduceURL(url string) string {
	request := domain.Request{
		url,
	}

	var response domain.Response
	resp, err:= h.client.R().SetResult(&response).SetBody(request).Post(domain.WebHost)
	if err != nil {
		log.Println(err)
		return ""
	}

	if resp.IsError() {
		log.Println(errors.New("Response code is" + strconv.Itoa(resp.StatusCode())))
		return ""
	}

	return response.URL
}

func (h *httpClient) RestoreURL(url string) string {
	var response domain.Response

	resp, err:= h.client.R().SetResult(&response).Get(url)
	if err !=nil {
		log.Println(err)
	}
	if resp.IsSuccess() {
		return response.URL
	} else {
		return resp.RawResponse.Status
	}
}

func (h *httpClient) OpenBrowser(url string) {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Println(err)
	}
}

