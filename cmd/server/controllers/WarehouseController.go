package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/warehouse"
)

type requestWarehousePost struct {
	Address            string  `json:"address"`
	Telephone          string  `json:"telephone"`
	WarehouseCode      string  `json:"warehouse_code"`
	MinimunCapacity    int64   `json:"minimun_capacity"`
	MinimunTemperature float64 `json:"minimun_temperature"`
}

type requestWarehousePatch struct {
	MinimunCapacity    int64   `json:"minimun_capacity"`
	MinimunTemperature float64 `json:"minimun_temperature"`
}

type Warehouse struct {
	service warehouse.Service
}

func NewWarehouse(s warehouse.Service) *Warehouse {
	return &Warehouse{
		service: s,
	}
}

// Warehouse godoc
// @Summary      Create warehouse
// @Description  create warehouse
// @Tags         Warehouse
// @Accept       json
// @Produce      json
// @Param Warehouse body requestWarehousePost true "Create warehouse"
// @Success      201  {object}  warehouse.WarehouseModel
// @Failure      409  {object}  httputil.HTTPError
// @Failure      422  {object}  httputil.HTTPError
// @Router /warehouses [post]
func (w Warehouse) CreateWarehouse() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var wh requestWarehousePost

		if err := ctx.BindJSON(&wh); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, err.Error())
			return
		}

		newWh, err := w.service.Create(wh.Address, wh.Telephone, wh.WarehouseCode, wh.MinimunTemperature, wh.MinimunCapacity)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		ctx.JSON(http.StatusCreated, &newWh)
	}
}

func (w Warehouse) GetAllWarehouse() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		shw, err := w.service.GetAll()

		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		ctx.JSON(http.StatusOK, &shw)
	}
}

func (w Warehouse) GetWarehouseByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if paramId, check := ctx.Params.Get("id"); check {
			id, err := strconv.Atoi(paramId)

			if err != nil {
				ctx.JSON(http.StatusInternalServerError, "erro: internal error")
				log.Println(err)
				return
			}

			wh, err := w.service.GetById(int64(id))

			if err != nil {
				ctx.JSON(http.StatusBadRequest, err.Error())
				return
			}

			ctx.JSON(http.StatusOK, wh)
		}
	}

}

// Warehouse godoc
// @Summary      Delete Warehouse
// @Description  Delete Warehouse by id
// @Tags         Warehouse
// @Accept       json
// @Produce      json
// @Param id path int true "Warehouse ID"
// @Success      204
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Router /warehouses/{id} [delete]
func (w Warehouse) DeleteWarehouse() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if paramId, check := ctx.Params.Get("id"); check {
			id, err := strconv.Atoi(paramId)

			if err != nil {
				ctx.JSON(http.StatusInternalServerError, "erro: internal error")
				log.Println(err)
				return
			}

			err = w.service.Delete(int64(id))

			if err != nil {
				ctx.JSON(http.StatusBadRequest, err.Error())
				return
			}

			ctx.JSON(http.StatusNoContent, gin.H{})
		}
	}

}

// Warehouse godoc
// @Summary      Update warehouse
// @Description  Update warehouse
// @Tags         Warehouse
// @Accept       json
// @Produce      json
// @Param id path int true "Warehouse ID"
// @Param Warehouse body requestWarehousePatch true "Update warehouse"
// @Success      201  {object} warehouse.WarehouseModel
// @Failure      409  {object}  httputil.HTTPError
// @Failure      422  {object}  httputil.HTTPError
// @Router /warehouses/{id} [patch]
func (w Warehouse) UpdateWarehouse() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if paramId, check := ctx.Params.Get("id"); check {
			id, err := strconv.Atoi(paramId)

			var body requestWarehousePatch
			var patchWh warehouse.WarehouseModel

			if err := ctx.BindJSON(&body); err != nil {
				ctx.JSON(http.StatusBadRequest, err.Error())
				return
			}

			if err != nil {
				ctx.JSON(http.StatusInternalServerError, "erro: internal error")
				log.Println(err)
				return
			}

			patchWh, err = w.service.UpdateTempAndCap(int64(id), body.MinimunTemperature, body.MinimunCapacity)

			if err != nil {
				ctx.JSON(http.StatusBadRequest, err.Error())
				return
			}

			ctx.JSON(http.StatusOK, patchWh)
			return
		}

		ctx.JSON(http.StatusUnprocessableEntity, "error: id is mandatory")
	}
}
