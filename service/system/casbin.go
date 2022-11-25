package system

import (
	"github.com/niudaii/gin-server/global"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"go.uber.org/zap"
	"sync"
)

type CasbinService struct{}

var (
	cachedEnforcer *casbin.CachedEnforcer
	once           sync.Once
)

func (s *CasbinService) Casbin() *casbin.CachedEnforcer {
	once.Do(func() {
		a, _ := gormadapter.NewAdapterByDB(global.DB)
		text := `
		[request_definition]
		r = sub, obj, act
		
		[policy_definition]
		p = sub, obj, act
		
		[role_definition]
		g = _, _
		
		[policy_effect]
		e = some(where (p.eft == allow))
		
		[matchers]
		m = r.sub == p.sub && keyMatch2(r.obj,p.obj) && r.act == p.act
		`
		m, err := model.NewModelFromString(text)
		if err != nil {
			global.Logger.Error("字符串加载模型失败", zap.Error(err))
			return
		}
		cachedEnforcer, _ = casbin.NewCachedEnforcer(m, a)
		cachedEnforcer.SetExpireTime(60 * 60)
		_ = cachedEnforcer.LoadPolicy()
	})
	return cachedEnforcer
}
