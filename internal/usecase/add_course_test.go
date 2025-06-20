package usecase

import (
	"context"
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joaolima7/uow-goexpert/internal/repository"
	"github.com/stretchr/testify/assert"
)

func TestAddCourse(t *testing.T) {
	dtb, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	assert.NoError(t, err)

	dtb.Exec("DROP TABLE IF EXISTS courses;")
	dtb.Exec("DROP TABLE IF EXISTS categories;")

	dtb.Exec("CREATE TABLE IF NOT EXISTS categories (id INT AUTO_INCREMENT PRIMARY KEY, name VARCHAR(255) NOT NULL);")
	dtb.Exec("CREATE TABLE IF NOT EXISTS courses (id INT AUTO_INCREMENT PRIMARY KEY, name VARCHAR(255) NOT NULL, category_id INT, FOREIGN KEY (category_id) REFERENCES categories(id));")

	input := InputUseCase{
		CategoryName:     "Category 1",
		CourseName:       "Course 1",
		CourseCategoryID: 1,
	}

	ctx := context.Background()
	useCase := NewAddCourseUseCase(repository.NewCourseRepository(dtb), repository.NewCategoryRepository(dtb))
	err = useCase.Execute(ctx, input)
	assert.NoError(t, err)
}
