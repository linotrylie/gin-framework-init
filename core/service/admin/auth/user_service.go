package auth

import (
	"equity/core/dao/authdao"
	"equity/core/model/authmodel"
	"equity/core/model/authmodel/request"
	"equity/core/model/authmodel/response"
	"equity/core/model/commonreq"
	"equity/pkg/app"
	"equity/utils"
	"equity/utils/passwordutil"
	"equity/utils/strutil"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type UserService struct {
	userDao *authdao.UserDao
	roleDao *authdao.RoleDao
	nodeDao *authdao.NodeDao
}

// Login 登录获取JWT token 和过期时间
func (u *UserService) Login(c *gin.Context, req *request.UserLoginParam) (token string, tokenExp int, userName string, err error) {
	user, err := u.GetUser(c, req)
	if err != nil {
		return "", 0, "", err
	}

	if user == nil || user.Id < 1 {
		return "", 0, "", errors.New(fmt.Sprintf("该账号不存在,name=%s", req.UserName))
	}

	if user.IsDelete == authmodel.AuthUserIsDeleteTrue {
		return "", 0, "", errors.New(fmt.Sprintf("该账号已经被删除,id=%d", user.Id))
	}

	if user.Disable == authmodel.AuthUserDisableTrue {
		return "", 0, "", errors.New(fmt.Sprintf("该账号已经被禁用,id=%d", user.Id))
	}

	// 校验密码
	validatePwd := passwordutil.ValidatePasswd(req.Password, user.Salt, user.Password)
	if validatePwd != true {
		return "", 0, "", errors.New(fmt.Sprintf("密码错误,id=%d", user.Id))
	}

	// 生成token
	token, err = app.GenerateToken(strconv.Itoa(int(user.Id)), user.Salt, user.Id)
	if err != nil {
		return "", 0, "", err
	}

	if token == "" {
		return "", 0, "", errors.New("获取 token 错误")
	}

	// 获取 token 过期时间
	exp, err := app.GetJwtPayload(token)
	if err != nil {
		return token, 0, "", err
	}
	return token, exp.Exp, user.Name, nil
}

// GetUser 获取用户信息
func (u *UserService) GetUser(c *gin.Context, req *request.UserLoginParam) (AuthUserRes *response.AuthUserRes, err error) {
	user, err := u.userDao.GetUser(c, req)
	if user == nil {
		return nil, err
	}
	authUserRes := &response.AuthUserRes{
		Id:           user.Id,
		IsRoot:       user.IsRoot,
		Name:         user.Name,
		Account:      user.Account,
		Password:     user.Password,
		Salt:         user.Salt,
		CreateTime:   user.CreateTime,
		UpdateTime:   user.UpdateTime,
		Disable:      user.Disable,
		DeleteTime:   user.DeleteTime,
		LoginAddress: user.LoginAddress,
		IsDelete:     user.IsDelete,
	}
	return authUserRes, nil
}

// CreateUser 新增管理员
func (u *UserService) CreateUser(c *gin.Context, req *request.CreateUserParam) error {
	err := u.validateUser(c, req)
	if err != nil {
		return err
	}

	if utils.ContainsDuplicate(req.RoleIds) {
		return errors.New("角色id不能重复")
	}

	passwordAndSalt := u.generatePasswordAndSalt(req.Password)
	userModel := authmodel.TAuthUser{
		IsRoot:   authmodel.AuthUserIsRootFalse,
		Name:     req.UserName,
		Account:  req.Account,
		Password: passwordAndSalt.Password,
		Salt:     passwordAndSalt.Salt,
		IsDelete: authmodel.AuthUserIsDeleteFalse,
	}
	return u.userDao.CreateUser(c, &userModel, req.RoleIds)

}

// UpdateUser 编辑管理员
func (u *UserService) UpdateUser(c *gin.Context, req *request.UpdateUserParam) error {
	if req.Id < 1 {
		return errors.New("用户id不能为空")
	}

	if utils.ContainsDuplicate(req.RoleIds) {
		return errors.New("角色id不能重复")
	}

	// 判断用户是否存在
	user, err := u.userDao.GetUserById(c, req.Id)
	if err != nil {
		return err
	}
	if user == nil || user.Id != req.Id {
		return errors.New("获取用户失败")
	}
	count, err := u.userDao.UserNameIsExist(c, req.UserName, req.Id)
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("该用户名已经被占用")
	}
	return u.userDao.UpdateUser(c, *req)
}

