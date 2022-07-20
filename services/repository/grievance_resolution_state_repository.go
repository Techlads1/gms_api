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




type GrievanceResolutionStateRepository struct {
	db *pgxpool.Pool
}


func NewGrievanceResolutionStateRepository() *GrievanceResolutionStateRepository {

	db, err := database.Connect()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}

	return &GrievanceResolutionStateRepository{
		db: db,
	}

}


func (connect *GrievanceResolutionStateRepository) Store(arg *entity.GrievanceResolutionState) (int, error) {

	var Id int

	query := "INSERT INTO grievance_resolution_state " +
		"(name, description, updated_at, created_at) " +
		"VALUES($1,$2,$3,$4) " +
		"RETURNING id"

	err := connect.db.QueryRow(context.Background(), query,
		arg.Name, arg.Description,
		arg.UpdatedAt, arg.CreatedAt).Scan(&Id)

	return Id, err

}

//Get gets single Department
func (connect *GrievanceResolutionStateRepository) Show(id int) (*entity.GrievanceResolutionState, error) {

	var query = "SELECT name, description, updated_at, created_at FROM grievance_resolution_state WHERE id = $1"

	var data entity.GrievanceResolutionState

	data.Id = id

	err := connect.db.QueryRow(context.Background(), query, id).
		Scan(&data.Name, &data.Description, &data.UpdatedAt, &data.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &data, err

}

//Update for updating Department
func (connect *GrievanceResolutionStateRepository) Update(arg *entity.GrievanceResolutionState) (int, error) {

	query := "UPDATE grievance_resolution_state SET name = $1, description = $2, updated_at = $3" +
		" WHERE id = $4"

	_, err := connect.db.Exec(context.Background(), query, arg.Name,
		arg.Description, time.Now(), arg.Id)

	return arg.Id, err

}

//List for listing Departments
func (connect *GrievanceResolutionStateRepository) List() ([]*entity.GrievanceResolutionState, error) {

	var entities []*entity.GrievanceResolutionState

	var query = "SELECT id, name, description, updated_at, created_at " +
		"FROM grievance_resolution_state"

	rows, err := connect.db.Query(context.Background(), query)

	if err != nil {
		return nil, errors.New("error listing grievance resolution states")
	}

	for rows.Next() {

		var data entity.GrievanceResolutionState

		if err := rows.Scan(&data.Id, &data.Name, &data.Description, &data.UpdatedAt, &data.CreatedAt); err != nil {
			log.Errorf("error scanning %v", err)
		}

		entities = append(entities, &data)
	}

	return entities, nil

}

//Delete for deleting Department
func (connect *GrievanceResolutionStateRepository) Delete(id int) error {

	query := "DELETE FROM grievance_resolution_state WHERE id = $1"

	_, err := connect.db.Exec(context.Background(), query, id)

	return err
}
