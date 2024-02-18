package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"text/template"
)

const ServiceTemplate = `package {{.PackageName}}

import (
	"context"
	"gocourse21/internal/core/db"
	"gocourse21/internal/domains/{{.PackageName}}/model"
)

func New{{.EntityName}}Service(repo *Repository) *service {
	return &service{
		repository: repo,
		fields:     repo.Schema().Fields(),
	}
}

type service struct {
	repository *Repository
	fields     []string
}

func (s *service) Create(ctx context.Context, entity *model.{{.EntityName}}) (db.PrimaryKey, error) {
	pk, err := s.repository.CreateOne(ctx, entity)
	if err != nil {
		return nil, err
	}

	return pk, nil
}

func (s *service) GetOne(ctx context.Context, pk db.PrimaryKey) (*model.{{.EntityName}}, error) {
	return s.repository.FindOne(ctx, s.fields, pk)
}
`

type ServiceData struct {
	PackageName     string
	EntityName      string
	ImportModelPath string
}

// GenerateServiceFiles приймає slice сутностей типу interface{} і генерує для них сервісні файли.
func GenerateServiceFiles(entities ...any) {
	for _, entity := range entities {
		entityType := reflect.TypeOf(entity)
		if entityType.Kind() == reflect.Ptr {
			entityType = entityType.Elem()
		}

		entityName := entityType.Name()
		lowerEntityName := strings.ToLower(entityName)

		data := ServiceData{
			PackageName:     lowerEntityName,
			EntityName:      entityName,
			ImportModelPath: fmt.Sprintf("gocourse21/internal/domains/%s/model", lowerEntityName),
		}

		generateServiceFile(data)
	}
}

// generateServiceFile виконує фактичну генерацію файлу на основі ServiceData.
func generateServiceFile(data ServiceData) {
	targetDir := fmt.Sprintf("./internal/domains/%s", data.PackageName)
	targetFilePath := filepath.Join(targetDir, fmt.Sprintf("service.go"))

	if err := os.MkdirAll(targetDir, os.ModePerm); err != nil {
		panic(fmt.Errorf("error creating directory: %w", err))
	}

	file, err := os.Create(targetFilePath)
	if err != nil {
		panic(fmt.Errorf("error creating file: %w", err))
	}
	defer file.Close()

	tmpl, err := template.New("service").Parse(ServiceTemplate)
	if err != nil {
		panic(fmt.Errorf("error parsing template: %w", err))
	}

	if err := tmpl.Execute(file, data); err != nil {
		panic(fmt.Errorf("error executing template: %w", err))
	}

	fmt.Printf("Service file generated for %s at %s\n", data.EntityName, targetFilePath)
}
