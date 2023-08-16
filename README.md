# IM
用golang实现的im聊天系统,用了gin,gorm,websocket,mysql,rabbitmq,redis

### 结构
routes->api->service->repository

### 实现在线聊天的大概思路
用户通过在前端页面点击chat聊天,就会与服务器建立一个长连接,并在后端创建一个client实例,在client_manager注册并被管理.假设用户A发送一条消息给用户B,就由后端接收到,并由client_manager确认用户B是否在线,在线则直接发给用户B,否则则交给发送到mq,交给后台程序处理.
