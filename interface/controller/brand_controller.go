package controller

import (
	"go_worlder_system/errs"
	"go_worlder_system/interface/gateway/database"
	"go_worlder_system/interface/presenter"
	inputdata "go_worlder_system/usecase/input/data"
	inputport "go_worlder_system/usecase/input/port"
	"go_worlder_system/usecase/interactor"
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"
)

// brandParam ...
type brandParam struct {
	UserID      string
	BrandName   string
	CategoryID  string
	Slogan      string
	Description string
	LogoImage   string
	IsDraft     string
}

// BrandController ...
type BrandController struct {
	inputport inputport.BrandInputPort
	param     *brandParam
}

// NewBrandController ...
func NewBrandController(sqlHandler database.SQLHandler) *BrandController {
	param := &brandParam{}
	initializeParam(param)
	return &BrandController{
		inputport: interactor.NewBrandInteractor(
			presenter.NewBrandPresenter(),
			database.NewUserDatabase(sqlHandler),
			database.NewBrandDatabase(sqlHandler),
			database.NewProductDatabase(sqlHandler),
			database.NewCategoryDatabase(sqlHandler),
		),
		param: param,
	}
}

// Index ...
// @summary Get brands
// @description Get brands by user id
// @tags Brand
// @produce json
// @param Authorization header string true "jwt token"
// @success 200 {array} outputdata.Brand "Brand list"
// @failure 404 {string} string "Brand list is not found"
// @router /brands [get]
func (controller *BrandController) Index(c Context) error {
	userID := c.UserID()
	oBrandList, err := controller.inputport.Index(userID)
	if err != nil {
		log.Println(err)
		c.String(statusCode(err), errs.Cause(err).Error())
		return err
	}
	return c.JSON(http.StatusOK, oBrandList)
}

// Show ...
// @summary Get a brand
// @description Get a brand by id
// @tags Brand
// @produce json
// @param Authorization header string true "jwt token"
// @param id path string true "Brand ID"
// @success 200 {object} outputdata.Brand "Brand"
// @failure 404 {string} string "The Brand is not found"
// @router /brands/{id} [get]
func (controller *BrandController) Show(c Context) error {
	id := c.Param(idParam)
	oBrand, err := controller.inputport.Show(id)
	if err != nil {
		log.Println(err)
		c.String(statusCode(err), errs.Cause(err).Error())
		return err
	}
	return c.JSON(http.StatusOK, oBrand)
}

// New ...
// @summary
// @description
// @tags Brand
// @produce json
// @param Authorization header string true "jwt token"
// @success 200
// @router /brands/new [get]
func (controller *BrandController) New(c Context) error {
	return c.JSON(http.StatusOK, nil)
}

// Edit ...
// @summary Get the brand
// @description Get the brand if the user is owner
// @tags Brand
// @produce json
// @param Authorization header string true "jwt token"
// @param id path string true "Brand ID:registered"
// @success 200 {object} outputdata.Brand "brand"
// @failure 404 {string} string "The user can't get the brand"
// @router /brands/{id}/edit [get]
func (controller *BrandController) Edit(c Context) error {
	id := c.Param(idParam)
	userID := c.UserID()
	oBrand, err := controller.inputport.Edit(id, userID)
	if err != nil {
		log.Println(err)
		c.String(statusCode(err), errs.Cause(err).Error())
		return err
	}
	return c.JSON(http.StatusOK, oBrand)
}

