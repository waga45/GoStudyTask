package study

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type UserX struct {
	Id          int64      `db:"id"`
	Name        string     `db:"name"`
	Age         int        `db:"age"`
	Phone       string     `db:"phone"`
	ProfileInfo string     `db:"profile_info"`
	CreateTime  *time.Time `db:"create_time"`
	UpdateTime  *time.Time `db:"update_time"`
}

type UserBaseInfoX struct {
	UserId  int64   `db:"userId"`
	Name    string  `db:"userName"`
	Balance float64 `db:"balance"`
}

func SQLXTest() {
	mysqlHost := "root:12345678@tcp(127.0.0.1:3306)/test_gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(mysqlHost), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	//SQLX
	forSqlx(db)
}

// sqlx操作
func forSqlx(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println(err)
		return
	}

	sqlxDB := sqlx.NewDb(sqlDB, "mysql")
	defer sqlxDB.Close()
	selectOne(sqlxDB)
	selectListIn(sqlxDB)
	selectNameQuery(sqlxDB)
}

// Get 只能查询一条
func selectOne(sqlxDB *sqlx.DB) {
	ql := "select  * from smt_user su where su.id=?"
	bean := UserX{}
	ERR := sqlxDB.Get(&bean, ql, 2)
	if ERR != nil {
		fmt.Println(ERR)
	}
	fmt.Println(bean)
}

func selectListIn(sqlxDB *sqlx.DB) {
	sl := "select su.id as userId,su.name as userName,sub.balance " +
		"from smt_user su " +
		"join smt_user_balance sub on sub.user_id=su.id " +
		"where su.id in (?)"

	q, args, err := sqlx.In(sl, []int{2, 3, 4})
	if err != nil {
		fmt.Println(err)
		return
	}
	//组装原始sql
	q = sqlxDB.Rebind(q)
	fmt.Println(q)
	list := make([]UserBaseInfoX, 0)
	//执行
	err = sqlxDB.Select(&list, q, args...)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("selectList:", list)
}

// 使用map参数
func selectNameQuery(sqlxDB *sqlx.DB) {
	sl := "select * from smt_user where name = :name and age= :age "
	list := make([]UserX, 0)
	rows, err := sqlxDB.NamedQuery(sl, map[string]interface{}{
		"name": "OnConflict",
		"age":  19,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	//需要关闭
	defer rows.Close()
	//需要手动解析
	for rows.Next() {
		var user UserX
		err = rows.StructScan(&user)
		if err != nil {
			fmt.Println("解析失败")
			continue
		}
		list = append(list, user)
	}
	fmt.Println("selectNameQuery:", list)
}
