package commands

import (
	"io/ioutil"
	"testing"

	"github.com/digitalocean/doctl"
	"github.com/digitalocean/godo"
)

func TestDropletActionsChangeKernel(t *testing.T) {
	client := &godo.Client{
		DropletActions: &doctl.DropletActionsServiceMock{
			ChangeKernelFn: func(id, kernelID int) (*godo.Action, *godo.Response, error) {
				if got, expected := id, 1; got != expected {
					t.Errorf("ChangeKernelFn() id = %d; expected %d", got, expected)
				}
				if got, expected := kernelID, 2; got != expected {
					t.Errorf("ChangeKernelFn() kernelID = %d; expected %d", got, expected)
				}
				return &testAction, nil, nil
			},
		},
	}

	withTestClient(client, func(c *TestConfig) {
		ns := "test"
		c.Set(ns, doctl.ArgDropletID, 1)
		c.Set(ns, doctl.ArgKernelID, 2)

		RunDropletActionChangeKernel(ns, ioutil.Discard)
	})
}
func TestDropletActionsDisableBackups(t *testing.T) {
	client := &godo.Client{
		DropletActions: &doctl.DropletActionsServiceMock{
			DisableBackupsFn: func(id int) (*godo.Action, *godo.Response, error) {
				if got, expected := id, 1; got != expected {
					t.Errorf("DisableBackupsFn() id = %d; expected %d", got, expected)
				}
				return &testAction, nil, nil
			},
		},
	}

	withTestClient(client, func(c *TestConfig) {
		ns := "test"
		c.Set(ns, doctl.ArgDropletID, 1)

		RunDropletActionDisableBackups(ns, ioutil.Discard)
	})

}
func TestDropletActionsEnableIPv6(t *testing.T) {
	client := &godo.Client{
		DropletActions: &doctl.DropletActionsServiceMock{
			EnableIPv6Fn: func(id int) (*godo.Action, *godo.Response, error) {
				if got, expected := id, 1; got != expected {
					t.Errorf("EnableIPv6Fn() id = %d; expected %d", got, expected)
				}
				return &testAction, nil, nil
			},
		},
	}

	withTestClient(client, func(c *TestConfig) {
		ns := "test"
		c.Set(ns, doctl.ArgDropletID, 1)

		RunDropletActionEnableIPv6(ns, ioutil.Discard)
	})
}

func TestDropletActionsEnablePrivateNetworking(t *testing.T) {
	client := &godo.Client{
		DropletActions: &doctl.DropletActionsServiceMock{
			EnablePrivateNetworkingFn: func(id int) (*godo.Action, *godo.Response, error) {
				if got, expected := id, 1; got != expected {
					t.Errorf("EnablePrivateNetworkingFn() id = %d; expected %d", got, expected)
				}
				return &testAction, nil, nil
			},
		},
	}

	withTestClient(client, func(c *TestConfig) {
		ns := "test"
		c.Set(ns, doctl.ArgDropletID, 1)

		RunDropletActionEnablePrivateNetworking(ns, ioutil.Discard)
	})
}
func TestDropletActionsGet(t *testing.T) {
	client := &godo.Client{
		DropletActions: &doctl.DropletActionsServiceMock{
			GetFn: func(dropletID, actionID int) (*godo.Action, *godo.Response, error) {
				if got, expected := dropletID, 1; got != expected {
					t.Errorf("GetFn() droplet id = %d; expected %d", got, expected)
				}
				if got, expected := actionID, 2; got != expected {
					t.Errorf("GetFn() action id = %d; expected %d", got, expected)
				}
				return &testAction, nil, nil
			},
		},
	}

	withTestClient(client, func(c *TestConfig) {
		ns := "test"
		c.Set(ns, doctl.ArgDropletID, 1)
		c.Set(ns, doctl.ArgActionID, 2)

		RunDropletActionGet(ns, ioutil.Discard)
	})
}

func TestDropletActionsPasswordReset(t *testing.T) {
	client := &godo.Client{
		DropletActions: &doctl.DropletActionsServiceMock{
			PasswordResetFn: func(id int) (*godo.Action, *godo.Response, error) {
				if got, expected := id, 1; got != expected {
					t.Errorf("PasswordResetFn() id = %d; expected %d", got, expected)
				}
				return &testAction, nil, nil
			},
		},
	}

	withTestClient(client, func(c *TestConfig) {
		ns := "test"
		c.Set(ns, doctl.ArgDropletID, 1)

		RunDropletActionPasswordReset(ns, ioutil.Discard)
	})
}

