package actions

import (
	"github.com/gobuffalo/pop/nulls"
	"github.com/piffio/myvol/models"
)

func (as *ActionSuite) Test_Profile_ShowProfile() {
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
	res := as.HTML("/profile").Get()
	as.Equal(200, res.Code)
	as.Contains(res.Body.String(), "Sergio Visinoni")

	as.Session.Clear()

	res = as.HTML("/profile").Get()
	as.Equal(302, res.Code)
	as.Contains(res.Body.String(), "<a href=\"/\">Found</a>")
}
