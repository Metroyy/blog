---
title: MAC完全移除PostgreSQL
desc: MAC完全移除PostgreSQL
time: 2023-03-03
tags: MAC,PostgreSQL
---

**执行PostgreSQL卸载程序（可能提示输入密码）**
```bash
open /Library/PostgreSQL/15/uninstall-postgresql.app
```

**删除PostgreSQL文件夹**
```bash
 sudo rm -rf /Library/PostgreSQL
```

**删除配置文件**
```bash
 sudo rm /etc/postgres-reg.in
```