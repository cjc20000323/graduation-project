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
	Id       string     `json:"id"`       //关系型数据库id
	Upload   []Resource `json:"Upload"`   //上传的资源对象，json格式
	Buy      []Resource `json:"Buy"`      //购买的资源对象
	Control  []Token    `json:"Control"`  //控制的代币
	Share    []Token    `json:"Share"`    //他人分享的代币
	Lend     []Token    `json:"Lend"`     //分享给他人的代币
	Projects []Project  `json:"Projects"` //用户上传的项目
	Score    float64    `json:"Score"`    //积分
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
type Token struct {
	Id         string   `json:"id"`           //唯一标识token
	Asset      Resource `json:"asset"`        //唯一标识token所包含的资产
	NotForSale bool     `json:"not_for_sale"` //是否可卖
	Owner      User     `json:"owner"`        //所有人
	Bid        User     `json:"bid"`          //投标人
	Notes      string   `json:"notes"`        //简要描述
	Value      float64  `json:"value"`        //价值
	Time       string   `json:"time"`
	Share      []User   `json:"share"` //分享给的对象
}
type Project struct {
	Id    string     `json:"id"`    //唯一标识project
	Name  string     `json:"name"`  //项目名称
	Owner User       `json:"owner"` //所有者的id
	Use   Resource   `json:"use"`   //采用的解决方案id
	Bid   []Resource `json:"bid"`   //可考虑的所有解决方案，存储目前考虑资源的id而不是代币的id
	Time  string     `json:"time"`  //标记创建时间
}

type ResourcePost struct {
	//ObjectType string    `json:"docType"` //用于CouchDB
	Id       string `json:"id"`       //资源在关系型数据库中的id，从id中可知类型
	Hash     string `json:"Hash"`     //文件在IPFS系统中的Hash值
	Uploader string `json:"Uploader"` //标记卖方id
	Time     string `json:"Time"`     //标记上链时间
	Cost     string `json:"Cost"`     //交易需要话费的积分，可设置为零\
	GetScore string `json:"GetScore"` //交易需要话费的积分，可设置为零\
	Type     string `json:"Type"`
}

type UserPost struct {
	//ObjectType  string        `json:"docType"` //用于CouchDB
	Id     string     `json:"id"`     //关系型数据库id
	Upload []Resource `json:"Upload"` //上传的资源对象，json格式
	Buy    []Resource `json:"Buy"`    //购买的资源对象
	Score  string     `json:"Score"`  //积分
}
type TokenPost struct {
	Id         string   `json:"id"`           //唯一标识token
	Asset      string   `json:"asset"`        //唯一标识token所包含的资产
	NotForSale string   `json:"not_for_sale"` //是否可卖
	Owner      string   `json:"owner"`        //所有人
	Bid        string   `json:"bid"`          //投标人
	Notes      string   `json:"notes"`        //简要描述
	Value      string   `json:"value"`        //价值
	CreateId   string   `json:"createId"`
	Time       string   `json:"time"`
	Share      []string `json:"share"` //分享给的对象
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

type TokenDealPost struct {
	Sell_id     string `json:"Sell_id"`
	Buy_id      string `json:"Buy_id"`
	Resource_id string `json:"Resource_id"` //上传的资源对象，json格式
	Cost        string `json:"Cost"`        //积分
	Time        string `json:"Time"`        //交易完成时间（上链时间更准）
	Type        string `json:"Type"`
}

type BidPost struct {
	Id     string `json:"id"`
	Bidder string `json:"bidder"`
}

type DeletePost struct {
	Id string `json:"id"`
}

type GivePost struct {
	From    string `json:"from"`
	To      string `json:"to"`
	TokenId string `json:"tokenId"`
}

type TransferPost struct {
	TokenId string `json:"tokenId"`
	UserId  string `json:"userId"`
}

type ChangeResourcePost struct {
	UserId  string `json:"userId"`
	TokenId string `json:"tokenId"`
	NewCost string `json:"newCost"`
}

type ProjectPost struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Owner string `json:"owner"`
	Time  string `json:"time"`
	Notes string `json:"notes"`
}

