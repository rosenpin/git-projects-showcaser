package parsers

import (
	"fmt"
	"reflect"

	"gitlab.com/rosenpin/git-project-showcaser/models"
)

// CreateProjectUsingTags uses the project tags to dynamically create a project object using reflections
func CreateProjectUsingTags(mappedProject map[string]interface{}, tag string) (*models.Project, error) {
	p := models.Project{}
	t := reflect.TypeOf(p)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get(tag)

		var val interface{}
		switch field.Name {
		case "IsFork":
			val, _ = mappedProject[tag].(bool)
		case "Stars":
			val, _ = mappedProject[tag].(float64)
		case "Name":
			val, _ = mappedProject[tag].(string)
		case "Description":
			val, _ = mappedProject[tag].(string)
		case "Link":
			val, _ = mappedProject[tag].(string)
		case "Language":
			val, _ = mappedProject[tag].(string)
		case "Forks":
			val, _ = mappedProject[tag].(float64)
		default:
			return nil, fmt.Errorf("Error creating project object dynamically, invalid field name %v", field.Name)
		}

		switch val.(type) {
		case bool:
			reflect.ValueOf(&p).Elem().FieldByName(field.Name).SetBool(val.(bool))
		case string:
			reflect.ValueOf(&p).Elem().FieldByName(field.Name).SetString(val.(string))
		case float64:
			reflect.ValueOf(&p).Elem().FieldByName(field.Name).SetFloat(val.(float64))
		default:
			return nil, fmt.Errorf("Error creating project object dynamically, invalid type %T", val)
		}
	}

	return &p, nil
}
