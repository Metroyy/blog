---
title: Hexo博客建站
desc: Hexo博客建站
time: 2023-02-10
tags: Hexo
---

**安装Hexo**
```bash 
npm install -g hexo-cli ---全局安装
npm install hexo ---局部安装
```

> **安装以后，可以使用以下两种方式执行 Hexo：**
> ```bash
> npx hexo <command>
> ```
> **Hexo所在的目录下的`node_modules`添加到环境变量中即可`hexo <command>`：**
> ```bash
>   echo 'PATH="$PATH:./node_modules/.bin"' >> ~/.profile
> ```

**初始化Hexo**
```bash
hexo init <folder>
cd <folder>
npm install
```

**完成初始化后，目录如下：**
```yaml
hexo/
  |- node_modules/  # hexo需要的模块，不需要上传GitHub
  |- themes/        	# 主题文件，需要上传GitHub的dev分支
  |- sources/       	# 博文md文件，需要上传GitHub的dev分支
  |- public/        	# 生成的静态页面，由hexo deploy自动上传到gh-page分支
  |- package.json  	# 记录hexo需要的包信息，不需要上传GitHub
  |- _config.yml    	# 全局配置文件，需要上传GitHub的dev分支
  |- .gitignore       	# hexo生成默认的.gitignore，它已经配置好了不需要上传的hexo文件
```

### 安装主题MengD
```bash
git clone https://github.com/lete114/hexo-theme-MengD.git themes/MengD ---稳定版
npm install hexo-theme-mengd --save ---npm安装主题稳定版
```

**编辑hexo配置文件_config.yml**
```yaml
theme: MengD
```
> **可选插件**
> ```bash
> npm install hexo-generator-search --save ---本地搜索
> ```
> ```bash
> npm install hexo-minify --save ---资源压缩(提高博客速度)
> ```

### 启动Hexo
**生成静态页面**
```bash
hexo g
```

**启动hexo-server**
```bash
npm install hexo-server --save
hexo server
```

### 参考文献
[Hexo官方文档](https://hexo.io/zh-cn/docs/)
[使用Hexo+服务器搭建个人博客](https://www.jianshu.com/p/6ae883f9291c)
[hexo初探---让写作飞起来](https://www.jianshu.com/p/e7c58f57f60e)
[MengD](https://mengd.js.org/)
