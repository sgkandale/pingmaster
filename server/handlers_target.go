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
	// if !s.Sesssions.TokenExists(userReq.TokenId) {
	// 	c.JSON(
	// 		http.StatusUnauthorized,
	// 		ServerResponse{
	// 			Status:  ResponseStatus_Error,
	// 			Message: ResponseMessage_InvalidToken,
	// 		},
	// 	)
	// 	return
	// }

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

	err = s.TargetsPool.Add(newTarget)
	if err != nil {
		c.JSON(
			http.StatusConflict,
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
			Message:  ResponseMessage_TargetCreated,
			Response: newTarget,
		},
	)
}
