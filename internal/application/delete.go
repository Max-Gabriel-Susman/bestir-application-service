package application

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (api *API) DeleteApplication(ctx context.Context, incomingApplication IncomingApplication) (Application, error) {
	fmt.Println("Create Account B invoked")
	id := uuid.New()

	application := Application{
		ID:   id,
		Name: incomingApplication.Name,
	}

	err := api.Store.DeleteApplication(ctx, application)

	return application, err
}
