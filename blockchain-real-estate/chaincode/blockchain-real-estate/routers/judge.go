package routers

import (
	//"awesomeProject3/blockchain/blockchain-real-estate/chaincode/blockchain-real-estate/lib"
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/togettoyou/blockchain-real-estate/chaincode/blockchain-real-estate/lib"
)

func JudgeOwnToken(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 2 {
		return shim.Error("Please offer the right number of parameters.")
	}

	userId := args[0]
	tokenId := args[1]

	if tokenId == "" {
		return shim.Error("Please offer the token id.")
	}
	if userId == "" {
		return shim.Error("Please offer the user id.")
	}

	var user lib.User
	err := json.Unmarshal(QueryAccount(stub, []string{userId}).Payload, &user)
	if err != nil {
		return shim.Error("The user does not exist.")
	}
	var token lib.Token
	err = json.Unmarshal(ReadToken(stub, []string{tokenId}).Payload, &token)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}

	if token.Owner == userId {
		tokenByte, err := json.Marshal(token)
		if err != nil {
			return shim.Error(fmt.Sprintf("序列化成功创建的信息出错: %s", err))
		}
		return shim.Success(tokenByte)
	}

	return shim.Error("The user does not own the token.")
}

//判断是否购买过某资源
func JudgeBuyResource(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 2 {
		return shim.Error("Please offer the right number of parameters.")
	}

	userId := args[0]
	resourceId := args[1]

	if resourceId == "" {
		return shim.Error("Please offer the token id.")
	}
	if userId == "" {
		return shim.Error("Please offer the user id.")
	}

	var user lib.User
	err := json.Unmarshal(QueryAccount(stub, []string{userId}).Payload, &user)
	if err != nil {
		return shim.Error("The user does not exist.")
	}
	var resource lib.Resource
	err = json.Unmarshal(QueryResource(stub, []string{resourceId}).Payload, &resource)
	if err != nil {
		return shim.Error("The resource does not exist")
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

//判断目前是否拥有某资源
func JudgeOwnResource(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 2 {
		return shim.Error("Please offer the right number of parameters.")
	}

	userId := args[0]
	resourceId := args[1]

	if resourceId == "" {
		return shim.Error("Please offer the token id.")
	}
	if userId == "" {
		return shim.Error("Please offer the user id.")
	}

	var user lib.User
	err := json.Unmarshal(QueryAccount(stub, []string{userId}).Payload, &user)
	if err != nil {
		return shim.Error("The user does not exist.")
	}
	var resource lib.Resource
	err = json.Unmarshal(QueryResource(stub, []string{resourceId}).Payload, &resource)
	if err != nil {
		return shim.Error("The resource does not exist")
	}

	if resource.Owner == userId {
		resourceByte, err := json.Marshal(resource)
		if err != nil {
			return shim.Error(fmt.Sprintf("序列化成功创建的信息出错: %s", err))
		}
		return shim.Success(resourceByte)
	}

	return shim.Error("The user does not own this resource.")
}

//判断是否有被分享某个资源
func JudgeShareResource(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 2 {
		return shim.Error("Please offer the right number of parameters.")
	}

	userId := args[0]
	resourceId := args[1]

	if resourceId == "" {
		return shim.Error("Please offer the token id.")
	}
	if userId == "" {
		return shim.Error("Please offer the user id.")
	}

	var user lib.User
	err := json.Unmarshal(QueryAccount(stub, []string{userId}).Payload, &user)
	if err != nil {
		return shim.Error("The user does not exist.")
	}
	var resource lib.Resource
	err = json.Unmarshal(QueryResource(stub, []string{resourceId}).Payload, &resource)
	if err != nil {
		return shim.Error("The resource does not exist")
	}

	for _, v := range user.Share {
		var token lib.Token
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
	if len(args) != 2 {
		return shim.Error("Please offer the right number of parameters.")
	}

	userId := args[0]
	tokenId := args[1]

	if tokenId == "" {
		return shim.Error("Please offer the token id.")
	}
	if userId == "" {
		return shim.Error("Please offer the user id.")
	}

	var user lib.User
	err := json.Unmarshal(QueryAccount(stub, []string{userId}).Payload, &user)
	if err != nil {
		return shim.Error("The user does not exist.")
	}
	var token lib.Token
	err = json.Unmarshal(ReadToken(stub, []string{tokenId}).Payload, &token)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}

	for _, v := range user.Share {
		if v == token.Id {
			tokenByte, err := json.Marshal(token)
			if err != nil {
				return shim.Error(fmt.Sprintf("序列化成功创建的信息出错: %s", err))
			}
			return shim.Success(tokenByte)
		}
	}

	return shim.Error("The user is not shared this token.")
}

func JudgeLendToken(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 2 {
		return shim.Error("Please offer the right number of parameters.")
	}

	userId := args[0]
	tokenId := args[1]

	if tokenId == "" {
		return shim.Error("Please offer the token id.")
	}
	if userId == "" {
		return shim.Error("Please offer the user id.")
	}

	var user lib.User
	err := json.Unmarshal(QueryAccount(stub, []string{userId}).Payload, &user)
	if err != nil {
		return shim.Error("The user does not exist.")
	}
	var token lib.Token
	err = json.Unmarshal(ReadToken(stub, []string{tokenId}).Payload, &token)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}

	for _, v := range user.Lend {
		if v == token.Id {
			tokenByte, err := json.Marshal(token)
			if err != nil {
				return shim.Error(fmt.Sprintf("序列化成功创建的信息出错: %s", err))
			}
			return shim.Success(tokenByte)
		}
	}

	return shim.Error("The user does not lend this token.")
}

func JudgeOwnProject(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 2 {
		return shim.Error("Please offer the right number of parameters.")
	}

	userId := args[0]
	projectId := args[1]

	if projectId == "" {
		return shim.Error("Please offer the token id.")
	}
	if userId == "" {
		return shim.Error("Please offer the user id.")
	}

	var user lib.User
	err := json.Unmarshal(QueryAccount(stub, []string{userId}).Payload, &user)
	if err != nil {
		return shim.Error("The user does not exist.")
	}
	var project lib.Project
	err = json.Unmarshal(QueryProject(stub, []string{projectId}).Payload, &project)
	if err != nil {
		return shim.Error("The project does not exist.")
	}
	if project.Owner == userId {
		projectByte, err := json.Marshal(project)
		if err != nil {
			return shim.Error(fmt.Sprintf("序列化成功创建的信息出错: %s", err))
		}
		return shim.Success(projectByte)
	}

	return shim.Error("The user does not upload this project.")
}
