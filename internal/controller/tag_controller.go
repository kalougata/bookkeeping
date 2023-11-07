package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gookit/goutil"
	"github.com/kalougata/bookkeeping/internal/model"
	"github.com/kalougata/bookkeeping/internal/service"
	"github.com/kalougata/bookkeeping/pkg/e"
	"github.com/kalougata/bookkeeping/pkg/response"
	"github.com/kalougata/bookkeeping/pkg/validator"
)

type TagController struct {
	srv *service.TagService
}

func (tc *TagController) Create() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		data := &model.TagInReq{}
		userId := ctx.GetRespHeader("userId")
		if goutil.IsEmpty(userId) {
			return response.Handle(ctx, e.ErrUnauthorized(), nil)
		}
		if err := validator.Checker(ctx, data); err != nil {
			return response.Handle(ctx, err, nil)
		}
		if err := tc.srv.Create(ctx.Context(), data); err != nil {
			return response.Handle(ctx, err, nil)
		}

		return response.Handle(ctx, nil, data)
	}
}

func NewTagController(srv *service.TagService) *TagController {
	return &TagController{srv}
}
