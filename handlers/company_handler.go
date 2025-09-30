package handlers

import (
	"net/http"
	"strconv"

	"example-go-gin/domain"
	"example-go-gin/repository"

	"github.com/gin-gonic/gin"
)

type CompanyHandler struct {
	comRepository *repository.CompanyRepository
}

func NewCompanyHandler(comRepository *repository.CompanyRepository) *CompanyHandler {
	return &CompanyHandler{comRepository}
}

// @Summary get all items in the Company list
// @ID get-all-companies
// @Description  Return all Companies
// @Tags         Company
// @Accept       json
// @Produce json
// @Success 200 {object} domain.Company
// @Failure 404 {object} domain.Message
// @Router /companies [get]
func (ch *CompanyHandler) GetCompanies(c *gin.Context) {
	companies, err := ch.comRepository.GetCompanies()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, companies)
}

// @Summary get item by id in the Company list
// @ID get-company-by-id
// @Description  Return Company by Id
// @Tags         Company
// @Accept       json
// @Produce      json
// @Success 200 {object} domain.Company
// @Failure 404 {object} domain.Message
// @Router /companies/{id} [get]
func (uh *CompanyHandler) GetUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	user, err := uh.comRepository.FindByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// @Summary create a company item
// @ID create-company
// @Description  Create a Company
// @Tags         Company
// @Accept       json
// @Produce json
// @Param company body domain.Company true "Company item"
// @Success 201 {object} domain.Company
// @Failure 400 {object} domain.Message
// @Router /companies [post]
func (uh *CompanyHandler) CreateUser(c *gin.Context) {
	var user domain.Company
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := uh.comRepository.Save(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

// @Summary create a company item
// @ID create-company
// @Description  Update a Company
// @Tags         Company
// @Accept       json
// @Produce json
// @Param company body domain.Company true "Company item"
// @Success 201 {object} domain.Company
// @Failure 400 {object} domain.Message
// @Router /companies [post]
func (uh *CompanyHandler) UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID", "id": id})
		return
	}
	var user domain.Company
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//user.ID = id
	err = uh.comRepository.Update(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

// @Summary delete a company item by ID
// @ID delete-company-by-id
// @Description  Delete a Company
// @Tags         Company
// @Accept       json
// @Produce json
// @Param id path string true "company ID"
// @Success 200 {object} domain.Company
// @Failure 404 {object} domain.Message
// @Router /companies/{id} [delete]
func (uh *CompanyHandler) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	err = uh.comRepository.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// return error message if todo is not found
	//var message domain.Message

	r := domain.Message{"User deleted successfully"}

	//c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
	c.JSON(http.StatusOK, r)
}
