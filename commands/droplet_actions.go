package commands

import (
	"io"
	"strconv"

	"github.com/digitalocean/doctl"
	"github.com/digitalocean/godo"
	"github.com/spf13/cobra"
)

type actionFn func(client *godo.Client) (*godo.Action, error)

func performAction(out io.Writer, fn actionFn) error {
	client := doctl.DoctlConfig.GetGodoClient()

	a, err := fn(client)
	if err != nil {
		return err
	}

	return doctl.DisplayOutput(a, out)
}

// DropletAction creates the droplet-action command.
func DropletAction() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "droplet-action",
		Aliases: []string{"da"},
		Short:   "droplet action commands",
		Long:    "droplet-action is used to access droplet action commands",
	}

	cmdDropletActionGet := cmdBuilder(RunDropletActionGet, "get", "get droplet action", writer, "g")
	cmd.AddCommand(cmdDropletActionGet)
	addIntFlag(cmdDropletActionGet, doctl.ArgDropletID, 0, "Droplet ID")
	addIntFlag(cmdDropletActionGet, doctl.ArgActionID, 0, "Action ID")

	cmdDropletActionDisableBackups := cmdBuilder(RunDropletActionDisableBackups,
		"disable-backups", "disable backups", writer)
	cmd.AddCommand(cmdDropletActionDisableBackups)
	addIntFlag(cmdDropletActionDisableBackups, doctl.ArgDropletID, 0, "Droplet ID")

	cmdDropletActionReboot := cmdBuilder(RunDropletActionReboot,
		"reboot", "reboot droplet", writer)
	cmd.AddCommand(cmdDropletActionReboot)
	addIntFlag(cmdDropletActionReboot, doctl.ArgDropletID, 0, "Droplet ID")

	cmdDropletActionPowerCycle := cmdBuilder(RunDropletActionPowerCycle,
		"power-cycle", "power cycle droplet", writer)
	cmd.AddCommand(cmdDropletActionPowerCycle)
	addIntFlag(cmdDropletActionPowerCycle, doctl.ArgDropletID, 0, "Droplet ID")

	cmdDropletActionShutdown := cmdBuilder(RunDropletActionShutdown,
		"shutdown", "shutdown droplet", writer)
	cmd.AddCommand(cmdDropletActionShutdown)
	addIntFlag(cmdDropletActionShutdown, doctl.ArgDropletID, 0, "Droplet ID")

	cmdDropletActionPowerOff := cmdBuilder(RunDropletActionPowerOff,
		"power-off", "power off droplet", writer)
	cmd.AddCommand(cmdDropletActionPowerOff)
	addIntFlag(cmdDropletActionPowerOff, doctl.ArgDropletID, 0, "Droplet ID")

	cmdDropletActionPowerOn := cmdBuilder(RunDropletActionPowerOn,
		"power-on", "power on droplet", writer)
	cmd.AddCommand(cmdDropletActionPowerOn)
	addIntFlag(cmdDropletActionPowerOn, doctl.ArgDropletID, 0, "Droplet ID")

	cmdDropletActionPasswordReset := cmdBuilder(RunDropletActionPasswordReset,
		"power-reset", "power reset droplet", writer)
	cmd.AddCommand(cmdDropletActionPasswordReset)
	addIntFlag(cmdDropletActionPasswordReset, doctl.ArgDropletID, 0, "Droplet ID")

	cmdDropletActionEnableIPv6 := cmdBuilder(RunDropletActionEnableIPv6,
		"enable-ipv6", "enable ipv6", writer)
	cmd.AddCommand(cmdDropletActionEnableIPv6)
	addIntFlag(cmdDropletActionEnableIPv6, doctl.ArgDropletID, 0, "Droplet ID")

	cmdDropletActionEnablePrivateNetworking := cmdBuilder(RunDropletActionEnablePrivateNetworking,
		"enable-private-networking", "enable private networking", writer)
	cmd.AddCommand(cmdDropletActionEnablePrivateNetworking)
	addIntFlag(cmdDropletActionEnablePrivateNetworking, doctl.ArgDropletID, 0, "Droplet ID")

	cmdDropletActionUpgrade := cmdBuilder(RunDropletActionUpgrade,
		"upgrade", "upgrade droplet", writer)
	cmd.AddCommand(cmdDropletActionUpgrade)
	addIntFlag(cmdDropletActionUpgrade, doctl.ArgDropletID, 0, "Droplet ID")

	cmdDropletActionRestore := cmdBuilder(RunDropletActionRestore,
		"restore", "restore backup", writer)
	cmd.AddCommand(cmdDropletActionRestore)
	addIntFlag(cmdDropletActionRestore, doctl.ArgDropletID, 0, "Droplet ID")
	addIntFlag(cmdDropletActionRestore, doctl.ArgImageID, 0, "Image ID")

	cmdDropletActionResize := cmdBuilder(RunDropletActionResize,
		"resize", "resize droplet", writer)
	cmd.AddCommand(cmdDropletActionResize)
	addIntFlag(cmdDropletActionResize, doctl.ArgDropletID, 0, "Droplet ID")
	addIntFlag(cmdDropletActionResize, doctl.ArgImageID, 0, "Image ID")
	addBoolFlag(cmdDropletActionResize, doctl.ArgResizeDisk, false, "Resize disk")

	cmdDropletActionRebuild := cmdBuilder(RunDropletActionRebuild,
		"rebuild", "rebuild droplet", writer)
	cmd.AddCommand(cmdDropletActionRebuild)
	addIntFlag(cmdDropletActionRebuild, doctl.ArgDropletID, 0, "Droplet ID")
	addIntFlag(cmdDropletActionRebuild, doctl.ArgImageID, 0, "Image ID")

	cmdDropletActionRename := cmdBuilder(RunDropletActionRename,
		"rename", "rename droplet", writer)
	cmd.AddCommand(cmdDropletActionRename)
	addIntFlag(cmdDropletActionRename, doctl.ArgDropletID, 0, "Droplet ID")
	addStringFlag(cmdDropletActionRename, doctl.ArgDropletName, "", "Droplet name")

	cmdDropletActionChangeKernel := cmdBuilder(RunDropletActionChangeKernel,
		"change-kernel", "change kernel", writer)
	cmd.AddCommand(cmdDropletActionChangeKernel)
	addIntFlag(cmdDropletActionChangeKernel, doctl.ArgDropletID, 0, "Droplet ID")
	addIntFlag(cmdDropletActionChangeKernel, doctl.ArgKernelID, 0, "Kernel ID")

	cmdDropletActionSnapshot := cmdBuilder(RunDropletActionSnapshot,
		"snapshot", "snapshot droplet", writer)
	cmd.AddCommand(cmdDropletActionSnapshot)
	addIntFlag(cmdDropletActionSnapshot, doctl.ArgDropletID, 0, "Droplet ID")
	addIntFlag(cmdDropletActionSnapshot, doctl.ArgSnapshotName, 0, "Snapshot name")

	return cmd
}

