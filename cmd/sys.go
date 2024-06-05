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
	"strings"

	"github.com/ganttee/snmptool/snmp"
	"github.com/ganttee/snmptool/snmp/util"
	"github.com/spf13/cobra"
)

// sysCmd represents the sys command
var sysCmd = &cobra.Command{
	Use:   "sys",
	Short: "Show system brief information",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		d, err := snmp.GetSystem(getSNMPParams())
		if err != nil {
			util.HandleError(err)
		} else {
			data := []string{"sysName: " + d.Name, "sysDescr: " + d.Desc, "sysObjectID: " + d.OId + " (" + util.GetDeviceType(d.OId) + ")", "sysContract: " + d.Contract,
				"sysLocation: " + d.Location, "sysServices: " + d.Services, "sysUpTime: " + d.UpTime}
			fmt.Println(strings.Join(data, "\n"))
		}
	},
}

func init() {
	rootCmd.AddCommand(sysCmd)
}
