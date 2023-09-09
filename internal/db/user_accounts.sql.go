// Code generated by sqlc. DO NOT EDIT.
// source: user_accounts.sql

package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const createConfirmToken = `-- name: CreateConfirmToken :one
INSERT INTO user_account_confirm_token (email) VALUES ($1) RETURNING confirm_token_id, email
`

func (q *Queries) CreateConfirmToken(ctx context.Context, email string) (UserAccountConfirmToken, error) {
	row := q.db.QueryRowContext(ctx, createConfirmToken, email)
	var i UserAccountConfirmToken
	err := row.Scan(&i.ConfirmTokenID, &i.Email)
	return i, err
}

const createInvitedProjectMember = `-- name: CreateInvitedProjectMember :one
INSERT INTO project_member_invited (project_id, user_account_invited_id) VALUES ($1, $2)
  RETURNING project_member_invited_id, project_id, user_account_invited_id
`

type CreateInvitedProjectMemberParams struct {
	ProjectID            uuid.UUID `json:"project_id"`
	UserAccountInvitedID uuid.UUID `json:"user_account_invited_id"`
}

func (q *Queries) CreateInvitedProjectMember(ctx context.Context, arg CreateInvitedProjectMemberParams) (ProjectMemberInvited, error) {
	row := q.db.QueryRowContext(ctx, createInvitedProjectMember, arg.ProjectID, arg.UserAccountInvitedID)
	var i ProjectMemberInvited
	err := row.Scan(&i.ProjectMemberInvitedID, &i.ProjectID, &i.UserAccountInvitedID)
	return i, err
}

const createInvitedUser = `-- name: CreateInvitedUser :one
INSERT INTO user_account_invited (email) VALUES ($1) RETURNING user_account_invited_id, email, invited_on, has_joined
`

func (q *Queries) CreateInvitedUser(ctx context.Context, email string) (UserAccountInvited, error) {
	row := q.db.QueryRowContext(ctx, createInvitedUser, email)
	var i UserAccountInvited
	err := row.Scan(
		&i.UserAccountInvitedID,
		&i.Email,
		&i.InvitedOn,
		&i.HasJoined,
	)
	return i, err
}

const createUserAccount = `-- name: CreateUserAccount :one
INSERT INTO user_account(full_name, initials, email, username, created_at, password_hash, role_code, active)
  VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING user_id, created_at, email, username, password_hash, profile_bg_color, full_name, initials, profile_avatar_url, role_code, bio, active
`

type CreateUserAccountParams struct {
	FullName     string    `json:"full_name"`
	Initials     string    `json:"initials"`
	Email        string    `json:"email"`
	Username     string    `json:"username"`
	CreatedAt    time.Time `json:"created_at"`
	PasswordHash string    `json:"password_hash"`
	RoleCode     string    `json:"role_code"`
	Active       bool      `json:"active"`
}

func (q *Queries) CreateUserAccount(ctx context.Context, arg CreateUserAccountParams) (UserAccount, error) {
	row := q.db.QueryRowContext(ctx, createUserAccount,
		arg.FullName,
		arg.Initials,
		arg.Email,
		arg.Username,
		arg.CreatedAt,
		arg.PasswordHash,
		arg.RoleCode,
		arg.Active,
	)
	var i UserAccount
	err := row.Scan(
		&i.UserID,
		&i.CreatedAt,
		&i.Email,
		&i.Username,
		&i.PasswordHash,
		&i.ProfileBgColor,
		&i.FullName,
		&i.Initials,
		&i.ProfileAvatarUrl,
		&i.RoleCode,
		&i.Bio,
		&i.Active,
	)
	return i, err
}

const deleteConfirmTokenForEmail = `-- name: DeleteConfirmTokenForEmail :exec
DELETE FROM user_account_confirm_token WHERE email = $1
`

func (q *Queries) DeleteConfirmTokenForEmail(ctx context.Context, email string) error {
	_, err := q.db.ExecContext(ctx, deleteConfirmTokenForEmail, email)
	return err
}

const deleteInvitedUserAccount = `-- name: DeleteInvitedUserAccount :one
DELETE FROM user_account_invited WHERE user_account_invited_id = $1 RETURNING user_account_invited_id, email, invited_on, has_joined
`