// RunDropletActionGet returns a droplet action by id.
func RunDropletActionGet(ns string, out io.Writer) error {
	fn := func(client *godo.Client) (*godo.Action, error) {
		dropletID := doctl.DoctlConfig.GetInt(ns, doctl.ArgDropletID)
		actionID := doctl.DoctlConfig.GetInt(ns, doctl.ArgActionID)

		a, _, err := client.DropletActions.Get(dropletID, actionID)
		return a, err
	}

	return performAction(out, fn)
}

// RunDropletActionDisableBackups disables backups for a droplet.
func RunDropletActionDisableBackups(ns string, out io.Writer) error {
	fn := func(client *godo.Client) (*godo.Action, error) {
		id := doctl.DoctlConfig.GetInt(ns, doctl.ArgDropletID)

		a, _, err := client.DropletActions.DisableBackups(id)
		return a, err
	}

	return performAction(out, fn)
}

// RunDropletActionReboot reboots a droplet.
func RunDropletActionReboot(ns string, out io.Writer) error {
	fn := func(client *godo.Client) (*godo.Action, error) {
		id := doctl.DoctlConfig.GetInt(ns, doctl.ArgDropletID)

		a, _, err := client.DropletActions.Reboot(id)
		return a, err
	}

	return performAction(out, fn)
}

// RunDropletActionPowerCycle power cycles a droplet.
func RunDropletActionPowerCycle(ns string, out io.Writer) error {
	fn := func(client *godo.Client) (*godo.Action, error) {
		id := doctl.DoctlConfig.GetInt(ns, doctl.ArgDropletID)
		a, _, err := client.DropletActions.PowerCycle(id)
		return a, err
	}

	return performAction(out, fn)
}

// RunDropletActionShutdown shuts a droplet down.
func RunDropletActionShutdown(ns string, out io.Writer) error {
	fn := func(client *godo.Client) (*godo.Action, error) {
		id := doctl.DoctlConfig.GetInt(ns, doctl.ArgDropletID)

		a, _, err := client.DropletActions.Shutdown(id)
		return a, err
	}

	return performAction(out, fn)
}

// RunDropletActionPowerOff turns droplet power off.
func RunDropletActionPowerOff(ns string, out io.Writer) error {
	fn := func(client *godo.Client) (*godo.Action, error) {
		id := doctl.DoctlConfig.GetInt(ns, doctl.ArgDropletID)

		a, _, err := client.DropletActions.PowerOff(id)
		return a, err
	}

	return performAction(out, fn)
}

// RunDropletActionPowerOn turns droplet power on.
func RunDropletActionPowerOn(ns string, out io.Writer) error {
	fn := func(client *godo.Client) (*godo.Action, error) {
		id := doctl.DoctlConfig.GetInt(ns, doctl.ArgDropletID)

		a, _, err := client.DropletActions.PowerOn(id)
		return a, err
	}

	return performAction(out, fn)
}

