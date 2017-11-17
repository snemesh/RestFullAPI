package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

var myDb *gorm.DB

func init() {
	//open a db connection
	var err error
	myDb, err = gorm.Open("mysql", "root:root@/homestead?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic("We can't get an access to the database. Please check credentials!")
	}
	//Migrate the schema
	myDb.AutoMigrate(&articleModel{})
}

type (
	// articleModel describes a articleModel type
	articleModel struct {
		gorm.Model
		Title     string `json:"title,omitempty"`
		Body 	  string `json:"body,omitempty"`
	}
	// transformed output data
	transformedArticle struct {
		ID        uint   `json:"id"`
		Title     string `json:"title"`
		Body 	  string `json:"body"`
	}

)

func main() {
	router := gin.Default()
	v1 := router.Group("/api/")
	{
		v1.GET("articles", getArticles)
		v1.GET("article/:id", getArticle)
		v1.POST("article", postArticle)
		v1.PUT("article/:id", updateArticle)
		v1.DELETE("article/:id", deleteArticle)

	}
	router.Run()
}

// Transform Article model for representing
func transformArticle(art []articleModel) []transformedArticle {

	var _article []transformedArticle

	//Casting Article model to transformedArticle. Output data will have only 3 fields ID, Title, Body
	for _, item := range art {
		_article = append(_article,

			transformedArticle{ID: item.ID,
			Title: item.Title,
			Body: item.Body})
	}

	return _article

}

// Get all article from the table
func getArticles (c *gin.Context)  {

	var article []articleModel
	var _article []transformedArticle

	myDb.Find(&article)
	_article = transformArticle(article)

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _article})

}

// Get just only one record about article. ID takes from URI parameters
func getArticle (c *gin.Context)  {
	var article articleModel
	articleID := c.Param("id")

	//We try to find article with needed ID
	myDb.First(&article, articleID)

	// If we don't find - return message that we didn't find anything
	if article.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Record about the record с ID = " + articleID + " - not found"})
		return
	}

	// JSON respond with status
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": article})


}

//Creating new article with Title and Body
func postArticle (c *gin.Context)  {

	var article articleModel

	newTitle := c.PostForm("title")
	newBody := c.PostForm("body")

	// If we didn't pass Title we will put empty field on the new record
	if newTitle != "" {
		article.Title = newTitle
	} else {
		article.Title = ""
	}

	// If we didn't pass Body we will put empty field on the new record
	if newBody != "" {
		article.Body = newBody
	} else {
		article.Body = ""
	}

	myDb.Save(&article) //Save new record to the database

	// JSON respond with status
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Article item created successfully!", "resourceId": article.ID})

}

// Update article information
func updateArticle (c *gin.Context)  {

	var article articleModel
	articleID := c.Param("id")

	// We are trying to find record with ID
	myDb.First(&article, articleID)

	// Checking, is this record was found?
	if article.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No article found!"})
		return
	}

	// If the record was found and the field Title is not empty, we update this field
	if article.Title != c.PostForm("title") && c.PostForm("title") != "" {
		myDb.Model(&article).Update("title", c.PostForm("title"))
	}

	// If the record was found and the field Body is not empty, we update this field
	if article.Body != c.PostForm("body") && c.PostForm("body") != "" {
		myDb.Model(&article).Update("body", c.PostForm("body"))
	}

	// JSON respond with status
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Todo updated successfully!"})

}

func deleteArticle (c *gin.Context)  {

	var article articleModel
	articleID := c.Param("id")

	// We are trying to find record with ID
	myDb.First(&article, articleID) //Find article with ID

	// Checking, is this record was found?
	if article.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Article with ID = " + articleID + " not found!"})
		return
	}

	//Delete the record
	myDb.Delete(&article)

	// JSON respond with status
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "The article with ID = " + articleID +  " deleted successfully!"})
	
}

