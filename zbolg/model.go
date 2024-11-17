package zbolg

// ZbpConfig
// @Description: 设置里面的表
type ZbpConfig struct {
	ConfID    int    `gorm:"column:conf_ID;primaryKey;autoIncrement"`
	ConfName  string `gorm:"column:conf_Name;type:varchar(250);charset:utf8mb4;collate:utf8mb4_general_ci"`
	ConfKey   string `gorm:"column:conf_Key;type:varchar(250);charset:utf8mb4;collate:utf8mb4_general_ci"`
	ConfValue string `gorm:"column:conf_Value;type:longtext;charset:utf8mb4;collate:utf8mb4_general_ci"`
}

func (ZbpConfig) TableName() string {
	return "zbp_config"
}

type Category struct {
	CateID          int    `gorm:"column:cate_ID;primaryKey;autoIncrement"`                                              // 分类ID
	CateName        string `gorm:"column:cate_Name;type:varchar(250);charset:utf8mb4;collate:utf8mb4_general_ci"`        // 分类名称
	CateOrder       int    `gorm:"column:cate_Order;type:int(11)"`                                                       // 分类顺序
	CateType        int    `gorm:"column:cate_Type;type:int(11)"`                                                        // 分类类型
	CateCount       int    `gorm:"column:cate_Count;type:int(11)"`                                                       // 分类中的项目数
	CateAlias       string `gorm:"column:cate_Alias;type:varchar(250);charset:utf8mb4;collate:utf8mb4_general_ci"`       // 分类别名
	CateGroup       string `gorm:"column:cate_Group;type:varchar(250);charset:utf8mb4;collate:utf8mb4_general_ci"`       // 分类组
	CateIntro       string `gorm:"column:cate_Intro;type:text;charset:utf8mb4;collate:utf8mb4_general_ci"`               // 分类介绍
	CateRootID      int    `gorm:"column:cate_RootID;type:int(11)"`                                                      // 分类根ID
	CateParentID    int    `gorm:"column:cate_ParentID;type:int(11)"`                                                    // 分类父ID
	CateCreateTime  int    `gorm:"column:cate_CreateTime;type:int(11)"`                                                  // 创建时间
	CatePostTime    int    `gorm:"column:cate_PostTime;type:int(11)"`                                                    // 发布时间
	CateUpdateTime  int    `gorm:"column:cate_UpdateTime;type:int(11)"`                                                  // 更新时间
	CateTemplate    string `gorm:"column:cate_Template;type:varchar(250);charset:utf8mb4;collate:utf8mb4_general_ci"`    // 分类模板
	CateLogTemplate string `gorm:"column:cate_LogTemplate;type:varchar(250);charset:utf8mb4;collate:utf8mb4_general_ci"` // 分类日志模板
	CateMeta        string `gorm:"column:cate_Meta;type:longtext;charset:utf8mb4;collate:utf8mb4_general_ci"`            // 分类的Meta信息
}

func (Category) TableName() string {
	return "zbp_category"
}

type ZPBModule struct {
	ModID          int    `gorm:"column:mod_ID;primaryKey;autoIncrement"`                                           // 主键，自增
	ModName        string `gorm:"column:mod_Name;type:varchar(250);charset:utf8mb4;collate:utf8mb4_general_ci"`     // 模块名称
	ModFileName    string `gorm:"column:mod_FileName;type:varchar(250);charset:utf8mb4;collate:utf8mb4_general_ci"` // 文件名
	ModContent     string `gorm:"column:mod_Content;type:text;charset:utf8mb4;collate:utf8mb4_general_ci"`          // 内容
	ModSidebarID   int    `gorm:"column:mod_SidebarID;type:int(11)"`                                                // 侧边栏ID
	ModHtmlID      string `gorm:"column:mod_HtmlID;type:varchar(250);charset:utf8mb4;collate:utf8mb4_general_ci"`   // HTML ID
	ModType        string `gorm:"column:mod_Type;type:varchar(5);charset:utf8mb4;collate:utf8mb4_general_ci"`       // 类型
	ModMaxLi       int    `gorm:"column:mod_MaxLi;type:int(11)"`                                                    // 最大数量
	ModSource      string `gorm:"column:mod_Source;type:varchar(250);charset:utf8mb4;collate:utf8mb4_general_ci"`   // 来源
	ModIsHideTitle int    `gorm:"column:mod_IsHideTitle;type:tinyint(4)"`                                           // 是否隐藏标题
	ModMeta        string `gorm:"column:mod_Meta;type:longtext;charset:utf8mb4;collate:utf8mb4_general_ci"`         // Meta信息
}

