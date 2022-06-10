/**
 * @Author: 夜央 Oh oh oh oh oh oh (https://github.com/togettoyou)
 * @Email: zoujh99@qq.com
 * @Date: 2020/3/4 1:51 下午
 * @Description: 读写账本工具
 */
package utils

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/togettoyou/blockchain-real-estate/chaincode/blockchain-real-estate/lib"
)

//写入账本
func WriteLedger(obj interface{}, stub shim.ChaincodeStubInterface, objectType string, keys []string) error {
	//创建复合主键
	var key string
	if val, err := stub.CreateCompositeKey(objectType, keys); err != nil {
		return errors.New(fmt.Sprintf("%s-创建复合主键出错 %s", objectType, err))
	} else {
		key = val
	}
	// 序列化对象
	bytes, err := json.Marshal(obj)
	if err != nil {
		return errors.New(fmt.Sprintf("%s-序列化json数据失败出错: %s", objectType, err))
	}
	//写入区块链账本
	if err := stub.PutState(key, bytes); err != nil {
		return errors.New(fmt.Sprintf("%s-写入区块链账本出错: %s", objectType, err))
	}
	return nil
}

//删除账本  没用
func DelLedger(stub shim.ChaincodeStubInterface, objectType string, keys []string) error {
	//创建复合主键
	var key string
	if val, err := stub.CreateCompositeKey(objectType, keys); err != nil {
		return errors.New(fmt.Sprintf("%s-创建复合主键出错 %s", objectType, err))
	} else {
		key = val
	}
	//写入区块链账本
	if err := stub.DelState(key); err != nil {
		return errors.New(fmt.Sprintf("%s-删除区块链账本出错: %s", objectType, err))
	}
	return nil
}

//根据复合主键查询数据(适合获取全部，多个，单个数据)
//将keys拆分查询
func GetStateByPartialCompositeKeys(stub shim.ChaincodeStubInterface, objectType string, keys []string) (results [][]byte, err error) {
	if len(keys) == 0 {
		// 传入的keys长度为0，则查找并返回所有数据
		// 通过主键从区块链查找相关的数据，相当于对主键的模糊查询
		resultIterator, err := stub.GetStateByPartialCompositeKey(objectType, keys)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("%s-获取全部数据出错: %s", objectType, err))
		}
		defer resultIterator.Close()

		//检查返回的数据是否为空，不为空则遍历数据，否则返回空数组
		for resultIterator.HasNext() {
			val, err := resultIterator.Next()
			if err != nil {
				return nil, errors.New(fmt.Sprintf("%s-返回的数据出错: %s", objectType, err))
			}

			results = append(results, val.GetValue())
		}
	} else {
		// 传入的keys长度不为0，查找相应的数据并返回
		for _, v := range keys {
			// 创建组合键
			key, err := stub.CreateCompositeKey(objectType, []string{v})
			if err != nil {
				return nil, errors.New(fmt.Sprintf("%s-创建组合键出错: %s", objectType, err))
			}
			// 从账本中获取数据
			bytes, err := stub.GetState(key)
			if err != nil {
				return nil, errors.New(fmt.Sprintf("%s-获取数据出错: %s", objectType, err))
			}

			if bytes != nil {
				results = append(results, bytes)
			}
		}
	}

	return results, nil
}

func GetHistoryForKeys(stub shim.ChaincodeStubInterface, objectType string, keys []string) (results []byte, err error) {
	type AuditHistory struct {
		TxId  string       `json:"txId"`
		Value lib.Resource `json:"value"`
	}
	var history []AuditHistory
	var resource lib.Resource
	for _, v := range keys { //参数其实就一个
		// 创建组合键
		key, err := stub.CreateCompositeKey(objectType, []string{v})
		if err != nil {
			return nil, errors.New(fmt.Sprintf("%s创建组合键出错: %s", objectType, err))
		}
		// 从账本中获取数据

		resultsIterator, err := stub.GetHistoryForKey(key)

		defer resultsIterator.Close()
		for resultsIterator.HasNext() {
			historyData, err := resultsIterator.Next()
			if err != nil {
				return nil, errors.New(fmt.Sprintf("%s迭代出错: %s", objectType, err))
			}
			var tx AuditHistory
			tx.TxId = historyData.TxId
			json.Unmarshal(historyData.Value, &resource)
			if historyData.Value == nil { //marble has been deleted
				var emptyResource lib.Resource
				tx.Value = emptyResource //copy nil marble
			} else {
				json.Unmarshal(historyData.Value, &resource) //un stringify it aka JSON.parse()
				tx.Value = resource                          //copy marble over
			}
			history = append(history, tx)
		}
	}
	historyAsBytes, _ := json.Marshal(history) //convert to array of bytes
	return historyAsBytes, nil
}

func GetHistoryForKeysToken(stub shim.ChaincodeStubInterface, objectType string, keys []string) (results []byte, err error) {
	type AuditHistory struct {
		TxId  string    `json:"txId"`
		Value lib.Token `json:"value"`
	}
	var history []AuditHistory
	var resource lib.Token
	for _, v := range keys { //参数其实就一个
		// 创建组合键
		key, err := stub.CreateCompositeKey(objectType, []string{v})
		if err != nil {
			return nil, errors.New(fmt.Sprintf("%s创建组合键出错: %s", objectType, err))
		}
		// 从账本中获取数据

		resultsIterator, err := stub.GetHistoryForKey(key)

		defer resultsIterator.Close()
		for resultsIterator.HasNext() {
			historyData, err := resultsIterator.Next()
			if err != nil {
				return nil, errors.New(fmt.Sprintf("%s迭代出错: %s", objectType, err))
			}
			var tx AuditHistory
			tx.TxId = historyData.TxId
			json.Unmarshal(historyData.Value, &resource)
			if historyData.Value == nil { //marble has been deleted
				var emptyResource lib.Token
				tx.Value = emptyResource //copy nil marble
			} else {
				json.Unmarshal(historyData.Value, &resource) //un stringify it aka JSON.parse()
				tx.Value = resource                          //copy marble over
			}
			history = append(history, tx)
		}
	}
	historyAsBytes, _ := json.Marshal(history) //convert to array of bytes
	return historyAsBytes, nil
}

