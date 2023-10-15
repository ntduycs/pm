package shared

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/goccy/go-json"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

const (
	JiraServer = "https://jira.banvien.com.vn"
)

var (
	httpClient = &http.Client{}
)

type Config struct {
	BaseUrl  string
	Method   string
	Headers  map[string]string
	Data     interface{}
	Callback func(*http.Response)
}

func Get[T any](config *Config, defaultResp T) (T, error) {
	config.Method = http.MethodGet
	return Do[T](config, defaultResp)
}

func Post[T any](config *Config, defaultResp T) (T, error) {
	config.Method = http.MethodPost
	return Do[T](config, defaultResp)
}

func Put[T any](config *Config, defaultResp T) (T, error) {
	config.Method = http.MethodPut
	return Do[T](config, defaultResp)
}

func Delete[T any](config *Config, defaultResp T) (T, error) {
	config.Method = http.MethodDelete
	return Do[T](config, defaultResp)
}

func Do[T any](config *Config, empty T) (T, error) {
	var request *http.Request
	var requestBody io.Reader
	var err error

	if strings.EqualFold(config.Method, http.MethodGet) {
		if config.Data != nil {
			config.BaseUrl = withQueryParams(config.BaseUrl, config.Data.(map[string]interface{}))
		}
	} else {
		requestBody, err = withBody(config.Data)
	}

	if err != nil {
		return empty, err
	}

	request, err = http.NewRequest(config.Method, config.BaseUrl, requestBody)

	if err != nil {
		return empty, err
	}

	header, err := withHeaders(config.Headers)
	if err != nil {
		return empty, err
	}

	request.Header = header

	Logger.Debug("Sending request", zap.String("url", config.BaseUrl), zap.String("method", config.Method))

	response, err := httpClient.Do(request)
	if err != nil {
		return empty, err
	}

	defer func() {
		_ = response.Body.Close()
	}()

	responseBody, err := io.ReadAll(response.Body)

	if err != nil {
		return empty, err
	}

	if response.StatusCode > 299 {
		return empty, fmt.Errorf("http call ended up with status code %d", response.StatusCode)
	}

	resp := new(T)

	err = json.Unmarshal(responseBody, resp)

	if err != nil {
		return empty, err
	}

	if config.Callback != nil {
		go func() {
			defer func() {
				if e := recover(); e != nil {
					Logger.Error("panic occurred but being recovered", zap.Error(errors.New(fmt.Sprintf("%v", e))))
				}
			}()

			config.Callback(response)
		}()
	}

	return *resp, nil
}

func withHeaders(headers map[string]string) (http.Header, error) {
	if headers == nil {
		return nil, nil
	}

	if len(headers) == 0 {
		return nil, errors.New("headers is empty")
	}

	header := http.Header{}
	for key, value := range headers {
		header.Add(key, value)
	}

	return header, nil
}

func withBody(body interface{}) (io.Reader, error) {
	if body == nil {
		return nil, nil
	}

	switch body.(type) {
	case string:
		return strings.NewReader(body.(string)), nil
	case []byte:
		return bytes.NewReader(body.([]byte)), nil
	case io.Reader:
		return body.(io.Reader), nil
	case struct{}:
		marshalled, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		return bytes.NewReader(marshalled), nil
	default:
		return nil, fmt.Errorf("body type %T is not supported", body)
	}
}

func withQueryParams(baseUrl string, queryParams map[string]interface{}) string {
	if queryParams == nil {
		return baseUrl
	}

	if len(queryParams) == 0 {
		return baseUrl
	}

	if !strings.Contains(baseUrl, "?") {
		baseUrl += "?"
	}

	for key, value := range queryParams {
		baseUrl += fmt.Sprintf("%s=%v&", key, value)
	}

	return baseUrl[:len(baseUrl)-1]
}
