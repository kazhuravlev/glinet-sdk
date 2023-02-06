package glinet

type Modem struct {
	Ports       []string `json:"ports"`
	ModemID     int      `json:"modem_id"`
	DataPort    string   `json:"data_port"`
	ControlPort string   `json:"control_port"`
	QmiPort     string   `json:"qmi_port"`
	Name        string   `json:"name"`
	Imei        string   `json:"IMEI"`
	Bus         string   `json:"bus"`
	HwVersion   string   `json:"hw_version"`
	SimNum      string   `json:"sim_num"`
	Mnc         string   `json:"mnc"`
	Mcc         string   `json:"mcc"`
	Carrier     string   `json:"carrier"`
	Up          string   `json:"up"`
	SIMStatus   int      `json:"SIM_status"`
	Operators   []string `json:"operators"`
}

type GetModemInfoResp struct {
	Passthrough           bool    `json:"passthrough"`
	HintModifyWifiChannel int     `json:"hint_modify_wifi_channel"`
	Modems                []Modem `json:"modems"`
}

type RouterClient struct {
	Remote     bool   `json:"remote"`
	Mac        string `json:"mac"`
	Favorite   bool   `json:"favorite"`
	IP         string `json:"ip"`
	Up         string `json:"up"`
	Down       string `json:"down"`
	TotalUp    string `json:"total_up"`
	TotalDown  string `json:"total_down"`
	QosUp      string `json:"qos_up"`
	QosDown    string `json:"qos_down"`
	Blocked    bool   `json:"blocked"`
	Iface      string `json:"iface"`
	Name       string `json:"name"`
	OnlineTime string `json:"online_time"`
	Alive      string `json:"alive"`
	NewOnline  bool   `json:"new_online"`
	Online     bool   `json:"online"`
	Vendor     string `json:"vendor"`
	Node       string `json:"node"`
}

type GetClientListResp struct {
	Clients []RouterClient `json:"clients"`
}

type GetNetworkStatusResp struct {
	Reachable  bool `json:"reachable"`
	RebootFlag bool `json:"reboot_flag"`
}
