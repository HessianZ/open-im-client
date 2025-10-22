# Open IM Client

Open IM Client 是一个封装了访问 Open IM Server HTTP API 的开源客户端库。

## 特性

- 封装完整的 Open IM Server HTTP API
- 易于集成和使用的客户端接口
- 遵循 Open IM Server 相同的开源协议

## 安装

```bash
# 使用 go get 安装
go get -u github.com/HessianZ/open-im-client
```

## 使用示例

```go
// 在这里添加使用示例
import "github.com/HessianZ/open-im-client/openim"

func main() {
    apiBaseUrl := "https://open-im.example.com"
    adminID := "imAdmin"
    adminSecret := "openim123"
    debug := true
    client := openim.NewOpenIMClient(apiBaseUrl, adminID, adminSecret, debug)

    resp := client.SendMessage(...)
}

```

## 开源协议

本项目采用 Apache 2.0 开源协议，与 Open IM Server 保持一致。

## 贡献

欢迎提交 Issue 和 Pull Request 来改进这个项目。
