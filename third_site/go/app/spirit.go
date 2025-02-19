package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var (
	SPIRIT_HOST = "https://www.printspirit.cn"
)

type ThirdApp struct {
	Host, Uid, Pass    string 
	Token        string
	Expired_time int64
}

func NewThirdApp(host, uid, pass string) *ThirdApp {
	return &ThirdApp{
		Host: host,
		Uid : uid, 
		Pass: pass,
	}
}

func (app *ThirdApp) getAccessToken() (string, error) {

	if app.Expired_time > time.Now().Unix() {
		return app.Token, nil
	}
	url := fmt.Sprintf("%s/api/get-access-token?userid=%v&passwd=%s", app.Host, app.Uid, app.Pass)

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	rc := make(map[string]interface{})
	err = json.NewDecoder(resp.Body).Decode(&rc)
	if err != nil {
		return "", err
	}
	if rc["rc"].(string) != "OK" {
		return "", fmt.Errorf(rc["errmsg"].(string))
	}
	app.Expired_time = time.Now().Unix() + int64(rc["expirt"].(float64))
	app.Token = rc["token"].(string)
	return app.Token, nil
}

type TpInfo struct {
	Name     string `json:"name"`
	Subclass string `json:"subclass"`
	Id       string `json:"id"`
}

func (app *ThirdApp) GetList(subclass string) (*[]TpInfo, error) {
	token, err := app.getAccessToken()
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/api/get-label-list?token=%s&subclass=%s", app.Host, token, subclass)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	rc := struct {
		Rc     string   `json:"rc"`
		Errmsg string   `json:"errmsg"`
		Data   []TpInfo `json:"data"`
	}{}
	err = json.NewDecoder(resp.Body).Decode(&rc)
	if err != nil {
		return nil, err
	}
	if rc.Rc != "OK" {
		return nil, fmt.Errorf(rc.Errmsg)
	}
	return &rc.Data, nil
}

func (app *ThirdApp) GetEditUrl(subclass, tpid string) (string, error) {
	token, err := app.getAccessToken()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/third-edit?subclass=%s&tpid=%s&token=%s", app.Host, subclass, tpid, token), nil
}
