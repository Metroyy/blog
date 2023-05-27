---
title: docker部署
desc: docker部署
time: 2023-02-27
tags: docker
---

**更改时区**

```bash
timedatectl set-timezone Asia/Shanghai
```

**检查更改**

```bash
timedatectl
```

**更新源**

```bash
yum -y update
```

**卸载旧版本 docker**

```bash
yum remove docker  docker-common docker-selinux docker-engine
```

**设置docker yum源**

```bash
yum-config-manager --add-repo http://mirrors.aliyun.com/docker-ce/linux/centos/docker-ce.repo
yum makecache fast
```

**安装docker**

```bash
yum -y install docker-ce
```

> **境外安装docker**
>
> ```bash
> curl -fsSL <https://get.docker.com> | bash -s docker
> ```

**启动并加入开机启动**

```bash
systemctl start docker
systemctl enable docker
```

**验证安装**

```bash
docker version
```

**可选扩展Docker-Compose**

```bash
sudo curl -L https://github.com/docker/compose/releases/download/1.16.1/docker-compose-`uname -s`-`uname -m` -o /usr/local/bin/docker-compose ---github安装最新版
sudo curl -L https://get.daocloud.io/docker/compose/releases/download/1.25.1/docker-compose-`uname -s`-`uname -m` -o /usr/local/bin/docker-compose ---如果跑不动则换daocloud
```

```bash
yum -y install epel-release ---再不行换pip安装
yum -y install python-pip
pip --version ---查看版本
pip install --upgrade pip ---更新pip
pip install docker-compose ---安装
```

**设置权限**

```bash
sudo chmod +x /usr/local/bin/docker-compose
```

