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

	//"github.com/smarkm/snmptool/snmp"
	"github.com/smarkm/snmptool/snmp/util"
	"github.com/soniah/gosnmp"
	"github.com/spf13/cobra"
)

func printValue(pdu gosnmp.SnmpPDU) error {
	fmt.Printf("%s = ", pdu.Name)

	switch pdu.Type {
	case gosnmp.OctetString:
		b := pdu.Value.([]byte)
		fmt.Printf("STRING: \"%s\"\n", string(b))
	case gosnmp.IPAddress:
		b := pdu.Value.(string)
		fmt.Printf("IpAddress: \"%s\"\n", string(b))
	case gosnmp.Boolean:
		fallthrough
	case gosnmp.Integer:
		b := pdu.Value.(int)
		fmt.Printf("INTEGER: %d\n", b)
	case gosnmp.BitString:
		fallthrough
	case gosnmp.Null:
		fallthrough
	case gosnmp.ObjectIdentifier:
		fallthrough
	case gosnmp.ObjectDescription:
		fallthrough
	case gosnmp.Counter32:
		fallthrough
	case gosnmp.Gauge32:
		fallthrough
	case gosnmp.TimeTicks:
		fallthrough
	case gosnmp.Opaque:
		fallthrough
	case gosnmp.NsapAddress:
		fallthrough
	case gosnmp.Counter64:
		fallthrough
	case gosnmp.Uinteger32:
		fallthrough
	case gosnmp.OpaqueFloat:
		fallthrough
	case gosnmp.OpaqueDouble:
		fallthrough
	case gosnmp.NoSuchObject:
		fallthrough
	case gosnmp.NoSuchInstance:
		fallthrough
	case gosnmp.EndOfMibView:
		fallthrough
	default:
		fmt.Printf("TYPE %d: %d\n", pdu.Type, gosnmp.ToBigInt(pdu.Value))
	}
	return nil
}

// walkCmd represents the walk command
var walkCmd = &cobra.Command{
	Use:   "walk",
	Short: "Execute SNMP WALK",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		s := getSNMPParams()
		err := s.Connect()
		if err != nil {
			return
		}
		defer s.Conn.Close()
		target := ParseOIDName(oid)
		err = s.Walk(target, printValue)
		if err != nil {
			util.HandleError(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(walkCmd)
	walkCmd.Flags().StringVarP(&oid, "oid", "o", "", "root oid")
	getCmd.MarkFlagRequired("oid")
}
