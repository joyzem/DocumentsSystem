package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/joyzem/documents/services/product/domain"
	"github.com/joyzem/documents/services/product/dto"
	"github.com/joyzem/documents/services/product/frontend/utils"
	"github.com/levigross/grequests"
)

// Обработчик запроса на страницу товаров
func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	// Получение адреса сервера backend
	// Так как адрес сервиса товаров будет часто использоваться в обработчиках,
	// будет целесообразно вынести его в отдельную функцию.
	// Это также необходимо и по той причине, что при создании Docker-контейнера
	// адрес бэкэнд сервера изменится
	productsUrl := fmt.Sprintf("%s/products", utils.GetBackendAddress())
	// Выполнение GET-запроса на сервер backend
	resp, err := grequests.Get(productsUrl, nil)
	if err != nil {
		// Обработка ошибки в случае неудачного запроса
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var products dto.GetProductsResponse
	// Декодирование ответа от сервера backend в структуру GetProductsResponse
	if err := resp.JSON(&products); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if products.Err != "" {
		// Обработка ошибки, переданной в ответе сервера backend
		http.Error(w, products.Err, http.StatusInternalServerError)
		return
	}
	// Получение списка единиц измерения из backend
	// Вынесен в отдельную функцию для повторных вызовов
	units, err := utils.GetUnitsFromBackend()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if units.Err != "" {
		http.Error(w, units.Err, http.StatusInternalServerError)
		return
	}
	// Создание списка товаров с указанием соответствующих единиц измерения
	type ProductItemTemplate struct {
		Product domain.Product
		Unit    domain.Unit
	}
	templateItems := []ProductItemTemplate{}
	for _, product := range products.Products {
		templateItem := ProductItemTemplate{}
		templateItem.Product = product
		var productUnit domain.Unit
		for _, unit := range units.Units {
			if unit.Id == product.UnitId {
				productUnit = unit
				break
			}
		}
		templateItem.Unit = productUnit
		templateItems = append(templateItems, templateItem)
	}

	// Загрузка шаблона страницы товаров и его рендеринг с передачей списка товаров в качестве аргумента
	productPage, _ := template.ParseFiles("../static/html/products.html")
	productPage.Execute(w, templateItems)
}

// Удаление товара
func DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем значение id товара, который нужно удалить из формы
	id, _ := strconv.Atoi(r.FormValue("id"))
	// Составляем URL-адрес, чтобы отправить запрос на удаление товара на бэкэнд
	productsUrl := fmt.Sprintf("%s/products", utils.GetBackendAddress())

	// Отправляем запрос на удаление товара на бэкэнд с помощью библиотеки grequests
	// Вторым аргументом передаем объект grequests.RequestOptions, содержащий JSON структуру
	resp, _ := grequests.Delete(productsUrl, &grequests.RequestOptions{
		JSON: dto.DeleteProductRequest{Id: id},
	})

	// Получаем ответ в формате JSON и декодируем его в структуру deleteResponse
	var deleteResponse dto.DeleteProductResponse
	resp.JSON(&deleteResponse)

	// Если при удалении товара произошла ошибка, возвращаем ошибку в ответе HTTP
	if deleteResponse.Err != "" {
		http.Error(w, deleteResponse.Err, http.StatusInternalServerError)
		return
	}

	// Если товар успешно удален, перенаправляем пользователя на страницу со списком товаров
	http.Redirect(w, r, "/product/products", http.StatusSeeOther)
}

// Обработчик страницы добавления товара
func CreateProductGetHandler(w http.ResponseWriter, r *http.Request) {
	// получить единицы измерения
	units, _ := utils.GetUnitsFromBackend()
	if units.Err != "" {
		http.Error(w, units.Err, http.StatusInternalServerError)
		return
	}
	// шаблон добавления товара
	createProductPage, _ := template.ParseFiles("../static/html/create-product.html")
	createProductPage.Execute(w, units.Units)
}

