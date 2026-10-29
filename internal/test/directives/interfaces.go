package directives

type Substituteable interface {
	Substitute(env map[string]string) error
}

type Assertable interface {
	Check()
	Message() string
	Details() string
}

type AssertableSubstitutable interface {
	Assertable
	Substituteable
}
