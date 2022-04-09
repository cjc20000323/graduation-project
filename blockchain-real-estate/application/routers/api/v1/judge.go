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

type UserTokenPost struct {
	UserId  string `json:"userId"`
	TokenId string `json:"tokenId"`
}

type UserResourcePost struct {
	UserId     string `json:"userId"`
	ResourceId string `json:"resourceId"`
}

type UserProjectPost struct {
	UserId    string `json:"userId"`
	ProjectId string `json:"projectId"`
}

func JudgeOwnToken(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(UserTokenPost)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	var bodyBytes [][]byte

	bodyBytes = append(bodyBytes, []byte(body.UserId))
	bodyBytes = append(bodyBytes, []byte(body.TokenId))

	//调用智能合约
	resp, err := bc.ChannelQuery("judgeOwnToken", bodyBytes)
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

func JudgeBuyResource(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(UserResourcePost)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	var bodyBytes [][]byte

	bodyBytes = append(bodyBytes, []byte(body.UserId))
	bodyBytes = append(bodyBytes, []byte(body.ResourceId))

	//调用智能合约
	resp, err := bc.ChannelQuery("judgeBuyResource", bodyBytes)
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

func JudgeOwnResource(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(UserResourcePost)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	var bodyBytes [][]byte

	bodyBytes = append(bodyBytes, []byte(body.UserId))
	bodyBytes = append(bodyBytes, []byte(body.ResourceId))

	//调用智能合约
	resp, err := bc.ChannelQuery("judgeOwnResource", bodyBytes)
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

func JudgeShareResource(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(UserResourcePost)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	var bodyBytes [][]byte

	bodyBytes = append(bodyBytes, []byte(body.UserId))
	bodyBytes = append(bodyBytes, []byte(body.ResourceId))

	//调用智能合约
	resp, err := bc.ChannelQuery("judgeShareResource", bodyBytes)
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

func JudgeShareToken(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(UserTokenPost)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	var bodyBytes [][]byte

	bodyBytes = append(bodyBytes, []byte(body.UserId))
	bodyBytes = append(bodyBytes, []byte(body.TokenId))

	//调用智能合约
	resp, err := bc.ChannelQuery("judgeShareToken", bodyBytes)
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

func JudgeLendToken(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(UserTokenPost)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	var bodyBytes [][]byte

	bodyBytes = append(bodyBytes, []byte(body.UserId))
	bodyBytes = append(bodyBytes, []byte(body.TokenId))

	//调用智能合约
	resp, err := bc.ChannelQuery("judgeLendToken", bodyBytes)
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

func JudgeOwnProject(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(UserProjectPost)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	var bodyBytes [][]byte

	bodyBytes = append(bodyBytes, []byte(body.UserId))
	bodyBytes = append(bodyBytes, []byte(body.ProjectId))

	//调用智能合约
	resp, err := bc.ChannelQuery("judgeOwnProject", bodyBytes)
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
