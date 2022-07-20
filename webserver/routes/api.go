package routes

import (
	"github.com/tzdit/sample_api/webserver/api"

	"github.com/labstack/echo/v4"
)

//APIRouters Init Router
func APIRouters(app *echo.Echo) {

	//Protected api should be defined in this group
	//This api is only accessed by authenticated user
	//api := app.Group("/api/v1", middlewares.Jwt()) //remove the middleware if you want to make public /github.com/tzdit/sample_api/api/v1
	aim := app.Group("/api/v1/grm") //remove the middleware if you want to make public

	departments := aim.Group("/departments")
	{
		departments.POST("/create", api.CreateDepartment)
		departments.GET("/list", api.ListDepartments)
		departments.GET("/get/:id", api.GetDepartment)
		departments.POST("/update", api.UpdateDepartment)
		departments.POST("/delete", api.DeleteDepartment)
	}

	grievant_category := aim.Group("/grievant_categories")
	{
		grievant_category.POST("/store", api.StoreGrievantCategory)
		grievant_category.GET("/list", api.ListGrievantCategories)
		grievant_category.GET("/show/:id", api.ShowGrievantCategory)
		grievant_category.POST("/update", api.UpdateGrievantCategory)
		grievant_category.POST("/delete", api.DeleteGrievantCategory)
	}

	grievant_group := aim.Group("/grievant_groups")
	{
		grievant_group.POST("/store", api.StoreGrievantGroup)
		grievant_group.GET("/list", api.ListGrievantGroups)
		grievant_group.GET("/show/:id", api.ShowGrievantGroup)
		grievant_group.POST("/update", api.UpdateGrievantGroup)
		grievant_group.POST("/delete", api.DeleteGrievantGroup)
	}

	grievance_category := aim.Group("/grievance_categories")
	{
		grievance_category.POST("/store", api.StoreGrievanceCategory)
		grievance_category.GET("/list", api.ListGrievanceCategories)
		grievance_category.GET("/show/:id", api.ShowGrievanceCategory)
		grievance_category.POST("/update", api.UpdateGrievanceCategory)
		grievance_category.POST("/delete", api.DeleteGrievanceCategory)
	}

	grievance_sub_category := aim.Group("/grievance_sub_categories")
	{
		grievance_sub_category.POST("/store", api.StoreGrievanceSubCategory)
		grievance_sub_category.GET("/list", api.ListGrievanceSubCategories)
		grievance_sub_category.GET("/show/:id", api.ShowGrievanceSubCategory)
		grievance_sub_category.POST("/update", api.UpdateGrievanceSubCategory)
		grievance_sub_category.POST("/delete", api.DeleteGrievanceSubCategory)
	}

	grievance_filling_mode := aim.Group("/grievance_filling_modes")
	{
		grievance_filling_mode.POST("/store", api.StoreGrievanceFillingMode)
		grievance_filling_mode.GET("/list", api.ListGrievanceFillingModes)
		grievance_filling_mode.GET("/show/:id", api.ShowGrievanceFillingMode)
		grievance_filling_mode.POST("/update", api.UpdateGrievanceFillingMode)
		grievance_filling_mode.POST("/delete", api.DeleteGrievanceFillingMode)
	}

	grievance_state := aim.Group("/grievance_states")
	{
		grievance_state.POST("/store", api.StoreGrievanceState)
		grievance_state.GET("/list", api.ListGrievanceStates)
		grievance_state.GET("/show/:id", api.ShowGrievanceState)
		grievance_state.POST("/update", api.UpdateGrievanceState)
		grievance_state.POST("/delete", api.DeleteGrievanceState)
	}

	grievance_state_action := aim.Group("/grievance_state_actions")
	{
		grievance_state_action.POST("/store", api.StoreGrievanceStateAction)
		grievance_state_action.GET("/list", api.ListGrievanceStateActions)
		grievance_state_action.GET("/show/:id", api.ShowGrievanceStateAction)
		grievance_state_action.POST("/update", api.UpdateGrievanceStateAction)
		grievance_state_action.POST("/delete", api.DeleteGrievanceStateAction)
	}

	grievance_state_transition := aim.Group("/grievance_state_transitions")
	{
		grievance_state_transition.POST("/store", api.StoreGrievanceStateTransition)
		grievance_state_transition.GET("/list", api.ListGrievanceStateTransitions)
		grievance_state_transition.GET("/show/:id", api.ShowGrievanceStateTransition)
		grievance_state_transition.POST("/update", api.UpdateGrievanceStateTransition)
		grievance_state_transition.POST("/delete", api.DeleteGrievanceStateTransition)
	}

	grievance_appeal_reason := aim.Group("/grievance_appeal_reasons")
	{
		grievance_appeal_reason.POST("/store", api.StoreGrievanceAppealReason)
		grievance_appeal_reason.GET("/list", api.ListGrievanceAppealReasons)
		grievance_appeal_reason.GET("/show/:id", api.ShowGrievanceAppealReason)
		grievance_appeal_reason.POST("/update", api.UpdateGrievanceAppealReason)
		grievance_appeal_reason.POST("/delete", api.DeleteGrievanceAppealReason)
	}

	grievance_resolution_state := aim.Group("/grievance_resolution_state")
	{
		grievance_resolution_state.POST("/store", api.StoreGrievanceResolutionState)
		grievance_resolution_state.GET("/list", api.ListGrievanceResolutionStates)
		grievance_resolution_state.GET("/show/:id", api.ShowGrievanceResolutionState)
		grievance_resolution_state.POST("/update", api.UpdateGrievanceResolutionState)
		grievance_resolution_state.POST("/delete", api.DeleteGrievanceResolutionState)
	}

	grievance_resolution := aim.Group("/grievance_resolution")
	{
		grievance_resolution.POST("/store", api.StoreGrievanceResolution)
		grievance_resolution.GET("/list", api.ListGrievanceResolutions)
		grievance_resolution.GET("/list_new", api.ListNewGrievanceResolutions)
		grievance_resolution.GET("/list_approved", api.ListApprovedGrievanceResolutions)
		grievance_resolution.GET("/list_denied", api.ListDeniedGrievanceResolutions)
		grievance_resolution.GET("/show/:id", api.ShowGrievanceResolution)
		grievance_resolution.POST("/update", api.UpdateGrievanceResolution)
		grievance_resolution.POST("/delete", api.DeleteGrievanceResolution)
	}

	grievance_forward := aim.Group("/grievance_forward")
	{
		grievance_forward.POST("/store", api.StoreGrievanceForward)
		grievance_forward.GET("/list", api.ListGrievanceForwards)
		grievance_forward.GET("/list_new", api.ListNewGrievanceForwards)
		grievance_forward.GET("/list_approved", api.ListApprovedGrievanceForwards)
		grievance_forward.GET("/list_denied", api.ListDeniedGrievanceForwards)
		grievance_forward.GET("/show/:id", api.ShowGrievanceForward)
		grievance_forward.POST("/update", api.UpdateGrievanceForward)
		grievance_forward.POST("/delete", api.DeleteGrievanceForward)
	}

	grievance_time_extension := aim.Group("/grievance_time_extension")
	{
		grievance_time_extension.POST("/store", api.StoreGrievanceTimeExtension)
		grievance_time_extension.GET("/list", api.ListGrievanceTimeExtensions)
		grievance_time_extension.GET("/list_new", api.ListNewGrievanceTimeExtensions)
		grievance_time_extension.GET("/list_approved", api.ListApprovedGrievanceTimeExtensions)
		grievance_time_extension.GET("/list_denied", api.ListDeniedGrievanceTimeExtensions)
		grievance_time_extension.GET("/show/:id", api.ShowGrievanceTimeExtension)
		grievance_time_extension.POST("/update", api.UpdateGrievanceTimeExtension)
		grievance_time_extension.POST("/delete", api.DeleteGrievanceTimeExtension)
	}

	grievance := aim.Group("/grievances")
	{
		grievance.POST("/store", api.StoreGrievance)
		grievance.GET("/list", api.ListGrievances)
		grievance.GET("/show/:id", api.ShowGrievance)
		grievance.POST("/update", api.UpdateGrievance)
		grievance.POST("/delete", api.DeleteGrievance)
	}


}
