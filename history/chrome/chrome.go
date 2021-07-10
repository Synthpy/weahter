package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// //	db, err := sql.Open("sqlite3", "/Users/xigua/Library/Application Support/Google/Chrome/Default/Historyc")
	// db, err := sql.Open("sqlite3", "/Users/xigua/Library/Safari/History.db")
	// checkErr(err)
	// now := int(time.Now().Unix())
	// fmt.Println(now)
	// //res, err := db.Query("select id, url, title from urls limit 5")
	// //res, err := db.Query("select id, url, visit_count from history_items limit 5")
	// t := time.Now()
	// //t = t.AddDate(0, 0, 1)
	// fmt.Println(t)
	// tm1 := int(time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).Unix())
	// tm2 := tm1 - 978307200
	// fmt.Println(tm2)
	// checkErr(err)
	// res, err := db.Query("select b.url,b.visit_count, a.history_item,CAST(a.visit_time as int), a.title from history_visits as a  join history_items as b on a.history_item=b.id where a.visit_time>?", tm2)

	// }
	// defer db.Close()

	res, err := getSafariHistory()
	checkErr(err)
	var visit_count string
	var url string
	var title string
	var visit_time string
	for res.Next() {
		res.Scan(&url, &visit_count, &visit_time, &title)
		fmt.Println(url, visit_count, visit_time, title)
		fmt.Println("--------")

	}

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func getSafariHistory() (*sql.Rows, error) {
	db, err := sql.Open("sqlite3", "/Users/xigua/Library/Safari/History.db")
	checkErr(err)
	t := time.Now()
	tm1 := int(time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).Unix())                                                                                                                                     // 获取当天0点的时间戳
	tm2 := tm1 - 978307200                                                                                                                                                                                                   //nsdate（apple一种时间格式）转换为unix时间戳
	res, err := db.Query("select b.url,b.visit_count, datetime(a.visit_time+978307200,'unixepoch','localtime'), a.title from history_visits as a  join history_items as b on a.history_item=b.id where a.visit_time>?", tm2) //查询Safari历史记录
	checkErr(err)
	defer db.Close()
	return res, err

}

func getChromeHistroy() (*sql.Rows, error) {
	ut := 11644473600
	t := time.Now()
	nu := int(time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).Unix()) // 获取当天0点的时间戳

	nasni := (ut + nu) * 1000000
	db, err := sql.Open("sqlite3", "/Users/xigua/Library/Application Support/Google/Chrome/Default/Historyc")
	checkErr(err)
	res, err := db.Query("select url, visit_count,datetime(last_visit_time / 1000000 + (strftime('%s', '1601-01-01')), 'unixepoch'),title from urls where last_visit_time>?", nasni)
	checkErr(err)
	defer db.Close()
	return res, err
}