// Обработчик запроса на добавление товара
func CreateProductPostHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем параметры товара из формы
	productName := r.FormValue("name")                             // наименование товара
	productPrice, priceError := strconv.Atoi(r.FormValue("price")) // цена товара
	unitId, unitErr := strconv.Atoi(r.FormValue("unit_id"))        // id ед.изм.
	// Проверяем корректность параметров товара
	if productName == "" || priceError != nil || unitErr != nil {
		http.Error(w, "fields validation error", http.StatusUnprocessableEntity)
		return
	}

	// Создаем запрос на добавление товара в бэкэнд
	request := dto.CreateProductRequest{
		Name:   productName,
		Price:  productPrice,
		UnitId: unitId,
	}
	productUrl := fmt.Sprintf("%s/products", utils.GetBackendAddress())
	resp, _ := grequests.Post(productUrl, &grequests.RequestOptions{
		JSON: request,
	})

	// Обрабатываем ответ от бэкэнда
	var data dto.CreateProductResponse
	resp.JSON(&data)
	if data.Err != "" {
		http.Error(w, data.Err, http.StatusInternalServerError)
	} else {
		http.Redirect(w, r, "/product/products", http.StatusSeeOther)
	}
}

// Обработчик страницы обновления продукта
func UpdateProductGetHandler(w http.ResponseWriter, r *http.Request) {
	// парсинг id из адреса страницы
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	// получение товара с бэка
	url := fmt.Sprintf("%s/products/%d", utils.GetBackendAddress(), id)
	resp, _ := grequests.Get(url, &grequests.RequestOptions{
		JSON: dto.ProductByIdRequest{Id: id},
	})
	var product dto.ProductByIdResponse
	resp.JSON(&product)
	// товар не найден
	if product.Err != "" {
		http.Error(w, product.Err, http.StatusBadRequest)
		return
	}
	// получение единиц измерения с бэка
	unitsResp, _ := utils.GetUnitsFromBackend()
	if unitsResp.Err != "" {
		http.Error(w, unitsResp.Err, http.StatusInternalServerError)
		return
	}
	// структура данных для шаблона
	type UpdateProductTemplate struct {
		Product *domain.Product
		Units   []domain.Unit
	}
	data := UpdateProductTemplate{
		Product: product.Product,
		Units:   unitsResp.Units,
	}
	// шаблон страницы
	updateProductPage, _ := template.ParseFiles("../static/html/update-product.html")
	updateProductPage.Execute(w, data)
}

// Обработчик запроса на обновление товара
func UpdateProductPostHandler(w http.ResponseWriter, r *http.Request) {
	productId, _ := strconv.Atoi(r.FormValue("id"))                // получаем id товара, который нужно обновить
	productName := r.FormValue("name")                             // получаем имя товара из формы
	productPrice, priceError := strconv.Atoi(r.FormValue("price")) // получаем цену товара из формы
	unitId, unitErr := strconv.Atoi(r.FormValue("unit_id"))        // получаем id единицы измерения из формы
	// валидация формы
	if len(productName) == 0 || priceError != nil || unitErr != nil { // проверяем, что поля заполнены корректно
		http.Error(w, "fields_validation_error", http.StatusUnprocessableEntity) // возвращаем ошибку
		return
	}

	product := domain.Product{ // создаем новый объект товара для обновления
		Id:     productId,
		Name:   productName,
		Price:  productPrice,
		UnitId: unitId,
	}

	productUrl := fmt.Sprintf("%s/products", utils.GetBackendAddress()) // формируем URL для запроса к API

	request := dto.UpdateProductRequest{ // создаем объект запроса на обновление товара
		Product: product,
	}

	resp, _ := grequests.Put(productUrl, &grequests.RequestOptions{ // отправляем PUT запрос на API
		JSON: request,
	})

	var updateResponse dto.UpdateProductResponse // создаем объект ответа на запрос обновления
	resp.JSON(&updateResponse)                   // получаем ответ от API и парсим его в объект ответа на запрос обновления
	if updateResponse.Err != "" {                // если в ответе есть ошибка, то возвращаем ее
		http.Error(w, updateResponse.Err, http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/product/products", http.StatusSeeOther) // перенаправляем пользователя на страницу товаров

}