// SetUserStatus 设置用户状态
func (u *UserService) SetUserStatus(c *gin.Context, req *request.SetUserStatusParam) error {
	user, err := u.userDao.GetUserById(c, req.UserId)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New(fmt.Sprintf("获取用户失败,用户id=%d", req.UserId))
	}
	if user.Id != req.UserId {
		return errors.New(fmt.Sprintf("获取用户失败,用户参数Uid=%d,结果uid=%d", req.UserId, user.Id))
	}

	// 值没有差异,不需更新
	if user.Disable == int8(req.Status) {
		return nil
	}

	// 超级管理员不能进禁用
	if user.IsRoot == authmodel.AuthUserIsRootTrue {
		return errors.New(fmt.Sprintf("超级管理员不能禁用,uid=%d", req.UserId))
	}

	user.Disable = int8(req.Status)
	user.UpdateTime = int32(time.Now().Unix())
	rowsAffected, err := u.userDao.SetUserStatus(c, user)
	if err != nil {
		return err
	}
	if rowsAffected != 1 {
		return errors.New(fmt.Sprintf("设置用户状态,成功数=%d", rowsAffected))
	}
	return nil
}

// ChangePassword 修改密码
func (u *UserService) ChangePassword(c *gin.Context, req *request.ChangeUserPWDParam) error {
	if len(req.Password) < 6 {
		return errors.New("密码至少6位")
	}

	if req.Password != req.ConfirmPassword {
		return errors.New("密码和确认密码不相同")
	}
	user, err := u.userDao.GetUserById(c, req.UserId)
	if err != nil {
		return err
	}
	if user == nil || user.Id != req.UserId {
		return errors.New(fmt.Sprintf("获取用户失败,用户id=%d", req.UserId))
	}
	if user.IsDelete == authmodel.AuthUserIsDeleteTrue {
		return errors.New(fmt.Sprintf("该用户已经删除,用户id=%d", req.UserId))
	}

	user.UpdateTime = int32(time.Now().Unix())
	passwordAndSalt := u.generatePasswordAndSalt(req.Password)
	user.Password = passwordAndSalt.Password
	user.Salt = passwordAndSalt.Salt

	rowsAffected, err := u.userDao.ChangePassword(c, user)
	if err != nil {
		return err
	}
	if rowsAffected != 1 {
		return errors.New(fmt.Sprintf("修改用户密码,成功数=%d", rowsAffected))
	}
	return nil
}

// Delete 设置用户状态
func (u *UserService) Delete(c *gin.Context, req *commonreq.PrimaryIdParam) error {
	// 不能删除自己的账号
	uidInt, err := utils.GetUidByToken(c)
	if err != nil {
		return err
	}
	if req.Id == uidInt {
		return errors.New(fmt.Sprintf("不能删除自己的账号,用户id=%d", req.Id))
	}

	user, err := u.userDao.GetUserById(c, int32(req.Id))
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New(fmt.Sprintf("用户已不存在,用户id=%d", req.Id))
	}
	if user.IsDelete == authmodel.AuthUserIsDeleteTrue {
		return errors.New(fmt.Sprintf("用户已删除,用户id=%d", req.Id))
	}
	if user.Id != int32(req.Id) {
		return errors.New(fmt.Sprintf("获取用户失败,用户id=%d", req.Id))
	}
	if user.IsRoot == authmodel.AuthUserIsRootTrue {
		return errors.New(fmt.Sprintf("超级管理员不能删除,用户id=%d", req.Id))
	}
	user.IsDelete = authmodel.AuthUserIsDeleteTrue
	user.DeleteTime = int32(time.Now().Unix())
	rowsAffected, err := u.userDao.Delete(c, user)
	if err != nil {
		return nil
	}
	if rowsAffected != 1 {
		return errors.New(fmt.Sprintf("删除用户失败,成功数=%d", rowsAffected))
	}
	return nil
}

