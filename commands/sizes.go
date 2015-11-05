package commands

import (
	"io"

	"github.com/digitalocean/doctl"
	"github.com/digitalocean/godo"
	"github.com/spf13/cobra"
)

// Size creates the size commands heirarchy.
func Size() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "size",
		Short: "size commands",
		Long:  "size is used to access size commands",
	}

	cmdSizeList := cmdBuilder(RunSizeList, "list", "list sizes", writer)
	cmd.AddCommand(cmdSizeList)

	return cmd
}

// RunSizeList all sizes.
func RunSizeList(ns string, out io.Writer) error {
	client := doctl.DoctlConfig.GetGodoClient()

	f := func(opt *godo.ListOptions) ([]interface{}, *godo.Response, error) {
		list, resp, err := client.Sizes.List(opt)
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

	list := make([]godo.Size, len(si))
	for i := range si {
		list[i] = si[i].(godo.Size)
	}

	return doctl.DisplayOutput(list, out)
}
