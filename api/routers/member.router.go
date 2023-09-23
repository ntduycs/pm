package routers

import "project-management/controllers"

type MemberRouter struct {
	controller *controllers.MemberController
}

func NewMemberRouter(controller *controllers.MemberController) *MemberRouter {
	return &MemberRouter{
		controller: controller,
	}
}

func (r *MemberRouter) Register(props *RouterProps) {
	router := props.App.Group(props.Prefix)

	router.Get("/:id", r.controller.GetMember)
}

func (r *MemberRouter) GetPrefix() string {
	return "/members"
}
