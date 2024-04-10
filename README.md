# Meeting

> 基于 golang、webrtc 实现在线会议

# 核心模块
- webrtc 实时通讯技术
```shell
go get -u github.com/pion/webrtc/v4
```

## 系统模块

+ [x] 会议管理
    + [x] 会议列表
    + [x] 创建会议
    + [x] 会议编辑
    + [x] 会议删除
+ [x] 用户管理
    + [x] 登录
+ [ ] WebRTC
    + [x] data channels
    + [x] 屏幕共享
    + [ ] 一对一音视频通信