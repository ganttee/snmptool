package snmp

const (
	CdpInterfaceIfIndexOid       = ".1.3.6.1.4.1.9.9.23.1.1.1.1.1"
	CdpInterfaceIfEnableOid      = ".1.3.6.1.4.1.9.9.23.1.1.1.1.2"
	CdpInterfaceIfMsgIntervalOid = ".1.3.6.1.4.1.9.9.23.1.1.1.1.3"
	CdpInterfaceIfGroupOid       = ".1.3.6.1.4.1.9.9.23.1.1.1.1.4"
	CdpInterfaceIfPortOid        = ".1.3.6.1.4.1.9.9.23.1.1.1.1.5"
	CdpInterfaceIfNameOid        = ".1.3.6.1.4.1.9.9.23.1.1.1.1.6"
)

var (
	Enable = []string{"true(1)", "false(2)"}
)

type CDPItem struct {
	IfIndex     int
	Enable      int
	MsgInterval int
	Group       int
	Port        int
	Name        string
}

//ParseEnable
func (i *CDPItem) ParseEnable() string {
	if i.Enable > 0 {
		return Enable[i.Enable-1]
	} else {
		return "invalid value"
	}
}

//GetLLdpLocalTable get loclTable
func GetCDPTable(ip string, communit string) (ips []*CDPItem, err error) {
	oids := []string{CdpInterfaceIfIndexOid, CdpInterfaceIfEnableOid, CdpInterfaceIfMsgIntervalOid, CdpInterfaceIfGroupOid, CdpInterfaceIfPortOid, CdpInterfaceIfNameOid}
	tableRows, err := GetTable(ip, communit, oids)
	if err != nil {
		return
	}
	for _, row := range tableRows {
		item := new(CDPItem)
		item.IfIndex = GetSnmpInt(row[CdpInterfaceIfIndexOid])
		item.Enable = GetSnmpInt(row[CdpInterfaceIfEnableOid])
		item.MsgInterval = GetSnmpInt(row[CdpInterfaceIfMsgIntervalOid])
		item.Group = GetSnmpInt(row[CdpInterfaceIfGroupOid])
		item.Port = GetSnmpInt(row[CdpInterfaceIfPortOid])
		item.Name = GetSnmpString(row[CdpInterfaceIfNameOid])
		ips = append(ips, item)
	}
	return
}
