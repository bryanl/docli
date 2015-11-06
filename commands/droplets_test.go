package commands

import (
	"io/ioutil"
	"testing"

	"github.com/digitalocean/doctl"
	"github.com/digitalocean/godo"
	"github.com/stretchr/testify/assert"
)

var (
	testImage = godo.Image{
		ID:      1,
		Slug:    "slug",
		Regions: []string{"test0"},
	}
	testImageList = []godo.Image{testImage}
)

func TestDropletActionList(t *testing.T) {
	client := &godo.Client{
		Droplets: &doctl.DropletsServiceMock{
			ActionsFn: func(id int, opts *godo.ListOptions) ([]godo.Action, *godo.Response, error) {
				assert.Equal(t, 1, id)

				resp := &godo.Response{
					Links: &godo.Links{
						Pages: &godo.Pages{},
					},
				}
				return testActionList, resp, nil
			},
		},
	}

	withTestClient(client, func(c *TestConfig) {
		ns := "test"
		c.Set(ns, doctl.ArgDropletID, 1)
		err := RunDropletActions(ns, ioutil.Discard)
		assert.NoError(t, err)
	})
}

func TestDropletBackupList(t *testing.T) {
	client := &godo.Client{
		Droplets: &doctl.DropletsServiceMock{
			BackupsFn: func(id int, opts *godo.ListOptions) ([]godo.Image, *godo.Response, error) {
				assert.Equal(t, 1, id)

				resp := &godo.Response{
					Links: &godo.Links{
						Pages: &godo.Pages{},
					},
				}
				return testImageList, resp, nil
			},
		},
	}

	withTestClient(client, func(c *TestConfig) {
		ns := "test"
		c.Set(ns, doctl.ArgDropletID, 1)
		err := RunDropletBackups(ns, ioutil.Discard)
		assert.NoError(t, err)
	})
}

func TestDropletCreate(t *testing.T) {
	client := &godo.Client{
		Droplets: &doctl.DropletsServiceMock{
			CreateFn: func(cr *godo.DropletCreateRequest) (*godo.Droplet, *godo.Response, error) {
				expected := &godo.DropletCreateRequest{
					Name:     "droplet",
					Image:    godo.DropletCreateImage{Slug: "image"},
					Region:   "dev0",
					Size:     "1gb",
					UserData: "#cloud-config",
					SSHKeys:  []godo.DropletCreateSSHKey{},
				}

				assert.Equal(t, cr, expected, "create requests did not match")

				return &testDroplet, nil, nil
			},
		},
	}

	withTestClient(client, func(c *TestConfig) {
		ns := "test"
		c.Set(ns, doctl.ArgDropletName, "droplet")
		c.Set(ns, doctl.ArgRegionSlug, "dev0")
		c.Set(ns, doctl.ArgSizeSlug, "1gb")
		c.Set(ns, doctl.ArgImage, "image")
		c.Set(ns, doctl.ArgUserData, "#cloud-config")

		err := RunDropletCreate(ns, ioutil.Discard)
		assert.NoError(t, err)
	})
}

func TestDropletCreateUserDataFile(t *testing.T) {
	userData, err := ioutil.ReadFile("../testdata/cloud-config.yml")
	if err != nil {
		t.Fatal(err)
	}

	client := &godo.Client{
		Droplets: &doctl.DropletsServiceMock{
			CreateFn: func(cr *godo.DropletCreateRequest) (*godo.Droplet, *godo.Response, error) {
				expected := &godo.DropletCreateRequest{
					Name:     "droplet",
					Image:    godo.DropletCreateImage{Slug: "image"},
					Region:   "dev0",
					Size:     "1gb",
					UserData: string(userData),
					SSHKeys:  []godo.DropletCreateSSHKey{},
				}

				assert.Equal(t, cr, expected, "create requests did not match")

				return &testDroplet, nil, nil
			},
		},
	}

	withTestClient(client, func(c *TestConfig) {
		ns := "test"

		c.Set(ns, doctl.ArgDropletName, "droplet")
		c.Set(ns, doctl.ArgRegionSlug, "dev0")
		c.Set(ns, doctl.ArgSizeSlug, "1gb")
		c.Set(ns, doctl.ArgImage, "image")
		c.Set(ns, doctl.ArgUserDataFile, "../testdata/cloud-config.yml")

		err := RunDropletCreate(ns, ioutil.Discard)
		assert.NoError(t, err)
	})
}

