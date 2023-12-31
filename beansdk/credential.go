package beansdk

type Credential struct {
	SecretID  string
	SecretKey string
}

func newCredential(secretID, secretKey string) *Credential {
	return &Credential{secretID, secretKey}
}

func (c *Credential) GetSecretId() string {
	if c.SecretID == "" {
		return "visitor"
	}
	return c.SecretID
}
