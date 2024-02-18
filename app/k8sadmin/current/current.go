/**
User: cr-mao
Date: 2023/02/18 17:08
Email: crmao@qq.com
Desc: current.go
*/
package current

import (
	"context"
)

func AdminUserToken(ctx context.Context) string {
	return ctx.Value("admin_user_token").(string)
}
