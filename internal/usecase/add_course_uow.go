package usecase

import (
	"context"

	"github.com/joaolima7/uow-goexpert/internal/entity"
	"github.com/joaolima7/uow-goexpert/internal/repository"
	"github.com/joaolima7/uow-goexpert/pkg/uow"
)

type InputUseCaseUow struct {
	CategoryName     string
	CourseName       string
	CourseCategoryID int
}

type AddCourseUseCaseUow struct {
	Uow uow.UowInterface
}

func NewAddCourseUseCaseUow(uow uow.UowInterface) *AddCourseUseCaseUow {
	return &AddCourseUseCaseUow{
		Uow: uow,
	}
}

func (a *AddCourseUseCaseUow) Execute(ctx context.Context, input InputUseCase) error {
	return a.Uow.Do(ctx, func(uow uow.UowInterface) error {
		category := entity.Category{
			Name: input.CategoryName,
		}
		repoCategory := a.GetCategoryRepository(ctx)
		err := repoCategory.Insert(ctx, category)
		if err != nil {
			return err
		}

		course := entity.Course{
			Name:       input.CourseName,
			CategoryID: input.CourseCategoryID,
		}

		repoCourse := a.GetCourseRepository(ctx)
		err = repoCourse.Insert(ctx, course)
		if err != nil {
			return err
		}

		return nil
	})

	// category := entity.Category{
	// 	Name: input.CategoryName,
	// }

	// err := a.CategoryRepositoy.Insert(ctx, category)
	// if err != nil {
	// 	return err
	// }

	// course := entity.Course{
	// 	Name:       input.CourseName,
	// 	CategoryID: input.CourseCategoryID,
	// }
	// err = a.CourseRepository.Insert(ctx, course)
	// if err != nil {
	// 	return err
	// }

	// return nil
}

func (a *AddCourseUseCaseUow) GetCategoryRepository(ctx context.Context) repository.CategoryRepositoryInterface {
	repo, err := a.Uow.GetRepository(ctx, "CategoryRepository")
	if err != nil {
		panic(err)
	}

	return repo.(repository.CategoryRepositoryInterface)
}

func (a *AddCourseUseCaseUow) GetCourseRepository(ctx context.Context) repository.CourseRepositoryInterface {
	repo, err := a.Uow.GetRepository(ctx, "CourseRepository")
	if err != nil {
		panic(err)
	}

	return repo.(repository.CourseRepositoryInterface)
}
