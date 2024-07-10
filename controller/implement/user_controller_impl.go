package implement

import (
	"azizdev/helper"
	"azizdev/model/web"
	"azizdev/model/web/create"
	"azizdev/model/web/update"
	"azizdev/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func (controller *UserControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userCreateRequest := create.UserCreateRequest{}
	helper.ReadFromRequestBody(request, &userCreateRequest)

	userResponse := controller.UserService.Create(request.Context(), userCreateRequest)
	webResponse := web.WebResponse{
		Code:    200,
		Status:  "OK",
		Massage: "Successfully",
		Data:    userResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userUpdteRequest := update.UserUpdateRequest{}
	helper.ReadFromRequestBody(request, &userUpdteRequest)

	userId := params.ByName("id")
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	userUpdteRequest.Id = id
	userResponse := controller.UserService.Update(request.Context(), userUpdteRequest)
	webResponse := web.WebResponse{
		Code:    200,
		Status:  "OK",
		Massage: "Successfully",
		Data:    userResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := params.ByName("id")
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	controller.UserService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:    200,
		Status:  "OK",
		Massage: "Successfully",
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := params.ByName("id")
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)
	userResponse := controller.UserService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:    200,
		Status:  "OK",
		Massage: "Successfully",
		Data:    userResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userResponse := controller.UserService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:    200,
		Status:  "OK",
		Massage: "Successfully",
		Data:    userResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}
