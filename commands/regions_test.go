package commands

import (
	"io/ioutil"
	"testing"

	"github.com/digitalocean/doctl"
	"github.com/digitalocean/godo"
	"github.com/stretchr/testify/assert"
)

var (
	testRegion     = godo.Region{Slug: "dev0"}
	testRegionList = []godo.Region{testRegion}
)

func TestRegionsList(t *testing.T) {
	didList := false

	client := &godo.Client{
		Regions: &doctl.RegionsServiceMock{
			ListFn: func(opt *godo.ListOptions) ([]godo.Region, *godo.Response, error) {
				didList = true

				resp := &godo.Response{
					Links: &godo.Links{
						Pages: &godo.Pages{},
					},
				}
				return testRegionList, resp, nil
			},
		},
	}

	withTestClient(client, func(c *TestConfig) {
		ns := "test"
		RunRegionList(ns, ioutil.Discard)
		assert.True(t, didList)
	})
}
