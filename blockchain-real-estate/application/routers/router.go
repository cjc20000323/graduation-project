/**
 * @Author: 夜央 Oh oh oh oh oh oh (https://github.com/togettoyou)
 * @Email: zoujh99@qq.com
 * @Date: 2020/3/3 11:17 下午
 * @Description: 路由配置
 */
package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	v1 "github.com/togettoyou/blockchain-real-estate/application/routers/api/v1"
	"net/http"
	"strings"
)

//初始化路由信息
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(Cors())
	//swagger文档
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	apiV1 := r.Group("/api/v1")
	{
		apiV1.POST("/createUser", v1.CreateUser)
		apiV1.POST("/uploadResource", v1.UploadResource)
		apiV1.POST("/createDeal", v1.CreateDeal)

		apiV1.POST("/queryResource", v1.QueryResource)
		apiV1.POST("/queryAccount", v1.QueryAccount)
		apiV1.POST("/queryAllAccount", v1.QueryAllAccount)
		apiV1.POST("/queryUpload", v1.QueryUpload)
		apiV1.POST("/queryBuyResources", v1.QueryBuyResources)

		apiV1.POST("/queryDealResource", v1.QueryDealResource)
		apiV1.POST("/queryAllResource", v1.QueryAllResource)
		apiV1.POST("/uploadfile", v1.Uploadfile)
		apiV1.POST("/createToken", v1.CreateToken)
		apiV1.POST("/bidToken", v1.BidToken)
		apiV1.POST("/deleteToken", v1.DeleteToken)
		apiV1.POST("/giveToken", v1.GiveToken)
		apiV1.POST("/transferToken", v1.TransferToken)
		apiV1.POST("/shareToken", v1.ShareToken)
		apiV1.POST("/redeemToken", v1.RedeemToken)
		apiV1.POST("/changeResourceScore", v1.ChangeResourceScore)
		apiV1.POST("/refuseTransferToken", v1.RefuseTransferToken)
		apiV1.POST("/changeTokenSaleState", v1.ChangeTokenSaleState)
		apiV1.POST("/createProject", v1.CreateProject)
		apiV1.POST("/addResourceToProject", v1.AddResourceToProject)
		apiV1.POST("/chooseResourceForProject", v1.ChooseResourceForProject)
		apiV1.POST("/readToken", v1.ReadToken)
		apiV1.POST("/getUserTokenList", v1.GetUserTokenList)
		apiV1.POST("/queryUserBiddenToken", v1.QueryUserBiddenToken)
		apiV1.POST("/queryUserResource", v1.QueryUserResource)
		apiV1.POST("/queryUserSharedToken", v1.QueryUserSharedToken)
		apiV1.POST("/queryUserLendToken", v1.QueryUserLendToken)
		apiV1.POST("/queryTokenShare", v1.QueryTokenShare)
		apiV1.POST("/queryProject", v1.QueryProject)
		apiV1.POST("/queryUserProject", v1.QueryUserProject)
		apiV1.POST("/queryProjectResource", v1.QueryProjectResource)
		apiV1.POST("/judgeOwnToken", v1.JudgeOwnToken)
		apiV1.POST("/judgeBuyResource", v1.JudgeBuyResource)
		apiV1.POST("/judgeOwnResource", v1.JudgeOwnResource)
		apiV1.POST("/judgeShareResource", v1.JudgeShareResource)
		apiV1.POST("/judgeShareToken", v1.JudgeShareToken)
		apiV1.POST("/judgeLendToken", v1.JudgeLendToken)
		apiV1.POST("/judgeOwnProject", v1.JudgeOwnProject)
		//apiV1.POST("/downloadfile", v1.Downloadfile)下载对应ipfs/hash即可 但是下载不知道文件名，所以必须要求上传是zip文件

	}
	// s
	r.StaticFS("/web", http.Dir("./dist/"))
	return r
}

//允许跨域请求
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method               //请求方法
		origin := c.Request.Header.Get("Origin") //请求头部
		var headerKeys []string                  // 声明请求头keys
		for k, _ := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Origin", "*")                                       // 这是允许访问所有域
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE") //服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
			// header的类型
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			// 允许跨域设置                                                                                                      可以返回其他子段
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar") // 跨域关键设置 让浏览器可以解析
			c.Header("Access-Control-Max-Age", "172800")                                                                                                                                                           // 缓存请求信息 单位为秒
			c.Header("Access-Control-Allow-Credentials", "false")                                                                                                                                                  //  跨域请求是否需要带cookie信息 默认设置为true
			c.Set("content-type", "application/json")                                                                                                                                                              // 设置返回格式是json
		}

		// 放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}
		// 处理请求
		c.Next()
	}
}