func TestDropletActionsPowerCycle(t *testing.T) {
	client := &godo.Client{
		DropletActions: &doctl.DropletActionsServiceMock{
			PowerCycleFn: func(id int) (*godo.Action, *godo.Response, error) {
				if got, expected := id, 1; got != expected {
					t.Errorf("PowerCycleFn() id = %d; expected %d", got, expected)
				}
				return &testAction, nil, nil
			},
		},
	}

	withTestClient(client, func(c *TestConfig) {
		ns := "test"
		c.Set(ns, doctl.ArgDropletID, 1)

		RunDropletActionPowerCycle(ns, ioutil.Discard)
	})

}
func TestDropletActionsPowerOff(t *testing.T) {
	client := &godo.Client{
		DropletActions: &doctl.DropletActionsServiceMock{
			PowerOffFn: func(id int) (*godo.Action, *godo.Response, error) {
				if got, expected := id, 1; got != expected {
					t.Errorf("PowerOffFn() id = %d; expected %d", got, expected)
				}
				return &testAction, nil, nil
			},
		},
	}

	withTestClient(client, func(c *TestConfig) {
		ns := "test"
		c.Set(ns, doctl.ArgDropletID, 1)

		RunDropletActionPowerOff(ns, ioutil.Discard)
	})
}
func TestDropletActionsPowerOn(t *testing.T) {
	client := &godo.Client{
		DropletActions: &doctl.DropletActionsServiceMock{
			PowerOnFn: func(id int) (*godo.Action, *godo.Response, error) {
				if got, expected := id, 1; got != expected {
					t.Errorf("PowerOnFn() id = %d; expected %d", got, expected)
				}
				return &testAction, nil, nil
			},
		},
	}

	withTestClient(client, func(c *TestConfig) {
		ns := "test"
		c.Set(ns, doctl.ArgDropletID, 1)

		RunDropletActionPowerOn(ns, ioutil.Discard)
	})

}
func TestDropletActionsReboot(t *testing.T) {
	client := &godo.Client{
		DropletActions: &doctl.DropletActionsServiceMock{
			RebootFn: func(id int) (*godo.Action, *godo.Response, error) {
				if got, expected := id, 1; got != expected {
					t.Errorf("RebootFn() id = %d; expected %d", got, expected)
				}
				return &testAction, nil, nil
			},
		},
	}

	withTestClient(client, func(c *TestConfig) {
		ns := "test"
		c.Set(ns, doctl.ArgDropletID, 1)

		RunDropletActionReboot(ns, ioutil.Discard)
	})
}

func TestDropletActionsRebuildByImageID(t *testing.T) {
	client := &godo.Client{
		DropletActions: &doctl.DropletActionsServiceMock{
			RebuildByImageIDFn: func(id, imageID int) (*godo.Action, *godo.Response, error) {
				if got, expected := id, 1; got != expected {
					t.Errorf("RebuildByImageIDFn() id = %d; expected %d", got, expected)
				}
				if got, expected := imageID, 2; got != expected {
					t.Errorf("RebuildByImageIDFn() image id = %d; expected %d", got, expected)
				}
				return &testAction, nil, nil
			},
		},
	}

	withTestClient(client, func(c *TestConfig) {
		ns := "test"
		c.Set(ns, doctl.ArgDropletID, 1)
		c.Set(ns, doctl.ArgImage, "2")

		RunDropletActionRebuild(ns, ioutil.Discard)
	})
}

func TestDropletActionsRebuildByImageSlug(t *testing.T) {
	client := &godo.Client{
		DropletActions: &doctl.DropletActionsServiceMock{
			RebuildByImageSlugFn: func(id int, slug string) (*godo.Action, *godo.Response, error) {
				if got, expected := id, 1; got != expected {
					t.Errorf("RebuildByImageSlugFn() id = %d; expected %d", got, expected)
				}
				if got, expected := slug, "slug"; got != expected {
					t.Errorf("RebuildByImageSlugFn() slug = %q; expected %q", got, expected)
				}
				return &testAction, nil, nil
			},
		},
	}

	withTestClient(client, func(c *TestConfig) {
		ns := "test"
		c.Set(ns, doctl.ArgDropletID, 1)
		c.Set(ns, doctl.ArgImage, "slug")

		RunDropletActionRebuild(ns, ioutil.Discard)
	})

}
func TestDropletActionsRename(t *testing.T) {
	client := &godo.Client{
		DropletActions: &doctl.DropletActionsServiceMock{
			RenameFn: func(id int, name string) (*godo.Action, *godo.Response, error) {
				if got, expected := id, 1; got != expected {
					t.Errorf("RenameFn() id = %d; expected %d", got, expected)
				}
				if got, expected := name, "name"; got != expected {
					t.Errorf("RenameFn() name = %q; expected %q", got, expected)
				}
				return &testAction, nil, nil
			},
		},
	}

	withTestClient(client, func(c *TestConfig) {
		ns := "test"
		c.Set(ns, doctl.ArgDropletID, 1)
		c.Set(ns, doctl.ArgDropletName, "name")

		RunDropletActionRename(ns, ioutil.Discard)
	})
}

