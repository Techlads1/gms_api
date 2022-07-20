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




type GrievanceForwardRepository struct {
	db *pgxpool.Pool
}


func NewGrievanceForwardRepository() *GrievanceForwardRepository {

	db, err := database.Connect()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}

	return &GrievanceForwardRepository{
		db: db,
	}

}


func (connect *GrievanceForwardRepository) Store(arg *entity.GrievanceForward) (int, error) {

	var Id int

	query := "INSERT INTO grievance_forward " +
		"(grievance_id, state,  fromgfu_id, togfu_id, description, comment, updated_at, created_at) " +
		"VALUES($1,$2,$3,$4,$5,$6,$7,$8) " +
		"RETURNING id"

	err := connect.db.QueryRow(context.Background(), query,
		arg.GrievanceId, arg.State, arg.FromGFUId, arg.ToGFUId, arg.Description, arg.Comment,
		arg.UpdatedAt, arg.CreatedAt).Scan(&Id)

	return Id, err

}


//Get gets single Department
func (connect *GrievanceForwardRepository) Show(id int) (*entity.GrievanceForward, error) {

	var query = "SELECT grievance_id, state, fromgfu_id, togfu_id, description, comment, updated_at, created_at FROM grievance_forward WHERE id = $1"

	var data entity.GrievanceForward

	data.Id = id

	err := connect.db.QueryRow(context.Background(), query, id).
		Scan(&data.GrievanceId, &data.State, &data.FromGFUId, &data.ToGFUId, &data.Description, &data.Comment, &data.UpdatedAt, &data.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &data, err

}


//Update for updating Department
func (connect *GrievanceForwardRepository) Update(arg *entity.GrievanceForward) (int, error) {

	query := "UPDATE grievance_forward SET grievance_id = $1, state = $2, fromgfu_id = $3, togfu_id = $4, description = $5, comment = $6, updated_at = $7" +
		" WHERE id = $8"

	_, err := connect.db.Exec(context.Background(), query, arg.GrievanceId, arg.State, arg.FromGFUId, arg.ToGFUId,
		arg.Description, arg.Comment, time.Now(), arg.Id)

	return arg.Id, err

}

//List for listing Departments
func (connect *GrievanceForwardRepository) List() ([]*entity.GrievanceForward, error) {

	var entities []*entity.GrievanceForward

	var query = "SELECT id, grievance_id, state,  fromgfu_id, togfu_id, description, comment, updated_at, created_at " +
		"FROM grievance_forward"

	rows, err := connect.db.Query(context.Background(), query)

	if err != nil {
		return nil, errors.New("error listing grievance Forwards")
	}

	for rows.Next() {

		var data entity.GrievanceForward

		if err := rows.Scan(&data.Id, &data.GrievanceId, &data.State, &data.FromGFUId, &data.ToGFUId, &data.Description, &data.Comment, &data.UpdatedAt, &data.CreatedAt); err != nil {
			log.Errorf("error scanning %v", err)
		}

		entities = append(entities, &data)
	}

	return entities, nil

}

func (connect *GrievanceForwardRepository) ListNew() ([]*entity.GrievanceForward, error) {

	var entities []*entity.GrievanceForward

	var query = "SELECT id, grievance_id, state,  fromgfu_id, togfu_id, description, comment, updated_at, created_at " +
		"FROM grievance_forward WHERE state = '1' "

	rows, err := connect.db.Query(context.Background(), query)

	if err != nil {
		return nil, errors.New("error listing new grievance Forwards")
	}

	for rows.Next() {

		var data entity.GrievanceForward

		if err := rows.Scan(&data.Id, &data.GrievanceId, &data.State, &data.FromGFUId, &data.ToGFUId, &data.Description, &data.Comment, &data.UpdatedAt, &data.CreatedAt); err != nil {
			log.Errorf("error scanning %v", err)
		}

		entities = append(entities, &data)
	}

	return entities, nil

}


func (connect *GrievanceForwardRepository) ListApproved() ([]*entity.GrievanceForward, error) {

	var entities []*entity.GrievanceForward

	var query = "SELECT id, grievance_id, state,  fromgfu_id, togfu_id, description, comment, updated_at, created_at " +
		"FROM grievance_forward WHERE state = '2' "

	rows, err := connect.db.Query(context.Background(), query)

	if err != nil {
		return nil, errors.New("error listing approved grievance Forwards")
	}

	for rows.Next() {

		var data entity.GrievanceForward

		if err := rows.Scan(&data.Id, &data.GrievanceId, &data.State, &data.FromGFUId, &data.ToGFUId, &data.Description, &data.Comment, &data.UpdatedAt, &data.CreatedAt); err != nil {
			log.Errorf("error scanning %v", err)
		}

		entities = append(entities, &data)
	}

	return entities, nil

}

func (connect *GrievanceForwardRepository) ListDenied() ([]*entity.GrievanceForward, error) {

	var entities []*entity.GrievanceForward

	var query = "SELECT id, grievance_id, state,  fromgfu_id, togfu_id, description, comment, updated_at, created_at " +
		"FROM grievance_forward WHERE state = '3' "

	rows, err := connect.db.Query(context.Background(), query)

	if err != nil {
		return nil, errors.New("error listing denied grievance Forwards")
	}

	for rows.Next() {

		var data entity.GrievanceForward

		if err := rows.Scan(&data.Id, &data.GrievanceId, &data.State, &data.FromGFUId, &data.ToGFUId, &data.Description, &data.Comment, &data.UpdatedAt, &data.CreatedAt); err != nil {
			log.Errorf("error scanning %v", err)
		}

		entities = append(entities, &data)
	}

	return entities, nil

}

//Delete for deleting Department
func (connect *GrievanceForwardRepository) Delete(id int) error {

	query := "DELETE FROM grievance_forward WHERE id = $1"

	_, err := connect.db.Exec(context.Background(), query, id)

	return err
}

