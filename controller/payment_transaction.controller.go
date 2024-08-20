package controller

import (
	"database/sql"
	"fmt"
	"net/http"
	db "payment/db/sqlc"
	"payment/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PaymentTransactionController struct {
	serviceManager *services.ServiceManager
}

type CreatePaymentTransactionDto struct {
	PatrxNo         string  `json:"patrx_no"`
	PatrxCreatedOn  *string `json:"patrx_created_on"`
	PatrxDebet      *int32  `json:"patrx_debet"`
	PatrxCredit     *int32  `json:"patrx_credit"`
	PatrxNotes      *string `json:"patrx_notes"`
	PatrxAcctnoFrom *string `json:"patrx_acctno_from"`
	PatrxAcctnoTo   *string `json:"patrx_acctno_to"`
	PatrxPatrxRef   *string `json:"patrx_patrx_ref"`
	PatrxTratyID    *int32  `json:"patrx_traty_id"`
}

type UpdatePaymentTransactionParams struct {
	PatrxNo         string  `json:"patrx_no"`
	PatrxCreatedOn  *string `json:"patrx_created_on"`
	PatrxDebet      *int32  `json:"patrx_debet"`
	PatrxCredit     *int32  `json:"patrx_credit"`
	PatrxNotes      *string `json:"patrx_notes"`
	PatrxAcctnoFrom *string `json:"patrx_acctno_from"`
	PatrxAcctnoTo   *string `json:"patrx_acctno_to"`
	PatrxPatrxRef   *string `json:"patrx_patrx_ref"`
	PatrxTratyID    *int32  `json:"patrx_traty_id"`
}

// constructor
func NewPaymentTransactionController(servicesManager services.ServiceManager) *PaymentTransactionController {
	return &PaymentTransactionController{
		serviceManager: &servicesManager,
	}
}

// get list
func (handler *PaymentTransactionController) GetListPaymentTransaction(e *gin.Context) {
	entities, err := handler.serviceManager.PaymentTransactionService.ListPaymentTransactions(e)
	if err != nil {
		e.JSON(http.StatusInternalServerError, err)
	}

	e.JSON(http.StatusOK, entities)
}

// get by No
func (handler *PaymentTransactionController) GetPaymentTransactionTypeByNo(e *gin.Context) {
	No, _ := strconv.Atoi(e.Param("No"))
	pt, err := handler.serviceManager.PaymentTransactionService.GetPaymentTransactionByNo(e, fmt.Sprint(No))
	if err != nil {
		e.JSON(http.StatusInternalServerError, err)
		return
	}
	e.JSON(http.StatusOK, pt)
}

func (handler *EntityController) CreatePaymentTransaction(e *gin.Context) {
	var payload *CreatePaymentTransactionDto

	if err := e.ShouldBindJSON(&payload); err != nil {
		//c.JSON(http.StatusUnprocessableEntity, gin.H{"status": "fail", "message": err})
		e.JSON(http.StatusUnprocessableEntity, err)
		return
	}

	args := db.CreatePaymentTransactionParams{
		PatrxNo:         payload.PatrxNo,
		PatrxCreatedOn:  payload.PatrxCreatedOn,
		PatrxDebet:      payload.PatrxDebet,
		PatrxCredit:     payload.PatrxCredit,
		PatrxNotes:      payload.PatrxNotes,
		PatrxAcctnoFrom: payload.PatrxAcctnoFrom,
		PatrxAcctnoTo:   payload.PatrxAcctnoTo,
		PatrxPatrxRef:   payload.PatrxPatrxRef,
		PatrxTratyID:    payload.PatrxTratyID,
	}

	pt, err := handler.serviceManager.PaymentTransactionService.CreatePaymentTransaction(e, args)
	if err != nil {
		e.JSON(http.StatusInternalServerError, err)
		return
	}

	e.JSON(http.StatusCreated, pt)

}

func (handler *EntityController) DeletePaymentTransaction(e *gin.Context) {
	No, _ := strconv.Atoi(e.Param("id"))

	_, err := handler.serviceManager.PaymentTransactionService.GetPaymentTransactionByNo(e, fmt.Sprint(No))

	if err != nil {
		if err == sql.ErrNoRows {
			e.JSON(http.StatusNotFound, err)
			return
		}
		e.JSON(http.StatusInternalServerError, err)
		return
	}

	err = handler.serviceManager.PaymentTransactionService.DeletePaymentTransaction(e, fmt.Sprint(No))
	if err != nil {
		e.JSON(http.StatusInternalServerError, err)
		return
	}
	e.JSON(http.StatusNoContent, gin.H{"status": "success", "message": "data has been deleted"})

}