func TestDropletActionsResize(t *testing.T) {
	client := &godo.Client{
		DropletActions: &doctl.DropletActionsServiceMock{
			ResizeFn: func(id int, slug string, resize bool) (*godo.Action, *godo.Response, error) {
				if got, expected := id, 1; got != expected {
					t.Errorf("ResizeFn() id = %d; expected %d", got, expected)
				}
				if got, expected := slug, "slug"; got != expected {
					t.Errorf("ResizeFn() name = %q; expected %q", got, expected)
				}
				if got, expected := resize, true; got != expected {
					t.Errorf("ResizeFn() resize = %t; expected %t", got, expected)
				}
				return &testAction, nil, nil
			},
		},
	}

	withTestClient(client, func(c *TestConfig) {
		ns := "test"
		c.Set(ns, doctl.ArgDropletID, 1)
		c.Set(ns, doctl.ArgImageSlug, "slug")
		c.Set(ns, doctl.ArgResizeDisk, true)

		RunDropletActionResize(ns, ioutil.Discard)
	})
}

func TestDropletActionsRestore(t *testing.T) {
	client := &godo.Client{
		DropletActions: &doctl.DropletActionsServiceMock{
			RestoreFn: func(id, imageID int) (*godo.Action, *godo.Response, error) {
				if got, expected := id, 1; got != expected {
					t.Errorf("RestoreFn() id = %d; expected %d", got, expected)
				}
				if got, expected := imageID, 2; got != expected {
					t.Errorf("RestoreFn() imageID = %d; expected %d", got, expected)
				}
				return &testAction, nil, nil
			},
		},
	}

	withTestClient(client, func(c *TestConfig) {
		ns := "test"
		c.Set(ns, doctl.ArgDropletID, 1)
		c.Set(ns, doctl.ArgImageID, 2)

		RunDropletActionRestore(ns, ioutil.Discard)
	})
}

func TestDropletActionsShutdown(t *testing.T) {
	client := &godo.Client{
		DropletActions: &doctl.DropletActionsServiceMock{
			ShutdownFn: func(id int) (*godo.Action, *godo.Response, error) {
				if got, expected := id, 1; got != expected {
					t.Errorf("ShutdownFn() id = %d; expected %d", got, expected)
				}
				return &testAction, nil, nil
			},
		},
	}

	withTestClient(client, func(c *TestConfig) {
		ns := "test"
		c.Set(ns, doctl.ArgDropletID, 1)

		RunDropletActionShutdown(ns, ioutil.Discard)
	})
}

func TestDropletActionsSnapshot(t *testing.T) {
	client := &godo.Client{
		DropletActions: &doctl.DropletActionsServiceMock{
			SnapshotFn: func(id int, name string) (*godo.Action, *godo.Response, error) {
				if got, expected := id, 1; got != expected {
					t.Errorf("ShutdownFn() id = %d; expected %d", got, expected)
				}
				if got, expected := name, "name"; got != expected {
					t.Errorf("ShutdownFn() name = %q; expected %q", got, expected)
				}
				return &testAction, nil, nil
			},
		},
	}

	withTestClient(client, func(c *TestConfig) {
		ns := "test"
		c.Set(ns, doctl.ArgDropletID, 1)
		c.Set(ns, doctl.ArgSnapshotName, "name")

		RunDropletActionSnapshot(ns, ioutil.Discard)
	})
}

func TestDropletActionsUpgrade(t *testing.T) {
	client := &godo.Client{
		DropletActions: &doctl.DropletActionsServiceMock{
			UpgradeFn: func(id int) (*godo.Action, *godo.Response, error) {
				if got, expected := id, 1; got != expected {
					t.Errorf("RenameFn() id = %d; expected %d", got, expected)
				}
				return &testAction, nil, nil
			},
		},
	}

	withTestClient(client, func(c *TestConfig) {
		ns := "test"
		c.Set(ns, doctl.ArgDropletID, 1)

		RunDropletActionUpgrade(ns, ioutil.Discard)
	})
}
