package interfaces

import (
	"flamingo/core/cms/domain"
	"flamingo/framework/web"
	"flamingo/framework/web/responder"
)

type (
	// ViewController demonstrates a product view controller
	ViewController struct {
		*responder.ErrorAware  `inject:""`
		*responder.RenderAware `inject:""`
		PageService            domain.PageService `inject:""`
	}

	// ViewData for rendering
	ViewData struct {
		CmsPage *domain.Page
	}
)

// Get Response for Product matching sku param
func (vc *ViewController) Get(c web.Context) web.Response {
	var page, err = vc.PageService.Get(c, c.MustParam1("name"))
	if err != nil {
		return vc.Error(c, err)
	}
	if page == nil {
		return vc.ErrorNotFound(c, nil)
	}
	template, err := c.Param1("template")
	if err != nil {
		template = "pages/cms/view"
	}
	return vc.Render(c, template, ViewData{CmsPage: page})
}
