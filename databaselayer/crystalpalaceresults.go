package main

import (
	"database/sql"
	"fmt"
	"github.com/coopernurse/gorp"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
	"os"
	"strconv"
	"time"
)

var dbmap = initDb()

type Result struct {
	Id              int64 `db:"result_id"`
	Created         int64
	Season          string
	Round           string
	Date            string
	Kickofftime     string
	AwayorHome      string
	Oppenent        string
	Resultshalftime string
	Resultsfulltime string
}

func initDb() *gorp.DbMap {

	dbUrl := os.Getenv("DATABASE_URL_THECROYDONPROJECT") //export DATABASE_URL_THECROYDONPROJECT="dbname=databasename user=databaseusername password=password host=localhost port=15432 sslmode=disable"

	fmt.Println("DB URL Connection is --> " + dbUrl)

	db, err := sql.Open("postgres", dbUrl)
	checkErr(err, "sql.Open failed")

	// construct a gorp DbMap
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}

	dbmap.AddTableWithName(Result{}, "eagles1").SetKeys(true, "Id")
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

func main() {

	defer dbmap.Db.Close()
	router := Router()
	router.Run(":8000")
}

func Router() *gin.Engine {

	router := gin.Default()
	router.GET("/results", allresults)        //curl -i http://localhost:8000/results
	router.POST("/result", postresultentry)   //	curl -i -X POST -H "Content-Type: application/json" -d "{\"Season\":\"1945/46\",\"Round\":\"15\",\"Date\":\"10-09-1946\",\"Kickofftime\":\"13:00\",\"AwayorHome\":\"A\",\"Oppenent\":\"Arsenal\",\"Resultshalftime\":\"1:2\",\"Resultsfulltime\":\"2:2\"}" http://localhost:8000/result
	router.GET("/results/:id", resultdetails) //	//curl -i http://localhost:8000/results/{result number}

	return router
}

func createresultentry(season, round, date, kickofftime, awayorhome, opponent, resulthalftime, resultfultime string) Result {

	result := Result{

		Created:         time.Now().UnixNano(),
		Season:          season,
		Round:           round,
		Date:            date,
		Kickofftime:     kickofftime,
		AwayorHome:      awayorhome,
		Oppenent:        opponent,
		Resultshalftime: resulthalftime,
		Resultsfulltime: resultfultime,
	}
	err := dbmap.Insert(&result)
	checkErr(err, "Insert failed")

	return result
}

func postresultentry(c *gin.Context) {

	var json Result

	c.Bind(&json) // This will infer what binder to use depending on the content-type header.

	result := createresultentry(json.Season, json.Round, json.Date, json.Kickofftime, json.AwayorHome, json.Oppenent, json.Resultshalftime, json.Resultsfulltime)

	if result.Season == json.Season {
		c.JSON(201, result)
	} else {
		c.JSON(500, gin.H{"result": "An error occured"})
	}

	c.JSON(201, result)

}

func getresult(result_id int) Result {

	result := Result{}
	err := dbmap.SelectOne(&result, "select * from eagles1 where result_id=$1", result_id)
	checkErr(err, "SelectOne failed")
	return result
}

func resultdetails(c *gin.Context) {
	result_id := c.Params.ByName("id")
	r_id, _ := strconv.Atoi(result_id)
	result := getresult(r_id)

	content := gin.H{"Season": result.Season, "Round": result.Round, "Date": result.Date, "Kickofftime": result.Kickofftime, "AwayorHome": result.AwayorHome, "Oppenent": result.Oppenent, "Resultshalftime": result.Resultshalftime, "Resultsfulltime": result.Resultsfulltime}

	c.JSON(200, content)
}

func allresults(c *gin.Context) {

	var result []Result

	_, err := dbmap.Select(&result, "select * from eagles1 order by result_id")

	checkErr(err, "Select failed")

	content := gin.H{}

	for k, v := range result {
		content[strconv.Itoa(k)] = v
	}
	c.JSON(200, content)

}
