/*
Copyright 2016 The Doctl Authors All rights reserved.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package commands

// cmdOption allow configuration of a command.
type cmdOption func(*Command)

// aliasOpt adds aliases for a command.
func aliasOpt(aliases ...string) cmdOption {
	return func(c *Command) {
		if c.Aliases == nil {
			c.Aliases = []string{}
		}

		for _, a := range aliases {
			c.Aliases = append(c.Aliases, a)
		}
	}
}

// displayerType sets the columns for display for a command.
func displayerType(d Displayable) cmdOption {
	return func(c *Command) {
		c.fmtCols = d.Cols()
	}
}

// hiddenCmd make a command hidden.
func hiddenCmd() cmdOption {
	return func(c *Command) {
		c.Hidden = true
	}
}

// docCategories adds documentation categories to a command.
func docCategories(categories ...string) cmdOption {
	return func(c *Command) {
		c.DocCategories = categories
	}
}
