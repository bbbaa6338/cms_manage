package panel

import (
	"fmt"
	"testing"
)

func panelTest() {
	p := NewPanel()
	err := p.Login()

	if err != nil {
		fmt.Println("登录失败: " + err.Error())
	} else {

		sites := []string{"glzsbz.com", "wkstny.com", "zml1976.com", "xiongxincailiao.com", "zgznsq.com", "hgmrjtss.com", "skgdsb.com", "knwsfwx.com", "liqingf.com", "cqyfxx.com", "lyjlnk.com", "zzqjdc.com", "ylsqgj.com", "huazhongchaxun.com", "yfby888.com", "mayibanjia365.com", "tjsjgsbxg.com", "shy5188.com", "fullerence.com", "firedreamphoto.com", "shjd-edu.com", "slpaishuiban.com", "ynhrpzs.com", "zsb018.com", "zshongx.com", "shcdcc.com", "lcmygg.com", "hdsjxsb.com", "1cy37.com", "wsroujiamo.com", "jcks888.com", "wlguolv0038.com", "yachhf.com", "tjchangronggg.com", "bccsoy.com"}

		sites = []string{"www.hyxwangshunxnyyxgst.com", "www.kslzfsa.com"}
		// // 删除
		//for _, site := range sites {
		//	err = p.DelWebsite(site)
		//	if err != nil {
		//		fmt.Println(site + "删除网站失败: " + err.Error())
		//	} else {
		//		fmt.Println(site + "删除网站成功")
		//	}
		//}

		//fmt.Println("登录成功")
		//

		for _, site := range sites {

			err = p.CreateWebsite(site)
			if err != nil {
				fmt.Println(site + "创建网站失败: " + err.Error())
				break
			} else {
				err = p.ZBolg()
				if err != nil {
					fmt.Println(site + "伪静态设置失败: " + err.Error())
					break
				} else {
					fmt.Println(site + "创建网站成功")
				}
			}
		}

		////////////////////
	}
}

func TestPanelTest(t *testing.T) {
	panelTest()
}
