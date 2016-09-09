/*
Copyright (C) 2016 Red Hat, Inc.

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
	"fmt"
	"os"

	"github.com/docker/machine/libmachine"
	"github.com/golang/glog"
	"github.com/jimmidyson/minishift/pkg/minikube/cluster"
	"github.com/jimmidyson/minishift/pkg/minikube/constants"
	"github.com/pkg/browser"
	"github.com/spf13/cobra"
)

var (
	consoleURLMode bool
)

// consoleCmd represents the console command
var consoleCmd = &cobra.Command{
	Use:   "console",
	Short: "Opens/displays the OpenShift console URL for your local cluster",
	Long:  `Opens/displays the OpenShift console URL for your local cluster`,
	Run: func(cmd *cobra.Command, args []string) {
		api := libmachine.NewClient(constants.Minipath, constants.MakeMiniPath("certs"))
		defer api.Close()
		url, err := cluster.GetConsoleURL(api)
		if err != nil {
			glog.Errorln("Error accessing the OpenShift console (is minishift running?): Error: ", err)
			os.Exit(1)
		}
		if consoleURLMode {
			fmt.Fprintln(os.Stdout, url)
		} else {
			fmt.Fprintln(os.Stdout, "Opening OpenShift console in default browser...")
			browser.OpenURL(url)
		}
	},
}

func init() {
	consoleCmd.Flags().BoolVar(&consoleURLMode, "url", false, "Display the OpenShift console in the CLI instead of opening it in the default browser")
	RootCmd.AddCommand(consoleCmd)
}
