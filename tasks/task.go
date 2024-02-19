package main

import (
	"fmt"
	"github.com/google/uuid"
	"reflect"
	"time"
)

type Brand struct {
	ID        uuid.UUID
	Name      string
	Code      string
	IsMain    bool
	Alias     string
	Order     int
	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}
type Log struct {
	VersionNum int64
	Entity     string
	Was        interface{}
	Become     interface{}
	Changed    []string
	Time       time.Time
}

type ChangeHistory struct {
	History map[string][]Log
}

func NewChangeHistory() *ChangeHistory {
	return &ChangeHistory{
		History: make(map[string][]Log),
	}
}

func (h *ChangeHistory) LogChanges(entityID uuid.UUID, become any) error {
	var was any
	entityIDStr := entityID.String()

	if logs, exists := h.History[entityIDStr]; exists && len(logs) > 0 {
		lastLog := logs[len(logs)-1]
		was = lastLog.Become
	}

	versionNum := int64(len(h.History[entityIDStr]))
	changedFields := findChangedFields(was, become)

	entityType := reflect.TypeOf(become)
	if entityType.Kind() != reflect.Ptr {
		return fmt.Errorf(`should be a pointer'`)
	}
	entityName := entityType.Name()

	log := Log{
		VersionNum: versionNum + 1,
		Entity:     entityName,
		Was:        was,
		Become:     become,
		Changed:    changedFields,
		Time:       time.Now(),
	}

	h.History[entityIDStr] = append(h.History[entityIDStr], log)

	return nil
}

func findChangedFields(was any, become any) (changed []string) {
	wasVal := reflect.ValueOf(was)
	becomeVal := reflect.ValueOf(become)

	if wasVal.Kind() == reflect.Ptr {
		wasVal = wasVal.Elem()
	}

	if becomeVal.Kind() == reflect.Ptr {
		becomeVal = becomeVal.Elem()
	}

	for i := 0; i < becomeVal.NumField(); i++ {
		fieldName := becomeVal.Type().Field(i).Name
		becomeFieldVal := becomeVal.Field(i).Interface()

		var wasFieldVal interface{}
		if wasVal.IsValid() && i < wasVal.NumField() {
			wasFieldVal = wasVal.Field(i).Interface()
		}

		if !reflect.DeepEqual(wasFieldVal, becomeFieldVal) {
			changed = append(changed, fieldName)
		}
	}

	return changed
}

func main() {
	history := NewChangeHistory()

	brandID := uuid.New()

	oldBrand := Brand{
		ID:        brandID,
		Name:      "Old Brand",
		Code:      "OB",
		IsMain:    true,
		Alias:     "OldB",
		Order:     1,
		CreatedAt: time.Now(),
	}

	newBrand := Brand{
		ID:        brandID,
		Name:      "New Brand",
		Code:      "NB",
		IsMain:    false,
		Alias:     "NewB",
		Order:     2,
		CreatedAt: time.Now(),
	}

	// Логування початкової версії
	_ = history.LogChanges(brandID, &oldBrand)

	// Логування оновленої версії
	_ = history.LogChanges(brandID, &newBrand)

	// Виведення історії змін для brandID
	for _, log := range history.History[brandID.String()] {
		fmt.Printf("Version: %d, Changed: %v, Time: %v\n", log.VersionNum, log.Changed, log.Time)
	}
}
