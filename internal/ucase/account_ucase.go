package ucase

import (
	"database/sql"
	"errors"
	"net/http"
	"voca/internal/presentations"
	"voca/internal/repositories"
	"voca/pkg/util"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type Ucases struct {
	Store      repositories.Store
	TokenMaker util.Maker
}

func (u *Ucases) CreateAccount(ctx *gin.Context) {

	var req presentations.CreateAccountRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	authPayload := ctx.MustGet(util.AuthorizationPayloadKey).(*util.Payload)
	arg := presentations.CreateAccountParams{
		Owner:    authPayload.Username,
		Currency: req.Currency,
		Balance:  0,
	}

	account, err := u.Store.CreateAccount(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique_violation":
				ctx.JSON(http.StatusForbidden, util.ErrorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

func (u *Ucases) GetAccountByID(ctx *gin.Context) {

	var req presentations.GetAccountRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	account, err := u.Store.GetAccount(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, util.ErrorResponse(err))
		}
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}

	authPayload := ctx.MustGet(util.AuthorizationPayloadKey).(*util.Payload)
	if account.Owner != authPayload.Username {
		err := errors.New("account doesn't belong to the authenticated user")
		ctx.JSON(http.StatusUnauthorized, util.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

func (u *Ucases) ListAccounts(ctx *gin.Context) {

	var req presentations.ListAccountRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	authPayload := ctx.MustGet(util.AuthorizationPayloadKey).(*util.Payload)

	arg := presentations.ListAccountsParams{
		Owner:  authPayload.Username,
		Limit:  req.PageID,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	accounts, err := u.Store.ListAccounts(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, accounts)

}
