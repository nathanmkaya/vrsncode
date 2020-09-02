/*
Copyright © 2020 Nathan Mkaya <nathanmkaya@gmail.com>

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
package cmd

import (
	"github.com/nathanmkaya/vrsncode/pkg"
	"github.com/spf13/cobra"
)

var step int

var incrementCmd = &cobra.Command{
	Use:   "increment",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		pkg.UpdateVersionCode(pkg.Fetch(keyFile, packageName), step)
	},
}

func init() {
	rootCmd.AddCommand(incrementCmd)
	incrementCmd.Flags().IntVarP(&step, "step", "s", 0, "step")
}
