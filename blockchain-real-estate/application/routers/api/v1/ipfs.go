package v1

import (
	"github.com/togettoyou/blockchain-real-estate/application/pkg/app"
	"bytes"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/ipfs/go-ipfs-api"
	"fmt"
	"io/ioutil"

)
type FileGet struct{
	Hash        string        `json:"Hash"`
}

var sh *shell.Shell

func Uploadfile(c *gin.Context) {
	appG := app.Gin{C: c}
	form,err:=c.MultipartForm()
	files:=form.File["files"]
	if err != nil {
		appG.Response(http.StatusBadRequest, "上传文件失败","fail" )
	}

	var resultlist []interface{}
	for _,f:=range files{
		fmt.Println(f.Filename)

		fileContent,_:= f.Open()
		var byteContainer [] byte
		byteContainer = make([]byte,1000000)
		fileContent.Read(byteContainer)

		result :=UploadIPFS(byteContainer)
		resultlist=append(resultlist, result)
	}
	appG.Response(http.StatusOK, "上传文件成功", resultlist)
}

/*
func Downloadfile(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(FileGet)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}

	result:=CatIPFS(body.Hash)
	ioutil.WriteFile("file.pptx",result,777)
	appG.Response(http.StatusOK, "下载文件成功",result)
}

 */


func UploadIPFS(file []byte) string {
	sh = shell.NewShell("localhost:5001")

	hash, err := sh.Add(bytes.NewBuffer(file))
	if err != nil {
		fmt.Println("上传ipfs时错误：", err)
	}
	return hash
}

func CatIPFS(hash string) []byte {
	sh = shell.NewShell("localhost:5001")

	read, err := sh.Cat(hash)
	if err != nil {
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(read)

	return body
}