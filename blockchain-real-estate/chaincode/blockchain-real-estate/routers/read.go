package routers

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/togettoyou/blockchain-real-estate/chaincode/blockchain-real-estate/lib"
	"github.com/togettoyou/blockchain-real-estate/chaincode/blockchain-real-estate/utils"
	"reflect"
	//"github.com/togettoyou/blockchain-real-estate/chaincode/blockchain-real-estate/utils"
)

//查询某用户上传资源
func QueryUpload(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var resourcelist []string
	var account lib.User
	id := args[0]
	if reflect.TypeOf(id).Name() != "string" {
		return shim.Error("Please check the type of input.")
	}
	if id == "" {
		return shim.Error("Please offer the user id.")
	}
	err := json.Unmarshal(QueryAccount(stub, args).Payload, &account)
	if err != nil {
		return shim.Error(fmt.Sprintf("用户信息验证失败%s", err))
	}

	for _, r := range account.Upload {
		//unMarshal  []byte 字符串到结构体json等
		//marshal  结构体变json byte
		resourcelist = append(resourcelist, r)
	}

	resourcelistByte, err := json.Marshal(resourcelist)
	if err != nil {
		return shim.Error(fmt.Sprintf("QueryAccountList-序列化出错: %s", err))
	}
	return shim.Success(resourcelistByte)
}

func QueryBuyResources(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var resourcelist []string
	var account lib.User
	id := args[0]
	if reflect.TypeOf(id).Name() != "string" {
		return shim.Error("Please check the type of input.")
	}
	if id == "" {
		return shim.Error("Please offer the user id.")
	}
	err := json.Unmarshal(QueryAccount(stub, args).Payload, &account)
	if err != nil {
		return shim.Error(fmt.Sprintf("用户信息验证失败%s", err))
	}

	for _, r := range account.Buy {
		//unMarshal  []byte 字符串到结构体json等
		//marshal  结构体变json byte
		resourcelist = append(resourcelist, r)
	}

	resourcelistByte, err := json.Marshal(resourcelist)
	if err != nil {
		return shim.Error(fmt.Sprintf("QueryAccountList-序列化出错: %s", err))
	}
	return shim.Success(resourcelistByte)
}

func QueryAccount(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var account lib.User //其实就一个
	id := args[0]
	if id == "" {
		return shim.Error("Please offer the user id.")
	}
	result, err := utils.GetStateByPartialCompositeKeys(stub, lib.UserKey, args)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	if result == nil {
		return shim.Error("The user does not exist.")
	}

	err = json.Unmarshal(result[0], &account)
	if err != nil {
		return shim.Error(fmt.Sprintf("QueryAccountList-反序列化出错: %s", err))
	}

	accountByte, err := json.Marshal(account)
	if err != nil {
		return shim.Error(fmt.Sprintf("QueryAccountList-序列化出错: %s", err))
	}

	return shim.Success(accountByte)
}

func QueryResource(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var resource lib.Resource //其实就一个
	id := args[0]
	if reflect.TypeOf(id).Name() != "string" {
		return shim.Error("Please check the type of input.")
	}
	if id == "" {
		return shim.Error("Please offer the resource id.")
	}
	result, err := utils.GetStateByPartialCompositeKeys(stub, lib.ResourceKey, args)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	err = json.Unmarshal(result[0], &resource)
	if err != nil {
		return shim.Error(fmt.Sprintf("QueryResource-反序列化出错: %s", err))
	}

	resourceByte, err := json.Marshal(resource)
	if err != nil {
		return shim.Error(fmt.Sprintf("QueryResource-序列化出错: %s", err))
	}
	return shim.Success(resourceByte)
}

func QueryAllAccount(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var accountList []lib.User
	results, err := utils.GetStateByPartialCompositeKeys(stub, lib.UserKey, args)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}

	for _, v := range results {
		if v != nil {
			var account lib.User
			err := json.Unmarshal(v, &account)
			if err != nil {
				return shim.Error(fmt.Sprintf("QueryAccountList-反序列化出错: %s", err))
			}
			accountList = append(accountList, account)
		}
	}

	accountListByte, err := json.Marshal(accountList)
	if err != nil {
		return shim.Error(fmt.Sprintf("QueryAccountList-序列化出错: %s", err))
	}
	return shim.Success(accountListByte)
}

func QueryAllResource(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var resourcelist []lib.Resource
	results, err := utils.GetStateByPartialCompositeKeys(stub, lib.ResourceKey, args)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}

	for _, v := range results {
		if v != nil {
			var resource lib.Resource
			err := json.Unmarshal(v, &resource)
			if err != nil {
				return shim.Error(fmt.Sprintf("Queryresourcelist-反序列化出错: %s", err))
			}
			resourcelist = append(resourcelist, resource)
		}
	}

	resourcelistByte, err := json.Marshal(resourcelist)
	if err != nil {
		return shim.Error(fmt.Sprintf("QueryAccountList-序列化出错: %s", err))
	}
	return shim.Success(resourcelistByte)
}

