# Open IM Client
OpenIM REST API SDK。

Open IM Client 是一个封装了访问 Open IM Server REST API 的开源GO客户端库。


## 特性

- 封装 Open IM Server REST API
- 易于集成和使用的客户端接口
- 遵循 Open IM Server 相同的开源协议

## 安装

```bash
# 使用 go get 安装
go get -u github.com/HessianZ/open-im-client
```

## 使用示例

```go
import (
    "log"
    "github.com/HessianZ/open-im-client/openim"
	"github.com/HessianZ/open-im-client/openim/apistruct"
)

func main() {
    apiBaseUrl := "https://open-im.example.com"
    adminID := "imAdmin"
    adminSecret := "openim123"
    debug := true
    client := openim.NewOpenIMClient(apiBaseUrl, adminID, adminSecret, debug)

    req := apistruct.SendMsg{
        SendID:           "100001",
        RecvID:           "100002",
        SenderPlatformID: openim.PlatformID_Admin,
        ContentType:      openim.MessageType_Text,
        SessionType:      openim.ConversationType_Single,
        Content:          map[string]any{
            "content": "Hello World!",
        },
        // Ex:               `{"extra": "hohoho~"}`,
    }
    resp, err := client.SendMessage(req)
    if err != nil {
        log.Fatal(err)
        return
    }
    log.Printf("%+v", resp)
}

```

## 完成状态
| 模块 | 状态 |
| ----------- | ----------- |
| 认证管理 | ✅已完成 |
| 用户管理 | 🚧进行中 |
| 关系链管理 | ✅已完成 |
| 群组管理 | 🚧进行中 |
| 会话管理 | ✅已完成 |
| 消息管理 | 🚧进行中 |

## 开源协议

本项目采用 Apache 2.0 开源协议，与 Open IM Server 保持一致。

## 贡献

欢迎提交 Issue 和 Pull Request 来改进这个项目。
