package main

import (
	"github.com/coopernurse/gorp"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

var dbmap = initDb()

func main() {

	defer dbmap.Db.Close()

	//router := gin.Default()
	router := Router()
	//router.GET("/articles", ArticlesList)  	//curl -i http://localhost:8000/articles
	//router.POST("/articles", ArticlePost)  //	//curl -i -X POST -H "Content-Type: application/json" -d "{ \"Title\": \"Thea\", \"Content\": \"Queen\" }" http://localhost:8000/articles
	//router.GET("/articles/:id", ArticlesDetail) //	//curl -i http://localhost:8000/articles/{article number}

	router.Run(":8000")
}

func Router() *gin.Engine {

	router := gin.Default()
	router.GET("/articles", ArticlesList)       //curl -i http://localhost:8000/articles
	router.POST("/articles", ArticlePost)       //	//curl -i -X POST -H "Content-Type: application/json" -d "{ \"Title\": \"Thea\", \"Content\": \"Queen\" }" http://localhost:8000/articles
	router.GET("/articles/:id", ArticlesDetail) //	//curl -i http://localhost:8000/articles/{article number}

	return router
}

type Article struct {
	Id      int64 `db:"article_id"`
	Created int64
	Title   string
	Content string
}

func createArticle(title, body string) Article {
	article := Article{
		Created: time.Now().UnixNano(),
		Title:   title,
		Content: body,
	}

	err := dbmap.Insert(&article)
	checkErr(err, "Insert failed")
	return article
}

func getArticle(article_id int) Article {
	article := Article{}
	err := dbmap.SelectOne(&article, "select * from articles3 where article_id=$1", article_id)
	checkErr(err, "SelectOne failed")
	return article
}

func ArticlesList(c *gin.Context) {
	var articles []Article
	_, err := dbmap.Select(&articles, "select * from articles3 order by article_id")
	checkErr(err, "Select failed")
	content := gin.H{}
	for k, v := range articles {
		content[strconv.Itoa(k)] = v
	}
	c.JSON(200, content)

}

func ArticlesDetail(c *gin.Context) {
	article_id := c.Params.ByName("id")
	a_id, _ := strconv.Atoi(article_id)
	article := getArticle(a_id)
	content := gin.H{"title": article.Title, "content": article.Content}
	c.JSON(200, content)
}

func ArticlePost(c *gin.Context) {
	var json Article

	c.Bind(&json) // This will infer what binder to use depending on the content-type header.
	article := createArticle(json.Title, json.Content)
	if article.Title == json.Title {
		content := gin.H{
			"result":  "Success",
			"title":   article.Title,
			"content": article.Content,
		}
		c.JSON(201, content)
	} else {
		c.JSON(500, gin.H{"result": "An error occured"})
	}

}

func initDb() *gorp.DbMap {

	dbUrl := os.Getenv("DATABASE_URL_THECROYDONPROJECT") //export DATABASE_URL_THECROYDONPROJECT="dbname=databasename user=databaseusername password=password host=localhost port=15432 sslmode=disable"

	fmt.Println("DB URL Connection is --> " + dbUrl)

	db, err := sql.Open("postgres", dbUrl)
	//db, err := sql.Open("sqlite3", "/tmp/post_db.bin")
	checkErr(err, "sql.Open failed")

	// construct a gorp DbMap
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}

	dbmap.AddTableWithName(Article{}, "articles3").SetKeys(true, "Id")
	err = dbmap.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed")

	return dbmap
}

func checkErr(err error, msg string) {
	if err != nil {
		//log.Fatalln(msg, err)

		log.Print(err.Error())
	}
}
