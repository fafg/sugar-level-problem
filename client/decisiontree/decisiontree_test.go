package decisiontree

import (
	"sugar-level-client/models"
	"testing"
)

func test(t *testing.T, user *models.User, expectedClassification SugarClassification) {
	Init()
	classification := CheckUserLevel(user)
	if classification != expectedClassification {
		t.Errorf("classification should be %v, but it is: %v", expectedClassification, classification)
	}
}

func Test_low_low_high_Values_Expected_Result_As_Low(t *testing.T) {
	var user models.User
	user.ID = "3f45f836-7ead-4f75-9067-f83f0f6a0d7e"
	user.Samples = []models.Sample {
		{ Time: "morning", Value: 0.49 }, //low morning -> 0.50 - 0.78
		{ Time: "afternoon", Value: 0.01 }, //low afternoon -> 0.15 - 0.24
		{ Time: "evening", Value: 0.37 }, //high evening -> 0.12 - 0.21
	}

	test(t, &user, Low)
}

func Test_low_normal_low_Values_Expected_Result_As_Low(t *testing.T) {
	var user models.User
	user.ID = "3f45f836-7ead-4f75-9067-f83f0f6a0d7e"
	user.Samples = []models.Sample {
		{ Time: "morning", Value: 0.49 }, //low morning -> 0.50 - 0.78
		{ Time: "afternoon", Value: 0.15 }, //normal afternoon -> 0.15 - 0.24
		{ Time: "evening", Value: 0.11 }, //low evening -> 0.12 - 0.21
	}

	test(t, &user, Low)
}

func Test_low_normal_normal_Values_Expected_Result_As_Normal(t *testing.T) {
	var user models.User
	user.ID = "3f45f836-7ead-4f75-9067-f83f0f6a0d7e"
	user.Samples = []models.Sample {
		{ Time: "morning", Value: 0.49 }, //low morning -> 0.50 - 0.78
		{ Time: "afternoon", Value: 0.15 }, //normal afternoon -> 0.15 - 0.24
		{ Time: "evening", Value: 0.14 }, //normal evening -> 0.12 - 0.21
	}

	test(t, &user, Normal)
}

func Test_low_normal_high_Values_Expected_Result_As_Normal(t *testing.T) {
	var user models.User
	user.ID = "3f45f836-7ead-4f75-9067-f83f0f6a0d7e"
	user.Samples = []models.Sample {
		{ Time: "morning", Value: 0.49 }, //low morning -> 0.50 - 0.78
		{ Time: "afternoon", Value: 0.15 }, //normal afternoon -> 0.15 - 0.24
		{ Time: "evening", Value: 0.22 }, //high evening -> 0.12 - 0.21
	}

	test(t, &user, Normal)
}

func Test_low_high_low_Values_Expected_Result_As_Low(t *testing.T) {
	var user models.User
	user.ID = "3f45f836-7ead-4f75-9067-f83f0f6a0d7e"
	user.Samples = []models.Sample {
		{ Time: "morning", Value: 0.49 }, //low morning -> 0.50 - 0.78
		{ Time: "afternoon", Value: 0.25 }, //high afternoon -> 0.15 - 0.24
		{ Time: "evening", Value: 0.11 }, //low evening -> 0.12 - 0.21
	}

	test(t, &user, Low)
}

func Test_low_high_normal_Values_Expected_Result_As_Normal(t *testing.T) {
	var user models.User
	user.ID = "3f45f836-7ead-4f75-9067-f83f0f6a0d7e"
	user.Samples = []models.Sample {
		{ Time: "morning", Value: 0.49 }, //low morning -> 0.50 - 0.78
		{ Time: "afternoon", Value: 0.25 }, //high afternoon -> 0.15 - 0.24
		{ Time: "evening", Value: 0.12 }, //normal evening -> 0.12 - 0.21
	}

	test(t, &user, Normal)
}

func Test_low_high_high_Values_Expected_Result_As_High(t *testing.T) {
	var user models.User
	user.ID = "3f45f836-7ead-4f75-9067-f83f0f6a0d7e"
	user.Samples = []models.Sample {
		{ Time: "morning", Value: 0.49 }, //low morning -> 0.50 - 0.78
		{ Time: "afternoon", Value: 0.25 }, //high afternoon -> 0.15 - 0.24
		{ Time: "evening", Value: 0.22 }, //high evening -> 0.12 - 0.21
	}

	test(t, &user, High)
}

