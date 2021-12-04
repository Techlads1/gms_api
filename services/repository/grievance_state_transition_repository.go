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




type GrievanceStateTransitionRepository struct {
	db *pgxpool.Pool
}


func NewGrievanceStateTransitionRepository() *GrievanceStateTransitionRepository {

	db, err := database.Connect()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}

	return &GrievanceStateTransitionRepository{
		db: db,
	}

}


func (connect *GrievanceStateTransitionRepository) Create(arg *entity.GrievanceStateTransition) (int, error) {

	var Id int

	query := "INSERT INTO grievance_state_transitions " +
		"(description, from_state_id, to_state_id, days, update_at, created_at) " +
		"VALUES($1,$2,$3,$4,$5,$6) " +
		"RETURNING id"

	err := connect.db.QueryRow(context.Background(), query,
		arg.Description, arg.FromStateId, arg.ToStateId, arg.Days,
		arg.UpdatedAt, arg.CreatedAt).Scan(&Id)

	return Id, err

}

//Get gets single Department
func (connect *GrievanceStateTransitionRepository) Get(id int) (*entity.GrievanceStateTransition, error) {

	var query = "SELECT description, from_state_id, to_state_id, days, updated_at, created_at FROM grievance_state_transitions WHERE id = $1"

	var data entity.GrievanceStateTransition

	data.Id = id

	err := connect.db.QueryRow(context.Background(), query, id).
		Scan(&data.Description, &data.FromStateId, &data.ToStateId, &data.UpdatedAt, &data.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &data, err

}

//Update for updating Department
func (connect *GrievanceStateTransitionRepository) Update(arg *entity.GrievanceStateTransition) (int, error) {

	query := "UPDATE grievance_state_transitions SET description = $1, from_state_id = $2, to_state_id = $3, updated_at = $4" +
		" WHERE id = $5"

	_, err := connect.db.Exec(context.Background(), query, arg.Description,
		arg.FromStateId, arg.ToStateId, time.Now(), arg.Id)

	return arg.Id, err

}

//List for listing Departments
func (connect *GrievanceStateTransitionRepository) List() ([]*entity.GrievanceStateTransition, error) {

	var entities []*entity.GrievanceStateTransition

	var query = "SELECT id, description, from_state_id, to_state_id, updated_at, created_at " +
		"FROM grievance_state_transitions"

	rows, err := connect.db.Query(context.Background(), query)

	if err != nil {
		return nil, errors.New("error listing grievance stateTransitions")
	}

	for rows.Next() {

		var data entity.GrievanceStateTransition

		if err := rows.Scan(&data.Id, &data.Days, &data.FromStateId, &data.ToStateId, &data.Days, &data.UpdatedAt, &data.CreatedAt); err != nil {
			log.Errorf("error scanning %v", err)
		}

		entities = append(entities, &data)
	}

	return entities, nil

}

//Delete for deleting Department
func (connect *GrievanceStateTransitionRepository) Delete(id int) error {

	query := "DELETE FROM grievance_state_actions WHERE id = $1"

	_, err := connect.db.Exec(context.Background(), query, id)

	return err
}
