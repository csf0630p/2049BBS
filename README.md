2049bbs，一个无需手机号和邮箱即可注册发言的论坛。Fork 自 [goyoubbs](https://github.com/ego008/goyoubbs)。

## 本地开发

安装 [go](https://golang.org/dl/)，然后 clone 本仓库。

```bash
go get -v github.com/terminus2049/2049bbs
```

然后 cd 到相应目录，一般是 `go/src/github.com/terminus2049/2049bbs`。

```bash
go run main.go
```

然后在浏览器打开 `127.0.0.1:8082` 即可，或者直接编译，运行 `sudo ./2049bbs`，

## 部署

编译二进制文件 `go build`，~~非 Linux 平台为交叉编译 `GOOS=linux GOARCH=amd64 go build`~~，由于使用了 gojieba 分词引擎，不能跨平台编译，请使用在线api功能、移除相关组件后再尝试跨平台编译。

将编译好的二进制文件与 config、static 和 view 三个文件夹的文件放在同一个文件夹内，运行 `./goyoubbs`。

服务器配置：在生产环境中，建议打开 `https`，把 `config.yaml` 中 `HttpsOn: false` 改为 `true`。也可以自行申请 cloudflare 证书，相应配置可以参考 [config-2049.yaml](https://github.com/Terminus2049/2049BBS/blob/master/config/config-2049.yaml).

## 备份

需要备份 `mydata.db` 和 `/static/avatar` 文件。
