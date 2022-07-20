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




type GrievanceTimeExtensionRepository struct {
	db *pgxpool.Pool
}


func NewGrievanceTimeExtensionRepository() *GrievanceTimeExtensionRepository {

	db, err := database.Connect()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}

	return &GrievanceTimeExtensionRepository{
		db: db,
	}

}


func (connect *GrievanceTimeExtensionRepository) Store(arg *entity.GrievanceTimeExtension) (int, error) {

	var Id int

	query := "INSERT INTO grievance_time_extension " +
		"(grievance_id, gfu_id, state, description, comment, updated_at, created_at) " +
		"VALUES($1,$2,$3,$4,$5,$6,$7) " +
		"RETURNING id"

	err := connect.db.QueryRow(context.Background(), query,
		arg.GrievanceId, arg.GFUId, arg.State, arg.Description, arg.Comment,
		arg.UpdatedAt, arg.CreatedAt).Scan(&Id)

	return Id, err

}


//Get gets single Department
func (connect *GrievanceTimeExtensionRepository) Show(id int) (*entity.GrievanceTimeExtension, error) {

	var query = "SELECT grievance_id, gfu_id, state, description, comment, updated_at, created_at FROM grievance_time_extension WHERE id = $1"

	var data entity.GrievanceTimeExtension

	data.Id = id

	err := connect.db.QueryRow(context.Background(), query, id).
		Scan(&data.GrievanceId, &data.GFUId, &data.State, &data.Description, &data.Comment, &data.UpdatedAt, &data.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &data, err

}


//Update for updating Department
func (connect *GrievanceTimeExtensionRepository) Update(arg *entity.GrievanceTimeExtension) (int, error) {

	query := "UPDATE grievance_time_extension SET grievance_id = $1, gfu_id = $2, state = $3,  description = $4, comment = $5, updated_at = $6" +
		" WHERE id = $7"

	_, err := connect.db.Exec(context.Background(), query, arg.GrievanceId, arg.GFUId, arg.State,
		arg.Description, arg.Comment, time.Now(), arg.Id)

	return arg.Id, err

}

//List for listing Departments
func (connect *GrievanceTimeExtensionRepository) List() ([]*entity.GrievanceTimeExtension, error) {

	var entities []*entity.GrievanceTimeExtension

	var query = "SELECT id, grievance_id, gfu_id, state, description, comment, updated_at, created_at " +
		"FROM grievance_time_extension"

	rows, err := connect.db.Query(context.Background(), query)

	if err != nil {
		return nil, errors.New("error listing grievance TimeExtensions")
	}

	for rows.Next() {

		var data entity.GrievanceTimeExtension

		if err := rows.Scan(&data.Id, &data.GrievanceId, &data.GFUId, &data.State, &data.Description, &data.Comment, &data.UpdatedAt, &data.CreatedAt); err != nil {
			log.Errorf("error scanning %v", err)
		}

		entities = append(entities, &data)
	}

	return entities, nil

}

func (connect *GrievanceTimeExtensionRepository) ListNew() ([]*entity.GrievanceTimeExtension, error) {

	var entities []*entity.GrievanceTimeExtension

	var query = "SELECT id, grievance_id, gfu_id, state, description, comment, updated_at, created_at " +
		"FROM grievance_time_extension WHERE state = '1' "

	rows, err := connect.db.Query(context.Background(), query)

	if err != nil {
		return nil, errors.New("error listing new grievance TimeExtensions")
	}

	for rows.Next() {

		var data entity.GrievanceTimeExtension

		if err := rows.Scan(&data.Id, &data.GrievanceId, &data.GFUId, &data.State, &data.Description, &data.Comment, &data.UpdatedAt, &data.CreatedAt); err != nil {
			log.Errorf("error scanning %v", err)
		}

		entities = append(entities, &data)
	}

	return entities, nil

}


func (connect *GrievanceTimeExtensionRepository) ListApproved() ([]*entity.GrievanceTimeExtension, error) {

	var entities []*entity.GrievanceTimeExtension

	var query = "SELECT id, grievance_id, gfu_id, state, description, comment, updated_at, created_at " +
		"FROM grievance_time_extension WHERE state = '2' "

	rows, err := connect.db.Query(context.Background(), query)

	if err != nil {
		return nil, errors.New("error listing approved grievance TimeExtensions")
	}

	for rows.Next() {

		var data entity.GrievanceTimeExtension

		if err := rows.Scan(&data.Id, &data.GrievanceId, &data.GFUId, &data.State, &data.Description, &data.Comment, &data.UpdatedAt, &data.CreatedAt); err != nil {
			log.Errorf("error scanning %v", err)
		}

		entities = append(entities, &data)
	}

	return entities, nil

}

func (connect *GrievanceTimeExtensionRepository) ListDenied() ([]*entity.GrievanceTimeExtension, error) {

	var entities []*entity.GrievanceTimeExtension

	var query = "SELECT id, grievance_id, gfu_id, state, description, comment, updated_at, created_at " +
		"FROM grievance_time_extension WHERE state = '3' "

	rows, err := connect.db.Query(context.Background(), query)

	if err != nil {
		return nil, errors.New("error listing denied grievance TimeExtensions")
	}

	for rows.Next() {

		var data entity.GrievanceTimeExtension

		if err := rows.Scan(&data.Id, &data.GrievanceId, &data.GFUId, &data.State, &data.Description, &data.Comment, &data.UpdatedAt, &data.CreatedAt); err != nil {
			log.Errorf("error scanning %v", err)
		}

		entities = append(entities, &data)
	}

	return entities, nil

}

//Delete for deleting Department
func (connect *GrievanceTimeExtensionRepository) Delete(id int) error {

	query := "DELETE FROM grievance_time_extension WHERE id = $1"

	_, err := connect.db.Exec(context.Background(), query, id)

	return err
}

