# API

IP: 127.0.0.1

port: 31415

## /api/v1/user/login

### POST 都是明文

```json
{
  "email": "",
  "password": ""
}
```

#### email与数据库不匹配

```json
{
  "code": 422,
  "msg": "用户不存在"
}
```

#### password与数据库不匹配

```json
{
  "code": 422,
  "msg": "密码错误"
}
```

#### 登录成功

```json
{
  "nickname": "",
  "email": "",
  "msg": "用户登录成功"
}
```

### 其他

可以使用get先获得需要的json格式

```json
{
  "email": "",
  "password": "",
  "msg": "请使用此格式的POST来登录"
}
```

## /api/v1/user/register

### POST 都是明文

```json
{
  "email": "",
  "password": "",
  "nickname": ""
}
```

#### email与数据库匹配

```json
{
  "code": 422,
  "msg": "用户已存在"
}
```

#### 注册成功

```json
{
  "nickname": "",
  "email": "",
  "msg": "用户注册成功"
}
```

### 其他

可以使用get先获得需要的json格式

```json
{
  "email": "",
  "password": "",
  "msg": "请使用此格式的POST来登录"
}
```

## /api/v1/img/

+ gas/
+ greedy/
+ dijkstra/

用户输入
```txt
"1,2 2,4"
```

post一个query

```txt
http://127.0.0.1:31415/api/v1/img/gas?cp=1,2&cp=2,4
```

得到
```json
{
  "coordinate_points": [
    "1,2",
    "2,4"
  ],
  "img": "gas.png",
  "path": "http://127.0.0.1:31415/img/gas.png"
}
```
img标签展示path的地址