package main

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func main() {
	r := gin.Default()

	rdb_title := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       1,  // use default DB
	})

	rdb_main_actor := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       2,  // use default DB
	})

	rdb_actor := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       3,  // use default DB
	})

	ctx := context.Background()

	r.GET("/title/:id", func(c *gin.Context) {
		id := c.Param("id")
		val, _ := rdb_title.Get(ctx, id).Result()
		c.JSON(http.StatusOK, strings.Replace(val, "'", `"`, -1))
	})

	r.GET("/main_actors/:id", func(c *gin.Context) {
		id := c.Param("id")
		val, _ := rdb_main_actor.Get(ctx, id).Result()
		c.JSON(http.StatusOK, strings.Replace(val, "'", `"`, -1))
	})

	r.GET("/actor/:id", func(c *gin.Context) {
		id := c.Param("id")
		val, _ := rdb_actor.Get(ctx, id).Result()
		c.JSON(http.StatusOK, strings.Replace(val, "'", `"`, -1))
	})

	r.Run(":8002") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
