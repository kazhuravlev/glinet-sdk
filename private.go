package glinet

import (
	"context"
	"fmt"
	"github.com/imroc/req/v3"
	"net/http"
)

func unmarshalResp(resp *req.Response, obj any) error {
	if resp.GetStatusCode() != http.StatusOK {
		return fmt.Errorf("bad status code: %w", ErrUnexpected)
	}

	var status struct {
		Code Code `json:"code"`
	}
	if err := resp.UnmarshalJson(&status); err != nil {
		return err
	}

	if status.Code == CodeBadToken {
		return ErrUnauthorized
	}

	if obj != nil {
		if err := resp.UnmarshalJson(&obj); err != nil {
			return err
		}
	}

	return nil
}

func fetchToken(ctx context.Context, addr string, password string) (string, error) {
	client := req.C().EnableInsecureSkipVerify()
	resp, err := client.R().
		SetContext(ctx).
		SetFormData(map[string]string{"pwd": password}).
		Post(fmt.Sprintf("https://%s/api/router/login", addr))
	if err != nil {
		return "", err
	}

	var m struct {
		Token string `json:"token"`
	}
	if err := resp.UnmarshalJson(&m); err != nil {
		return "", err
	}

	return m.Token, nil
}
