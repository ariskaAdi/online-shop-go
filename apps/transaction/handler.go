package transaction

import (
	infrafiber "ariskaAdi-online-shop/infra/fiber"
	"ariskaAdi-online-shop/infra/response"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	svc service
}

func newHandler(svc service) handler {
	return handler{
		svc: svc,
	}
}

// CreateTransaction creates a new transaction with given request and product stock.
// It starts a new transaction, creates a new transaction entity, validates the transaction,
// updates the product stock, and commits the transaction to the database.
// If any error occurs, it will rollback the transaction.
func (h handler) CreateTransaction(ctx *fiber.Ctx) error {

	var req = CreateTransactionRequestPayload{}

	if err := ctx.BodyParser(&req); err != nil {
		return infrafiber.NewResponse(
			infrafiber.WithMessage(err.Error()),
			infrafiber.WithError(err),
			infrafiber.WithHttpCode(http.StatusBadRequest),
		).Send(ctx)

	}

	userPublicId := ctx.Locals("PUBLIC_ID")
	req.UserPublicId = fmt.Sprintf("%v", userPublicId)

	if err := h.svc.CreateTransaction(ctx.UserContext(), req); err != nil {
		myErr, ok := response.ErrorMapping[err.Error()]
		if !ok {
			myErr = response.ErrorGeneral
		}

		return infrafiber.NewResponse(
			infrafiber.WithMessage(err.Error()),
			infrafiber.WithError(myErr),
		).Send(ctx)
	}

	return infrafiber.NewResponse(
		infrafiber.WithHttpCode(http.StatusCreated),
		infrafiber.WithMessage("create transactions success"),
	).Send(ctx)
}


// GetTransactionByUser returns transaction histories of a user by user public id. If no transactions found, returns an empty array and nil error. If error occurs, returns an empty array and the error.
func (h handler) GetTransactionByUser(ctx *fiber.Ctx) error {
	userPublicId := fmt.Sprintf("%v", ctx.Locals("PUBLIC_ID"))

	trxs, err := h.svc.TransactionsHistories(ctx.UserContext(), userPublicId)
	if err != nil {
		myErr, ok := response.ErrorMapping[err.Error()]
		if !ok {
			myErr = response.ErrorGeneral
		}

		return infrafiber.NewResponse(
			infrafiber.WithMessage(err.Error()),
			infrafiber.WithError(myErr),
		).Send(ctx)
	}

	var response = []TransactionHisotryResponse{}

	for _, trx := range trxs {
		response = append(response, trx.ToTransactionHistoryResponse())
	}

	return infrafiber.NewResponse(
		infrafiber.WithHttpCode(http.StatusCreated),
		infrafiber.WithPayload(response),
		infrafiber.WithMessage("get transaction histories success"),
	).Send(ctx)
}