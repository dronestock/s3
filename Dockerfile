FROM storezhang/alpine


LABEL author="storezhang<华寅>" \
    email="storezhang@gmail.com" \
    qq="160290688" \
    wechat="storezhang" \
    description="Drone持续集成系统S3插件，通过S3插件来完成文件上传、删除、修改等操作，可以方便的在CI/CD流程中集成"


# 复制文件
COPY s3 /usr/local/bin


RUN set -ex \
    \
    \
    \
    # 增加执行权限
    && chmod +x /usr/local/bin/s3 \
    \
    \
    \
    && rm -rf /var/cache/apk/*



# 执行命令
ENTRYPOINT /usr/local/bin/s3
