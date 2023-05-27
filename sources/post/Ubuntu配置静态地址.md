---
title: Ubuntu配置静态地址
desc: Ubuntu配置静态地址
time: 2023-04-07
tags: Ubuntu
---

**Ubuntu自17.10后，放弃在/etc/network/interfaces里配置IP**
**改为/etc/netplan/00-install-config.yaml**
**查看网卡配置信息**
```yaml
ifconfig
```

**修改配置文件**
```yaml
vi /etc/netplan/00-install-config.yaml
```

```yaml
network:
  ethernets:
    ens160:
      addresses:
      - 10.0.2.20/24 ---静态IP地址和掩码
      gateway4: 10.0.2.254 ---网关地址
      nameservers:
        addresses:
        - 114.114.114.114,8.8.8.8 ---DNS服务器地址，多个DNS用英文逗号隔开
        search: []
  version: 2
```

**使配置生效**
```yaml
sudo netplan apply
```