func Test_normal_low_low_Values_Expected_Result_As_Low(t *testing.T) {
	var user models.User
	user.ID = "3f45f836-7ead-4f75-9067-f83f0f6a0d7e"
	user.Samples = []models.Sample {
		{ Time: "morning", Value: 0.50 }, //normal morning -> 0.50 - 0.78
		{ Time: "afternoon", Value: 0.14 }, //low afternoon -> 0.15 - 0.24
		{ Time: "evening", Value: 0.11 }, //low evening -> 0.12 - 0.21
	}

	test(t, &user, Low)
}

func Test_normal_low_normal_Values_Expected_Result_As_Normal(t *testing.T) {
	var user models.User
	user.ID = "3f45f836-7ead-4f75-9067-f83f0f6a0d7e"
	user.Samples = []models.Sample {
		{ Time: "morning", Value: 0.50 }, //normal morning -> 0.50 - 0.78
		{ Time: "afternoon", Value: 0.14 }, //low afternoon -> 0.15 - 0.24
		{ Time: "evening", Value: 0.14 }, //normal evening -> 0.12 - 0.21
	}

	test(t, &user, Normal)
}

func Test_normal_low_high_Values_Expected_Result_As_Normal(t *testing.T) {
	var user models.User
	user.ID = "3f45f836-7ead-4f75-9067-f83f0f6a0d7e"
	user.Samples = []models.Sample {
		{ Time: "morning", Value: 0.50 }, //normal morning -> 0.50 - 0.78
		{ Time: "afternoon", Value: 0.14 }, //low afternoon -> 0.15 - 0.24
		{ Time: "evening", Value: 0.22 }, //high evening -> 0.12 - 0.21
	}

	test(t, &user, Normal)
}

func Test_normal_high_high_Values_Expected_Result_As_High(t *testing.T) {
	var user models.User
	user.ID = "3f45f836-7ead-4f75-9067-f83f0f6a0d7e"
	user.Samples = []models.Sample {
		{ Time: "morning", Value: 0.50 }, //normal morning -> 0.50 - 0.78
		{ Time: "afternoon", Value: 0.25 }, //high afternoon -> 0.15 - 0.24
		{ Time: "evening", Value: 0.22 }, //high evening -> 0.12 - 0.21
	}

	test(t, &user, High)
}

func Test_normal_high_normal_Values_Expected_Result_As_Normal(t *testing.T) {
	var user models.User
	user.ID = "3f45f836-7ead-4f75-9067-f83f0f6a0d7e"
	user.Samples = []models.Sample {
		{ Time: "morning", Value: 0.50 }, //normal morning -> 0.50 - 0.78
		{ Time: "afternoon", Value: 0.25 }, //high afternoon -> 0.15 - 0.24
		{ Time: "evening", Value: 0.21 }, //normal evening -> 0.12 - 0.21
	}

	test(t, &user, Normal)
}

func Test_normal_high_low_Values_Expected_Result_As_Normal(t *testing.T) {
	var user models.User
	user.ID = "3f45f836-7ead-4f75-9067-f83f0f6a0d7e"
	user.Samples = []models.Sample {
		{ Time: "morning", Value: 0.50 }, //normal morning -> 0.50 - 0.78
		{ Time: "afternoon", Value: 0.25 }, //high afternoon -> 0.15 - 0.24
		{ Time: "evening", Value: 0.11 }, //low evening -> 0.12 - 0.21
	}

	test(t, &user, Normal)
}

func Test_normal_normal_high_Values_Expected_Result_As_Normal(t *testing.T) {
	var user models.User
	user.ID = "3f45f836-7ead-4f75-9067-f83f0f6a0d7e"
	user.Samples = []models.Sample {
		{ Time: "morning", Value: 0.50 }, //normal morning -> 0.50 - 0.78
		{ Time: "afternoon", Value: 0.15 }, //normal afternoon -> 0.15 - 0.24
		{ Time: "evening", Value: 0.22 }, //high evening -> 0.12 - 0.21
	}

	test(t, &user, Normal)
}

