package task3

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"math/rand"
	"strconv"
	"time"
)

type User struct {
	Id         int        `gorm:"primary_key;autoIncrement:true"`
	Name       string     `gorm:"type:varchar(50)"`
	PostNum    int        `gorm:"type:int"`
	CreateTime *time.Time `gorm:"type:datetime"`
}

func (u *User) TableName() string {
	return "user"
}

// 文章
type Post struct {
	Id         int
	UserId     int
	Title      *string
	Content    *string
	RateNum    int        `gorm:"type:int"` //评论数
	RawStatus  int        `gorm:"type:Integer"`
	CreateTime *time.Time `gorm:"type:datetime"`
	UpdateTime *time.Time `gorm:"type:datetime"`
}
type PostPage struct {
	Count int64
	List  []Post
}

// 钩子-创建文章 更新用户表统计数
func (c *Post) AfterCreate(tx *gorm.DB) (err error) {
	//创建程，更新用户文章数
	r := tx.Model(&User{}).Where(&User{Id: c.UserId}).Update("post_num", gorm.Expr("post_num + ?", 1))
	if r.Error != nil {
		fmt.Println("同步更新用户表数据失败")
		return r.Error
	}
	return nil
}

// 评论
type Comment struct {
	Id         int
	PostId     int
	Content    string
	State      int        `gorm:"type:int"`
	CreateTime *time.Time `gorm:"type:datetime"`
}
type CommentPage struct {
	Count int64
	List  []Comment
}

// Comment删除钩子，更新文章的评论数
func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	r := tx.Model(&Post{}).Where(&Post{Id: c.PostId}).Update("rate_num", gorm.Expr("rate_num - ？", 1))
	if r.Error != nil {
		return r.Error
	}
	return nil
}

// Comment更新钩子
func (c *Comment) AfterUpdate(tx *gorm.DB) (err error) {
	if c.State == 0 {
		r := tx.Model(&Post{}).Where(&Post{Id: c.PostId}).Update("rate_num", gorm.Expr("rate_num - ？", 1))
		if r.Error != nil {
			return r.Error
		}
	}
	return nil
}

// 创建自动新增
func (c *Comment) AfterCreate(tx *gorm.DB) (err error) {
	r := tx.Model(&Post{}).Where(&Post{Id: c.PostId}).Update("rate_num", gorm.Expr("rate_num + ?", 1))
	if r.Error != nil {
		return r.Error
	}
	return nil
}
func TestGorm() {
	url := "root:12345678@tcp(127.0.0.1:3306)/test_gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	sqlDB, _ := db.DB()
	//设置空闲连接池中的最大连接数:cite[4]:cite[6]:cite[8]
	sqlDB.SetMaxIdleConns(8)
	//设置数据库的最大打开连接数:cite[4]:cite[6]:cite[8]
	sqlDB.SetMaxOpenConns(1000)
	//设置连接可复用的最大时间:cite[4]:cite[8]
	sqlDB.SetConnMaxLifetime(time.Hour)
	//设置连接最大空闲时间:cite[8]
	sqlDB.SetConnMaxIdleTime(time.Minute * 30)

	autoGenTable(db)
	//生成从测试数据
	//genRandomTestData(db)
	//查询用户1 的文章，分页
	postPage, _ := selectPostByUserId(db, 1, 1, 10)
	fmt.Println("用户1下面的文章：", postPage)
	if len(postPage.List) > 0 {
		//文章下的评论
		commentPage, _ := selectCommentByPostId(db, postPage.List[0].Id, 1, 10)
		fmt.Println("评论数据：", commentPage)
	}
	//热门文章列表
	postList, _ := selectHotTopNPost(db, 5)
	fmt.Println("热门前5的文章列表：", postList)
	for _, v := range *postList {
		fmt.Println(*v.Title)
	}
}