func TestDropletDelete(t *testing.T) {
	client := &godo.Client{
		Droplets: &doctl.DropletsServiceMock{
			DeleteFn: func(id int) (*godo.Response, error) {
				assert.Equal(t, id, testDroplet.ID, "droplet ids did not match")
				return nil, nil
			},
		},
	}

	withTestClient(client, func(c *TestConfig) {
		ns := "test"
		c.Set(ns, doctl.ArgDropletID, testDroplet.ID)

		err := RunDropletDelete(ns, ioutil.Discard)
		assert.NoError(t, err)
	})
}

func TestDropletGet(t *testing.T) {
	client := &godo.Client{
		Droplets: &doctl.DropletsServiceMock{
			GetFn: func(id int) (*godo.Droplet, *godo.Response, error) {
				assert.Equal(t, id, testDroplet.ID, "droplet ids did not match")
				return &testDroplet, nil, nil
			},
		},
	}

	withTestClient(client, func(c *TestConfig) {
		ns := "test"
		c.Set(ns, doctl.ArgDropletID, testDroplet.ID)

		err := RunDropletGet(ns, ioutil.Discard)
		assert.NoError(t, err)
	})
}

func TestDropletKernelList(t *testing.T) {
	client := &godo.Client{
		Droplets: &doctl.DropletsServiceMock{
			KernelsFn: func(id int, opts *godo.ListOptions) ([]godo.Kernel, *godo.Response, error) {
				if got, expected := id, 1; got != expected {
					t.Errorf("KernelsFn() id = %d; expected %d", got, expected)
				}

				resp := &godo.Response{
					Links: &godo.Links{
						Pages: &godo.Pages{},
					},
				}
				return testKernelList, resp, nil
			},
		},
	}

	withTestClient(client, func(c *TestConfig) {
		ns := "test"
		c.Set(ns, doctl.ArgDropletID, testDroplet.ID)

		err := RunDropletKernels(ns, ioutil.Discard)
		assert.NoError(t, err)
	})
}

func TestDropletNeighbors(t *testing.T) {
	didRun := false
	client := &godo.Client{
		Droplets: &doctl.DropletsServiceMock{
			NeighborsFn: func(id int) ([]godo.Droplet, *godo.Response, error) {
				didRun = true
				assert.Equal(t, id, 1)

				resp := &godo.Response{
					Links: &godo.Links{
						Pages: &godo.Pages{},
					},
				}
				return testDropletList, resp, nil
			},
		},
	}

	withTestClient(client, func(c *TestConfig) {
		ns := "test"
		c.Set(ns, doctl.ArgDropletID, testDroplet.ID)

		err := RunDropletNeighbors(ns, ioutil.Discard)
		assert.NoError(t, err)
		assert.True(t, didRun)
	})
}

func TestDropletSnapshotList(t *testing.T) {
	client := &godo.Client{
		Droplets: &doctl.DropletsServiceMock{
			SnapshotsFn: func(id int, opts *godo.ListOptions) ([]godo.Image, *godo.Response, error) {
				assert.Equal(t, id, 1)

				resp := &godo.Response{
					Links: &godo.Links{
						Pages: &godo.Pages{},
					},
				}
				return testImageList, resp, nil
			},
		},
	}

	withTestClient(client, func(c *TestConfig) {
		ns := "test"
		c.Set(ns, doctl.ArgDropletID, testDroplet.ID)

		err := RunDropletSnapshots(ns, ioutil.Discard)
		assert.NoError(t, err)
	})
}

func TestDropletsList(t *testing.T) {
	didRun := false
	client := &godo.Client{
		Droplets: &doctl.DropletsServiceMock{
			ListFn: func(opts *godo.ListOptions) ([]godo.Droplet, *godo.Response, error) {
				didRun = true
				resp := &godo.Response{
					Links: &godo.Links{
						Pages: &godo.Pages{},
					},
				}
				return testDropletList, resp, nil
			},
		},
	}

	withTestClient(client, func(c *TestConfig) {
		ns := "test"
		err := RunDropletList(ns, ioutil.Discard)
		assert.NoError(t, err)
		assert.True(t, didRun)
	})
}
