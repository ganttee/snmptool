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
	"os"
	"strconv"

	"github.com/soniah/gosnmp"

	"github.com/smarkm/snmptool/snmp"
	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

const version = "V0.0.2"

var cfgFile string

//Common var
var (
	IP        string
	Community string
	snmpver   string
	oid       string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "snmptool",
	Short: "Simple snmp tool",
	Long:  ``,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	UseGlobleFlags(rootCmd)

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".snmptool" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".snmptool")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

//UseGlobleFlags use
func UseGlobleFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&IP, "ip", "i", "127.0.0.1", "target ip")
	cmd.Flags().StringVarP(&Community, "community", "c", "public", "community")
	cmd.Flags().StringVarP(&snmpver, "snmpver", "v", "2c", "snmp version")

}

//i2S string
func i2S(n int) string {
	return strconv.Itoa(n)
}

//ParseOIDName parse named oid
func ParseOIDName(oid string) (target string) {
	target = snmp.OIDs[oid]
	if target == "" {
		target = oid
	}
	return
}

//ParseSNMPVer RT
func ParseSNMPVer() gosnmp.SnmpVersion {
	switch snmpver {
	case "1":
		return gosnmp.Version1
	case "2c":
		return gosnmp.Version2c
	case "3":
		return gosnmp.Version3

	}
	return gosnmp.Version2c
}
