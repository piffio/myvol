package actions

import (
	"github.com/gobuffalo/pop/nulls"
	"github.com/piffio/myvol/models"
)

func createTestUser(as *ActionSuite) *models.User {
	u := &models.User{
		Email:      nulls.NewString("sergio@example.com"),
		Name:       "Sergio Visinoni",
		Provider:   "gplus",
		ProviderID: "123456abc",
	}

	err := as.DB.Create(u)
	as.NoError(err)

	user := &models.User{}

	err = as.DB.First(user)
	as.NoError(err)

	as.Session.Set("current_user_id", user.ID)

	return u
}

func createTestDevice(as *ActionSuite, u *models.User) *models.Device {
	d := &models.Device{
		Model:       "Raspi",
		Description: "Testing my rasputin",
		Serial:      "00000000deadbeef",
		UserID:      u.ID,
	}

	err := as.DB.Create(d)
	as.NoError(err)

	return d
}

func (as *ActionSuite) Test_DevicesResource_List() {
	u := createTestUser(as)
	_ = createTestDevice(as, u)

	res := as.HTML("/devices").Get()
	as.Equal(200, res.Code)
	as.Contains(res.Body.String(), "<h1>Devices</h1>")
	as.Contains(res.Body.String(), "Raspi - Testing my rasputin")
}

func (as *ActionSuite) Test_DevicesResource_Show() {
	u := createTestUser(as)
	d := createTestDevice(as, u)

	res := as.HTML("/devices/" + d.ID.String()).Get()
	as.Equal(200, res.Code)

	as.Contains(res.Body.String(), "Model: Raspi")
	as.Contains(res.Body.String(), "Description: Testing my rasputin")
}

func (as *ActionSuite) Test_DevicesResource_New() {
	_ = createTestUser(as)

	res := as.HTML("/devices/new").Get()
	as.Equal(200, res.Code)

	as.Contains(res.Body.String(), "<label>Model</label>")
	as.Contains(res.Body.String(), "<label>Description</label>")
	as.Contains(res.Body.String(), "<label>Serial</label>")
}

func (as *ActionSuite) Test_DevicesResource_Create() {
	u := createTestUser(as)

	d := &models.Device{
		Model:       "Raspi",
		Description: "Testing my rasputin",
		Serial:      "00000000deadbeef",
		UserID:      u.ID,
	}

	res := as.HTML("/devices").Post(d)

	err := as.DB.First(d)
	as.NoError(err)
	as.Equal(302, res.Code)
	as.Equal("/devices/"+d.ID.String(), res.Location())
}

func (as *ActionSuite) Test_DevicesResource_Edit() {
	u := createTestUser(as)
	d := createTestDevice(as, u)

	res := as.HTML("/devices/" + d.ID.String() + "/edit").Get()
	as.Equal(200, res.Code)
	as.Contains(res.Body.String(), "<h1>Edit Device</h1>")
	as.Contains(res.Body.String(), "Testing my rasputin")
}

func (as *ActionSuite) Test_DevicesResource_Update() {
	u := createTestUser(as)
	d := createTestDevice(as, u)

	updated := &models.Device{
		Model:       "Raspi",
		Description: "Updating my rasputin",
		Serial:      "00000000deadbeef",
	}

	res := as.HTML("/devices/" + d.ID.String()).Put(updated)
	as.Equal(200, res.Code)
	as.Contains(res.Body.String(), "Updating my rasputin")
}

func (as *ActionSuite) Test_DevicesResource_Destroy() {
	u := createTestUser(as)
	d := createTestDevice(as, u)

	res := as.HTML("/devices/" + d.ID.String()).Delete()

	as.Equal(302, res.Code)
	as.Equal("/devices", res.Location())
	as.NotContains(res.Body.String(), "Testing my rasputin")
}
