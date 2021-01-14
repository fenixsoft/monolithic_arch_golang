package main

// 通过rice.go在编译期将静态资源进行嵌入，保证单个执行文件可直接运行
//go:generate go run github.com/GeertJohan/go.rice/rice embed-go
import (
	"flag"
	rice "github.com/GeertJohan/go.rice"
	"github.com/fenixsoft/monolithic_arch_golang/controller"
	"github.com/fenixsoft/monolithic_arch_golang/infrasturcture/config"
	"github.com/fenixsoft/monolithic_arch_golang/infrasturcture/db"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io/ioutil"
)

var conf = flag.String("c", "", "Fenix's Bookstore的配置文件，格式同Java版本的Spring Application YAML配置")

func main() {
	// 处理启动参数传入的配置
	var data []byte
	var err error
	var logger = logrus.StandardLogger()

	flag.Parse()
	if *conf != "" {
		data, err = ioutil.ReadFile(*conf)
	} else {
		data, err = rice.MustFindBox("resource").Bytes("application.yml")
	}
	if err != nil {
		panic("读取配置文件失败：" + err.Error())
	} else if err = config.LoadConfiguration(data); err != nil {
		panic("解析配置文件失败：" + err.Error())
	}

	// 初始化数据库
	ddl, _ := rice.MustFindBox("resource/db/" + config.GetConfiguration().Database).String("schema.sql")
	dml, _ := rice.MustFindBox("resource/db/" + config.GetConfiguration().Database).String("data.sql")
	db.InitDB(ddl, dml)
	logger.Info("初始化数据库完毕")

	// 初始化路由与HTTP服务
	gin.DefaultWriter = logger.WriterLevel(logrus.DebugLevel)
	gin.DefaultErrorWriter = logger.WriterLevel(logrus.ErrorLevel)
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	// router.Use(gin.Logger())  // 这个日志中间件太话唠了，建议不加载
	router.Use(gin.Recovery())
	controller.Register(router)
	_ = router.Run(":" + config.GetConfiguration().Port)
}
