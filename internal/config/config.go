package config

type AppConfig struct {
	AppSettings struct {
		UpdateSysctl bool `yaml:"update_sysctl"`
		InstallWarp  bool `yaml:"install_warp"`
	} `yaml:"app_settings"`

	XrayParams struct {
		Port        int      `yaml:"port"`
		Dest        string   `yaml:"dest"`
		ServerNames []string `yaml:"server_names"`
		Flow        string   `yaml:"flow"`
	} `yaml:"xray_params"`

	Warp struct {
		Enabled      bool   `yaml:"enabled"`
		Mode         string `yaml:"mode"`
		Port         int    `yaml:"port"`
		FullTunnel   bool   `yaml:"full_tunnel"`
		RoutingRules struct {
			Domains []string `yaml:"domains"`
			Ips     []string `yaml:"ips"`
		} `yaml:"routing_rules"`
	} `yaml:"warp"`
}
