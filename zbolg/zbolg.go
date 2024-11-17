package zbolg

import (
	"cmsManage/utils/reqRequest"
	"cmsManage/utils/setPWD"
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"regexp"
)

const (
	zBolgUserName = "wiki888"          // 默认用户名
	zBolgPassword = "L5Zz4687UfFhKcH7" // 默认密码
)

type ZBolg struct {
	WebSite    string // 站点，如：baidu.com
	InstallUrl string // 完整安装路径，如：https://www.baidu.com/install.html
	LoginUrl   string // 管理员地址，如：https://www.baidu.com/admin

	// 用于链接数据库的操作
	Host string // IP 地址，如：0.0.0.0
	Port int    // 端口，如：8080

	// Login Info
	//Username string // 用户名
	//Password string // 密码
	Auth string // 认证信息

	AuthorCount int // 用户，zBolg 里面新增的用户多少，默认为 5 - 8 个

	Session *reqRequest.Request // 用于请求的 session

	DB         *gorm.DB // 数据库连接
	dbName     string   // 数据库名
	dbPassword string   // 数据库密码

	csrfToken string         // csrfToken, 后期数据的操作都需要这个 token
	regexTag  *regexp.Regexp // 正则表达式，用于匹配 tag

}

// NewZBolg 创建一个 ZBolg 实例，并进行检查和初始化
func NewZBolg(zBolg ZBolg) (*ZBolg, error) {

	if zBolg.WebSite == "" {
		return nil, errors.New(" WebSite is empty")
	}

	if zBolg.InstallUrl == "" {
		return nil, errors.New(" InstallUrl is empty")
	}

	if zBolg.LoginUrl == "" {
		return nil, errors.New(" AdminUrl is empty")
	}

	if zBolg.Host == "" {
		return nil, errors.New(" Host is empty")
	}

	if zBolg.Port == 0 {
		return nil, errors.New(" Port is empty")
	}
	//
	//if zBolg.Username == "" {
	//	return nil, errors.New(" Username is empty")
	//}
	//
	//if zBolg.Password == "" {
	//	return nil, errors.New(" Password is empty")
	//}

	dbName, dbPwd := setPWD.GetSiteDbname(zBolg.WebSite)
	//dsn := "your_user:your_password@tcp(127.0.0.1:3306)/your_db?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbName, dbPwd, zBolg.Host, zBolg.Port, dbName)

	// 连接 MySQL 数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, errors.New(fmt.Sprintf(" failed to connect database:%s", err.Error()))
	}
	zBolg.DB = db
	zBolg.dbName = dbName
	zBolg.dbPassword = dbPwd
	zBolg.regexTag = regexp.MustCompile(`>([^<]+)<`)

	if zBolg.Session == nil {
		// 初始化请求 session
		zBolg.Session = reqRequest.NewRequest()
	}

	return &zBolg, nil
}

// checkTableExists
//
//	@Description: 检查表是否存在
//	@receiver z
//	@param db
//	@param tableName
//	@return bool
func (z *ZBolg) checkTableExists(db *gorm.DB, tableName string) bool {
	var exists bool
	err := db.Raw("SELECT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = ?)", tableName).Scan(&exists).Error
	if err != nil {
		//log.Println("Error checking table existence:", err)
		return false
	}
	return exists
}
