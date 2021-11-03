package repository

import (
	"context"
	"errors"
	"fmt"

	"os"
	"time"

	"github.com/tzdit/sample_api/package/log"
	"github.com/tzdit/sample_api/services/database"
	"github.com/tzdit/sample_api/services/entity"

	"github.com/jackc/pgx/v4/pgxpool"
)

//DepartmentConn for connection to postgresql repo
type DepartmentConn struct {
	conn *pgxpool.Pool
}

//NewDepartment connection
func NewDepartment() *DepartmentConn {
	conn, err := database.Connect()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}
	return &DepartmentConn{
		conn: conn,
	}
}

//Create for creating new Department
func (con *DepartmentConn) Create(dep *entity.Department) (int, error) {
	var departmentID int
	query := "INSERT INTO departments " +
		"(department_title,department_description,department_size" +
		"created_by,created_at) " +
		"VALUES($1,$2,$3,$4,$5,$6) " +
		"RETURNING id"
	err := con.conn.QueryRow(context.Background(), query,
		dep.DepartmentTitle, dep.DepartmentDescription, dep.DepartmentSize,
		dep.CreatedBy, dep.CreatedAt).Scan(&departmentID)
	return departmentID, err
}

//Get gets single Department
func (con *DepartmentConn) Get(id int) (*entity.Department, error) {

	var query = "SELECT department_title, department_description,department_size, campus_id, created_by, created_at FROM departments WHERE department_id = $1"
	var ent entity.Department
	ent.Id = id

	err := con.conn.QueryRow(context.Background(), query, id).
		Scan(&ent.DepartmentTitle, &ent.DepartmentDescription, &ent.DepartmentSize, &ent.CampusId, &ent.CreatedBy, &ent.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &ent, err
}

//Update for updating Department
func (con *DepartmentConn) Update(dep *entity.Department) (int, error) {
	query := "UPDATE departments SET department_title = $1, department_description = $2, department_size = $3, campus_id = $4" +
		"updated_by = $5, updated_at = $6" +
		" WHERE department_id = $7"
	_, err := con.conn.Exec(context.Background(), query, dep.DepartmentTitle,
		dep.DepartmentDescription, dep.DepartmentSize, dep.CampusId, dep.UpdatedBy, time.Now(), dep.Id)
	return dep.Id, err
}

//List for listing Departments
func (con *DepartmentConn) List() ([]*entity.Department, error) {

	var entities []*entity.Department
	var query = "SELECT department_id, department_title, department_description,department_size, campus_id, created_by, created_at " +
		"FROM departments " +
		"WHERE deleted_by IS NULL"

	rows, err := con.conn.Query(context.Background(), query)
	if err != nil {
		return nil, errors.New("error listing departments")
	}
	for rows.Next() {
		var ent entity.Department
		if err := rows.Scan(&ent.Id, &ent.DepartmentTitle, &ent.DepartmentDescription, &ent.DepartmentSize, &ent.CampusId, &ent.CreatedBy, &ent.CreatedAt); err != nil {
			log.Errorf("error scanning %v", err)
		}
		entities = append(entities, &ent)
	}
	return entities, nil
}

//Delete for deleting Department
func (con *DepartmentConn) Delete(id, deletedBy int) error {
	query := "UPDATE departments SET deleted_by = $1, deleted_at = $2 WHERE department_id = $3"
	_, err := con.conn.Exec(context.Background(), query, deletedBy, time.Now(), id)
	return err
}
