package httpadapter

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"contract-service/models"
	"contract-service/service"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	service *service.ContractService
}

func NewHandlers(svc *service.ContractService) *Handlers {
	return &Handlers{service: svc}
}

func (h *Handlers) Register(router *gin.Engine) {
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	persons := router.Group("/persons")
	persons.POST("", h.createPerson)
	persons.POST("/search", h.searchPersonContracts)
	persons.GET(":id", h.getPerson)
	persons.DELETE(":id", h.deletePerson)

	legalPersons := router.Group("/legal-persons")
	legalPersons.POST("", h.createLegalPerson)
	legalPersons.POST("/search", h.searchLegalPersonContracts)
	legalPersons.GET(":id", h.getLegalPerson)
	legalPersons.DELETE(":id", h.deleteLegalPerson)

	contracts := router.Group("/contracts")
	contracts.POST("", h.createContract)
	contracts.GET(":id", h.getContract)
	contracts.PUT(":id", h.updateContract)
	contracts.DELETE(":id", h.deleteContract)

	contracts.POST(":id/persons", h.addPersonToContract)
	contracts.DELETE(":id/persons/:personId", h.removePersonFromContract)
	contracts.POST(":id/legal-persons", h.addLegalPersonToContract)
	contracts.DELETE(":id/legal-persons/:legalPersonId", h.removeLegalPersonFromContract)
}

