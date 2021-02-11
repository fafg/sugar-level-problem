package reports

import (
	"sugar-level-client/client"
	"testing"
)

func Test_Producer_Report(t *testing.T) {
	var backend, err = client.NewFromUrlAddress("http://localhost:8080/results")

	if err == nil {
		NewFromBackendClient(backend)
		var report, _ = ProduceSugarReport()
		if report.Low.Count == 0 || report.Low.Count != len(report.Low.Users) {
			t.Errorf("the amount of lower users are different of the counter")
		}
		if report.Normal.Count == 0 || report.Normal.Count != len(report.Normal.Users) {
			t.Errorf("the amount of lower users are different of the counter")
		}
		if report.High.Count == 0 || report.High.Count != len(report.High.Users) {
			t.Errorf("the amount of lower users are different of the counter")
		}

		if (report.Low.Count + report.Normal.Count + report.High.Count) != 300 {
			t.Errorf("users missed to be classified")
		}
	}
}

func Test_ProduceReport_Async(t *testing.T) {
	var backend, err = client.NewFromUrlAddress("http://localhost:8080/results")

	if err == nil {
		NewFromBackendClient(backend)
		var report = ProduceSugarReportAsync()
		if report.Low.Count == 0 || report.Low.Count != len(report.Low.Users) {
			t.Errorf("the amount of lower users are different of the counter")
		}
		if report.Normal.Count == 0 || report.Normal.Count != len(report.Normal.Users) {
			t.Errorf("the amount of lower users are different of the counter")
		}
		if report.High.Count == 0 || report.High.Count != len(report.High.Users) {
			t.Errorf("the amount of lower users are different of the counter")
		}

		if (report.Low.Count + report.Normal.Count + report.High.Count) != 300 {
			t.Errorf("users missed to be classified")
		}
	}
}