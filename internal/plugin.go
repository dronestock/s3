package internal

import (
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/dronestock/drone"
	"github.com/dronestock/s3/internal/internal/config"
	"github.com/dronestock/s3/internal/internal/step"
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
)

type Plugin struct {
	drone.Base

	// 源
	Source config.Source `default:"${SOURCE}" json:"source,omitempty"`
	// 自身配置
	S3 config.S3 `default:"${S3}" json:"s3,omitempty"`
	// 密钥
	Secret config.Secret `default:"${SECRET}" json:"secret,omitempty"`

	client *s3.Client
}

func New() drone.Plugin {
	return new(Plugin)
}

func (p *Plugin) Config() drone.Config {
	return p
}

func (p *Plugin) Setup() (err error) {
	options := s3.Options{
		Credentials:      credentials.NewStaticCredentialsProvider(p.Secret.Ak, p.Secret.Sk, p.Secret.Session),
		Region:           p.S3.Region,
		EndpointResolver: s3.EndpointResolverFromURL(p.S3.Endpoint),
		UsePathStyle:     true,
		HTTPClient:       p.Http(),
	}
	p.client = s3.New(options)

	return
}

func (p *Plugin) Steps() drone.Steps {
	return drone.Steps{
		drone.NewStep(step.NewUpload(&p.Source, &p.S3, p.client, p.Logger)).Name("上传文件").Build(),
	}
}

func (p *Plugin) Fields() gox.Fields[any] {
	return gox.Fields[any]{
		field.New("source", p.Source),
		field.New("secret", p.Secret),
		field.New("s3", p.S3),
	}
}
