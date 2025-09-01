package task3

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //导入驱动 内部有init函数
	"math/rand"
	"strconv"
	"strings"
)

type Students struct {
	Id    int
	Name  string
	Age   int
	Grade string
}

func JustTestCurd() {
	url := "root:123456@tcp(127.0.0.1:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("链接失败:", err)
		return
	}
	//insertStudentTest(db)
	//selectAgeBig18(db)
	//updateStudentInfo(db)
	//deleteAgeLatter15(db)
	transAmount(db)
}

// 插入
func insertStudentTest(db *sql.DB) {
	fmt.Println("测试单条插入...")
	sqlStr := "insert into students(id,name,age,grade) values (?,?,?,?)"
	result, err := db.Exec(sqlStr, nil, "张三", 20, "三年级")
	if err != nil {
		fmt.Println(err)
		return
	}
	count, _ := result.RowsAffected()
	lastId, _ := result.LastInsertId()
	fmt.Println("单条插入完成，条数：", count, ", ID:", lastId)
	willAddList := make([]Students, 0)
	for i := 0; i < 10; i++ {
		n := fmt.Sprintf("测试%d号", rand.Intn(100))
		willAddList = append(willAddList, Students{Name: n, Age: rand.Intn(30), Grade: "六年级"})
	}
	ss := strings.Builder{}
	ss.WriteString("insert into students(id,name,age,grade) values ")
	for i, v := range willAddList {
		if i == len(willAddList)-1 {
			ss.WriteString("(NULL,'" + v.Name + "'," + string(strconv.Itoa(v.Age)) + ",'" + v.Grade + "');")
		} else {
			ss.WriteString("(NULL,'" + v.Name + "'," + string(strconv.Itoa(v.Age)) + ",'" + v.Grade + "'),")
		}
	}
	fmt.Println("开始执行批量插入")
	fmt.Println(ss.String())
	result, err = db.Exec(ss.String())
	if err != nil {
		fmt.Println(err)
		return
	}
	effectCount, _ := result.RowsAffected()
	fmt.Println("批量插入成功，影响行数：", effectCount)
}

// 查询年龄大于18的数据
func selectAgeBig18(db *sql.DB) {
	sqlStr := "select * from students where age > ?"
	row, err := db.Query(sqlStr, 18)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer row.Close()
	students := make([]Students, 0)
	for row.Next() {
		var id, age int
		var name, grade string
		row.Scan(&id, &name, &age, &grade)
		students = append(students, Students{Id: id, Name: name, Age: age, Grade: grade})
	}
	fmt.Println("selectAgeBig18结果：", students)
}

// 更新
func updateStudentInfo(db *sql.DB) {
	updateStr := "update students set grade='四年级' where name = ?"
	result, err := db.Exec(updateStr, "张三")
	if err != nil {
		fmt.Println(err)
		return
	}
	effectCount, _ := result.RowsAffected()
	fmt.Println("更新张三年级为四年级结果：", effectCount)
}

// 删除小于15岁的学生
func deleteAgeLatter15(db *sql.DB) {
	deleteStr := "delete from students where age < ?"
	result, err := db.Exec(deleteStr, 15)
	if err != nil {
		fmt.Println(err)
		return
	}
	effectCount, _ := result.RowsAffected()
	fmt.Println("删除小于15岁的学生结果：", effectCount)
}
func selectUserBalance(db *sql.DB, id int) float64 {
	sql := "select balance from accounts where id =?"
	row := db.QueryRow(sql, id)
	if row.Err() != nil {
		fmt.Println(row.Err())
		return 0
	}
	var balance float64
	err := row.Scan(&balance)
	if err != nil {
		fmt.Println("数据不存在")
		return 0
	}
	return balance
}

// 事务转账
func transAmount(db *sql.DB) {
	//用户A  B的余额
	var transAmont = float64(100)
	var fromId = 1
	var toId = 2
	var fromBalance = selectUserBalance(db, fromId)
	fmt.Println("用户:", fromId, " 转账前余额:", fromBalance)
	if fromBalance <= 0 {
		fmt.Println("转出账户余额不足")
		return
	}
	var toBalance = selectUserBalance(db, toId)
	fmt.Println("用户:", toId, " 转账前余额:", toBalance)
	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err)
		return
	}
	if fromBalance < transAmont {
		fmt.Println("转出账户余额不足")
		tx.Rollback()
		return
	}
	//修改转出账户
	updateFromBalanceStr := "update accounts set balance = (balance-?) where id = ?"
	result, err := tx.Exec(updateFromBalanceStr, transAmont, fromId)
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return
	}
	c, _ := result.RowsAffected()
	fmt.Println("修改转出账户余额：", c)
	//转入账户加值
	updateToBalanceStr := "update accounts set balance = (balance+?) where id = ?"
	result, err = tx.Exec(updateToBalanceStr, transAmont, toId)
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return
	}
	c, _ = result.RowsAffected()
	fmt.Println("修改转入账户余额：", c)
	//记录日志
	insertStr := "insert into transactions (id,from_account_id,to_account_id,amount) values (NULL,?,?,?);"
	result, err = tx.Exec(insertStr, fromId, toId, transAmont)
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return
	}
	c, _ = result.RowsAffected()
	fmt.Println("转账日志记录成功：", c)
	tx.Commit()
	fmt.Println("转账完成！")
}
