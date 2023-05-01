package errs

func SetUserError(err error, customErrorMessage string) error {
	if err == nil {
		return err
	}

	return &CustomError{
		ErrorMessage:       err.Error(),
		CustomErrorMessage: customErrorMessage,
	}
}

func GetUserError(err error) string {
	userFacingError, ok := err.(*CustomError)
	if !ok {
		return err.Error()
	}

	if userFacingError == nil {
		return err.Error()
	}

	if userFacingError.CustomErrorMessage != "" {
		return userFacingError.CustomErrorMessage
	}

	return userFacingError.ErrorMessage
}

func GetSystemError(err error) string {
	userFacingError, _ := err.(*CustomError)

	if userFacingError != nil {
		return userFacingError.ErrorMessage
	}

	return err.Error()
}

func (c *CustomError) Error() string {
	return c.ErrorMessage
}
