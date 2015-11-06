package commands

import (
	"io/ioutil"
	"testing"

	"github.com/digitalocean/doctl"
	"github.com/digitalocean/godo"
	"github.com/stretchr/testify/assert"
)

func TestFloatingIPActionsGet(t *testing.T) {
	client := &godo.Client{
		FloatingIPActions: &doctl.FloatingIPActionsServiceMock{
			GetFn: func(ip string, actionID int) (*godo.Action, *godo.Response, error) {
				assert.Equal(t, "127.0.0.1", ip)
				assert.Equal(t, 2, actionID)
				return &testAction, nil, nil
			},
		},
	}

	withTestClient(client, func(c *TestConfig) {
		ns := "test"
		c.Set(ns, doctl.ArgIPAddress, "127.0.0.1")
		c.Set(ns, doctl.ArgActionID, 2)

		RunFloatingIPActionsGet(ns, ioutil.Discard)
	})

}

func TestFloatingIPActionsAssign(t *testing.T) {
	client := &godo.Client{
		FloatingIPActions: &doctl.FloatingIPActionsServiceMock{
			AssignFn: func(ip string, dropletID int) (*godo.Action, *godo.Response, error) {

				assert.Equal(t, ip, "127.0.0.1")
				assert.Equal(t, dropletID, 2)

				return &testAction, nil, nil
			},
		},
	}

	withTestClient(client, func(c *TestConfig) {
		ns := "test"
		c.Set(ns, doctl.ArgIPAddress, "127.0.0.1")
		c.Set(ns, doctl.ArgDropletID, 2)

		RunFloatingIPActionsAssign(ns, ioutil.Discard)
	})
}

func TestFloatingIPActionsUnassign(t *testing.T) {
	client := &godo.Client{
		FloatingIPActions: &doctl.FloatingIPActionsServiceMock{
			UnassignFn: func(ip string) (*godo.Action, *godo.Response, error) {

				assert.Equal(t, ip, "127.0.0.1")

				return &testAction, nil, nil
			},
		},
	}

	withTestClient(client, func(c *TestConfig) {
		ns := "test"
		c.Set(ns, doctl.ArgIPAddress, "127.0.0.1")

		RunFloatingIPActionsUnassign(ns, ioutil.Discard)
	})
}
