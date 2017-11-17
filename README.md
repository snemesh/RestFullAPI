# RestFullAPI service developed using Golang

To install all needed packages please go to your workspace $GOPATH/src and run these command below:
$ go get gopkg.in/gin-gonic/gin.v1
$ go get -u github.com/jinzhu/gorm
$ go get github.com/go-sql-driver/mysql

# MySQL settings
To setup connection with MySQL please use right credantials here:

Exemple: root:root@/homestead?charset=utf8&parseTime=True&loc=Local
  
"root:root" there are login and password for the database. Please use your own parametrs.

# API
In generic the crud application has the API’s as follows:
GET    /api/articles             --> Get all articles
GET    /api/article/:id          --> Get an article with ID 
POST   /api/article              --> Create new article
PUT    /api/article/:id          --> Update the article with ID ("title" or "body")
DELETE /api/article/:id          --> Delete article with ID

# License

Copyright (c) 2017 Sergey Nemesh.

Use of the code provided on this repository is subject to the MIT License.

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
