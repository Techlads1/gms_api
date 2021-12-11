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




type GrievanceStateRepository struct {
	db *pgxpool.Pool
}


func NewGrievanceStateRepository() *GrievanceStateRepository {

	db, err := database.Connect()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}

	return &GrievanceStateRepository{
		db: db,
	}

}


func (connect *GrievanceStateRepository) Store(arg *entity.GrievanceState) (int, error) {

	var Id int

	query := "INSERT INTO grievance_states " +
		"(name, description, code_name, days, sequence_number, updated_at, created_at) " +
		"VALUES($1,$2,$3,$4,$5,$6,$7) " +
		"RETURNING id"

		last_sequence,_ := connect.GetLastSequence()

	err := connect.db.QueryRow(context.Background(), query,
		arg.Name, arg.Description, arg.CodeName, arg.Days, last_sequence+1,
		arg.UpdatedAt, arg.CreatedAt).Scan(&Id)
	
	return Id, err

}

//Get gets single Department
func (connect *GrievanceStateRepository) Show(id int) (*entity.GrievanceState, error) {

	var query = "SELECT name, description, code_name, days, sequence_number, updated_at, created_at FROM grievance_states WHERE id = $1"

	var data entity.GrievanceState

	data.Id = id

	err := connect.db.QueryRow(context.Background(), query, id).
		Scan(&data.Name, &data.Description, &data.CodeName, &data.Days, &data.SequenceNumber, &data.UpdatedAt, &data.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &data, err

}

//Update for updating Department
func (connect *GrievanceStateRepository) Update(arg *entity.GrievanceState) (int, error) {

	query := "UPDATE grievance_states SET name = $1, description = $2, code_name = $3, days = $4, updated_at = $5" +
		" WHERE id = $6"

	_, err := connect.db.Exec(context.Background(), query, arg.Name,
		arg.Description, arg.CodeName, arg.Days, time.Now(), arg.Id)

	return arg.Id, err

}

//List for listing Departments
func (connect *GrievanceStateRepository) List() ([]*entity.GrievanceState, error) {

	var entities []*entity.GrievanceState

	var query = "SELECT id, name, description, code_name, days, sequence_number, updated_at, created_at " +
		"FROM grievance_states"

	rows, err := connect.db.Query(context.Background(), query)

	if err != nil {
		return nil, errors.New("error listing grievance states")
	}

	for rows.Next() {

		var data entity.GrievanceState

		if err := rows.Scan(&data.Id, &data.Name, &data.Description, &data.CodeName, &data.Days, &data.SequenceNumber, &data.UpdatedAt, &data.CreatedAt); err != nil {
			log.Errorf("error scanning %v", err)
		}

		entities = append(entities, &data)
	}

	return entities, nil

}

//Delete for deleting Department
func (connect *GrievanceStateRepository) Delete(id int) error {

	query := "DELETE FROM grievance_states WHERE id = $1"

	_, err := connect.db.Exec(context.Background(), query, id)

	return err
}

func (connect *GrievanceStateRepository) GetLastSequence() (int, error){

	var last_sequence int 

	err := connect.db.QueryRow(context.Background(), "SELECT MAX(sequence_number) as last_sequence_number FROM grievance_states").Scan(&last_sequence)

	return last_sequence,err
}