package models_test

import (
	"github.com/gobuffalo/pop/nulls"
	"github.com/piffio/myvol/models"
)

func (as *ModelSuite) Test_User_Create() {
	count, err := as.DB.Count("users")
	as.NoError(err)
	as.Equal(0, count)

	u := &models.User{
		Email:      nulls.NewString("sergio@example.com"),
		Name:       "Sergio Visinoni",
		Provider:   "gplus",
		ProviderID: "123456abc",
	}

	err = as.DB.Save(u)
	as.NoError(err)

	count, err = as.DB.Count("users")
	as.NoError(err)
	as.Equal(1, count)
}
