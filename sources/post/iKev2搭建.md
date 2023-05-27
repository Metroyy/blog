---
title: iKev2搭建
desc: iKev2搭建
time: 2023-02-19
tags: iKev2
---



### 安装

**系统环境**
`Ubuntu 20.0.4`

**更新源**
```bash
 sudo apt-get update
```

**拉取源代码**
```bash
 wget <https://get.vpnsetup.net> -O vpn.sh && sudo sh vpn.sh
```

### **命令**
**列出已有的客户端**
```bash
 sudo ikev2.sh --listclients
```

**添加客户端证书**
```bash
 sudo ikev2.sh --addclient client name
```

**导出已有的客户端配置**
```bash
 sudo ikev2.sh --exportclient client name
```

**吊销客户端证书**
```bash
 sudo ikev2.sh --revokeclient client name
```

**删除客户端**
```bash
 sudo ikev2.sh --deleteclient client name
```

**重启IPsec服务**
```bash
 service ipsec restart service xl2tpd restart
```

**检查 IPsec VPN 服务器状态**
```bash
 ipsec status
```

**查看当前已建立的 VPN 连接**
```bash
 ipsec trafficstatus
```

**移除iKev2**
```bash
 sudo ikev2.sh --removeikev2
```

**检查 Libreswan (IPsec) 和 xl2tpd 日志是否有错误**
```bash
grep pluto /var/log/auth.log ---Ubuntu & Debian
grep xl2tpd /var/log/syslog

grep pluto /var/log/secure ---CentOS/RHEL, Rocky Linux, AlmaLinux, Oracle Linux & Amazon Linux 2
grep xl2tpd /var/log/messages

grep pluto /var/log/messages ---Alpine Linux
grep xl2tpd /var/log/messages
```

### 参考文献
[ikev2](https://github.com/hwdsl2/setup-ipsec-vpn)
[错误码对照](https://github.com/hwdsl2/setup-ipsec-vpn/blob/master/docs/clients-zh.md#%E6%A3%80%E6%9F%A5%E6%97%A5%E5%BF%97%E5%8F%8A-vpn-%E7%8A%B6%E6%80%81)