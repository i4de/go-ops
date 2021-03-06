package service

import (
	"context"
	"errors"
	"go-ops/internal/model/entity"
	"go-ops/internal/service/internal/dao"
	v1 "go-ops/pkg/api/v1"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/v2/util/guid"
)

type (
	sApp struct{}
)

var (
	insApp = sApp{}
)

func App() *sApp {
	return &insApp
}

func (self *sApp) GetApp(ctx context.Context, appid string) (r *entity.App, err error) {
	var app *entity.App

	err = dao.App.Ctx(ctx).Where("appid = ?", appid).Scan(&app)
	if err != nil {
		return
	}

	if app == nil {
		err = errors.New("not found:" + appid)
		return
	}

	r = app
	return
}

func (self *sApp) Create(ctx context.Context, req *v1.AddAppReq) (res *v1.AddAppRes, err error) {

	appid := guid.S()
	apikey := guid.S() + guid.S()
	seckey := guid.S() + guid.S() + guid.S() + guid.S()
	app := &entity.App{
		Name:   req.Name,
		Owner:  req.Owner,
		Appid:  appid,
		ApiKey: apikey,
		SecKey: seckey,
	}

	_, err = dao.App.Ctx(ctx).Data(app).Insert()
	if err != nil {
		return
	}

	res = &v1.AddAppRes{
		Appid:  appid,
		SecKey: seckey,
		ApiKey: apikey,
		Name:   req.Name,
		Owner:  req.Owner,
	}
	return
}

func (self *sApp) Query(ctx context.Context, req *v1.QueryAppReq) (res *v1.QueryAppRes, err error) {

	m := g.Map{}

	if req.Name != "" {
		m["name"] = req.Name

	}

	if req.Owner != "" {
		m["owner"] = req.Owner
	}

	list := make([]*entity.App, 0)

	err = dao.App.Ctx(ctx).Where(m).Page(req.PageNum, req.PageSize).Scan(&list)

	if err != nil {
		return
	}

	count, err := dao.App.Ctx(ctx).Where(m).Count()
	if err != nil {
		return
	}

	res = &v1.QueryAppRes{
		List: list,
	}

	res.Total = count
	res.PageNum = req.PageNum
	res.PageSize = req.PageSize

	if res.Total%res.PageSize > 0 {
		res.PageTotal = 1
	}
	res.PageTotal += res.Total / res.PageSize
	return
}

func (self *sApp) SingleQuery(ctx context.Context, req *v1.QuerySingleAppReq) (res *v1.AddAppRes, err error) {
	var app *entity.App
	err = dao.App.Ctx(ctx).Where("appid = ?", req.AppId).Scan(&app)
	if err != nil {
		return
	}
	if app == nil {
		return
	}
	res = &v1.AddAppRes{
		Appid:  app.Appid,
		Name:   app.Name,
		Owner:  app.Owner,
		ApiKey: app.ApiKey,
		SecKey: app.SecKey,
		Status: app.Status,
	}

	return
}

func (self *sApp) Update(ctx context.Context, req *v1.UpdateAppReq) (res *v1.AddAppRes, err error) {

	var app *entity.App

	err = dao.App.Ctx(ctx).Where("appid = ?", req.Appid).Scan(&app)
	if err != nil {
		return
	}

	if app == nil {
		return
	}

	if req.Name != "" {
		app.Name = req.Name
	}

	if req.Owner != "" {
		app.Owner = req.Owner

	}

	if req.Status != 0 {
		app.Status = req.Status
	}

	_, err = dao.App.Ctx(ctx).Data(app).Where("appid = ?", req.Appid).Update()

	if err != nil {
		return
	}

	res = &v1.AddAppRes{
		Appid:  req.Appid,
		Name:   app.Name,
		Owner:  app.Owner,
		ApiKey: app.ApiKey,
		Status: app.Status,
	}

	return
}

func (self *sApp) Delete(ctx context.Context, req *v1.DeleteAppReq) (err error) {
	_, err = dao.App.Ctx(ctx).WhereIn("appid", req.Appids).Delete()
	return
}
