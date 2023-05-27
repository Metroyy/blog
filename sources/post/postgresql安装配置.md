---
title: Debian PostgreSQL安装配置
desc: Debian PostgreSQL安装配置
time: 2023-03-14
tags: Debian,PostgreSQL
---

**切换为root用户**

```bash
 sudo su
```

**更新APT包索引**

```bash
 sudo apt update
```

**安装 PostgreSQL服务**

```bash
 sudo apt install postgresql
```

**安装校验**

```bash
 sudo -u postgres psql -c "SELECT version();"
```

**默认情况下PostgreSQL服务器仅在本地监听127.0.0.1:5432
如果需要远程连接PostgreSQL服务器，则需要将服务器设置为公网IP上监听，并编辑配置文件**

**配置文件路径**

```bash
 vi /etc/postgresql/13/main/postgresql.conf
```

**修改如下内容**

```bash
- CONECTIONS AND AUTHENTICATION
listen_addresses =  ’*’  #what IP address(es) to listen on;
```

**保存文件并重启PostgreSQL**

```bash
sudo service postgresql restart
```

**使用ss验证更改**

```bash
 ss -nlt | grep 5432
LISTEN 0  128   0.0.0.0:5432   0.0.0.0:*
LISTEN 0  128    [::]:5432    [::]:*
```

**编辑pg\_hba.conf文件使其接受远程登录**

```bash
 vi /etc/postgresql/13/main/pg_hba.conf
```

**找到#IPv4 local connections插入命令**

```bash
 host  all  all  0.0.0.0/0  md5
```

**切换Postgres用户**

```bash
 su postgres
```

**链接数据库**

```bash
 psql
```

**修改postersql默认密码**

```bash
 \password postgres 根据命令提示输入密码
```

**创建数据库用户**

```bash
 CREATE USER zhangsan WITH PASSWORD ‘zhangsan123;
```

**创建数据库**

```bash
 CREATE DATABASE test_db OWNER zhangsan;
```


