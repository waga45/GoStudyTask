package study

import (
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"math/rand"
	"sync"
	"time"
)

type User struct {
	Id          int64
	Name        string     `gorm:"type:varchar(50)"`
	Age         int        `gorm:"type:integer"`
	Phone       string     `gorm:"type:varchar(11)"`
	ProfileInfo string     `gorm:"type:varchar(1024)"`
	CreateTime  *time.Time `gorm:"type:datetime"`
	UpdateTime  *time.Time `gorm:"type:datetime"`
}

func (u *User) TableName() string {
	return "smt_user"
}

type UserBalance struct {
	Id         int64
	UserId     int64
	Balance    float64 `gorm:"type:decimal(10,2)"`
	CreateTime time.Time
	UpdateTime *time.Time
}

func (u *UserBalance) TableName() string {
	return "smt_user_balance"
}

// 复合体结构
type UserBaseInfo struct {
	UserId  int64  `gorm:"column:userId"`
	Name    string `gorm:"column:userName"`
	Balance float64
}

func genNewBalance(userId int64, amount float64) *UserBalance {
	return &UserBalance{UserId: userId, Balance: amount, CreateTime: time.Now(), UpdateTime: nil}
}

type Content struct {
	Id         int64 `gorm:"primary_key;autoIncrement:false"`
	title      string
	content    string `gorm:"type:text"`
	CreateTime time.Time
	UpdateTime *time.Time
}

func (c *Content) TableName() string {
	return "smt_content"
}

func MainTestGorm() {
	mysqlHost := "root:12345678@tcp(127.0.0.1:3306)/test_gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(mysqlHost), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	//自动建表
	autoCreateTable(db)
	//测试添加
	//addUser(db, "初始用户", 10, "18821234509", "有点意思")
	//更新
	//updateData(db)
	//删除
	//deleteData(db)
	//查询
	//selectData(db)
	//事务
	updateByTransaction(db)

}

func autoCreateTable(db *gorm.DB) {
	db.AutoMigrate(&User{})
	db.AutoMigrate(&UserBalance{})
	db.AutoMigrate(&Content{})
}

// 事务
func updateByTransaction(db *gorm.DB) {
	var fromUserId = 2            //转出人员id
	var toUserId = 3              //转入人员id
	var transAmount = float64(60) //转入金额
	var wg = sync.WaitGroup{}
	//模拟2个go进行操作
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			fmt.Printf("%d开始操作转账, 转出用户：%d,转入用户：%d, 转出金额：%f \n", index, fromUserId, toUserId, transAmount)
			//自动事务
			err := db.Transaction(func(tx *gorm.DB) error {
				//查询转出人金额是否足够
				var fromUser UserBalance
				rs := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("user_id=?", fromUserId).Find(&fromUser)
				if rs.Error != nil {
					fmt.Println("当前有正在转账进程，中断")
					return rs.Error
				}
				if &fromUser == nil {
					return errors.New("转出人账户不存在，请检查")
				}
				if fromUser.Balance < transAmount {
					return errors.New("转出余额不足，请检查")
				}
				tx.Model(&fromUser).Update("balance", (fromUser.Balance - transAmount))
				var toUser UserBalance
				ts := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("user_id=?", toUserId).Find(&toUser)
				if ts.Error != nil {
					return ts.Error
				}
				updateResult := tx.Model(&toUser).Update("balance", (toUser.Balance + transAmount))
				if updateResult.Error != nil {
					fmt.Println("转账失败")
					return updateResult.Error
				}
				fmt.Println(index, " 转账成功")
				//返回nil提交事务
				return nil
			})
			if err != nil {
				fmt.Println(err)
			}
		}(i)
	}
	wg.Wait()
	var fromUserB UserBalance
	db.Model(&UserBalance{}).Where("id=?", fromUserId).Find(&fromUserB)
	fmt.Println("转出人当前余额：", fromUserB)
	var toUserB UserBalance
	db.Model(&UserBalance{}).Where("id=?", toUserId).Find(&toUserB)
	fmt.Println("转入人当前余额：", toUserB)
}

func deleteData(db *gorm.DB) {
	//删除 根据ID删除
	result := db.Delete(&User{}, 12)
	fmt.Println("删除数据：", result.RowsAffected)
	//带显示条件删除
	result1 := db.Where("name = ? and age =?", "初始用户72", 66).Delete(&User{})
	fmt.Println("组合条件where删除数据：", result1.RowsAffected)
}

