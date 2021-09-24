package interceptor

type MetadataKey string

const (
	MetadataKeyAuthorization MetadataKey = "authorization"
	MetadataKeyUserID        MetadataKey = "user_id"
	MetadataKeyRoleID        MetadataKey = "role_id"
)

func (x MetadataKey) String() string {
	return string(x)
}
