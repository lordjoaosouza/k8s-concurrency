package picalculator

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lordjoaosouza/k3s-concurrency/presenter/req"
	"github.com/lordjoaosouza/k3s-concurrency/presenter/res"
	"github.com/rs/zerolog"
	"io"
	"net/http"
	"os"
	"sync"
	"time"
)

type Service struct {
	logger *zerolog.Logger
}

func New(logger *zerolog.Logger) *Service {
	return &Service{
		logger: logger,
	}
}

func (svc Service) CalculatePI(ctx context.Context, quantity int) (*res.CalculatePi, error) {
	hostName, err := os.Hostname()
	if err != nil {
		return nil, err
	}

	pi := 0.0
	denominator := 1.0
	for i := 0; i < quantity; i++ {
		if i%2 == 0 {
			pi += 4 / denominator
		} else {
			pi -= 4 / denominator
		}

		denominator += 2.0
	}

	return &res.CalculatePi{
		HostName: hostName,
		PiValue:  pi,
	}, nil
}

type RequisitionValues struct {
	Success  bool   `json:"success"`
	Message  string `json:"message"`
	DataJson Data   `json:"data"`
}

type Data struct {
	PiValue  float64 `json:"piValue"`
	HostName string  `json:"hostName"`
}

func (svc Service) TestAPI(ctx context.Context, request *req.TestAPIRequest) (*res.TestAPIResponse, error) {
	quantityRequisition := request.Quantity
	ip := request.IP
	c := make(chan bool)

	startTime := time.Now().UnixMilli()
	go startTest(c, quantityRequisition, ip)
	hits := 0
	fails := 0

	for {
		isHit, open := <-c
		
		if !open {
			break
		}

		if isHit {
			hits++
		} else {
			fails++
		}

	}
	afterTime := time.Now().UnixMilli()

	resultTime := afterTime - startTime
	return &res.TestAPIResponse{
		Time:  resultTime,
		Hits:  hits,
		Fails: fails,
	}, nil
}

func startTest(c chan bool, quantityRequisition int, ip string) {
	var wg sync.WaitGroup
	wg.Add(quantityRequisition)

	for i := 0; i < quantityRequisition; i++ {
		go func() {
			c <- makeRequest(ip)
			wg.Done()
		}()
	}

	wg.Wait()
	close(c)
}

func makeRequest(ip string) bool {
	resp, err := http.Get(ip)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	reqValues := RequisitionValues{}

	if err := json.Unmarshal(respBody, &reqValues); err != nil {
		fmt.Println(err.Error())
		return false
	}

	fmt.Println("Host name: ", reqValues.DataJson.HostName, " \nValue: ", reqValues.DataJson.PiValue)
	return true
}
