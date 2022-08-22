package server

import (
	"fmt"
	"net/http"

	"pingmaster/user"

	"github.com/gin-gonic/gin"
)

func (s Server) registerUser(c *gin.Context) {
	// init variables
	userReq := user.User{}
	userExists := false

	// read request
	err := c.ShouldBindJSON(&userReq)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			ServerResponse{
				Status:  ResponseStatus_Error,
				Message: ResponseMessage_ReadRequestError,
			},
		)
		return
	}

	// verify user fields
	err = userReq.Prepare()
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			ServerResponse{
				Status:  ResponseStatus_Error,
				Message: err.Error(),
			},
		)
		return
	}

	// check if user name exists in database
	userExists, err = s.Database.CheckUserExistance(
		c.Request.Context(),
		userReq,
	)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			ServerResponse{
				Status:  ResponseStatus_Error,
				Message: err.Error(),
			},
		)
		return
	}
	if userExists {
		c.JSON(
			http.StatusConflict,
			ServerResponse{
				Status:  ResponseStatus_Error,
				Message: ResponseMessage_UsernameExistsError,
			},
		)
		return
	}

	err = s.Database.InsertUser(
		c.Request.Context(),
		userReq,
	)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			ServerResponse{
				Status:  ResponseStatus_Error,
				Message: err.Error(),
			},
		)
		return
	}

	c.JSON(
		http.StatusCreated,
		ServerResponse{
			Status:   ResponseStatus_Success,
			Message:  fmt.Sprintf("user registered with name '%s'", userReq.Name),
			Response: userReq,
		},
	)
}

func (s Server) login(c *gin.Context) {
	userReq := user.User{}

	err := c.BindJSON(&userReq)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			ServerResponse{
				Status:  ResponseStatus_Error,
				Message: ResponseMessage_ReadRequestError,
			},
		)
		return
	}

	err = userReq.Prepare()
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			ServerResponse{
				Status:  ResponseStatus_Error,
				Message: err.Error(),
			},
		)
		return
	}
}

func (s Server) logout(c *gin.Context) {}

func (s Server) getHost(c *gin.Context) {}

func (s Server) addHost(c *gin.Context) {}

func (s Server) updateHost(c *gin.Context) {}

func (s Server) deleteHost(c *gin.Context) {}
