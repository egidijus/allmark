// Copyright 2015 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package handlers

import (
	"fmt"
	"net/http"

	"github.com/egidijus/allmark/common/route"
	"github.com/egidijus/allmark/web/header"
	"github.com/egidijus/allmark/web/orchestrator"
	"github.com/egidijus/allmark/web/view/templates"
	"github.com/egidijus/allmark/web/view/viewmodel"
)

func Tags(headerWriter header.HeaderWriter,
	navigationOrchestrator *orchestrator.NavigationOrchestrator,
	tagsOrchestrator *orchestrator.TagsOrchestrator,
	templateProvider templates.Provider) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// set headers
		headerWriter.Write(w, header.CONTENTTYPE_HTML)

		hostname := getBaseURLFromRequest(r)

		tagmapTemplate, err := templateProvider.GetTagMapTemplate(hostname)
		if err != nil {
			fmt.Fprintf(w, "Template not found. Error: %s", err)
			return
		}

		// Page parameters
		pageType := "tagmap"
		headline := "Tags"
		pageTitle := tagsOrchestrator.GetPageTitle(headline)

		pageModel := viewmodel.Model{}
		pageModel.Type = pageType
		pageModel.Title = headline
		pageModel.PageTitle = pageTitle
		pageModel.ToplevelNavigation = navigationOrchestrator.GetToplevelNavigation()
		pageModel.BreadcrumbNavigation = navigationOrchestrator.GetBreadcrumbNavigation(route.New())
		pageModel.TagCloud = tagsOrchestrator.GetTagCloud()

		tagsPageModel := viewmodel.Tags{}
		tagsPageModel.Model = pageModel
		tagsPageModel.Tags = tagsOrchestrator.GetTags()

		renderTemplate(tagmapTemplate, tagsPageModel, w)
	})
}
