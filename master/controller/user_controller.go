package controller

import (
	"context"
	service "eh_teh_mewa/master/service"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
)

type UserControllerImpl struct {
	UserService service.UsersService
}

func NewUsersController(usersService service.UsersService) UsersController {
	return &UserControllerImpl{UserService: usersService}
}

func (controller *UserControllerImpl) Create(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	myTemplate := template.Must(template.ParseFiles(
		"master/views/user/create.gohtml", "view/layout/app.gohtml",
		"view/layout/bodyTop.gohtml", "view/layout/footer.gohtml", "view/layout/head.gohtml", "view/layout/header.gohtml",
		"view/layout/sidebar.gohtml"))
	myTemplate.ExecuteTemplate(w, "createUser", map[string]interface{}{
		"type": [...]string{
			"admin", "kasir",
		},
	})
}

func (controller *UserControllerImpl) Store(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	//TODO implement me
	panic("implement me")
}

func (controller *UserControllerImpl) Update(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	//TODO implement me
	panic("implement me")
}

func (controller *UserControllerImpl) Delete(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	//TODO implement me
	panic("implement me")
}

func (controller *UserControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	myTemplate := template.Must(template.ParseFiles("master/views/user/index.gohtml", "view/layout/app.gohtml",
		"view/layout/bodyTop.gohtml", "view/layout/footer.gohtml", "view/layout/head.gohtml", "view/layout/header.gohtml",
		"view/layout/sidebar.gohtml"))

	serv := controller.UserService.FindAll(context.Background())
	myTemplate.ExecuteTemplate(w, "indexUser", serv)
	//var filepath = path.Join("view", "users_show.html")
	//tmpl, err := template.ParseFiles(filepath)
	//if err != nil {
	//	panic(err)
	//}
	//db := model.GetConnect()
	//UserRepo := repository.NewUsersRepository()
	//userService := service.NewUsersService(UserRepo, db)
	//UserResponses := userService.FindAll(context.Background())
	//datas := helperMain.StructSliceToMap_Users(UserResponses)
	//err1 := tmpl.Execute(w, datas)
	//if err1 != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//}
	//fmt.Println(UserResponses)
}

func (controller *UserControllerImpl) FindById(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	//TODO implement me
	panic("implement me")
}
