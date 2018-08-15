package default_service

func (o *DefaultController) GetAllPublic() {
	o.GetLogger().Printf("log from other service!!!!!")
}
