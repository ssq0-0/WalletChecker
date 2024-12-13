package httpClient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type HttpClient struct {
	Client *http.Client
}

func NewHttpClient(proxyURL string) (*HttpClient, error) {
	var transport *http.Transport
	if proxyURL != "" {
		proxy, err := url.Parse(proxyURL)
		if err == nil {
			transport = &http.Transport{
				Proxy: http.ProxyURL(proxy),
			}
		}
	}

	if transport == nil {
		transport = &http.Transport{}
	}

	client := &http.Client{
		Transport: transport,
	}

	return &HttpClient{
		Client: client,
	}, nil
}

func (h *HttpClient) SendJSONRequest(urlRequest, method string, reqBody, respBody interface{}) error {
	var req *http.Request
	var err error

	if reqBody != nil {
		jsonData, err := json.Marshal(reqBody)
		if err != nil {
			return err
		}

		req, err = http.NewRequest(method, urlRequest, bytes.NewBuffer(jsonData))
		if err != nil {
			return err
		}

		req.Header.Set("Content-Type", "application/json")
	} else {
		req, err = http.NewRequest(method, urlRequest, nil)
		if err != nil {
			return err
		}
	}

	resp, err := h.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err := h.checkAndParseResp(resp, respBody); err != nil {
		return err
	}

	return nil
}

func (h *HttpClient) checkAndParseResp(resp *http.Response, respBody interface{}) error {
	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("unexpected status code: %d\nBody: %s", resp.StatusCode, string(bodyBytes))
	}
	if respBody != nil {
		if err := json.NewDecoder(resp.Body).Decode(respBody); err != nil {
			return err
		}
	}
	return nil
}
