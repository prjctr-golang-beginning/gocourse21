package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"text/template"
)

const RepositoryTemplate = `package {{.PackageName}}

import (
	"database/sql"
	"gocourse21/internal/core/db"
	"gocourse21/internal/domains/{{.PackageName}}/model"
	"sync"
)

var (
	schema *db.TableSchema
	once   sync.Once
)

func Schema() *db.TableSchema {
	once.Do(func() {
		schema = db.NewTableSchema(&model.{{.EntityName}}{})
	})

	return schema
}

func New{{.EntityName}}Repository(pool sql.Conn) *Repository {
	return &Repository{
		Repository: sql.NewRepository[model.{{.EntityName}}](pool, Schema()),
	}
}

type Repository struct {
	sql.Repository[model.{{.EntityName}}]
}
`

type RepositoryData struct {
	PackageName string
	EntityName  string
}

// GenerateRepositoryFiles приймає slice сутностей типу interface{} і генерує для них репозиторійні файли.
func GenerateRepositoryFiles(entities ...any) {
	for _, entity := range entities {
		entityType := reflect.TypeOf(entity)
		if entityType.Kind() == reflect.Ptr {
			entityType = entityType.Elem()
		}
		entityName := entityType.Name()
		lowerEntityName := strings.ToLower(entityName)

		data := RepositoryData{
			PackageName: lowerEntityName,
			EntityName:  entityName,
		}

		generateRepositoryFile(data)
	}
}

// generateRepositoryFile виконує фактичну генерацію файлу репозиторію на основі RepositoryData.
func generateRepositoryFile(data RepositoryData) {
	targetDir := fmt.Sprintf("./internal/domains/%s", data.PackageName)
	targetFilePath := filepath.Join(targetDir, fmt.Sprintf("repository.go"))
	fmt.Println(targetDir)
	if err := os.MkdirAll(targetDir, os.ModePerm); err != nil {
		panic(fmt.Errorf("error creating directory: %w", err))
	}

	file, err := os.Create(targetFilePath)
	if err != nil {
		panic(fmt.Errorf("error creating file: %w", err))
	}
	defer file.Close()

	tmpl, err := template.New("repository").Parse(RepositoryTemplate)
	if err != nil {
		panic(fmt.Errorf("error parsing template: %w", err))
	}

	if err := tmpl.Execute(file, data); err != nil {
		panic(fmt.Errorf("error executing template: %w", err))
	}

	fmt.Printf("Repository file generated for %s at %s\n", data.EntityName, targetFilePath)
}
