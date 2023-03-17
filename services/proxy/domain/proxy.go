package domain

type Proxy struct {
	ProxyHeader    ProxyHeader     `json:"proxy_header"`
	ProxyBodyItems []ProxyBodyItem `json:"proxy_body_items"`
}
