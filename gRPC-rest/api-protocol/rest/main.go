package main

import (
	"fmt"
	cont "grpc-rest/api-protocol/rest/movies"
	handler "grpc-rest/api-protocol/rest/routes"
	"grpc-rest/pkg/repository"
	rep "grpc-rest/pkg/repository"
	serv "grpc-rest/pkg/service"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectionMysql() *gorm.DB {

	configDB := map[string]string{

		"DB_Username": "root",
		"DB_Password": "",
		"DB_Port":     "3306",
		"DB_Host":     "localhost",
		"DB_Name":     "catalog",
	}

	fmt.Println(configDB)

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		configDB["DB_Username"],
		configDB["DB_Password"],
		configDB["DB_Host"],
		configDB["DB_Port"],
		configDB["DB_Name"])

	db, e := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if e != nil {
		panic(e)
	}
	repository.InitMigrate(db)
	return db
}

func main() {
	Conn := ConnectionMysql()
	Repository := rep.NewLogRepository(Conn)
	Service := serv.NewService(Repository)
	Handler := cont.NewController(Service)
	e := echo.New()
	handler.HandlerApi(e, Handler)

	e.Logger.Fatal(e.Start(":8000"))
}
