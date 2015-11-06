package commands

import (
	"errors"
	"fmt"
	"io"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/digitalocean/doctl"
	"github.com/digitalocean/godo"
	"github.com/spf13/cobra"
)

const (
	sshNoAddress = "could not find droplet address"
)

var (
	errSSHInvalidOptions = fmt.Errorf("neither id or name were supplied")
)

// SSH creates the ssh commands heirarchy
func SSH() *cobra.Command {
	usr, err := user.Current()
	if err != nil {
		logrus.Fatal(err.Error())
	}
	path := filepath.Join(usr.HomeDir, ".ssh", "id_rsa")

	cmdSSH := cmdBuilder(RunSSH, "ssh", "ssh to droplet", writer)
	addIntFlag(cmdSSH, doctl.ArgDropletID, 0, "droplet id")
	addStringFlag(cmdSSH, doctl.ArgDropletName, "", "droplet name")
	addStringFlag(cmdSSH, doctl.ArgSSHUser, "root", "ssh user")
	addStringFlag(cmdSSH, doctl.ArgsSSHKeyPath, path, "path to private ssh key")
	addIntFlag(cmdSSH, doctl.ArgsSSHPort, 22, "port sshd is running on")

	return cmdSSH
}

// RunSSH finds a droplet to ssh to given input parameters (name or id).
func RunSSH(ns string, out io.Writer) error {
	client := doctl.DoctlConfig.GetGodoClient()
	id := doctl.DoctlConfig.GetInt(ns, doctl.ArgDropletID)
	name := doctl.DoctlConfig.GetString(ns, doctl.ArgDropletName)
	user := doctl.DoctlConfig.GetString(ns, doctl.ArgSSHUser)
	keyPath := doctl.DoctlConfig.GetString(ns, doctl.ArgsSSHKeyPath)
	port := doctl.DoctlConfig.GetInt(ns, doctl.ArgsSSHPort)

	var droplet *godo.Droplet
	var err error

	switch {
	case id > 0 && len(name) < 1:
		droplet, err = getDropletByID(client, id)
		if err != nil {
			return err
		}
	case len(name) > 0 && id < 1:
		var droplets []godo.Droplet
		droplets, err = listDroplets(client)
		for _, d := range droplets {
			if d.Name == name {
				droplet = &d
				break
			}
		}

		if droplet == nil {
			return errors.New("could not find droplet by name")
		}

	default:
		return errSSHInvalidOptions
	}

	user = defaulSSHUser(droplet)
	publicIP := extractDropletPublicIP(droplet)

	if len(publicIP) < 1 {
		return errors.New(sshNoAddress)
	}

	runner := doctl.DoctlConfig.SSH(user, publicIP, keyPath, port)
	return runner.Run()

	// return doctl.DoctlConfig.SSH(user, publicIP, keyPath, port)
}

func removeEmptyOptions(in []string) []string {
	var out []string
	if len(in) == 1 && in[0] == "[]" {
		return out
	}

	for _, s := range in {
		if len(s) > 0 {
			out = append(out, s)
		}
	}

	return out
}

func defaulSSHUser(droplet *godo.Droplet) string {
	slug := strings.ToLower(droplet.Image.Slug)
	if strings.Contains(slug, "coreos") {
		return "core"
	}

	return "root"
}