func (h *Handlers) createPerson(c *gin.Context) {
	var dto Person

	// Read raw body for debugging (log it) and then reset the body for binding
	if data, err := io.ReadAll(c.Request.Body); err == nil {
		log.Printf("createPerson raw body: %s", string(data))
		c.Request.Body = io.NopCloser(bytes.NewBuffer(data))
	}

	if err := c.ShouldBindJSON(&dto); err != nil {
		log.Printf("createPerson bind error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("createPerson after bind: %+v", dto)

	if err := validatePerson(dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	person := dtoToPerson(dto)

	if err := h.service.CreatePerson(context.Background(), &person); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("createPerson created: %+v", person)

	log.Printf("saved person with id: %d", person.ID)

	c.JSON(http.StatusCreated, personToDto(person))
}

func (h *Handlers) searchPersonContracts(c *gin.Context) {
	var filter SearchPersonContractsFilter
	if err := c.ShouldBindJSON(&filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	contracts, err := h.service.SearchPersonContracts(c.Request.Context(), models.SearchPersonContractsFilter{
		ContractType:      filter.ContractType,
		StartDateFrom:     filter.StartDateFrom,
		StartDateTo:       filter.StartDateTo,
		PersonName:        filter.PersonName,
		PersonLastname:    filter.PersonLastname,
		PersonPatronym:    filter.PersonPatronym,
		PersonDateOfBirth: filter.PersonDateOfBirth,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fmt.Printf("contracts len %d", len(contracts))

	dtos := make([]Contract, len(contracts))
	for i, c := range contracts {
		dtos[i] = contractToDto(c)
	}

	c.JSON(http.StatusOK, dtos)
}

func (h *Handlers) getPerson(c *gin.Context) {
	id, ok := parseID(c, "id")
	if !ok {
		return
	}

	person, err := h.service.GetPerson(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, personToDto(*person))
}

func (h *Handlers) deletePerson(c *gin.Context) {
	id, ok := parseID(c, "id")
	if !ok {
		return
	}

	if err := h.service.DeletePerson(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *Handlers) createLegalPerson(c *gin.Context) {
	var dto LegalPerson
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := validateLegalPerson(dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	legalPerson := dtoToLegalPerson(dto)

	if err := h.service.CreateLegalPerson(context.Background(), &legalPerson); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("saved legal person with id: %d", legalPerson.ID)

	c.JSON(http.StatusCreated, legalPersonToDto(legalPerson))
}

func (h *Handlers) searchLegalPersonContracts(c *gin.Context) {
	var filter SearchLegalPersonContractsFilter
	if err := c.ShouldBindJSON(&filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("FILTER: %+v", filter)

	contracts, err := h.service.SearchLegalPersonContracts(c.Request.Context(), models.SearchLegalPersonContractsFilter{
		ContractType:         filter.ContractType,
		StartDateFrom:        filter.StartDateFrom,
		StartDateTo:          filter.StartDateTo,
		LegalPersonShortName: filter.LegalPersonShortName,
		LegalPersonName:      filter.LegalPersonName,
		LocalEDRPOU:          filter.LocalEDRPOU,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fmt.Printf("contracts len %d", len(contracts))

	dtos := make([]Contract, len(contracts))
	for i, c := range contracts {
		dtos[i] = contractToDto(c)
	}

	c.JSON(http.StatusOK, dtos)
}

func (h *Handlers) getLegalPerson(c *gin.Context) {
	id, ok := parseID(c, "id")
	if !ok {
		return
	}

	legalPerson, err := h.service.GetLegalPerson(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, legalPersonToDto(*legalPerson))
}

func (h *Handlers) deleteLegalPerson(c *gin.Context) {
	id, ok := parseID(c, "id")
	if !ok {
		return
	}

	if err := h.service.DeleteLegalPerson(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *Handlers) createContract(c *gin.Context) {
	var dto Contract
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := validateContract(dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	contract := dtoToContract(dto)

	if err := h.service.CreateContract(context.Background(), &contract); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("saved contract with id: %d", contract.ID)

	c.JSON(http.StatusCreated, contractToDto(contract))
}

func (h *Handlers) getContract(c *gin.Context) {
	id, ok := parseID(c, "id")
	if !ok {
		return
	}

	details, err := h.service.GetContractDetails(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, contractDetailsToDto(details))
}

func (h *Handlers) updateContract(c *gin.Context) {
	id, ok := parseID(c, "id")
	if !ok {
		return
	}

	var dto Contract
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dto.ID = id
	contract := dtoToContract(dto)

	if err := h.service.UpdateContract(c.Request.Context(), &contract); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, contractToDto(contract))
}

func (h *Handlers) deleteContract(c *gin.Context) {
	id, ok := parseID(c, "id")
	if !ok {
		return
	}

	if err := h.service.DeleteContract(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *Handlers) addPersonToContract(c *gin.Context) {
	contractID, ok := parseID(c, "id")
	if !ok {
		return
	}

	var relation struct {
		PersonID uint `json:"person_id"`
	}
	if err := c.ShouldBindJSON(&relation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if relation.PersonID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "person_id is required"})
		return
	}

	if err := h.service.AddPersonToContract(c.Request.Context(), contractID, relation.PersonID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}

func (h *Handlers) removePersonFromContract(c *gin.Context) {
	contractID, ok := parseID(c, "id")
	if !ok {
		return
	}

	personID, ok := parseID(c, "personId")
	if !ok {
		return
	}

	if err := h.service.RemovePersonFromContract(c.Request.Context(), contractID, personID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *Handlers) addLegalPersonToContract(c *gin.Context) {
	contractID, ok := parseID(c, "id")
	if !ok {
		return
	}

	var relation struct {
		LegalPersonID uint `json:"legal_person_id"`
	}
	if err := c.ShouldBindJSON(&relation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if relation.LegalPersonID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "legal_person_id is required"})
		return
	}

	if err := h.service.AddLegalPersonToContract(c.Request.Context(), contractID, relation.LegalPersonID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}

func (h *Handlers) removeLegalPersonFromContract(c *gin.Context) {
	contractID, ok := parseID(c, "id")
	if !ok {
		return
	}

	legalPersonID, ok := parseID(c, "legalPersonId")
	if !ok {
		return
	}

	if err := h.service.RemoveLegalPersonFromContract(c.Request.Context(), contractID, legalPersonID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func parseID(c *gin.Context, key string) (uint, bool) {
	value, err := strconv.ParseUint(c.Param(key), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return 0, false
	}

	return uint(value), true
}
