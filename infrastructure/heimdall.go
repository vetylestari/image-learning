package infrastructure

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/gojek/heimdall/v7"
	"github.com/gojek/heimdall/v7/hystrix"
	"github.com/sirupsen/logrus"
)

type Configuration struct {
	Timeout              time.Duration
	CommandName          string
	HystrixTimeout       time.Duration
	MaxConcurrentRequest int
	ErrorPercentTreshold int
	RetryCount           int
}
type Client struct {
	hystrixClient *hystrix.Client
}

var logger *logrus.Logger

func InitLogger(logr *logrus.Logger) {
	logger = logr
}

func InitHeimdall() heimdall.Doer {
	logger := InitLog()
	InitLogger(logger)
	timeout, _ := strconv.ParseInt(os.Getenv("REST_TIMEOUT"), 10, 0)
	hystrixTimeout, _ := strconv.ParseInt(os.Getenv("REST_HYSTRIX_TIMEOUT"), 10, 0)
	maxConcurrentRequest, _ := strconv.ParseInt(os.Getenv("REST_MAX_CONCURRENT_REQUEST"), 10, 0)
	errorPercentage, _ := strconv.ParseInt(os.Getenv("REST_ERROR_PERCENTAGE_TRESHOLD"), 10, 0)
	retryCount, _ := strconv.ParseInt(os.Getenv("REST_RETRY_COUNT"), 10, 0)

	return NewClient(&Configuration{
		Timeout:              time.Duration(timeout * int64(time.Second)),
		CommandName:          fmt.Sprintf("%s call 3rd party API", os.Getenv("APP_NAME")),
		HystrixTimeout:       time.Duration(hystrixTimeout * int64(time.Second)),
		MaxConcurrentRequest: int(maxConcurrentRequest),
		ErrorPercentTreshold: int(errorPercentage),
		RetryCount:           int(retryCount),
	})
}

func NewClient(config *Configuration) heimdall.Doer {
	return &Client{
		hystrixClient: config.setupHystrix(),
	}
}

func (config *Configuration) setupHystrix() *hystrix.Client {
	// First set a backoff mechanism. Constant backoff increases the backoff at a constant rate
	backoffInterval := 10 * time.Millisecond
	// Define a maximum jitter interval. It must be more than 1*time.Millisecond
	maximumJitterInterval := 5 * time.Millisecond
	backoff := heimdall.NewConstantBackoff(backoffInterval, maximumJitterInterval)

	retrier := heimdall.NewRetrier(backoff)

	httpClient := &http.Client{
		Timeout: config.Timeout,
	}

	return hystrix.NewClient(
		hystrix.WithHTTPTimeout(config.Timeout),
		hystrix.WithCommandName(config.CommandName),
		hystrix.WithHystrixTimeout(config.HystrixTimeout),
		hystrix.WithMaxConcurrentRequests(config.MaxConcurrentRequest),
		hystrix.WithErrorPercentThreshold(config.ErrorPercentTreshold),
		hystrix.WithSleepWindow(10),
		hystrix.WithRequestVolumeThreshold(10),
		hystrix.WithHTTPClient(httpClient),
		hystrix.WithRetryCount(config.RetryCount),
		hystrix.WithRetrier(retrier),
	)
}

func (client *Client) Do(req *http.Request) (*http.Response, error) {
	resp, err := client.hystrixClient.Do(req)
	if err != nil {
		return nil, err
	}
	// defer resp.Body.Close()
	return resp, nil
}

type Request struct {
	Client      heimdall.Doer
	Headers     *http.Header
	Method      string
	Body        interface{}
	URL         string
	Result      *map[string]any
	QueryParams map[string]string
}

func PerformRequest(ctx context.Context, data Request) (int, error) {
	var (
		body []byte
		err  error
	)

	queryParams := url.Values{}
	var queryParamsEncoded string
	url, _ := url.Parse(data.URL)

	if len(data.QueryParams) > 0 && data.QueryParams != nil {
		for key, value := range data.QueryParams {
			queryParams.Add(key, value)
		}
		queryParamsEncoded = queryParams.Encode()
		url.RawQuery = queryParamsEncoded
	}

	if data.Body != nil {
		body, err = json.Marshal(data.Body)
		if err != nil {
			return 0, err
		}
	}

	req, err := http.NewRequest(data.Method, url.String(), bytes.NewBuffer(body))
	if err != nil {
		return 0, err
	}

	req = req.WithContext(ctx)
	req.Header = http.Header{}
	if data.Headers != nil {
		req.Header = *data.Headers
	}
	if req.Header.Get("content-type") == "" {
		req.Header.Add("content-type", "application/json")
	}

	tmBeforeRequest := time.Now()

	resp, err := data.Client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("unable connect to the service, got: %s", err.Error())
	}

	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, err
	}
	defer resp.Body.Close()
	// responseBody := resp.Body
	// defer responseBody.Close()
	// responseParsed, _ := lib.ParseResponseToJSON(responseBody)

	logger.WithField("service", "3rd party service").
		WithField("url", req.URL).
		WithField("headers", req.Header).
		WithField("method", req.Method).
		WithField("rt", time.Since(tmBeforeRequest).Milliseconds()).
		WithField("httpStatus", resp.StatusCode).
		WithField("request", json.RawMessage(body)).
		WithField("queryParams", queryParamsEncoded).
		WithField("response", json.RawMessage(result)).Info("http call info")

	// data.Result = json.RawMessage(result)

	// if data.Result == nil {
	// 	return resp.StatusCode, nil
	// }

	err = json.Unmarshal(result, data.Result)
	if err != nil {
		logger.WithField("err", err).Error("failed to parse json response")
		return resp.StatusCode, err
	}

	return resp.StatusCode, nil
}
