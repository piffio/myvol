package models_test

import (
	"github.com/gobuffalo/pop/nulls"
	"github.com/piffio/myvol/models"
)

func (as *ModelSuite) Test_Device_Create() {
	u := &models.User{
		Email:      nulls.NewString("sergio@example.com"),
		Name:       "Sergio Visinoni",
		Provider:   "gplus",
		ProviderID: "123456abc",
	}

	err := as.DB.Create(u)
	as.NoError(err)

	count, err := as.DB.Count("devices")
	as.NoError(err)
	as.Equal(0, count)

	d := &models.Device{
		Model:       "Raspi",
		Description: "The nice raspi in my living room",
		Serial:      "000000007967f6d4",
		UserID:      u.ID,
	}

	err = as.DB.Create(d)
	as.NoError(err)

	count, err = as.DB.Count("devices")
	as.NoError(err)
	as.Equal(1, count)
}