// RunDropletActionPasswordReset resets the droplet root password.
func RunDropletActionPasswordReset(ns string, out io.Writer) error {
	fn := func(client *godo.Client) (*godo.Action, error) {
		id := doctl.DoctlConfig.GetInt(ns, doctl.ArgDropletID)

		a, _, err := client.DropletActions.PasswordReset(id)
		return a, err
	}

	return performAction(out, fn)
}

// RunDropletActionEnableIPv6 enables IPv6 for a droplet.
func RunDropletActionEnableIPv6(ns string, out io.Writer) error {
	fn := func(client *godo.Client) (*godo.Action, error) {
		id := doctl.DoctlConfig.GetInt(ns, doctl.ArgDropletID)

		a, _, err := client.DropletActions.EnableIPv6(id)
		return a, err
	}

	return performAction(out, fn)
}

// RunDropletActionEnablePrivateNetworking enables private networking for a droplet.
func RunDropletActionEnablePrivateNetworking(ns string, out io.Writer) error {
	fn := func(client *godo.Client) (*godo.Action, error) {
		id := doctl.DoctlConfig.GetInt(ns, doctl.ArgDropletID)

		a, _, err := client.DropletActions.EnablePrivateNetworking(id)
		return a, err
	}

	return performAction(out, fn)
}

// RunDropletActionUpgrade upgrades a droplet.
func RunDropletActionUpgrade(ns string, out io.Writer) error {
	fn := func(client *godo.Client) (*godo.Action, error) {
		id := doctl.DoctlConfig.GetInt(ns, doctl.ArgDropletID)

		a, _, err := client.DropletActions.Upgrade(id)
		return a, err
	}

	return performAction(out, fn)
}

// RunDropletActionRestore restores a droplet using an image id.
func RunDropletActionRestore(ns string, out io.Writer) error {
	fn := func(client *godo.Client) (*godo.Action, error) {
		id := doctl.DoctlConfig.GetInt(ns, doctl.ArgDropletID)
		image := doctl.DoctlConfig.GetInt(ns, doctl.ArgImageID)

		a, _, err := client.DropletActions.Restore(id, image)
		return a, err
	}

	return performAction(out, fn)
}

// RunDropletActionResize resizesx a droplet giving a size slug and
// optionally expands the disk.
func RunDropletActionResize(ns string, out io.Writer) error {
	fn := func(client *godo.Client) (*godo.Action, error) {
		id := doctl.DoctlConfig.GetInt(ns, doctl.ArgDropletID)
		size := doctl.DoctlConfig.GetString(ns, doctl.ArgImageSlug)
		disk := doctl.DoctlConfig.GetBool(ns, doctl.ArgResizeDisk)

		a, _, err := client.DropletActions.Resize(id, size, disk)
		return a, err
	}

	return performAction(out, fn)
}

// RunDropletActionRebuild rebuilds a droplet using an image id or slug.
func RunDropletActionRebuild(ns string, out io.Writer) error {
	fn := func(client *godo.Client) (*godo.Action, error) {
		id := doctl.DoctlConfig.GetInt(ns, doctl.ArgDropletID)
		image := doctl.DoctlConfig.GetString(ns, doctl.ArgImage)

		var a *godo.Action
		var err error
		if i, aerr := strconv.Atoi(image); aerr == nil {
			a, _, err = client.DropletActions.RebuildByImageID(id, i)
		} else {
			a, _, err = client.DropletActions.RebuildByImageSlug(id, image)
		}
		return a, err
	}

	return performAction(out, fn)
}

// RunDropletActionRename renames a droplet.
func RunDropletActionRename(ns string, out io.Writer) error {
	fn := func(client *godo.Client) (*godo.Action, error) {
		id := doctl.DoctlConfig.GetInt(ns, doctl.ArgDropletID)
		name := doctl.DoctlConfig.GetString(ns, doctl.ArgDropletName)

		a, _, err := client.DropletActions.Rename(id, name)
		return a, err
	}

	return performAction(out, fn)
}

// RunDropletActionChangeKernel changes the kernel for a droplet.
func RunDropletActionChangeKernel(ns string, out io.Writer) error {
	fn := func(client *godo.Client) (*godo.Action, error) {
		id := doctl.DoctlConfig.GetInt(ns, doctl.ArgDropletID)
		kernel := doctl.DoctlConfig.GetInt(ns, doctl.ArgKernelID)

		a, _, err := client.DropletActions.ChangeKernel(id, kernel)
		return a, err
	}

	return performAction(out, fn)
}

// RunDropletActionSnapshot creates a snapshot for a droplet.
func RunDropletActionSnapshot(ns string, out io.Writer) error {
	fn := func(client *godo.Client) (*godo.Action, error) {
		id := doctl.DoctlConfig.GetInt(ns, doctl.ArgDropletID)
		name := doctl.DoctlConfig.GetString(ns, doctl.ArgSnapshotName)

		a, _, err := client.DropletActions.Snapshot(id, name)
		return a, err
	}

	return performAction(out, fn)
}
