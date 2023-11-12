package ucase

import (
	"database/sql"
	"errors"
	"net/http"
	"voca/internal/presentations"
	"voca/pkg/util"

	"github.com/gin-gonic/gin"
)

func (u *Ucases) CreateProduct(ctx *gin.Context) {

	var req presentations.CreateProductParams

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	authPayload := ctx.MustGet(util.AuthorizationPayloadKey).(*util.Payload)
	arg := presentations.CreateProductParams{
		UserProduct:        authPayload.Username,
		ProductName:        req.ProductName,
		ProductDescription: req.ProductDescription,
		Quantity:           0,
		Price:              0,
	}

	product, err := u.Store.CreateProduct(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func (u *Ucases) ListProducts(ctx *gin.Context) {
	var req presentations.ListProductsParams

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	authPayload := ctx.MustGet(util.AuthorizationPayloadKey).(*util.Payload)

	arg := presentations.ListProductsParams{
		UserProduct: authPayload.Username,
		Limit:       req.Limit,
		Offset:      (req.Limit - 1) * req.Offset,
	}

	product, err := u.Store.ListProducts(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func (u *Ucases) GetProductById(ctx *gin.Context) {
	var req presentations.GetProductRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	product, err := u.Store.GetProduct(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, util.ErrorResponse(err))
		}
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}

	authPayload := ctx.MustGet(util.AuthorizationPayloadKey).(*util.Payload)
	if product.UserProduct != authPayload.Username {
		err := errors.New("account doesn't belong to the authenticated user")
		ctx.JSON(http.StatusUnauthorized, util.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, product)
}
