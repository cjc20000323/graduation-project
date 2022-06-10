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

func typeof(v interface{}) string {
	return fmt.Sprintf("%T", v)
}

type DealGet struct {
	Sell_Id     string `json:"sell_Id"`
	Buy_id      string `json:"buy_id"`
	Resource_id string `json:"resource_id"`
}

type UserGet struct {
	Id string `json:"Id"`
}

type ResourceGet struct {
	Id string `json:"Id"`
}

type TokenGet struct {
	Id string `json:"id"`
}

type CoinRecordGet struct {
	Id string `json:"Id"`
}

type ProjectGet struct {
	Id string `json:"id"`
}

type TokenDealGet struct {
	Id string `json:"Id"`
}

// @Summary 根据id查询资源
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
func QueryResource(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(ResourceGet)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	var bodyBytes [][]byte

	bodyBytes = append(bodyBytes, []byte(body.Id))

	//调用智能合约
	resp, err := bc.ChannelQuery("queryResource", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	res := []byte("[" + string(resp.Payload[:]) + "]") //因为链码放回不是列表，query和queryresource有
	// 反序列化json
	var data []map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(res).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}

	appG.Response(http.StatusOK, "成功", data)
}

// @Summary 查询所有用户，可能用于restful外部接口
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
func QueryAllAccount(c *gin.Context) {
	appG := app.Gin{C: c}
	//调用智能合约

	var bodyBytes [][]byte
	resp, err := bc.ChannelQuery("queryAllAccount", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}

	// 反序列化json
	var data []map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "成功", data)
}

func QueryAllResource(c *gin.Context) {
	appG := app.Gin{C: c}
	//调用智能合约

	var bodyBytes [][]byte
	resp, err := bc.ChannelQuery("queryAllResource", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}

	// 反序列化json
	var data []map[string]interface{}
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
func QueryAccount(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(UserGet)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	var bodyBytes [][]byte

	bodyBytes = append(bodyBytes, []byte(body.Id))

	//调用智能合约
	resp, err := bc.ChannelQuery("queryAccount", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}

	res := []byte("[" + string(resp.Payload[:]) + "]") //因为链码放回不是列表，query和queryresource有
	// 反序列化json
	var data []map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(res).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "成功", data)
}

func QueryUpload(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(UserGet)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	var bodyBytes [][]byte

	bodyBytes = append(bodyBytes, []byte(body.Id))

	//调用智能合约
	resp, err := bc.ChannelQuery("queryUpload", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	// 反序列化json
	var data []map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "成功", data)
}

func QueryBuyResources(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(UserGet)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	var bodyBytes [][]byte

	bodyBytes = append(bodyBytes, []byte(body.Id))

	//调用智能合约
	resp, err := bc.ChannelQuery("queryBuyResources", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}

	// 反序列化json
	var data []map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "成功", data)
}

func QueryDealResource(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(ResourceGet)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	var bodyBytes [][]byte

	bodyBytes = append(bodyBytes, []byte(body.Id))
	//调用智能合约
	resp, err := bc.ChannelQuery("queryDealResource", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	// 反序列化json
	var data []map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "成功", data)
}

//成功之后不显示所有信息

func ReadToken(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(TokenGet)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	var bodyBytes [][]byte

	bodyBytes = append(bodyBytes, []byte(body.Id))

	//调用智能合约
	resp, err := bc.ChannelQuery("readToken", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}

	res := []byte("[" + string(resp.Payload[:]) + "]") //因为链码放回不是列表，query和queryresource有
	// 反序列化json
	var data []map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(res).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "成功", data)
}

func GetUserTokenList(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(UserGet)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	var bodyBytes [][]byte

	bodyBytes = append(bodyBytes, []byte(body.Id))

	//调用智能合约
	resp, err := bc.ChannelQuery("getUserTokenList", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	// 反序列化json
	var data []map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "成功", data)
}

func QueryUserBiddenToken(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(UserGet)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	var bodyBytes [][]byte

	bodyBytes = append(bodyBytes, []byte(body.Id))

	//调用智能合约
	resp, err := bc.ChannelQuery("queryUserBiddenToken", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	// 反序列化json
	var data []map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "成功", data)
}

func QueryUserResource(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(UserGet)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	var bodyBytes [][]byte

	bodyBytes = append(bodyBytes, []byte(body.Id))

	//调用智能合约
	resp, err := bc.ChannelQuery("queryUserResource", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	// 反序列化json
	var data []map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "成功", data)
}

