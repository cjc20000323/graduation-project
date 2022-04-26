package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/togettoyou/blockchain-real-estate/chaincode/blockchain-real-estate/lib"
	"github.com/togettoyou/blockchain-real-estate/chaincode/blockchain-real-estate/routers"
	"github.com/togettoyou/blockchain-real-estate/chaincode/blockchain-real-estate/utils"
	"time"
)

type BlockChainRealEstate struct {
}

//链码初始化
func (t *BlockChainRealEstate) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("链码初始化")
	timeLocal, err := time.LoadLocation("Asia/Chongqing")
	if err != nil {
		return shim.Error(fmt.Sprintf("时区设置失败%s", err))
	}
	time.Local = timeLocal
	//初始化默认数据
	NewUser := &lib.User{
		Id:    "Admin_1",
		Score: float64(20),
	}
	// 写入账本
	if err := utils.WriteLedger(NewUser, stub, lib.UserKey, []string{NewUser.Id}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}

	return shim.Success(nil)
}

//实现Invoke接口调用智能合约
func (t *BlockChainRealEstate) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	funcName, args := stub.GetFunctionAndParameters()
	switch funcName {
	case "createUser":
		return routers.CreateUser(stub, args)
	case "uploadResource":
		return routers.UploadResource(stub, args)
	case "queryResource":
		return routers.QueryResource(stub, args)
	case "queryAccount":
		return routers.QueryAccount(stub, args)
	case "queryAllAccount":
		return routers.QueryAllAccount(stub, args)
	case "queryUpload":
		return routers.QueryUpload(stub, args)
	case "createDeal":
		return routers.CreateDeal(stub, args)
	case "queryBuyResources":
		return routers.QueryBuyResources(stub, args)
	case "queryDealResource":
		return routers.QueryDealResource(stub, args)
	case "queryAllResource":
		return routers.QueryAllResource(stub, args)
	case "judgeOwnResource":
		return routers.JudgeOwnResource(stub, args)
	case "judgeBuyResource":
		return routers.JudgeBuyResource(stub, args)
	case "judgeShareResource":
		return routers.JudgeShareResource(stub, args)
	case "judgeOwnToken":
		return routers.JudgeOwnToken(stub, args)
	case "judgeShareToken":
		return routers.JudgeShareToken(stub, args)
	case "createToken":
		return routers.CreateToken(stub, args)
	case "readToken":
		return routers.ReadToken(stub, args)
	case "deleteToken":
		return routers.DeleteToken(stub, args)
	case "bidToken":
		return routers.BidToken(stub, args)
	case "transferToken":
		return routers.TransferToken(stub, args)
	case "getUserTokenList":
		return routers.GetUserTokenList(stub, args)
	case "shareToken":
		return routers.ShareToken(stub, args)
	case "redeemToken":
		return routers.RedeemToken(stub, args)
	case "changeResourceScore":
		return routers.ChangeResourceScore(stub, args)
	case "refuseTransferToken":
		return routers.RefuseTransferToken(stub, args)
	case "changeTokenSaleState":
		return routers.ChangeTokenSaleState(stub, args)
	case "queryUserBiddenToken":
		return routers.QueryUserBiddenToken(stub, args)
	case "queryUserResource":
		return routers.QueryUserResource(stub, args)
	case "giveToken":
		return routers.GiveToken(stub, args)
	case "judgeLendToken":
		return routers.JudgeLendToken(stub, args)
	case "queryUserLendToken":
		return routers.QueryUserLendToken(stub, args)
	case "queryUserSharedToken":
		return routers.QueryUserSharedToken(stub, args)
	case "queryTokenShare":
		return routers.QueryTokenShare(stub, args)
	case "judgeOwnProject":
		return routers.JudgeOwnProject(stub, args)
	case "queryProject":
		return routers.QueryProject(stub, args)
	case "createProject":
		return routers.CreateProject(stub, args)
	case "queryProjectResource":
		return routers.QueryProjectResource(stub, args)
	case "addResourceToProject":
		return routers.AddResourceToProject(stub, args)
	case "chooseResourceForProject":
		return routers.ChooseResourceForProject(stub, args)
	case "queryUserProject":
		return routers.QueryUserProject(stub, args)
	case "recordDeal":
		return routers.RecordDeal(stub, args)
	case "queryDeal":
		return routers.QueryDeal(stub, args)
	case "queryAllDealSum":
		return routers.QueryAllDealSum(stub, args)
	case "queryALLResourceSum":
		return routers.QueryAllResourceSum(stub, args)
	case "queryAllUserSum":
		return routers.QueryAllUserSum(stub, args)
	default:
		return shim.Error(fmt.Sprintf("没有该功能: %s", funcName))
	}
}

func main() {
	err := shim.Start(new(BlockChainRealEstate))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
