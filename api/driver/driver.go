package driver

import (
	"github.com/google/uuid"
	"github.com/o-ga09/GO_TEMPLATE_RESTAPI/api/gateway/repository"
	"gorm.io/gorm"
)

type DBDriver struct {
	conn *gorm.DB
}

func NewDBdriver(conn *gorm.DB) DBDriver {
	return DBDriver{conn}
}

func (d *DBDriver) GetAll() (repository.RepositoryJsons, error) {
	users := repository.RepositoryJsons{}
	err := d.conn.Find(&users).Error
	
	return users, err
}

func (d *DBDriver) GetById(id uuid.UUID) (repository.RepositoryJson, error) {
	suuid := id.String()
	user := repository.RepositoryJson{}
	err := d.conn.Where("userid = ?",suuid).Find(&user).Error

	return user, err
}

func (d *DBDriver) Create(param repository.RepositoryJson) error {
	err := d.conn.Create(param).Error

	return err
}
func (d *DBDriver) Update(id uuid.UUID, param repository.RepositoryParamJson) error {
	suuid := id.String()
	err := d.conn.Model(&repository.RepositoryJson{}).Where("userid = ?",suuid).Updates(&param).Error

	return err
}
func (d *DBDriver) Delete(id uuid.UUID) error {
	suuid := id.String()
	err := d.conn.Delete(&repository.RepositoryJson{},suuid).Error

	return err
}