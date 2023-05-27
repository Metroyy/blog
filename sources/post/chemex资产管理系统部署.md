---
title: chemex资产管理系统部署
desc: chemex资产管理系统部署
time: 2023-02-21
tags: chemex
---

### 环境

```yaml
git：用于管理版本，部署和升级必要工具。
PHP：仅支持 PHP8.1或以上。
composer：PHP 的包管理工具，用于安装必要的依赖包。
MySQL 5.7：数据库引擎，理论上 MariaDB 10.2 + 兼容支持。

ext-zip：扩展。
ext-json：扩展。
ext-fileinfo：扩展。
ext-ldap：扩展。
ext-bcmath：扩展。
ext-mysqli：扩展。
ext-xml：扩展。
ext-xmlrpc：扩展。
以上扩展安装过程注意版本必须与 PHP 版本一致。
```

### 安装

```bash
mkdir chemex
cd chemex
git clone <https://gitee.com/celaraze/chemex.git> .
git submodule init
git submodule update
cp .env.example .env ---根据.env文件注释配置
composer update -vvv ---确保起始路径在 /public 目录
```

**设置伪静态规则**

```bash
try_files $uri $uri/ /index.php?$args;
```

**启动服务**

```bash
php artisan chemex:install ---报错则检查php扩展插件
http://[domain] ---默认账号密码admin
```

### 参考文献
[chemex资产管理系统](https://github.com/celaraze/chemex)
[chemex第三方文档](https://www.yuque.com/xpesir/chemexdocs)