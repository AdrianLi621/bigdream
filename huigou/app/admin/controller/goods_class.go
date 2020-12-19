package controller

import (
	"bigdream/huigou/app/admin/service"
	"bigdream/huigou/model"
	. "bigdream/huigou/pkg"
	"fmt"
	"github.com/gin-gonic/gin"
	"runtime"
)

type GoodsClass struct {
	GcName   string `form:"gc_name" json:"gc_name" binding:"required"`
	GcPid    int    `form:"gc_pid" json:"gc_pid"`
	GcSort   int    `form:"gc_sort" json:"gc_sort"`
	GcLevel  int    `form:"gc_level" json:"gc_level"`
	GcPids   string `form:"gc_pids" json:"gc_pids"`
	GcImgUrl string `form:"gc_img_url" json:"gc_img_url"`
}
type SelectGoodsClass struct {
	Page     int    `form:"page" json:"page" binding:"required"`
	PageSize int    `form:"page_size" json:"page_size" binding:"required"`
	GcName   string `form:"gc_name" json:"gc_name" json:"gc_name"`
}

/**
获取分类列表
*/
func GoodsClassList(ctx *gin.Context) {
	var SelectForm SelectGoodsClass
	if err := ctx.ShouldBindJSON(&SelectForm); err != nil {
		BadResponse(ctx,0,nil,err.Error())
		return
	}
	condition:=make(map[string]interface{})
	condition["is_delete"]=0
	condition["gc_name_like"]="%"+SelectForm.GcName+"%"

	count:=service.CountGoodsClass(condition)
	var list []model.GoodsClass
	if count>0 {
		list=service.SelectGoodsClass(condition,SelectForm.Page,SelectForm.PageSize,"gc_id desc")
	}
	response:=make(map[string]interface{})
	response["page"]=SelectForm.Page
	response["page_size"]=SelectForm.PageSize
	response["list"]=list
	response["count"]=count

	SuccessResponse(ctx,1,response,"")
}

/**
添加分类
*/
func AddGoodsClass(ctx *gin.Context) {
	var GcForm GoodsClass
	if err := ctx.ShouldBind(&GcForm); err != nil {
		BadResponse(ctx,0,nil,err.Error())
	}
	fmt.Println(GcForm)
	runtime.Goexit()
	InsData := make(map[string]interface{})
	InsData["gc_name"] = "99"
	InsData["gc_pid"] = 1
	InsData["gc_sort"] = 1
	InsData["gc_level"] = 1
	InsData["gc_pids"] = "11"
	InsData["gc_img_url"] = "22"

	service.InsertGoodsClass(InsData)
}
