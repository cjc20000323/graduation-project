package routers

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/togettoyou/blockchain-real-estate/chaincode/blockchain-real-estate/lib"
	"github.com/togettoyou/blockchain-real-estate/chaincode/blockchain-real-estate/utils"
	//"github.com/togettoyou/blockchain-real-estate/chaincode/blockchain-real-estate/utils"
)

//查询某用户上传资源
func QueryUpload(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Please offer the right number of parameters.")
	}

	var resourcelist []string
	var account lib.User
	id := args[0]

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
	if len(args) != 1 {
		return shim.Error("Please offer the right number of parameters.")
	}

	var resourcelist []string
	var account lib.User
	id := args[0]
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
	if len(args) != 1 {
		return shim.Error("Please offer the right number of parameters.")
	}

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
	if len(args) != 1 {
		return shim.Error("Please offer the right number of parameters.")
	}

	var resource lib.Resource //其实就一个
	id := args[0]
	if id == "" {
		return shim.Error("Please offer the resource id.")
	}
	result, err := utils.GetStateByPartialCompositeKeys(stub, lib.ResourceKey, args)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	if result == nil {
		return shim.Error("The resource does not exist.")
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
	if len(args) != 0 {
		return shim.Error("Please offer the right number of parameters.")
	}

	var accountList []lib.User
	results, err := utils.GetStateByPartialCompositeKeys(stub, lib.UserKey, args)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	if results == nil {
		return shim.Error("There are no accounts now.")
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
	if len(args) != 0 {
		return shim.Error("Please offer the right number of parameters.")
	}

	var resourcelist []lib.Resource
	results, err := utils.GetStateByPartialCompositeKeys(stub, lib.ResourceKey, args)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	if results == nil {
		return shim.Error("There are no resources now.")
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
	if len(args) != 1 {
		return shim.Error("Please offer the right number of parameters.")
	}

	id := args[0]
	if id == "" {
		return shim.Error("Please offer the token id.")
	}
	var token lib.Token
	result, err := utils.GetStateByPartialCompositeKeys2(stub, lib.TokenKey, args)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	if result == nil {
		return shim.Error("The token does not exist.")
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
	if len(args) != 1 {
		return shim.Error("Please offer the right number of parameters.")
	}

	id := args[0] //传入用户的id
	if id == "" {
		return shim.Error("Please offer the user id.")
	}
	var account lib.User
	result, err := utils.GetStateByPartialCompositeKeys(stub, lib.UserKey, args)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	if result == nil {
		return shim.Error("The user does not exist.")
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
	tokenListByte, err := json.Marshal(tokenList)
	if err != nil {
		return shim.Error(fmt.Sprintf("getUserTokenList-序列化出错: %s", err))
	}
	return shim.Success(tokenListByte)
}

//查询用户目前有人投标的代币
func QueryUserBiddenToken(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Please offer the right number of parameters.")
	}

	userId := args[0]

	if userId == "" {
		return shim.Error("Please offer the user id.")
	}

	var user lib.User
	err := json.Unmarshal(QueryAccount(stub, []string{userId}).Payload, &user)
	if err != nil {
		return shim.Error("The user does not exist.")
	}

	var tokenList []lib.Token

	for _, v := range user.Control {
		var token lib.Token
		err := json.Unmarshal(ReadToken(stub, []string{v}).Payload, &token)
		if err != nil {
			shim.Error(fmt.Sprintf("%s", err))
		}
		if token.Bid != "" {
			tokenList = append(tokenList, token)
		}
	}

	tokenListByte, err := json.Marshal(tokenList)
	if err != nil {
		return shim.Error(fmt.Sprintf("getUserTokenList-序列化出错: %s", err))
	}
	return shim.Success(tokenListByte)
}

//查询当前用户拥有的资源（通过控制的代币查询）
func QueryUserResource(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Please offer the right number of parameters.")
	}

	userId := args[0]

	if userId == "" {
		return shim.Error("Please offer the user id.")
	}

	var user lib.User
	err := json.Unmarshal(QueryAccount(stub, []string{userId}).Payload, &user)
	if err != nil {
		return shim.Error("The user does not exist.")
	}

	var resourceList []lib.Resource

	for _, v := range user.Control {
		var token lib.Token
		var resource lib.Resource
		err := json.Unmarshal(ReadToken(stub, []string{v}).Payload, &token)
		if err != nil {
			shim.Error(fmt.Sprintf("%s", err))
		}
		err = json.Unmarshal(QueryResource(stub, []string{token.Asset}).Payload, &resource)
		if err != nil {
			shim.Error(fmt.Sprintf("%s", err))
		}
		resourceList = append(resourceList, resource)
	}

	resourceListByte, err := json.Marshal(resourceList)
	if err != nil {
		return shim.Error(fmt.Sprintf("getUserTokenList-序列化出错: %s", err))
	}
	return shim.Success(resourceListByte)
}

//查询用户被分享的代币
func QueryUserSharedToken(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Please offer the right number of parameters.")
	}

	userId := args[0]

	if userId == "" {
		return shim.Error("Please offer the user id.")
	}

	var user lib.User
	err := json.Unmarshal(QueryAccount(stub, []string{userId}).Payload, &user)
	if err != nil {
		return shim.Error("The user does not exist.")
	}

	var tokenList []lib.Token

	for _, v := range user.Share {
		var token lib.Token
		err := json.Unmarshal(ReadToken(stub, []string{v}).Payload, &token)
		if err != nil {
			shim.Error(fmt.Sprintf("%s", err))
		}
		tokenList = append(tokenList, token)
	}

	tokenListByte, err := json.Marshal(tokenList)
	if err != nil {
		return shim.Error(fmt.Sprintf("getUserTokenList-序列化出错: %s", err))
	}
	return shim.Success(tokenListByte)
}

//查询用户分享给他人的代币
func QueryUserLendToken(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Please offer the right number of parameters.")
	}

	userId := args[0]

	if userId == "" {
		return shim.Error("Please offer the user id.")
	}

	var user lib.User
	err := json.Unmarshal(QueryAccount(stub, []string{userId}).Payload, &user)
	if err != nil {
		return shim.Error("The user does not exist.")
	}

	var tokenList []lib.Token

	for _, v := range user.Lend {
		var token lib.Token
		err := json.Unmarshal(ReadToken(stub, []string{v}).Payload, &token)
		if err != nil {
			shim.Error(fmt.Sprintf("%s", err))
		}
		tokenList = append(tokenList, token)
	}

	tokenListByte, err := json.Marshal(tokenList)
	if err != nil {
		return shim.Error(fmt.Sprintf("getUserTokenList-序列化出错: %s", err))
	}
	return shim.Success(tokenListByte)
}

