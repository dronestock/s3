package config

type Wrapper struct {
	// 本地上传目录
	Folder string `default:"${FOLDER=.}" json:"folder,omitempty"`
	// 存储桶地址
	Endpoint string `default:"${ENDPOINT}" validate:"required,url" json:"endpoint,omitempty"`
	// 区域
	Region string `default:"${REGIN=ap-chengdu}" json:"region,omitempty"`
	// 桶
	Bucket string `default:"${BUCKET}" json:"bucket,omitempty"`

	// 分隔符
	Separator string `default:"${SEPARATOR=/}" json:"separator,omitempty"`
	// 是否清空存储桶
	Clear *bool `default:"${CLEAR=true}" json:"clear,omitempty"`
	// 路径前缀，所有文件上传都会在这上面加上前缀
	Prefix string `default:"${PREFIX}" json:"prefix,omitempty"`
	// 路径后缀，所有文件上传都会在这上面加上后缀
	Suffix string `default:"${SUFFIX}" json:"suffix,omitempty"`
}
