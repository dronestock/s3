package config

type Secret struct {
	// 授权，类型于用户名
	Ak string `default:"${AK}" validate:"required" json:"ak,omitempty"`
	// 授权，类型于密码
	Sk string `default:"${SK}" validate:"required" json:"sk,omitempty"`
	// 会话
	Session string `json:"session,omitempty"`
}