type AddResourcePost struct {
	UserId    string `json:"userId"`
	TokenId   string `json:"tokenId"`
	ProjectId string `json:"projectId"`
}

type ChooseResourcePost struct {
	UserId     string `json:"userId"`
	ResourceId string `json:"resourceId"`
	ProjectId  string `json:"projectId"`
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
	if body.Id == "" || body.Hash == "" || body.Uploader == "" || body.Time == "" || body.Type == "" {
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
	bodyBytes = append(bodyBytes, []byte(body.Type))

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

func CreateToken(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(TokenPost)

	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}

	if body.Id == "" || body.NotForSale == "" || body.Time == "" || body.Owner == "" || body.Asset == "" || body.Value == "" {
		appG.Response(http.StatusBadRequest, "失败", "存在空参数")
		return
	}

	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.Id))
	bodyBytes = append(bodyBytes, []byte(body.Asset))
	bodyBytes = append(bodyBytes, []byte(body.NotForSale))
	bodyBytes = append(bodyBytes, []byte(body.Owner))
	bodyBytes = append(bodyBytes, []byte(body.Bid))
	bodyBytes = append(bodyBytes, []byte(body.Notes))
	bodyBytes = append(bodyBytes, []byte(body.Value))
	bodyBytes = append(bodyBytes, []byte(body.CreateId))
	bodyBytes = append(bodyBytes, []byte(body.Time))

	//调用智能合约
	resp, err := bc.ChannelExecute("createToken", bodyBytes)
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

func BidToken(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(BidPost)

	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}

	if body.Id == "" || body.Bidder == "" {
		appG.Response(http.StatusBadRequest, "失败", "存在空参数")
		return
	}

	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.Id))
	bodyBytes = append(bodyBytes, []byte(body.Bidder))

	//调用智能合约
	resp, err := bc.ChannelExecute("bidToken", bodyBytes)
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

func DeleteToken(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(DeletePost)

	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}

	if body.Id == "" {
		appG.Response(http.StatusBadRequest, "失败", "存在空参数")
		return
	}

	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.Id))

	//调用智能合约
	resp, err := bc.ChannelExecute("deleteToken", bodyBytes)
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

func GiveToken(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(GivePost)

	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}

	if body.From == "" || body.To == "" || body.TokenId == "" {
		appG.Response(http.StatusBadRequest, "失败", "存在空参数")
		return
	}

	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.From))
	bodyBytes = append(bodyBytes, []byte(body.To))
	bodyBytes = append(bodyBytes, []byte(body.TokenId))

	//调用智能合约
	resp, err := bc.ChannelExecute("giveToken", bodyBytes)
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

func TransferToken(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(TransferPost)

	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}

	if body.UserId == "" || body.TokenId == "" {
		appG.Response(http.StatusBadRequest, "失败", "存在空参数")
		return
	}

	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.TokenId))
	bodyBytes = append(bodyBytes, []byte(body.UserId))

	//调用智能合约
	resp, err := bc.ChannelExecute("transferToken", bodyBytes)
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

func ShareToken(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(GivePost)

	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}

	if body.From == "" || body.TokenId == "" || body.To == "" {
		appG.Response(http.StatusBadRequest, "失败", "存在空参数")
		return
	}

	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.From))
	bodyBytes = append(bodyBytes, []byte(body.To))
	bodyBytes = append(bodyBytes, []byte(body.TokenId))

	//调用智能合约
	resp, err := bc.ChannelExecute("shareToken", bodyBytes)
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

func RedeemToken(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(GivePost)

	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}

	if body.From == "" || body.TokenId == "" || body.To == "" {
		appG.Response(http.StatusBadRequest, "失败", "存在空参数")
		return
	}

	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.From))
	bodyBytes = append(bodyBytes, []byte(body.To))
	bodyBytes = append(bodyBytes, []byte(body.TokenId))

	//调用智能合约
	resp, err := bc.ChannelExecute("redeemToken", bodyBytes)
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

