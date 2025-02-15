package main

import (
	"context"
	"github.com/digitalocean/godo"
)


func DigitalOceanTeamHealthcheck(ctx context.Context, config Config) (bool, error) {
	client := godo.NewFromToken(config.AuthToken)

	_, resp, err := client.Account.Get(ctx)
	if err != nil {
		return false, err
	}

	if resp.StatusCode != 200 {
		return false, err
	}

	return true, nil
}
