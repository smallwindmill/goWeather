package controllers

import (
  // "time"
  "fmt"
  // "strconv"
  // "mindpluswebserver/models"
  "github.com/astaxie/beego/httplib"
  "github.com/astaxie/beego"

  "encoding/json"
)

type WeatherController struct {
  beego.Controller
}


type weatherData struct {
  Speech  string `json:"speech"`
  Format  string `json:"format,omitempty"`
  // Buffer  map[string]string
  // Speech  []byte `json:"speech,omitempty"`
  Token  string `json:"token,omitempty"`
  Rate int `json:"rate"`
  Channel int `json:"channel"`
  Len int `json:"len"`
  Cuid string `json:"cuid"`
}


// 获取存储数据的集合
/*func (c *WeatherController) GetCachedData() {
  // tokenid := models.GetServerTokenid()
  // fmt.Println("token_id_global======", tokenid)
  // fmt.Println("token_id_global======", token_id_global)
  // fmt.Println("token_id_global---", Token_id)
  // c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", c.Ctx.Request.Header.Get("Origin"))
  key1 := c.Input().Get("key1")
  key2 := c.Input().Get("key2")

  c.Data["json"] = GetServerTokenid(key1, key2)
  c.ServeJSON()
  // c.Ctx.WriteJson(token_id)
  // 获取前端传递的参数
}*/




// 天气
func (c *WeatherController) Weather() {
  // c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", c.Ctx.Request.Header.Get("Origin"))

  city := c.Input().Get("cityid")
  cityname := c.Input().Get("cityname")
  url := GetConfig("weather")

  appid  := c.Input().Get("appid")
  //? c.Input().Get("appid") : "29549952"
  appsecret := c.Input().Get("appsecret")
  //? c.Input().Get("appsecret") : "2WngPSdn"

  // var token_arr map[string]string
  // 读取参数

  cached_id := GetCachedData(city)
  cached_name := GetCachedData(cityname)
  fmt.Println("cached_id===", cached_id, cached_name)

  if cached_id != "" {
    c.Ctx.WriteString(cached_id)
    fmt.Println("using cache")
    return
  }

  if cached_name != "" {
    c.Ctx.WriteString(cached_name)
    fmt.Println("using cache")
    return
  }

  // cached_name
  fmt.Println("url===", url)
  req_twice := httplib.Get(url)
  // 账号和秘钥在http://www.tianqiapi.com/?action=v1 申请获取
  req_twice.Param("appid", appid)
  req_twice.Param("appsecret", appsecret)
  req_twice.Param("cityid", city)
  req_twice.Param("city", cityname)
  req_twice.Param("version", "v6")
  req_twice.Param("vue", "1")
  str, err := req_twice.String()

  // SetCachedData(city || cityname, str)
  if err != nil {
    fmt.Println("weather error== ",err)
    c.Ctx.WriteString("500")
    return
  }


  var coll map[string]interface{}
  if err := json.Unmarshal([]byte(str), &coll); err == nil {
      fmt.Println("==============json str 转map=======================")
      // fmt.Println(coll)
      if city != ""{
        SetCachedData(city, str)
      }else{
        // 当请求为城市名，根据返回的城市ID再保存一次数据，避免下一次相同的请求
        SetCachedData(cityname, str)
        // fmt.Println("cached id=====",coll["cityid"][0:6])
        id_inner := coll["cityid"]
        // SetCachedData(id_inner, str)
        fmt.Println("city_id======", id_inner)
      }
  }
  c.Ctx.WriteString(str)
}


