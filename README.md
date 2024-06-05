## Snmptool
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/9e720e9e45f6456abfa18d27d0b8136c)](https://app.codacy.com/app/ganttee/snmptool?utm_source=github.com&utm_medium=referral&utm_content=ganttee/snmptool&utm_campaign=Badge_Grade_Dashboard)
[![ASL 2.0](https://img.shields.io/hexpm/l/plug.svg)](https://github.com/ganttee/snmptool/blob/master/LICENSE)
[![Build Status](https://travis-ci.org/ganttee/snmptool.svg?branch=master)](https://travis-ci.org/ganttee/snmptool)
[![](http://shields.katacoda.com/katacoda/smark/count.svg)](https://www.katacoda.com/smark/scenarios/snmptool)
[![Go Report Card](https://goreportcard.com/badge/github.com/ganttee/snmptool)](https://goreportcard.com/report/github.com/ganttee/snmptool)


## Usage
```
Simple snmp tool

Usage:
  snmptool [command]

Available Commands:
  bgp         Show BGP brief information
  cdp         Show CDP brief infromation
  get         Execute SNMP GET function
  help        Help about any command
  iftable     Show Iftable brief information
  interface   Show interface biref information
  ipaddress   Show ip address table
  lldp        Show LLDP brief information
  oids        Show Name-OID mapping
  ospf        Show OSPF biref information
  storage     Show storage biref information
  sys         Show system brief information
  trap        Start trap receiver
  version     Show version
  walk        Execute SNMP WALK

Flags:
  -c, --community string   community (default "public")
  -h, --help               help for snmptool
  -i, --ip string          target ip (default "127.0.0.1")

Use "snmptool [command] --help" for more information about a command.
```