//func QueryDeal(stub shim.ChaincodeStubInterface, args []string) pb.Response {//查询所有记录先查所有资源再分别查交易记录，没有必要
//	results, err := utils.GetStateByPartialCompositeKeys(stub, lib.DealKey, args)
//	if err != nil {
//		return shim.Error(fmt.Sprintf("%s", err))
//	}
//
//	dealListByte, err := json.Marshal(results)
//	if err != nil {
//		return shim.Error(fmt.Sprintf("QueryAccountList-序列化出错: %s", err))
//	}
//	return shim.Success(dealListByte)
//}

func QueryDealResource(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	result, err := utils.GetHistoryForKeys(stub, lib.ResourceKey, args) //查找该资源得历史数据
	if err != nil {
		return shim.Error(fmt.Sprintf("获取历史信息错误: %s", err))
	}
	//for _, v := range result {
	//	if v != nil {
	//		var resource lib.Resource
	//		err := json.Unmarshal(v, &resource)
	//		if err != nil {
	//			return shim.Error(fmt.Sprintf("QueryDealResource-反序列化出错: %s", err))
	//		}
	//		resourceList = append(resourceList, resource)
	//	}
	//}
	//resourceListByte, err := json.Marshal(resourceList)
	//if err != nil {
	//	return shim.Error(fmt.Sprintf("QueryDealResource-序列化出错: %s", err))
	//}
	return shim.Success(result)
}

func ReadToken(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	id := args[0]
	if reflect.TypeOf(id).Name() != "string" {
		return shim.Error("Please check the type of input.")
	}
	if id == "" {
		return shim.Error("Please offer the token id.")
	}
	var token lib.Token
	result, err := utils.GetStateByPartialCompositeKeys(stub, lib.TokenKey, args)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	err = json.Unmarshal(result[0], &token)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	tokenByte, err := json.Marshal(token)
	if err != nil {
		return shim.Error(fmt.Sprintf("ReadToken-序列化出错: %s", err))
	}
	return shim.Success(tokenByte)
}

func GetUserTokenList(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	id := args[0] //传入用户的id
	if reflect.TypeOf(id).Name() != "string" {
		return shim.Error("Please check the type of input.")
	}
	if id == "" {
		return shim.Error("Please offer the user id.")
	}
	var account lib.User
	result, err := utils.GetStateByPartialCompositeKeys(stub, lib.UserKey, args)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	err = json.Unmarshal(result[0], &account)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	var tokenList []lib.Token
	for _, v := range account.Control {
		var token lib.Token
		err := json.Unmarshal(ReadToken(stub, []string{v}).Payload, &token)
		if err != nil {
			shim.Error(fmt.Sprintf("%s", err))
		}
		tokenList = append(tokenList, token)
	}
	/*for _, v := range account.Buy {
		var token lib.Token
		err := json.Unmarshal(ReadToken(stub, []string{v}).Payload, &token)
		if err != nil {
			shim.Error(fmt.Sprintf("%s", err))
		}
		tokenList = append(tokenList, token)
	}*/
	tokenListByte, err := json.Marshal(tokenList)
	if err != nil {
		return shim.Error(fmt.Sprintf("getUserTokenList-序列化出错: %s", err))
	}
	return shim.Success(tokenListByte)
}

/*func JudgeOwnToken(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	userId := args[0]
	tokenId := args[1]

	if reflect.TypeOf(tokenId).Name() != "string" || reflect.TypeOf(userId).Name() != "string" {
		return shim.Error("Please check the type of input.")
	}

	if tokenId == "" {
		return shim.Error("Please offer the token id.")
	}
	if userId == "" {
		return shim.Error("Please offer the user id.")
	}

	var user lib.User
	result, err := utils.GetStateByPartialCompositeKeys(stub, lib.UserKey, args)
	if err != nil {
		return shim.Error("The user does not exist.")
	}
	err = json.Unmarshal(result[0], &user)
	if err != nil {
		shim.Error(fmt.Sprintf("%s", err))
	}
	var token lib.Token
	result, err = utils.GetStateByPartialCompositeKeys(stub, lib.TokenKey, args)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	err = json.Unmarshal(result[0], &token)
	if err != nil {
		shim.Error(fmt.Sprintf("%s", err))
	}

	for _, v := range user.Control {
		if v == tokenId {
			tokenByte, err := json.Marshal(token)
			if err != nil {
				return shim.Error(fmt.Sprintf("序列化成功创建的信息出错: %s", err))
			}
			return shim.Success(tokenByte)
		}
	}

	return shim.Error("The user does not own the token.")
}*/
