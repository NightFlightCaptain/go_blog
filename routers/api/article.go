package api

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"go_blog/pkg/app"
	"go_blog/pkg/e"
	"go_blog/pkg/util"
	"go_blog/service"
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

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		c.JSON(http.StatusOK, app.GetResponse(e.INVALID_PARAMS, nil))
		return
	}

	data, code := articleService.GetArticle(id)
	if code != e.SUCCESS {
		c.JSON(http.StatusOK, app.GetResponse(code, nil))
		return
	}
	c.JSON(http.StatusOK, app.GetResponse(code, data))
}

func GetArticles(c *gin.Context) {
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

	if valid.HasErrors() {
		for _, err := range valid.Errors {
			log.Printf("err.key :%s, err.message: %s ", err.Key, err.Message)
		}
		c.JSON(http.StatusOK, app.GetResponse(e.INVALID_PARAMS, nil))
		return
	}
	offset, limit := util.GetPage(c)
	data, code := articleService.GetArticles(offset, limit, maps)
	c.JSON(http.StatusOK, app.GetResponse(code, data))
}

type AddArticleForm struct {
	TagId         int    `form:"tag_id" valid:"Required;Min(1)"`
	Title         string `form:"title" valid:"Required;MaxSize(100)"`
	Desc          string `form:"desc" valid:"Required;MaxSize(256)"`
	Content       string `form:"content" valid:"Required;MaxSize(65535)"`
	CreatedBy     string `form:"created_by"  valid:"Required;MaxSize(100)"`
	CoverImageUrl string `form:"cover_image_url" valid:"Required;MaxSize(255)"`
	State         int    `form:"state" valid:"Range(0,1)"`
}

func AddArticle(c *gin.Context) {

	var articleForm AddArticleForm
	httpCode, code := app.BindAndValid(c, &articleForm)
	if code != e.SUCCESS {
		c.JSON(httpCode, app.GetResponse(code, nil))
		return
	}

	_, code = tagService.ExistTagById(articleForm.TagId)
	if code != e.SUCCESS {
		c.JSON(http.StatusOK, app.GetResponse(code, nil))
		return
	}

	article := service.Article{
		TagId:         articleForm.TagId,
		Title:         articleForm.Title,
		Desc:          articleForm.Desc,
		Content:       articleForm.Content,
		CoverImageUrl: articleForm.CoverImageUrl,
		State:         articleForm.State,
		CreatedBy:     articleForm.CreatedBy,
	}
	code = articleService.AddArticle(article)

	c.JSON(http.StatusOK, app.GetResponse(code, nil))
}

type EditArticleForm struct {
	Id            int    `form:"id" valid:"Required;Min(1)"`
	TagId         int    `form:"tag_id" valid:"Min(1)"`
	Title         string `form:"title" valid:"MaxSize(100)"`
	Desc          string `form:"desc" valid:"MaxSize(256)"`
	Content       string `form:"content" valid:"MaxSize(65535)"`
	ModifiedBy    string `form:"modified_by"  valid:"Required;MaxSize(100)"`
	CoverImageUrl string `form:"cover_image_url" valid:"MaxSize(255)"`
	State         int    `form:"state" valid:"Range(0,1)"`
}

func EditArticle(c *gin.Context) {
	var articleForm EditArticleForm

	httpCode, errCode := app.BindAndValid(c, &articleForm)
	if errCode != e.SUCCESS {
		c.JSON(httpCode, app.GetResponse(errCode, nil))
		return
	}

	article := service.Article{
		Id:            articleForm.Id,
		TagId:         articleForm.TagId,
		Title:         articleForm.Title,
		Desc:          articleForm.Desc,
		Content:       articleForm.Content,
		CoverImageUrl: articleForm.CoverImageUrl,
		State:         articleForm.State,
		ModifiedBy:    articleForm.ModifiedBy,
	}
	code := articleService.EditArticle(article)
	c.JSON(http.StatusOK, app.GetResponse(code, nil))

}

func DeleteArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID不合法")

	if valid.HasErrors() {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s ", err.Key, err.Message)
		}
		c.JSON(http.StatusOK, app.GetResponse(e.INVALID_PARAMS, nil))
		return
	}

	code := articleService.DeleteArticle(id)
	c.JSON(http.StatusOK, app.GetResponse(code, nil))
}
