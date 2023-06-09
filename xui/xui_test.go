package xui

import (
	"fmt"
	"log"
	"testing"
	"xui-sub-sync/config"
)

func TestLoginForCookie(t *testing.T) {
	// Define the form data
	globalConfig := config.GlobalConfig
	username := globalConfig.Servers.NodeDetail[0].Username
	password := globalConfig.Servers.NodeDetail[0].Password
	loginUrl := globalConfig.Servers.NodeDetail[0].LoginUrl
	//baseUrl := globalConfig.Servers.NodeDetail[0].BaseUrl

	cookie := LoginForCookie(loginUrl, username, password)
	log.Println(cookie)

}

func TestLoginAllNodeCookies(t *testing.T) {
	globalConfig := config.GlobalConfig
	cookies := LoginAllNodeCookies(globalConfig)
	fmt.Println(cookies)

}

func TestGetServerNodeList(t *testing.T) {
	globalConfig := config.GlobalConfig
	globalConfig = LoginAllNodeCookies(globalConfig)

	url := globalConfig.Servers.NodeDetail[0].BaseUrl
	cookie := globalConfig.Servers.NodeDetail[0].Cookie

	list, err := GetServerNodeList(url, cookie)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v\n", list)

}

func TestGetBaseUrlFromUrl(t *testing.T) {
	url := GetBaseUrlFromUrl("https://xui.131433.xyz/admin/xui/inbounds")
	fmt.Println(url)

}
