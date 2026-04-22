package generator

func NewMasterConfig(uuid, privKey, shortID string, warpEnabled bool) *XrayConfig {
	config := &XrayConfig{
		Log: Log{LogLevel: "warning"},
		Routing: Routing{
			DomainStrategy: "IPIfNonMatch",
			Rules: []Rule{
				{
					Type:        "field",
					Domain:      []string{"geosite:category-ads-all"},
					OutboundTag: "block",
				},
			},
		},
		Inbounds: []Inbound{
			{
				Listen:   "0.0.0.0",
				Port:     443,
				Protocol: "vless",
				Settings: InboundSettings{
					Clients: []Client{
						{
							Email: "main",
							ID:    uuid,
							Flow:  "xtls-rprx-vision",
						},
					},
					Decryption: "none",
				},
				StreamSettings: StreamSettings{
					Network:  "tcp",
					Security: "reality",
					RealitySettings: RealitySettings{
						Show:        false,
						Dest:        "amd.com:443",
						Xver:        0,
						ServerNames: []string{"amd.com", "www.amd.com"},
						PrivateKey:  privKey,
						ShortIds:    []string{shortID},
					},
				},
				Sniffing: Sniffing{
					Enabled:      true,
					DestOverride: []string{"http", "tls"},
				},
			},
		},
		Outbounds: []Outbound{
			{
				Protocol: "freedom",
				Tag:      "direct",
			},
			{
				Protocol: "blackhole",
				Tag:      "block",
			},
		},
		Policy: Policy{
			Levels: map[string]PolicyLevel{
				"0": {Handshake: 3, ConnIdle: 180},
			},
		},
	}

	// Если WARP включен, добавляем outbound и правила маршрутизации
	if warpEnabled {
		config.Outbounds = append(config.Outbounds, Outbound{
			Protocol: "socks",
			Tag:      "warp-out",
			Settings: OutboundSettings{
				Servers: []WarpServer{
					{Address: "127.0.0.1", Port: 4000},
				},
			},
		})

		// Добавляем правила для Google через WARP, как в вашем примере
		warpRule := Rule{
			Type:        "field",
			OutboundTag: "warp-out",
			Domain: []string{
				"geosite:google",
				"geosite:youtube",
				"domain:google.com",
				"domain:googleapis.com",
				"domain:gstatic.com",
				"domain:googlevideo.com",
				"domain:googleusercontent.com",
			},
			IP: []string{"geoip:google"},
		}
		
		// Вставляем правило в начало списка правил маршрутизации
		config.Routing.Rules = append([]Rule{warpRule}, config.Routing.Rules...)
	}

	return config
}