func ChangeResourceScore(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(ChangeResourcePost)

	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}

	if body.UserId == "" || body.TokenId == "" || body.NewCost == "" {
		appG.Response(http.StatusBadRequest, "失败", "存在空参数")
		return
	}

	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.UserId))
	bodyBytes = append(bodyBytes, []byte(body.TokenId))
	bodyBytes = append(bodyBytes, []byte(body.NewCost))

	//调用智能合约
	resp, err := bc.ChannelExecute("changeResourceScore", bodyBytes)
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

func RefuseTransferToken(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(TransferPost)

	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}

	if body.UserId == "" || body.TokenId == "" {
		appG.Response(http.StatusBadRequest, "失败", "存在空参数")
		return
	}

	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.TokenId))
	bodyBytes = append(bodyBytes, []byte(body.UserId))

	//调用智能合约
	resp, err := bc.ChannelExecute("refuseTransferToken", bodyBytes)
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

func ChangeTokenSaleState(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(TransferPost)

	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}

	if body.UserId == "" || body.TokenId == "" {
		appG.Response(http.StatusBadRequest, "失败", "存在空参数")
		return
	}

	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.UserId))
	bodyBytes = append(bodyBytes, []byte(body.TokenId))

	//调用智能合约
	resp, err := bc.ChannelExecute("changeTokenSaleState", bodyBytes)
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

func CreateProject(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(ProjectPost)

	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}

	if body.Id == "" || body.Name == "" || body.Time == "" || body.Owner == "" {
		appG.Response(http.StatusBadRequest, "失败", "存在空参数")
		return
	}

	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.Id))
	bodyBytes = append(bodyBytes, []byte(body.Name))
	bodyBytes = append(bodyBytes, []byte(body.Owner))
	bodyBytes = append(bodyBytes, []byte(body.Time))
	bodyBytes = append(bodyBytes, []byte(body.Notes))

	//调用智能合约
	resp, err := bc.ChannelExecute("createProject", bodyBytes)
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

func AddResourceToProject(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(AddResourcePost)

	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}

	if body.ProjectId == "" || body.TokenId == "" || body.UserId == "" {
		appG.Response(http.StatusBadRequest, "失败", "存在空参数")
		return
	}

	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.UserId))
	bodyBytes = append(bodyBytes, []byte(body.TokenId))
	bodyBytes = append(bodyBytes, []byte(body.ProjectId))

	//调用智能合约
	resp, err := bc.ChannelExecute("addResourceToProject", bodyBytes)
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

func ChooseResourceForProject(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(ChooseResourcePost)

	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}

	if body.ProjectId == "" || body.ResourceId == "" || body.UserId == "" {
		appG.Response(http.StatusBadRequest, "失败", "存在空参数")
		return
	}

	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.UserId))
	bodyBytes = append(bodyBytes, []byte(body.ResourceId))
	bodyBytes = append(bodyBytes, []byte(body.ProjectId))

	//调用智能合约
	resp, err := bc.ChannelExecute("chooseResourceForProject", bodyBytes)
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

func RecordDeal(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(TokenDealPost)

	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}

	if body.Sell_id == "" || body.Buy_id == "" || body.Resource_id == "" || body.Cost == "" || body.Type == "" {
		appG.Response(http.StatusBadRequest, "失败", "存在空参数")
		return
	}

	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.Sell_id))
	bodyBytes = append(bodyBytes, []byte(body.Buy_id))
	bodyBytes = append(bodyBytes, []byte(body.Resource_id))
	bodyBytes = append(bodyBytes, []byte(body.Cost))
	bodyBytes = append(bodyBytes, []byte(body.Type))

	//调用智能合约
	resp, err := bc.ChannelExecute("recordDeal", bodyBytes)
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
