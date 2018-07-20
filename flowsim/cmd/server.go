// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
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
	// "fmt"
	"github.com/spf13/cobra"
	"github.com/mami-project/trafic/flowsim/tcp"
)

var serverIp string
var serverPort int
var serverSingle bool
var serverTos int

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start an flow server",
	Long: `Start an ABR server.
It will basically sit there and wait for the client to request bunches of data
over a TCP connection`,
	Run: func(cmd *cobra.Command, args []string) {
		flow.Server(serverIp, serverPort, serverSingle, serverTos)
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
	serverCmd.PersistentFlags().StringVarP(&serverIp, "ip", "I", "127.0.0.1", "IP address or host name bound to the flowsim server")
	serverCmd.PersistentFlags().IntVarP(&serverPort, "port", "p", 8081, "TCP port bound to the flowsim server")
	serverCmd.PersistentFlags().BoolVarP(&serverSingle,"one-off", "1", false, "Just accept one connection and quit (default is run until killed)")
	serverCmd.PersistentFlags().IntVarP(&serverTos, "TOS", "T", 128, "Value of the TOS field in the IP layer")
}
