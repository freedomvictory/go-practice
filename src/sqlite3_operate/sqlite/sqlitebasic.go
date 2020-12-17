package sqlite
import(
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"time"
)




func sqlite3_demo(){
	db, err := sql.Open("sqlite3", "foo.db")
	checkErr(err)

	sqlTable := `
	 CREATE TABLE IF NOT EXISTS userinfo(
        uid INTEGER PRIMARY KEY AUTOINCREMENT,
        username VARCHAR(64) NULL,
        departname VARCHAR(64) NULL,
        created DATE NULL
    );
	`
	db.Exec(sqlTable)

	stmt, err := db.Prepare("INSERT INTO userinfo(username, departname, created) values(?,?,?)")
	checkErr(err)
	defer stmt.Close()

	res, err := stmt.Exec("wangshubo", "guowuyuan", "2017-04-21")
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	log.Printf("LastInsertId:%v\n" ,id)

	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)

	res, err = stmt.Exec("wangshubo_new", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)
	log.Printf("affect: %v", affect)

	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)

	var uid int
	var username , department string
	var created time.Time

	for rows.Next(){
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		log.Printf("uid:%d, username:%s, department:%s, created:%v\n", uid, username, department, created)
	}
	rows.Close()


	db.Close()

}

func checkErr(err error){
	if err != nil{
		panic(err)
	}
}