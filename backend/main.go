package main

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Post struct {
	Uid      int    `json:"uid"`
	Text     string `json:"text"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	Favorite bool   `json:"favorite"`
}

func PostsHandler(c *gin.Context) {
	fmt.Print("Handling Request for posts")

	posts := []Post{
		{
			1,
			"First",
			"Miguel",
			"miguel_avatar",
			false,
		},
		{
			2,
			"Second",
			"Laura",
			"laura_avatar",
			false,
		},
	}

	bs, err := json.Marshal(posts)
	if err != nil {
		//ReturnError(w, err)
		return
	}
	fmt.Fprint(c.Writer, string(bs))
}

func main() {

	fmt.Println("Listening on :8080")

	router := gin.Default()
	router.GET("/posts", PostsHandler)

	static := router.Group("/")
	{
		static.Static("/test", "../frontend")
	}
	router.Run(":8080")

	// Simple static webserver:
	// router.ServeFiles("/frontend", http.Dir("../frontend"))
	// log.Fatal(http.ListenAndServe(":8080", router))
	// log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("../frontend"))))
}
