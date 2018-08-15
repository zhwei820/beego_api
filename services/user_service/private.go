package user_service

func (this *UserController) loginTest() {
	username := this.GetString("username")
	this.GetLogger().Printf("this is a message with trace id")
	this.GetLogger().Printf("username: %s try to login.", username)

}