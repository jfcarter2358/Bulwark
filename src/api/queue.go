package api

import (
	"bulwark/queue"
	"bulwark/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func QueueDelete(ctx *gin.Context) {
	name := ctx.Param("name")
	err := queue.Delete(name)
	if err != nil {
		utils.Error(err, ctx, http.StatusInternalServerError)
		return
	}
	ctx.Status(http.StatusOK)
}

func QueueCreate(ctx *gin.Context) {
	name := ctx.Param("name")
	if err := queue.Create(name); err != nil {
		utils.Error(err, ctx, http.StatusInternalServerError)
		return
	}
	ctx.Status(http.StatusCreated)
}

func QueuePing(ctx *gin.Context) {
	name := ctx.Param("name")
	if err := queue.Ping(name); err != nil {
		utils.Error(err, ctx, http.StatusInternalServerError)
		return
	}
	ctx.Status(http.StatusOK)
}

func QueuePop(ctx *gin.Context) {
	name := ctx.Param("name")
	datum, err := queue.Pop(name)
	if err != nil {
		utils.Error(err, ctx, http.StatusInternalServerError)
		return
	}
	if datum == nil {
		ctx.JSON(http.StatusNoContent, gin.H{})
		return
	}
	ctx.JSON(http.StatusOK, datum)
}

func QueuePush(ctx *gin.Context) {
	name := ctx.Param("name")
	var datum interface{}
	if err := ctx.ShouldBindJSON(&datum); err != nil {
		utils.Error(err, ctx, http.StatusInternalServerError)
		return
	}
	if err := queue.Push(name, datum); err != nil {
		utils.Error(err, ctx, http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, datum)
}
