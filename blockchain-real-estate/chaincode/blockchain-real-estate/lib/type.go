/**
 * @Author: 夜央 Oh oh oh oh oh oh (https://github.com/togettoyou)
 * @Email: zoujh99@qq.com
 * @Date: 2020/3/4 2:00 下午
 * @Description: 定义的数据结构体、常量
 */
package lib

import (
	"time"
)

type Resource struct {
	//ObjectType string    `json:"docType"` //用于CouchDB
	Id   string `json:"id"`   //资源在关系型数据库中的id，从id中可知类型  paper_1223131
	Hash string `json:"Hash"` //文件在IPFS系统中的Hash值
	//Uploader   string    `json:"Uploader"`//标记卖方id
	Owner string `json:"Owner"`
	Time  string `json:"Time"` //标记上链时间
	//State      string     `json:"State"`//只有未有对应解决方案的项目需求为true   没用
	Cost float64 `json:"Cost"` //交易需要话费的积分，可设置为零
	//type
}

type User struct {
	//ObjectType  string        `json:"docType"` //用于CouchDB
	Id      string   `json:"id"`      //关系型数据库id
	Upload  []string `json:"Upload"`  //上传的资源对象，json格式
	Buy     []string `json:"Buy"`     //购买的资源对象
	Control []string `json:"Control"` //控制的代币
	Share   []string `json:"Share"`   //他人分享的代币
	Score   float64  `json:"Score"`   //积分
}

//其实并不需要交易记录表
type Deal struct {
	//ObjectType  string        `json:"docType"` //用于CouchDB
	//Id          string        `json:"id"`     //联合主键  三个
	Sell_id      string    `json:"Sell_id"`
	Buy_id       string    `json:"Buy_id"`
	Rescource_id string    `json:"Rescource_id"` //上传的资源对象，json格式
	Cost         float64   `json:"Cost"`         //积分
	Time         time.Time `json:"Time"`         //交易完成时间（上链时间更准）
}

type Token struct {
	Id         string  `json:"id"`
	Asset      string  `json:"asset"`
	NotForSale bool    `json:"not_for_sale"`
	Owner      string  `json:"owner"`
	Bid        string  `json:"bid"`
	Notes      string  `json:"notes"`
	Value      float64 `json:"value"`
}

//账户，虚拟管理员和若干业主账号
type Account struct {
	AccountId string  `json:"accountId"` //账号ID
	UserName  string  `json:"userName"`  //账号名
	Balance   float64 `json:"balance"`   //余额
}

//var SellingStatusConstant = func() map[string]string {
//	return map[string]string{
//		"saleStart": "销售中", //正在销售状态,等待买家光顾
//		"cancelled": "已取消", //被卖家取消销售或买家退款操作导致取消
//		"expired":   "已过期", //销售期限到期
//		"delivery":  "交付中", //买家买下并付款,处于等待卖家确认收款状态,如若卖家未能确认收款，买家可以取消并退款
//		"done":      "完成",  //卖家确认接收资金，交易完成
//	}
//}

////买家参与销售
////销售对象不能是买家发起的
////Buyer和CreateTime作为复合键,保证可以通过buyer查询到名下所有参与的销售
//type SellingBuy struct {
//	Buyer      string  `json:"buyer"`      //参与销售人、买家(买家AccountId)
//	CreateTime string  `json:"createTime"` //创建时间
//	Selling    Selling `json:"selling"`    //销售对象
//}

const (
	UserKey     = "user-key"
	ResourceKey = "resource-key"
	DealKey     = "deal-key"

	AccountKey = "account-key"
	TokenKey   = "token-key"
	//RealEstateKey      = "real-estate-key"
	//SellingKey         = "selling-key"
	//SellingBuyKey      = "selling-buy-key"
	//DonatingKey        = "donating-key"
	//DonatingGranteeKey = "donating-grantee-key"
)
