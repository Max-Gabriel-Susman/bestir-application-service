package application

import "github.com/google/uuid"

/*
 applicaiton should capture all the information needed to provision and
 manage an application on the bestir network
*/
type Application struct {
	ID   uuid.UUID `db:"id" json:"id"`
	Name string    `db:"name" json:"name"`
}

type IncomingApplication struct {
	Name string `json:"name" required:"true"`
	// IdempotencyKey null.String `json:"-" db:"idempotency_key"`
}

type ApplicationGroup struct {
}
