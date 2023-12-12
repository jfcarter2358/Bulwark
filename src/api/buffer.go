package api

import (
	"bulwark/buffer"
	"bulwark/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BufferDelete(ctx *gin.Context) {
	name := ctx.Param("name")
	err := buffer.Delete(name)
	if err != nil {
		utils.Error(err, ctx, http.StatusInternalServerError)
		return
	}
	ctx.Status(http.StatusOK)
}

func BufferCreate(ctx *gin.Context) {
	name := ctx.Param("name")
	if err := buffer.Create(name); err != nil {
		utils.Error(err, ctx, http.StatusInternalServerError)
		return
	}
	ctx.Status(http.StatusCreated)
}

func BufferPing(ctx *gin.Context) {
	name := ctx.Param("name")
	if err := buffer.Ping(name); err != nil {
		utils.Error(err, ctx, http.StatusInternalServerError)
		return
	}
	ctx.Status(http.StatusOK)
}

func BufferGet(ctx *gin.Context) {
	name := ctx.Param("name")
	datum, err := buffer.Get(name)
	if err != nil {
		utils.Error(err, ctx, http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, datum)
}

func BufferSet(ctx *gin.Context) {
	name := ctx.Param("name")
	var datum interface{}
	if err := ctx.ShouldBindJSON(&datum); err != nil {
		utils.Error(err, ctx, http.StatusInternalServerError)
		return
	}
	if err := buffer.Set(name, datum); err != nil {
		utils.Error(err, ctx, http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, datum)
}
