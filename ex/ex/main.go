package main

import (
	"fmt"
	"strings"
)

func main() {

	siteString := `shcdcc.com
ylsqgj.com
1cy37.com
shjd-edu.com
zsb018.com
bccsoy.com
mayibanjia365.com
wlguolv0038.com
fullerence.com
wkstny.com
knwsfwx.com
shy5188.com
www.kslzfsa.com
ynhrpzs.com
zzqjdc.com
firedreamphoto.com
zgznsq.com
huazhongchaxun.com
xzxlxh.com
yachhf.com
zml1976.com
tjsjgsbxg.com
hgmrjtss.com
hdsjxsb.com
slpaishuiban.com
zshongx.com
wsroujiamo.com
glzsbz.com
jcks888.com
yfby888.com
liqingf.com
lyjlnk.com
skgdsb.com
xiongxincailiao.com
cqyfxx.com
gdlingji.com
www.hyxwangshunxnyyxgst.com
lcmygg.com
tjchangronggg.com`

	sites := strings.Split(siteString, "\n")
	for _, site := range sites {
		str := ""

		str = fmt.Sprintf(`cp /opt/zbolgcms/asfd.zip /opt/1panel/apps/openresty/openresty/www/sites/%s/index/ && unzip -o /opt/1panel/apps/openresty/openresty/www/sites/%s/index/asfd.zip -d /opt/1panel/apps/openresty/openresty/www/sites/%s/index/ && rm /opt/1panel/apps/openresty/openresty/www/sites/%s/index/asfd.zip && chown -R 1000:1000 /opt/1panel/apps/openresty/openresty/www/sites/%s/index/
`, site, site, site, site, site)
		//
		//// 登录地址
		//str = "http://" + site + "/zb_system/jsj.php"
		//
		//// 应用中心地址
		//str = "http://" + site + "/zb_users/plugin/AppCentre/client.php"
		//
		//// 分类管理
		//str = "http://" + site + "/zb_system/admin/index.php?act=CategoryMng"
		//
		//// 蜘蛛查看
		str = "http://" + site + "/zb_users/plugin/cat_spider/main.php?act=xi"

		// 网站设置
		//str = "http://" + site + "/zb_system/admin/index.php?act=SettingMng"

		// 站内
		//str = "https://zhannei.baidu.com/cse/site?q=" + site + "%2F&click=1&s=&nsid="
		fmt.Println(str)
	}
}