// Create ...
// @summary Register a new brand
// @description Register a new brand. ID is generated by random string generation
// @tags Brand
// @accept mpfd
// @produce json
// @param Authorization header string true "jwt token"
// @param brand-name formData string true "brand name"
// @param category-id formData int true "category id: min=1, max=1"
// @param slogan formData string true "slogan"
// @param description formData string true "description"
// @param logo-image formData file false "brand logo image"
// @param is-draft formData boolean true "draft or not"
// @success 200 {string} outputdata.Brand "brand"
// @failure 409 {string} string "The user can't register a brand"
// @router /brands [post]
func (controller *BrandController) Create(c Context) error {
	userID := c.UserID()
	categoryID, err := strconv.ParseUint(c.FormValue(controller.param.CategoryID), 10, 64)
	if err != nil {
		log.WithFields(log.Fields{}).Error(err)
		return err
	}
	name := c.FormValue(controller.param.BrandName)
	slogan := c.FormValue(controller.param.Slogan)
	description := c.FormValue(controller.param.Description)
	logoImage, _ := c.FormFile(controller.param.LogoImage)
	isDraft, err := strconv.ParseBool(c.FormValue(controller.param.IsDraft))
	if err != nil {
		log.WithFields(log.Fields{}).Error(err)
		c.String(statusCode(err), errs.Cause(err).Error())
		return err
	}
	iBrand := &inputdata.Brand{
		UserID:      userID,
		Name:        name,
		CategoryID:  uint(categoryID),
		Slogan:      slogan,
		Description: description,
		LogoImage:   logoImage,
		IsDraft:     isDraft,
	}
	iNewBrand := &inputdata.NewBrand{Brand: iBrand}
	err = controller.inputport.Create(iNewBrand)
	if err != nil {
		log.Println(err)
		c.String(statusCode(err), errs.Cause(err).Error())
		return err
	}
	return c.JSON(http.StatusOK, iBrand)
}

// Update ...
// @summary Edit a brand
// @description Edit a brand got by the brand id
// @tags Brand
// @accept mpfd
// @produce json
// @param Authorization header string true "jwt token"
// @param id path string true "Brand ID"
// @param brand-name formData string true "brand name"
// @param category-id formData int true "category id: min=1, max=1"
// @param slogan formData string true "slogan"
// @param description formData string true "description"
// @param logo-image formData file false "brand logo image"
// @param is-draft formData boolean true "draft or not"
// @success 200 {object} outputdata.Brand "brand"
// @failure 409 {string} string "The user can't edit the brand"
// @router /brands/{id} [patch]
func (controller *BrandController) Update(c Context) error {
	id := c.Param(idParam)
	userID := c.UserID()
	categoryID, err := strconv.ParseUint(c.FormValue(controller.param.CategoryID), 10, 64)
	if err != nil {
		log.WithFields(log.Fields{}).Error(err)
		return err
	}
	name := c.FormValue(controller.param.BrandName)
	slogan := c.FormValue(controller.param.Slogan)
	description := c.FormValue(controller.param.Description)
	logoImage, _ := c.FormFile(controller.param.LogoImage)
	isDraft, err := strconv.ParseBool(c.FormValue(controller.param.IsDraft))
	if err != nil {
		log.Println(err)
		c.String(statusCode(err), errs.Cause(err).Error())
		return err
	}
	iBrand := &inputdata.Brand{
		UserID:      userID,
		Name:        name,
		CategoryID:  uint(categoryID),
		Slogan:      slogan,
		Description: description,
		LogoImage:   logoImage,
		IsDraft:     isDraft,
	}
	iUpdatedBrand := &inputdata.UpdatedBrand{
		ID:    id,
		Brand: iBrand,
	}
	err = controller.inputport.Update(iUpdatedBrand)
	if err != nil {
		log.Println(err)
		c.String(statusCode(err), errs.Cause(err).Error())
		return err
	}
	return c.JSON(http.StatusOK, iBrand)
}

// Delete ...
// @summary Delete a brand
// @description Delete a brand got by the id
// @tags Brand
// @produce json
// @param Authorization header string true "jwt token"
// @param id path string true "Brand ID:registered"
// @success 200
// @failure 409 {string} string "The user can't delete the brand"
// @router /brands/{id} [delete]
func (controller *BrandController) Delete(c Context) error {
	id := c.Param(idParam)
	userID := c.UserID()
	err := controller.inputport.Delete(id, userID)
	if err != nil {
		log.Println(err)
		c.String(statusCode(err), errs.Cause(err).Error())
		return err
	}
	return c.JSON(http.StatusOK, nil)
}