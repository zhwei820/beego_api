package user_service

func (this *UserController) LoginTest() {
	username := this.GetString("username")
	this.GetLogger().Printf("this is a message with trace id")
	this.GetLogger().Printf("username: %s try to login.", username)

}