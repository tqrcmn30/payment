package controller

import (
	"database/sql"
	"net/http"
	db "payment/db/sqlc"
	"payment/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserAccountController struct {
	serviceManager *services.ServiceManager
}

type CreateUserAccountDto struct {
	UsacAccountNo string  `json:"usac_account_no"`
	UsacBalance   *int32  `json:"usac_balance"`
	UsacCreatedOn *string `json:"usac_created_on"`
	UsacButyID    *int32  `json:"usac_buty_id"`
	UsacUserID    *int32  `json:"usac_user_id"`
}
type UpdateUserAccountDto struct {
	UsacAccountNo string  `json:"usac_account_no"`
	UsacBalance   *int32  `json:"usac_balance"`
	UsacCreatedOn *string `json:"usac_created_on"`
	UsacButyID    *int32  `json:"usac_buty_id"`
	UsacUserID    *int32  `json:"usac_user_id"`
}

// constructor
func NewUserAccountController(servicesManager services.ServiceManager) *UserAccountController {
	return &UserAccountController{
		serviceManager: &servicesManager,
	}
}

// get list
func (handler *UserAccountController) GetListUserAccounts(e *gin.Context) {
	UserAcc, err := handler.serviceManager.UserAccountService.ListUserAccounts(e)
	if err != nil {
		e.JSON(http.StatusInternalServerError, err)
	}

	e.JSON(http.StatusOK, UserAcc)
}

// Get by ID
func (handler *UserAccountController) GetUserAccountByID(e *gin.Context) {
	id, _ := strconv.Atoi(e.Param("id"))
	UserAcc, err := handler.serviceManager.UserAccountService.GetUserAccountByID(e, int32(id))
	if err != nil {
		e.JSON(http.StatusInternalServerError, err)
		return
	}
	e.JSON(http.StatusOK, UserAcc)
}

// Create
func (handler *EntityController) CreateUserAccount(e *gin.Context) {
	var payload *CreateUserAccountDto

	if err := e.ShouldBindJSON(&payload); err != nil {
		//c.JSON(http.StatusUnprocessableEntity, gin.H{"status": "fail", "message": err})
		e.JSON(http.StatusUnprocessableEntity, err)
		return
	}

	args := db.CreateUserAccountParams{
		UsacAccountNo: payload.UsacAccountNo,
		UsacBalance:   payload.UsacBalance,
		UsacCreatedOn: payload.UsacCreatedOn,
		UsacButyID:    payload.UsacButyID,
		UsacUserID:    payload.UsacUserID,
	}

	UserAcc, err := handler.serviceManager.UserAccountService.CreateUserAccount(e, args)
	if err != nil {
		e.JSON(http.StatusInternalServerError, err)
		return
	}

	e.JSON(http.StatusCreated, UserAcc)

}

// Delete
func (handler *EntityController) DeleteUserAccount(e *gin.Context) {
	UserId, _ := strconv.Atoi(e.Param("id"))

	_, err := handler.serviceManager.UserAccountService.GetUserAccountByID(e, int32(UserId))

	if err != nil {
		if err == sql.ErrNoRows {
			e.JSON(http.StatusNotFound, err)
			return
		}
		e.JSON(http.StatusInternalServerError, err)
		return
	}

	err = handler.serviceManager.UserAccountService.DeleteUserAccount(e, int32(UserId))
	if err != nil {
		e.JSON(http.StatusInternalServerError, err)
		return
	}
	e.JSON(http.StatusNoContent, gin.H{"status": "success", "message": "data has been deleted"})

}

// Update
func (handler *EntityController) UpdateUserAccount(e *gin.Context) {
	var payload *UpdateUserAccountDto
	UserId, _ := strconv.Atoi(e.Param("id"))

	if err := e.ShouldBindJSON(&payload); err != nil {
		e.JSON(http.StatusUnprocessableEntity, err)
		return
	}

	args := &db.UpdateUserAccountParams{
		UsacID:        int32(UserId),
		UsacAccountNo: payload.UsacAccountNo,
		UsacBalance:   payload.UsacBalance,
		UsacCreatedOn: payload.UsacCreatedOn,
		UsacButyID:    payload.UsacButyID,
		UsacUserID:    payload.UsacUserID,
	}

	UserAccount := handler.serviceManager.UserAccountService.UpdateUserAccount(e, *args)
	e.JSON(http.StatusCreated, UserAccount)
}
