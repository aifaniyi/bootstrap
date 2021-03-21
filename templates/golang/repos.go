package gotemplate

// Templates
var (
	// repos
	DbService = `
	package db

// Service : interface for fetching database repository
// objects.
// New repositories should be registered by adding
// a corresponding GetXYZRepo service signature.
// An implementation should be provided in the
// db implementation file.
type Service interface {
	%s
}
`
	DbServiceImpl = `
	package db

import (
	"gorm.io/gorm"
)

// Impl : struct containing database repository
// objects. New repositories should be registered
// in the Impl struct and a corresponding GetXYZRepo
// implementation should be provided
type Impl struct {
	conn *gorm.DB
	%s
}

// NewService : create new database service
func NewService(conn *gorm.DB) Service {
	return &Impl{
		conn: conn,
	}
}

// Migrate :
func Migrate(db *gorm.DB, isDev string) {
	entities := []interface{}{
		%s
	}

	if isDev == "true" {
		// add migrations
		db.AutoMigrate(entities...)
	} else {
		for _, model := range entities {
			if !db.Migrator().HasTable(model) {
				db.AutoMigrate(model)
			}
		}
	}
}

%s`
	DbServiceGetRepoFunction = `
	// Get%sRepo : get %s repository
func (i *Impl) Get%sRepo() %srepo.Service {
	if i.%sRepo == nil {
		i.%sRepo = %srepo.New%sRepo(i.conn)
	}

	return i.%sRepo
}`
	DbRepoService = `
	package %srepo

// Service :
type Service interface {
	Create(%s *models.%s) (*models.%s, error)
	Read(id, size int) ([]models.%s, error)
	Update(%s *models.%s) error
	Delete(%s *models.%s) error
}`
	DbRepoServiceImpl = `
	package %srepo

// Impl :
type Impl struct {
	db *gorm.DB
}

// New%sRepo : create new %s repository
func New%sRepo(session *gorm.DB) Service {
	return &Impl{
		db: session,
	}
}

// Create : create a new %s
func (u *Impl) Create(%s *models.%s) (*models.%s, error) {
	if result := u.db.Create(%s); result.Error != nil {
		return nil, fmt.Errorf("error creating %s: %%v", result.Error)
	}
	return %s, nil
}

// Read : read %s
func (u *Impl) Read(id, size int) ([]models.%s, error) {
	%s := []models.%s{}
	if result := u.db.Limit(size).Where("id > ?", id).Order("id").Find(&%s); result.Error != nil {
		if result.RowsAffected < 1 {
			return nil, fmt.Errorf("%s with id [%%d] does not exist: %%v", id, result.Error)
		}
		return nil, fmt.Errorf("error reading %s by id: %%v", result.Error)
	}
	return %s, nil
}

// Update : update %s
func (u *Impl) Update(%s *models.%s) error {
	return u.db.Transaction(func(tx *gorm.DB) error {
		// fetch record
		_%s := &models.%s{}
		if result := u.db.Where("id = ?", %s.ID).First(_%s); result.Error != nil {
			return result.Error
		}

		// save record
		if result := u.db.Model(_%s).Updates(%s); result.Error != nil {
			return result.Error
		}

		return nil
	})
}

// Delete : delete %s
func (u *Impl) Delete(%s *models.%s) error {
	if result := u.db.Unscoped().Delete(%s); result.Error != nil {
		return result.Error
	}

	return nil
}`
)
