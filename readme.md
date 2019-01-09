# 不学网后端程序

[不学网前端项目地址](https://github.com/noxue/noxue-ui)

### 需求列表

* [需求思维导图在线预览](https://www.processon.com/view/link/5c227a30e4b0df5abaa60247)

### 开发规范

1. 所有文件小写字母组成，或下划线分隔
1. 方法名 采用pascal命名法，格式 类名+操作，例如：UserInsert，ArticleSelect
1. 局部变量采用骆驼命名法，全局变量采用pascal命名法


### 返回状态码

*错误信息格式
```json
{
  "status":1000,
  "msg": "授权已过期"
}
```

| http状态码 | 错误码 | 具体信息 |
| ------ | ------ | ------ |
| 200 | 0 | 无错误 |
| 401 | 1000 | 授权已过期 |
| 401 | 1001 | 授权超过刷新时间，请重新登陆 |

* 10开头表示用户相关的错误

### 开发进度

**2018/12/27**

* 完善mgodb库功能，批量查询时可以显示指定的字段，直接指定排序规则
* 完成用户组dao层的编写和测试

**2018/12/28**

* 放弃mgodb，因为功能太不完善，只提供了非常基础的功能，所以自行按照需求封装ormgo库操作mongodb
* 用ormgo重写user相关的所有dao

**2018/12/30**

* 完成用户相关所有service层功能编写

**2018/12/31**

* 添加图片验证码
* 添加邮箱和手机验证码功能（未实现具体发送功能）
* 注册功能（包括验证手机或邮箱收到的验证码）
