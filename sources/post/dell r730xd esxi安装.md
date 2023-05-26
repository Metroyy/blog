---
title: "DELL R730XD安装ESXI"
desc: "DELL R730XD安装ESXI"
time: "2023-03-15"
tags: "R730XD,ESXI"
---

### 本机环境
**型号：DELL R730XD**
**CPU：2673v4 * 2**
**内存：32G * 16**
**硬盘：SAS 6T * 12**
**显卡：Tesla p40**

### 软件版本
ESXI 7.0U3

### 安装ESXI 7.0U3g
**idrac口进入管理面板**
![img](../../img/esxi/3.png)

**欢迎页面，回车下一步**
![img](../../img/esxi/4.png)

**用户协议，F11下一步**
![img](../../img/esxi/5.png)

**选择ESXI安装位置，回车下一步**
![img](../../img/esxi/6.png)

**键盘布局，保持默认值，回车下一步**
![img](../../img/esxi/7.png)

**输入root密码，回车下一步**
![img](../../img/esxi/8.png)

**F11开始安装**
![img](../../img/esxi/9.png)

**等待安装完成后拔掉u盘，回车重启**
![img](../../img/esxi/10.png)
![img](../../img/esxi/11.png)

**加载ESXI**
![img](../../img/esxi/12.png)

**加载完成，F2进入配置页**
![img](../../img/esxi/13.png)

**输入管理员密码，回车进入**
![img](../../img/esxi/14.png)

**完成网络配置后，浏览器输入地址进入ESXI控制台**
![img](../../img/esxi/15.png)


### 参考文献
[VM官方下载地址](https://customerconnect.vmware.com/cn/downloads/info/slug/datacenter_cloud_infrastructure/vmware_vsphere/7_0)
[VM官方查询地址](https://www.vmware.com/resources/compatibility/search.php)
[rufus](http://rufus.ie/en/)
