package generator

import (
	"fmt"
	"net/url"

	"github.com/Qwertymart/xray-ctl/internal/config"
)

func NewMasterConfig(appConf *config.AppConfig, uuid, privKey, shortID string) *XrayConfig {
	xc := &XrayConfig{
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
				Port:     appConf.XrayParams.Port,
				Protocol: "vless",
				Settings: InboundSettings{
					Clients: []Client{
						{
							Email: "main",
							ID:    uuid,
							Flow:  appConf.XrayParams.Flow,
						},
					},
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
			{Protocol: "freedom", Tag: "direct"},
			{Protocol: "blackhole", Tag: "block"},
		},
		Policy: Policy{
			Levels: map[string]PolicyLevel{
				"0": {Handshake: 3, ConnIdle: 180},
			},
		},
	}

	if appConf.Warp.Enabled {
		xc.Outbounds = append(xc.Outbounds, Outbound{
			Protocol: "socks",
			Tag:      "warp-out",
			Settings: OutboundSettings{
				Servers: []WarpServer{
					{Address: "127.0.0.1", Port: appConf.Warp.Port},
				},
			},
		})

		var warpRule Rule
		if appConf.Warp.FullTunnel {
			warpRule = Rule{
				Type:        "field",
				Network:     "tcp,udp",
				OutboundTag: "warp-out",
			}
		} else {
			warpRule = Rule{
				Type:        "field",
				OutboundTag: "warp-out",
				Domain:      appConf.Warp.RoutingRules.Domains,
				IP:          appConf.Warp.RoutingRules.Ips,
			}
		}
		xc.Routing.Rules = append([]Rule{warpRule}, xc.Routing.Rules...)
	}

	return xc
}

func GenerateVlessLink(appConf *config.AppConfig, ip, uuid, pubKey, sid, name string) string {
	params := url.Values{}
	params.Add("encryption", "none")
	params.Add("flow", appConf.XrayParams.Flow)
	params.Add("security", "reality")
	params.Add("sni", appConf.XrayParams.ServerNames[0])
	params.Add("fp", "chrome")
	params.Add("pbk", pubKey)
	params.Add("sid", sid)
	params.Add("type", "tcp")

	link := fmt.Sprintf("vless://%s@%s:%d?%s#%s",
		uuid,
		ip,
		appConf.XrayParams.Port,
		params.Encode(),
		url.PathEscape(name),
	)

	return link
}
