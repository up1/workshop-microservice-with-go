package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"service01/tracing"
	"strconv"
	"time"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/eapache/go-resiliency/retrier"
	"github.com/gorilla/mux"
)

type Response struct {
	Result string `json:"result"`
}

func CallService02(w http.ResponseWriter, r *http.Request) {
	var state = mux.Vars(r)["state"]
	// Call service 2
	resp := callService02(state, r.Context())
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(resp.Text)))
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(resp.Text))
}

type service2Response struct {
	Text string `json:"result"`
}

var fallbackService02 = service2Response{
	Text: "May the source be with you, always."}

func callService02(state string, ctx context.Context) service2Response {
	// Start a new opentracing child span
	child := tracing.StartSpanFromContextWithLogEvent(ctx, "callService02", "Client send")
	defer tracing.CloseSpan(child, "Client Receive")

	req, err := http.NewRequest("GET", "http://localhost:9002/hello/"+state, nil)
	body, err := performHTTPRequestCircuitBreaker(tracing.UpdateContext(ctx, child), "service02", req)
	if err == nil {
		resp := service2Response{}
		json.Unmarshal(body, &resp)
		return resp
	} else {
		return fallbackService02
	}
}

func callUsingCircuitBreaker(breakerName string, url string, method string) ([]byte, error) {
	output := make(chan []byte, 1)
	errors := hystrix.Go(breakerName, func() error {

		req, _ := http.NewRequest(method, url, nil)
		err := callWithRetries(req, output)

		return err // For hystrix, forward the err from the retrier. It's nil if OK.
	}, func(err error) error {
		fmt.Printf("In fallback function for breaker %v, error: %v\n", breakerName, err.Error())
		circuit, _, _ := hystrix.GetCircuit(breakerName)
		fmt.Printf("Circuit state is: %v\n", circuit.IsOpen())
		return err
	})

	select {
	case out := <-output:
		fmt.Printf("Call in breaker %v successful\n", breakerName)
		return out, nil

	case err := <-errors:
		fmt.Printf("Got error on channel in breaker %v. Msg: %v\n", breakerName, err.Error())
		return nil, err
	}
}

var client http.Client
var RETRIES = 3

func callWithRetries(req *http.Request, output chan []byte) error {

	r := retrier.New(retrier.ConstantBackoff(RETRIES, 100*time.Millisecond), nil)
	attempt := 0
	err := r.Run(func() error {
		attempt++
		resp, err := client.Do(req)
		if err == nil && resp.StatusCode < 299 {
			responseBody, err := ioutil.ReadAll(resp.Body)
			if err == nil {
				output <- responseBody
				return nil
			}
			return err
		} else if err == nil {
			err = fmt.Errorf("Status was %v\n", resp.StatusCode)
		}
		fmt.Printf("Retrier failed, attempt %v\n", attempt)

		return err
	})
	return err
}

func performHTTPRequestCircuitBreaker(ctx context.Context, breakerName string, req *http.Request) ([]byte, error) {
	output := make(chan []byte, 1)
	errors := hystrix.Go(breakerName, func() error {
		tracing.AddTracingToReqFromContext(ctx, req)
		err := callWithRetries(req, output)
		return err
	}, func(err error) error {
		fmt.Errorf("In fallback function for breaker %v, error: %v\n", breakerName, err.Error())
		return err
	})

	select {
	case out := <-output:
		fmt.Errorf("Call in breaker %v successful\n", breakerName)
		return out, nil

	case err := <-errors:
		fmt.Errorf("Got error on channel in breaker %v. Msg: %v\n", breakerName, err.Error())
		return nil, err
	}
}