func updateData(db *gorm.DB) {
	lastUser := User{}
	db.Model(&User{}).Last(&lastUser)
	fmt.Println("最后一条 用户：", lastUser)
	//更新
	db.Model(&lastUser).Update("name", "更新最后一个用户")
	fmt.Println("单个字段更新完成")
	//多列更新
	result := db.Model(&lastUser).Updates(map[string]interface{}{
		"name": "有一次修改",
		"age":  20,
	})
	if result.Error != nil {
		fmt.Println("更新失败：", result.Error)
		return
	}
	fmt.Println("Updates 多列更新完成")
	//批量更新  clause.OnConflict  通过create主键唯一实现
	needUpdateUsers := []User{}
	needUpdateUsers = append(needUpdateUsers, User{Id: 2, Name: "批量更新值", Age: 18})
	needUpdateUsers = append(needUpdateUsers, User{Id: 3, Name: "OnConflict", Age: 19})
	needUpdateUsers = append(needUpdateUsers, User{Id: 4, Name: "good", Age: 20})
	result1 := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},                     //冲突列-一般是主建
		DoUpdates: clause.AssignmentColumns([]string{"name", "age"}), //需要更新的列
	}).Create(&needUpdateUsers)
	if result1.Error != nil {
		fmt.Println("批量更新失败：", result1.Error)
		return
	}
	fmt.Println("批量更新完成")
}
func selectData(db *gorm.DB) {
	u := User{}
	result := db.First(&u)
	if result.Error != nil {
		fmt.Println(result.Error)
		return
	}
	fmt.Println("第一个用户：", u)
	u1 := User{}
	db.Last(&u1)
	fmt.Println("最后一个用户：", u1)

	//查询年龄为10的用户
	users := make([]User, 0)
	db.Where("age = ?", 1).Find(&users)
	fmt.Println("Find查询集合结果：", users)
	//查询ID=2的用户名称和余额
	db.Where("name like ?", "%294%").Find(&users)
	fmt.Println("Find LIKE 结果查询集合结果：", users)
	//查询ID 在2-6只见数据
	db.Where("id in ?", []int{2, 3, 4, 5, 6}).Order("id desc").Find(&users)
	fmt.Println("Find In 结果查询集合结果：", users)
	//直接使用结构体作为条件，有点类似java的lamda表达式
	db.Select("id,name,age").Where(&User{Id: 12}).Find(&users)
	fmt.Println("Find 结构体条件 查询集合结果：", users)
	var count int64
	db.Table("smt_user").Where("id=?", 2).Count(&count)
	fmt.Println("总数据条数：", count)

	db.Distinct("name").Limit(5).Offset(0).Find(&users)
	fmt.Println("Find 分页查询结果：", users)

	baseUserInfos := make([]UserBaseInfo, 0)
	//join查询组合体
	db.Model(&UserBaseInfo{}).
		Raw("select su.id as userId,su.name as userName,sub.balance from smt_user su "+
			"LEFT JOIN smt_user_balance sub on sub.user_id=su.id "+
			"where su.id in ?", []int{2, 7}).
		Scan(&baseUserInfos)
	fmt.Println("RAW JOIN组合查询结果:", baseUserInfos)
}

func addUser(db *gorm.DB, name string, age int, phone string, profile string) {
	var t = time.Now()
	u := User{Name: name, Age: age, Phone: phone, ProfileInfo: profile, CreateTime: &t}
	//保存一个
	fmt.Println("开始执行单个插入")
	tx := db.Debug().Create(&u)
	if tx.Error != nil {
		panic(tx.Error)
		return
	}
	//创建钱包账号
	userBalance := genNewBalance(u.Id, 100)
	db.Debug().Create(userBalance)
	//批量保存
	tempUsers := make([]User, 0)
	for i := 0; i < 10; i++ {
		var t = time.Now()
		tempUsers = append(tempUsers, User{Name: fmt.Sprintf("%s%d", name, rand.Intn(1000)), Age: rand.Intn(100), Phone: phone, ProfileInfo: profile, CreateTime: &t})
	}
	fmt.Println("开始执行批量插入")
	tx1 := db.CreateInBatches(&tempUsers, 100)
	if tx1.Error != nil {
		panic(tx1.Error)
		return
	}
	fmt.Println("插入完成")
	for _, user := range tempUsers {
		userBalance := genNewBalance(user.Id, 100)
		db.Create(userBalance)
	}
	fmt.Println("账号创建完成")
}