func (ZPBModule) TableName() string {
	return "zbp_module"
}

type ZPBMember struct {
	MemID         int    `gorm:"column:mem_ID;primaryKey;autoIncrement"`                                           // 会员ID，主键自增
	MemGuid       string `gorm:"column:mem_Guid;type:varchar(36);charset:utf8mb4;collate:utf8mb4_general_ci"`      // GUID
	MemLevel      int    `gorm:"column:mem_Level;type:tinyint(4)"`                                                 // 会员等级
	MemStatus     int    `gorm:"column:mem_Status;type:tinyint(4)"`                                                // 会员状态
	MemName       string `gorm:"column:mem_Name;type:varchar(250);charset:utf8mb4;collate:utf8mb4_general_ci"`     // 会员姓名
	MemPassword   string `gorm:"column:mem_Password;type:varchar(250);charset:utf8mb4;collate:utf8mb4_general_ci"` // 密码
	MemEmail      string `gorm:"column:mem_Email;type:varchar(250);charset:utf8mb4;collate:utf8mb4_general_ci"`    // 邮箱
	MemHomePage   string `gorm:"column:mem_HomePage;type:varchar(250);charset:utf8mb4;collate:utf8mb4_general_ci"` // 主页
	MemIP         string `gorm:"column:mem_IP;type:varchar(250);charset:utf8mb4;collate:utf8mb4_general_ci"`       // IP地址
	MemCreateTime int    `gorm:"column:mem_CreateTime;type:int(11)"`                                               // 创建时间
	MemPostTime   int    `gorm:"column:mem_PostTime;type:int(11)"`                                                 // 发帖时间
	MemUpdateTime int    `gorm:"column:mem_UpdateTime;type:int(11)"`                                               // 更新时间
	MemAlias      string `gorm:"column:mem_Alias;type:varchar(250);charset:utf8mb4;collate:utf8mb4_general_ci"`    // 昵称
	MemIntro      string `gorm:"column:mem_Intro;type:text;charset:utf8mb4;collate:utf8mb4_general_ci"`            // 介绍
	MemArticles   int    `gorm:"column:mem_Articles;type:int(11)"`                                                 // 文章数量
	MemPages      int    `gorm:"column:mem_Pages;type:int(11)"`                                                    // 页数
	MemComments   int    `gorm:"column:mem_Comments;type:int(11)"`                                                 // 评论数
	MemUploads    int    `gorm:"column:mem_Uploads;type:int(11)"`                                                  // 上传数
	MemTemplate   string `gorm:"column:mem_Template;type:varchar(250);charset:utf8mb4;collate:utf8mb4_general_ci"` // 模板
	MemMeta       string `gorm:"column:mem_Meta;type:longtext;charset:utf8mb4;collate:utf8mb4_general_ci"`         // Meta信息
}

func (ZPBMember) TableName() string {
	return "zbp_member"
}

