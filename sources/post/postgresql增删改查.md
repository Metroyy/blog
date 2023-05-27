---
title: postgresql增删改查
desc: postgresql增删改查
time: 2023-03-16
tags: postgresql
---

### 增

**语法**

```sql
INSERT INTO 表名 (列1, 列2, 列3,...列N)
VALUES ('值1', '值2', '值3',...'值N');
```

**例**

```sql
INSERT INTO "mima" ("desc", "user_name", "static_postion", "make_postion") 
VALUES ('RUhhdV6JkQ', 'Laura Vargas', 'K37kDJXSsZ', 'H1FwzDeyAK');
```

### 删

**语法**

```sql
DELETE FROM 表名 
WHERE 条件;
```

**例**

```sql
DELETE FROM mima WHERE id = 1;
```

### 改

**语法**

```sql
UPDATE 表名 SET 列1 = '值1', 列2 = '值2'...., 列N = '值N' 
WHERE 条件;
```

**例**

```sql
UPDATE mima
SET "desc" = '银行卡密码',user_name = 'YY',static_postion= '123yy',make_postion = '中国银行'
WHERE id = 1;
```

### 查

**语法**

```sql
SELECT 列1, 列2,...列N FROM 表名
WHERE 条件;
```

**例**

```sql
#查询所有
SELECT * FROM 表名;
SELECT * FROM mima;

#查询某列
SELECT 列名 FROM 表名;
SELECT desc FROM mima;

#查询多列
SELECT 列1,列2,列3 FROM 表名;
SELECT id,user_name FROM mima;

```
