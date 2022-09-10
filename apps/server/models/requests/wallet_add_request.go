package requests

type WalletAddRequest struct {
	BaseRequest `mapstructure:",squash"`
	Key         string `json:"key" mapstructure:"key"`
	Wallet      string `json:"wallet" mapstructure:"wallet"`
}