type ZPBPost struct {
	LogID         int    `gorm:"column:log_ID;type:int(11);primary_key;comment:'主键'"`         // 主键
	LogCateID     int    `gorm:"column:log_CateID;type:int(11);not null;comment:'分类ID'"`      // 分类ID
	LogAuthorID   int    `gorm:"column:log_AuthorID;type:int(11);not null;comment:'作者ID'"`    // 作者ID
	LogTag        string `gorm:"column:log_Tag;type:varchar(250);not null;comment:'标签'"`      // 标签
	LogStatus     int    `gorm:"column:log_Status;type:tinyint(4);not null;comment:'状态'"`     // 状态
	LogType       int    `gorm:"column:log_Type;type:int(11);not null;comment:'类型'"`          // 类型
	LogAlias      string `gorm:"column:log_Alias;type:varchar(250);not null;comment:'别名'"`    // 别名
	LogIsTop      int    `gorm:"column:log_IsTop;type:tinyint(4);not null;comment:'是否置顶'"`    // 是否置顶
	LogIsLock     int    `gorm:"column:log_IsLock;type:tinyint(4);not null;comment:'是否锁定'"`   // 是否锁定
	LogTitle      string `gorm:"column:log_Title;type:varchar(250);not null;comment:'标题'"`    // 标题
	LogIntro      string `gorm:"column:log_Intro;type:text;not null;comment:'简介'"`            // 简介
	LogContent    string `gorm:"column:log_Content;type:longtext;not null;comment:'内容'"`      // 内容
	LogCreateTime int    `gorm:"column:log_CreateTime;type:int(11);not null;comment:'创建时间'"`  // 创建时间
	LogPostTime   int    `gorm:"column:log_PostTime;type:int(11);not null;comment:'发表时间'"`    // 发表时间
	LogUpdateTime int    `gorm:"column:log_UpdateTime;type:int(11);not null;comment:'更新时间'"`  // 更新时间
	LogCommNums   int    `gorm:"column:log_CommNums;type:int(11);not null;comment:'评论数'"`     // 评论数
	LogViewNums   int    `gorm:"column:log_ViewNums;type:int(11);not null;comment:'浏览数'"`     // 浏览数
	LogTemplate   string `gorm:"column:log_Template;type:varchar(250);not null;comment:'模板'"` // 模板
	LogMeta       string `gorm:"column:log_Meta;type:longtext;not null;comment:'Meta 信息'"`    // Meta 信息
}

func (ZPBPost) TableName() string {
	return "zbp_post"
}

type ZPBTag struct {
	TagID         int    `gorm:"column:tag_ID;type:int(11);primary_key;comment:'主键'"`           // 主键
	TagName       string `gorm:"column:tag_Name;type:varchar(250);not null;comment:'标签名称'"`     // 标签名称
	TagOrder      int    `gorm:"column:tag_Order;type:int(11);not null;comment:'标签排序'"`         // 标签排序
	TagType       int    `gorm:"column:tag_Type;type:int(11);not null;comment:'标签类型'"`          // 标签类型
	TagCount      int    `gorm:"column:tag_Count;type:int(11);not null;comment:'标签使用次数'"`       // 标签使用次数
	TagAlias      string `gorm:"column:tag_Alias;type:varchar(250);not null;comment:'标签别名'"`    // 标签别名
	TagGroup      string `gorm:"column:tag_Group;type:varchar(250);not null;comment:'标签分组'"`    // 标签分组
	TagIntro      string `gorm:"column:tag_Intro;type:text;not null;comment:'标签简介'"`            // 标签简介
	TagCreateTime int    `gorm:"column:tag_CreateTime;type:int(11);not null;comment:'创建时间'"`    // 创建时间
	TagPostTime   int    `gorm:"column:tag_PostTime;type:int(11);not null;comment:'发布时间'"`      // 发布时间
	TagUpdateTime int    `gorm:"column:tag_UpdateTime;type:int(11);not null;comment:'更新时间'"`    // 更新时间
	TagTemplate   string `gorm:"column:tag_Template;type:varchar(250);not null;comment:'标签模板'"` // 标签模板
	TagMeta       string `gorm:"column:tag_Meta;type:longtext;not null;comment:'Meta信息'"`       // Meta 信息
}

func (ZPBTag) TableName() string {
	return "zbp_tag"
}
