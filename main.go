package main

import (
	"learnGin/app/check"
	"learnGin/app/html"
	"learnGin/app/login"
	"learnGin/app/parMidWar"
	"learnGin/app/parameter"
	"learnGin/app/token"
	"learnGin/routers"
	"log"
)

func main() {
	log.Println("load routers")
	routers.Include(parameter.RoutersBasic, html.RoutersHtml,
		parMidWar.RoutersPart, login.RoutersLogin,
		check.RoutersCheck, token.RoutersToken)

	log.Println("init routers")
	rou, err := routers.Init()
	if err != nil {
		panic(err)
	}

	log.Println("4. listen port default 8080")
	if err := rou.Run(":8080"); err != nil {
		log.Printf("start service failed %v\n", err)
	}
}
