package main

import (
	"net/http"

	prout "imuslab.com/aroz_online/mod/prouter"
)

func permissionInit(){
	//Register the permission handler, require authentication except listgroup
	adminRouter := prout.NewModuleRouter(prout.RouterOption{
		ModuleName: "System Setting", 
		AdminOnly: true, 
		UserHandler: userHandler, 
		DeniedHandler: func(w http.ResponseWriter, r *http.Request){
			sendErrorResponse(w, "Permission Denied");
		},
	});

	//Must be handled by default router
	http.HandleFunc("/system/permission/listgroup", permissionHandler.HandleListGroup)
	adminRouter.HandleFunc("/system/permission/newgroup", permissionHandler.HandleGroupCreate)
	adminRouter.HandleFunc("/system/permission/editgroup", permissionHandler.HandleGroupEdit)
	adminRouter.HandleFunc("/system/permission/delgroup", permissionHandler.HandleGroupRemove)

	registerSetting(settingModule{
		Name:         "Permission Groups",
		Desc:         "Handle the permission of access in groups",
		IconPath:     "SystemAO/users/img/small_icon.png",
		Group:        "Users",
		StartDir:     "SystemAO/users/group.html",
		RequireAdmin: true,
	})
}


