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




type GrievanceResolutionRepository struct {
	db *pgxpool.Pool
}


func NewGrievanceResolutionRepository() *GrievanceResolutionRepository {

	db, err := database.Connect()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}

	return &GrievanceResolutionRepository{
		db: db,
	}

}


func (connect *GrievanceResolutionRepository) Store(arg *entity.GrievanceResolution) (int, error) {

	var Id int

	query := "INSERT INTO grievance_resolution " +
		"(grievance_id, gfu_id, state, description, comment, updated_at, created_at) " +
		"VALUES($1,$2,$3,$4,$5,$6,$7) " +
		"RETURNING id"

	err := connect.db.QueryRow(context.Background(), query,
		arg.GrievanceId, arg.GFUId, arg.State, arg.Description, arg.Comment,
		arg.UpdatedAt, arg.CreatedAt).Scan(&Id)

	return Id, err

}


//Get gets single Department
func (connect *GrievanceResolutionRepository) Show(id int) (*entity.GrievanceResolution, error) {

	var query = "SELECT grievance_id, gfu_id, state, description, comment, updated_at, created_at FROM grievance_resolution WHERE id = $1"

	var data entity.GrievanceResolution

	data.Id = id

	err := connect.db.QueryRow(context.Background(), query, id).
		Scan(&data.GrievanceId, &data.GFUId, &data.State, &data.Description, &data.Comment, &data.UpdatedAt, &data.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &data, err

}


//Update for updating Department
func (connect *GrievanceResolutionRepository) Update(arg *entity.GrievanceResolution) (int, error) {

	query := "UPDATE grievance_resolution SET grievance_id = $1, gfu_id = $2, state = $3,  description = $4, comment = $5, updated_at = $6" +
		" WHERE id = $7"

	_, err := connect.db.Exec(context.Background(), query, arg.GrievanceId, arg.GFUId, arg.State,
		arg.Description, arg.Comment, time.Now(), arg.Id)

	return arg.Id, err

}

//List for listing Departments
func (connect *GrievanceResolutionRepository) List() ([]*entity.GrievanceResolution, error) {

	var entities []*entity.GrievanceResolution

	var query = "SELECT id, grievance_id, gfu_id, state, description, comment, updated_at, created_at " +
		"FROM grievance_resolution"

	rows, err := connect.db.Query(context.Background(), query)

	if err != nil {
		return nil, errors.New("error listing grievance resolutions")
	}

	for rows.Next() {

		var data entity.GrievanceResolution

		if err := rows.Scan(&data.Id, &data.GrievanceId, &data.GFUId, &data.State, &data.Description, &data.Comment, &data.UpdatedAt, &data.CreatedAt); err != nil {
			log.Errorf("error scanning %v", err)
		}

		entities = append(entities, &data)
	}

	return entities, nil

}

func (connect *GrievanceResolutionRepository) ListNew() ([]*entity.GrievanceResolution, error) {

	var entities []*entity.GrievanceResolution

	var query = "SELECT id, grievance_id, gfu_id, state, description, comment, updated_at, created_at " +
		"FROM grievance_resolution WHERE state = '1' "

	rows, err := connect.db.Query(context.Background(), query)

	if err != nil {
		return nil, errors.New("error listing new grievance resolutions")
	}

	for rows.Next() {

		var data entity.GrievanceResolution

		if err := rows.Scan(&data.Id, &data.GrievanceId, &data.GFUId, &data.State, &data.Description, &data.Comment, &data.UpdatedAt, &data.CreatedAt); err != nil {
			log.Errorf("error scanning %v", err)
		}

		entities = append(entities, &data)
	}

	return entities, nil

}


func (connect *GrievanceResolutionRepository) ListApproved() ([]*entity.GrievanceResolution, error) {

	var entities []*entity.GrievanceResolution

	var query = "SELECT id, grievance_id, gfu_id, state, description, comment, updated_at, created_at " +
		"FROM grievance_resolution WHERE state = '2' "

	rows, err := connect.db.Query(context.Background(), query)

	if err != nil {
		return nil, errors.New("error listing approved grievance resolutions")
	}

	for rows.Next() {

		var data entity.GrievanceResolution

		if err := rows.Scan(&data.Id, &data.GrievanceId, &data.GFUId, &data.State, &data.Description, &data.Comment, &data.UpdatedAt, &data.CreatedAt); err != nil {
			log.Errorf("error scanning %v", err)
		}

		entities = append(entities, &data)
	}

	return entities, nil

}

func (connect *GrievanceResolutionRepository) ListDenied() ([]*entity.GrievanceResolution, error) {

	var entities []*entity.GrievanceResolution

	var query = "SELECT id, grievance_id, gfu_id, state, description, comment, updated_at, created_at " +
		"FROM grievance_resolution WHERE state = '3' "

	rows, err := connect.db.Query(context.Background(), query)

	if err != nil {
		return nil, errors.New("error listing denied grievance resolutions")
	}

	for rows.Next() {

		var data entity.GrievanceResolution

		if err := rows.Scan(&data.Id, &data.GrievanceId, &data.GFUId, &data.State, &data.Description, &data.Comment, &data.UpdatedAt, &data.CreatedAt); err != nil {
			log.Errorf("error scanning %v", err)
		}

		entities = append(entities, &data)
	}

	return entities, nil

}

//Delete for deleting Department
func (connect *GrievanceResolutionRepository) Delete(id int) error {

	query := "DELETE FROM grievance_resolution WHERE id = $1"

	_, err := connect.db.Exec(context.Background(), query, id)

	return err
}

