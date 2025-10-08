package main

import (
	"fmt"
	"html/template"

	"pankreatitmed/internal/app/config"
	"pankreatitmed/internal/app/dsn"
	"pankreatitmed/internal/app/handler"
	"pankreatitmed/internal/app/repository"
	"pankreatitmed/internal/app/services"
	"pankreatitmed/internal/pkg"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	router := gin.Default()
	conf, err := config.NewConfig()
	if err != nil {
		logrus.Fatalf("error loading config: %v", err)
	}
	router.SetFuncMap(template.FuncMap{
		// true, если указатель не nil и значение != 0
		"nzf": func(p *float64) bool { return p != nil && *p != 0 },

		// безопасно достаём значение (0, если nil)
		"valf": func(p *float64) float64 {
			if p == nil {
				return 0
			}
			return *p
		},
	})
	router.LoadHTMLGlob("templates/*")

	postgresString := dsn.FromEnv()
	fmt.Println(postgresString)

	rep, errRep := repository.New(postgresString)
	if errRep != nil {
		logrus.Fatalf("error initializing repository: %v", errRep)
	}

	svcs := services.NewServices(services.Reps{
		CriteriaRepo:      rep,
		MedOrdersRepo:     rep,
		MedOrderItemsRepo: rep,
		MedUsersRepo:      rep,
	})

	hand := handler.NewHandler(svcs)

	application := pkg.NewApp(conf, router, hand)
	application.RunApp()
}
