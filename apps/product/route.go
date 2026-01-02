package product

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func Init(router fiber.Router, db *sqlx.DB) {
	repo := newRepository(db)
	svc := newService(repo)
	handler := newHandler(svc)

	productRouter := router.Group("/product")
	{
		productRouter.Post("/", handler.CreateProduct)
		productRouter.Get("/", handler.GetAllProducts)
		productRouter.Get("/:sku", handler.GetPoductDetail)
	}
}