package server

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"pingmaster/helpers"
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
		log.Print(err)
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

func (s Server) getPings(c *gin.Context) {

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

	var (
		beforeInt int64 = time.Now().Unix()
		limitInt  int64 = Default_PingCountLimit
	)

	urlQuery := c.Request.URL.Query()
	nameArr := urlQuery["name"]
	beforeArr := urlQuery["before"]
	limitArr := urlQuery["limit"]

	if len(nameArr) == 0 || nameArr[0] == "" {
		c.JSON(
			http.StatusBadRequest,
			ServerResponse{
				Status:   ResponseStatus_Error,
				Message:  "name not provided in query params",
				Response: map[string]string{"hint": "add 'name' query param with value being the name of the target"},
			},
		)
		return
	}
	if len(beforeArr) > 0 {
		if beforeArr[0] != "" {
			beforeTime, err := time.Parse(helpers.Default_TimeFormat, beforeArr[0])
			if err != nil {
				c.JSON(
					http.StatusBadRequest,
					ServerResponse{
						Status:   ResponseStatus_Error,
						Message:  err.Error(),
						Response: map[string]string{"hint": "timestamp should have a layout like '2006-01-02 15:04:05'"},
					},
				)
				return
			}
			beforeInt = beforeTime.Unix()
		}
	}
	if len(limitArr) > 0 {
		if limitArr[0] != "" {
			limitInt, err = strconv.ParseInt(limitArr[0], 10, 64)
			if err != nil {
				c.JSON(
					http.StatusBadRequest,
					ServerResponse{
						Status:   ResponseStatus_Error,
						Message:  err.Error(),
						Response: map[string]string{"hint": "limit should be integer"},
					},
				)
				return
			}
		}
	}

	genericTarget := target.GenericTarget{
		Name: nameArr[0],
		User: userReq,
	}

	err = s.Database.GetTargetDetails(c.Request.Context(), &genericTarget)
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

	pings, err := s.Database.GetPings(c.Request.Context(), &genericTarget, beforeInt, limitInt)
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
			Response: pings,
		},
	)
}
