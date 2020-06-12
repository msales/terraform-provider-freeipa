package freeipa

import ipa "github.com/msales/goipa"

type Config struct {
	Host     string
	Username string
	Password string
	BaseDN   string
	Insecure bool
}

type Connection struct {
	client     *ipa.Client
	ldapClient *ipa.LdapClient
	session    string
}

func (cfg *Config) NewConnection() (*Connection, error) {
	c := ipa.Client{KeyTab: "", Host: cfg.Host, Insecure: cfg.Insecure}

	sess, err := c.Login(cfg.Username, cfg.Password)

	if err != nil {
		return nil, err
	}

	ldapClient, err := ipa.LdapConnect(cfg.Host, cfg.BaseDN, cfg.Username, cfg.Password)

	return &Connection{client: &c, ldapClient: ldapClient, session: sess}, nil
}
