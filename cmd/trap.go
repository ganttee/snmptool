/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"log"

	"github.com/ganttee/snmptool/snmp"
	"github.com/spf13/cobra"
)

// trapCmd represents the trap command
var trapCmd = &cobra.Command{
	Use:   "trap",
	Short: "Start trap receiver",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		log.Printf("will start trap receiver on port %d ...\n", port)
		r := &snmp.TrapReciver{Address: fmt.Sprintf("0.0.0.0:%d", port)}
		defer r.Close()
		r.Start()
	},
}

func init() {
	rootCmd.AddCommand(trapCmd)
	trapCmd.Flags().Uint16VarP(&port, "port", "p", 162, "trap receiver port")
}
