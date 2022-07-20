package repository

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/k0kubun/pp"
	"github.com/tzdit/sample_api/package/log"
	"github.com/tzdit/sample_api/services/database"
	"github.com/tzdit/sample_api/services/entity"
)




type GrievantGroupRepository struct {
	db *pgxpool.Pool
}


func NewGrievantGroupRepository() *GrievantGroupRepository {

	db, err := database.Connect()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}

	return &GrievantGroupRepository{
		db: db,
	}

}

func (connect *GrievantGroupRepository) Store(arg *entity.GrievantGroup) (int, error) {

	var Id int

	query := "INSERT INTO grievant_groups " +
		"(name, description, grievant_category_id, updated_at, created_at) " +
		"VALUES($1,$2,$3,$4,$5) " +
		"RETURNING id"

	err := connect.db.QueryRow(context.Background(), query,
		&arg.Name, &arg.Description,&arg.GrievantCategoryId,
		&arg.UpdatedAt, &arg.CreatedAt).Scan(&Id)
		pp.Printf(err.Error())

	return Id, err

}

//Get gets single Department
func (connect *GrievantGroupRepository) Show(id int) (*entity.GrievantGroup, error) {

	var query = "SELECT name, description, grievant_category_id, updated_at, created_at FROM grievant_groups WHERE id = $1"

	var data entity.GrievantGroup

	data.Id = id

	err := connect.db.QueryRow(context.Background(), query, id).
		Scan(&data.Name, &data.Description, &data.GrievantCategoryId, &data.UpdatedAt, &data.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &data, err

}

//Update for updating Department
func (connect *GrievantGroupRepository) Update(arg *entity.GrievantGroup) (int, error) {

	query := "UPDATE grievant_groups SET name = $1, description = $2, grievant_category_id = $3, updated_at = $4" +
		" WHERE id = $5"

	_, err := connect.db.Exec(context.Background(), query, arg.Name,
		arg.Description, arg.GrievantCategoryId, time.Now(), arg.Id)

	return arg.Id, err

}

//List for listing Departments
func (connect *GrievantGroupRepository) List() ([]*entity.GrievantGroup, error) {

	var entities []*entity.GrievantGroup

	var query = "SELECT id, name, description, grievant_category_id, updated_at, created_at " +
		"FROM grievant_groups"

	rows, err := connect.db.Query(context.Background(), query)

	if err != nil {
		return nil, errors.New("error listing grievant groups")
	}

	for rows.Next() {

		var data entity.GrievantGroup

		if err := rows.Scan(&data.Id, &data.Name, &data.Description, &data.GrievantCategoryId, &data.UpdatedAt, &data.CreatedAt); err != nil {
			log.Errorf("error scanning %v", err)
		}

		entities = append(entities, &data)
	}

	return entities, err

}

//Delete for deleting Department
func (connect *GrievantGroupRepository) Delete(id int) error {
	
	query := "DELETE FROM grievant_groups WHERE id = $1"

	_, err := connect.db.Exec(context.Background(), query, id)

	return err
}
