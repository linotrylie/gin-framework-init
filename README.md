gin-framework-init管理后台

## 模块说明

## configs: 配置文件

## core : 核心模块

- api: 控制器层
    - admin : 管理员后台
    - mch : 商户后台
    - openapi : 开放平台

- consts:定义静态变量层
- dao: 数据库操作层
- middleware: 中间件层
- model: 数据表层
- router: 路由层

- service: 服务层
    - admin : 管理员后台
    - common : 公共服务

- codegenerator：代码生成器

### docs: 文档集合

### global: 全局变量

- db
- redis
- setting
- tracer

### pkg : 项目相关的模块包

### storage: 项目生成的临时文件

### scripts : 各类构建、安装、分析等操作的脚本

## 安全

- 提交代码前,必须本地运行成功
- 进行单元测试: ``` go test -v ./... ```

## 开发规约

### 数据库

- 更新字段有数据结构对应的0值时,不要使用结构体,而是map
- 更新数据,如果能预期更新条数的,应明确判断更新条数的范围,eg: rowsAffected == 1,rowsAffected > 1 等

- 更新数据字段,统一使用 id
- 新模块定义,数据表业务类型不使用0代表业务类型,避免在使用 ozzo-validate 时,传了0值但是报数据不能为空的问题
- 数据校验,整型数据,必须定义校验最大值和最小值,避免超出范围导致数据精度丢失和业务的bug
- 数据库model下的代码,不要手动更新,自动生成代码时,会覆盖
- 数据库字段注释需要使用逗号的,使用英文逗号,避免自动生成代码时,不能通过代码检查
- 枚举注释写法:类型:1=目录,2=菜单,3=按钮,方便使用工具生成代码

## 数据对比

- 整型数据对比时,需要将要比较的两个数据转为相同的类型,如 i int32 = 1,j int = 1,但是 i != j

### 中文标点符号

- 禁止使用中文逗号和中文问号,因为做了代码检测,避免Gorm中误用导致sql错误

### 容器化部署设置GOMAXPROCS

- 引入Uber公司推出的 uber-go/automaxprocs

## 运维

- 测试环境或开发环境
    - 配置文件：configs/config.yaml
    - runMode: debug

- 生产环境
    - 配置文件：configs/config.yaml
    - runMode: release,(生产环境一定要配置为:release,不然将会获取到测试环境的配置信息)

- 所需软件：
    - mysql5.7
    - go1.18
    - redis

- 健康检查：
    - http://ip:port/v1/system/health
    - 返回数据:{"code":0,"msg":"ok!"}

  
  