package neodb_go_client

import "errors"

var NoShelfTypeErr = errors.New("no shelf type")
var NeoDBAPIRequestErr = errors.New("neodb api request error")
var InternalErr = errors.New("internal error")
var NeoDBTokenErr = errors.New("neodb token error")
