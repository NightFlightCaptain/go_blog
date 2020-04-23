package api

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"go_blog/models"
	"go_blog/pkg/e"
	"net/http"
)

func GetTags(c *gin.Context) {
	//id:=c.Query("id")
	//name:=c.Query("name")

}

// @Summary 新增文章标签
// @Produce json
// @Param name query string true "Name"
// @Param state query int false "State"
// @Param created_by query int false "CreatedBy"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /tag [post]
func AddTag(c *gin.Context) {
	name := c.Query("name")
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()
	createdBy := c.Query("createdBy")

	vaild := validation.Validation{}
	vaild.Required(name, "name").Message("名称不能为空")
	vaild.MaxSize(name, 100, "name").Message("名称最长为100")
	vaild.Required(createdBy, "createdBy").Message("创建人不能为空")
	vaild.MaxSize(createdBy, 100, "createdBy").Message("创建人名称最长为100")
	vaild.Range(state, 0, 1, "state").Message("state只能为0或1")

	code := e.INVALID_PARAMS
	if !vaild.HasErrors() {
		if !models.ExistTagByName(name) {
			code = e.SUCCESS
			models.AddTag(name, state, createdBy)
		} else {
			code = e.ERROR_EXIST_TAG
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": nil,
	})
}

func EditTag(c *gin.Context) {
	id := com.StrTo(c.Query("id")).MustInt()
	name := c.Query("name")
	modifiedBy := c.Query("modifiedBy")
	vaild := validation.Validation{}

	state := -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		vaild.Range(state, 0, 1, "state").Message("状态只能为0或1")
	}
	vaild.Required(name, "name").Message("名称不能为空")
	vaild.MaxSize(name, 100, "name").Message("名称最长为100")
	vaild.Required(modifiedBy, "modifiedBy").Message("修改人不能为空")
	vaild.MaxSize(modifiedBy, 100, "modifiedBy").Message("修改人名称最长为100")

	code := e.INVALID_PARAMS
	if !vaild.HasErrors() {
		code = e.SUCCESS
		if models.ExistTagById(id) {
			data := make(map[string]interface{})
			data["modified_by"] = modifiedBy
			if name != "" {
				data["name"] = name
			}
			if state != -1 {
				data["state"] = state
			}
			models.EditTag(id, data)
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]interface{}),
	})

}

func DeleteTag(c *gin.Context) {
	id := com.StrTo(c.Query("id")).MustInt()

	vaild := validation.Validation{}
	vaild.Min(id, 1, "id").Message("ID不合法")

	code := e.INVALID_PARAMS
	if !vaild.HasErrors() {
		code = e.SUCCESS
		if !models.ExistTagById(id) {
			code = e.ERROR_NOT_EXIST_TAG
		} else {
			models.DeleteTag(id)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]interface{}),
	})

}
