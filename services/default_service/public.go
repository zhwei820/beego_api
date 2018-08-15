package default_service

func (this *DefaultController) GetAllPublic() {
	this.GetLogger().Printf("log from other service!!!!!")
}
