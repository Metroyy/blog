---
title: VGPU配置
desc: VGPU配置
time: 2023-04-12
tags: VGPU
---


### esxi安装vgpu驱动

**设置显卡禁止直通**
![img](../../img/vgpu/2.png)

**上传驱动vib文件**
![img](../../img/vgpu/3.png)

**进入维护模式**
![img](../../img/vgpu/4.png)

**设置读写权限**
![img](../../img/vgpu/5.png)

**安装驱动**
```bash
esxcli software vib install -v (vib文件路径) ---安装驱动
esxcli software vib update -v (vib文件路径) ---更新驱动
```
![img](../../img/vgpu/6.png)

**验证驱动**
```bash
nvidia-smi
nvidia-smi -e 0 ---若ECC不为OFF
```
![img](../../img/vgpu/7.png)
**若红框为Xorg**
**设置显卡共享类型**
![img](../../img/vgpu/8.png)

**切换直接共享**
![img](../../img/vgpu/9.png)

**再次验证驱动，如图即可**
```bash
nvidia-smi
lspci | grep NVIDIA
```
![img](../../img/vgpu/10.png)
![img](../../img/vgpu/11.png)

### 安装授权服务器licserve
**lic版本对照表**
![img](../../img/vgpu/1.png)

**配置CentOS虚拟机**
```bash
8a:50:13:0c:ae:06 ---修改虚拟网卡MAC
2018-07-19 18:40:00 ---修改系统时间并关闭ntp同步
yum update -y ---更新系统
yum intsall java -y ---安装java
yum install tomcat tomcat-webapps -y ---安装tomcat 
```

**本机上传安装文件到虚拟机**
```bash
scp setup.bin root@IP:文件夹
```
![img](../../img/vgpu/12.png)

**验证上传**
![img](../../img/vgpu/13.png)

**安装licserve**
```bash
./setup.bin ---运行安装文件
/usr/share/tomcat ---提示输入tomcat路径时输入默认路径
```
![img](../../img/vgpu/14.png)

**服务设置**
```bash
systemctl stop firewalld ---停止防护墙
systemctl disable  firewalld  ---禁用防火墙自启服务
systemctl start tomcat.service ---启动tomcat服务
systemctl enable tomcat.service ---设置tomcat服务自启
systemctl status tomcat.service ---查看服务状态
systemctl list-unit-files ---查看所有自启服务  we前后 q退出
```

**验证授权服务器**
```bash
http://ip:8080/licserver ---浏览器进入管理页面
```
![img](../../img/vgpu/15.png)

**导入vib授权文件**
![img](../../img/vgpu/16.png)

**验证授权**
![img](../../img/vgpu/17.png)

**授权服务器打开NTP自动同步**
```bash
yum -y install chrony ---安装软件
systemctl enable chronyd  ---开机自启
systemctl start chronyd ---启动
timedatectl status ---查看时间同步状态
timedatectl set-ntp true ---开启网络时间同步
date ---查看系统时间是否正确
```

**下载对应的guid驱动，并安装**
![img](../../img/vgpu/18.png)


**进入驱动面板，输入授权服务器ip及7070授权端口**
![img](../../img/vgpu/19.png)

**管理面板验证授权成功**
![img](../../img/vgpu/20.png)
![img](../../img/vgpu/21.png)
