package base_service

type ErrorController struct {
	BaseController
}

//
//func (c *ErrorController) Error404() {
//
//	c.WriteJsonWithStatusCode(404, Response{
//		Errcode: 404,
//		Errmsg:  "Not Found",
//	})
//
//}
//func (c *ErrorController) Error401() {
//	c.WriteJsonWithStatusCode(401, Response{
//		Errcode: 401,
//		Errmsg:  "Permission denied",
//	})
//
//}
//func (c *ErrorController) Error403() {
//
//	c.WriteJsonWithStatusCode(403, Response{
//		Errcode: 403,
//		Errmsg:  "Forbidden",
//	})
//
//}
