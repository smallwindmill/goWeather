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
  cityname := c.Input().Get("city")
  url := GetConfig("weather")
  if cityname == "" {
    cityname = c.Input().Get("cityname")
  }


  appid := c.Input().Get("appid")
  appsecret := c.Input().Get("appsecret")
  user := c.Input().Get("user")

  if(appid == "" && user == "df"){
    appid = "26385774"
    // 36135261    I0cJXOu8
  }

  if(appsecret == "" && user == "df"){
    appsecret = "IWPdTbX4"
  }


  if(city == "" && cityname == ""){
    c.Ctx.WriteString("{\"errcode\":404,\"errmsg\":\"城市不能为空\"}")
    return
  }

  // var token_arr map[string]string
  // 读取参数

  cached_id := GetCachedData(city)
  cached_name := GetCachedData(cityname)
  // fmt.Println("cached_id===", cached_id, cached_name)

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
  fmt.Println("url===", url, appid, city, cityname)
  req_twice := httplib.Get(url)
  // 账号和秘钥在http://www.tianqiapi.com/?action=v1 申请获取
  req_twice.Param("appid", appid)
  req_twice.Param("appsecret", appsecret)
  req_twice.Param("cityid", city)
  req_twice.Param("city", cityname)
  req_twice.Param("version", "v6")
  req_twice.Param("vue", "1")
  str, err := req_twice.String() //发送请求
  fmt.Println("cached name=====",appid, appsecret)

  // SetCachedData(city || cityname, str)
  if err != nil {
    fmt.Println("weather error== ",err)
    c.Ctx.ResponseWriter.WriteHeader(500)
    return
  }


  var coll map[string]interface{}
  if err := json.Unmarshal([]byte(str), &coll); err == nil {
      fmt.Println("==============json str 转map=======================")
      // fmt.Println(coll)
      name_inner := coll["city"]
      id_inner := coll["cityid"]

      // 当请求为城市名，根据返回的城市ID再保存一次数据，避免下一次相同的请求
      // 判断返回信息中是否存在错误提示，有错误提示时不缓存
      // fmt.Println("cached id=====",coll["cityid"][0:6])
      // 请求城市与返回城市不一致时，返回错误
      fmt.Println("cached id=====",name_inner, id_inner, cityname, city)

      if name_inner != nil {
        if((name_inner.(string) != cityname && cityname != "") || (id_inner.(string) != city && city != "")) {
           fmt.Println("city is wrong ")
           // c.Ctx.ResponseWriter.WriteHeader(500)
           c.Ctx.WriteString("{\"errcode\":404,\"errmsg\":\"城市"+cityname+"不存在\"}")
           return
        }
      }

      if id_inner != nil {
        SetCachedData(name_inner.(string), str)
        SetCachedData(id_inner.(string), str)  //断言，转换interface为string
      }
      fmt.Println("city_id======", id_inner)
      c.Ctx.WriteString(str)
  }
}