//查询代币分享给哪些用户
func QueryTokenShare(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Please offer the right number of parameters.")
	}

	tokenId := args[0]

	if tokenId == "" {
		return shim.Error("Please offer the user id.")
	}

	var token lib.Token
	err := json.Unmarshal(ReadToken(stub, []string{tokenId}).Payload, &token)
	if err != nil {
		return shim.Error("The user does not exist.")
	}

	var userList []lib.User

	for _, v := range token.Share {
		var user lib.User
		err := json.Unmarshal(QueryAccount(stub, []string{v}).Payload, &user)
		if err != nil {
			shim.Error(fmt.Sprintf("%s", err))
		}
		userList = append(userList, user)
	}

	userListByte, err := json.Marshal(userList)
	if err != nil {
		return shim.Error(fmt.Sprintf("getUserTokenList-序列化出错: %s", err))
	}
	return shim.Success(userListByte)
}

func QueryProject(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Please offer the right number of parameters.")
	}

	id := args[0]
	if id == "" {
		return shim.Error("Please offer the project id.")
	}

	var project lib.Project
	result, err := utils.GetStateByPartialCompositeKeys(stub, lib.ProjectKey, args)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	if result == nil {
		return shim.Error("The project does not exist.")
	}
	err = json.Unmarshal(result[0], &project)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	projectByte, err := json.Marshal(project)
	if err != nil {
		return shim.Error(fmt.Sprintf("ReadToken-序列化出错: %s", err))
	}
	return shim.Success(projectByte)
}

