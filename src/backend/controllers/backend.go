package controllers

import (
	"fmt"
	"net/http"

	"github.com/ittechman101/gobackend/src/parser"

	"github.com/gin-gonic/gin"
)

// POST /book_text
// post book by text
func BookByText(c *gin.Context) {
	// Validate input
	text := c.PostForm("book_text")
	fmt.Println(text)

	topTenWords := parser.ParseText(text)

	fmt.Println(topTenWords[0].Word)
	//j, _ := json.Marshal(topTenWords)
	//log.Print(j)
	c.JSON(http.StatusOK, topTenWords)
}
