package zbolg

import (
	"fmt"
	"testing"
)

func zBolgInstallTest() {

	//sites := []string{"glzsbz.com", "wkstny.com", "zml1976.com", "xiongxincailiao.com", "zgznsq.com", "hgmrjtss.com", "skgdsb.com", "knwsfwx.com", "liqingf.com", "cqyfxx.com", "lyjlnk.com", "zzqjdc.com", "ylsqgj.com", "huazhongchaxun.com", "yfby888.com", "mayibanjia365.com", "tjsjgsbxg.com", "shy5188.com", "fullerence.com", "firedreamphoto.com", "shjd-edu.com", "slpaishuiban.com", "ynhrpzs.com", "zsb018.com", "zshongx.com", "shcdcc.com", "lcmygg.com", "hdsjxsb.com", "1cy37.com", "wsroujiamo.com", "jcks888.com", "wlguolv0038.com", "yachhf.com", "tjchangronggg.com", "bccsoy.com"}
	//
	//for _, site := range sites {
	//	if site == "bccsoy.com" {
	//		continue
	//	}
	//	zb, err := NewZBolg(
	//		ZBolg{
	//			WebSite:    site,
	//			InstallUrl: "http://" + site + "/zb_install/jsj.php",
	//			LoginUrl:   "http://" + site + "/zb_system/jsj.php",
	//			Host:       "38.181.29.193",
	//			Port:       23301,
	//			//Username:   "admin",
	//			//Password:   "admin",
	//		})
	//	if err != nil {
	//		fmt.Println(site + "初始化失败: " + err.Error())
	//		panic(err)
	//	}
	//
	//	db, _ := zb.DB.DB()
	//	defer db.Close()
	//
	//	err = zb.Install()
	//	if err != nil {
	//		fmt.Println(site + "安装失败: " + err.Error())
	//		panic(err)
	//	}
	//
	//	err = zb.Login()
	//	if err != nil {
	//		fmt.Println(site + "登录失败: " + err.Error())
	//		panic(err)
	//	}
	//	fmt.Println(site + "登录成功, 安装成功")
	//
	//	// 查询
	//	//_, _, _, err = zb.QueryInformation()
	//	//if err != nil {
	//	//	panic(err)
	//	//}
	//	//
	//	//QueryInformation("abcdz.cc", "38.181.25.59", 23301)
	//
	//	err = zb.Plugin()
	//	fmt.Println(err)
	//
	//	category := "分类1&2啊啊"
	//	zb.AddCategory(category, "分类1的简介", "0")
	//
	//	err = zb.WebSetting("title", "sub_title", "copy_right")
	//	err = zb.SetSidebar()
	//
	//	//err = zb.ArticleUpdate("abcdz.cc", "1", "11",
	//	//	"豪华家用轿车之争:奥迪A6L值得入手吗?",
	//	//	`<article><p>在2024年的北京国际车展上，当众多参展商竞相展示其电动汽车和自动驾驶技术的最新成果时，一种对经典驾驶激情的呼唤在豪华车领域中回响。在这片充满未来科技感的海洋中，奥迪A6L如同一股清流，以其独特的魅力，唤醒了人们对豪华轿车原始驾驶乐趣的向往。奥迪A6L，作为豪华C级车市场的中坚力量，以一场不言而喻的试驾体验，向世人展示了它在设计美学、高科技配置、动力性能与燃油经济性方面的全面升级，以及对安全性的极致追求，再次证明了其在市场中的不可替代地位。</p><p><strong>奥迪A6L设计美学</strong></p><p>奥迪A6L的设计美学，是现代与未来的碰撞，是商务与运动的完美平衡。它以家族式的前卫设计语言，突破传统商务轿车的界限，展现了一种前所未有的科幻动感。革新性的进气格栅，与矩阵式LED头灯的巧妙融合，不仅塑造了极具辨识度的前脸，还透露出奥迪品牌对未来设计趋势的深刻洞察。车身侧面，动感的腰线与轻微溜背设计，打破了传统轿车的沉闷，而大尺寸轮毂与精致的门把手设计，无疑为A6L平添了几分豪华与运动的双重气质。尾部设计同样不凡，矩阵式LED尾灯在夜幕中熠熠生辉，搭配镀铬装饰条，将运动感与尊贵感巧妙融合，令人过目难忘。奥迪A6L的设计，不仅是对当前审美的诠释，更是对未来趋势的预示，即使时光流转，这份设计的魅力依然历久弥新。</p><div class="business-container-rtb"></div><p><strong>奥迪A6L内饰深度解析</strong></p><p style="text-align:center"><img src="{#ZC_BLOG_HOST#}zb_users/upload/2024/11/20241115192555173166995560314.jpg" alt="豪华轿车市场的新星，奥迪A6L年轻车主的挚爱" title="豪华轿车市场的新星，奥迪A6L年轻车主的挚爱" /></p><div class="business-container-rtb"></div><p>进入车内，奥迪A6L的内饰设计将科技与豪华推向了一个新的高度。双联屏中控设计，不仅引领了内饰设计的新风尚，更是将科技便利性与驾驶乐趣完美结合。全液晶仪表盘的引入，让驾驶信息一目了然，而智能车载系统的集成，则让多媒体、导航等功能触手可及，驾驶者只需指尖轻轻滑动，即可轻松掌握全局。此外，贯穿式空调出风口设计，不仅体现了设计的美感，更兼顾了实用功能，确保了车厢内温度的均匀舒适。车内装饰细节上的钢琴烤漆处理，更是将豪华氛围提升到了一个新的层次，每一次触摸，都是对品质生活的深刻理解。</p><p><strong>动力与操控的豪华典范</strong></p><p>在动力与操控方面，奥迪A6L再次彰显了豪华品牌的实力与底蕴。它提供了多样化的动力选项，包括2.0T高低功率发动机及3.0T发动机，每一种选择背后，都是对动力与效率的极致追求。无论是日常通勤还是长途旅行，A6L都能以高效的动力响应，给予驾驶者充足的信心。同时，丰富的驾驶模式选项，赋予了A6L适应各种路况的能力，无论是追求速度的激情，还是注重经济的考量，奥迪A6L都能从容应对，满足不同驾驶者的需求。</p><p><strong>豪华担当，全维度守护安全</strong></p><p>安全性能方面，奥迪A6L更是展现了豪华品牌的责任与担当。通过全面升级的安全配置，从主动安全预警系统到被动安全防护措施，无一不体现了对驾乘人员安全的全方位关怀。无论是城市拥堵的街道，还是高速疾驰的公路，A6L总能以敏锐的反应和周密的保护，为每一次出行保驾护航。</p><p>总而言之，奥迪A6L在豪华C级车市场中，凭借其在设计、科技、动力、安全等多方面的卓越表现，不仅为消费者提供了超越期待的驾驶体验，更是重新定义了豪华轿车的标准。在未来的道路上，奥迪A6L将继续引领潮流，以不断进化的姿态，陪伴每一位追求品质生活的驾驶者，开启一段又一段精彩旅程。在汽车世界多元化发展的今天，奥迪A6L无疑是那颗璀璨夺目的星，照亮着豪华轿车市场的前行之路。</p><p>#奥迪A6L #商务出行 #豪华典范 #好车推荐 #高颜值 #实力派 #家用车怎么选</p></article>`,
	//	//	"恒岳,奥迪,奥迪A6",
	//	//	"【南京恒岳奥迪】位于浦口区浦珠北路42号，买奥迪！就到中升恒岳！")
	//}

	/// 添加分类及插件及模块
	site := "tjsjgsbxg.com"

	categories := map[string]string{
		"市场行情": "实时更新奥迪A6二手车的市场价格动态和趋势分析。",
		"估值工具": "提供在线工具，帮助用户快速评估奥迪A6二手车的市场价值。",
		"车型对比": "对比不同年份和配置的奥迪A6二手车报价，方便用户挑选适合的车辆。",
		"估值指南": "详解影响奥迪A6二手车报价的关键因素，如车龄、里程、车况等。",
		"地区差异": "展示不同地区奥迪A6二手车报价的价格差异及市场特点。",
		"谈判技巧": "分享在奥迪A6二手车交易中与卖家或买家谈判报价的技巧和经验。",
		"热门车型": "介绍热门奥迪A6二手车型及其当前市场报价。",
		"历史查询": "帮助用户查询奥迪A6二手车的历史记录，避免高估或低估车辆价值。",
		"案例分析": "通过真实交易案例，解析影响奥迪A6二手车报价的实际因素及定价策略。",
		"优惠汇总": "整合各大平台和经销商的奥迪A6二手车报价优惠活动，为用户提供参考。",
	}

	zb, err := NewZBolg(
		ZBolg{
			WebSite:    site,
			InstallUrl: "http://" + site + "/zb_install/jsj.php",
			LoginUrl:   "http://" + site + "/zb_system/jsj.php",
			Host:       "38.181.29.193",
			Port:       23301,
		})
	if err != nil {
		fmt.Println(site + "初始化失败: " + err.Error())
		panic(err)
	}

	db, _ := zb.DB.DB()
	defer db.Close()

	err = zb.Login()
	if err != nil {
		fmt.Println(site + "登录失败: " + err.Error())
		panic(err)
	}

	err = zb.Plugin()
	if err != nil {
		fmt.Println(site + "插件失败: " + err.Error())
		panic(err)
	}

	err = zb.WebSetting("title", "sub_title", "copy_right")
	if err != nil {
		fmt.Println(site + "网站设置失败: " + err.Error())
		panic(err)
	}

	err = zb.SetSidebar()
	if err != nil {
		fmt.Println(site + "侧边栏设置失败: " + err.Error())
		panic(err)
	}

	for s, s2 := range categories {
		err := zb.AddCategory(s, s2, "0")
		if err != nil {
			fmt.Println(site + "添加分类失败: " + s + err.Error())
			panic(err)
		}
	}

}

func Test(t *testing.T) {
	zBolgInstallTest()
}
