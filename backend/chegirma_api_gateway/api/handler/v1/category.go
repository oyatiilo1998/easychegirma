package v1

import (
	"context"
	"net/http"

	"chegirma_api_gateway/ek_variables/ek_setting_service"
	"chegirma_api_gateway/genproto/setting_service"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// @Router /v1/category [post]
// @Summary Create category
// @Description API for creating category
// @Tags category
// @Accept json
// @Produce json
// @Param category body ek_setting_service.CreateUpdateCategorySwag  true "Category"
// @Success 201 {object} setting_service.CreatedResponse
// @Failure 400 {object} ek_variables.FailureResponse
// @Failure 404 {object} ek_variables.FailureResponse
// @Failure 500 {object} ek_variables.FailureResponse
// @Failure 503 {object} ek_variables.FailureResponse
func (h *handlerV1) CreateCategory(c *gin.Context) {
	var (
		category   ek_setting_service.CreateUpdateCategorySwag
		categoryID = primitive.NewObjectID().Hex()
		// userInfo, _ = h.UserInfo(c, true)
	)

	if err := c.BindJSON(&category); HandleHTTPError(c, http.StatusBadRequest, "DiscussionLogicService.Action.Create.BindingAction", err) {
		return
	}
	resp, err := h.services.CategoryService().Create(
		context.Background(),
		&setting_service.Category{
			Id:     categoryID,
			Name:   category.Name,
			RuName: category.RuName,
			Code:   category.Code,
		},
	)

	if HandleRPCError(c, "SettingService.CreateCategory.InterService", err) {
		return
	}
	// actionHistory := &ek_analytic_service.ActionHistory{
	// 	UserID:         userInfo.ID,
	// 	UserUniqueName: userInfo.Login,
	// 	Action:         ek_variables.EntityCreated,
	// 	EntityID:       categoryID,
	// 	EntityName:     "entity",
	// }
	// h.CreateActionHistory(c, actionHistory)

	c.JSON(http.StatusCreated, resp)
}

// @Router /v1/category/{category_id} [get]
// @Summary Get Category
// @Description API for getting category
// @Tags category
// @Accept json
// @Produce json
// @Param category_id path string  true "category_id"
// @Success 200 {object} setting_service.Category
// @Failure 400 {object} ek_variables.FailureResponse
// @Failure 404 {object} ek_variables.FailureResponse
// @Failure 500 {object} ek_variables.FailureResponse
// @Failure 503 {object} ek_variables.FailureResponse
func (h *handlerV1) GetCategory(c *gin.Context) {
	var (
		categoryID = c.Param("category_id")
		_, err     = primitive.ObjectIDFromHex(categoryID)
	)
	if HandleHTTPError(c, http.StatusBadRequest, "SettingService.GetCategory.ParseCategoryID", err) {
		return
	}
	category, err := h.services.CategoryService().Get(context.Background(), &setting_service.CategoryGetReq{
		Id: categoryID,
	})

	if HandleRPCError(c, "SettingService.GetCategory.InterService", err) {
		return
	}

	c.JSON(http.StatusOK, category)
}

// @Router /v1/category [get]
// @Summary Getting All categories
// @Description API for getting all categoryes
// @Tags category
// @Accept json
// @Produce json
// @Param name query string  false "name"
// @Param soato query string  false "soato"
// @Param page query integer false "page"
// @Param limit query integer false "limit"
// @Success 200 {object} setting_service.GetAllCategoriesResponse
// @Failure 400 {object} ek_variables.FailureResponse
// @Failure 404 {object} ek_variables.FailureResponse
// @Failure 500 {object} ek_variables.FailureResponse
// @Failure 503 {object} ek_variables.FailureResponse
func (h *handlerV1) GetAllCategories(c *gin.Context) {
	var (
		name       = c.Query("name")
		soatoQuery = c.Query("soato")
		soato      int
	)
	page, err := ParseQueryParam(c, h.log, "page", "1")
	if err != nil {
		return
	}

	limit, err := ParseQueryParam(c, h.log, "limit", "20")
	if err != nil {
		return
	}
	if soatoQuery != "" {
		soato, err = ParseQueryParam(c, h.log, "soato", "0")
		if err != nil {
			return
		}
	}
	categories, err := h.services.CategoryService().GetAll(
		context.Background(),
		&setting_service.GetAllCategoriesRequest{
			Name:  name,
			Soato: uint32(soato),
			Page:  uint32(page),
			Limit: uint32(limit),
		})

	if HandleRPCError(c, "SettingService.GetAllCategories.InternalService", err) {
		return
	}

	c.JSON(http.StatusOK, categories)
}

// @Router /v1/category/{category_id} [put]
// @Summary Update category
// @Description API for updating category
// @Tags category
// @Accept json
// @Produce json
// @Param category_id path string  true "category_id"
// @Param category body ek_setting_service.CreateUpdateCategorySwag true "category"
// @Success 200 {object} setting_service.CreatedResponse
// @Failure 400 {object} ek_variables.FailureResponse
// @Failure 404 {object} ek_variables.FailureResponse
// @Failure 500 {object} ek_variables.FailureResponse
// @Failure 503 {object} ek_variables.FailureResponse
func (h *handlerV1) UpdateCategory(c *gin.Context) {
	var (
		category   setting_service.Category
		categoryID = c.Param("category_id")
		_, err     = primitive.ObjectIDFromHex(categoryID)
	)
	if HandleHTTPError(c, http.StatusBadRequest, "SettingService.UpdateCategory.ParsingObjectID", err) {
		return
	}
	if err = c.ShouldBindJSON(&category); HandleHTTPError(c, http.StatusBadRequest, "SettingService.UpdateCategory.InvalidRequest", err) {
		return
	}
	category.Id = categoryID

	resp, err := h.services.CategoryService().Update(
		context.Background(),
		&category)

	if HandleRPCError(c, "SettingService.UpdateCategory.InternalService", err) {
		return
	}

	c.JSON(http.StatusOK, resp)
}
