package controllers
// package controllers
	// gorm.Model

import (
	// "github.com/astaxie/beego/httplib"
	_"path/filepath"
  _"github.com/astaxie/beego/config"

  "os"
  "time"
	"fmt"
	"strconv"
	"strings"
  "encoding/json"
)

var cached_datas0 map[string]interface{}
var cached_datas map[string]string

// cached_datas["access_token"] = ""

func init() {
	cached_datas = make(map[string]string)
	// 每隔六小时，清空一次缓存数据，避免天气数据不准确
	go func() {
     for {
        fmt.Println("1===", cached_datas)
        fmt.Println("===================reset cached=======================")
        cached_datas = make(map[string]string)
        time.Sleep(time.Second * 60 * 60 * 6)
        fmt.Println("2===", cached_datas)
     }
  }()
}



func GetServerTokenid (key, value string) (map[string]interface{}){
// func GetServerTokenid (key, value string) (string){
	return cached_datas0
}

// func GetCachedData (key, value string) (map[string]interface{}){
func GetCachedData (key string) (string){
  // cached_datas = make(map[string]string)
  cached_inner, exits := cached_datas[key]

	// fmt.Println("cached_datas========", cached_inner)
	if(exits){
		fmt.Println("pass====","pass")
		return cached_inner
	}else{
		return ""
	}
}


// func SetCachedData (key1, value string) (map[string]interface{}){
func SetCachedData (key , value string){
		fmt.Println("set cached data===============")
		// fmt.Println("MapToJson(cached_datas)===", MapToJson(cached_datas))
		//convert_str,_ := strconv.Unquote(value)
		cached_datas[key] = value
		// fmt.Println("fromUnicodeToZn(value)===", convert_str)
		WriteFile("./.cached.json", MapToJson(cached_datas))
		// fmt.Println("cached_datas=====", cached_datas)
}


func WriteFile(path, str string) {
    fmt.Println("path==", path)
    fmt.Println("str==", str[0:6])

    f, err := os.Create(path)
    if err != nil {
        fmt.Println(err)
        return
    }
    l, err := f.WriteString(str)
    if err != nil {
        fmt.Println(err)
        f.Close()
        return
    }
    fmt.Println(l, "bytes written successfully")
    err = f.Close()
    if err != nil {
        fmt.Println(err)
        return
    }
}

// Convert json string to map
func JsonToMap(jsonStr string) (map[string]string, error) {
	m := make(map[string]string)
	err := json.Unmarshal([]byte(jsonStr), &m)
	if err != nil {
		fmt.Printf("Unmarshal with error: %+v\n", err)
		return nil, err
	}

	for k, v := range m {
		fmt.Printf("%v: %v\n", k, v)
	}

	return m, nil
}

// Convert map json string
func MapToJson(m map[string]string) (string) {
	jsonByte, err := json.Marshal(m)
	if err != nil {
		fmt.Printf("Marshal with error: %+v\n", err)
		return ""
	}

	return string(jsonByte)
}


 func fromUnicodeToZn(sText string)(string) {
    textQuoted := strconv.QuoteToASCII(sText)
    textUnquoted := textQuoted[1 : len(textQuoted)-1]
    fmt.Println(textUnquoted)

    // strconv.QuoteToASCII("苏苏")     zh——Unicode
    // strconv.Unquote(s1)    Unicode——zh

    sUnicodev := strings.Split(textUnquoted, "\\u")
    var context string
    for _, v := range sUnicodev {
        if len(v) < 1 {
            continue
        }
        temp, err := strconv.ParseInt(v, 16, 32)
        if err != nil {
            panic(err)
        }
        context += fmt.Sprintf("%c", temp)
    }
    fmt.Println(context)
    return context
}
