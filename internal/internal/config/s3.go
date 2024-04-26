package config

type S3 struct {
	// 存储桶地址
	Endpoint string `default:"${ENDPOINT}" validate:"required,url" json:"endpoint,omitempty"`
	// 区域
	Region string `default:"${REGIN=ap-chengdu}" json:"region,omitempty"`
	// 桶
	Bucket string `default:"${BUCKET}" json:"bucket,omitempty"`
	// 分隔符
	Separator string `default:"${SEPARATOR=/}" json:"separator,omitempty"`
}
