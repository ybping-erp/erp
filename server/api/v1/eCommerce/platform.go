package eCommerce

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/eCommerce"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    eCommerceReq "github.com/flipped-aurora/gin-vue-admin/server/model/eCommerce/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type PlatformApi struct {
}

var platformService = service.ServiceGroupApp.ECommerceServiceGroup.PlatformService


// CreatePlatform 创建电商平台
// @Tags Platform
// @Summary 创建电商平台
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body eCommerce.Platform true "创建电商平台"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /platform/createPlatform [post]
func (platformApi *PlatformApi) CreatePlatform(c *gin.Context) {
	var platform eCommerce.Platform
	err := c.ShouldBindJSON(&platform)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := platformService.CreatePlatform(&platform); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeletePlatform 删除电商平台
// @Tags Platform
// @Summary 删除电商平台
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body eCommerce.Platform true "删除电商平台"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /platform/deletePlatform [delete]
func (platformApi *PlatformApi) DeletePlatform(c *gin.Context) {
	var platform eCommerce.Platform
	err := c.ShouldBindJSON(&platform)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := platformService.DeletePlatform(platform); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeletePlatformByIds 批量删除电商平台
// @Tags Platform
// @Summary 批量删除电商平台
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除电商平台"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /platform/deletePlatformByIds [delete]
func (platformApi *PlatformApi) DeletePlatformByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := platformService.DeletePlatformByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdatePlatform 更新电商平台
// @Tags Platform
// @Summary 更新电商平台
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body eCommerce.Platform true "更新电商平台"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /platform/updatePlatform [put]
func (platformApi *PlatformApi) UpdatePlatform(c *gin.Context) {
	var platform eCommerce.Platform
	err := c.ShouldBindJSON(&platform)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := platformService.UpdatePlatform(platform); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindPlatform 用id查询电商平台
// @Tags Platform
// @Summary 用id查询电商平台
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query eCommerce.Platform true "用id查询电商平台"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /platform/findPlatform [get]
func (platformApi *PlatformApi) FindPlatform(c *gin.Context) {
	var platform eCommerce.Platform
	err := c.ShouldBindQuery(&platform)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if replatform, err := platformService.GetPlatform(platform.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"replatform": replatform}, c)
	}
}

// GetPlatformList 分页获取电商平台列表
// @Tags Platform
// @Summary 分页获取电商平台列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query eCommerceReq.PlatformSearch true "分页获取电商平台列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /platform/getPlatformList [get]
func (platformApi *PlatformApi) GetPlatformList(c *gin.Context) {
	var pageInfo eCommerceReq.PlatformSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := platformService.GetPlatformInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}
