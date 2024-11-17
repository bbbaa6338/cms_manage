package setPWD

import (
	"fmt"
	"github.com/duke-git/lancet/v2/cryptor"
	"github.com/duke-git/lancet/v2/strutil"
	"strings"
)

const (
	zBolgEncrypt = "zBolgEncryptBolg" // 加密 zBolg 的数据库的 key
)

// GetSiteDbname
//
//	@Description: 获取域名的数据库名字及 en
//	@receiver receiver
//	@param site
//	@return string
//	@return string
func GetSiteDbname(site string) (string, string) {
	if strings.Contains(site, "://") {
		site = strings.Split(site, "://")[1]
	}
	site = strutil.RemoveWhiteSpace(site, true)
	dbName := strings.ReplaceAll(site, ".", "_")
	dbName = strings.ReplaceAll(dbName, "-", "_")

	encrypted := cryptor.AesEcbEncrypt([]byte(dbName), []byte(zBolgEncrypt))
	encStr := fmt.Sprintf("%x", encrypted)
	encStr16 := encStr[:16]

	return dbName, encStr16
}
