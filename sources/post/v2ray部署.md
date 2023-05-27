---
title: V2ray部署
desc: V2ray部署
time: 2023-02-27
tags: V2ray
---

**安装v2ray**

```bash
bash <(curl -sL https://raw.githubusercontent.com/hiifeng/v2ray/main/install_v2ray.sh)
```

**v2rayUI后台管理界面安装**

```bash
bash <(curl -Ls https://blog.sprov.xyz/v2-ui.sh)
```

> **若提示curl: command not found**
>
> ```bash
> ubuntu/debian: apt-get update -y && apt-get install curl -y 
> centos: yum update -y && yum install curl -y
> ```

**修改时区**

```bash
cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
```

**查看修改是否完成**

```bash
date -R
```

**启动v2ray**

```bash
sudo systemctl restart v2ray
```

**查看状态**

```bash
service v2ray status
```

### 参考文献
[Xray核心](https://github.com/XTLS/Xray-core)
[v2ray Win客户端](https://github.com/2dust/v2rayN)
[v2ray Android客户端](https://github.com/2dust/v2flyNG)