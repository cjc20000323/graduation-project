/**
 * @Author: 夜央 Oh oh oh oh oh oh (https://github.com/togettoyou)
 * @Email: zoujh99@qq.com
 * @Date: 2020/3/5 4:18 下午
 * @Description: 账户相关合约路由
 */
package routers

import (
	"encoding/json"
	"fmt"
	//idworker "github.com/gitstliu/go-id-worker"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/togettoyou/blockchain-real-estate/chaincode/blockchain-real-estate/lib"
	"github.com/togettoyou/blockchain-real-estate/chaincode/blockchain-real-estate/utils"
	"strconv"
	"time"
)

const Layout = "2006-01-02 15:04:05" //时间常量

//新建用户
func CreateUser(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 2 {
		return shim.Error("Please offer the right number of parameters.")
	}

	Id := args[0]
	Score := args[1] //初始积分
	if Id == "" || Score == "" {
		return shim.Error("参数存在空值")
	}

	// 参数数据格式转换
	var formattedTotalScore float64
	if val, err := strconv.ParseFloat(Score, 64); err != nil {
		return shim.Error(fmt.Sprintf("积分参数格式转换出错: %s", err))
	} else {
		formattedTotalScore = val
	}

	//判断用户是否存在在关系型处理

	NewUser := &lib.User{
		Id:    args[0],
		Score: formattedTotalScore,
	}
	// 写入账本
	if err := utils.WriteLedger(NewUser, stub, lib.UserKey, []string{NewUser.Id}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	//将成功创建的信息返回
	NewUserByte, err := json.Marshal(NewUser)
	if err != nil {
		return shim.Error(fmt.Sprintf("序列化成功创建的信息出错: %s", err))
	}
	// 成功返回
	return shim.Success(NewUserByte)
}

//创建数据资源
func UploadResource(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 7 {
		return shim.Error("Please offer the right number of parameters.")
	}

	Id := args[0]
	Hash := args[1]
	Owner := args[2]
	Time := args[3]
	Cost := args[4]
	GetScore := args[5]
	Kind := args[6]

	if Id == "" || Hash == "" || Owner == "" || Time == "" || Cost == "" || Kind == "" {
		return shim.Error("参数存在空值")
	}

	// 参数数据格式转换
	var formattedCost float64
	if val, err := strconv.ParseFloat(Cost, 64); err != nil {
		return shim.Error(fmt.Sprintf("积分参数格式转换出错: %s", err))
	} else {
		formattedCost = val
	}
	var formattedScore float64
	if val, err := strconv.ParseFloat(GetScore, 64); err != nil {
		return shim.Error(fmt.Sprintf("积分参数格式转换出错: %s", err))
	} else {
		formattedScore = val
	}

	NewResource := &lib.Resource{
		Id:    args[0],
		Hash:  args[1],
		Owner: args[2],
		Time:  args[3],
		Cost:  formattedCost,
		Type:  args[6],
	}
	// 写入资源账本
	if err := utils.WriteLedger(NewResource, stub, lib.ResourceKey, []string{NewResource.Id}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	//写入用户帐本
	var account lib.User
	err := json.Unmarshal(QueryAccount(stub, []string{args[2]}).Payload, &account)
	if err != nil {
		return shim.Error(fmt.Sprintf("用户信息验证失败%s", err))
	}

	old_resources := account.Upload
	new_resources := append(old_resources, NewResource.Id) //只要ID
	account.Upload = new_resources
	account.Score = account.Score + formattedScore

	if err := utils.WriteLedger(account, stub, lib.UserKey, []string{account.Id}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	//将成功创建的信息返回
	accountByte, err := json.Marshal(account)
	if err != nil {
		return shim.Error(fmt.Sprintf("序列化成功创建的信息出错: %s", err))
	}
	// 成功返回
	return shim.Success(accountByte)
}

func CreateDeal(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	Sell_id := args[0]
	Buy_id := args[1]
	Rescource_id := args[2]
	Cost := args[3] //多少买就必须多少走
	if Sell_id == "" || Buy_id == "" || Rescource_id == "" || Cost == "" {
		return shim.Error("参数存在空值")
	}

	// 参数数据格式转换
	var formattedCost float64
	if val, err := strconv.ParseFloat(Cost, 64); err != nil {
		return shim.Error(fmt.Sprintf("积分参数格式转换出错: %s", err))
	} else {
		formattedCost = val
	}
	// time转换。先检查time格式，再转换成time
	timeLocal, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return shim.Error(fmt.Sprintf("时区设置失败%s", err))
	}
	time.Local = timeLocal
	formattedTime := time.Now()

	//查找卖方
	var account_seller lib.User
	err_seller := json.Unmarshal(QueryAccount(stub, []string{Sell_id}).Payload, &account_seller)
	if err_seller != nil {
		return shim.Error(fmt.Sprintf("用户信息验证失败%s", err_seller))
	}

	//查找买方
	var account_buyer lib.User
	err_buyer := json.Unmarshal(QueryAccount(stub, []string{Buy_id}).Payload, &account_buyer)
	if err_buyer != nil {
		return shim.Error(fmt.Sprintf("用户信息验证失败%s", err_buyer))
	}

	//查找资源
	var resource lib.Resource
	err_resource := json.Unmarshal(QueryResource(stub, []string{Rescource_id}).Payload, &resource)
	if err_resource != nil {
		return shim.Error(fmt.Sprintf("资源信息验证失败%s", err_resource))
	}

	if resource.Owner != Sell_id { //买方不是拥有者
		return shim.Error(fmt.Sprintf("卖方不是拥有者%s", err))
	}
	//更新卖方积分，买方buy列表和积分，更新resource的上次操作时间和拥有者
	resource.Time = formattedTime.Format(Layout)
	resource.Owner = Buy_id

	old_buyresources := account_buyer.Buy

	var havebuy = false
	for _, v := range old_buyresources { //可用二分查找优化，因为go语言没有in之类的函数
		if v == resource.Id {
			havebuy = true
			break
		}
	}
	if havebuy == false {
		new_resources := append(old_buyresources, resource.Id) //只添加Id
		account_buyer.Buy = new_resources
		//购买过但是现在所有人不是自己，不重复添加购买过信息，但是仍能更换所有权和积分
	}
	account_buyer.Score = account_buyer.Score - formattedCost
	account_seller.Score = account_seller.Score + formattedCost

	//这种存储可能导致冗余，但是如果不冗余就需要扫所有资源找到买过的资源，所以这样存储并没有问题，空间换时间，寻找买过
	if err := utils.WriteLedger(account_buyer, stub, lib.UserKey, []string{account_buyer.Id}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	if err := utils.WriteLedger(account_seller, stub, lib.UserKey, []string{account_seller.Id}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	//更改resource的状态
	if err := utils.WriteLedger(resource, stub, lib.ResourceKey, []string{resource.Id}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	//将成功创建的信息返回
	accountByte, err := json.Marshal(account_buyer)
	if err != nil {
		return shim.Error(fmt.Sprintf("序列化成功创建的信息出错: %s", err))
	}
	// 成功返回
	return shim.Success(accountByte)
}

func CreateToken(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 9 {
		return shim.Error("Please offer the right number of parameters.")
	}

	id := args[0]
	assets := args[1]
	notforsale := args[2]
	owner := args[3]
	bid := args[4]   //一个token的bid可以是没有的，也可以有
	notes := args[5] //对token的描述可以有也可以没有
	value := args[6]
	createrId := args[7]
	Time := args[8]

	if id == "" {
		return shim.Error("Please offer the token id.")
	}
	if assets == "" {
		return shim.Error("Please offer the assets of the token.")
	}
	if notforsale == "" {
		return shim.Error("Please offer the message that whether the token is for sale.")
	}
	if owner == "" {
		return shim.Error("Please offer the owner of the token.")
	}
	if Time == "" {
		return shim.Error("Please offer the time.")
	}
	result, err := strconv.ParseFloat(value, 64)
	if result < 0 {
		shim.Error("The value of the token is not legal.")
	}
	if createrId == "" {
		return shim.Error("Please offer the creater of the token.")
	}
	if createrId != owner {
		return shim.Error("The creater is not the owner. Please check these two ids.")
	}
	var user lib.User
	err = json.Unmarshal(QueryAccount(stub, []string{createrId}).Payload, &user)
	if err != nil {
		return shim.Error("The user does not exist.")
	}
	var resource lib.Resource
	err = json.Unmarshal(QueryResource(stub, []string{assets}).Payload, &resource)
	if err != nil {
		return shim.Error("The resource does not exist")
	}
	if resource.Cost != result {
		return shim.Error("The value is wrong.")
	}
	var token lib.Token
	err = json.Unmarshal(ReadToken(stub, []string{id}).Payload, &token)
	if err == nil {
		return shim.Error("The token has existed")
	}

	saleMessage, err := strconv.ParseBool(notforsale)
	if err != nil {
		return shim.Error("The format of Notforsale is not boolean")
	}
	NewToken := &lib.Token{
		Id:         id,
		Asset:      assets,
		NotForSale: saleMessage,
		Owner:      owner,
		Bid:        bid,
		Notes:      notes,
		Value:      result,
		Time:       args[8],
	}

	if err := utils.WriteLedger(NewToken, stub, lib.TokenKey, []string{NewToken.Asset, NewToken.Time}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	if err := utils.WriteLedger(NewToken, stub, lib.TokenKey, []string{NewToken.Id}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}

	//向User拥有的资产列表中写入
	var account lib.User
	err = json.Unmarshal(QueryAccount(stub, []string{args[3]}).Payload, &account)
	if err != nil {
		return shim.Error(fmt.Sprintf("用户信息验证失败%s", err))
	}
	old_resources := account.Upload
	new_resources := append(old_resources, NewToken.Asset)
	oldControls := account.Control
	newControls := append(oldControls, NewToken.Id)
	account.Upload = new_resources
	account.Control = newControls
	if err := utils.WriteLedger(account, stub, lib.UserKey, []string{account.Id}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}

	NewTokenByte, err := json.Marshal(NewToken)
	if err != nil {
		return shim.Error(fmt.Sprintf("序列化成功创建的信息出错: %s", err))
	}
	return shim.Success(NewTokenByte)
}

func BidToken(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 2 {
		return shim.Error("Please offer the right number of parameters.")
	}

	id := args[0]
	bidder := args[1] //投标人的id

	if id == "" {
		return shim.Error("Please offer the token id.")
	}
	if bidder == "" {
		return shim.Error("Please offer the bidder id.")
	}
	var token lib.Token
	var user lib.User
	err := json.Unmarshal(QueryAccount(stub, []string{bidder}).Payload, &user)
	if err != nil {
		return shim.Error("The bidder does not exist.")
	}
	err = json.Unmarshal(ReadToken(stub, []string{id}).Payload, &token)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	err = json.Unmarshal(JudgeOwnToken(stub, []string{bidder, id}).Payload, &token)
	if err == nil {
		return shim.Error("You own this token.")
	}
	if token.NotForSale {
		return shim.Error("This token is not for sale.")
	}
	if token.Bid != "" {
		return shim.Error("Sorry, this token is bid by others, please try again later.")
	}
	token.Bid = bidder
	if err := utils.WriteLedger(token, stub, lib.TokenKey, []string{token.Id}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	tokenByte, err := json.Marshal(token)
	if err != nil {
		return shim.Error(fmt.Sprintf("序列化成功创建的信息出错: %s", err))
	}
	return shim.Success(tokenByte)
}

func DeleteToken(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Please offer the right number of parameters.")
	}

	id := args[0] //token的id
	if id == "" {
		return shim.Error("Please offer the id of token that you want to delete.")
	}

	var token lib.Token
	var user lib.User
	err := json.Unmarshal(ReadToken(stub, args).Payload, &token)
	if err != nil {
		return shim.Error("Can not read the token.")
	}

	result, err := utils.GetStateByPartialCompositeKeys(stub, lib.UserKey, []string{token.Owner})
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}

	err = json.Unmarshal(result[0], &user)
	if err != nil {
		return shim.Error("Can not find the user")
	}

	//oldResources := user.Upload
	//var newResources []string
	var newControl []string
	/*for _, v := range oldResources {
		if v != id {
			newResources = append(newResources, v)
		}
	}*/
	for _, v := range user.Control {
		if v != id {
			newControl = append(newControl, v)
		}
	}
	//user.Upload = newResources
	user.Control = newControl
	if err := utils.WriteLedger(user, stub, lib.UserKey, []string{user.Id}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}

	err = utils.DelLedger(stub, lib.TokenKey, []string{id})
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}

	accountByte, err := json.Marshal(user)
	if err != nil {
		return shim.Error(fmt.Sprintf("序列化成功创建的信息出错: %s", err))
	}
	// 成功返回
	return shim.Success(accountByte)
}

func GiveToken(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 3 {
		return shim.Error("Please offer the right number of parameters.")
	}

	var from lib.User
	var to lib.User
	var token lib.Token
	fromId := args[0]
	toId := args[1]
	tokenId := args[2]

	if fromId == "" {
		return shim.Error("Please offer the user id.")
	}
	if toId == "" {
		return shim.Error("Please offer the id of the user that you want to share with.")
	}
	if tokenId == "" {
		return shim.Error("Please offer the token id.")
	}

	err := json.Unmarshal(QueryAccount(stub, []string{fromId}).Payload, &from)
	if err != nil {
		return shim.Error("The user does not exist.")
	}
	err = json.Unmarshal(QueryAccount(stub, []string{toId}).Payload, &to)
	if err != nil {
		return shim.Error("The user you want to share with does not exist.")
	}
	err = json.Unmarshal(ReadToken(stub, []string{tokenId}).Payload, &token)
	if err != nil {
		return shim.Error("The token you want to share does not exist.")
	}
	if token.Owner != fromId {
		return shim.Error("You don't have the token.")
	}
	if token.Bid != "" {
		return shim.Error("This token has already been bid.")
	}

	var newToken lib.Token
	var resource lib.Resource
	var newNotForSale string
	if token.NotForSale {
		newNotForSale = "true"
	} else {
		newNotForSale = "false"
	}
	newValue := strconv.FormatFloat(token.Value, 'f', 2, 64)
	err = json.Unmarshal(CreateToken(stub, []string{"new" + tokenId, token.Asset, newNotForSale, toId, "",
		token.Notes, newValue, toId, time.Now().String()}).Payload, &newToken)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	err = json.Unmarshal(QueryResource(stub, []string{token.Asset}).Payload, &resource)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}

	to.Control = append(to.Control, tokenId)
	resource.Owner = toId

	if err := utils.WriteLedger(to, stub, lib.UserKey, []string{to.Id}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	if err := utils.WriteLedger(resource, stub, lib.ResourceKey, []string{resource.Id}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}

	var newAccount lib.User
	err = json.Unmarshal(DeleteToken(stub, []string{tokenId}).Payload, &newAccount)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}

	newToken.Id = tokenId
	if err := utils.WriteLedger(newToken, stub, lib.TokenKey, []string{tokenId}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	err = utils.DelLedger(stub, lib.TokenKey, []string{"new" + tokenId})
	if err != nil {
		return shim.Error("Del Wrong.")
	}

	NewTokenByte, err := json.Marshal(newToken)
	if err != nil {
		return shim.Error(fmt.Sprintf("序列化成功创建的信息出错: %s", err))
	}
	return shim.Success(NewTokenByte)

}

func TransferToken(stub shim.ChaincodeStubInterface, args []string) pb.Response { //代币的所有者通过此链码将自己的代币转给其他人
	if len(args) != 2 {
		return shim.Error("Please offer the right number of parameters.")
	}

	tokenId := args[0]
	userId := args[1]

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
	if token.NotForSale {
		return shim.Error("This token is not for sale.")
	}
	if token.Owner != userId {
		return shim.Error("You are not the owner of this token, so the change is illegal.")
	}
	if token.Bid == "" {
		return shim.Error("This token has not been bid yet.")
	}

	/*currWorker := &idworker.IdWorker{}
	currWorker.InitIdWorker(1000, 1)
	newID, err := currWorker.NextId()*/
	//以上的一小段代码利用uuid方法，和snowflake类似

	//newTokenId := strconv.FormatInt(newID, 10)
	var newNotForSale string
	if token.NotForSale {
		newNotForSale = "true"
	} else {
		newNotForSale = "false"
	}
	newValue := strconv.FormatFloat(token.Value, 'f', 2, 64)

	var newToken lib.Token
	var bidder lib.User
	var resource lib.Resource
	err = json.Unmarshal(CreateToken(stub, []string{"new" + tokenId, token.Asset, newNotForSale, token.Bid, "",
		token.Notes, newValue, token.Bid, time.Now().String()}).Payload, &newToken)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}

	err = json.Unmarshal(QueryAccount(stub, []string{token.Bid}).Payload, &bidder)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	err = json.Unmarshal(QueryResource(stub, []string{token.Asset}).Payload, &resource)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	bidder.Score -= token.Value
	bidder.Control = append(bidder.Control, token.Id)
	bidder.Buy = append(bidder.Buy, token.Asset) //可能会有重复，但不影响查询
	resource.Owner = bidder.Id
	if err := utils.WriteLedger(bidder, stub, lib.UserKey, []string{bidder.Id}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	if err := utils.WriteLedger(resource, stub, lib.ResourceKey, []string{resource.Id}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}

	var transformer lib.User
	var newAccount lib.User
	err = json.Unmarshal(QueryAccount(stub, []string{userId}).Payload, &transformer)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	transformer.Score += token.Value
	err = json.Unmarshal(DeleteToken(stub, []string{tokenId}).Payload, &newAccount)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	newAccount.Score = transformer.Score
	if err := utils.WriteLedger(newAccount, stub, lib.UserKey, []string{transformer.Id}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	newToken.Id = tokenId
	if err := utils.WriteLedger(newToken, stub, lib.TokenKey, []string{tokenId}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}

	err = utils.DelLedger(stub, lib.TokenKey, []string{"new" + tokenId})
	if err != nil {
		return shim.Error("Del Wrong.")
	}
	NewTokenByte, err := json.Marshal(newToken)
	if err != nil {
		return shim.Error(fmt.Sprintf("序列化成功创建的信息出错: %s", err))
	}
	return shim.Success(NewTokenByte)
}

func TransferToken2(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 2 {
		return shim.Error("Please offer the right number of parameters.")
	}

	tokenId := args[0]
	userId := args[1]

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
	if token.NotForSale {
		return shim.Error("This token is not for sale.")
	}
	if token.Owner != userId {
		return shim.Error("You are not the owner of this token, so the change is illegal.")
	}
	if token.Bid == "" {
		return shim.Error("This token has not been bid yet.")
	}

	var transformer lib.User
	var bidder lib.User
	var resource lib.Resource
	err = json.Unmarshal(QueryAccount(stub, []string{userId}).Payload, &transformer)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	transformer.Score += token.Value
	err = json.Unmarshal(QueryAccount(stub, []string{token.Bid}).Payload, &bidder)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	err = json.Unmarshal(QueryResource(stub, []string{token.Asset}).Payload, &resource)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	bidder.Score -= token.Value
	bidder.Buy = append(bidder.Buy, token.Asset)
	bidder.Control = append(bidder.Control, token.Id)
	token.Bid = ""
	token.Owner = bidder.Id
	resource.Owner = bidder.Id

	var newControl []string
	for _, v := range transformer.Control {
		if v != token.Id {
			newControl = append(newControl, v)
		}
	}
	transformer.Control = newControl

	if err := utils.WriteLedger(transformer, stub, lib.UserKey, []string{transformer.Id}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	if err := utils.WriteLedger(bidder, stub, lib.UserKey, []string{bidder.Id}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	if err := utils.WriteLedger(token, stub, lib.TokenKey, []string{token.Id}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	if err := utils.WriteLedger(token, stub, lib.TokenKey, []string{token.Asset, token.Time}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	if err := utils.WriteLedger(resource, stub, lib.ResourceKey, []string{resource.Id}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}

	NewTokenByte, err := json.Marshal(token)
	if err != nil {
		return shim.Error(fmt.Sprintf("序列化成功创建的信息出错: %s", err))
	}
	return shim.Success(NewTokenByte)
}

func ShareToken(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 3 {
		return shim.Error("Please offer the right number of parameters.")
	}

	var from lib.User
	var to lib.User
	var token lib.Token
	fromId := args[0]
	toId := args[1]
	tokenId := args[2]

	if fromId == "" {
		return shim.Error("Please offer the user id.")
	}
	if toId == "" {
		return shim.Error("Please offer the id of the user that you want to share with.")
	}
	if tokenId == "" {
		return shim.Error("Please offer the token id.")
	}

	if fromId == toId {
		return shim.Error("The from is the same as the to.")
	}

	err := json.Unmarshal(QueryAccount(stub, []string{fromId}).Payload, &from)
	if err != nil {
		return shim.Error("The user does not exist.")
	}
	err = json.Unmarshal(QueryAccount(stub, []string{toId}).Payload, &to)
	if err != nil {
		return shim.Error("The user you want to share with does not exist.")
	}
	err = json.Unmarshal(ReadToken(stub, []string{tokenId}).Payload, &token)
	if err != nil {
		return shim.Error("The token you want to share does not exist.")
	}
	if token.Owner != fromId {
		return shim.Error("You don't have the token.")
	}

	var anotherToken lib.Token
	err = json.Unmarshal(JudgeShareToken(stub, []string{toId, token.Id}).Payload, &anotherToken)
	if err != nil {
		to.Share = append(to.Share, token.Id)
		token.Share = append(token.Share, toId)
		err = json.Unmarshal(JudgeLendToken(stub, []string{fromId, token.Id}).Payload, &anotherToken)
		if err != nil {
			from.Lend = append(from.Lend, token.Id)
		}
	} else {
		return shim.Error("The token has shared with this user.")
	}
	/*var flag bool
	flag = false
	for _, v := range to.Share {
		if v == tokenId {
			flag = true
			break
		}
	}*/
	//这里可以换成写好的判断的智能合约
	/*if !flag {
		to.Share = append(to.Share, tokenId)
		token.Share = append(token.Share, toId)
		var anotherToken lib.Token
		err = json.Unmarshal(JudgeLendToken(stub, []string{fromId, tokenId}).Payload, &anotherToken)
		if err != nil {
			from.Lend = append(from.Lend, tokenId)
		}
	}*/

	if err := utils.WriteLedger(to, stub, lib.UserKey, []string{to.Id}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	if err := utils.WriteLedger(from, stub, lib.UserKey, []string{from.Id}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	if err = utils.WriteLedger(token, stub, lib.TokenKey, []string{token.Id}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	if err = utils.WriteLedger(token, stub, lib.TokenKey, []string{token.Asset, token.Time}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}

	toByte, err := json.Marshal(to)
	return shim.Success(toByte) //返回被分享对象的json
}

func RedeemToken(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 3 {
		return shim.Error("Please offer the right number of parameters.")
	}

	var from lib.User
	var to lib.User
	var token lib.Token
	fromId := args[0]
	toId := args[1]
	tokenId := args[2]

	if fromId == "" {
		return shim.Error("Please offer the user id.")
	}
	if toId == "" {
		return shim.Error("Please offer the id of the user that you want to share with.")
	}
	if tokenId == "" {
		return shim.Error("Please offer the token id.")
	}

	err := json.Unmarshal(QueryAccount(stub, []string{fromId}).Payload, &from)
	if err != nil {
		return shim.Error("The user does not exist.")
	}
	err = json.Unmarshal(QueryAccount(stub, []string{toId}).Payload, &to)
	if err != nil {
		return shim.Error("The user you want to share with does not exist.")
	}
	err = json.Unmarshal(ReadToken(stub, []string{tokenId}).Payload, &token)
	if err != nil {
		return shim.Error("The token you want to share does not exist.")
	}
	if token.Owner != fromId {
		return shim.Error("You don't have the token.")
	}

	var newShares []string
	for _, v := range to.Share {
		if v != token.Id {
			newShares = append(newShares, v)
		}
	}
	var shares []string
	for _, v := range token.Share {
		if v != toId {
			shares = append(shares, v)
		}
	}
	to.Share = newShares
	token.Share = shares
	if len(shares) == 0 {
		var lend []string
		for _, v := range from.Lend {
			if v != token.Id {
				lend = append(lend, v)
			}
		}
		from.Lend = lend
	}

	if err := utils.WriteLedger(to, stub, lib.UserKey, []string{to.Id}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	if err := utils.WriteLedger(from, stub, lib.UserKey, []string{from.Id}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	if err = utils.WriteLedger(token, stub, lib.TokenKey, []string{token.Id}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	if err = utils.WriteLedger(token, stub, lib.TokenKey, []string{token.Asset, token.Time}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}

	toByte, err := json.Marshal(to)
	return shim.Success(toByte)
}

func ChangeResourceScore(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 3 {
		return shim.Error("Please offer the right number of parameters.")
	}

	userId := args[0]
	tokenId := args[1]
	newCost := args[2]

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

	err = json.Unmarshal(JudgeOwnToken(stub, []string{userId, token.Id}).Payload, &token)
	if err != nil {
		return shim.Error("The user does not own the token")
	}

	var resource lib.Resource
	err = json.Unmarshal(QueryResource(stub, []string{token.Asset}).Payload, &resource)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}

	if val, err := strconv.ParseFloat(newCost, 64); err != nil {
		return shim.Error(fmt.Sprintf("积分参数格式转换出错: %s", err))
	} else {
		resource.Cost = val
		token.Value = val
	}

	if err := utils.WriteLedger(token, stub, lib.TokenKey, []string{token.Id}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	if err = utils.WriteLedger(token, stub, lib.TokenKey, []string{token.Asset, token.Time}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	if err := utils.WriteLedger(resource, stub, lib.ResourceKey, []string{resource.Id}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}

	tokenByte, err := json.Marshal(token)
	return shim.Success(tokenByte)

}

func RefuseTransferToken(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 2 {
		return shim.Error("Please offer the right number of parameters.")
	}

	tokenId := args[0]
	userId := args[1]

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

	err = json.Unmarshal(JudgeOwnToken(stub, []string{userId, token.Id}).Payload, &token)
	if err != nil {
		return shim.Error("The user does not own the token")
	}

	if token.NotForSale {
		return shim.Error("This token is not for sale.")
	}
	if token.Owner != userId {
		return shim.Error("You are not the owner of this token, so the change is illegal.")
	}
	if token.Bid == "" {
		return shim.Error("This token has not been bid yet.")
	}

	token.Bid = ""

	if err := utils.WriteLedger(token, stub, lib.TokenKey, []string{token.Id}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	if err = utils.WriteLedger(token, stub, lib.TokenKey, []string{token.Asset, token.Time}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}

	tokenByte, err := json.Marshal(token)
	if err != nil {
		return shim.Error(fmt.Sprintf("序列化成功创建的信息出错: %s", err))
	}
	return shim.Success(tokenByte)
}

func ChangeTokenSaleState(stub shim.ChaincodeStubInterface, args []string) pb.Response {
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

	err = json.Unmarshal(JudgeOwnToken(stub, []string{userId, token.Id}).Payload, &token)
	if err != nil {
		return shim.Error("The user does not own the token")
	}

	token.NotForSale = !token.NotForSale

	if err := utils.WriteLedger(token, stub, lib.TokenKey, []string{token.Id}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	if err := utils.WriteLedger(token, stub, lib.TokenKey, []string{token.Asset, token.Time}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}

	tokenByte, err := json.Marshal(token)
	if err != nil {
		return shim.Error(fmt.Sprintf("序列化成功创建的信息出错: %s", err))
	}
	return shim.Success(tokenByte)
}

func CreateProject(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 4 {
		return shim.Error("Please offer the right number of parameters.")
	}

	id := args[0]
	name := args[1]
	owner := args[2]
	Time := args[3]
	notes := args[4]

	if id == "" {
		return shim.Error("Please offer the project id.")
	}
	if name == "" {
		return shim.Error("Please offer the name of the project.")
	}
	if owner == "" {
		return shim.Error("Please offer the owner of the project.")
	}
	if Time == "" {
		return shim.Error("Please offer the time.")
	}

	var project lib.Project
	err := json.Unmarshal(QueryProject(stub, []string{id}).Payload, &project)
	if err == nil {
		return shim.Error("The project has existed")
	}

	NewProject := &lib.Project{
		Id:    id,
		Name:  name,
		Owner: owner,
		Time:  args[3],
		Notes: notes,
		Use:   "",
	}

	if err := utils.WriteLedger(NewProject, stub, lib.ProjectKey, []string{id}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}

	var account lib.User
	err = json.Unmarshal(QueryAccount(stub, []string{owner}).Payload, &account)
	if err != nil {
		return shim.Error(fmt.Sprintf("用户信息验证失败%s", err))
	}
	account.Projects = append(account.Projects, id)
	if err := utils.WriteLedger(account, stub, lib.UserKey, []string{account.Id}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}

	NewProjectByte, err := json.Marshal(NewProject)
	if err != nil {
		return shim.Error(fmt.Sprintf("序列化成功创建的信息出错: %s", err))
	}
	return shim.Success(NewProjectByte)
}

func AddResourceToProject(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 3 {
		return shim.Error("Please offer the right number of parameters.")
	}

	userId := args[0]
	tokenId := args[1]
	projectId := args[2]

	if tokenId == "" {
		return shim.Error("Please offer the token id.")
	}
	if userId == "" {
		return shim.Error("Please offer the user id.")
	}
	if projectId == "" {
		return shim.Error("Please offer the project id.")
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
	var project lib.Project
	err = json.Unmarshal(QueryProject(stub, []string{projectId}).Payload, &project)
	if err != nil {
		return shim.Error("The project does not exist.")
	}

	err = json.Unmarshal(JudgeOwnToken(stub, []string{userId, tokenId}).Payload, &token)
	if err != nil {
		return shim.Error("The user does not own the token")
	}

	project.Bid = append(project.Bid, token.Asset)
	if err := utils.WriteLedger(project, stub, lib.ProjectKey, []string{projectId}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}

	projectByte, err := json.Marshal(project)
	if err != nil {
		return shim.Error(fmt.Sprintf("序列化成功创建的信息出错: %s", err))
	}
	return shim.Success(projectByte)
}

func ChooseResourceForProject(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 3 {
		return shim.Error("Please offer the right number of parameters.")
	}

	userId := args[0]
	resourceId := args[1]
	projectId := args[2]

	if resourceId == "" {
		return shim.Error("Please offer the resource id.")
	}
	if userId == "" {
		return shim.Error("Please offer the user id.")
	}
	if projectId == "" {
		return shim.Error("Please offer the project id.")
	}

	var user lib.User
	err := json.Unmarshal(QueryAccount(stub, []string{userId}).Payload, &user)
	if err != nil {
		return shim.Error("The user does not exist.")
	}
	var resource lib.Resource
	err = json.Unmarshal(QueryResource(stub, []string{resourceId}).Payload, &resource)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	var project lib.Project
	err = json.Unmarshal(QueryProject(stub, []string{projectId}).Payload, &project)
	if err != nil {
		return shim.Error("The project does not exist.")
	}

	err = json.Unmarshal(JudgeOwnProject(stub, []string{userId, projectId}).Payload, &project)
	if err != nil {
		return shim.Error("The user does not upload the project.")
	}

	for _, v := range project.Bid {
		if v == resourceId {
			project.Use = resourceId
			if err := utils.WriteLedger(project, stub, lib.ProjectKey, []string{projectId}); err != nil {
				return shim.Error(fmt.Sprintf("%s", err))
			}
			projectByte, err := json.Marshal(project)
			if err != nil {
				return shim.Error(fmt.Sprintf("序列化成功创建的信息出错: %s", err))
			}
			return shim.Success(projectByte)
		}
	}

	return shim.Error("The project can't choose this resource.")

}

func RecordDeal(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 5 {
		return shim.Error("Please offer the right number of parameters.")
	}

	sell_id := args[0]
	buy_id := args[1]
	rescource_id := args[2]
	cost := args[3] //多少买就必须多少走
	Type := args[4]

	if sell_id == "" {
		return shim.Error("Please offer the sell_id.")
	}
	if buy_id == "" {
		return shim.Error("Please offer the buy_id.")
	}
	if rescource_id == "" {
		return shim.Error("Please offer Resource_id.")
	}
	if cost == "" {
		return shim.Error("Please offer Cost.")
	}
	if Type == "" {
		return shim.Error("Please offer type of deal.")
	}

	if sell_id == buy_id {
		return shim.Error("The id of seller and buyer can't be the same.")
	}

	if Type != "share" && Type != "recommend" && Type != "transfer" && Type != "give" && Type != "redeem" {
		return shim.Error("The type of trade is illegal.")
	}

	value, err := strconv.ParseFloat(cost, 64)
	if err != nil {
		return shim.Error("The value of the deal is not legal.")
	}

	NewDeal := &lib.TokenDeal{
		Sell_id:     sell_id,
		Buy_id:      buy_id,
		Resource_id: rescource_id,
		Cost:        value,
		Type:        args[4],
		Time:        time.Now().String(),
	}

	if err := utils.WriteLedger(NewDeal, stub, lib.DealKey, []string{NewDeal.Resource_id, NewDeal.Sell_id, NewDeal.Buy_id, Type, time.Now().String()}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	if err := utils.WriteLedger(NewDeal, stub, lib.DealKey, []string{NewDeal.Sell_id, NewDeal.Resource_id, NewDeal.Buy_id, Type, time.Now().String()}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	if err := utils.WriteLedger(NewDeal, stub, lib.DealKey, []string{NewDeal.Buy_id, NewDeal.Resource_id, NewDeal.Sell_id, Type, time.Now().String()}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}

	NewDealByte, err := json.Marshal(NewDeal)
	if err != nil {
		return shim.Error(fmt.Sprintf("序列化成功创建的信息出错: %s", err))
	}
	return shim.Success(NewDealByte)

}

func CreateCoinRecord(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 6 {
		return shim.Error("Please offer the right number of parameters.")
	}

	user := args[0]
	direction := args[1]
	Type := args[2]
	name := args[3] //多少买就必须多少走
	value := args[4]
	Time := args[5]

	cost, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return shim.Error("The value of the record is not legal.")
	}

	NewDeal := &lib.CoinRecord{
		User:      user,
		Direction: direction,
		Type:      Type,
		Name:      name,
		Value:     cost,
		Time:      Time,
	}

	if err := utils.WriteLedger(NewDeal, stub, lib.CoinRecordKey, []string{NewDeal.User, NewDeal.Direction, NewDeal.Type, NewDeal.Name, time.Now().String()}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}

	NewDealByte, err := json.Marshal(NewDeal)
	if err != nil {
		return shim.Error(fmt.Sprintf("序列化成功创建的信息出错: %s", err))
	}
	return shim.Success(NewDealByte)

}
