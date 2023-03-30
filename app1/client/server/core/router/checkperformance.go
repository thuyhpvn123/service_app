package router

import (
	"gitlab.com/meta-node/client/models"
	"gitlab.com/meta-node/client/utils"
)
func (caller *CallData)checkperformance()*Result{
	result := &Result{
		Success: true,
		Data:map[string]interface{}{
			"device-info": 1,
			"device-performance": "",
		},
	
	}
	header := models.Header{ Success:true,Data: result.Data}
	kq1 := utils.NewResultTransformer(header)

	go caller.sentToClient("desktop","check-performance", false,kq1)

	return result
}
// func (caller *CallData)getPersonInfo(data ...interface{})*Result{
// 	result := &Result{
// 		Success: true,
// 		Data:"hard code",
	
// 	}
// 	header := models.Header{ Success:true,Count: 0, Data: result.Data}
// 	kq := utils.NewResultTransformer(header)

// 	go caller.sentToClient("get-person-info", true,kq)

// 	return result
// }
// func (caller *CallData)getMySetting(data ...interface{})*Result{
// 	result := &Result{
// 		Success: true,
// 		Data:"hard code",
	
// 	}
// 	go caller.sentToClient("get-my-setting", true,result.Data)

// 	return result
// }

