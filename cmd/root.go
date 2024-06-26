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
	"time"

	"github.com/ganttee/snmptool/snmp"
	"github.com/ganttee/snmptool/snmp/util"
	"github.com/gosnmp/gosnmp"
	g "github.com/gosnmp/gosnmp"
	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

const version = "V0.0.3"

var cfgFile string

// Common var
var (
	IP        string
	Community string
	snmpver   string
	oid       string
	Port      uint16
	port      uint16
	//v3 parameters
	UserName     string
	MsgFlags     g.SnmpV3MsgFlags
	AuthProtocol g.SnmpV3AuthProtocol
	AuthPass     string
	PrivProtocol g.SnmpV3PrivProtocol
	PrivPass     string

	SecurityLevel   string
	AuthProtocolStr string
	PrivProtocolStr string
	UseDefaulParam  bool
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
	rootCmd.PersistentFlags().StringVarP(&IP, "ip", "i", "127.0.0.1", "target ip")
	rootCmd.PersistentFlags().Uint16VarP(&Port, "port", "P", 161, "target port")
	rootCmd.PersistentFlags().StringVarP(&Community, "community", "c", "public", "community")
	rootCmd.PersistentFlags().StringVarP(&snmpver, "version", "v", "2c", "snmp version default is '2c'")
	rootCmd.PersistentFlags().StringVarP(&UserName, "username", "u", "", "username")
	rootCmd.PersistentFlags().StringVarP(&SecurityLevel, "level", "l", "noAuthNoPriv", "security level (noAuthNoPriv|authNoPriv|authPriv)")
	rootCmd.PersistentFlags().StringVarP(&AuthProtocolStr, "authProtocol", "a", "MD5", "authentication protocol (MD5|SHA)")
	rootCmd.PersistentFlags().StringVarP(&AuthPass, "authPass", "A", "", "auth password")
	rootCmd.PersistentFlags().StringVarP(&PrivProtocolStr, "privProtocol", "x", "DES", "privacy protocol (DES|AES)")
	rootCmd.PersistentFlags().StringVarP(&PrivPass, "privPass", "X", "", "privacy password")
	rootCmd.PersistentFlags().BoolVarP(&UseDefaulParam, "default", "", false, "use default snmp params, use 'st config' to list items")

}

func getSNMPParams() g.GoSNMP {
	var p SNMPParams
	if UseDefaulParam {
		p = DefaulSNMPParams()
		snmpver = p.SNMPVersion
		Community = p.Community
		SecurityLevel = p.Level
		UserName = p.Username
		AuthProtocolStr = p.AuthProtocol
		AuthPass = p.AuthPass
		PrivProtocolStr = p.PrivProtocol
		PrivPass = p.PrivPass
	}
	msgFlags := g.AuthPriv
	authProto := g.MD5
	privProto := g.DES
	v3Params := &g.UsmSecurityParameters{UserName: UserName}
	switch SecurityLevel {
	case "noAuthNoPriv":
		msgFlags = g.NoAuthNoPriv
	case "authNoPriv":
		msgFlags = g.AuthNoPriv
		v3Params.AuthenticationProtocol = authProto
		v3Params.AuthenticationPassphrase = AuthPass
	case "authPriv":
		msgFlags = g.AuthPriv
		v3Params.AuthenticationProtocol = authProto
		v3Params.AuthenticationPassphrase = AuthPass
		v3Params.PrivacyProtocol = privProto
		v3Params.PrivacyPassphrase = PrivPass
	}

	switch AuthProtocolStr {
	case "SHA":
		authProto = g.SHA
	}
	switch PrivProtocolStr {
	case "AES":
		privProto = g.AES
	}

	switch snmpver {
	case "3":
		return g.GoSNMP{
			Target:             IP,
			Port:               Port,
			Version:            g.Version3,
			SecurityModel:      g.UserSecurityModel,
			MsgFlags:           msgFlags,
			Timeout:            time.Duration(2) * time.Second,
			SecurityParameters: v3Params,
		}

	case "2c":
		return g.GoSNMP{
			Port:      Port,
			Community: Community,
			Version:   g.Version2c,
			Timeout:   time.Duration(2) * time.Second,
			Retries:   3,
			MaxOids:   g.MaxOids,
			Target:    IP,
		}
	case "1":
		return g.GoSNMP{
			Port:      Port,
			Community: Community,
			Version:   g.Version1,
			Timeout:   time.Duration(2) * time.Second,
			Retries:   3,
			MaxOids:   g.MaxOids,
			Target:    IP,
		}
	}
	return *g.Default
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
			util.HandleError(err)
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

// i2S string
func i2S(n int) string {
	return strconv.Itoa(n)
}

// ParseOIDName parse named oid
func ParseOIDName(oid string) (target string) {
	target = snmp.OIDs[oid]
	if target == "" {
		target = oid
	}
	return
}

// ParseSNMPVer RT
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
