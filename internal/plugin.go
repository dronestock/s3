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
	config.Wrapper
	config.Secret

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
		Credentials:      credentials.NewStaticCredentialsProvider(p.Ak, p.Sk, p.Session),
		Region:           p.Region,
		EndpointResolver: s3.EndpointResolverFromURL(p.Endpoint),
		UsePathStyle:     true,
		HTTPClient:       p.Http(),
	}
	p.client = s3.New(options)

	return
}

func (p *Plugin) Steps() drone.Steps {
	return drone.Steps{
		drone.NewStep(step.NewUpload(&p.Wrapper, p.client, p.Logger)).Name("上传文件").Build(),
	}
}

func (p *Plugin) Fields() gox.Fields[any] {
	return gox.Fields[any]{
		field.New("folder", p.Folder),
		field.New("secret", p.Secret),
		field.New("endpoint", p.Endpoint),
		field.New("separator", p.Separator),
		field.New("clear", p.Clear),
		field.New("prefix", p.Prefix),
		field.New("suffix", p.Suffix),
	}
}
