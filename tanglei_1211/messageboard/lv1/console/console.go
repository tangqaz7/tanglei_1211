package console

//插入数据
import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"tanglei_1211/messageboard/configs"
)

var DB = DBInit()
//初始化数据库
func DBInit() *sql.DB {
	var username string
	db, err := sql.Open("mysql", configs.Connection)
	stmt, _ := db.Query("select * from data where ID = ?", username)
	fmt.Printf(
		"初始化用户数据库:%v",
		configs.Error(err),
	)
	var msgs []string
	for stmt.Next() {
		var m []string
		_ = stmt.Scan(`json:"username"`,`json:"password"`)
		msgs = append(msgs, m...)
		fmt.Println(msgs,"hhhhh")
	}

	return db
}

func User() (*sql.Stmt,error) {
	stmt, err := DB.Prepare("insert into data(ID ,PassWord, Name, RID, RName, RMessage, Good, IsGood) values (?,?,?,?,?,?,?,?)")

	return stmt,err
}
func Exec(stmt *sql.Stmt, userinfo *configs.User)  error{
	_, err := stmt.Exec(&userinfo.ID,
		userinfo.Name,
		userinfo.PassWord,
		userinfo.RID,
		userinfo.RName,
		userinfo.RMessage,
		userinfo.Good,
		userinfo.IsGood,
		)
	return err
}