func QueryUserProject(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Please offer the right number of parameters.")
	}

	userId := args[0]
	if userId == "" {
		return shim.Error("Please offer the user id.")
	}

	var user lib.User
	err := json.Unmarshal(QueryAccount(stub, []string{userId}).Payload, &user)
	if err != nil {
		return shim.Error("The user does not exist.")
	}

	var projectList []lib.Project

	for _, v := range user.Projects {
		var project lib.Project
		err := json.Unmarshal(QueryProject(stub, []string{v}).Payload, &project)
		if err != nil {
			shim.Error(fmt.Sprintf("%s", err))
		}
		projectList = append(projectList, project)
	}

	projectListByte, err := json.Marshal(projectList)
	if err != nil {
		return shim.Error(fmt.Sprintf("getUserTokenList-序列化出错: %s", err))
	}
	return shim.Success(projectListByte)
}

func QueryProjectResource(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Please offer the right number of parameters.")
	}

	projectId := args[0]
	if projectId == "" {
		return shim.Error("Please offer the project id.")
	}

	var project lib.Project
	err := json.Unmarshal(QueryProject(stub, []string{projectId}).Payload, &project)
	if err != nil {
		return shim.Error("The project does not exist.")
	}

	var resourceList []lib.Resource

	for _, v := range project.Bid {
		var resource lib.Resource
		err := json.Unmarshal(QueryResource(stub, []string{v}).Payload, &resource)
		if err != nil {
			shim.Error(fmt.Sprintf("%s", err))
		}
		resourceList = append(resourceList, resource)
	}

	resourceListByte, err := json.Marshal(resourceList)
	if err != nil {
		return shim.Error(fmt.Sprintf("getUserTokenList-序列化出错: %s", err))
	}
	return shim.Success(resourceListByte)
}

func QueryDeal(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	result, err := utils.GetStateByPartialCompositeKeys2(stub, lib.DealKey, args)
	if err != nil {
		return shim.Error(fmt.Sprintf("获取历史信息错误: %s", err))
	}
	if err != nil {
		return shim.Error(fmt.Sprintf("QueryDeal-序列化出错: %s", err))
	}
	var dealList []lib.TokenDeal
	for _, v := range result {
		if v != nil {
			var tokenDeal lib.TokenDeal
			err := json.Unmarshal(v, &tokenDeal)
			if err != nil {
				return shim.Error(fmt.Sprintf("Queryresourcelist-反序列化出错: %s", err))
			}
			dealList = append(dealList, tokenDeal)
		}
	}

	deallistByte, err := json.Marshal(dealList)
	if err != nil {
		return shim.Error(fmt.Sprintf("QueryDeal-序列化出错: %s", err))
	}
	return shim.Success(deallistByte)
}

func QueryAllDealSum(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 0 {
		return shim.Error("Please offer the right number of parameters.")
	}

	results, err := utils.GetStateByPartialCompositeKeys(stub, lib.DealKey, args)

	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	if results == nil {
		sum := 0
		sumByte, err := json.Marshal(sum)
		if err != nil {
			return shim.Error(fmt.Sprintf("QueryAllDealSum-序列化出错: %s", err))
		}
		return shim.Success(sumByte)
	}
	var sum = 0
	for _, v := range results {
		if v != nil {
			sum++
		}
	}
	sumByte, err := json.Marshal(sum)
	if err != nil {
		return shim.Error(fmt.Sprintf("QueryAllDealSum-序列化出错: %s", err))
	}
	return shim.Success(sumByte)

}

func QueryAllResourceSum(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 0 {
		return shim.Error("Please offer the right number of parameters.")
	}

	results, err := utils.GetStateByPartialCompositeKeys(stub, lib.ResourceKey, args)

	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	if results == nil {
		sum := 0
		sumByte, err := json.Marshal(sum)
		if err != nil {
			return shim.Error(fmt.Sprintf("QueryAllResourceSum-序列化出错: %s", err))
		}
		return shim.Success(sumByte)
	}
	var sum = 0
	for _, v := range results {
		if v != nil {
			sum++
		}
	}
	sumByte, err := json.Marshal(sum)
	if err != nil {
		return shim.Error(fmt.Sprintf("QueryAllResourceSum-序列化出错: %s", err))
	}
	return shim.Success(sumByte)
}

