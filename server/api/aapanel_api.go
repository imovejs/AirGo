package api

import (
	"AirGo/global"
	"AirGo/model"
	"AirGo/service"
	"AirGo/utils/response"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 当前节点设置
func SSNodeInfo(ctx *gin.Context) {
	//验证key
	if global.Server.System.MuKey != ctx.Query("muKey") {
		return
	}
	//节点号
	nodeID := ctx.Param("nodeID")
	nodeIDInt, _ := strconv.Atoi(nodeID)
	nodeInfo, err := service.SSNodeInfo(nodeIDInt)
	if err != nil {
		response.SSUsersFail(ctx)
		return
	}
	response.SSUsersOK(nodeInfo, ctx)

}

// 可连接的用户
func SSUsers(ctx *gin.Context) {
	//验证key
	if global.Server.System.MuKey != ctx.Query("muKey") {
		return
	}
	//节点号
	nodeID := ctx.Query("node_id")
	nodeIDInt, _ := strconv.Atoi(nodeID)
	//节点属于哪些goods
	goods, err := service.FindGoodsByNodeID(nodeIDInt)
	if err != nil {
		return
	}
	//goods属于哪些用户
	users, err := service.FindUsersByGoods(&goods)
	if err != nil {
		response.SSUsersFail(ctx)
		return
	}
	//fmt.Println("users", users)
	response.SSUsersOK(users, ctx)
}

// 上报用户的流量使用情况
func SSUsersTraffic(ctx *gin.Context) {
	//验证key
	if global.Server.System.MuKey != ctx.Query("muKey") {
		return
	}
	node_id_str := ctx.Query("node_id")
	node_id, _ := strconv.Atoi(node_id_str)
	//fmt.Println("用户的流量使用情况node_id:", node_id)

	var trafficReq model.TrafficReq
	err := ctx.ShouldBind(&trafficReq)
	if err != nil {
		return
	}
	//用户的流量使用情况: {"data":[{"user_id":1,"u":445782,"d":1757834}]}
	//fmt.Println("用户的流量使用情况trafficReq:", trafficReq.Data)
	response.SSUsersOK("ok", ctx)

	var userIds []int
	var userArr []model.User
	var trafficLog = model.TrafficLog{
		NodeID: node_id,
	}
	for _, v := range trafficReq.Data {
		//每个用户流量
		var user model.User
		userIds = append(userIds, v.UserID)
		user.ID = v.UserID
		user.SubscribeInfo.U = v.U
		user.SubscribeInfo.D = v.D
		userArr = append(userArr, user)
		//该节点总流量
		trafficLog.U = trafficLog.U + v.U
		trafficLog.D = trafficLog.D + v.D
	}
	//插入流量统计统计
	err = service.NewTrafficLog(&trafficLog)
	if err != nil {
		return
	}
	//更新用户流量信息
	//fmt.Println("更新用户流量信息userArr:", userArr)
	if len(userArr) == 0 {
		return
	}
	err = service.UpdateUserTrafficInfo(userArr, userIds)
	if err != nil {
		fmt.Println("更新用户流量信息err", err)
		return
	}
}

// 上报用户的当前在线IP
func SSUsersAliveIP(ctx *gin.Context) {

}