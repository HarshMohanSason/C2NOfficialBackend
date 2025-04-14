package tests

import(
	"testing"
	"c2nofficialsitebackend/utils"
	)

func TestPassword(t *testing.T){

    passwords := []string{
        "HarshMohanSason",
        "",
    }

	for _, tc := range passwords {
		t.Run("CheckGeneratePassword", func(t *testing.T){
			generatedPassword, err := utils.GenerateHashedPassword(tc)
			err = utils.VerifyPasswords(generatedPassword, tc)
			if err != nil{
				t.Errorf("Error occured: %v", err)
			}
			t.Logf("Result: %+v", generatedPassword)
		})
	}
}
