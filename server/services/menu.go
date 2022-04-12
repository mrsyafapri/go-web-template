package services

import (
	"database/sql"
	"go-web-template/server/models"
	params "go-web-template/server/params/menu"
	repositories "go-web-template/server/repositories/menu"
	"time"
)

type MenuServices struct {
	MenuRepository repositories.MenuRepository
	DB             *sql.DB
}

func NewMenuService(db *sql.DB) *MenuServices {
	repository := repositories.NewMenuRepository(db)
	return &MenuServices{
		MenuRepository: repository,
		DB:             db,
	}
}

func (m *MenuServices) CreateNewMenu(request *params.MenuCreate) (bool, error) {
	menu := request.ParseToModel()
	err := m.MenuRepository.Save(menu)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (m *MenuServices) GetAllMenu() (*[]params.MenuSingleView, error) {
	menus, err := m.MenuRepository.FindAll()

	if err != nil {
		return nil, err
	}

	return makeMenuListView(menus), nil
}

func (m *MenuServices) GetMenuByID(id string) (*params.MenuSingleView, error) {
	menu, err := m.MenuRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	return makeMenuSingleView(menu), nil
}

func (m *MenuServices) UpdateMenuByID(requrest *params.MenuUpdate) (bool, error) {
	model := requrest.ParseToModel()
	err := m.MenuRepository.UpdateByID(model)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (m *MenuServices) DeleteMenuByID(id string) (bool, error) {
	err := m.MenuRepository.DeleteByID(id)

	if err != nil {
		return false, err
	}

	return true, nil
}

func makeMenuListView(models *[]models.Menu) *[]params.MenuSingleView {
	var menuListview []params.MenuSingleView
	for _, model := range *models {
		menuListview = append(menuListview, *makeMenuSingleView(&model))
	}
	return &menuListview
}

func makeMenuSingleView(model *models.Menu) *params.MenuSingleView {
	return &params.MenuSingleView{
		ID:        model.ID,
		Name:      model.Name,
		Category:  model.Category,
		Desc:      model.Desc,
		CreatedAt: model.CreatedAt.Format(time.RFC3339),
		UpdatedAt: model.UpdatedAt.Format(time.RFC3339),
	}
}
