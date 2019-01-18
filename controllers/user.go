package controllers

import (
	"github.com/gin-gonic/gin"
	"gin-frame/models"
	. "gin-frame/utils"
	"net/http"
)


func UserSetup(v1 *gin.RouterGroup) {
	r := v1.Group("/user")

	r.GET("login", login)
	r.POST("logout", logout)

}

// @Summary 登录
// @Accept json
// @Produce  json
// @Param name query string true "name"
// @Router /login [get]
func login(c *gin.Context) {
	user := models.Users{Name: "cyh", Age: 18, Email: "abcdefg@qq.com"}
	err := Orm.Create(&user).Error
	if err != nil {
		Log.Errorln(err)
	}

	var searchUser models.Users
	err = Orm.Where("name = ?", "cyh").First(&searchUser).Error
	if err != nil {
		Log.Errorln(err)
	} else {
		Log.WithField("user", searchUser).Info("test")
	}

	searchUser.Age = 24
	err = Orm.Save(&searchUser).Error
	if err != nil {
		Log.Errorln(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"user": searchUser,
	})
}

// @Summary Print
// @Accept json
// @Produce  json
// @Param name query string true "name"
// @Router /logout [get]
func logout(c *gin.Context) {

}
