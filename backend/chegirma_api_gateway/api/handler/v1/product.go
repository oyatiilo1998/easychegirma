package v1

import (
	"context"
	"net/http"

	"chegirma_api_gateway/ek_variables"
	"chegirma_api_gateway/ek_variables/chegirma_setting_service"
	"chegirma_api_gateway/genproto/setting_service"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// @Router /v1/product [post]
// @Summary Create product
// @Description API for creating product
// @Tags product
// @Accept json
// @Produce json
// @Param product body chegirma_setting_service.CreateUpdateProductSwag  true "Product"
// @Success 201 {object} setting_service.CreatedResponse
// @Failure 400 {object} ek_variables.FailureResponse
// @Failure 404 {object} ek_variables.FailureResponse
// @Failure 500 {object} ek_variables.FailureResponse
// @Failure 503 {object} ek_variables.FailureResponse
func (h *handlerV1) CreateProduct(c *gin.Context) {
	var (
		product   chegirma_setting_service.CreateUpdateProductSwag
		productID = primitive.NewObjectID().Hex()
		// userInfo, _ = h.UserInfo(c, true)
	)

	if err := c.BindJSON(&product); HandleHTTPError(c, http.StatusBadRequest, "DiscussionLogicService.Action.Create.BindingAction", err) {
		return
	}
	resp, err := h.services.ProductService().Create(
		context.Background(),
		&setting_service.CreateUpdateProduct{
			Id:             productID,
			Name:           product.Name,
			Description:    product.Description,
			Url:            product.URL,
			ImageUrl:       product.ImageURL,
			DiscountAmount: uint32(product.DiscountAmount),
			PriceBefore:    uint64(product.PriceBefore),
			PriceAfter:     uint64(product.PriceAfter),
			CategoryId:     product.CategoryID,
			OwnerId:        product.OwnerId,
			FromTime:       product.FromTime.Format(ek_variables.TimeLayout),
			ToTime:         product.ToTime.Format(ek_variables.TimeLayout),
		},
	)

	if HandleRPCError(c, "SettingService.CreateProduct.InterService", err) {
		return
	}
	// actionHistory := &ek_analytic_service.ActionHistory{
	// 	UserID:         userInfo.ID,
	// 	UserUniqueName: userInfo.Login,
	// 	Action:         ek_variables.EntityCreated,
	// 	EntityID:       productID,
	// 	EntityName:     "entity",
	// }
	// h.CreateActionHistory(c, actionHistory)

	c.JSON(http.StatusCreated, resp)
}

// @Router /v1/product/{product_id} [get]
// @Summary Get Product
// @Description API for getting product
// @Tags product
// @Accept json
// @Produce json
// @Param product_id path string  true "product_id"
// @Success 200 {object} setting_service.Product
// @Failure 400 {object} ek_variables.FailureResponse
// @Failure 404 {object} ek_variables.FailureResponse
// @Failure 500 {object} ek_variables.FailureResponse
// @Failure 503 {object} ek_variables.FailureResponse
func (h *handlerV1) GetProduct(c *gin.Context) {
	var (
		productID = c.Param("product_id")
		_, err    = primitive.ObjectIDFromHex(productID)
	)
	if HandleHTTPError(c, http.StatusBadRequest, "SettingService.GetProduct.ParseProductID", err) {
		return
	}
	product, err := h.services.ProductService().Get(context.Background(), &setting_service.GetReq{
		Id: productID,
	})

	if HandleRPCError(c, "SettingService.GetProduct.InterService", err) {
		return
	}

	c.JSON(http.StatusOK, product)
}

// @Router /v1/product [get]
// @Summary Getting All products
// @Description API for getting all productes
// @Tags product
// @Accept json
// @Produce json
// @Param name query string  false "name"
// @Param category_id query string  false "category_id"
// @Param owner_id query string  false "owner_id"
// @Param page query integer false "page"
// @Param limit query integer false "limit"
// @Success 200 {object} setting_service.GetAllProductsResponse
// @Failure 400 {object} ek_variables.FailureResponse
// @Failure 404 {object} ek_variables.FailureResponse
// @Failure 500 {object} ek_variables.FailureResponse
// @Failure 503 {object} ek_variables.FailureResponse
func (h *handlerV1) GetAllProducts(c *gin.Context) {
	var (
		name       = c.Query("name")
		categoryID = c.Query("category_id")
		ownerID    = c.Query("owner_id")
	)
	page, err := ParseQueryParam(c, h.log, "page", "1")
	if err != nil {
		return
	}

	limit, err := ParseQueryParam(c, h.log, "limit", "20")
	if err != nil {
		return
	}
	products, err := h.services.ProductService().GetAll(
		context.Background(),
		&setting_service.GetAllProductsRequest{
			Name:       name,
			CategoryId: categoryID,
			OwnerId:    ownerID,
			Page:       uint32(page),
			Limit:      uint32(limit),
		})

	if HandleRPCError(c, "SettingService.GetAllProducts.InternalService", err) {
		return
	}

	c.JSON(http.StatusOK, products)
}

// @Router /v1/product/{product_id} [put]
// @Summary Update product
// @Description API for updating product
// @Tags product
// @Accept json
// @Produce json
// @Param product_id path string  true "product_id"
// @Param product body chegirma_setting_service.CreateUpdateProductSwag true "product"
// @Success 200 {object} setting_service.CreatedResponse
// @Failure 400 {object} ek_variables.FailureResponse
// @Failure 404 {object} ek_variables.FailureResponse
// @Failure 500 {object} ek_variables.FailureResponse
// @Failure 503 {object} ek_variables.FailureResponse
func (h *handlerV1) UpdateProduct(c *gin.Context) {
	var (
		product   setting_service.CreateUpdateProduct
		productID = c.Param("product_id")
		_, err    = primitive.ObjectIDFromHex(productID)
	)
	if HandleHTTPError(c, http.StatusBadRequest, "SettingService.UpdateProduct.ParsingObjectID", err) {
		return
	}
	if err = c.ShouldBindJSON(&product); HandleHTTPError(c, http.StatusBadRequest, "SettingService.UpdateProduct.InvalidRequest", err) {
		return
	}
	product.Id = productID

	resp, err := h.services.ProductService().Update(
		context.Background(),
		&product)

	if HandleRPCError(c, "SettingService.UpdateProduct.InternalService", err) {
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router /v1/product/{product_id} [delete]
// @Summary Delete product
// @Description API for deleting product
// @Tags product
// @Accept json
// @Produce json
// @Param product_id path string  true "product_id"
// @Success 200 {object} ek_variables.SuccessResponse
// @Failure 400 {object} ek_variables.FailureResponse
// @Failure 404 {object} ek_variables.FailureResponse
// @Failure 500 {object} ek_variables.FailureResponse
// @Failure 503 {object} ek_variables.FailureResponse
func (h *handlerV1) DeleteProduct(c *gin.Context) {
	productID := c.Param("product_id")

	_, err := primitive.ObjectIDFromHex(productID)
	if HandleHTTPError(c, http.StatusBadRequest, "Error while parsing uuid, product id incorrect format", err) {
		return
	}

	product, err := h.services.ProductService().Delete(
		context.Background(),
		&setting_service.DeleteProductRequest{Id: productID})

	if HandleRPCError(c, "ProductService.Product.DeleteProduct", err) {
		return
	}

	c.JSON(http.StatusOK, product)
}
