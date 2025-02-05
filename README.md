# 抽奖秒杀系统

## 项目简介

这是一个基于前后端分离的抽奖秒杀系统。前端使用 Vue.js 实现，后端使用 Go 语言实现，并使用 Redis 作为缓存数据库。系统包含抽奖转盘、奖品管理、并发压力测试等功能。

## 目录结构

```
.
├── lottery_backend       # 后端代码
│   ├── controller        # 控制器
│   ├── database          # 数据库相关
│   ├── model             # 数据模型
│   ├── router            # 路由
│   ├── stress_test       # 并发压力测试
│   ├── test              # 测试代码
│   └── utils             # 工具函数
├── lottery_frontend      # 前端代码
│   ├── public            # 公共资源
│   ├── src               # 源代码
│   │   ├── assets        # 静态资源
│   │   ├── components    # 组件
│   │   ├── router        # 路由配置
│   │   ├── views         # 视图
│   │   ├── App.vue       # 根组件
│   │   ├── main.js       # 入口文件
│   │   └── axiosAPI      # API 请求
└── README.md             # 项目说明文件
```

## 前端部分

前端使用 Vue.js 实现，主要功能包括：
- 抽奖转盘展示
- 奖品信息展示
- 抽奖结果展示

### 主要文件

- `src/views/HomeView.vue`：抽奖转盘的主要视图组件。
- `src/router/index.js`：前端路由配置。

## 后端部分

后端使用 Go 语言实现，主要功能包括：
- 奖品信息管理
- 抽奖逻辑实现
- 并发压力测试

### 主要文件

- `controller`：包含奖品信息获取和抽奖逻辑的控制器。
- `model`：包含数据库模型和初始化逻辑。
- `router`：包含后端路由配置。
- `stress_test`：包含并发压力测试代码。
- `utils`：包含工具函数，如二分查找函数。

## 数据库

使用 MySQL 作为数据库，包含一个 `inventory` 表，用于存储奖品信息。

### 数据库初始化脚本

- `database/init.sql`：创建 `inventory` 表并插入初始数据。

## 运行步骤

### 前端部分

1. 安装依赖：

    ```bash
    cd lottery_frontend
    npm install
    ```

2. 启动前端服务：

    ```bash
    npm run serve
    ```

### 后端部分

1. 安装依赖：

    ```bash
    go mod tidy
    ```

2. 初始化数据库：

    ```bash
    # 确保 MySQL 服务已启动，并执行以下命令导入初始化脚本
    mysql -u root -p < lottery_backend/database/init.sql
    ```

3. 启动后端服务：

    ```bash
    go run lottery_backend/main.go
    ```

### 并发压力测试

1. 运行压力测试：

    ```bash
    go test -v lottery_backend/test/gift_test.go
    ```

## 注意事项

- 确保 Redis 服务已启动，并配置正确的连接信息。
- 确保 MySQL 服务已启动，并配置正确的连接信息。

## 结语

希望这个项目能帮助你更好地理解前后端分离的抽奖秒杀系统的实现。如果有任何问题或建议，欢迎提 Issue 或 Pull Request。
