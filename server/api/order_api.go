package api

import (
	"AirGo/model"
	"AirGo/service"
	"AirGo/utils/response"
	"fmt"

	"github.com/gin-gonic/gin"
)

// 获取全部订单，分页获取
func GetAllOrder(ctx *gin.Context) {
	var params model.QueryParamsWithDate
	err := ctx.ShouldBind(&params)
	res, err := service.GetAllOrder(&params)
	if err != nil {
		fmt.Println("获取全部订单 err:", err)
		response.Fail("订单获取错误"+err.Error(), nil, ctx)
		return
	}
	//fmt.Println("获取全部订单res:", res)
	response.OK("全部订单获取成功", res, ctx)
}

// 获取用户订单by user id
func GetOrderByUserID(ctx *gin.Context) {
	uID, ok := ctx.Get("uID")
	if !ok {
		response.Fail("订单获取，uID参数错误", nil, ctx)
		return
	}
	uIDInt := uID.(int)
	var params model.PaginationParams
	err := ctx.ShouldBind(&params)
	res, err := service.GetOrderByUserID(uIDInt, &params)
	if err != nil {
		fmt.Println("获取订单 err:", err)
		response.Fail("订单获取错误"+err.Error(), nil, ctx)
		return
	}
	response.OK("订单获取成功", res, ctx)
}

// 完成未支付订单
func CompletedOrder(ctx *gin.Context) {
	var order model.Orders
	err := ctx.ShouldBind(&order)
	if err != nil {
		response.Fail("完成未支付订单参数错误"+err.Error(), nil, ctx)
		return
	}
	order.TradeStatus = "completed"   //更新数据库订单状态,自定义结束状态completed
	err = service.UpdateOrder(&order) //更新数据库状态
	if err != nil {
		response.Fail("更新数据库状态错误"+err.Error(), nil, ctx)
		return
	}
	err = service.UpdateUserSubscribe(&order) //更新用户订阅信息
	if err != nil {
		response.Fail("更新用户订阅信息错误"+err.Error(), nil, ctx)
		return
	}
	response.OK("完成未支付订单成功", nil, ctx)

}