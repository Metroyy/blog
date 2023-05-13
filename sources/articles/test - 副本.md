---
title: "这是第二篇文章标题"
img: "https://cdn.gamma.app/cdn-cgi/image/quality=80,fit=scale-down,onerror=redirect,width=500/a6uyzivr086smdy/bcde7aa37d204865b27a758dcb271fa6/original/u-4010697962-4143975424-fm-253-app-120-f-JPEG-fmt-auto-q-75.jpg"
desc: "这是第二篇文章描述"
time: "2023-5-13"
tags: "java,python"
---

### 查询信息命令
```yaml
show this ---查看系统当前模式下生效的配置信息
show ver ---查看设备信息
show vlan ---查看vlan信息
show line ---查看线路的配置信息
show int sta ---查看接口状态信息
show int switchport ---查看接口模式信息
show int link sta ---查看接口震荡情况
show service ---查看服务的开关状态
show sessions ---查看telnet实例信息
show ftp-server ---查看ftp信息
show tcp connect ----查看系统当前 IPv4 TCP 连接的基本信息
show tcp connect statistics ---查看系统当前 IPv4 TCP 连接的统计信息
show ip interface br ---查看ip信息
show arp ---查看mac地址
show ip route ---查看路由表
show arp 0.0.0.0 ---查看ip地址从哪个接口学习的
show ip dhcp pool ---查看地址池分配情况
show ip dhcp binding ---查看dhcp分配列表
show ip dhcp server statistics ---查看dhcp 地址池分配统计
show ip dhcp conflict ---查看dhcp冲突列表
tracert -d 0.0.0.0 ---路由追踪
```

### 接口命令
```yaml
errdisable recovery ---接口状态disabled时重新开启
int TenGigabitEthernet 0/x ---打开/关闭接口
no shutdown/shutdown

interface Ten/FortyGigabitEthernet 0/1x ---将接口添加到vlan x
switchport
switchport mode access
switchport access vlan x

vlan x ---批量添加接口到vlan x
add int range Ten/FortyGigabitEthernet 0/1-x

interface Ten/FortyGigabitEthernet 0/x ---修改接口速率
speed 100/1000/10G
half-duplex ---半双工
full-duplex ---全双工
```
awdwad
![img](../sources/img/title.png)
awwwwwwwwwwwwwwwwwd
### VLAN命令
```yaml
int vlan x ---创建vlanx
name xxx ---vlan更名
vlan range a,b,c ---批量创建vlan a,b,c
vlan range 1-10 ---批量创建vlan1-10

vlan x ---配置vlan接口的ip地址
ip address 10.0.2.254 24

interface vlan x ---删除vlan接口的ip地址
no ip address
```
![img](../sources/img/vgpu/1.png)
![img](../sources/img/vgpu/3.png)
![img](../sources/img/vgpu/4.png)