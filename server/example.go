package main

import (
	"database/sql"
	"fmt"

	"github.com/kataras/iris"
	_ "github.com/mattn/go-sqlite3"
)

type Event struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Start    string `json:"start"`
	Category int    `json:"category"`
}

type News struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	URI   string `json:"uri"`
	EID   int    `json:"eid"`
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	app := iris.Default()
	event1 := Event{
		ID:       1,
		Title:    "XXX事件",
		Content:  "XXX在XXX因XXX被XXX",
		Start:    "12月14日",
		Category: 0,
	}
	event2 := Event{
		ID:       2,
		Title:    "XXX事件2",
		Content:  "XXX在XXX因XXX被XXX",
		Start:    "12月15日",
		Category: 0,
	}
	var events = [...]Event{event1, event2}

	db, err := sql.Open("sqlite3", "./vptest.db")
	checkErr(err)
	rows, err := db.Query("SELECT * FROM event")
	checkErr(err)

	for rows.Next() {
		event := Event{}
		err = rows.Scan(&event.ID, &event.Title, &event.Content, &event.Start, &event.Category)
		checkErr(err)
		fmt.Println(event)
	}

	news1 := News{
		ID:    1,
		Title: "XXX*XXX*XXX",
		URI:   "www.baidu.com",
		EID:   1,
	}

	app.Get("/news/list", func(ctx iris.Context) {
		ctx.JSON(iris.Map{
			"code": 200,
			"data": events,
		})
	})
	app.Get("/news/id/{id:int}", func(ctx iris.Context) {
		var id int
		id, _ = ctx.Params().GetInt("id")
		if id == 1 {
			news1.Title = news1.Title + "HHHHHH"
			ctx.JSON(iris.Map{
				"code": 200,
				"data": news1,
			})
		} else {
			ctx.JSON(iris.Map{
				"code": 200,
				"data": nil,
			})
		}
	})
	app.Get("/life/list", func(ctx iris.Context) {
		ctx.JSON(iris.Map{
			"message": "pong",
		})
	})
	// listen and serve on http://0.0.0.0:8080.
	app.Run(iris.Addr("localhost:8080"))
}