func QueryAllUserSum(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 0 {
		return shim.Error("Please offer the right number of parameters.")
	}

	results, err := utils.GetStateByPartialCompositeKeys(stub, lib.UserKey, args)

	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	if results == nil {
		sum := 0
		sumByte, err := json.Marshal(sum)
		if err != nil {
			return shim.Error(fmt.Sprintf("QueryAllDealSum-序列化出错: %s", err))
		}
		return shim.Success(sumByte)
	}
	var sum = 0
	for _, v := range results {
		if v != nil {
			sum++
		}
	}
	sumByte, err := json.Marshal(sum)
	if err != nil {
		return shim.Error(fmt.Sprintf("QueryAllDealSum-序列化出错: %s", err))
	}
	return shim.Success(sumByte)
}

func QueryAllBlockSum(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 0 {
		return shim.Error("Please offer the right number of parameters.")
	}

	results, err := utils.GetStateByPartialCompositeKeys(stub, lib.DealKey, args)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}

	var sum = 0

	for _, v := range results {
		var deal lib.Deal
		err := json.Unmarshal(v, &deal)
		if err != nil {
			return shim.Error(fmt.Sprintf("QueryAllBlockSum-反序列化出错: %s", err))
		}
		result, err := utils.GetHistoryForKeys(stub, lib.ResourceKey, args) //查找该资源得历史数据
		for _, _ = range result {
			sum++
		}
	}

	results, err = utils.GetStateByPartialCompositeKeys(stub, lib.UserKey, args)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}

	for _, v := range results {
		var deal lib.Deal
		err := json.Unmarshal(v, &deal)
		if err != nil {
			return shim.Error(fmt.Sprintf("QueryAllBlockSum-反序列化出错: %s", err))
		}
		result, err := utils.GetHistoryForKeys(stub, lib.ResourceKey, args) //查找该资源得历史数据
		for _, _ = range result {
			sum++
		}
	}

	results, err = utils.GetStateByPartialCompositeKeys(stub, lib.ResourceKey, args)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}

	for _, v := range results {
		var deal lib.Deal
		err := json.Unmarshal(v, &deal)
		if err != nil {
			return shim.Error(fmt.Sprintf("QueryAllBlockSum-反序列化出错: %s", err))
		}
		result, err := utils.GetHistoryForKeys(stub, lib.ResourceKey, args) //查找该资源得历史数据
		for _, _ = range result {
			sum++
		}
	}

	results, err = utils.GetStateByPartialCompositeKeys(stub, lib.TokenKey, args)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}

	for _, v := range results {
		var deal lib.Deal
		err := json.Unmarshal(v, &deal)
		if err != nil {
			return shim.Error(fmt.Sprintf("QueryAllBlockSum-反序列化出错: %s", err))
		}
		result, err := utils.GetHistoryForKeys(stub, lib.ResourceKey, args) //查找该资源得历史数据
		for _, _ = range result {
			sum++
		}
	}

	sumByte, err := json.Marshal(sum)
	if err != nil {
		return shim.Error(fmt.Sprintf("QueryAllBlockSum-序列化出错: %s", err))
	}
	return shim.Success(sumByte)
}

func QueryBuyDeal(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Please offer the right number of parameters.")
	}
	result, err := utils.GetStateByPartialCompositeKeys2(stub, lib.DealKey, args)
	if err != nil {
		return shim.Error(fmt.Sprintf("获取历史信息错误: %s", err))
	}
	var dealList []lib.TokenDeal
	for _, v := range result {
		if v != nil {
			var tokenDeal lib.TokenDeal
			err := json.Unmarshal(v, &tokenDeal)
			if err != nil {
				return shim.Error(fmt.Sprintf("Queryresourcelist-反序列化出错: %s", err))
			}
			if tokenDeal.Type == "buy" || tokenDeal.Type == "transfer" {
				dealList = append(dealList, tokenDeal)
			}
		}
	}

	deallistByte, err := json.Marshal(dealList)
	if err != nil {
		return shim.Error(fmt.Sprintf("QueryDeal-序列化出错: %s", err))
	}
	return shim.Success(deallistByte)
}

