package adminm

/**
Role签名token签名表,用于对比权限变化  后台更新权限，对应更新这个角色的Sign
*/

type RoleSign struct {
	ID   int64
	Uid  string `json:"uid"`
	Sign string `json:"sign"`
}
