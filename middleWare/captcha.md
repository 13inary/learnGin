```
package main

import (
"bytes"
"github.com/dchest/captcha"
"github.com/gin-contrib/sessions"
"github.com/gin-contrib/sessions/cookie"
"github.com/gin-gonic/gin"
"net/http"

"time"
)
/ 中 间 件 , 处 理 session
func Session(keyPairs string) gin.HandlerFunc {
store := SessionConfig()

return sessions.Sessions(keyPairs, store)
}

func SessionConfig() sessions.Store {
sessionMaxAge := 3600
sessionSecret := "topgoer"

var store sessions.Store
store = cookie.NewStore([]byte(sessionSecret))
store.Options(sessions.Options{

MaxAge: sessionMaxAge, //seconds

Path: "/",
return store
}

func Captcha(c *gin.Context, length ...int) {

1 := captcha.DefaultLen

w,h := 107, 36

if len(length) == 1 {
1 = length[0]

}

if len(length) == 2 {
w = length[1]

}

if len(length) == 3 {
h = length[2]
}

captchaId := captcha.NewLen(1)
session := sessions.Default(c)

session.Set("captcha", captchaId)

_= session.Save()

_=Serve(c.Writer, c.Request, captchaId, ".png", "zh", false, w, h)
}

func CaptchaVerify(c *gin.Context, code string) bool {
session := sessions.Default(c)
if captchaId := session.Get("captcha"); captchald != nil {
session.Delete("captcha")
_ = session.Save()
if captcha.VerifyString(captchald.(string), code) {
return true
1 else {
return false
}
1 else {

return false
}

func Serve(w http.ResponseWriter, r *http.Request, id, ext, lang string,
download bool, width, height int) error {
w.Header () .Set("Cache-Control", "no-cache, no-store, must-revalidate")
w.Header () .Set("Pragma", "no-cache")

w.Header () .Set("Expires", "O")
var content bytes .Buffer
switch ext {
case ".png":
w.Header ().Set("Content-Type", "image/png")
_ = captcha.wWriteImage(&content, id, width, height)

case ".wav'":
w.Header ().Set("Content-Type", "audio/x-wav")

default:

= captcha.WriteAudio(&content, id, lang)

return captcha.ErrNotFound
}
if download {
w.Header ().Set("Content-Type", "application/octet-stream")
}
http.ServeContent(w, r, id+ext, time.Time{},
bytes.NewReader (content.Bytes()))

return nil
}


func main() {

router := gin.Default()

router.LoadHTMLGlob("./*.html")

router.Use(Session("topgoer"))

router .GET("/captcha", func(c *gin.Context) {
Captcha(c, 4)

})

router.GET("/", func(c *gin.Context) {
C.HTML(http.StatuSOK, “index.html“ , nil )

})

router .GET("/captcha/verify/:value", func(c *gin.Context) {

value := c.Param("value")
if captchaverify(c, value) {
C.JSON(http.StatusOK, gin.H{"status": 0,
1 else {
Cc.JSON(http.StatusOK, gin.H{"status": 1,

"mSg" :

"mSg" :

"success"})

"failed"})
router .Run(":8080")
}
```

