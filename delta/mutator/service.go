//nolint:cyclop
package mutator

import (
	"github.com/integration-system/isp-mdb-lib/delta"
	"github.com/pkg/errors"
)

const (
	itemIdPath = "ITEM_ID"
)

type Service struct{}

func NewService() Service {
	return Service{}
}

func (s Service) Apply(data map[string]any, changelogs []delta.Changelog) error {
	for i, changelog := range changelogs {
		var err error
		switch changelog.Operation {
		case delta.DeleteBlock:
			err = s.deleteBlockOperation(data, changelog)
		case delta.ChangeField:
			if changelog.Path.IsArray {
				err = s.changeFieldOperationForArray(data, changelog)
			} else {
				err = s.changeFieldOperationForObject(data, changelog)
			}
		case delta.DeleteField:
			if changelog.Path.IsArray {
				err = s.deleteFieldOperationForArray(data, changelog)
			} else {
				err = s.deleteFieldOperationForObject(data, changelog)
			}
		case delta.DeleteArrayItem:
			err = s.deleteArrayItemOperation(data, changelog)
		case delta.AddArrayItem:
			err = s.addArrayItemOperation(data, changelog)
		case delta.ExtraOperation:
			continue
		default:
			err = errors.New("unexpected operation")
		}
		if err != nil {
			return errors.WithMessagef(err, "apply operation [%d]: %s", i, changelog.Operation)
		}
	}
	return nil
}

func (s Service) changeFieldOperationForArray(data map[string]any, change delta.Changelog) error {
	err := s.validateArrayPath(change.Path)
	if err != nil {
		return errors.WithMessage(err, "validate array path")
	}

	if change.Field == "" {
		return errors.New("field is empty")
	}

	if change.NewValue == nil {
		return errors.New("newValue is empty")
	}

	dataBlock, exist := data[change.Path.BlockName]
	if !exist {
		dataBlock = make([]any, 0)
	}

	itemBlock, ok := dataBlock.([]any)
	if !ok {
		return errors.Errorf("%s is not array in data", change.Path.BlockName)
	}

	itemIdIsExist := false
	for _, item := range itemBlock {
		item, ok := item.(map[string]any)
		if !ok {
			return errors.Errorf("%s is not array of objects in data", change.Path.BlockName)
		}

		itemId, foundItemId := item[itemIdPath]
		if !foundItemId {
			return errors.Errorf("%s without %s in data", change.Path.BlockName, itemIdPath)
		}

		if itemId == change.Path.ItemId {
			itemIdIsExist = true
			item[change.Field] = change.NewValue
			break
		}
	}
	if !itemIdIsExist {
		itemBlock = append(itemBlock, map[string]any{
			itemIdPath:   change.Path.ItemId,
			change.Field: change.NewValue,
		})
	}

	data[change.Path.BlockName] = itemBlock
	return nil
}

func (s Service) changeFieldOperationForObject(data map[string]any, change delta.Changelog) error {
	err := s.validateObjectPath(change.Path)
	if err != nil {
		return errors.WithMessage(err, "validate object path")
	}

	if change.Field == "" {
		return errors.New("field is empty")
	}

	if change.NewValue == nil {
		return errors.New("newValue is empty")
	}

	dataBlock, exist := data[change.Path.BlockName]
	if !exist {
		dataBlock = make(map[string]any)
	}

	itemBlock, ok := dataBlock.(map[string]any)
	if !ok {
		return errors.Errorf("block %s is not has field %s", change.Path.BlockName, change.Field)
	}

	itemBlock[change.Field] = change.NewValue
	data[change.Path.BlockName] = itemBlock
	return nil
}

func (s Service) deleteFieldOperationForArray(data map[string]any, change delta.Changelog) error {
	err := s.validateArrayPath(change.Path)
	if err != nil {
		return errors.WithMessage(err, "validate array path")
	}

	if change.Field == "" {
		return errors.New("field is empty")
	}

	dataBlock, exist := data[change.Path.BlockName]
	if !exist {
		return errors.Errorf("%s is not exist in data", change.Path.BlockName)
	}

	itemBlock, ok := dataBlock.([]any)
	if !ok {
		return errors.Errorf("%s is not array in data", change.Path.BlockName)
	}

	itemIdIsExist := false
	for _, item := range itemBlock {
		item, ok := item.(map[string]any)
		if !ok {
			return errors.Errorf("%s is not array of objects in data", change.Path.BlockName)
		}

		itemId, foundItemId := item[itemIdPath]
		if !foundItemId {
			return errors.Errorf("%s without %s in data", change.Path.BlockName, itemIdPath)
		}

		if itemId == change.Path.ItemId {
			_, found := item[change.Field]
			if !found {
				return nil // nothing to delete
			}

			itemIdIsExist = true
			delete(item, change.Field)
			break
		}
	}
	if !itemIdIsExist {
		return errors.Errorf("%s doesn't have %s = %s", change.Path.BlockName, itemIdPath, change.Path.ItemId)
	}

	data[change.Path.BlockName] = itemBlock
	return nil
}

