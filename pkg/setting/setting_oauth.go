package setting

type OAuthInfo struct {
	ClientId, ClientSecret string
	Scopes                 []string
	AuthUrl, TokenUrl      string
	Enabled                bool
	EmailAttributeName     string
	EmailAttributePath     string
	RoleAttributePath      string
	AllowedDomains         []string
	HostedDomain           string
	ApiUrl                 string
	AllowSignup            bool
	Name                   string
	TlsClientCert          string
	TlsClientKey           string
	TlsClientCa            string
	TlsSkipVerify          bool
	GroupMappingsFile      string
	GroupMappings          []OAuthGroupMapping
}

type OAuthGroupMapping struct {
	Filter         string `toml:"filter"`
	Role           string `toml:"role"`
	OrgID          int64  `toml:"org_id"`
	IsGrafanaAdmin *bool  `toml:"grafana_admin"`
}

type OAuther struct {
	OAuthInfos map[string]*OAuthInfo
}

var OAuthService *OAuther
