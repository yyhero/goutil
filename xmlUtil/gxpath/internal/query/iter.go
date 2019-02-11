package query

import "github.com/yyhero/goutil/xmlUtil/gxpath/xpath"

type Iterator interface {
	Current() xpath.NodeNavigator
}
