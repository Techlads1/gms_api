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




type GrievanceCategoryRepository struct {
	db *pgxpool.Pool
}


func NewGrievanceCategoryRepository() *GrievanceCategoryRepository {

	db, err := database.Connect()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}

	return &GrievanceCategoryRepository{
		db: db,
	}

}


func (connect *GrievanceCategoryRepository) Store(arg *entity.GrievanceCategory) (int, error) {

	var Id int

	query := "INSERT INTO grievance_categories " +
		"(name, description, code_name, updated_at, created_at) " +
		"VALUES($1,$2,$3,$4,$5) " +
		"RETURNING id"

	err := connect.db.QueryRow(context.Background(), query,
		arg.Name, arg.Description, arg.CodeName,
		arg.UpdatedAt, arg.CreatedAt).Scan(&Id)

	return Id, err

}

//Get gets single Department
func (connect *GrievanceCategoryRepository) Show(id int) (*entity.GrievanceCategory, error) {

	var query = "SELECT name, description, code_name, updated_at, created_at FROM grievance_categories WHERE id = $1"

	var data entity.GrievanceCategory

	data.Id = id

	err := connect.db.QueryRow(context.Background(), query, id).
		Scan(&data.Name, &data.Description, &data.CodeName, &data.UpdatedAt, &data.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &data, err

}

//Update for updating Department
func (connect *GrievanceCategoryRepository) Update(arg *entity.GrievanceCategory) (int, error) {

	query := "UPDATE grievance_categories SET name = $1, description = $2, code_name = $3, updated_at = $4" +
		" WHERE id = $5"

	_, err := connect.db.Exec(context.Background(), query, arg.Name,
		arg.Description, arg.CodeName, time.Now(), arg.Id)

	return arg.Id, err

}

//List for listing Departments
func (connect *GrievanceCategoryRepository) List() ([]*entity.GrievanceCategory, error) {

	var entities []*entity.GrievanceCategory

	var query = "SELECT id, name, description, code_name, updated_at, created_at " +
		"FROM grievance_categories"

	rows, err := connect.db.Query(context.Background(), query)

	if err != nil {
		return nil, errors.New("error listing grievance categories")
	}

	for rows.Next() {

		var data entity.GrievanceCategory

		if err := rows.Scan(&data.Id, &data.Name, &data.Description, &data.CodeName, &data.UpdatedAt, &data.CreatedAt); err != nil {
			log.Errorf("error scanning %v", err)
		}

		entities = append(entities, &data)
	}

	return entities, nil

}

//Delete for deleting Department
func (connect *GrievanceCategoryRepository) Delete(id int) error {

	query := "DELETE FROM grievance_categories WHERE id = $1"

	_, err := connect.db.Exec(context.Background(), query, id)

	return err
}
