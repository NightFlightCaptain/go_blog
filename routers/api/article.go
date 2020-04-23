package api

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"go_blog/models"
	"go_blog/pkg"
	"go_blog/pkg/e"
	"go_blog/pkg/setting"
	"go_blog/pkg/util"
	"log"
	"net/http"
)

// @Summary Add a new article
// @Description get Article by id
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} pkg.Response
// @Failure 400 {object} pkg.Response
// @Router /article/{id} [get]
func GetArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID不合法")

	code := e.INVALID_PARAMS
	var data interface{}
	if !valid.HasErrors() {
		if !models.ExistArticleByID(id) {
			code = e.ERROR_NOT_EXIST_ARTICLE
		} else {
			data = models.GetArticle(id)
			code = e.SUCCESS
		}
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key:%s,err.message:%s", err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, pkg.GetResponse(code, data))
}

func GetArticles(c *gin.Context) {
	data := make(map[string]interface{})
	maps := make(map[string]interface{})

	valid := validation.Validation{}
	var state = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
		valid.Range(state, 0, 1, "state").Message("state只能为0或1")
	}

	var tagId = -1
	if arg := c.Query("tag_id"); arg != "" {
		tagId = com.StrTo(arg).MustInt()
		maps["tag_id"] = tagId
		valid.Min(tagId, 1, "tag_id").Message("标签ID不合法")
	}

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS
		data["lists"] = models.GetArticles(util.GetPage(c), setting.PageSize, maps)
		data["total"] = models.GetArticleTotal(maps)
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key :%s, err.message: %s ", err.Key, err.Message)
		}
	}
	c.JSON(http.StatusOK, pkg.GetResponse(code, data))
}

func AddArticle(c *gin.Context) {
	tagId := com.StrTo(c.Query("tag_id")).MustInt()
	title := c.Query("title")
	desc := c.Query("desc")
	content := c.Query("content")
	state := com.StrTo(c.Query("state")).MustInt()

	valid := validation.Validation{}
	valid.Min(tagId, 1, "tag_id").Message("标签ID不合法")
	valid.Required(title, "title").Message("标题不能为空")
	valid.Required(desc, "desc").Message("简述不能为空")
	valid.Required(content, "content").Message("内容不能为空")
	valid.Range(state, 0, 1, "state").Message("状态只能为0或1")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if models.ExistTagById(tagId) {
			data := make(map[string]interface{})
			data["tag_id"] = tagId
			data["title"] = title
			data["desc"] = desc
			data["content"] = content
			data["state"] = state
			models.AddArticle(data)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s ", err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, pkg.GetResponse(code, nil))
}

func EditArticle(c *gin.Context) {
	valid := validation.Validation{}

	id := com.StrTo(c.Param("id")).MustInt()
	tagId := com.StrTo(c.Query("tag_id")).MustInt()
	title := c.Query("title")
	desc := c.Query("desc")
	content := c.Query("content")
	modifiedBy := c.Query("modified_by")

	var state = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(c.Query("state")).MustInt()
		valid.Range(state, 0, 1, "state").Message("状态只能为0或1")
	}

	valid.Min(id, 1, "ID").Message("ID不合法")
	valid.MaxSize(title, 100, "title").Message("标题长度最大为100")
	valid.MaxSize(desc, 255, "desc").Message("简述长度最大为255")
	valid.MaxSize(content, 65535, "content").Message("文章长度最长为65535")
	valid.Required(modifiedBy, "modifiedBy").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "modifiedBy").Message("修改人长度最长为100")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if models.ExistArticleByID(id) {
			if models.ExistTagById(tagId) {
				data := make(map[string]interface{})
				if tagId != 0 {
					data["tag_id"] = tagId
				}
				if title != "" {
					data["title"] = title
				}
				if desc != "" {
					data["desc"] = desc
				}
				if content != "" {
					data["content"] = content
				}
				data["modifiedBy"] = modifiedBy
				models.EditArticle(id, data)
				code = e.SUCCESS
			} else {
				code = e.ERROR_NOT_EXIST_TAG
			}
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s ", err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, pkg.GetResponse(code, nil))

}

func DeleteArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID不合法")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if models.ExistArticleByID(id) {
			models.DeleteArticle(id)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s ", err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, pkg.GetResponse(code, nil))

}
