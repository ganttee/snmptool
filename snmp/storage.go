package snmp

import g "github.com/gosnmp/gosnmp"

// Const
const (
	WinStorageDescr = ".1.3.6.1.2.1.25.2.3.1.3"
	WinStorageUnits = ".1.3.6.1.2.1.25.2.3.1.4"
	WinStorageSize  = ".1.3.6.1.2.1.25.2.3.1.5"
	WinStorageUsed  = ".1.3.6.1.2.1.25.2.3.1.6"
)

// WinStorage  wap
type WinStorage struct {
	Descr string `json:"descr"`
	Units int    `json:"units"`
	Size  int64  `json:"size"`
	Used  int64  `json:"used"`
}

// GetWinStorage get loclTable
func GetWinStorage(s g.GoSNMP) (table []*WinStorage, err error) {
	oids := []string{WinStorageDescr, WinStorageUnits, WinStorageSize, WinStorageUsed}
	tableRows, err := GetTable(s, oids)
	if err != nil {
		return
	}
	for _, row := range tableRows {
		storage := new(WinStorage)
		storage.Descr = GetSnmpString(row[WinStorageDescr])
		storage.Units = GetSnmpInt(row[WinStorageUnits])
		storage.Size = GetSnmpInt64(row[WinStorageSize])
		storage.Used = GetSnmpInt64(row[WinStorageUsed])
		table = append(table, storage)
	}
	return
}