func autoGenTable(db *gorm.DB) {
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Post{})
	db.AutoMigrate(&Comment{})
}
func caculateOffset(pageIndex *int, pageSize *int) int {
	if *pageIndex <= 0 {
		*pageIndex = 1
	}
	if *pageSize <= 0 {
		*pageSize = 10
	}
	var offsetNum = (*pageIndex - 1) * (*pageSize)
	return offsetNum
}

// 插入测试数据
func genRandomTestData(db *gorm.DB) {
	//插入用户数据
	users := make([]User, 0)
	fmt.Println("开始插入测试用户")
	for i := 0; i < 10; i++ {
		n := "测试用户" + string(strconv.Itoa(rand.Intn(100)))
		t := time.Now()
		users = append(users, User{Name: n, PostNum: 0, CreateTime: &t})
	}
	r := db.CreateInBatches(&users, 100)
	if r.Error != nil {
		fmt.Println(r.Error)
		return
	}
	fmt.Println("测试用户插入完成，共计:", len(users))
	//插入文章数据
	postList := make([]Post, 0)
	for i := 0; i < len(users); i++ {
		user := users[i]
		title := "《Web3之路》" + string(strconv.Itoa(rand.Intn(1000)))
		content := "内容都是一样的，大家学习路线都是一样的，准备好了吗？" + string(strconv.Itoa(rand.Intn(1000000)))
		t := time.Now()
		postList = append(postList, Post{UserId: user.Id, Title: &title, Content: &content, CreateTime: &t, UpdateTime: nil})
	}
	fmt.Println("开始插入用户文章")
	r = db.CreateInBatches(&postList, 100)
	if r.Error != nil {
		fmt.Println(r.Error)
		return
	}
	fmt.Println("文章插入完成,共计：", len(postList))
	fmt.Println("开始插入文章评论")
	//插入评论数据
	comments := make([]Comment, 0)
	for i := 0; i < len(postList); i++ {
		//奇数才生成
		if i%2 == 0 {
			post := postList[i]
			t := time.Now()
			comments = append(comments, Comment{PostId: post.Id, Content: "写的真好，你一定是大佬吧？", State: 1, CreateTime: &t})
		}
	}
	r = db.CreateInBatches(&comments, 100)
	if r.Error != nil {
		fmt.Println(r.Error)
		return
	}
	fmt.Println("文章评论插入完成，共计：", len(comments))
}

// 用户查询文章
func selectPostByUserId(db *gorm.DB, userId int, pageIndex int, pageSize int) (*PostPage, error) {
	var posts []Post
	var offsetNum = caculateOffset(&pageIndex, &pageSize)
	var count int64
	db.Model(&Post{}).Where(&Post{UserId: userId}).Count(&count)
	db.Debug().Model(&Post{}).Where(&Post{UserId: userId}).Offset(offsetNum).Limit(pageSize).Find(&posts)
	return &PostPage{Count: count, List: posts}, nil
}

// 根据文章查询评论
func selectCommentByPostId(db *gorm.DB, postId int, pageIndex int, pageSize int) (*CommentPage, error) {
	var offsetNum = caculateOffset(&pageIndex, &pageSize)
	var count int64
	var comments []Comment
	db.Model(&Comment{}).Where(&Comment{PostId: postId}).Count(&count)
	db.Model(&Comment{}).Where(&Comment{PostId: postId}).Offset(offsetNum).Limit(pageSize).Find(&comments)
	return &CommentPage{Count: count, List: comments}, nil
}

// 此方式效率极其低下
func selectHotTopNPost(db *gorm.DB, topN int) (*[]Post, error) {
	var list = make([]Post, 0)
	db.Debug().Model(&Post{}).Raw("select * from ("+
		"select pt.id, pt.title, count(ct.id) as rate_num, pt.create_time, pt.update_time "+
		"from posts pt "+
		"LEFT JOIN comments ct on ct.post_id=pt.id and ct.state=1 "+
		"where pt.raw_status=1 "+
		"GROUP BY pt.id "+
		") t "+
		"order by t.rate_num desc "+
		"limit ?", topN).
		Scan(&list)
	return &list, nil
}
