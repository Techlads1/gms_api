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




type GrievanceStateActionRepository struct {
	db *pgxpool.Pool
}


func NewGrievanceStateActionRepository() *GrievanceStateActionRepository {

	db, err := database.Connect()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}

	return &GrievanceStateActionRepository{
		db: db,
	}

}


func (connect *GrievanceStateActionRepository) Store(arg *entity.GrievanceStateAction) (int, error) {

	var Id int

	query := "INSERT INTO grievance_state_actions " +
		"(name, role_perform_action, grievance_state_id, updated_at, created_at) " +
		"VALUES($1,$2,$3,$4,$5) " +
		"RETURNING id"

	err := connect.db.QueryRow(context.Background(), query,
		arg.Name, arg.RolePerformAction, arg.StateId,
		arg.UpdatedAt, arg.CreatedAt).Scan(&Id)

	return Id, err

}

//Get gets single Department
func (connect *GrievanceStateActionRepository) Show(id int) (*entity.GrievanceStateAction, error) {

	var query = "SELECT name, role_perform_action, grievance_state_id, updated_at, created_at FROM grievance_state_actions WHERE id = $1"

	var data entity.GrievanceStateAction

	data.Id = id

	err := connect.db.QueryRow(context.Background(), query, id).
		Scan(&data.Name, &data.RolePerformAction, &data.StateId, &data.UpdatedAt, &data.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &data, err

}

//Update for updating Department
func (connect *GrievanceStateActionRepository) Update(arg *entity.GrievanceStateAction) (int, error) {

	query := "UPDATE grievance_state_actions SET name = $1, role_perform_action = $2, grievance_state_id = $3, updated_at = $4" +
		" WHERE id = $5"

	_, err := connect.db.Exec(context.Background(), query, arg.Name,
		arg.RolePerformAction, arg.StateId, time.Now(), arg.Id)

	return arg.Id, err

}

//List for listing Departments
func (connect *GrievanceStateActionRepository) List() ([]*entity.GrievanceStateAction, error) {

	var entities []*entity.GrievanceStateAction

	var query = "SELECT id, name, role_perform_action, grievance_state_id, updated_at, created_at " +
		"FROM grievance_state_actions"

	rows, err := connect.db.Query(context.Background(), query)

	if err != nil {
		return nil, errors.New("error listing grievance stateActions")
	}

	for rows.Next() {

		var data entity.GrievanceStateAction

		if err := rows.Scan(&data.Id, &data.Name, &data.RolePerformAction, &data.StateId, &data.UpdatedAt, &data.CreatedAt); err != nil {
			log.Errorf("error scanning %v", err)
		}

		entities = append(entities, &data)
	}

	return entities, nil

}

//Delete for deleting Department
func (connect *GrievanceStateActionRepository) Delete(id int) error {

	query := "DELETE FROM grievance_state_actions WHERE id = $1"

	_, err := connect.db.Exec(context.Background(), query, id)

	return err
}
