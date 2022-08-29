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
	err = userReq.PrepareNew()
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

	err = userReq.CreateToken(s.TokenSecret)
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

	s.Sesssions.AddToken(userReq.TokenId)

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

	err = userReq.PrepareLogin()
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

	err = s.Database.GetUserDetails(
		c.Request.Context(),
		&userReq,
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

	err = userReq.VerifyPassword()
	if err != nil {
		c.JSON(
			http.StatusUnauthorized,
			ServerResponse{
				Status:  ResponseStatus_Error,
				Message: err.Error(),
			},
		)
		return
	}

	err = userReq.CreateToken(s.TokenSecret)
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

	s.Sesssions.AddToken(userReq.TokenId)

	c.JSON(
		http.StatusOK,
		ServerResponse{
			Status:   ResponseStatus_Success,
			Message:  ResponseMessage_UserLogin,
			Response: userReq,
		},
	)
}

func (s Server) logout(c *gin.Context) {
	authTokenArr := c.Request.Header["Authorization"]

	usr, err := user.DecodeToken(authTokenArr[0], s.TokenSecret)
	if err != nil {
		c.JSON(
			http.StatusUnauthorized,
			ServerResponse{
				Status:  ResponseStatus_Error,
				Message: err.Error(),
			},
		)
		return
	}

	deleted := s.Sesssions.DeleteToken(usr.TokenId)
	if !deleted {
		c.JSON(
			http.StatusInternalServerError,
			ServerResponse{
				Status:  ResponseStatus_Error,
				Message: ResponseMessage_GeneralError,
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		ServerResponse{
			Status:  ResponseStatus_Success,
			Message: ResponseMessage_UserLogout,
		},
	)
}

func (s Server) getHost(c *gin.Context) {}

func (s Server) addHost(c *gin.Context) {}

func (s Server) updateHost(c *gin.Context) {}

func (s Server) deleteHost(c *gin.Context) {}
