# 

## Controllers 
控制层 负责接受参数，验证数据，调度业务层（Servers）/ 数据访问层（Repositories）

## Models
数据模型层 定义数据模型 及模型关系
### relations
复用模型关系
### scopes
全局查询作用域


## Repositories
数据仓库层 不直接依赖数据库，负责获取数据（多种数据源），处理缓存等


## Services 
业务服务层 负责业务逻辑组装