func GetHistoryForKeysDeal(stub shim.ChaincodeStubInterface, objectType string, keys []string) (results []byte, err error) {
	type AuditHistory struct {
		TxId  string   `json:"txId"`
		Value lib.Deal `json:"value"`
	}
	var history []AuditHistory
	var resource lib.Deal
	for _, v := range keys { //参数其实就一个
		// 创建组合键
		key, err := stub.CreateCompositeKey(objectType, []string{v})
		if err != nil {
			return nil, errors.New(fmt.Sprintf("%s创建组合键出错: %s", objectType, err))
		}
		// 从账本中获取数据

		resultsIterator, err := stub.GetHistoryForKey(key)

		defer resultsIterator.Close()
		for resultsIterator.HasNext() {
			historyData, err := resultsIterator.Next()
			if err != nil {
				return nil, errors.New(fmt.Sprintf("%s迭代出错: %s", objectType, err))
			}
			var tx AuditHistory
			tx.TxId = historyData.TxId
			json.Unmarshal(historyData.Value, &resource)
			if historyData.Value == nil { //marble has been deleted
				var emptyResource lib.Deal
				tx.Value = emptyResource //copy nil marble
			} else {
				json.Unmarshal(historyData.Value, &resource) //un stringify it aka JSON.parse()
				tx.Value = resource                          //copy marble over
			}
			history = append(history, tx)
		}
	}
	historyAsBytes, _ := json.Marshal(history) //convert to array of bytes
	return historyAsBytes, nil
}

func GetHistoryForKeysUser(stub shim.ChaincodeStubInterface, objectType string, keys []string) (results []byte, err error) {
	type AuditHistory struct {
		TxId  string   `json:"txId"`
		Value lib.User `json:"value"`
	}
	var history []AuditHistory
	var resource lib.User
	for _, v := range keys { //参数其实就一个
		// 创建组合键
		key, err := stub.CreateCompositeKey(objectType, []string{v})
		if err != nil {
			return nil, errors.New(fmt.Sprintf("%s创建组合键出错: %s", objectType, err))
		}
		// 从账本中获取数据

		resultsIterator, err := stub.GetHistoryForKey(key)

		defer resultsIterator.Close()
		for resultsIterator.HasNext() {
			historyData, err := resultsIterator.Next()
			if err != nil {
				return nil, errors.New(fmt.Sprintf("%s迭代出错: %s", objectType, err))
			}
			var tx AuditHistory
			tx.TxId = historyData.TxId
			json.Unmarshal(historyData.Value, &resource)
			if historyData.Value == nil { //marble has been deleted
				var emptyResource lib.User
				tx.Value = emptyResource //copy nil marble
			} else {
				json.Unmarshal(historyData.Value, &resource) //un stringify it aka JSON.parse()
				tx.Value = resource                          //copy marble over
			}
			history = append(history, tx)
		}
	}
	historyAsBytes, _ := json.Marshal(history) //convert to array of bytes
	return historyAsBytes, nil
}

func GetHistoryForKeysResource(stub shim.ChaincodeStubInterface, objectType string, keys []string) (results []byte, err error) {
	type AuditHistory struct {
		TxId  string       `json:"txId"`
		Value lib.Resource `json:"value"`
	}
	var history []AuditHistory
	var resource lib.Resource
	for _, v := range keys { //参数其实就一个
		// 创建组合键
		key, err := stub.CreateCompositeKey(objectType, []string{v})
		if err != nil {
			return nil, errors.New(fmt.Sprintf("%s创建组合键出错: %s", objectType, err))
		}
		// 从账本中获取数据

		resultsIterator, err := stub.GetHistoryForKey(key)

		defer resultsIterator.Close()
		for resultsIterator.HasNext() {
			historyData, err := resultsIterator.Next()
			if err != nil {
				return nil, errors.New(fmt.Sprintf("%s迭代出错: %s", objectType, err))
			}
			var tx AuditHistory
			tx.TxId = historyData.TxId
			json.Unmarshal(historyData.Value, &resource)
			if historyData.Value == nil { //marble has been deleted
				var emptyResource lib.Resource
				tx.Value = emptyResource //copy nil marble
			} else {
				json.Unmarshal(historyData.Value, &resource) //un stringify it aka JSON.parse()
				tx.Value = resource                          //copy marble over
			}
			history = append(history, tx)
		}
	}
	historyAsBytes, _ := json.Marshal(history) //convert to array of bytes
	return historyAsBytes, nil
}

//根据复合主键查询数据(适合获取全部或指定的数据)
func GetStateByPartialCompositeKeys2(stub shim.ChaincodeStubInterface, objectType string, keys []string) (results [][]byte, err error) {
	// 通过主键从区块链查找相关的数据，相当于对主键的模糊查询
	resultIterator, err := stub.GetStateByPartialCompositeKey(objectType, keys)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%s-获取全部数据出错: %s", objectType, err))
	}
	defer resultIterator.Close()

	//检查返回的数据是否为空，不为空则遍历数据，否则返回空数组
	for resultIterator.HasNext() {
		val, err := resultIterator.Next()
		if err != nil {
			return nil, errors.New(fmt.Sprintf("%s-返回的数据出错: %s", objectType, err))
		}

		results = append(results, val.GetValue())
	}
	return results, nil
}
