package model

/*
import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var Database *sql.DB

func InitDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)

	if err != nil {
		panic(err)
	}

	//If we don't get any errors but somehow still don't get a db connection
	//we exit as well

	if db == nil {
		panic("db nil")
	}

	return db
}

func Migrate(db *sql.DB) {
	sql := `
	CREATE TABLE IF NOT EXISTS deviceinfo(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		mac VARCHAR,
		conn INTEGER,
		online INTEGER,
		heartbeat VARCHAR
	);
	`
	_, err := db.Exec(sql)
	// Exit if something goes wrong with our SQL statement above

	if err != nil {
		panic(err)
	}
}

func InsertMac(db *sql.DB, mac string) {
	stmt, err := db.Prepare("INSERT INTO deviceinfo(mac, online) values(?, ?) ")
	checkErr(err)

	res, err := stmt.Exec(mac, 0)
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println("id=", id)
}

func DeleteMac(db *sql.DB, mac string) {
	stmt, err := db.Prepare("DELETE from deviceinfo where mac=?")
	checkErr(err)

	res, err := stmt.Exec(mac)
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println("id=", id)
}

func UpdateDeviceOnlineStatus(db *sql.DB, mac string, online int) {
	stmt, err := db.Prepare("UPDATE deviceinfo SET online=? where mac=?")
	checkErr(err)

	res, err := stmt.Exec(online, mac)
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println("id=", id)
}

func UpdateDeviceLastHeartbeat(db *sql.DB, mac, heartbeat string) {
	stmt, err := db.Prepare("UPDATE deviceinfo SET heartbeat=? where mac=?")
	checkErr(err)

	res, err := stmt.Exec(heartbeat, mac)
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println("id=", id)
}

func QueryMac(db *sql.DB, mac string) (int, string, error) {

	var err error
	//var rows *sql.Rows

	var id, online int
	var heartbeat string

	if mac != "" {
		rows := db.QueryRow("select id,online,heartbeat from deviceinfo where mac=?", mac)
		err = rows.Scan(&id, &online, &heartbeat)
		fmt.Printf("id=%d, online=%d, mac=%s, heartbeat=%s\n",
			id, online, mac, heartbeat)
	} else {
		return -1, "", errors.New("No name assigned")
	}
	fmt.Println("QueryMac id", id)
	return id, heartbeat, err
}

func WalkDeviceInfo(db *sql.DB, num int) {

}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

//func main() {
//	db := InitDB("test.db")
//	defer db.Close()

//	Migrate(db)
//	InsertMac(db, "xxxxxddd")
//	QueryMac(db, "xxxxxddd")
//	UpdateDeviceOnlineStatus(db, "xxxxxddd", 1)
//	UpdateDeviceLastHeartbeat(db, "xxxxxddd", "test")
//	QueryMac(db, "xxxxxddd")
//	DeleteMac(db, "xxxxxddd")

//}
*/