func (q *Queries) DeleteInvitedUserAccount(ctx context.Context, userAccountInvitedID uuid.UUID) (UserAccountInvited, error) {
	row := q.db.QueryRowContext(ctx, deleteInvitedUserAccount, userAccountInvitedID)
	var i UserAccountInvited
	err := row.Scan(
		&i.UserAccountInvitedID,
		&i.Email,
		&i.InvitedOn,
		&i.HasJoined,
	)
	return i, err
}

const deleteProjectMemberInvitedForEmail = `-- name: DeleteProjectMemberInvitedForEmail :exec
DELETE FROM project_member_invited WHERE project_member_invited_id IN (
  SELECT pmi.project_member_invited_id FROM user_account_invited AS uai
  INNER JOIN project_member_invited AS pmi
  ON pmi.user_account_invited_id = uai.user_account_invited_id
  WHERE uai.email = $1
)
`

func (q *Queries) DeleteProjectMemberInvitedForEmail(ctx context.Context, email string) error {
	_, err := q.db.ExecContext(ctx, deleteProjectMemberInvitedForEmail, email)
	return err
}

const deleteUserAccountByID = `-- name: DeleteUserAccountByID :exec
DELETE FROM user_account WHERE user_id = $1
`

func (q *Queries) DeleteUserAccountByID(ctx context.Context, userID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteUserAccountByID, userID)
	return err
}

const deleteUserAccountInvitedForEmail = `-- name: DeleteUserAccountInvitedForEmail :exec
DELETE FROM user_account_invited WHERE email = $1
`

func (q *Queries) DeleteUserAccountInvitedForEmail(ctx context.Context, email string) error {
	_, err := q.db.ExecContext(ctx, deleteUserAccountInvitedForEmail, email)
	return err
}

const doesUserExist = `-- name: DoesUserExist :one
SELECT EXISTS(SELECT 1 FROM user_account WHERE email = $1 OR username = $2)
`

type DoesUserExistParams struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

