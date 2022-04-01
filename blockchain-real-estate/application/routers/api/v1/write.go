/**
 * @Author: 夜央 Oh oh oh oh oh oh (https://github.com/togettoyou)
 * @Email: zoujh99@qq.com
 * @Date: 2020/3/12 12:09 下午
 * @Description: 销售相关接口
 */
package v1

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	bc "github.com/togettoyou/blockchain-real-estate/application/blockchain"
	"github.com/togettoyou/blockchain-real-estate/application/pkg/app"
	"net/http"
)

type Resource struct {
	//ObjectType string    `json:"docType"` //用于CouchDB
	Id       string  `json:"id"`       //资源在关系型数据库中的id，从id中可知类型
	Hash     string  `json:"Hash"`     //文件在IPFS系统中的Hash值
	Uploader string  `json:"Uploader"` //标记卖方id
	Time     string  `json:"Time"`     //标记上链时间
	Cost     float64 `json:"Cost"`     //交易需要话费的积分，可设置为零
}

type User struct {
	//ObjectType  string        `json:"docType"` //用于CouchDB
	Id     string     `json:"id"`     //关系型数据库id
	Upload []Resource `json:"Upload"` //上传的资源对象，json格式
	Buy    []Resource `json:"Buy"`    //购买的资源对象
	Score  float64    `json:"Score"`  //积分
}
type Deal struct {
	//ObjectType  string        `json:"docType"` //用于CouchDB
	//Id          string        `json:"id"`     //联合主键  三个
	Sell_id     string  `json:"Sell_id"`
	Buy_id      string  `json:"Buy_id"`
	Resource_id string  `json:"Resource_id"` //上传的资源对象，json格式
	Cost        float64 `json:"Cost"`        //积分
	Time        string  `json:"Time"`        //交易完成时间（上链时间更准）
}

type ResourcePost struct {
	//ObjectType string    `json:"docType"` //用于CouchDB
	Id       string `json:"id"`       //资源在关系型数据库中的id，从id中可知类型
	Hash     string `json:"Hash"`     //文件在IPFS系统中的Hash值
	Uploader string `json:"Uploader"` //标记卖方id
	Time     string `json:"Time"`     //标记上链时间
	Cost     string `json:"Cost"`     //交易需要话费的积分，可设置为零\
	GetScore string `json:"GetScore"` //交易需要话费的积分，可设置为零\
}

type UserPost struct {
	//ObjectType  string        `json:"docType"` //用于CouchDB
	Id     string     `json:"id"`     //关系型数据库id
	Upload []Resource `json:"Upload"` //上传的资源对象，json格式
	Buy    []Resource `json:"Buy"`    //购买的资源对象
	Score  string     `json:"Score"`  //积分
}
type DealPost struct {
	//ObjectType  string        `json:"docType"` //用于CouchDB
	//Id          string        `json:"id"`     //联合主键  三个
	Sell_id     string `json:"Sell_id"`
	Buy_id      string `json:"Buy_id"`
	Resource_id string `json:"Resource_id"` //上传的资源对象，json格式
	Cost        string `json:"Cost"`        //积分
	Time        string `json:"Time"`        //交易完成时间（上链时间更准）
}

// @Summary 创建用户，id，初始积分
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/CreateUser [post]
func CreateUser(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(UserPost)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	if body.Id == "" || body.Score == "" {
		appG.Response(http.StatusBadRequest, "失败", "新建用户需要id和初始积分")
		return
	}

	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.Id))
	bodyBytes = append(bodyBytes, []byte(body.Score))

	//调用智能合约
	resp, err := bc.ChannelExecute("createUser", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	var data map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "成功", data)
}

// @Summary 上传资源
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/CreateUser [post]
func UploadResource(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(ResourcePost)

	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	if body.Id == "" || body.Hash == "" || body.Uploader == "" || body.Time == "" {
		appG.Response(http.StatusBadRequest, "失败", "除cost和getsocre外存在空参数")
		return
	}

	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.Id))
	bodyBytes = append(bodyBytes, []byte(body.Hash))
	bodyBytes = append(bodyBytes, []byte(body.Uploader))
	bodyBytes = append(bodyBytes, []byte(body.Time))
	bodyBytes = append(bodyBytes, []byte(body.Cost))
	bodyBytes = append(bodyBytes, []byte(body.GetScore))

	//调用智能合约
	resp, err := bc.ChannelExecute("uploadResource", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	var data map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "成功", data)
}

// @Summary 交易记录
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/CreateUser [post]
func CreateDeal(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(DealPost)

	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	if body.Sell_id == "" || body.Buy_id == "" || body.Resource_id == "" || body.Time == "" || body.Cost == "" {
		appG.Response(http.StatusBadRequest, "失败", "存在空参数")
		return
	}

	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.Sell_id))
	bodyBytes = append(bodyBytes, []byte(body.Buy_id))
	bodyBytes = append(bodyBytes, []byte(body.Resource_id))
	bodyBytes = append(bodyBytes, []byte(body.Cost))
	bodyBytes = append(bodyBytes, []byte(body.Time))

	//调用智能合约
	resp, err := bc.ChannelExecute("createDeal", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	var data map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "成功", data)
}

//成功之后不显示所有信息
