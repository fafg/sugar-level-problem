package reports

import (
	"sugar-level-client/client"
	"sugar-level-client/decisiontree"
	"sugar-level-client/models"
	"runtime"
	"sync"
)

var (
	_backend *client.BackendClient
)

func NewFromBackendClient(backendClient *client.BackendClient) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	_backend = backendClient
	decisiontree.Init()
}

func ProduceSugarReportAsync() *models.SugarReport {
	var lowResult models.SugarResult
	var normalResult models.SugarResult
	var highResult models.SugarResult

	var wg sync.WaitGroup

	if _backend != nil {
		users, err := _backend.GetUserData()
		if err == nil {
			wg.Add(3)

			go func() {
				lowResult = classifyUsersFiltered(users, decisiontree.Low)
				wg.Done()
			}()
			go func() {
				normalResult = classifyUsersFiltered(users, decisiontree.Normal)
				wg.Done()
			}()
			go func() {
				highResult = classifyUsersFiltered(users, decisiontree.High)
				wg.Done()
			}()

			wg.Wait()
		}
	}

	return &models.SugarReport{
		Low:    lowResult,
		Normal: normalResult,
		High:   highResult,
	}
}

func ProduceSugarReport() (*models.SugarReport, error) {
	var result = &models.SugarReport{
		Low:    models.SugarResult{},
		Normal: models.SugarResult{},
		High:   models.SugarResult{},
	}

	if _backend != nil {
		users, err := _backend.GetUserData()
		if err == nil {
			for _, user := range *users {
				if decisiontree.CheckUserLevel(&user) == decisiontree.Low {
					result.Low.Count = result.Low.Count + 1
					result.Low.Users = append(result.Low.Users, user.ID)
				} else if decisiontree.CheckUserLevel(&user) == decisiontree.Normal {
					result.Normal.Count = result.Normal.Count + 1
					result.Normal.Users = append(result.Normal.Users, user.ID)
				} else if decisiontree.CheckUserLevel(&user) == decisiontree.High {
					result.High.Count = result.High.Count + 1
					result.High.Users = append(result.High.Users, user.ID)
				}
			}
		} else {
			return nil, err
		}
	}

	return result, nil
}

func classifyUsersFiltered(users *[]models.User, classificationFilter decisiontree.SugarClassification) models.SugarResult {
	var result decisiontree.SugarClassification
	var SugarResult = models.SugarResult{}

	for _, user := range *users {
		result = decisiontree.CheckUserLevel(&user)
		if result == classificationFilter {
			SugarResult.Count = SugarResult.Count + 1
			SugarResult.Users = append(SugarResult.Users, user.ID)
		}
		result = decisiontree.Unknown
		runtime.Gosched()
	}

	return SugarResult
}