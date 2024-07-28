package app

import (
    "net/http"
    "time"
    "encoding/json"
    "fmt"
)

const (
    SPIRIT_HOST = "https://www.printspirit.cn"
    UID = "third_test"
    PASS= "third_test"
)

type AccessToken struct {
    Id string
    expired_time int64
}
var access_token = AccessToken{};

func GetAccessToken(uid, pass string ) (string, error) {
    
    if (access_token.expired_time > time.Now().Unix()) {
        return access_token.Id, nil
    }
	url := fmt.Sprintf("%s/api/get-access-token?userid=%v&passwd=%s", SPIRIT_HOST, uid, pass)

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
	if ( rc["rc"].(string)!="OK") {
	    return "", fmt.Errorf(rc["errmsg"].(string))
	}
	access_token.expired_time =  time.Now().Unix() + int64(rc["expirt"].(float64));
	access_token.Id = rc["token"].(string)
	return access_token.Id, nil
}

type TpInfo struct {
    Name string `json:"name"`
    Subclass string `json:"subclass"`
    Id string `json:"id"`
}

func GetList(subclass string) (*[]TpInfo, error) {
    token, err := GetAccessToken(UID, PASS)
    if err!=nil {
        return nil, err
    }
    url := fmt.Sprintf("%s/api/get-label-list?token=%s&subclass=%s", SPIRIT_HOST, token, subclass)
    
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	rc := struct {
	    Rc  string `json:"rc"`
	    Errmsg string `json:"errmsg"`
	    Data []TpInfo `json:"data"`
    }{}
	err = json.NewDecoder(resp.Body).Decode(&rc)
	if err != nil {
		return nil, err
	}
	if ( rc.Rc!="OK") {
	    return nil, fmt.Errorf(rc.Errmsg)
	}
	return &rc.Data, nil
}

func GetEditUrl(subclass, tpid string) (string, error){
    token, err := GetAccessToken(UID, PASS)
    if err!=nil {
        return "", err
    }
    return fmt.Sprintf("%s/third-edit?subclass=%s&tpid=%s&token=%s", SPIRIT_HOST, subclass, tpid, token), nil
}

