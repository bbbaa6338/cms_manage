package fake

import (
	"math/rand"
)

var array1 = []string{"赵", "钱", "孙", "李", "周", "吴", "郑", "王", "冯", "陈", "褚", "卫", "蒋", "沈", "韩", "杨", "朱",
	"秦", "尤", "许", "何", "吕", "施", "张", "孔", "曹", "严", "华", "金", "魏", "陶", "姜", "戚", "谢",
	"邹", "喻", "柏", "水", "窦", "章", "云", "苏", "潘", "葛", "奚", "范", "彭", "郎", "鲁", "韦", "昌",
	"马", "苗", "凤", "花", "方", "俞", "任", "袁", "柳", "酆", "鲍", "史", "唐", "费", "廉", "岑", "薛",
	"雷", "贺", "倪", "汤", "滕", "殷", "罗", "毕", "郝", "邬", "安", "常", "乐", "于", "时", "傅", "皮",
	"卞", "齐", "康", "伍", "余", "元", "卜", "顾", "孟", "平", "黄", "和", "穆", "萧", "尹", "姚", "邵",
	"堪", "汪", "祁", "毛", "禹", "狄", "米", "贝", "明", "臧", "计", "伏", "成", "戴", "谈", "宋", "茅",
	"庞", "熊", "纪", "舒", "屈", "项", "祝", "董", "梁"}

var array2 = []string{"秀", "娟", "英", "华", "慧", "巧", "美", "娜", "静", "淑", "惠", "珠", "翠", "雅", "玉", "萍", "红",
	"娥", "玲", "芬", "芳", "燕", "彩", "春", "菊", "兰", "凤", "梅", "琳", "素", "云", "莲", "真", "环",
	"雪", "荣", "爱", "妹", "霞", "香", "莺", "媛", "艳", "瑞", "凡", "佳", "嘉", "琼", "勤", "珍", "贞",
	"莉", "桂", "叶", "璧", "璐", "娅", "琦", "晶", "妍", "茜", "秋", "珊", "莎", "锦", "黛", "倩", "婷",
	"姣", "婉", "娴", "瑾", "颖", "露", "瑶", "怡", "婵", "雁", "蓓", "仪", "荷", "丹", "蓉", "眉", "君",
	"琴", "蕊", "薇", "菁", "梦", "岚", "苑", "柔", "竹", "霭", "凝", "晓", "欢", "霄", "枫", "芸", "菲",
	"寒", "欣", "滢", "伊", "亚", "宜", "可", "姬", "舒", "影", "荔", "枝", "思", "丽", "秀", "飘", "育",
	"馥", "琦", "晶", "妍", "茜", "秋", "珊", "莎", "锦", "黛", "青", "倩", "婷", "宁", "蓓", "纨", "苑",
	"婕", "馨", "瑗", "琰", "韵", "融", "园", "艺", "咏", "卿", "聪", "澜", "纯", "毓", "悦", "昭", "冰",
	"爽", "琬", "茗", "羽", "希", "伟", "刚", "勇", "毅", "俊", "峰", "强", "军", "平", "保", "东", "文",
	"辉", "明", "永", "健", "世", "广", "志", "义", "兴", "良", "海", "山", "仁", "波", "贵", "福", "生",
	"龙", "元", "全", "国", "胜", "学", "祥", "才", "发", "武", "新", "清", "飞", "彬", "富", "顺", "信",
	"子", "杰", "涛", "昌", "成", "康", "星", "天", "达", "安", "岩", "中", "茂", "进", "林", "有", "坚",
	"和", "彪", "诚", "先", "敬", "震", "振", "壮", "会", "思", "群", "豪", "心", "邦", "承", "乐", "功",
	"松", "善", "厚", "庆", "磊", "民", "友", "裕", "河", "哲", "江", "超", "亮", "政", "谦", "亨", "奇",
	"固", "之", "轮", "翰", "朗", "伯", "宏", "言", "鸣", "朋", "斌", "梁", "栋", "维", "启", "克", "伦",
	"翔", "旭", "鹏", "泽", "辰", "士", "以", "建", "家", "致", "树", "炎", "德", "行", "时", "泰", "盛"}

// GenerateCNName
//
//	@Description: 生成随机的中文名，长度为 2-3
//	@return string
func GenerateCNName() string {
	front := array1[rand.Intn(len(array1))]
	last := ""
	nameLength := rand.Intn(2) + 1 // 生成 1 或 2，表示名字部分长度为 1 或 2
	for i := 0; i < nameLength; i++ {
		last += array2[rand.Intn(len(array2))]
	}

	return front + last
}

