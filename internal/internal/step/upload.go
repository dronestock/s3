package step

import (
	"context"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/dronestock/s3/internal/internal/config"
	"github.com/goexl/gfx"
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
	"github.com/goexl/log"
)

type Upload struct {
	config *config.Wrapper
	paths  []string
	client *s3.Client
	logger log.Logger
}

func NewUpload(config *config.Wrapper, client *s3.Client, logger log.Logger) *Upload {
	return &Upload{
		config: config,
		client: client,
		logger: logger,
	}
}

func (u *Upload) Runnable() (runnable bool) {
	if paths, ae := gfx.All(u.config.Folder); nil == ae || 0 != len(paths) {
		runnable = true
		u.paths = paths
	}

	return
}

func (u *Upload) Run(ctx *context.Context) (err error) {
	for _, path := range u.paths {
		if err = u.upload(ctx, path); nil != err {
			return
		}
	}

	return
}

func (u *Upload) upload(ctx *context.Context, path string) (err error) {
	if really, re := filepath.Rel(u.config.Folder, path); nil != re {
		err = re
		u.logger.Error("获取文件相对路径出错", field.New("path", path), field.Error(err))
	} else if file, oe := os.Open(path); nil != oe {
		err = oe
	} else {
		err = u.put(ctx, really, file)
	}

	return
}

func (u *Upload) put(ctx *context.Context, path string, body io.Reader) (err error) {
	poi := new(s3.PutObjectInput)
	poi.Bucket = aws.String(u.config.Bucket)
	paths := strings.Split(path, string(filepath.Separator))
	if "" != u.config.Prefix {
		paths = append([]string{u.config.Prefix}, paths...)
	}
	if "" != u.config.Suffix {
		paths = append(paths, u.config.Suffix)
	}
	poi.Key = aws.String(strings.Join(paths, u.config.Separator))
	poi.Body = body

	fields := gox.Fields[any]{
		field.New("path", path),
	}
	if out, poe := u.client.PutObject(*ctx, poi); nil != poe {
		err = poe
		u.logger.Error("上传文件出错", fields.Add(field.Error(err))...)
	} else if nil == out {
		u.logger.Warn("上传文件失败", fields...)
	} else {
		u.logger.Debug("文件上传成功", fields...)
	}

	return
}
