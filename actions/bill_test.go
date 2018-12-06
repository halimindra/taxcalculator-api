package actions

func (as *ActionSuite) Test_Bill_Show() {
	res := as.JSON("/bill").Get()
	as.Equal(200, res.Code)
}
