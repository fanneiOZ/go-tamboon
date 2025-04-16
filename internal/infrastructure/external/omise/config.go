package omise

type ApiEndpoints struct {
	Api   string
	Vault string
}

type ApiConfig struct {
	SecretKey string
	PublicKey string
	Endpoints ApiEndpoints
}
