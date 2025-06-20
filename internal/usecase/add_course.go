package usecase

import (
	"context"

	"github.com/joaolima7/uow-goexpert/internal/entity"
	"github.com/joaolima7/uow-goexpert/internal/repository"
)

type InputUseCase struct {
	CategoryName     string
	CourseName       string
	CourseCategoryID int
}

type AddCourseUseCase struct {
	CourseRepository  repository.CourseRepositoryInterface
	CategoryRepositoy repository.CategoryRepositoryInterface
}

func NewAddCourseUseCase(course repository.CourseRepositoryInterface, category repository.CategoryRepositoryInterface) *AddCourseUseCase {
	return &AddCourseUseCase{
		CourseRepository:  course,
		CategoryRepositoy: category,
	}
}

func (a *AddCourseUseCase) Execute(ctx context.Context, input InputUseCase) error {
	category := entity.Category{
		Name: input.CategoryName,
	}

	err := a.CategoryRepositoy.Insert(ctx, category)
	if err != nil {
		return err
	}

	course := entity.Course{
		Name:       input.CourseName,
		CategoryID: input.CourseCategoryID,
	}
	err = a.CourseRepository.Insert(ctx, course)
	if err != nil {
		return err
	}

	return nil
}
