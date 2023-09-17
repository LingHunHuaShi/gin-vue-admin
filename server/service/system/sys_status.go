package system

import (
	"encoding/json"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
)

type SysStatusService struct {
}

func (s *SysStatusService) UploadSystemStatus(sysStatus system.SysStatus, serverUrl string) error {
	// 解析云端服务器websocket url
	u, err := url.Parse(serverUrl)
	if err != nil {
		log.Fatal("err")
		return err
	}
	log.Println("解析url成功")

	// 建立websocket链接
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal(err)
		return err
	}
	log.Println(" 成功建立websocket链接")
	defer conn.Close() //延时执行关闭websocket链接

	// 数据序列化
	statusJson, _ := json.Marshal(sysStatus)

	// json数据传递
	err = conn.WriteJSON(statusJson)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil

}
