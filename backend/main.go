package main

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

// Post ...
type Post struct {
	UID      int    `json:"uid"`
	Text     string `json:"text"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	Favorite bool   `json:"favorite"`
}

// PostsHandler ...
func PostsHandler(c *gin.Context) {
	fmt.Print("Handling Request for posts")

	posts := []Post{
		{
			1,
			"Wounded, captured and forced to build a weapon by his enemies, billionaire industrialist Tony Stark instead created an advanced suit of armor to save his life and escape captivity.",
			"Iron Man",
			"ironman.jpg",
			false,
		},
		{
			2,
			"Bitten by a radioactive spider, high school student Peter Parker gained the speed, strength and powers of a spider. Adopting the name Spider-Man, Peter hoped to start a career using his new ...",
			"Spider Man",
			"spidey.jpg",
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
	router.LoadHTMLFiles("frontend/*.html")
	router.Static("/front", "../frontend")

	router.Run(":8080")

	// Simple static webserver:
	// router.ServeFiles("/frontend", http.Dir("../frontend"))
	// log.Fatal(http.ListenAndServe(":8080", router))
	// log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("../frontend"))))
}
