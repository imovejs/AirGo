import service from "/@/utils/request";

export function useOrderApi(){
    return {
        //获取全部订单
        getAllOrderApi:(data?:object)=>{
            return service({
                url: '/order/getAllOrder',
                method: 'POST',
                data
            })
        },
        //获取用户订单
        getOrderApi:(data?:object)=>{
            return service({
                url: '/order/getOrderByUserID',
                method: 'POST',
                data
            })
        },
        //完成未支付订单
        completedOrderApi:(data?:object)=>{
            return service({
                url: '/order/completedOrder',
                method: 'POST',
                data
            })
        },

    }

}