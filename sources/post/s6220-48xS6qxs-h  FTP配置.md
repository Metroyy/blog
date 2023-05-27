---
title: S6220-48XS6QXS-H FTP配置
desc: S6220-48XS6QXS-H FTP配置
time: 2023-03-21
tags: S6220,FTP
---

**FTP配置**
```bash
enable ---进入特权模式
configure terminal ---进入全局配置模式
ftp-server enable ---开启ftp服务
ftp-server topdir flash:/ ---限制FTP客户端能够进行文件读写操作的目录范围
ftp-server username ruijie password ruijie ---配置FTP账号密码
```