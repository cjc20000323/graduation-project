package routers

import (
	//"awesomeProject3/blockchain/blockchain-real-estate/chaincode/blockchain-real-estate/lib"
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/togettoyou/blockchain-real-estate/chaincode/blockchain-real-estate/lib"
	"github.com/togettoyou/blockchain-real-estate/chaincode/blockchain-real-estate/utils"
)

func JudgeOwnToken(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	userId := args[0]
	tokenId := args[1]

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
}

func JudgeBuyResource(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	userId := args[0]
	resourceId := args[1]

	if resourceId == "" {
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
	var resource lib.Resource
	result, err = utils.GetStateByPartialCompositeKeys(stub, lib.ResourceKey, args)
	if err != nil {
		return shim.Error("The resource does not exist")
	}
	err = json.Unmarshal(result[0], &resource)
	if err != nil {
		shim.Error(fmt.Sprintf("%s", err))
	}

	for _, v := range user.Buy {
		if v == resourceId {
			resourceByte, err := json.Marshal(resource)
			if err != nil {
				return shim.Error(fmt.Sprintf("序列化成功创建的信息出错: %s", err))
			}
			return shim.Success(resourceByte)
		}
	}

	return shim.Error("The user hasn't bought this resource.")
}

func JudgeOwnResource(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	userId := args[0]
	resourceId := args[1]

	if resourceId == "" {
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
	var resource lib.Resource
	result, err = utils.GetStateByPartialCompositeKeys(stub, lib.ResourceKey, args)
	if err != nil {
		return shim.Error("The resource does not exist")
	}
	err = json.Unmarshal(result[0], &resource)
	if err != nil {
		shim.Error(fmt.Sprintf("%s", err))
	}

	var token lib.Token

	for _, v := range user.Control {
		err = json.Unmarshal(ReadToken(stub, []string{v}).Payload, &token)
		if token.Asset == resourceId {
			resourceByte, err := json.Marshal(resource)
			if err != nil {
				return shim.Error(fmt.Sprintf("序列化成功创建的信息出错: %s", err))
			}
			return shim.Success(resourceByte)
		}
	}

	return shim.Error("The user does not own this resource.")
}

func JudgeShareResource(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	userId := args[0]
	resourceId := args[1]

	if resourceId == "" {
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
	var resource lib.Resource
	result, err = utils.GetStateByPartialCompositeKeys(stub, lib.ResourceKey, args)
	if err != nil {
		return shim.Error("The resource does not exist")
	}
	err = json.Unmarshal(result[0], &resource)
	if err != nil {
		shim.Error(fmt.Sprintf("%s", err))
	}

	var token lib.Token

	for _, v := range user.Share {
		err = json.Unmarshal(ReadToken(stub, []string{v}).Payload, &token)
		if token.Asset == resourceId {
			resourceByte, err := json.Marshal(resource)
			if err != nil {
				return shim.Error(fmt.Sprintf("序列化成功创建的信息出错: %s", err))
			}
			return shim.Success(resourceByte)
		}
	}

	return shim.Error("The user is not shared this resource.")
}

func JudgeShareToken(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	userId := args[0]
	tokenId := args[1]

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

	for _, v := range user.Share {
		if v == tokenId {
			tokenByte, err := json.Marshal(token)
			if err != nil {
				return shim.Error(fmt.Sprintf("序列化成功创建的信息出错: %s", err))
			}
			return shim.Success(tokenByte)
		}
	}

	return shim.Error("The user is not shared this token.")
}
