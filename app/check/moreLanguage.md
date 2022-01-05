```
package main

import (
llfmt T

"github.com/gin-gonic/gin"

"github.com/go-playground/locales/en"
"github.com/go-playground/locales/zh"
"github.com/go-playground/locales/zh_Hant_TW"

ut "github.com/go-playground/universal-translator"
"gopkg.in/go-playground/validator.v9"

en_translations "gopkg.in/go-playground/validator.v9/translations/en"
zh_translations "gopkg.in/go-playground/validator.v9/translations/zh"

zh_tw_translations "gopkg.in/go-playground/validator.v9/translations/zh_tw"

var (
Uni *ut.UniversalTranslator

Validate *validator.Validate
)

type User struct {
Username string “form:"user_name" validate:'"required"®
Tagline string “form:"tag_line" validate:"required,lt=10""

Tagline2 string "form:"tag_line2" validate:'"required,gt=1""
}

func main() {
en := en.New()
zh := zh.New()
zh_tw := zh_Hant_TW.New()
Uni = ut.New(en, zh, zh_tw)

Validate = validator.New()

route := gin.Default()
route.POST("/51mh", startPage)
route.Run(":8080")
}

func startPage(c *gin.Context) {

// 这 部 分 应 放 到 中 间 件 中

locale := c.DefaultQuery("locale", "zh")

trans, _ := Uni.GetTranslator(locale)

switch locale {

case "zh":
zh_translations.RegisterDefaultTranslations(Validate, trans)
break

case "en":
en_translations.RegisterDefaultTranslations(Validate, trans)
break

case "zh_tw":
zh_tw_translations.RegisterDefaultTranslations(Validate, trans)

break
default:
zh_translations.RegisterDefaultTranslations(Validate, trans)

break
}

// 自 定 义 错 误 内 容
Validate.RegisterTranslation("required", trans, func(ut ut.Translator)
error {
return ut.Add("required", "{0} must have a value!", true) // see
universal-translator for details
}, func(ut ut.Translator, fe validator.FieldError) string {
t, _ = ut.T("required", fe.Field())
return t

})

// 这 块 应 该 放 到 公 共 验 证 方 法 中
user := User{}
c.ShouldBind(&user)

fmt.Println(user)

err := Validate.Struct(user)

if err != nil {
errs := err.(validator.ValidationErrors)
sliceErrs := []string{}
for _, e := range errs {

sliceErrs = append(sliceErrs, e.Translate(trans))
}
c.String (200, fmt.Sprintf("%#v", sliceErrs))

}
c.String(200, fmt.Sprintf("%#v", "user"))
}
```




```
正 确 的 链 接 : http:/7localhost :898@/testing?user_name= 枯 藤
&tag_line=9&tag_line2=33&locale=zh

http://localhost:8080/testing?user_name=fhf&
&tag_line=9&tag_line2=3&locale=en 返 回 英 文 的 验 证 信 息

http://Vlocalhost:898@/testing2?user_name= 枯 藤
&tag_1ine=9&tag_1ine2=3&locale=zh 返 回 中 文 的 验 证 信 息

查 看 更 多 的 功 能 可 以 查 看 官 网 gopkg.in/go-playground/validator .v9
```

