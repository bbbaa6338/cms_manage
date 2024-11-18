package main

import "fmt"

func main() {

	sites := []string{"glzsbz.com", "wkstny.com", "zml1976.com", "xiongxincailiao.com", "zgznsq.com", "hgmrjtss.com", "skgdsb.com", "knwsfwx.com", "liqingf.com", "cqyfxx.com", "lyjlnk.com", "zzqjdc.com", "ylsqgj.com", "huazhongchaxun.com", "yfby888.com", "mayibanjia365.com", "tjsjgsbxg.com", "shy5188.com", "fullerence.com", "firedreamphoto.com", "shjd-edu.com", "slpaishuiban.com", "ynhrpzs.com", "zsb018.com", "zshongx.com", "shcdcc.com", "lcmygg.com", "hdsjxsb.com", "1cy37.com", "wsroujiamo.com", "jcks888.com", "wlguolv0038.com", "yachhf.com", "tjchangronggg.com", "bccsoy.com"}

	for _, site := range sites {
		str := ""

		str = fmt.Sprintf(`cp /opt/zbolgcms/asfd.zip /opt/1panel/apps/openresty/openresty/www/sites/%s/index/ && unzip -o /opt/1panel/apps/openresty/openresty/www/sites/%s/index/asfd.zip -d /opt/1panel/apps/openresty/openresty/www/sites/%s/index/ && rm /opt/1panel/apps/openresty/openresty/www/sites/%s/index/asfd.zip && chown -R 1000:1000 /opt/1panel/apps/openresty/openresty/www/sites/%s/index/
`, site, site, site, site, site)

		// 登录地址
		str = "http://" + site + "/zb_system/jsj.php"

		// 应用中心地址
		str = "http://" + site + "/zb_users/plugin/AppCentre/client.php"

		// 分类管理
		str = "http://" + site + "/zb_system/admin/index.php?act=CategoryMng"
		fmt.Println(str)
	}
}
