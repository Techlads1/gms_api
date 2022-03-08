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
	"github.com/tzdit/sample_api/util/enums"
)




type GrievanceRepository struct {
	db *pgxpool.Pool
}


func NewGrievanceRepository() *GrievanceRepository {

	db, err := database.Connect()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}

	return &GrievanceRepository{
		db: db,
	}

}


func (connect *GrievanceRepository) Store(arg *entity.Grievance) (int, error) {

	var Id int

	query := "INSERT INTO grievances " +
		"(name, description, reference_number, comment,location_occurred, state,"+
		" grievance_filling_mode_id,grievance_sub_category_id, grievant_id, grievant_group_id, updated_at, created_at) " +
		"VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12) " +
		"RETURNING id"

	err := connect.db.QueryRow(context.Background(), query,
		arg.Name, arg.Description, arg.ReferenceNumber, arg.Comment, arg.LocationOccurred, arg.State,
		arg.FillingModeId, arg.GrievanceSubCategoryId, arg.GrievantId, arg.GrievantGroupId,
		arg.UpdatedAt, arg.CreatedAt).Scan(&Id)

		state_first_sequence,_ := connect.GetStateFirstSequence()

		id,_ :=connect.StoreGrievanceState(Id,state_first_sequence)
		pp.Print(id)
	return Id, err

}

//Get gets single Department
func (connect *GrievanceRepository) Show(id int) (*entity.Grievance, error) {

	var query = "SELECT name, description, reference_number, comment,location_occurred, state,"+
	" grievance_filling_mode_id,grievance_sub_category_id, grievant_id, grievant_group_id, updated_at, created_at FROM grievances WHERE id = $1"

	var data entity.Grievance

	data.Id = id

	err := connect.db.QueryRow(context.Background(), query, id).
		Scan(&data.Name, &data.Description, &data.ReferenceNumber, &data.Comment, &data.LocationOccurred,
			&data.State, &data.FillingModeId, &data.GrievanceSubCategoryId, &data.GrievantId, &data.GrievantGroupId,
			 &data.UpdatedAt, &data.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &data, err

}

//Update for updating Department
func (connect *GrievanceRepository) Update(arg *entity.Grievance) (int, error) {

	query := "UPDATE grievances SET name = $1, description = $2, reference_number, comment,location_occurred, state,"+
	" grievance_filling_mode_id,grievance_sub_category_id, grievant_id, grievant_group_id,  updated_at = $5" +
		" WHERE id = $6"

	_, err := connect.db.Exec(context.Background(), query, arg.Name,
		arg.Description,arg.ReferenceNumber, arg.Comment, arg.LocationOccurred, arg.State,
		arg.FillingModeId, arg.GrievanceSubCategoryId, arg.GrievantId, arg.GrievantGroupId, time.Now(), arg.Id)

	return arg.Id, err

}

//List for listing Departments
func (connect *GrievanceRepository) List() ([]*entity.Grievance, error) {

	var entities []*entity.Grievance

	var query = "SELECT id, name, description, reference_number, comment,location_occurred, state,"+
	" grievance_filling_mode_id,grievance_sub_category_id, grievant_id, grievant_group_id, updated_at, created_at " +
		"FROM grievances"

	rows, err := connect.db.Query(context.Background(), query)
pp.Print(err)
	if err != nil {
		return nil, errors.New("error listing grievances")
	}

	for rows.Next() {

		var data entity.Grievance

		if err := rows.Scan(&data.Id, &data.Name, &data.Description,&data.ReferenceNumber, &data.Comment, &data.LocationOccurred,
			&data.State, &data.FillingModeId, &data.GrievanceSubCategoryId, &data.GrievantId, &data.GrievantGroupId,
			 &data.UpdatedAt, &data.CreatedAt); err != nil {
			log.Errorf("error scanning %v", err)
		}

		entities = append(entities, &data)
	}

	return entities, nil

}

//Delete for deleting Department
func (connect *GrievanceRepository) Delete(id int) error {

	query := "DELETE FROM grievances WHERE id = $1"

	_, err := connect.db.Exec(context.Background(), query, id)

	return err
}


func (connect *GrievanceRepository) StoreGrievanceState(grievance_id, grievance_state_id int) (int, error) {

	query := "INSERT INTO grievances_has_states ( grievance_id, grievance_state_id, status, updated_at, created_at)"+
					"VALUES($1,$2,$3,$4,$5) RETURNING id"
	
	var Id int 
 var status = enums.Open
	err := connect.db.QueryRow(context.Background(), query,
	grievance_id, grievance_state_id, status.Name() , time.Now(), time.Now()).Scan(&Id)
		pp.Print(err)			
	return Id, err
}

func (connect *GrievanceRepository) GetStateFirstSequence() (int, error){

	var first_sequence int 

	err := connect.db.QueryRow(context.Background(), "SELECT id FROM public.grievance_states WHERE "+
	"sequence_number = (SELECT MIN(sequence_number) FROM public.grievance_states)").Scan(&first_sequence)

	return first_sequence,err
}