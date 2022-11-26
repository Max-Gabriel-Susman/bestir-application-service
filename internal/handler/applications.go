package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Max-Gabriel-Susman/bestir-identity-service/internal/application"
	"github.com/Max-Gabriel-Susman/bestir-identity-service/internal/foundation/web"
	"github.com/go-chi/chi/v5"
)

type Application struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type applicationGroup struct {
	*application.API
}

type ListApplicationsResponse struct {
	Applications []application.Application `json:"applications"`
}

func ApplicationEndpoints(app *web.App, api *application.API) {
	ag := applicationGroup{API: api}

	// app.Handle("GET", "/application", ag.GetApplication)
	app.Handle("GET", "/application", ag.ListApplications)
	app.Handle("POST", "/application", ag.CreateApplication)
}

func (ag applicationGroup) ListApplications(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	applications, err := ag.API.ListApplications(ctx)
	if err != nil {
		return err
	}

	return web.Respond(ctx, w, ListApplicationsResponse{
		Applications: applications,
	}, http.StatusOK)
}

// accounts := []account.Account{
// 	{
// 		ID:      69,
// 		Balance: 69,
// 	},
// 	{
// 		ID:      420,
// 		Balance: 420,
// 	},
// }

func (ag applicationGroup) CreateApplication(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	fmt.Println("Create Account a invoked")
	var input application.IncomingApplication
	if err := web.Decode(r.Body, &input); err != nil {
		return err
	}

	application, err := ag.API.CreateApplication(ctx, input)
	if err != nil {
		return err
	}

	return web.Respond(ctx, w, application, http.StatusCreated)
}

func (ag applicationGroup) GetApplication(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	applicationID := chi.URLParam(r, "application_id")
	if applicationID == "" {
		return nil
		// return handleMissingURLParameter(ctx, accountID, Account)
	}

	return nil
}