func QueryUserSharedToken(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(UserGet)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	var bodyBytes [][]byte

	bodyBytes = append(bodyBytes, []byte(body.Id))

	//调用智能合约
	resp, err := bc.ChannelQuery("queryUserSharedToken", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	// 反序列化json
	var data []map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "成功", data)
}

func QueryUserLendToken(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(UserGet)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	var bodyBytes [][]byte

	bodyBytes = append(bodyBytes, []byte(body.Id))

	//调用智能合约
	resp, err := bc.ChannelQuery("queryUserLendToken", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	// 反序列化json
	var data []map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "成功", data)
}

func QueryTokenShare(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(TokenGet)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	var bodyBytes [][]byte

	bodyBytes = append(bodyBytes, []byte(body.Id))

	//调用智能合约
	resp, err := bc.ChannelQuery("queryTokenShare", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	// 反序列化json
	var data []map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "成功", data)
}

func QueryProject(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(ProjectGet)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	var bodyBytes [][]byte

	bodyBytes = append(bodyBytes, []byte(body.Id))

	//调用智能合约
	resp, err := bc.ChannelQuery("queryProject", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}

	res := []byte("[" + string(resp.Payload[:]) + "]") //因为链码放回不是列表，query和queryresource有
	// 反序列化json
	var data []map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(res).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "成功", data)
}

func QueryUserProject(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(UserGet)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	var bodyBytes [][]byte

	bodyBytes = append(bodyBytes, []byte(body.Id))

	//调用智能合约
	resp, err := bc.ChannelQuery("queryUserProject", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	// 反序列化json
	var data []map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "成功", data)
}

func QueryProjectResource(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(ProjectGet)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	var bodyBytes [][]byte

	bodyBytes = append(bodyBytes, []byte(body.Id))

	//调用智能合约
	resp, err := bc.ChannelQuery("queryProjectResource", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	// 反序列化json
	var data []map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "成功", data)
}

func QueryDeal(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(TokenDealGet)

	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	var bodyBytes [][]byte

	bodyBytes = append(bodyBytes, []byte(body.Id))

	//调用智能合约
	resp, err := bc.ChannelQuery("queryDeal", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	// 反序列化json
	var data []map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "成功", data)
}

func QueryAllDealSum(c *gin.Context) {
	appG := app.Gin{C: c}
	//调用智能合约

	var bodyBytes [][]byte
	resp, err := bc.ChannelQuery("queryAllDealSum", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}

	// 反序列化json
	var data float64
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "成功", data)
}

func QueryAllResourceSum(c *gin.Context) {
	appG := app.Gin{C: c}
	//调用智能合约

	var bodyBytes [][]byte
	resp, err := bc.ChannelQuery("queryAllResourceSum", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}

	// 反序列化json
	var data float64
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "成功", data)
}

func QueryAllUserSum(c *gin.Context) {
	appG := app.Gin{C: c}
	//调用智能合约

	var bodyBytes [][]byte
	resp, err := bc.ChannelQuery("queryAllUserSum", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}

	// 反序列化json
	var data float64
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "成功", data)
}

func QueryAllBlockSum(c *gin.Context) {
	appG := app.Gin{C: c}
	//调用智能合约

	var bodyBytes [][]byte
	resp, err := bc.ChannelQuery("queryAllBlockSum", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}

	// 反序列化json
	var data float64
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "成功", data)
}

func QueryBuyDeal(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(TokenDealGet)

	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	var bodyBytes [][]byte

	bodyBytes = append(bodyBytes, []byte(body.Id))

	//调用智能合约
	resp, err := bc.ChannelQuery("queryBuyDeal", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	// 反序列化json
	var data []map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "成功", data)
}

func QueryCoinRecord(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(CoinRecordGet)

	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	var bodyBytes [][]byte

	bodyBytes = append(bodyBytes, []byte(body.Id))

	//调用智能合约
	resp, err := bc.ChannelQuery("queryCoinRecord", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	// 反序列化json
	var data []map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "成功", data)
}

func QueryBlock(c *gin.Context) {
	appG := app.Gin{C: c}
	//调用智能合约

	var bodyBytes [][]byte
	resp, err := bc.ChannelQuery("queryBlock", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}

	// 反序列化json
	var data float64
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "成功", data)
}
