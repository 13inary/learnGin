```
Sessions
gorillLa/sessions为自定义session后端提供cookie和文件系统session以及基础结构。
主要功能是:
简单的API:将其用作设置签名(以及可选的加密)cookie的简便方法。
内置的后端可将session存储在cookie或文件系统中。
F1Lash消息:一直持续读取的session值。
切换session持久性(又称“记住我“)和设置其他属性的便捷方法。
旋转身份验证和加密蝈钥的机制。
每个请求有多个session,即使使用不同的后端也是如此。
自定义session后端的接口和基础结构;可以使用通用API检索并批量保存来自不同商店的
session,
```



```
package main

import (
Ilfmtll
"net/http"

"github.com/gorilla/sessions"

// 初 始 化 一 个 cookie 存 储 对 象

/ something-very -Secret 应 该 是 一 个 你 自 己 的 密 匙 , 只 要 不 被 别 人 知 道 就 行

var store = sessions.NewCookieStore([]byte("something-very-secret"))

func main() {
http.HandleFunc("/save", SaveSession)
http.HandleFunc("/get", GetSession)
err := http.ListenAndServe(":8080", nil)
if err != nil {
fmt.Println("HTTP server failed,err:", err)

return

func SaveSession(w http.ResponseWriter, r *http.Request) {

// Get a session. We're ignoring the error resulted from decoding an

// existing session: Get() always returns a session, even if empty.

/ “ 获 取 一 个 session 对 象 ,session-name 是 session 的 名 字
session, err := store.Get(r, "session-name")
if err != nil {
http.Error(w, err.Error(), http.StatusInternalServerError)

return
}
}

// 在 session 中 存 储 值
session.Values["foo"] = "bar"
session.Values[42] = 43

// 保 存 更 改

session.Save(r, w)

func GetSession(w http.ResponseWriter, r *http.Request) {

session, err := store.Get(r, "session-name")

if err != nil {
http.Error(w, err.Error(), http.StatusInternalServerError)
return

]

foo := session.Values["foo"]

fmt .Println(foo )

}

/ 删 除
// 将 session 的 最 大 存 储 时 间 设 置 为 小 于 零 的 数 即 为 删 除
session.Options.MaxAge = -1

session.Save(r, w)
```

