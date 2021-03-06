// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// controlvmCmd represents the controlvm command
var controlvmCmd = &cobra.Command{
	Use:   "controlvm  <uuid|vmname>",
	Short: "Control VM",
	Run: func(cmd *cobra.Command, args []string) {
		var req *http.Request
		var err error

		if len(args) < 2 {
			cmd.Help()
			os.Exit(0)
		}

		host := viper.GetString("server.host")
		port := viper.GetString("server.port")

		// Create client
		client := &http.Client{}

		// Create request
		switch args[1] {
		case "poweroff":
			req, err = http.NewRequest("GET", "http://"+host+":"+port+"/virtualbox/stop/"+args[0], nil)
			assert(err)
		case "nictracefile1":
			req, err = http.NewRequest("GET", "http://"+host+":"+port+"/virtualbox/snapshot/restorecurrent/"+args[0], nil)
			assert(err)
		case "nictrace1":
			req, err = http.NewRequest("GET", "http://"+host+":"+port+"/virtualbox/snapshot/restorecurrent/"+args[0], nil)
			assert(err)
		}

		if req != nil {
			// Fetch Request
			resp, err := client.Do(req)
			assert(err)

			// Read Response Body
			respBody, _ := ioutil.ReadAll(resp.Body)

			// Display Results
			fmt.Print(string(respBody))
		} else {
			cmd.Help()
			os.Exit(1)
		}
	},
}

var poweroffCmd = &cobra.Command{
	Use:   "poweroff",
	Short: "Power off VM",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var nictracefile1Cmd = &cobra.Command{
	Use:   "nictracefile1 <filename>",
	Short: "Write pcap to file",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var nictrace1Cmd = &cobra.Command{
	Use:   "nictrace1 <on|off>",
	Short: "Start/Stop pcap capture",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	RootCmd.AddCommand(controlvmCmd)
	controlvmCmd.AddCommand(poweroffCmd)
	controlvmCmd.AddCommand(nictracefile1Cmd)
	controlvmCmd.AddCommand(nictrace1Cmd)

	controlvmCmd.SetUsageTemplate(controlvmTmplt)
	// poweroffCmd.SetUsageTemplate(snapshotTmplt)
	// nictracefile1Cmd.SetUsageTemplate(snapshotTmplt)
	// nictrace1Cmd.SetUsageTemplate(restoreCurrentTmplt)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// controlvmCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// controlvmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
