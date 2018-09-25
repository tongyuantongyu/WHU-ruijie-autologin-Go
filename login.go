package main

import (
        "fmt"
        "strings"
        "io/ioutil"
        "net/http"
        "github.com/mozillazg/request"
       )

func main() {
   c := &http.Client{}
   req := request.NewRequest(c)
   resp1,_ := req.Get("http://123.123.123.123")
   defer resp1.Body.Close()
   restext,_ := resp1.Text()
   cookieUrl := restext[32:len(restext)-12]
   resp2,_ := req.Get(cookieUrl)
   defer resp2.Body.Close()
   file,_ := ioutil.ReadFile("account.txt")
   read := string(file)
   para := strings.Split(read, "\r\n")
   req.Cookies = map[string]string{
    "EPORTAL_COOKIE_USERNAME": para[0],
    "EPORTAL_COOKIE_PASSWORD": para[1],
    "EPORTAL_COOKIE_DOMAIN": "false",
    "EPORTAL_COOKIE_SAVEPASSWORD": "true",
    "EPORTAL_AUTO_LAND": "",
    "EPORTAL_COOKIE_OPERATORPWD": "",
    "EPORTAL_COOKIE_SERVER": para[2],
    "EPORTAL_COOKIE_SERVER_NAME": "%E7%A7%BB%E5%8A%A8",
    "EPORTAL_USER_GROUP": "%E5%AD%A6%E7%94%9F%E7%BB%84"}
   qString := strings.Replace(strings.Replace(cookieUrl[41:], "&", "%26", -1), "=", "%3D", -1)
   req.Data = map[string]string{
    "userId": para[0],
    "password": para[1],
    "service": para[2],
    "queryString": qString,
    "operatorPwd": "",
    "operatorUserId": "",
    "validcode": "",
    "passwordEncrypt": "false"}
   resp,_ := req.Post("http://172.19.1.9:8080/eportal/InterFace.do?method=login")
   respText,_ := resp.Text()
   fmt.Println(respText)
}