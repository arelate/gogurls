// Copyright (c) 2020. All rights reserved.
// Use of this source code is governed by an MIT-style license that can be
// found in the LICENSE file.

package gogurls

import (
	"github.com/boggydigital/gogog/pkg/gogtypes"
	"net/url"
	"strconv"
)

func AccountProductsPage(
	page int,
	mt gogtypes.Media,
	sortOrder gogtypes.AccountProductsSortOrder,
	updated bool,
	hidden bool) *url.URL {

	accountProductsPage := &url.URL{
		Scheme: HttpsScheme,
		Host:   GogHost,
		Path:   accountProductsPagePath,
	}

	hiddenFlag, updatedFlag := "0", "0"
	if hidden {
		hiddenFlag = "1"
	}

	if updated {
		updatedFlag = "1"
	}

	if sortOrder == "" {
		sortOrder = gogtypes.AccountProductsSortPurchaseDate
	}

	q := accountProductsPage.Query()
	q.Add("mediaType", strconv.Itoa(int(mt)))
	q.Add("sortBy", string(sortOrder))
	q.Add("page", strconv.Itoa(page))
	q.Add("hiddenFlag", hiddenFlag)
	q.Add("isUpdated", updatedFlag)
	accountProductsPage.RawQuery = q.Encode()

	return accountProductsPage
}