package proc

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"strings"
)

type AccessToken struct {
	Token string `json:"access_token"`
	Expire int64 `json:"expire_in"`
}

const (
	appId = "wx2e85b3f5a5c2ca65"
	appSecurt = "2bb232c2eb6b0e1e8cbae7f90b2d5152"

)
var(
	token = "gE4FXRKcCr1YBm0AmvtR954c-j2ab_OD8mjM4m7mqTKGQGuAcV5oRchSql11ZcySGPI_pXKyJZgtGWVBWnf3N0yDw_wfQLfvjV_kWtf7M63gj2KHXLmzpp3Pqmyd1aPGELHgAJAWDZ"
)

//func init() {
//	tokenUrl := "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=wx2e85b3f5a5c2ca65&secret=2bb232c2eb6b0e1e8cbae7f90b2d5152"
//	resp,err := http.Get(tokenUrl)
//	if err != nil {
//		fmt.Println(err)
//	}
//	data,err := ioutil.ReadAll(resp.Body)
//	defer resp.Body.Close()
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	err = json.Unmarshal(data,&token)
//	if err != nil {
//		fmt.Println(err)
//	}
//}



func CreateMenu(postdata, token string){
	postUrl := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/menu/create?access_token=%s",token)

	resp,err := http.Post(postUrl,"application/json",strings.NewReader(postdata))
	if err != nil {
		fmt.Println(err)
	}
	data,err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.StatusCode)
	fmt.Println("data",string(data))
}

func Menu(){
	postJson := `
	{
    "button": [
        {
            "name": "扫码",
            "sub_button": [
                {
                    "type": "scancode_waitmsg",
                    "name": "扫码带提示",
                    "key": "rselfmenu_0_0",
                    "sub_button": [ ]
                },
                {
                    "type": "scancode_push",
                    "name": "扫码推事件",
                    "key": "rselfmenu_0_1",
                    "sub_button": [ ]
                }
            ]
        },
        {
            "name": "发图",
            "sub_button": [
                {
                    "type": "pic_sysphoto",
                    "name": "系统拍照发图",
                    "key": "rselfmenu_1_0",
                   "sub_button": [ ]
                 },
                {
                    "type": "pic_photo_or_album",
                    "name": "拍照或者相册发图",
                    "key": "rselfmenu_1_1",
                    "sub_button": [ ]
                },
                {
                    "type": "pic_weixin",
                    "name": "微信相册发图",
                    "key": "rselfmenu_1_2",
                    "sub_button": [ ]
                }
            ]
        },
        {
            "name": "发送位置",
            "type": "location_select",
            "key": "rselfmenu_2_0"
        }
    ]
}
	`
	CreateMenu(postJson,token)
}