// UserList 分页获取用户列表
func (u *UserService) UserList(c *gin.Context, req *request.SearchUserListParam) (err error, list *[]response.AuthUserBaseRes, total int64) {
	return u.userDao.UserList(c, req)
}

// UserDetail 获取用户所属详情
func (u *UserService) UserDetail(c *gin.Context, req *commonreq.PrimaryIdParam) (role *response.AuthUserDetailRes, err error) {
	userRoleList, err := u.roleDao.GetRoleByUserId(c, req.Id)
	if err != nil {
		return nil, err
	}
	var userRoleIds []*int32
	for _, userRole := range userRoleList {
		userRoleIds = append(userRoleIds, &userRole.RoleId)
	}

	// 用户所属的角色信息
	role = &response.AuthUserDetailRes{}
	role.Role = userRoleIds
	user, err := u.userDao.GetUserById(c, int32(req.Id))
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New(fmt.Sprintf("用户获取失败,uid=%d", req.Id))
	}

	// 用户基本信息
	userInfo := response.AuthUserBaseRes{
		Id:           user.Id,
		IsRoot:       user.IsRoot,
		Name:         user.Name,
		Account:      user.Account,
		CreateTime:   user.CreateTime,
		UpdateTime:   user.UpdateTime,
		LoginTime:    user.LoginTime,
		LoginIp:      user.LoginIp,
		Disable:      user.Disable,
		LoginAddress: user.LoginAddress,
	}
	role.Info = &userInfo
	return role, nil
}

// GetUserNodeUrls 获取用户应有的节点
func (u *UserService) GetUserNodeUrls(c *gin.Context, uid int32) ([]*string, error) {
	userRoleList, err := u.roleDao.GetRoleByUserId(c, uid)
	if err != nil {
		return nil, err
	}

	if len(userRoleList) < 1 {
		return nil, errors.New("该用户没有绑定任何角色")
	}

	// 角色id
	var userRoleIds []*int32
	for _, userRole := range userRoleList {
		userRoleIds = append(userRoleIds, &userRole.RoleId)
	}

	// 用户所属的角色信息
	nodeIds, err := u.nodeDao.GetNodeIdsByRoleIds(c, userRoleIds)
	if err != nil {
		return nil, err
	}

	if len(nodeIds) < 1 {
		return nil, errors.New("获取角色关联的节点失败")
	}

	urls, err := u.nodeDao.GetNodeUrls(c, nodeIds)
	if err != nil {
		return nil, err
	}

	// 获取所有节点url
	return urls, nil
}

// 生成密码和密码Salt
func (u *UserService) generatePasswordAndSalt(passwordPlaintext string) *response.AuthUserPwdAndSaltRes {
	salt := strutil.GetRandom(32, strutil.LetterNumbers)
	return &response.AuthUserPwdAndSaltRes{
		Password: passwordutil.MakePasswd(passwordPlaintext, salt),
		Salt:     strutil.GetRandom(32, strutil.LetterNumbers),
	}
}

// validateUser 验证用户信息
func (u *UserService) validateUser(c *gin.Context, req *request.CreateUserParam) error {
	if len(req.Password) < 6 {
		return errors.New("密码至少6位")
	}

	// 判断用户是否存在
	userCount, err := u.userDao.GetUserCountByName(c, req.UserName)
	if err != nil {
		return err
	}
	if userCount > 0 {
		return errors.New("该用户名已经被占用")
	}

	// 判断账号是否唯一
	count, err := u.userDao.GetUserByAccount(c, req.Account)
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("该账号已经被占用")
	}
	return nil
}
