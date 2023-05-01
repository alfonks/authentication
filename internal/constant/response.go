package constant

// error response
const (
	SignUpServerError         = "Maaf, gagal melakukan sign up. Silahkan hubungi custommer support"
	SignUpDuplicateEmailError = "Maaf, email yang digunakan sudah terdaftar"

	LoginUserDataInvalidOrNotFound = "Maaf, email atau password yang anda masukan salah"
	LoginInternalServerError       = "Maaf, terjadi kesalahan saat memproses login. Silahkan dicoba lagi"

	GenerateNewtokenPairRefreshTokenInvalid = "Gagal memproses request, silahkan hubungi support"
	GenerateNewTokenPairRefreshTokenExpired = "Silahkan login kembali"
	GenerateTokenFail                       = "Gagal memproses token"

	AccessTokenInvalidError = "Maaf access token salah"
	AccessTokenExpiredError = "Maaf sesi telah habis"
	AccessTokenUnauthorized = "Maaf anda tidak memiliki akses untuk fitur ini"
	AccessTokenNotFound     = "Maaf anda belum login"
)

// success response
const (
	SignUpSuccess           = "Berhasil membuat user baru! Silahkan login"
	LoginSuccess            = "Berhasil login"
	GenerateNewTokenSuccess = "Berhasil men-generate new pair token"
)
