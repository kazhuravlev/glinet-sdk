package glinet

import (
	"context"
	"fmt"
	"github.com/imroc/req/v3"
)

type Client struct {
	addr  string
	token string
	http  *req.Client
}

func New(addr, token string) (*Client, error) {
	httpClient := req.C().
		SetBaseURL(fmt.Sprintf("https://%s/cgi-bin/api", addr)).
		SetCommonHeader("Authorization", token).
		EnableInsecureSkipVerify()

	return &Client{
		addr:  addr,
		token: token,
		http:  httpClient,
	}, nil
}

// NewFromPassword will do auth and build a new regular Client.
func NewFromPassword(ctx context.Context, addr, password string) (*Client, error) {
	token, err := fetchToken(ctx, addr, password)
	if err != nil {
		return nil, fmt.Errorf("fetch token: %w", err)
	}

	return New(addr, token)
}

func (c *Client) GetPublicIP(ctx context.Context) (string, error) {
	resp, err := c.http.R().
		SetContext(ctx).
		Get("/internet/public_ip/get")
	if err != nil {
		return "", err
	}

	var res struct {
		ServerIP string `json:"serverip"`
	}
	if err := unmarshalResp(resp, &res); err != nil {
		return "", err
	}

	return res.ServerIP, nil
}

func (c *Client) ModemTurnOnAuto(ctx context.Context) error {
	request := map[string]string{
		"modem_id": "1",
		"bus":      "1-1.2",
	}

	resp, err := c.http.R().
		SetContext(ctx).
		SetFormData(request).
		Post("/modem/auto")
	if err != nil {
		return err
	}

	if err := unmarshalResp(resp, nil); err != nil {
		return err
	}

	return nil
}

func (c *Client) ModemTurnOff(ctx context.Context) error {
	request := map[string]string{
		"disable": "true",
	}

	resp, err := c.http.R().
		SetContext(ctx).
		SetFormData(request).
		Post("/modem/enable")
	if err != nil {
		return err
	}

	if err := unmarshalResp(resp, nil); err != nil {
		return err
	}

	return nil
}

func (c *Client) ModemTurnOn(ctx context.Context) error {
	request := map[string]string{
		"disable": "false",
	}

	resp, err := c.http.R().
		SetContext(ctx).
		SetFormData(request).
		Post("/modem/enable")
	if err != nil {
		return err
	}

	if err := unmarshalResp(resp, nil); err != nil {
		return err
	}

	return nil
}

func (c *Client) GetModemInfo(ctx context.Context) (*GetModemInfoResp, error) {
	resp, err := c.http.R().
		SetContext(ctx).
		Post("/modem/info")
	if err != nil {
		return nil, err
	}

	var res GetModemInfoResp
	if err := unmarshalResp(resp, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) GetClientList(ctx context.Context) (*GetClientListResp, error) {
	resp, err := c.http.R().
		SetContext(ctx).
		Get("/client/list")
	if err != nil {
		return nil, err
	}

	var res GetClientListResp
	if err := unmarshalResp(resp, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) GetNetworkStatus(ctx context.Context) (*GetNetworkStatusResp, error) {
	resp, err := c.http.R().
		SetContext(ctx).
		Get("/internet/reachable")
	if err != nil {
		return nil, err
	}

	var res GetNetworkStatusResp
	if err := unmarshalResp(resp, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) Token() string {
	return c.token
}
