---
title: phpipam搭建
desc: phpipam搭建
time: 2023-02-13
tags: phpipam
---

### 手动部署phpipam
**前置:无Apache，MySQL，php**
**禁用SELINUX，重启服务器**
``` bash
sed -i '/SELINUX/s/enforcing/disabled/' /etc/selinux/config && reboot
```

**关闭防火墙**
```bash
systemctl stop firewalld
```

**安装依赖包**
```bash
yum install epel-release -y
yum install wget vim net-tools httpd mariadb-server 
php php-cli php-gd php-common php-ldap php-pdo php-pear
php-snmp php-xml php-mysql php-mbstring git -y
```

### 初始化Apache
**修改/etc/httpd/conf/httpd.conf**
```yaml
ServerName localhost:80
```

**允许mod_rewrite URL重写**
```yaml
<Directory "/var/www/html">
 Options FollowSymLinks
 AllowOverride all
 Order allow,deny
 Allow from all
</Directory>
```

**检查配置文件，显示 OK 状态为正常**
```yaml
httpd -t -f /etc/httpd/conf/httpd.conf
```

**修改/etc/php.ini时区**
```yaml
date.timezone = Asia/Shanghai
```

**开机启动Apache**
```bash
systemctl enable httpd
```
**启动Apache**
```bash
systemctl start httpd
```

### 初始化mariadb

**开机启动mariadb**
```bash
systemctl enable mariadb
```

**启动mariadb**
```bash
systemctl start mariadb
```

**初始化mariaDB密码**
```bash
mysql_secure_installation
```

### 初始化phpipam
**下载phpipam**
```bash
cd /var/www/html/
git clone https://github.com/phpipam/phpipam.git .
git checkout 1.4
```

**权限配置**
```bash
chown apache:apache -R /var/www/html/
```

**拷贝文件**
```bash
cp /var/www/html/config.dist.php /var/www/html/config.dist.php.bak
mv /var/www/html/config.dist.php /var/www/html/config.php
```

**编辑/var/www/html/config.php**
```yaml
define('BASE', "/phpipam"); 
```

**重启httpd**
```bash
systemctl restart httpd
```

### docker部署phpipam

```bash
vi docker-compose.yml
#添加配置
```

> ```bash
> # WARNING: Replace the example passwords with secure secrets.
> # WARNING: 'my_secret_phpipam_pass' and 'my_secret_mysql_root_pass'
>
> version: '3'
>
> services:
>   phpipam-web:
>     privileged: true ---特级权限
>     image: phpipam/phpipam-www:latest
>     ports: ---添加一条映射到本地端口
>       - "64080:80"
>     environment:
>       - TZ=Asia/Shanghai
>       - IPAM_DATABASE_HOST=phpipam-mariadb
>       - IPAM_DATABASE_PASS=password ---修改密码
>       - IPAM_DATABASE_WEBHOST=%
>     restart: unless-stopped
>     volumes:
>       - phpipam-logo:/phpipam/css/images/logo
>     depends_on:
>       - phpipam-mariadb
>
>   phpipam-cron:
>     privileged: true ---特级权限
>     image: phpipam/phpipam-cron:latest
>     environment:
>       - TZ=Asia/Shanghai
>       - IPAM_DATABASE_HOST=phpipam-mariadb
>       - IPAM_DATABASE_PASS=password ---修改密码
>       - SCAN_INTERVAL=1h
>     restart: unless-stopped
>     depends_on:
>       - phpipam-mariadb
>
>   phpipam-mariadb:
>     privileged: true ---特级权限
>     image: mariadb:latest
>     ports: ---mysql端口映射
>       - "3306:3306" 
>     environment:
>       - MYSQL_ROOT_PASSWORD=password ---修改mysql密码
>     restart: unless-stopped
>     command: ---添加command 设置拉取的mysql镜像支持中文
>       - mysqld
>       - --character-set-server=utf8mb4
>       - --collation-server=utf8mb4_unicode_ci
>     volumes: ---容器内数据文件存放路径
>       - phpipam-db-data:/var/lib/mysql
> 
> volumes: ---通过卷标将数据持久化
>   phpipam-db-data:
>   phpipam-logo:
> ```

**运行容器**
```bash
docker-compose up -d
```

### 配置phpipam
**打开地址：http://IP/phpipam**
**选择新的phpipam安装**
![img](img/phpipam/1.png)
**安装pfpipam数据库**
![img](img/phpipam/2.png)
**设置数据库**
![img](img/phpipam/3.png)
![img](img/phpipam/4.png)
**填写系统信息**
![img](img/phpipam/5.png)
![img](img/phpipam/6.png)
**登录**
![img](img/phpipam/7.png)
![img](img/phpipam/8.png)
**登录成功后会跳转到主界面**
![img](img/phpipam/9.png)
**修改语言为中文，修改后重新登陆**
![img](img/phpipam/10.png)
![img](img/phpipam/11.png)
![img](img/phpipam/12.png)
**子网创建，删除默认测试子网**
![img](img/phpipam/13.png)
![img](img/phpipam/14.png)
![img](img/phpipam/15.png)
![img](img/phpipam/16.png)
**子网页面**
![img](img/phpipam/17.png)
**选中IP进行编辑**
![img](img/phpipam/18.png)