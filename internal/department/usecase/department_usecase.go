package usecase

import (
	"context"
	"go-gomanager/db"
)

type DepartmentUseCase struct {

		DB *db.Postgres
}


func (c *DepartmentUseCase) Create(ctx context.Context, request)