func Test_high_high_low_Values_Expected_Result_As_High(t *testing.T) {
	var user models.User
	user.ID = "3f45f836-7ead-4f75-9067-f83f0f6a0d7e"
	user.Samples = []models.Sample {
		{ Time: "morning", Value: 0.79 }, //high morning -> 0.50 - 0.78
		{ Time: "afternoon", Value: 0.25 }, //high afternoon -> 0.15 - 0.24
		{ Time: "evening", Value: 0.11 }, //low evening -> 0.12 - 0.21
	}

	test(t, &user, High)
}

func Test_high_low_low_Values_Expected_Result_As_Low(t *testing.T) {
	var user models.User
	user.ID = "3f45f836-7ead-4f75-9067-f83f0f6a0d7e"
	user.Samples = []models.Sample {
		{ Time: "morning", Value: 0.79 }, //high morning -> 0.50 - 0.78
		{ Time: "afternoon", Value: 0.14 }, //low afternoon -> 0.15 - 0.24
		{ Time: "evening", Value: 0.11 }, //low evening -> 0.12 - 0.21
	}

	test(t, &user, Low)
}

func Test_high_low_normal_Values_Expected_Result_As_Normal(t *testing.T) {
	var user models.User
	user.ID = "3f45f836-7ead-4f75-9067-f83f0f6a0d7e"
	user.Samples = []models.Sample {
		{ Time: "morning", Value: 0.79 }, //high morning -> 0.50 - 0.78
		{ Time: "afternoon", Value: 0.14 }, //low afternoon -> 0.15 - 0.24
		{ Time: "evening", Value: 0.12 }, //normal evening -> 0.12 - 0.21
	}

	test(t, &user, Normal)
}

func Test_high_low_high_Values_Expected_Result_As_High(t *testing.T) {
	var user models.User
	user.ID = "3f45f836-7ead-4f75-9067-f83f0f6a0d7e"
	user.Samples = []models.Sample {
		{ Time: "morning", Value: 0.79 }, //high morning -> 0.50 - 0.78
		{ Time: "afternoon", Value: 0.14 }, //low afternoon -> 0.15 - 0.24
		{ Time: "evening", Value: 0.22 }, //high evening -> 0.12 - 0.21
	}

	test(t, &user, High)
}

func Test_high_normal_high_Values_Expected_Result_As_High(t *testing.T) {
	var user models.User
	user.ID = "3f45f836-7ead-4f75-9067-f83f0f6a0d7e"
	user.Samples = []models.Sample {
		{ Time: "morning", Value: 0.79 }, //high morning -> 0.50 - 0.78
		{ Time: "afternoon", Value: 0.16 }, //normal afternoon -> 0.15 - 0.24
		{ Time: "evening", Value: 0.22 }, //high evening -> 0.12 - 0.21
	}

	test(t, &user, High)
}

func Test_high_normal_normal_Values_Expected_Result_As_Normal(t *testing.T) {
	var user models.User
	user.ID = "3f45f836-7ead-4f75-9067-f83f0f6a0d7e"
	user.Samples = []models.Sample {
		{ Time: "morning", Value: 0.79 }, //high morning -> 0.50 - 0.78
		{ Time: "afternoon", Value: 0.16 }, //normal afternoon -> 0.15 - 0.24
		{ Time: "evening", Value: 0.20 }, //normal evening -> 0.12 - 0.21
	}

	test(t, &user, Normal)
}

func Test_high_normal_low_Values_Expected_Result_As_Normal(t *testing.T) {
	var user models.User
	user.ID = "3f45f836-7ead-4f75-9067-f83f0f6a0d7e"
	user.Samples = []models.Sample {
		{ Time: "morning", Value: 0.79 }, //high morning -> 0.50 - 0.78
		{ Time: "afternoon", Value: 0.16 }, //normal afternoon -> 0.15 - 0.24
		{ Time: "evening", Value: 0.11 }, //low evening -> 0.12 - 0.21
	}

	test(t, &user, Normal)
}