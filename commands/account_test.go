package commands

import (
	"io/ioutil"
	"testing"

	"github.com/digitalocean/doctl"
	"github.com/digitalocean/godo"
	"github.com/stretchr/testify/assert"
)

var testAccount = &godo.Account{
	DropletLimit:  10,
	Email:         "user@example.com",
	UUID:          "1234",
	EmailVerified: true,
}

func TestAccountGet(t *testing.T) {
	accountDidGet := false

	client := &godo.Client{
		Account: &doctl.AccountServiceMock{
			GetFn: func() (*godo.Account, *godo.Response, error) {
				accountDidGet = true
				return testAccount, nil, nil
			},
		},
	}

	withTestClient(client, func(c *TestConfig) {
		ns := "test"

		err := RunAccountGet(ns, ioutil.Discard)
		assert.NoError(t, err)

		if !accountDidGet {
			t.Errorf("could not retrieve account")
		}
	})
}