func (s Service) deleteFieldOperationForObject(data map[string]any, change delta.Changelog) error {
	err := s.validateObjectPath(change.Path)
	if err != nil {
		return errors.WithMessage(err, "validate object path")
	}

	if change.Field == "" {
		return errors.New("field is empty")
	}

	dataBlock, exist := data[change.Path.BlockName]
	if !exist {
		return nil // nothing to delete
	}

	itemBlock, ok := dataBlock.(map[string]any)
	if !ok {
		return errors.Errorf("%s is not object in data", change.Path.BlockName)
	}

	_, found := itemBlock[change.Field]
	if !found {
		return nil // nothing to delete
	}

	delete(itemBlock, change.Field)
	data[change.Path.BlockName] = itemBlock
	return nil
}

func (s Service) deleteArrayItemOperation(data map[string]any, change delta.Changelog) error {
	err := s.validateArrayPath(change.Path)
	if err != nil {
		return errors.WithMessage(err, "validate array path")
	}

	dataBlock, exist := data[change.Path.BlockName]
	if !exist {
		return nil // nothing to delete
	}

	itemBlock, ok := dataBlock.([]any)
	if !ok {
		return errors.Errorf("%s is not array in data", change.Path.BlockName)
	}
	if len(itemBlock) == 0 {
		return nil // nothing to delete
	}

	deleteIndex := -1
	for i, item := range itemBlock {
		item, ok := item.(map[string]any)
		if !ok {
			return errors.Errorf("%s is not array of objects in data", change.Path.BlockName)
		}

		itemId, foundItemId := item[itemIdPath]
		if !foundItemId {
			return errors.Errorf("%s without %s in data", change.Path.BlockName, itemIdPath)
		}

		if itemId == change.Path.ItemId {
			deleteIndex = i
			break
		}
	}
	if deleteIndex == -1 {
		return nil // nothing to delete
	}

	data[change.Path.BlockName] = append(itemBlock[:deleteIndex], itemBlock[deleteIndex+1:]...)
	return nil
}

func (s Service) addArrayItemOperation(data map[string]any, change delta.Changelog) error {
	err := s.validateArrayPath(change.Path)
	if err != nil {
		return errors.WithMessage(err, "validate array path")
	}

	dataBlock, exist := data[change.Path.BlockName]
	if !exist {
		dataBlock = make([]any, 0)
	}

	itemBlock, ok := dataBlock.([]any)
	if !ok {
		return errors.Errorf("%s is not array in data", change.Path.BlockName)
	}

	for _, item := range itemBlock {
		item, ok := item.(map[string]any)
		if !ok {
			return errors.Errorf("%s is not array of objects in data", change.Path.BlockName)
		}

		itemId, foundItemId := item[itemIdPath]
		if !foundItemId {
			return errors.Errorf("%s without %s in data", change.Path.BlockName, itemIdPath)
		}

		if itemId == change.Path.ItemId {
			return errors.Errorf("%s already has %s %s", change.Path.BlockName, itemIdPath, change.Path.ItemId)
		}
	}

	itemBlock = append(itemBlock, map[string]any{
		itemIdPath: change.Path.ItemId,
	})
	data[change.Path.BlockName] = itemBlock
	return nil
}

func (s Service) validateObjectPath(path delta.Path) error {
	if path.BlockName == "" {
		return errors.New("empty block name")
	}

	if path.IsArray {
		return errors.Errorf("expetected object path for %s", path.BlockName)
	}

	if path.ItemId != "" {
		return errors.Errorf("itemId must be empty for %s ", path.BlockName)
	}

	return nil
}

func (s Service) validateArrayPath(path delta.Path) error {
	if path.BlockName == "" {
		return errors.New("empty blockName")
	}

	if !path.IsArray {
		return errors.Errorf("expetected array path for %s", path.BlockName)
	}

	if path.ItemId == "" {
		return errors.Errorf("itemId is required for array %s", path.BlockName)
	}

	return nil
}

func (s Service) deleteBlockOperation(data map[string]any, change delta.Changelog) error {
	err := s.validateObjectPath(change.Path)
	if err != nil {
		return errors.WithMessage(err, "validate object path")
	}

	delete(data, change.Path.BlockName)

	return nil
}
