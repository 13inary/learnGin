package routers

import (
	"io"
	"log"
	"os"
	"testGin/middleWare"
	"testGin/validators"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Option func(*gin.Engine)

var options = []Option{}

func Include(opts ...Option) {
	options = append(options, opts...)
}

func nameNotNullAndAdmin(f1 validator.FieldLevel) bool {

	if v, f := f1.Field().Interface().(string); f {
		return v != "" && !("admin" == v)
	}

	return true
}

func Init() (e *gin.Engine, err error) {
	// cancel debug mode
	//gin.SetMode(gin.ReleaseMode)

	// log
	//gin.DisableConsoleColor()
	f, err := os.Create("gin.log")
	if err != nil {
		log.Println("create log error")
		return nil, err
	}
	//gin.DefaultWriter = io.MultiWriter(f)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	log.Println("1. new router")
	// default use middleware Logger() and Recovery()
	e = gin.Default()
	//e= gin.New()

	// custom checker
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		log.Println("binding validator error")
		//return nil, err
	}
	// register validate
	v.RegisterValidation("NotNullAndAdmin", validators.NameNotNullAndAdmin)
	v.RegisterValidation("bookabledate", validators.BookableDate)

	// config
	// upload
	e.MaxMultipartMemory = 8 << 20
	// html
	e.LoadHTMLGlob("tem/**/*")
	// middleware
	e.Use(middleWare.Before())
	// static
	e.Static("/assets", "./assets")

	// execute routers
	for _, opt := range options {
		opt(e)
	}

	return e, nil
}
