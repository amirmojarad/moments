package controllers

import (
	"github.com/gin-gonic/gin"
	"moments/ent"
	"net/http"
)

func checkErrors(err error) (gin.H, int) {
	if ent.IsConstraintError(err) {
		return gin.H{
			"message": "constraint error.",
			"error":   err,
		}, http.StatusConflict
	} else if ent.IsValidationError(err) {
		return gin.H{
			"message": "validation error",
			"error":   err,
		}, http.StatusBadRequest
	} else if ent.IsNotFound(err) {
		return gin.H{
			"message": "not found error",
			"error":   err,
		}, http.StatusNotFound
	} else if ent.IsNotSingular(err) {
		return gin.H{
			"message": "not singular error",
			"error":   err,
		}, http.StatusConflict
	} else if ent.IsNotLoaded(err) {
		return gin.H{
			"message": "not loaded error",
			"error":   err,
		}, http.StatusInternalServerError
	}
	return gin.H{
		"message": "unknown error occured",
		"err":     err.Error(),
	}, http.StatusInternalServerError
}
