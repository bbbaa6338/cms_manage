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
		fmt.Println("登录成功")

		err = p.CreateWebsite("123123.com")
		if err != nil {
			fmt.Println("创建网站失败: " + err.Error())
		} else {
			err = p.ZBolg()
			if err != nil {
				fmt.Println("伪静态设置失败: " + err.Error())
			} else {
				fmt.Println("创建网站成功")
			}
		}
	}
}

func TestPanelTest(t *testing.T) {
	panelTest()
}
