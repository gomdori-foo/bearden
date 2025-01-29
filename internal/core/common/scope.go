package common

type Scope string

const (
	ScopeDefault Scope = "DEFAULT"
	ScopeTransient Scope = "TRANSIENT"
	ScopeRequest Scope = "REQUEST"
)