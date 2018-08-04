
2049bbs，一个无需手机号和邮箱即可注册发言的论坛。

## 安装

2049bbs Fork 自 [goyoubbs](https://github.com/ego008/goyoubbs)，本仓库对主程序 `goyoubbs` 不做任何修改，只维护 `config`、`static` 和 `view` 等配置文件和样式。

首先 clone 本仓库，然后在 [goyoubbs 仓库](https://github.com/ego008/goyoubbs/releases) 下载对应的主程序。

```bash
git clone https://github.com/Terminus2049/2049BBS.git
cd 2049BBS

## linus 系统 64位
wget https://github.com/ego008/goyoubbs/releases/download/master/goyoubbs-linux-amd64.zip
```

然后运行 `sudo ./goyoubbs`，在浏览器打开 `127.0.0.1/` 即可。

在生产环境中，建议打开 `https`，把 `config.yaml` 中 `HttpsOn: false` 改为 `true`。
