package handlers

import (
	"V2RayA/global"
	"V2RayA/models/v2ray"
	"V2RayA/tools"
	"errors"
	"github.com/gin-gonic/gin"
)

func PostV2ray(ctx *gin.Context) {
	tr := global.GetTouchRaw()
	if tr.ConnectedServer == nil {
		tools.ResponseError(ctx, errors.New("不能启动V2Ray, 请选择一个节点连接"))
		return
	}
	err := tools.RestartV2rayService()
	if err != nil {
		tools.ResponseError(ctx, err)
		return
	}
	err = tools.EnableV2rayService()
	if err != nil {
		tools.ResponseError(ctx, err)
		return
	}
	tools.ResponseSuccess(ctx, gin.H{"connectedServer": tr.ConnectedServer})
}

func DeleteV2ray(ctx *gin.Context) {
	if global.ServiceControlMode == v2ray.Docker {
		tools.ResponseError(ctx, errors.New("Docker模式下无法关闭V2Ray，但可以断开节点连接"))
		return
	}
	err := tools.StopV2rayService()
	if err != nil {
		tools.ResponseError(ctx, err)
		return
	}
	err = tools.DisableV2rayService()
	if err != nil {
		tools.ResponseError(ctx, err)
		return
	}
	tools.ResponseSuccess(ctx, gin.H{"lastConnectedServer": global.GetTouchRaw().ConnectedServer})
}
