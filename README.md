# RestFullAPI service developed using Golang

To install all needed packages please go to your workspace $GOPATH/src and run these command below:
$ go get gopkg.in/gin-gonic/gin.v1
$ go get -u github.com/jinzhu/gorm
$ go get github.com/go-sql-driver/mysql

To setup connection with MySQL please use right credantials here:

Exemple: root:root@/homestead?charset=utf8&parseTime=True&loc=Local
  
"root:root" there are login and password for the database. Please use your own parametrs.

In generic the crud application has the API’s as follows:
GET    /api/articles             --> Get all articles
GET    /api/article/:id          --> Get an article with ID 
POST   /api/article              --> Create new article
PUT    /api/article/:id          --> Update the article with ID ("title" or "body")
DELETE /api/article/:id          --> Delete article with ID
