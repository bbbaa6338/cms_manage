package zbolg

import (
	"cmsManage/utils/reqRequest"
	"fmt"
	"github.com/duke-git/lancet/v2/random"
	"github.com/duke-git/lancet/v2/slice"
	"github.com/duke-git/lancet/v2/strutil"
	"strings"
)

func (z *ZBolg) SetSidebar() (err error) {

	sidebar := []string{
		"calendar", "previous", "authors", "catalog", "tags", "archives",
		//"link",
	}

	//var sidebar1, sidebar2, sidebar3, sidebar4, sidebar5, sidebar6, sidebar7, sidebar8, sidebar8 string
	var data = make(map[string]string)
	for i := 0; i < 9; i++ {

		var key = "sidebar"
		if i != 0 {
			key = fmt.Sprintf("%s%d", key, i-1)
		}

		siRandom := random.RandInt(4, len(sidebar))
		var start = 0
		si := make([]string, siRandom)
		for {
			if start >= siRandom {
				break
			}
			v, _ := slice.Random(sidebar)
			if !slice.Contain(si, v) {
				si[start] = v
				start++
			}
		}

		value := strings.Join(si, "|") + "|link|"
		data[key] = value
	}

	reqUrl := strutil.Before(z.LoginUrl, z.WebSite) + z.WebSite + "/zb_system/cmd.php"
	resp, err := z.Session.Post(reqRequest.RequestOption{
		Url:  reqUrl,
		Data: data,
		Params: map[string]string{
			"act":       "SidebarSet",
			"csrfToken": z.csrfToken,
		},
	})
	if err != nil {
		return
	}

	status := resp.StatusCode
	if status != 200 {
		err = fmt.Errorf("SetSidebar error, status code is not 200, %d", status)
		return
	}
	return
}
