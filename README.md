# ginEssential

### 开发环境

go版本：1.17.6

mysql版本: 8.0.29-0

### 使用须知

需要在config文件夹中配置一个application.yml文件

文件内容格式如下：

````
server:
  port: 1016               //接收、输出端口
  datasource:              //数据源
  driverName: mysql        //数据库
  host: 127.0.0.1          //数据库地址
  port: 3306               //数据库端口
  database: ginessential   //数据库名
  username: ***           //数据库用户
  password: ***           //数据库密码
  charset: utf8            //字符集
  loc: Asia/Shanghai       //时间
````

并在数据库中创建以database的值为名的数据库，之后便可正常使用。





