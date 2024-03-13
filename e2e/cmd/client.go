package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/dghubble/sling"
	"github.com/wstiehler/tempocerto-backend/internal/domain/tempocerto"
	"github.com/wstiehler/tempocerto-backend/internal/infrastructure/logger"
	"go.uber.org/zap"
)

type ErrorInfo struct {
	Message string `json:"msg"`
}

type ErrorResponse struct {
	Errors []ErrorInfo `json:"errors"`
}

type HealthReturn struct {
	Status int `json:"status"`
}

type ProjectApi struct {
	url string
}

func NewProjectApi(url string) *ProjectApi {

	project := &ProjectApi{
		url: url,
	}

	return project
}

func (project ProjectApi) ApiHealth() (*HealthReturn, error) {
	logger, dispose := logger.New()
	defer dispose()

	response := new(HealthReturn)

	resp, err := sling.New().Base(project.url).Path("health").ReceiveSuccess(response)

	if err != nil {
		logger.Error("Error")
		logger.Error(err.Error())
		return nil, err
	}
	fmt.Printf("[Health] result: %v\n", resp)
	return response, nil
}

func (project ProjectApi) CreateCompany(r tempocerto.CompanyEntity) (*tempocerto.CompanyDTO, error) {
	logger, dispose := logger.New()
	defer dispose()

	response := new(tempocerto.CompanyDTO)

	resp, err := sling.New().Base(project.url).Post("/v1/company").BodyJSON(r).ReceiveSuccess(response)
	if err != nil {
		logger.Error("Create error")
		fmt.Println(response, resp, err)
		return nil, err
	}
	return response, nil
}

func (project ProjectApi) CreateWeeklySlots(r tempocerto.WeeklySlotEntity) (*tempocerto.WeeklySlotEntity, error) {
	logger, dispose := logger.New()
	defer dispose()

	// Codifica o corpo da solicitação como JSON
	requestBody, err := json.Marshal(r)
	if err != nil {
		logger.Error("Error encoding request body:", zap.String("error", err.Error()))
		return nil, err
	}

	// Cria a solicitação HTTP POST
	req, err := http.NewRequest("POST", project.url+"/v1/weekly-slots", bytes.NewBuffer(requestBody))
	if err != nil {
		logger.Error("Error creating POST request:", zap.String("error", err.Error()))
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json") // Define o cabeçalho Content-Type como application/json

	// Cria um cliente HTTP
	client := &http.Client{}

	// Envia a solicitação e obtém a resposta
	resp, err := client.Do(req)
	if err != nil {
		logger.Error("Error sending POST request:", zap.String("error", err.Error()))
		return nil, err
	}
	defer resp.Body.Close()

	// Lê o corpo da resposta
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error("Error reading response body:", zap.String("error", err.Error()))
		return nil, err
	}

	// Verifica se a resposta foi bem-sucedida (código de status 2xx)
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		logger.Error("Received non-successful response:", zap.String("error", resp.Status))
		return nil, fmt.Errorf("non-successful response: %d", resp.StatusCode)
	}

	// Decodifica o corpo da resposta JSON para uma instância de WeeklySlotEntity
	var response tempocerto.WeeklySlotEntity
	if err := json.Unmarshal(responseBody, &response); err != nil {
		logger.Error("Error decoding response body:", zap.String("error", err.Error()))
		return nil, err
	}

	// Retorna a entidade semanal criada
	return &response, nil
}
