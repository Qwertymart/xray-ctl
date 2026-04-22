package generator

import "github.com/Qwertymart/xray-ctl/internal/config"

func NewBridgeConfig(appConf *config.AppConfig, localUUID, localPrivKey, localShortID string) *XrayConfig {
	return &XrayConfig{
		Log: Log{LogLevel: "warning"},
		Routing: Routing{
			DomainStrategy: "AsIs",
			Rules: []Rule{
				{
					Type:        "field",
					OutboundTag: "proxy",
					Network:     "tcp,udp",
				},
			},
		},
		Inbounds: []Inbound{
			{
				Listen:   "0.0.0.0",
				Port:     appConf.XrayParams.Port,
				Protocol: "vless",
				Settings: InboundSettings{
					Clients:    []Client{{ID: localUUID, Flow: appConf.XrayParams.Flow}},
					Decryption: "none",
				},
				StreamSettings: StreamSettings{
					Network:  "tcp",
					Security: "reality",
					RealitySettings: RealitySettings{
						Show:        false,
						Dest:        appConf.XrayParams.Dest,
						Xver:        0,
						ServerNames: appConf.XrayParams.ServerNames,
						PrivateKey:  localPrivKey,
						ShortIds:    []string{localShortID},
					},
				},
				Sniffing: Sniffing{Enabled: true, DestOverride: []string{"http", "tls"}},
			},
		},
		Outbounds: []Outbound{
			{
				Protocol: "vless",
				Tag:      "proxy",
				Settings: OutboundSettings{
					Vnext: []Vnext{
						{
							Address: appConf.RemoteServer.Address,
							Port:    appConf.RemoteServer.Port,
							Users: []User{
								{
									ID:         appConf.RemoteServer.UUID,
									Flow:       appConf.XrayParams.Flow,
									Encryption: "none",
								},
							},
						},
					},
				},
				StreamSettings: &StreamSettings{
					Network:  "tcp",
					Security: "reality",
					RealitySettings: RealitySettings{
						ServerName:  appConf.RemoteServer.SNI,
						PublicKey:   appConf.RemoteServer.PublicKey,
						ShortId:     appConf.RemoteServer.ShortID,
						Fingerprint: "chrome",
					},
				},
			},
		},
	}
}
