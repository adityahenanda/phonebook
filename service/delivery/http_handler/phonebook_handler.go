package http_handler

import (
	"encoding/json"
	"phonebook/models"
	"phonebook/service/usecase"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)

func (idb *InDB) GetPhonebooks(c *gin.Context) {

	var response models.ResponsePhoneBooks

	//prevent sql injection with queryEscape
	pageParam := url.QueryEscape(c.DefaultQuery("page", "1"))
	limitParam := url.QueryEscape(c.DefaultQuery("limit", "100"))

	//parse page and limit into int using uint to check value whether string or negative number
	page, err := strconv.ParseUint(pageParam, 10, 64)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Message = err.Error()
		response.Status = "Error"
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	limit, err := strconv.ParseUint(limitParam, 10, 64)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Message = err.Error()
		response.Status = "Error"
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	//begin bussiness logic
	res, total, err := usecase.GetPhonebooksUsecase(idb.DB, int(limit), int(page))
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Message = err.Error()
		response.Status = "Error"
		response.Data = res
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response.Code = http.StatusOK
	response.Message = "Success"
	response.Status = "Success"
	response.Data = res
	response.TotalData = total
	c.JSON(http.StatusOK, response)
	return
}

func (idb *InDB) GetPhonebook(c *gin.Context) {

	var response models.ResponsePhoneBook

	paramID := c.Param("phoneBookID")
	//parse params id into int using uint to check value whether string or negative number
	id, err := strconv.ParseUint(paramID, 10, 64)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Message = err.Error()
		response.Status = "Error"
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	//begin bussiness logic
	res, err := usecase.GetPhonebookUsecase(idb.DB, int(id))
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Message = err.Error()
		response.Status = "Error"
		response.Data = res
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response.Code = http.StatusOK
	response.Message = "Success"
	response.Status = "Success"
	response.Data = res
	c.JSON(http.StatusOK, response)
	return
}

func (idb *InDB) CreatePhonebook(c *gin.Context) {
	var response models.ResponsePhoneBookID
	var req models.PhonebookRequest
	body, _ := ioutil.ReadAll(c.Request.Body)

	//get body request
	err := json.Unmarshal(body, &req)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Message = err.Error()
		response.Status = "Error"
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	//validate tag required
	v := validator.New()
	err = v.Struct(req)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Message = err.Error()
		response.Status = "Error"
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	//begin bussiness logic
	id, err := usecase.CreatePhonebookUsecase(idb.DB, req)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Message = err.Error()
		response.Status = "Error"
		response.ID = id
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response.Code = http.StatusOK
	response.Message = "Success"
	response.Status = "Success"
	response.ID = id
	c.JSON(http.StatusOK, response)
	return
}

func (idb *InDB) DestroyPhonebook(c *gin.Context) {

	var response models.ResponsePhoneBookID

	paramID := c.Param("phoneBookID")
	//parse params id into int using uint to check value whether string or negative number
	id, err := strconv.ParseUint(paramID, 10, 64)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Message = err.Error()
		response.Status = "Error"
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	//begin bussiness logic
	err = usecase.DestroyPhonebookUsecase(idb.DB, int(id))
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Message = err.Error()
		response.Status = "Error"
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response.Code = http.StatusOK
	response.Message = "Success"
	response.Status = "Success"
	response.ID = int(id)
	c.JSON(http.StatusOK, response)
	return
}

func (idb *InDB) DeletePhonebook(c *gin.Context) {

	var response models.ResponsePhoneBookID

	paramID := c.Param("phoneBookID")
	//parse params id into int using uint to check value whether string or negative number
	id, err := strconv.ParseUint(paramID, 10, 64)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Message = err.Error()
		response.Status = "Error"
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	//begin bussiness logic
	err = usecase.DeletePhonebookUsecase(idb.DB, int(id))
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Message = err.Error()
		response.Status = "Error"
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response.Code = http.StatusOK
	response.Message = "Success"
	response.Status = "Success"
	response.ID = int(id)
	c.JSON(http.StatusOK, response)
	return
}

func (idb *InDB) UpdatePhonebook(c *gin.Context) {

	var response models.ResponsePhoneBookID
	var req models.PhonebookRequest
	body, _ := ioutil.ReadAll(c.Request.Body)
	//get body request
	err := json.Unmarshal(body, &req)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Message = err.Error()
		response.Status = "Error"
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	paramID := c.Param("phoneBookID")
	//parse params id into int using uint to check value whether string or negative number
	id, err := strconv.ParseUint(paramID, 10, 64)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Message = err.Error()
		response.Status = "Error"
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	//begin bussiness logic
	req.PhonebookID = int(id)
	err = usecase.UpdatePhonebookUsecase(idb.DB, req)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Message = err.Error()
		response.Status = "Error"
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response.Code = http.StatusOK
	response.Message = "Success"
	response.Status = "Success"
	response.ID = int(id)
	c.JSON(http.StatusOK, response)
	return
}
