package commands

import (
	"errors"
	"io"

	"github.com/digitalocean/doctl"
	"github.com/digitalocean/godo"
	"github.com/spf13/cobra"
)

// Actions creates the action commands heirarchy.
func Actions() *cobra.Command {
	cmdActions := &cobra.Command{
		Use:   "action",
		Short: "action commands",
		Long:  "action is used to access action commands",
	}

	cmdActionGet := cmdBuilder(RunCmdActionGet, "get", "get action", writer, "g")
	cmdActions.AddCommand(cmdActionGet)
	addIntFlag(cmdActionGet, doctl.ArgActionID, 0, "Action ID")

	cmdActionList := cmdBuilder(RunCmdActionList, "list", "list actions", writer, "ls")
	cmdActions.AddCommand(cmdActionList)

	return cmdActions
}

// RunCmdActionList run action list.
func RunCmdActionList(ns string, out io.Writer) error {
	client := doctl.DoctlConfig.GetGodoClient()
	f := func(opt *godo.ListOptions) ([]interface{}, *godo.Response, error) {
		list, resp, err := client.Actions.List(opt)
		if err != nil {
			return nil, nil, err
		}

		si := make([]interface{}, len(list))
		for i := range list {
			si[i] = list[i]
		}

		return si, resp, err
	}

	si, err := doctl.PaginateResp(f)
	if err != nil {
		return err
	}

	list := make([]godo.Action, len(si))
	for i := range si {
		list[i] = si[i].(godo.Action)
	}

	return doctl.DisplayOutput(list, out)
}

// RunCmdActionGet runs action get.
func RunCmdActionGet(ns string, out io.Writer) error {
	client := doctl.DoctlConfig.GetGodoClient()
	id := doctl.DoctlConfig.GetInt(ns, doctl.ArgActionID)
	if id < 1 {
		return errors.New("invalid action id")
	}

	a, _, err := client.Actions.Get(id)
	if err != nil {
		return err
	}

	return doctl.DisplayOutput(a, out)
}
