package constant

const (
	UserLevelNormal = 3 - iota
	UserLevelAdmin
	UserLevelRoot
)

// cache key
const (
	KeyUserDataByEmail = "user-data-by-email-%v"
)
