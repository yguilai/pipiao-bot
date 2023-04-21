# 皮皮奥频道机器人

### 目录结构
```text
/api          相关proto文件及生成代码
/app
  /channelbot 频道机器人, 指令解析&消息回复
  /transport  消息转发服务, 对接qq官方websocket, 将消息转发至消息队列
  /wft        翻译服务(grpc), 拉取wf i18n词条, 及wm词条
  /wfs        wf服务(grpc), 防腐、解析. 
              对接wf官方服务、wm服务, 解析wf世界状态信息, 调用wm接口
/deploy       开发依赖环境, 使用docker compose
/pkg          封装的一些工具包
/third_party  第三方proto文件
```

### 代码生成

- 配置代码生成
```bash
make config
```

- wire依赖注入
```bash
make generate
```

- proto代码生成
```bash
make api
```

### Run
```bash
# 运行依赖环境
cd deploy/env
docker-compose up -d

# 启动消息转发服务, 其他服务类似操作 (需要提前在etcdkeeper配置中心加上相关配置)
cd app/transport
make dev
```