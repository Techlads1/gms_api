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




type GrievanceSubCategoryRepository struct {
	db *pgxpool.Pool
}


func NewGrievanceSubCategoryRepository() *GrievanceSubCategoryRepository {

	db, err := database.Connect()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}

	return &GrievanceSubCategoryRepository{
		db: db,
	}

}


func (connect *GrievanceSubCategoryRepository) Store(arg *entity.GrievanceSubCategory) (int, error) {

	var Id int

	query := "INSERT INTO grievance_sub_categories " +
		"(name, description, code_name, grievance_category_id, updated_at, created_at) " +
		"VALUES($1,$2,$3,$4,$5,$6) " +
		"RETURNING id"

	err := connect.db.QueryRow(context.Background(), query,
		arg.Name, arg.Description, arg.CodeName, arg.GrievanceCategoryId,
		arg.UpdatedAt, arg.CreatedAt).Scan(&Id)

	return Id, err

}

//Get gets single Department
func (connect *GrievanceSubCategoryRepository) Show(id int) (*entity.GrievanceSubCategory, error) {

	var query = "SELECT name, description, code_name, grievance_category_id, updated_at, created_at FROM grievance_sub_categories WHERE id = $1"

	var data entity.GrievanceSubCategory

	data.Id = id

	err := connect.db.QueryRow(context.Background(), query, id).
		Scan(&data.Name, &data.Description, &data.CodeName, &data.GrievanceCategoryId, &data.UpdatedAt, &data.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &data, err

}

//Update for updating Department
func (connect *GrievanceSubCategoryRepository) Update(arg *entity.GrievanceSubCategory) (int, error) {

	query := "UPDATE grievance_sub_categories SET name = $1, description = $2, code_name = $3, grievance_category_id = $4, updated_at = $5" +
		" WHERE id = $6"

	_, err := connect.db.Exec(context.Background(), query, arg.Name,
		arg.Description, arg.CodeName, arg.GrievanceCategoryId, time.Now(), arg.Id)

	return arg.Id, err

}

//List for listing Departments
func (connect *GrievanceSubCategoryRepository) List() ([]*entity.GrievanceSubCategory, error) {

	var entities []*entity.GrievanceSubCategory

	var query = "SELECT id, name, description, code_name, grievance_category_id, updated_at, created_at " +
		"FROM grievance_sub_categories"

	rows, err := connect.db.Query(context.Background(), query)

	if err != nil {
		return nil, errors.New("error listing grievance sub categories")
	}

	for rows.Next() {

		var data entity.GrievanceSubCategory

		if err := rows.Scan(&data.Id, &data.Name, &data.Description, &data.CodeName, &data.GrievanceCategoryId, &data.UpdatedAt, &data.CreatedAt); err != nil {
			log.Errorf("error scanning %v", err)
		}

		entities = append(entities, &data)
	}

	return entities, nil

}

//Delete for deleting Department
func (connect *GrievanceSubCategoryRepository) Delete(id int) error {

	query := "DELETE FROM grievance_sub_categories WHERE id = $1"

	_, err := connect.db.Exec(context.Background(), query, id)

	return err
}
