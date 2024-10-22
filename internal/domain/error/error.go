package error

type Error struct {
	description string
	originalErr error
}

func (err *Error) Description() string {
	return err.description
}

func (err *Error) Error() string {
	if err.originalErr != nil {
		return err.originalErr.Error()
	}
	return err.description
}

func NewError(s string) *Error {
	return &Error{
		description: s,
	}
}

func WrapError(err *Error, originalErr error) *Error {
	return &Error{
		description: err.description,
		originalErr: originalErr,
	}
}

func ValidationError(err error) *Error {
	return &Error{
		description: InvalidInputErr.description,
		originalErr: err,
	}
}

// エラー変数を定義
var (
	InvalidInputErr   = NewError("入力の内容が不正です")
	NotFoundErr       = NewError("リソースが見つかりませんでした")
	UnAuthorizedErr   = NewError("認証に失敗しました")
	ForbiddenErr      = NewError("アクセス権限がありません")
	CognitoFailureErr = NewError("Cognitoでエラーが発生しました")
	GeneralDBError    = NewError("データベース操作中にエラーが発生しました")
)
