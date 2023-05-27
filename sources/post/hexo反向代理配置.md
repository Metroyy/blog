---
title: Hexo反向代理配置
desc: Hexo反向代理配置
time: 2023-02-11
tags: Hexo,反向代理
---

<font color=gray>**Nginx配置**</font>
```yaml
    location / {
        proxy_pass http://127.0.0.1:端口;
    }

    location ~ .*\.(gif|jpg|jpeg|png|bmp|swf|flv|mp3|wma|js|css)${
        proxy_pass http://127.0.0.1:端口;
        expires      30d;
    }
	
    location ~ .*\.(js|css)${
        proxy_pass http://127.0.0.1:端口;
        expires      12h;
    }
```