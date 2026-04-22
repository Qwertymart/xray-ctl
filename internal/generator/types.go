package generator

type XrayConfig struct {
	Log      Log        `json:"log"`
	Routing  Routing    `json:"routing"`
	Inbounds []Inbound  `json:"inbounds"`
	Outbounds []Outbound `json:"outbounds"`
	Policy   Policy     `json:"policy"`
}

type Log struct {
	LogLevel string `json:"loglevel"`
}

type Routing struct {
	DomainStrategy string `json:"domainStrategy"`
	Rules          []Rule `json:"rules"`
}

type Rule struct {
	Type        string   `json:"type"`
	Domain      []string `json:"domain,omitempty"`
	IP          []string `json:"ip,omitempty"`
	Network     string   `json:"network,omitempty"`
	OutboundTag string   `json:"outboundTag"`
}

type Inbound struct {
	Listen         string         `json:"listen"`
	Port           int            `json:"port"`
	Protocol       string         `json:"protocol"`
	Settings       InboundSettings `json:"settings"`
	StreamSettings StreamSettings  `json:"streamSettings"`
	Sniffing       Sniffing       `json:"sniffing"`
}

type InboundSettings struct {
	Clients    []Client `json:"clients"`
	Decryption string   `json:"decryption"`
}

type Client struct {
	Email string `json:"email,omitempty"`
	ID    string `json:"id"`
	Flow  string `json:"flow,omitempty"`
}

type StreamSettings struct {
	Network         string          `json:"network"`
	Security        string          `json:"security"`
	RealitySettings RealitySettings `json:"realitySettings"`
}

type RealitySettings struct {
    Show        bool     `json:"show,omitempty"`
    Dest        string   `json:"dest,omitempty"`        
    Xver        int      `json:"xver,omitempty"`        
    ServerNames []string `json:"serverNames,omitempty"`
    PrivateKey  string   `json:"privateKey,omitempty"`
    PublicKey   string   `json:"publicKey,omitempty"` 
    ShortIds    []string `json:"shortIds,omitempty"`
    ShortId     string   `json:"shortId,omitempty"` 
    Fingerprint string   `json:"fingerprint,omitempty"`
    ServerName  string   `json:"serverName,omitempty"`
}

type Sniffing struct {
	Enabled      bool     `json:"enabled"`
	DestOverride []string `json:"destOverride"`
}

type Outbound struct {
	Protocol string           `json:"protocol"`
	Tag      string           `json:"tag"`
	Settings OutboundSettings `json:"settings,omitempty"`
	StreamSettings *StreamSettings `json:"streamSettings,omitempty"`
}

type OutboundSettings struct {
	Servers []WarpServer `json:"servers,omitempty"`
	Vnext   []Vnext      `json:"vnext,omitempty"`
}

type Vnext struct {
	Address string `json:"address"`
	Port    int    `json:"port"`
	Users   []User `json:"users"`
}

type User struct {
	ID         string `json:"id"`
	Flow       string `json:"flow"`
	Encryption string `json:"encryption"`
}

type WarpServer struct {
	Address string `json:"address"`
	Port    int    `json:"port"`
}

type Policy struct {
	Levels map[string]PolicyLevel `json:"levels"`
}

type PolicyLevel struct {
	Handshake int `json:"handshake"`
	ConnIdle  int `json:"connIdle"`
}