package middleware

import (
	"fmt"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

func UserPolicy() *casbin.Enforcer {

	adapter, err := gormadapter.NewAdapter("mysql", "root:@tcp(localhost:3306)/")
	if err != nil {
		panic(fmt.Sprintf("failed to create casbin enforcer: %v", err))
	}
	enforcer, err := casbin.NewEnforcer("./config/rbac_model.conf", adapter)
	if err != nil {
		panic(fmt.Sprintf("failed to create casbin enforcer: %v", err))
	}
	if hasPolicy := enforcer.HasPolicy("admin", "course", "read"); !hasPolicy {
		enforcer.AddPolicy("admin", "course", "read")
	}
	if hasPolicy := enforcer.HasPolicy("admin", "course", "write"); !hasPolicy {
		enforcer.AddPolicy("admin", "course", "write")
	}
	if hasPolicy := enforcer.HasPolicy("admin", "course", "delete"); !hasPolicy {
		enforcer.AddPolicy("admin", "course", "delete")
	}
	if hasPolicy := enforcer.HasPolicy("user", "course", "read"); !hasPolicy {
		enforcer.AddPolicy("user", "course", "read")
	}
	if hasPolicy := enforcer.HasPolicy("admin", "class", "read"); !hasPolicy {
		enforcer.AddPolicy("admin", "class", "read")
	}
	if hasPolicy := enforcer.HasPolicy("user", "class", "read"); !hasPolicy {
		enforcer.AddPolicy("user", "class", "read")
	}
	if hasPolicy := enforcer.HasPolicy("user", "class", "write"); !hasPolicy {
		enforcer.AddPolicy("user", "class", "write")
	}
	return enforcer
}
