# Fenix's Bookstore后端：Go语言实现

这是Fenix's Bookstore单体版本的Go语言版本，与Java版本具有完全相同的功能和非常相似的代码结构，具体信息可参考[Fenix's Bookstore的Spring Boot版本](https://icyfenix.cn/exploration/projects/monolithic_arch_springboot.html)。

## 运行程序

在已安装好Go语言环境，以及CGO编译器（使用的SQLite 3用到了CGO）的前提下，执行以下命令运行程序：

```bash
# 下载源码
git clone https://github.com/fenixsoft/monolithic_arch_golang.git && cd monolithic_arch_golang

# 编译
go build

# 运行程序
./monolithic_arch_golang
```

## 依赖信息

- [github.com/gin-gonic/gin](https://github.com/gin-gonic/gin) v1.6.3
  著名的Go语言Web框架，使用体检类似于[SparkJava](https://sparkjava.com/)。
- [gorm.io/gorm](https://gorm.io/gorm) v1.20.9
  著名的Go语言ORM框架。
- [github.com/mattn/go-sqlite3](https://github.com/mattn/go-sqlite3) v1.14.6 // indirect
  SQLite3数据库的Go语言版本。
- [github.com/GeertJohan/go.rice](https://github.com/GeertJohan/go.rice) v1.0.2
  资源打包框架，将静态的HTML、JS等资源文件嵌入二进制运行包中，获得单个文件运行的体验。
- [github.com/dgrijalva/jwt-go](https://github.com/dgrijalva/jwt-go) v3.2.0
  JWT令牌生成工具。
- [github.com/go-yaml/yaml](https://github.com/go-yaml/yaml) v2.1.0
  YAML文件解析工具。
- [github.com/sirupsen/logrus](https://github.com/sirupsen/logrus) v1.7.0
  日志框架。
- [golang.org/x/crypto](https://golang.org/x/crypto) v0.0.0-20201221181555
  Go的密码库套件，用了其中BCrypto算法。

## 协议

- 本作品代码部分采用[Apache 2.0协议](https://www.apache.org/licenses/LICENSE-2.0)进行许可。遵循许可的前提下，你可以自由地对代码进行修改，再发布，可以将代码用作商业用途。但要求你：
  - **署名**：在原有代码和衍生代码中，保留原作者署名及代码来源信息。
  - **保留许可证**：在原有代码和衍生代码中，保留Apache 2.0协议文件。

- 本作品文档部分采用[知识共享署名 4.0 国际许可协议](http://creativecommons.org/licenses/by/4.0/)进行许可。 遵循许可的前提下，你可以自由地共享，包括在任何媒介上以任何形式复制、发行本作品，亦可以自由地演绎、修改、转换或以本作品为基础进行二次创作。但要求你：
  - **署名**：应在使用本文档的全部或部分内容时候，注明原作者及来源信息。
  - **非商业性使用**：不得用于商业出版或其他任何带有商业性质的行为。如需商业使用，请联系作者。
  - **相同方式共享的条件**：在本文档基础上演绎、修改的作品，应当继续以知识共享署名 4.0国际许可协议进行许可。