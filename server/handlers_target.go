package server

import (
	"net/http"

	"pingmaster/target"
	"pingmaster/user"

	"github.com/gin-gonic/gin"
)

func (s Server) addTarget(c *gin.Context) {

	userReq, err := user.DecodeToken(
		// can index directly because handled in authmiddleware
		c.Request.Header["Authorization"][0],
		s.TokenSecret,
	)
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
	if !s.Sesssions.TokenExists(userReq.TokenId) {
		c.JSON(
			http.StatusUnauthorized,
			ServerResponse{
				Status:  ResponseStatus_Error,
				Message: ResponseMessage_InvalidToken,
			},
		)
		return
	}

	targetReq := target.GenericTarget{}
	err = c.ShouldBindJSON(&targetReq)
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

	newTarget, err := target.New(&targetReq, userReq)
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

	if s.TargetsPool.Contains(newTarget.GetPoolKey()) {
		c.JSON(
			http.StatusConflict,
			ServerResponse{
				Status:  ResponseStatus_Error,
				Message: ResponseMessage_TargetDuplicate,
			},
		)
		return
	}

	err = s.Database.InsertTarget(c.Request.Context(), newTarget)
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

	err = s.TargetsPool.Add(newTarget)
	if err != nil {
		c.JSON(
			http.StatusConflict,
			ServerResponse{
				Status:  ResponseStatus_Error,
				Message: err.Error(),
			},
		)
		// go s.Database.DeleteTarget(context.Background(), newTarget)
		return
	}

	c.JSON(
		http.StatusOK,
		ServerResponse{
			Status:   ResponseStatus_Success,
			Message:  ResponseMessage_TargetCreated,
			Response: newTarget,
		},
	)
}

func (s Server) getTarget(c *gin.Context) {

	userReq, err := user.DecodeToken(
		// can index directly because handled in authmiddleware
		c.Request.Header["Authorization"][0],
		s.TokenSecret,
	)
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
	if !s.Sesssions.TokenExists(userReq.TokenId) {
		c.JSON(
			http.StatusUnauthorized,
			ServerResponse{
				Status:  ResponseStatus_Error,
				Message: ResponseMessage_InvalidToken,
			},
		)
		return
	}

	urlQuery := c.Request.URL.Query()
	targetNameArr := urlQuery["name"]

	if len(targetNameArr) == 0 {
		s.getAllTargets(c, userReq)
		return
	} else if targetNameArr[0] == "" {
		s.getAllTargets(c, userReq)
		return
	} else {
		s.getOneTarget(c, targetNameArr[0], userReq)
		return
	}
}

func (s Server) getOneTarget(c *gin.Context, name string, usr *user.User) {
	tg, err := target.NewGeneric(name, usr)
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

	err = s.Database.GetTargetDetails(c.Request.Context(), tg)
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
		http.StatusOK,
		ServerResponse{
			Status:   ResponseStatus_Success,
			Response: tg,
		},
	)
}
func (s Server) getAllTargets(c *gin.Context, usr *user.User) {
	targets, err := s.Database.GetTargets(c.Request.Context(), *usr)
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
		http.StatusOK,
		ServerResponse{
			Status:   ResponseStatus_Success,
			Response: targets,
		},
	)
}
