package client

import (
	"fmt"

	"github.com/pen/go-gle/option"
	"github.com/pen/go-gle/spreadsheets"
	"github.com/pen/go-gle/util"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
)

type Client struct {
	gJWTConfig *jwt.Config
}

func New(options ...option.Option) (*Client, error) {
	opt := option.Apply(options, &option.Default{
		EnvName: "GOOGLE_API_KEY",
	})

	json := []byte(opt.JSON)

	if opt.JSON == "" {
		if opt.EnvName == "" {
			return nil, fmt.Errorf("missing opt.envName")
		}

		var err error

		json, err = util.GetEncodedEnv(opt.EnvName)
		if err != nil {
			return nil, fmt.Errorf("on GetEncodedEnv(): %w", err)
		}
	}

	jwtConfig, err := google.JWTConfigFromJSON(json, opt.Scopes...)
	if err != nil {
		return nil, fmt.Errorf("on JWTConfigFromJSON(): %w", err)
	}

	return &Client{
		gJWTConfig: jwtConfig,
	}, nil
}

func (c *Client) NewSpreadsheetsService() (*spreadsheets.Service, error) {
	return spreadsheets.NewService(c.gJWTConfig) //nolint:wrapcheck
}
