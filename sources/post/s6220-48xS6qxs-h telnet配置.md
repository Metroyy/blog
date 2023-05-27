---
title: S6220-48XS6QXS-H Telnet配置
desc: S6220-48XS6QXS-H Telnet配置
time: 2023-03-21
tags: S6220,Telnet
---

**配置交换机管理IP**
```bash
enable ---进入特权模式
configure terminal ---进入全局配置模式
enable service telnetserver ---默认开启
interface vlan1 ---进入vlan1接口
ip address 192.168.10.254 255.255.255.0 ---为vlan1接口设置管理IP
exit ---退回到全局配置模式
```

**配置MGMT**
```bash
int mgmt 0 ---配置mgmt0口
ip address 192.168.1.3 255.255.255.0 ---为mgmt配置管理ip
```

**配置telnet密码**
```bash
line vty 0 4 ---进入telnet密码配置模式，0，4表示允许共五个用户同时使用telnet登入到交换机
login ---启用需输入密码才能telnet
password ruijie ---将telnet密码设置为ruijie
exit ---回到全局配置模式
enable password ruijie ---配置进入特权模式的密码为ruijie
end ---退出到特权模式
write ---确认配置正确，保存配置
```