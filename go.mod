module github.com/omniboost/go-myob

go 1.25.0

require (
	github.com/gorilla/schema v1.4.1
	github.com/hashicorp/go-multierror v1.1.1
	github.com/pkg/errors v0.9.1
	golang.org/x/oauth2 v0.30.0
	gopkg.in/guregu/null.v3 v3.5.0
)

require github.com/hashicorp/errwrap v1.1.0 // indirect

replace github.com/gorilla/schema => github.com/omniboost/schema v1.1.1-0.20211111150515-2e872025e306
