package nz

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"time"
)

type Client struct {
	client *http.Client
}

func NewClient(time time.Duration) (*Client, error) {
	if time == 0 {
		return nil, errors.New("time is zero")
	}
	return &Client{
		client: &http.Client{
			Timeout: time,
			Transport: &Logg{
				logger: os.Stdout,
				next:   http.DefaultTransport,
			},
		},
	}, nil
}

func (c Client) GetData() ([]AssetData, error) {
	resp, err := c.client.Get("https://api.disneyapi.dev/character")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	//fmt.Println("Starus: ", resp.StatusCode)

	var r caracterResponse
	if err = json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return nil, err
	}

	return r.Data, nil
}
