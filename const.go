// Copyright (c) 2020. All rights reserved.
// Use of this source code is governed by an MIT-style license that can be
// found in the LICENSE file.

package gogurls

const (
	HttpsScheme = "https"
)

// hosts
const (
	GogHost    = "gog.com"
	WwwGogHost = "www." + GogHost
)

// paths
const (
	detailsPathTemplate     = "/account/{mediaType}Details/{id}.json"
	productsPagePath        = "/games/ajax/filtered"
	accountProductsPagePath = "/account/getFilteredProducts"
	wishlistPath            = "/account/wishlist/search"
)
