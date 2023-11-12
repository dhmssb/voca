package ucase

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"voca/internal/entity"
	"voca/internal/presentations"
	"voca/pkg/util"

	"github.com/gin-gonic/gin"
)

func (u *Ucases) CreateTransfer(ctx *gin.Context) {

	var req presentations.TransferRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	fromAccount, valid := u.validAccount(ctx, req.FromAccountID, req.Currency)

	if !valid {
		return
	}

	authPayload := ctx.MustGet(util.AuthorizationPayloadKey).(*util.Payload)
	if fromAccount.Owner != authPayload.Username {
		err := errors.New("from account doesnt belong to the authenticated user")
		ctx.JSON(http.StatusUnauthorized, util.ErrorResponse(err))
		return
	}

	_, valid = u.validAccount(ctx, req.ToAccountID, req.Currency)
	if !valid {
		return
	}

	arg := presentations.TransferTxParams{
		FromAccountID: req.FromAccountID,
		ToAccountID:   req.ToAccountID,
		Amount:        req.Amount,
	}

	result, err := u.Store.TransferTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (u *Ucases) validAccount(ctx *gin.Context, accountID int64, currency string) (entity.Account, bool) {
	account, err := u.Store.GetAccount(ctx, accountID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, util.ErrorResponse(err))
		}
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return account, false
	}

	if account.Currency != currency {
		err := fmt.Errorf("account [%d] currency mismatch: %s vs %s", accountID, account.Currency, currency)
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return account, false
	}

	return account, true

}
