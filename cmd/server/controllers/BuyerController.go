package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/buyer"
)

type requestBuyerPost struct {
	CardNumberId int64  `json:"card_number_id" binding:"required"`
	FirstName    string `json:"first_name"    binding:"required"`
	LastName     string `json:"last_name"     binding:"required"`
}

type requestBuyerPatch struct {
	CardNumberId int64  `json:"card_number_id" binding:"required"`
	LastName     string `json:"last_name"     binding:"required"`
}

type BuyerController struct {
	service buyer.Service
}

func NewBuyer(service buyer.Service) BuyerController {
	return BuyerController{service: service}
}

// Buyers godoc
// @Summary      Create buyer
// @Description  create buyer
// @Tags         Buyers
// @Accept       json
// @Produce      json
// @Param Buyer body requestBuyerPost true "Create buyer"
// @Success      201  {object} buyer.Buyer
// @Failure      409  {object}  httputil.HTTPError
// @Failure      422  {object}  httputil.HTTPError
// @Router /buyers [post]
func (c *BuyerController) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req requestBuyerPost
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
				"error":   err.Error(),
				"message": "Invalid request"})
			return
		}

		buyer, err := c.service.Store(req.CardNumberId, req.FirstName, req.LastName)
		if err != nil {
			ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(
			http.StatusCreated,
			gin.H{
				"data": buyer,
			})
	}
}

// Buyers godoc
// @Summary      List all buyers
// @Description  get buyers
// @Tags         Buyers
// @Accept       json
// @Produce      json
// @Success      200  {object} []buyer.Buyer
// @Failure      404  {object}  httputil.HTTPError
// @Router /buyers [get]
func (c *BuyerController) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		buyers, err := c.service.GetAll()

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"data": buyers})
	}
}

// Buyers godoc
// @Summary      List buyer by id
// @Description  get buyer by id
// @Tags         Buyers
// @Accept       json
// @Produce      json
// @Param id path int true "Buyer ID"
// @Success      200  {object} buyer.Buyer
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Router /buyers/{id} [get]
func (c *BuyerController) GetId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		buyer, err := c.service.GetId(id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"data": buyer,
		})
	}
}

// Buyers godoc
// @Summary      Update UpdateCardNumberLastName
// @Description  Update UpdateCardNumberLastName field by id
// @Tags         Buyers
// @Accept       json
// @Produce      json
// @Param id path int true "Buyers ID"
// @Param Buyer body requestBuyerPatch true "Update field"
// @Success      200  {object} buyer.Buyer
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Router /buyers/{id} [patch]
func (c *BuyerController) UpdateCardNumberLastName() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		var req requestBuyerPatch
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid request"})
			return
		}

		buyer, err := c.service.Update(id, req.CardNumberId, req.LastName)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(
			http.StatusOK,
			gin.H{
				"data": buyer,
			})
	}
}

// Buyers godoc
// @Summary      DeleteBuyer buyer
// @Description  DeleteBuyer buyer by id
// @Tags         Buyers
// @Accept       json
// @Produce      json
// @Param id path int true "Buyer ID"
// @Success      204
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Router /buyers/{id} [delete]
func (c *BuyerController) DeleteBuyer() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
			return
		}

		err = c.service.Delete(int64(id))
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusNoContent, gin.H{})
	}
}