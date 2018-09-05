package query

import "moqikaka.com/goutil/xmlUtil/gxpath/xpath"

type Iterator interface {
	Current() xpath.NodeNavigator
}
