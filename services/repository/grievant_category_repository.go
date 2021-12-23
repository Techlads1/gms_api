package repository

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/tzdit/sample_api/package/log"
	"github.com/tzdit/sample_api/services/database"
	"github.com/tzdit/sample_api/services/entity"
)

type GrievantCategoryRepository struct {
	db *pgxpool.Pool
}

func NewGrievantCategoryRepository() *GrievantCategoryRepository {

	db, err := database.Connect()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}

	return &GrievantCategoryRepository{
		db: db,
	}

}

func (connect *GrievantCategoryRepository) Store(arg *entity.GrievantCategory) (int, error) {

	var Id int

	query := "INSERT INTO grievant_categories " +
		"(name, description, updated_at, created_at) " +
		"VALUES($1,$2,$3,$4) " +
		"RETURNING id"

	err := connect.db.QueryRow(context.Background(), query,
		&arg.Name, &arg.Description,
		&arg.UpdatedAt, &arg.CreatedAt).Scan(&Id)

	return Id, err

}

//Get gets single Department
func (connect *GrievantCategoryRepository) Show(id int) (*entity.GrievantCategory, error) {

	var query = "SELECT name, description, updated_at, created_at FROM grievant_categories WHERE id = $1"

	var data entity.GrievantCategory

	data.Id = id

	err := connect.db.QueryRow(context.Background(), query, id).
		Scan(&data.Name, &data.Description, &data.UpdatedAt, &data.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &data, err

}

//Update for updating Department
func (connect *GrievantCategoryRepository) Update(arg *entity.GrievantCategory) (int, error) {

	query := "UPDATE grievant_categories SET name = $1, description = $2, updated_at = $3" +
		" WHERE id = $4"

	_, err := connect.db.Exec(context.Background(), query, arg.Name,
		arg.Description, time.Now(), arg.Id)

	return arg.Id, err

}

//List for listing Departments
func (connect *GrievantCategoryRepository) List() ([]*entity.GrievantCategory, error) {

	var entities []*entity.GrievantCategory

	var query = "SELECT id, name, description, updated_at, created_at " +
		"FROM grievant_categories"

	rows, err := connect.db.Query(context.Background(), query)

	if err != nil {
		return nil, errors.New("error listing grievant categories")
	}

	for rows.Next() {

		var data entity.GrievantCategory

		if err := rows.Scan(&data.Id, &data.Name, &data.Description, &data.UpdatedAt, &data.CreatedAt); err != nil {
			log.Errorf("error scanning %v", err)
		}

		entities = append(entities, &data)
	}

	return entities, err

}

//Delete for deleting Department
func (connect *GrievantCategoryRepository) Delete(id int) error {

	query := "DELETE FROM grievant_categories WHERE id = $1"

	_, err := connect.db.Exec(context.Background(), query, id)

	return err
}
