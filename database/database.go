package database

import (
    "fmt"
    "sync"
    "github.com/beego/beego/v2/client/orm"
    "github.com/beego/beego/v2/server/web"
    _ "github.com/go-sql-driver/mysql"
    "rbac-beego-api/models"
)

var (
    instance *Database
    once sync.Once
)

func InitDatabase() {
    once.Do(func() {
        // Get database configuration
        dbUser := web.AppConfig.DefaultString("mysqluser", "root")
        dbPass := web.AppConfig.DefaultString("mysqlpass", "")
        dbName := web.AppConfig.DefaultString("mysqldb", "bee_rbac")
        dbHost := web.AppConfig.DefaultString("mysqlhost", "127.0.0.1")
        dbPort := web.AppConfig.DefaultString("mysqlport", "3306")

        dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&collation=%s&%s", 
            dbUser, dbPass, dbHost, dbPort, dbName, 
            web.AppConfig.DefaultString("mysqlcharset", "utf8mb4"),
            web.AppConfig.DefaultString("mysqlcollation", "utf8mb4_unicode_ci"),
            web.AppConfig.DefaultString("mysqlparams", "parseTime=true&loc=Local"))

        // Register models
        orm.RegisterModel(
            new(models.User),
        )

        // Setup database connection
        orm.RegisterDriver("mysql", orm.DRMySQL)
        orm.RegisterDataBase("default", "mysql", dataSource)
    })
}


type Database struct {
    Ormer orm.Ormer
}

func GetInstance() *Database {
    once.Do(func() {
        instance = &Database{
            Ormer: orm.NewOrm(),
        }
    })
    return instance
}

func GetOrmer() orm.Ormer {
    return orm.NewOrm()
}