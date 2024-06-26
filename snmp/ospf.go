package snmp

import g "github.com/gosnmp/gosnmp"

// OspfIfTable oids
const (
	OspfIfIPAddressOid = ".1.3.6.1.2.1.14.7.1.1"
	OspfIfAdminStatOid = ".1.3.6.1.2.1.14.7.1.5"
	OspfIfStateOid     = ".1.3.6.1.2.1.14.7.1.12"
	OspfIfStatusOid    = ".1.3.6.1.2.1.14.7.1.17"
)

// OspfIfTable oids
const (
	OspfNbrIPAddrOid = ".1.3.6.1.2.1.14.10.1.1"
	OspfNbrRtrIDOid  = ".1.3.6.1.2.1.14.10.1.3"
	OspfNbrStateOid  = ".1.3.6.1.2.1.14.10.1.6"
)

// Values
var (
	IfState      = []string{"down (1)", "loopback (2)", "waiting (3)", "pointToPoint (4)", "designatedRouter (5)", "backupDesignatedRouter (6)", "otherDesignatedRouter (7)"}
	IfAdminState = []string{"enabled (1)", "disabled (2)"}
	NbrState     = []string{"down (1)", "attempt (2)", "init (3)", "twoWay (4)", "exchangeStart (5)", "exchange (6)", "loading (7)", "full (8)"}
)

// OspfItem ospf iftable item
type OspfItem struct {
	OspfIfIPAddress string
	OspfIfAdminStat int
	OspfIfState     int
	OspfIfStatus    int
}

// OspfNbrItem ospf iftable item
type OspfNbrItem struct {
	NbrIPAddress string
	NbrRtrID     string
	NbrState     int
}

// StateStr parse meaningful value
func (o *OspfItem) StateStr() string {
	if o.OspfIfState > 0 {
		return IfState[o.OspfIfState-1]
	}
	return "invalid value"
}

// AdminStateStr parse meaningful value
func (o *OspfItem) AdminStateStr() string {
	if o.OspfIfAdminStat > 0 {
		return IfState[o.OspfIfAdminStat-1]
	}
	return "invalid value"
}

// GetOspfIfTable get loclTable
func GetOspfIfTable(s g.GoSNMP) (ips []*OspfItem, err error) {
	oids := []string{OspfIfIPAddressOid, OspfIfAdminStatOid, OspfIfStateOid, OspfIfStatusOid}
	tableRows, err := GetTable(s, oids)
	if err != nil {
		return
	}
	for _, row := range tableRows {
		item := new(OspfItem)
		item.OspfIfIPAddress = GetSnmpString(row[OspfIfIPAddressOid])
		item.OspfIfAdminStat = GetSnmpInt(row[OspfIfAdminStatOid])
		item.OspfIfState = GetSnmpInt(row[OspfIfStateOid])
		ips = append(ips, item)
	}
	// for _, ipa := range ips {
	// 	v, err := GetOne(ip, communit, IfDescr+"."+strconv.Itoa(ipa.IfIndex))
	// 	if err != nil {
	// 		log.Println(err.Error())
	// 	} else {
	// 		ipa.IfDesc = GetSnmpString(v)
	// 	}
	// }
	return
}

// GetOspfNbrTable get Nbr table
func GetOspfNbrTable(s g.GoSNMP) (rs []*OspfNbrItem, err error) {
	oids := []string{OspfNbrIPAddrOid, OspfNbrRtrIDOid, OspfNbrStateOid}
	tableRows, err := GetTable(s, oids)
	if err != nil {
		return
	}
	for _, row := range tableRows {
		item := new(OspfNbrItem)
		item.NbrIPAddress = GetSnmpString(row[OspfNbrIPAddrOid])
		item.NbrRtrID = GetSnmpString(row[OspfNbrRtrIDOid])
		item.NbrState = GetSnmpInt(row[OspfNbrStateOid])
		rs = append(rs, item)
	}
	return
}

// NbrStateStr RT
func (o *OspfNbrItem) NbrStateStr() string {
	if o.NbrState > 0 {
		return NbrState[o.NbrState-1]
	}
	return "invalid value"
}
