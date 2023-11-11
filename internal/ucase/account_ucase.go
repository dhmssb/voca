package ucase

import (
	"net/http"
	"voca/internal/presentations"
	"voca/internal/repositories"
	"voca/pkg/util"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type Ucases struct {
	store repositories.Store
}

func (u *Ucases) CreateAccount(ctx *gin.Context) {

	var req presentations.CreateAccountRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	// authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	arg := presentations.CreateAccountParams{
		// Owner:    authPayload.Username,
		Currency: req.Currency,
		Balance:  0,
	}

	account, err := u.store.CreateAccount(ctx, arg)
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
