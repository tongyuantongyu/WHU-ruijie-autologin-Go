package main

import (
    "flag"
    "fmt"
    "github.com/mozillazg/request"
    "net/http"
    "os"
    "strings"
    "time"
)

var (
    help        bool
    id          string
    password    string
    nettype     int
    networkname string
)

func init() {
    flag.BoolVar(&help, "h", false, "Show help")
    flag.StringVar(&id, "i", "", "Your student id")
    flag.StringVar(&password, "p", "", "Your account password")
    flag.IntVar(&nettype, "t", 0, "Network type. 0: Cernet; 1:Telecom; 2: Mobile")
}

func usage() {
    _, _ = fmt.Fprintf(os.Stderr, `Ruijie Autologin
Usage: login -i account -p password [-t networktype]

Options:
`)
    flag.PrintDefaults()
}

func main() {
    flag.Parse()
    
    if help {
        flag.Usage()
        return
    }
    
    switch nettype {
    case 0:
        networkname = "Internet"
    case 1:
        networkname = "dianxin"
    case 2:
        networkname = "yidong"
    }
    
    c := &http.Client{}
    timeout := time.Duration(10 * time.Second)
    req := request.NewRequest(c)
    req.Client.Timeout = timeout
    resp1, err := req.Get("http://123.123.123.123")
    if err != nil {
        fmt.Println("No need to login")
        return
    }
    defer resp1.Body.Close()
    restext, _ := resp1.Text()
    cookieUrl := restext[32 : len(restext)-12]
    resp2, _ := req.Get(cookieUrl)
    defer resp2.Body.Close()
    req.Cookies = map[string]string{
        "EPORTAL_COOKIE_USERNAME":     id,
        "EPORTAL_COOKIE_PASSWORD":     password,
        "EPORTAL_COOKIE_DOMAIN":       "false",
        "EPORTAL_COOKIE_SAVEPASSWORD": "true",
        "EPORTAL_AUTO_LAND":           "",
        "EPORTAL_COOKIE_OPERATORPWD":  "",
        "EPORTAL_COOKIE_SERVER":       networkname,
        "EPORTAL_COOKIE_SERVER_NAME":  "%E7%A7%BB%E5%8A%A8",
        "EPORTAL_USER_GROUP":          "%E5%AD%A6%E7%94%9F%E7%BB%84"}
    qString := strings.Replace(strings.Replace(cookieUrl[41:], "&", "%26", -1), "=", "%3D", -1)
    req.Data = map[string]string{
        "userId":          id,
        "password":        password,
        "service":         networkname,
        "queryString":     qString,
        "operatorPwd":     "",
        "operatorUserId":  "",
        "validcode":       "",
        "passwordEncrypt": "false"}
    resp, _ := req.Post("http://172.19.1.9:8080/eportal/InterFace.do?method=login")
    respText, _ := resp.Text()
    fmt.Println("Got login reply:")
    fmt.Println(respText)
}
