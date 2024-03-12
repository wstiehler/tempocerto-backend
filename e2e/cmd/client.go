package cmd

import (
	"fmt"

	"github.com/dghubble/sling"
	"github.com/wstiehler/tempocerto-backend/internal/infrastructure/logger"
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

// func (project ProjectApi) CreateRole(r tempocerto.RoleEntity) (*tempocerto.RoleDTO, error) {
// 	logger, dispose := logger.New()
// 	defer dispose()

// 	response := new(tempocerto.RoleDTO)

// 	resp, err := sling.New().Base(project.url).Post("/v1/role").BodyJSON(r).ReceiveSuccess(response)
// 	if err != nil {
// 		logger.Error("Create error")
// 		fmt.Println(response, resp, err)
// 		return nil, err
// 	}
// 	return response, nil
// }

// func (project ProjectApi) GetRoleByID(id uint) (*tempocerto.RoleDTO, error) {
// 	logger, dispose := logger.New()
// 	defer dispose()

// 	response := new(tempocerto.RoleDTO)

// 	idString := strconv.Itoa(int(id))

// 	req, err := http.NewRequest("GET", project.url+"/v1/role/"+idString, nil)

// 	if err != nil {
// 		logger.Error("Failed to create request", zap.Error(err))
// 		return nil, err
// 	}

// 	httpClient := &http.Client{}
// 	resp, _ := httpClient.Do(req)

// 	err = newcheckHasError(resp)

// 	if err != nil {
// 		logger.Error("Error on get role by id", zap.String("error", err.Error()))
// 	}

// 	defer resp.Body.Close()

// 	responseBody, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		logger.Error("Failed to read response body", zap.Error(err))
// 		return nil, err
// 	}

// 	if err := json.Unmarshal(responseBody, response); err != nil {
// 		logger.Error("Failed to decode response body", zap.Error(err))
// 		return nil, err
// 	}

// 	return response, nil
// }

// func newcheckHasError(resp *http.Response) error {
// 	logger, dispose := logger.New()
// 	defer dispose()

// 	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
// 		logger.Error("Unexpected response status", zap.Int("status", resp.StatusCode))
// 		return fmt.Errorf("unexpected response status: %d", resp.StatusCode)
// 	}

// 	if resp.StatusCode >= 500 {
// 		logger.Error("Unexpected response status", zap.Int("status", resp.StatusCode))
// 		return fmt.Errorf("unexpected response status: %d", resp.StatusCode)
// 	}

// 	return nil
// }