func QueryCoinRecord(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Please offer the right number of parameters.")
	}

	result, err := utils.GetStateByPartialCompositeKeys2(stub, lib.CoinRecordKey, args)
	if err != nil {
		return shim.Error(fmt.Sprintf("获取历史信息错误: %s", err))
	}
	var recordList []lib.CoinRecord

	for _, v := range result {
		if v != nil {
			var record lib.CoinRecord
			err := json.Unmarshal(v, &record)
			if err != nil {
				return shim.Error(fmt.Sprintf("Queryresourcelist-反序列化出错: %s", err))
			}
			recordList = append(recordList, record)
		}
	}

	recordListByte, err := json.Marshal(recordList)
	if err != nil {
		return shim.Error(fmt.Sprintf("QueryDeal-序列化出错: %s", err))
	}
	return shim.Success(recordListByte)
}

func QueryBlock(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	results, err := utils.GetStateByPartialCompositeKeys(stub, lib.TokenKey, args)

	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	var idList []string
	for _, v := range results {
		var token lib.Token
		err := json.Unmarshal(v, &token)
		if err != nil {
			return shim.Error(fmt.Sprintf("QueryAccountList-反序列化出错: %s", err))
		}
		result, err := utils.GetHistoryForKeysToken(stub, lib.TokenKey, []string{token.Id})
		if err != nil {
			return shim.Error(fmt.Sprintf("%s", err))
		}
		var list []lib.TokenHistory
		err = json.Unmarshal(result, &list)
		if err != nil {
			return shim.Error(fmt.Sprintf("QueryAccountList-反序列化出错: %s", err))
		}
		for _, data := range list {
			if !IsContain(idList, data.TxId) {
				idList = append(idList, data.TxId)
			}
		}
	}
	results, err = utils.GetStateByPartialCompositeKeys(stub, lib.UserKey, args)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}

	for _, v := range results {
		var user lib.User
		err := json.Unmarshal(v, &user)
		if err != nil {
			return shim.Error(fmt.Sprintf("QueryAccountList-反序列化出错: %s", err))
		}
		result, err := utils.GetHistoryForKeysToken(stub, lib.UserKey, []string{user.Id})
		if err != nil {
			return shim.Error(fmt.Sprintf("%s", err))
		}
		var list []lib.UserHistory
		err = json.Unmarshal(result, &list)
		if err != nil {
			return shim.Error(fmt.Sprintf("QueryAccountList-反序列化出错: %s", err))
		}
		for _, data := range list {
			if !IsContain(idList, data.TxId) {
				idList = append(idList, data.TxId)
			}
		}
	}

	results, err = utils.GetStateByPartialCompositeKeys(stub, lib.ResourceKey, args)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}

	for _, v := range results {
		var resource lib.Resource
		err := json.Unmarshal(v, &resource)
		if err != nil {
			return shim.Error(fmt.Sprintf("QueryAccountList-反序列化出错: %s", err))
		}
		result, err := utils.GetHistoryForKeysToken(stub, lib.ResourceKey, []string{resource.Id})
		if err != nil {
			return shim.Error(fmt.Sprintf("%s", err))
		}
		var list []lib.ResourceHistory
		err = json.Unmarshal(result, &list)
		if err != nil {
			return shim.Error(fmt.Sprintf("QueryAccountList-反序列化出错: %s", err))
		}
		for _, data := range list {
			if !IsContain(idList, data.TxId) {
				idList = append(idList, data.TxId)
			}
		}
	}

	results, err = utils.GetStateByPartialCompositeKeys(stub, lib.DealKey, args)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}

	for _, v := range results {
		var deal lib.Deal
		err := json.Unmarshal(v, &deal)
		if err != nil {
			return shim.Error(fmt.Sprintf("QueryAccountList-反序列化出错: %s", err))
		}
		result, err := utils.GetHistoryForKeysToken(stub, lib.DealKey, []string{deal.Sell_id})
		if err != nil {
			return shim.Error(fmt.Sprintf("%s", err))
		}
		var list []lib.DealHistory
		err = json.Unmarshal(result, &list)
		if err != nil {
			return shim.Error(fmt.Sprintf("QueryAccountList-反序列化出错: %s", err))
		}
		for _, data := range list {
			if !IsContain(idList, data.TxId) {
				idList = append(idList, data.TxId)
			}
		}
	}

	sumByte, err := json.Marshal(len(idList))
	return shim.Success(sumByte)
}

func IsContain(items []string, item string) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}