func (q *Queries) DoesUserExist(ctx context.Context, arg DoesUserExistParams) (bool, error) {
	row := q.db.QueryRowContext(ctx, doesUserExist, arg.Email, arg.Username)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const getAllUserAccounts = `-- name: GetAllUserAccounts :many
SELECT user_id, created_at, email, username, password_hash, profile_bg_color, full_name, initials, profile_avatar_url, role_code, bio, active FROM user_account WHERE username != 'system'
`

func (q *Queries) GetAllUserAccounts(ctx context.Context) ([]UserAccount, error) {
	rows, err := q.db.QueryContext(ctx, getAllUserAccounts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []UserAccount
	for rows.Next() {
		var i UserAccount
		if err := rows.Scan(
			&i.UserID,
			&i.CreatedAt,
			&i.Email,
			&i.Username,
			&i.PasswordHash,
			&i.ProfileBgColor,
			&i.FullName,
			&i.Initials,
			&i.ProfileAvatarUrl,
			&i.RoleCode,
			&i.Bio,
			&i.Active,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getConfirmTokenByEmail = `-- name: GetConfirmTokenByEmail :one
SELECT confirm_token_id, email FROM user_account_confirm_token WHERE email = $1
`

func (q *Queries) GetConfirmTokenByEmail(ctx context.Context, email string) (UserAccountConfirmToken, error) {
	row := q.db.QueryRowContext(ctx, getConfirmTokenByEmail, email)
	var i UserAccountConfirmToken
	err := row.Scan(&i.ConfirmTokenID, &i.Email)
	return i, err
}

const getConfirmTokenByID = `-- name: GetConfirmTokenByID :one
SELECT confirm_token_id, email FROM user_account_confirm_token WHERE confirm_token_id = $1
`

func (q *Queries) GetConfirmTokenByID(ctx context.Context, confirmTokenID uuid.UUID) (UserAccountConfirmToken, error) {
	row := q.db.QueryRowContext(ctx, getConfirmTokenByID, confirmTokenID)
	var i UserAccountConfirmToken
	err := row.Scan(&i.ConfirmTokenID, &i.Email)
	return i, err
}

const getInvitedUserAccounts = `-- name: GetInvitedUserAccounts :many
SELECT user_account_invited_id, email, invited_on, has_joined FROM user_account_invited
`

func (q *Queries) GetInvitedUserAccounts(ctx context.Context) ([]UserAccountInvited, error) {
	rows, err := q.db.QueryContext(ctx, getInvitedUserAccounts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []UserAccountInvited
	for rows.Next() {
		var i UserAccountInvited
		if err := rows.Scan(
			&i.UserAccountInvitedID,
			&i.Email,
			&i.InvitedOn,
			&i.HasJoined,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getInvitedUserByEmail = `-- name: GetInvitedUserByEmail :one
SELECT user_account_invited_id, email, invited_on, has_joined FROM user_account_invited WHERE email = $1
`

func (q *Queries) GetInvitedUserByEmail(ctx context.Context, email string) (UserAccountInvited, error) {
	row := q.db.QueryRowContext(ctx, getInvitedUserByEmail, email)
	var i UserAccountInvited
	err := row.Scan(
		&i.UserAccountInvitedID,
		&i.Email,
		&i.InvitedOn,
		&i.HasJoined,
	)
	return i, err
}

const getMemberData = `-- name: GetMemberData :many
SELECT user_id, created_at, email, username, password_hash, profile_bg_color, full_name, initials, profile_avatar_url, role_code, bio, active FROM user_account
  WHERE username != 'system'
  AND user_id NOT IN (SELECT user_id FROM project_member WHERE project_id = $1)
`

func (q *Queries) GetMemberData(ctx context.Context, projectID uuid.UUID) ([]UserAccount, error) {
	rows, err := q.db.QueryContext(ctx, getMemberData, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []UserAccount
	for rows.Next() {
		var i UserAccount
		if err := rows.Scan(
			&i.UserID,
			&i.CreatedAt,
			&i.Email,
			&i.Username,
			&i.PasswordHash,
			&i.ProfileBgColor,
			&i.FullName,
			&i.Initials,
			&i.ProfileAvatarUrl,
			&i.RoleCode,
			&i.Bio,
			&i.Active,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getProjectsForInvitedMember = `-- name: GetProjectsForInvitedMember :many
SELECT project_id FROM user_account_invited AS uai
  INNER JOIN project_member_invited AS pmi
  ON pmi.user_account_invited_id = uai.user_account_invited_id
  WHERE uai.email = $1
`

func (q *Queries) GetProjectsForInvitedMember(ctx context.Context, email string) ([]uuid.UUID, error) {
	rows, err := q.db.QueryContext(ctx, getProjectsForInvitedMember, email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []uuid.UUID
	for rows.Next() {
		var project_id uuid.UUID
		if err := rows.Scan(&project_id); err != nil {
			return nil, err
		}
		items = append(items, project_id)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getRoleForUserID = `-- name: GetRoleForUserID :one
SELECT username, role.code, role.name FROM user_account
  INNER JOIN role ON role.code = user_account.role_code
WHERE user_id = $1
`

type GetRoleForUserIDRow struct {
	Username string `json:"username"`
	Code     string `json:"code"`
	Name     string `json:"name"`
}

func (q *Queries) GetRoleForUserID(ctx context.Context, userID uuid.UUID) (GetRoleForUserIDRow, error) {
	row := q.db.QueryRowContext(ctx, getRoleForUserID, userID)
	var i GetRoleForUserIDRow
	err := row.Scan(&i.Username, &i.Code, &i.Name)
	return i, err
}

const getUserAccountByEmail = `-- name: GetUserAccountByEmail :one
SELECT user_id, created_at, email, username, password_hash, profile_bg_color, full_name, initials, profile_avatar_url, role_code, bio, active FROM user_account WHERE email = $1
`

func (q *Queries) GetUserAccountByEmail(ctx context.Context, email string) (UserAccount, error) {
	row := q.db.QueryRowContext(ctx, getUserAccountByEmail, email)
	var i UserAccount
	err := row.Scan(
		&i.UserID,
		&i.CreatedAt,
		&i.Email,
		&i.Username,
		&i.PasswordHash,
		&i.ProfileBgColor,
		&i.FullName,
		&i.Initials,
		&i.ProfileAvatarUrl,
		&i.RoleCode,
		&i.Bio,
		&i.Active,
	)
	return i, err
}

const getUserAccountByID = `-- name: GetUserAccountByID :one
SELECT user_id, created_at, email, username, password_hash, profile_bg_color, full_name, initials, profile_avatar_url, role_code, bio, active FROM user_account WHERE user_id = $1
`

func (q *Queries) GetUserAccountByID(ctx context.Context, userID uuid.UUID) (UserAccount, error) {
	row := q.db.QueryRowContext(ctx, getUserAccountByID, userID)
	var i UserAccount
	err := row.Scan(
		&i.UserID,
		&i.CreatedAt,
		&i.Email,
		&i.Username,
		&i.PasswordHash,
		&i.ProfileBgColor,
		&i.FullName,
		&i.Initials,
		&i.ProfileAvatarUrl,
		&i.RoleCode,
		&i.Bio,
		&i.Active,
	)
	return i, err
}

const getUserAccountByUsername = `-- name: GetUserAccountByUsername :one
SELECT user_id, created_at, email, username, password_hash, profile_bg_color, full_name, initials, profile_avatar_url, role_code, bio, active FROM user_account WHERE username = $1
`

func (q *Queries) GetUserAccountByUsername(ctx context.Context, username string) (UserAccount, error) {
	row := q.db.QueryRowContext(ctx, getUserAccountByUsername, username)
	var i UserAccount
	err := row.Scan(
		&i.UserID,
		&i.CreatedAt,
		&i.Email,
		&i.Username,
		&i.PasswordHash,
		&i.ProfileBgColor,
		&i.FullName,
		&i.Initials,
		&i.ProfileAvatarUrl,
		&i.RoleCode,
		&i.Bio,
		&i.Active,
	)
	return i, err
}

const hasActiveUser = `-- name: HasActiveUser :one
SELECT EXISTS(SELECT 1 FROM user_account WHERE username != 'system' AND active = true)
`

func (q *Queries) HasActiveUser(ctx context.Context) (bool, error) {
	row := q.db.QueryRowContext(ctx, hasActiveUser)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const hasAnyUser = `-- name: HasAnyUser :one
SELECT EXISTS(SELECT 1 FROM user_account WHERE username != 'system')
`

func (q *Queries) HasAnyUser(ctx context.Context) (bool, error) {
	row := q.db.QueryRowContext(ctx, hasAnyUser)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const setFirstUserActive = `-- name: SetFirstUserActive :one
UPDATE user_account SET active = true WHERE user_id = (
  SELECT user_id from user_account WHERE active = false LIMIT 1
) RETURNING user_id, created_at, email, username, password_hash, profile_bg_color, full_name, initials, profile_avatar_url, role_code, bio, active
`

func (q *Queries) SetFirstUserActive(ctx context.Context) (UserAccount, error) {
	row := q.db.QueryRowContext(ctx, setFirstUserActive)
	var i UserAccount
	err := row.Scan(
		&i.UserID,
		&i.CreatedAt,
		&i.Email,
		&i.Username,
		&i.PasswordHash,
		&i.ProfileBgColor,
		&i.FullName,
		&i.Initials,
		&i.ProfileAvatarUrl,
		&i.RoleCode,
		&i.Bio,
		&i.Active,
	)
	return i, err
}

const setUserActiveByEmail = `-- name: SetUserActiveByEmail :one
UPDATE user_account SET active = true WHERE email = $1 RETURNING user_id, created_at, email, username, password_hash, profile_bg_color, full_name, initials, profile_avatar_url, role_code, bio, active
`

func (q *Queries) SetUserActiveByEmail(ctx context.Context, email string) (UserAccount, error) {
	row := q.db.QueryRowContext(ctx, setUserActiveByEmail, email)
	var i UserAccount
	err := row.Scan(
		&i.UserID,
		&i.CreatedAt,
		&i.Email,
		&i.Username,
		&i.PasswordHash,
		&i.ProfileBgColor,
		&i.FullName,
		&i.Initials,
		&i.ProfileAvatarUrl,
		&i.RoleCode,
		&i.Bio,
		&i.Active,
	)
	return i, err
}

const setUserPassword = `-- name: SetUserPassword :one
UPDATE user_account SET password_hash = $2 WHERE user_id = $1 RETURNING user_id, created_at, email, username, password_hash, profile_bg_color, full_name, initials, profile_avatar_url, role_code, bio, active
`

type SetUserPasswordParams struct {
	UserID       uuid.UUID `json:"user_id"`
	PasswordHash string    `json:"password_hash"`
}

func (q *Queries) SetUserPassword(ctx context.Context, arg SetUserPasswordParams) (UserAccount, error) {
	row := q.db.QueryRowContext(ctx, setUserPassword, arg.UserID, arg.PasswordHash)
	var i UserAccount
	err := row.Scan(
		&i.UserID,
		&i.CreatedAt,
		&i.Email,
		&i.Username,
		&i.PasswordHash,
		&i.ProfileBgColor,
		&i.FullName,
		&i.Initials,
		&i.ProfileAvatarUrl,
		&i.RoleCode,
		&i.Bio,
		&i.Active,
	)
	return i, err
}

const updateUserAccountInfo = `-- name: UpdateUserAccountInfo :one
UPDATE user_account SET bio = $2, full_name = $3, initials = $4, email = $5
  WHERE user_id = $1 RETURNING user_id, created_at, email, username, password_hash, profile_bg_color, full_name, initials, profile_avatar_url, role_code, bio, active
`

type UpdateUserAccountInfoParams struct {
	UserID   uuid.UUID `json:"user_id"`
	Bio      string    `json:"bio"`
	FullName string    `json:"full_name"`
	Initials string    `json:"initials"`
	Email    string    `json:"email"`
}

func (q *Queries) UpdateUserAccountInfo(ctx context.Context, arg UpdateUserAccountInfoParams) (UserAccount, error) {
	row := q.db.QueryRowContext(ctx, updateUserAccountInfo,
		arg.UserID,
		arg.Bio,
		arg.FullName,
		arg.Initials,
		arg.Email,
	)
	var i UserAccount
	err := row.Scan(
		&i.UserID,
		&i.CreatedAt,
		&i.Email,
		&i.Username,
		&i.PasswordHash,
		&i.ProfileBgColor,
		&i.FullName,
		&i.Initials,
		&i.ProfileAvatarUrl,
		&i.RoleCode,
		&i.Bio,
		&i.Active,
	)
	return i, err
}

const updateUserAccountProfileAvatarURL = `-- name: UpdateUserAccountProfileAvatarURL :one
UPDATE user_account SET profile_avatar_url = $2 WHERE user_id = $1
  RETURNING user_id, created_at, email, username, password_hash, profile_bg_color, full_name, initials, profile_avatar_url, role_code, bio, active
`

type UpdateUserAccountProfileAvatarURLParams struct {
	UserID           uuid.UUID      `json:"user_id"`
	ProfileAvatarUrl sql.NullString `json:"profile_avatar_url"`
}

func (q *Queries) UpdateUserAccountProfileAvatarURL(ctx context.Context, arg UpdateUserAccountProfileAvatarURLParams) (UserAccount, error) {
	row := q.db.QueryRowContext(ctx, updateUserAccountProfileAvatarURL, arg.UserID, arg.ProfileAvatarUrl)
	var i UserAccount
	err := row.Scan(
		&i.UserID,
		&i.CreatedAt,
		&i.Email,
		&i.Username,
		&i.PasswordHash,
		&i.ProfileBgColor,
		&i.FullName,
		&i.Initials,
		&i.ProfileAvatarUrl,
		&i.RoleCode,
		&i.Bio,
		&i.Active,
	)
	return i, err
}

const updateUserRole = `-- name: UpdateUserRole :one
UPDATE user_account SET role_code = $2 WHERE user_id = $1 RETURNING user_id, created_at, email, username, password_hash, profile_bg_color, full_name, initials, profile_avatar_url, role_code, bio, active
`

type UpdateUserRoleParams struct {
	UserID   uuid.UUID `json:"user_id"`
	RoleCode string    `json:"role_code"`
}

func (q *Queries) UpdateUserRole(ctx context.Context, arg UpdateUserRoleParams) (UserAccount, error) {
	row := q.db.QueryRowContext(ctx, updateUserRole, arg.UserID, arg.RoleCode)
	var i UserAccount
	err := row.Scan(
		&i.UserID,
		&i.CreatedAt,
		&i.Email,
		&i.Username,
		&i.PasswordHash,
		&i.ProfileBgColor,
		&i.FullName,
		&i.Initials,
		&i.ProfileAvatarUrl,
		&i.RoleCode,
		&i.Bio,
		&i.Active,
	)
	return i, err
}