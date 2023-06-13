package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //没用到包里的内容但是需要加载一下这个包,import该包，如果该包有init函数，则会先执行init函数
	"strings"
	"time"
)

const (
	username     = "root"
	password     = "12345678"
	ipAddr       = "127.0.0.1"
	port         = "3307"
	databaseName = "mytest"
)

// DB 数据库连接池对象
var DB *sql.DB

type PERSON struct {
	Id   int
	Name string
	Age  int
	Tall int
}

type DBCONF struct {
	Username        string
	Password        string
	Host            string
	Port            string
	DbName          string
	ConnMaxLiftTime int64 //设置数据库最大连接数
	MaxIdleConn     int   //设置上数据库最大闲置连接数
}

// InitDB 初始化连接
//func InitDB(c *DBCONF) {
//	//dsn := strings.Join([]string{username, ":", password, "@tcp", "(", ipAddr, ":", port, ")/", databaseName}, "")
//	dsn := strings.Join([]string{c.Username, ":", c.Password, "@tcp", "(", c.Host, ":", c.Port, ")/", c.DbName}, "")
//	DB, _ = sql.Open("mysql", dsn)
//	//设置数据库最大连接数
//	DB.SetConnMaxLifetime(time.Duration(c.ConnMaxLiftTime))
//	//设置上数据库最大闲置连接数
//	DB.SetMaxIdleConns(c.MaxIdleConn)
//	//验证连接
//	if err := DB.Ping(); err != nil {
//		fmt.Println("open database fail")
//		return
//	}
//	fmt.Println("connect success")
//}

func InitDB(c *DBCONF) *sql.DB {
	//dsn := strings.Join([]string{username, ":", password, "@tcp", "(", ipAddr, ":", port, ")/", databaseName}, "")
	dsn := strings.Join([]string{c.Username, ":", c.Password, "@tcp", "(", c.Host, ":", c.Port, ")/", c.DbName}, "")
	DB, _ := sql.Open("mysql", dsn)
	//设置数据库最大连接数
	DB.SetConnMaxLifetime(time.Duration(c.ConnMaxLiftTime))
	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(c.MaxIdleConn)
	//验证连接
	if err := DB.Ping(); err != nil {
		fmt.Println("open database fail")
		return nil
	}
	fmt.Println("connect success")
	return DB
}

// QueryRow 查询单行
func QueryRow() {
	var person PERSON

	err := DB.QueryRow("select * from person where id = ?", 1).
		Scan(&person.Id, &person.Name, &person.Age, &person.Tall)
	if err != nil {
		fmt.Println(err, "queryRow error")
	} else {
		fmt.Printf("p: %+v", person)
	}
}

// Query 查询多行
func Query() {
	var person PERSON
	rows, err := DB.Query("select * from person")
	if err != nil {
		return
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println("closing db object error", err)
		}
	}(rows) //关闭数据库连接
	fmt.Printf("err: %v, rows: %v, \n", err, rows)
	if err != nil {
		errors.New("query incr error")
	}

	//循环取出查询出的多条记录
	for rows.Next() {
		//这里的Scan操作需要传入实体对象的每个字段的地址
		e := rows.Scan(&person.Id, &person.Name, &person.Age, &person.Tall)
		fmt.Println(person)
		if e != nil {
			fmt.Println(json.Marshal(person))
		}
	}
}

func InsertOne(n string, a, t int) (result sql.Result, err error) {
	stmt, err := DB.Prepare(`insert into person(name,age,tall) values(?,?,?)`)
	if err != nil {
		fmt.Println("prepare catch a error")
	}
	defer func(stmt *sql.Stmt) {
		err = stmt.Close()
		if err != nil {
			fmt.Println("close catch a error")
		}
	}(stmt)

	exec, err := stmt.Exec(n, a, t)
	if err != nil {
		return nil, err
	}

	return exec, nil
}

// GetNowStr 获取当前时间的字符串
func GetNowStr() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func main() {
	conf := &DBCONF{
		"root",
		"12345678",
		"127.0.0.1",
		"3307",
		"mytest",
		100,
		10,
	}
	db := InitDB(conf)
	DB = db
	QueryRow()
	fmt.Println(GetNowStr())
	//Query()

	//res, _ := InsertOne("kkkk", 100, 190)
	//fmt.Println(res)
}
