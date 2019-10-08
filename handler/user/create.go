package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"go_restful_api/handler"
	"go_restful_api/pkg/errno"
)

func Create(c *gin.Context)  {
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	admin2 := c.Param("username")
	log.Infof("URL username: %s", admin2)

	desc := c.Query("desc")
	log.Infof("URL key param desc: %s", desc)

	contentType := c.GetHeader("Content-Type")
	log.Infof("Header Content-Type: %s", contentType)

	log.Debugf("username is: [%s], password is [%s]", r.Username, r.Password)


	if r.Username == "" {
		handler.SendResponse(c, errno.New(errno.ErrUserNotFound, fmt.Errorf("username can not found in db: xx.xx.xx.xx")), nil)
		return
	}

	if r.Password == "" {
		handler.SendResponse(c, fmt.Errorf("password is empty"), nil)
		return
	}

	rsp := CreateResponse{
		Username: r.Username,
	}

	// Show the user information.
	handler.SendResponse(c, nil, rsp)

	//var r struct{
	//	Username string `json:"username"`
	//	Password string `json:"password"`
	//}
	//
	//var err error
	//if err := c.Bind(&r); err != nil {
	//	c.JSON(http.StatusOK, gin.H{"error": errno.ErrBind})
	//	return
	//}
	//log.Debugf("username is: [%s], password is [%s]",  r.Username, r.Password)
	//
	//if r.Username == "" {
	//	err = errno.New(errno.ErrUserNotFound, fmt.Errorf("username can not found in db: xx.xx.xx.xx")).Add("This is add message.")
	//	log.Errorf(err, "Get an error")
	//}
	//
	//if errno.IsErrUserNotFound(err) {
	//	log.Debugf("err type is ErrUserNotFound")
	//}
	//
	//if r.Password == "" {
	//	err = fmt.Errorf("password is empty")
	//}
	//
	//code, message := errno.DecodeErr(err)
	//c.JSON(http.StatusOK, gin.H{"code": code, "message": message})
}