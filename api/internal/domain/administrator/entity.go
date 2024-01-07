package administrator

type Administrator struct {
	id    userUUID
	admin admin
}

// ドメイン バリューオブジェクト
type userUUID struct{ value string }
type admin struct{ value int }

// バリューオブジェクトの取得関数
func (u *Administrator) GetUUID() string { return u.id.value }
func (u *Administrator) GetAdmin() int   { return u.admin.value }

// 構造体生成関数
func NewAdministrator(id string, admin int) *Administrator {
	return newAdministrator(id, admin)
}

func newAdministrator(id string, administrator int) *Administrator {
	return &Administrator{
		id: userUUID{value: id},
		admin: admin{value: administrator},
	}
}