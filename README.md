# CloudDisk

> 轻量级云盘系统，基于go-zero、xorm实现。
> 
> B站视频链接：https://www.bilibili.com/video/BV1cr4y1s7H4/

使用到的命令
```text
# 创建API服务
goctl api new core
# 启动服务
go run core.go -f etc/core-api.yaml
# 使用api文件生成代码
goctl api go -api core.api -dir . -style go_zero
```

腾讯云COS后台地址：https://console.cloud.tencent.com/cos/bucket

腾讯云COS帮助文档：https://cloud.tencent.com/document/product/436/31215

## 系统模块
- [x] 用户模块
  - [x] 密码登录
  - [x] 刷新Authorization
  - [x] 邮箱注册
  - [x] 用户详情
  - [x] 用户容量
- [x] 存储池模块
  - [x] 中心存储池资源管理
    - [x] 文件上传
    - [x] 文件秒传
    - [x] 文件分片上传
    - [x] 对接 MinIO
    - [ ] 对接阿里对象存储
  - [x] 个人存储池资源管理
    - [x] 文件关联存储
    - [x] 文件列表
    - [x] 文件名称修改
    - [x] 文件夹创建
    - [x] 文件删除
    - [x] 文件移动
- [x] 文件分享模块
  - [x] 创建分享记录
  - [x] 获取资源详情
  - [x] 资源保存

