package handler

import (
	"net/http"

	"github.com/endalk200/blazingly-fast-url-shortener/shortener"
	"github.com/endalk200/blazingly-fast-url-shortener/store"
	"github.com/gin-gonic/gin"
)

// Request model definition
type UrlCreationRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
	UserId  string `json:"user_id" binding:"required"`
}

// @summary Create a short url
// @description Create a short url
// @tags urls
// @accept json
// @produce json
// @param urlCreationRequest body UrlCreationRequest true "Url Creation"
// @success 200 {object} string "short url created successfully"
// @router /urls [post]
func CreateShortUrl(c *gin.Context) {
	var creationRequest UrlCreationRequest

	if err := c.ShouldBindJSON(&creationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortUrl := shortener.GenerateShortLink(creationRequest.LongUrl, creationRequest.UserId)
	store.SaveUrlMapping(shortUrl, creationRequest.LongUrl, creationRequest.UserId)

	host := "http://localhost:9808/"

	c.JSON(200, gin.H{
		"message":   "short url created successfully",
		"short_url": host + shortUrl,
		"key":       shortUrl,
	})

}

// @summary Get all short urls
// @description Get all short urls
// @tags urls
// @produce json
// @success 200 {object} string "short url created successfully"
// @router /urls [get]
func GetShortUrls(c *gin.Context) {
	storedUrl := store.GetUrls()

	c.JSON(200, gin.H{
		"urls": storedUrl,
	})
}

// @summary Get a short url
// @description Get a short url
// @tags urls
// @produce json
// @param shortUrl path string true "short url"
// @success 200 {object} string "short url created successfully"
// @router /urls/{shortUrl} [get]
func GetUrlObj(c *gin.Context) {
	shortUrl := c.Param("shortUrl")

	initialUrl := store.RetrieveInitialUrl(shortUrl)

	c.JSON(200, gin.H{
		"short_url": shortUrl,
		"long_url":  initialUrl,
	})
}

// @summary Redirect to initial url
// @description Redirect to initial url
// @tags urls
// @produce json
// @param shortUrl path string true "short url"
// @success 200 {object} string "short url created successfully"
// @router /{shortUrl} [get]
func HandleShortUrlRedirect(c *gin.Context) {
	shortUrl := c.Param("shortUrl")

	initialUrl := store.RetrieveInitialUrl(shortUrl)

	c.Redirect(302, initialUrl)
}
