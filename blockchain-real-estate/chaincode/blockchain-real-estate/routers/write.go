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
	"reflect"

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
	Id := args[0]
	Hash := args[1]
	Owner := args[2]
	Time := args[3]
	Cost := args[4]
	GetScore := args[5]
	if reflect.TypeOf(Id).Name() != "string" || reflect.TypeOf(Hash).Name() != "string" || reflect.TypeOf(Owner).Name() != "string" ||
		reflect.TypeOf(Time).Name() != "string" || reflect.TypeOf(Time).Name() != "string" || reflect.TypeOf(Cost).Name() != "string" ||
		reflect.TypeOf(GetScore).Name() != "string" {
		return shim.Error("Please check the type of input")
	}
	if Id == "" || Hash == "" || Owner == "" || Time == "" || Cost == "" {
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
	id := args[0]
	assets := args[1]
	notforsale := args[2]
	owner := args[3]
	bid := args[4]   //一个token的bid可以是没有的，也可以有
	notes := args[5] //对token的描述可以有也可以没有
	value := args[6]
	createrId := args[7]
	if reflect.TypeOf(id).Name() != "string" || reflect.TypeOf(assets).Name() != "string" || reflect.TypeOf(notforsale).Name() != "string" ||
		reflect.TypeOf(owner).Name() != "string" || reflect.TypeOf(bid).Name() != "string" || reflect.TypeOf(notes).Name() != "string" ||
		reflect.TypeOf(value).Name() != "string" || reflect.TypeOf(bid).Name() != "string" {
		shim.Error("Please check the type of input.")
	}
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
	id := args[0]
	bidder := args[1] //投标人的id
	if reflect.TypeOf(id).Name() != "string" || reflect.TypeOf(bidder).Name() != "string" {
		return shim.Error("Please check the type of input.")
	}
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
	id := args[0] //token的id
	if reflect.TypeOf(id).Name() != "string" {
		return shim.Error("Please check the type of input.")
	}
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

	oldResources := user.Upload
	var newResources []string
	var newControl []string
	for _, v := range oldResources {
		if v != id {
			newResources = append(newResources, v)
		}
	}
	for _, v := range user.Buy {
		if v != id {
			newControl = append(newControl, v)
		}
	}
	user.Upload = newResources
	user.Buy = newControl
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

func TransferToken(stub shim.ChaincodeStubInterface, args []string) pb.Response { //代币的所有者通过此链码将自己的代币转给其他人
	tokenId := args[0]
	userId := args[1]
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
	err = json.Unmarshal(CreateToken(stub, []string{"new" + tokenId, token.Asset, newNotForSale, token.Bid, "",
		token.Notes, newValue, token.Bid}).Payload, &newToken)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	err = json.Unmarshal(QueryAccount(stub, []string{token.Bid}).Payload, &bidder)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	bidder.Score -= token.Value
	bidder.Control = append(bidder.Control, newToken.Id)
	bidder.Buy = append(bidder.Buy, newToken.Asset) //可能会有重复，但不影响查询
	if err := utils.WriteLedger(bidder, stub, lib.UserKey, []string{bidder.Id}); err != nil {
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

	NewTokenByte, err := json.Marshal(newToken)
	if err != nil {
		return shim.Error(fmt.Sprintf("序列化成功创建的信息出错: %s", err))
	}
	return shim.Success(NewTokenByte)
}
