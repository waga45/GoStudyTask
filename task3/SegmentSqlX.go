package task3

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/shopspring/decimal"
)

type Employees struct {
	Id         int             `db:"id"`
	Name       string          `db:"name"`
	Department string          `db:"department"`
	Salary     decimal.Decimal `db:"salary"`
}

func TestSqlX() {
	url := "root:123456@tcp(127.0.0.1:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sqlx.Open("mysql", url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	selectAllTechnical(db)
	selectHeightSalary(db)
	selectBooks(db)
}

// 表中所有部门为 "技术部" 的员工信息
func selectAllTechnical(db *sqlx.DB) {
	var employees []Employees
	strSql := "select * from employees where department =?"
	err := db.Select(&employees, strSql, "技术部")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("selectAllTechnical结果：", employees)
}

// 表中工资最高的员工信息
func selectHeightSalary(db *sqlx.DB) {
	var emp Employees
	strSql := "select * from employees order by salary desc limit 1"
	err := db.Get(&emp, strSql)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("selectHeightSalary结果：", emp)
}

type Books struct {
	Id     int             `db:"id"`             //id
	Title  string          `db:"title"`          //名称
	Author *string         `db:"author"`         //作者
	Price  decimal.Decimal `db:"price"`          //售价
	Salary decimal.Decimal `db:"employeeSalary"` //薪资
}

// select
func selectBooks(db *sqlx.DB) {
	strSql := "select bk.id,bk.title,bk.author,bk.price,es.salary as employeeSalary from books  bk " +
		"JOIN employees es on es.name=bk.author " +
		"where bk.author=? " +
		"and bk.price >=? "
	var books []Books
	err := db.Select(&books, strSql, "王五", 50)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("selectBooks结果：", books)
}
