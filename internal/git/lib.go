package git

import (
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"os/user"
)

type Git interface {
	Fetch(url string) error
}

type Github struct {
	client      *resty.Client
	accessToken string
	cacheDir    string
}

var _ Git = &Github{}

func NewGit(accessToken string) Git {
	client := resty.New()
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	cacheDir := fmt.Sprintf("%s/.kubectx-git/", usr.HomeDir)
	return &Github{
		client:      client,
		accessToken: accessToken,
		cacheDir:    cacheDir,
	}
}

func (g Github) Fetch(url string) error {
	res, err := g.client.R().
		SetBasicAuth(g.accessToken, "").
		Get(url)
	if err != nil {
		return err
	}
	if res.StatusCode() != 200 {
		return errors.New(fmt.Sprintf("Failed to fetch url: %s. Status code: %d", url, res.StatusCode()))
	}

	fmt.Println(res, err)
	return nil
}
