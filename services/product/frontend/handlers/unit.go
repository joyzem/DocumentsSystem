package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/joyzem/documents/services/base"
	"github.com/joyzem/documents/services/product/domain"
	"github.com/joyzem/documents/services/product/dto"
	"github.com/joyzem/documents/services/product/frontend/utils"
	"github.com/levigross/grequests"
)

// Обработчик страницы единиц измерения
func UnitsHandler(w http.ResponseWriter, r *http.Request) {
	// получение ед. измерения
	unitsRepo, _ := utils.GetUnitsFromBackend()
	if unitsRepo.Err != "" {
		http.Error(w, unitsRepo.Err, http.StatusInternalServerError)
		return
	}
	// создание шаблона
	unitsPage, _ := template.ParseFiles("../static/html/units.html")
	unitsPage.Execute(w, unitsRepo.Units)
}

// Обработчик удаления единицы измерения
func DeleteUnitHandler(w http.ResponseWriter, r *http.Request) {
	// парсинг id
	id, _ := strconv.Atoi(r.FormValue("id"))
	// тело запроса
	body := dto.DeleteUnitRequest{Id: id}
	// адрес ед.измерения
	unitsUrl := fmt.Sprintf("%s/units", utils.GetBackendAddress())
	// отправка запроса и получение ответа
	resp, _ := grequests.Delete(unitsUrl, &grequests.RequestOptions{
		JSON: body,
	})
	// парсинг ответа
	var deleteResponse dto.DeleteUnitResponse
	resp.JSON(&deleteResponse)
	// проверка на ошибку
	if deleteResponse.Err != "" {
		http.Error(w, deleteResponse.Err, http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/product/units", http.StatusSeeOther)
}

// Обработчик страницы добавления единицы измерения
func CreateUnitGetHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../static/html/create-unit.html")
}

// Обработчик добавления единицы измерения
func CreateUnitPostHandler(w http.ResponseWriter, r *http.Request) {
	// парсинг наименования
	unitName := r.FormValue("name")
	if unitName == "" {
		http.Error(w, base.FIELDS_VALIDATION_ERROR, http.StatusUnprocessableEntity)
		return
	}
	// создание запроса
	request := dto.CreateUnitRequest{
		Unit: unitName,
	}
	// адрес бэка
	unitsUrl := fmt.Sprintf("%s/units", utils.GetBackendAddress())
	resp, _ := grequests.Post(unitsUrl, &grequests.RequestOptions{
		JSON: request,
	})
	// парсинг ответа
	var data dto.CreateUnitResponse
	resp.JSON(&data)
	// проверка на ошибку
	if data.Err == "" {
		http.Redirect(w, r, "/product/units", http.StatusSeeOther)
	} else {
		http.Error(w, data.Err, http.StatusInternalServerError)
	}
}

// Обработчик страницы изменения единицы измерения
func UpdateUnitGetHandler(w http.ResponseWriter, r *http.Request) {
	// парсинг id
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	// получение единицы измерения с бэка
	url := fmt.Sprintf("%s/units/%d", utils.GetBackendAddress(), id)
	resp, _ := grequests.Get(url, &grequests.RequestOptions{
		JSON: dto.UnitByIdRequest{Id: id},
	})
	// поиск запрашиваемой единицы измерения
	var unit dto.UnitByIdResponse
	// unit не найден
	resp.JSON(&unit)
	if unit.Unit == nil {
		http.Error(w, "the unit does not exist", http.StatusBadRequest)
		return
	}
	// создание шаблона
	updateUnitPage, _ := template.ParseFiles("../static/html/update-unit.html")
	updateUnitPage.Execute(w, unit.Unit)
}

// Обработчик добавления единицы измерения
func UpdateUnitPostHandler(w http.ResponseWriter, r *http.Request) {
	// парсинг полей
	unitId, err := strconv.Atoi(r.FormValue("id"))
	unitName := r.FormValue("name")
	if unitName == "" || err != nil {
		http.Error(w, base.FIELDS_VALIDATION_ERROR, http.StatusUnprocessableEntity)
		return
	}
	unit := domain.Unit{
		Id:   unitId,
		Name: unitName,
	}
	// адрес бэка
	unitsUrl := fmt.Sprintf("%s/units", utils.GetBackendAddress())
	// создание запроса
	request := dto.UpdateUnitRequest{
		Unit: unit,
	}
	// отправка запроса и получение ответа
	updateResponse, _ := grequests.Put(unitsUrl, &grequests.RequestOptions{
		JSON: request,
	})
	// парсинг ответа
	var data dto.UpdateUnitResponse
	updateResponse.JSON(&data)
	if data.Err != "" {
		http.Error(w, data.Err, http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/product/units", http.StatusSeeOther)
}