var enArray1 = []string{
	"Anderson", "Baker", "Bell", "Bennett", "Brooks", "Brown", "Butler", "Campbell", "Carter", "Clark", "Collins", "Cook", "Cooper", "Cox", "Davis", "Diaz", "Edwards", "Evans", "Flores", "Foster", "Garcia", "Gonzalez", "Gray", "Green", "Hall", "Harris", "Hernandez", "Hill", "Howard", "Hughes", "Jackson", "James", "Jenkins", "Johnson", "Jones", "Kelly", "King", "Lee", "Lewis", "Long", "Lopez", "Martinez", "Miller", "Mitchell", "Moore", "Morgan", "Morris", "Murphy", "Nelson", "Nguyen", "Parker", "Perez", "Perry", "Peterson", "Phillips", "Powell", "Price", "Ramirez", "Reed", "Richardson", "Rivera", "Roberts", "Robinson", "Rodriguez", "Rogers", "Ross", "Russell", "Sanchez", "Sanders", "Scott", "Simmons", "Smith", "Stewart", "Sullivan", "Taylor", "Thomas", "Thompson", "Torres", "Turner", "Walker", "Ward", "Watson", "White", "Williams", "Wilson", "Wood", "Wright", "Young", "Adams", "Allen", "Bailey", "Barnes", "Bryant", "Burns", "Chapman", "Curtis", "Dixon", "Duncan", "Freeman", "Harper",
}
var enArray2 = []string{
	"Abby", "Adam", "Alec", "Alex", "Amos", "Andy", "Anna", "Aria", "Aron", "Axel", "Barb", "Beau", "Beth", "Bill", "Blair", "Brad", "Carl", "Cate", "Chad", "Clay", "Cody", "Cory", "Cruz", "Dale", "Dana", "Dawn", "Dean", "Drew", "Earl", "Eddy", "Ella", "Emil", "Emma", "Enzo", "Eric", "Erin", "Evan", "Ezra", "Finn", "Fred", "Gale", "Gary", "Gina", "Glen", "Greg", "Gwen", "Hank", "Hans", "Hugh", "Ivan", "Jack", "Jake", "Jane", "Jena", "Jess", "Jill", "Joan", "Joel", "John", "Jose", "Josh", "June", "Karl", "Kate", "Kent", "Kyle", "Lana", "Leon", "Levi", "Liam", "Lily", "Lisa", "Liza", "Lois", "Lucy", "Luke", "Mack", "Mark", "Matt", "Max", "Maya", "Meg", "Mike", "Mira", "Nate", "Neil", "Nick", "Noah", "Nora", "Owen", "Paul", "Pete", "Phil", "Quinn", "Raul", "Reed", "Rene", "Rick", "Rita", "Rob", "Rosa", "Ross", "Ruby", "Ruth", "Ryan", "Sam", "Sara", "Seth", "Shaw", "Stan", "Tate", "Tess", "Todd", "Tony", "Troy", "Troy", "Tyra", "Vera", "Vick", "Vlad", "Wade", "Walt", "Will", "Zane", "Zach", "Abel", "Alma", "Bart", "Bess", "Boyd", "Cain", "Cliff", "Cora", "Dave", "Davy", "Dora", "Doug", "Drew", "Dusty", "Eddy", "Elly", "Elmo", "Elsa", "Emma", "Evan", "Fran", "Gary", "Glen", "Greg", "Hank", "Hugo", "Ida", "Inez", "Iris", "Isla", "Jake", "Jane", "Jeff", "Jess", "Jody", "Joey", "Judy", "Karl", "Kate", "Kirk", "Kris", "Lara", "Lisa", "Lori", "Louis", "Lucy", "Luke", "Marc", "Milo", "Nina", "Noel", "Omar", "Opal", "Owen", "Raul", "Reid", "Rick", "Rolf", "Rosa", "Roy", "Rudy", "Ryan", "Saul", "Sean", "Seth", "Troy", "Vlad", "Wade", "Will", "Zara", "Zeke", "Zora", "Zuri", "Abel", "Axel",
}

// GenerateENName
//
//	@Description: 生成随机英文名，长度需要传递
//	@param count
//	@return string
func GenerateENName() string {
	// 从数组中随机选择姓和名字
	firstName := array2[rand.Intn(len(enArray2))]
	lastName := array1[rand.Intn(len(enArray1))]

	return firstName + " " + lastName
}
