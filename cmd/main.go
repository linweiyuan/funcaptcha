package main

import (
	"fmt"

	"github.com/linweiyuan/funcaptcha"
)

func main() {
	token, _ := funcaptcha.GetOpenAIToken()
	fmt.Println(token)

	token, _ = funcaptcha.GetOpenAITokenWithBx(`[{"key":"enhanced_fp","value":[{"key":"navigator_battery_charging","value":true}]},{"key":"fe","value":["DNT:1","L:zh-CN","D:24","PR:1","S:1920,1080","AS:1920,1080","TO:-480","SS:true","LS:true","IDB:true","B:false","ODB:true","CPUC:unknown","PK:Linux x86_64","CFP:1186665521","FR:false","FOS:false","FB:false","JSF:Arial,Courier,Courier New,Helvetica,Times,Times New Roman","P:Chrome PDF Viewer,Chromium PDF Viewer,Microsoft Edge PDF Viewer,PDF Viewer,WebKit built-in PDF","T:0,false,false","H:16","SWF:false"]}]`)
	fmt.Println(token)
}
