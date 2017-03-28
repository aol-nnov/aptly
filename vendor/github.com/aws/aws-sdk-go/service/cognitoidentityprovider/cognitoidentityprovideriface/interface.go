// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package cognitoidentityprovideriface provides an interface to enable mocking the Amazon Cognito Identity Provider service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package cognitoidentityprovideriface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

// CognitoIdentityProviderAPI provides an interface to enable mocking the
// cognitoidentityprovider.CognitoIdentityProvider service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Cognito Identity Provider.
//    func myFunc(svc cognitoidentityprovideriface.CognitoIdentityProviderAPI) bool {
//        // Make svc.AddCustomAttributes request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := cognitoidentityprovider.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockCognitoIdentityProviderClient struct {
//        cognitoidentityprovideriface.CognitoIdentityProviderAPI
//    }
//    func (m *mockCognitoIdentityProviderClient) AddCustomAttributes(input *cognitoidentityprovider.AddCustomAttributesInput) (*cognitoidentityprovider.AddCustomAttributesOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockCognitoIdentityProviderClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type CognitoIdentityProviderAPI interface {
	AddCustomAttributes(*cognitoidentityprovider.AddCustomAttributesInput) (*cognitoidentityprovider.AddCustomAttributesOutput, error)
	AddCustomAttributesWithContext(aws.Context, *cognitoidentityprovider.AddCustomAttributesInput, ...request.Option) (*cognitoidentityprovider.AddCustomAttributesOutput, error)
	AddCustomAttributesRequest(*cognitoidentityprovider.AddCustomAttributesInput) (*request.Request, *cognitoidentityprovider.AddCustomAttributesOutput)

	AdminAddUserToGroup(*cognitoidentityprovider.AdminAddUserToGroupInput) (*cognitoidentityprovider.AdminAddUserToGroupOutput, error)
	AdminAddUserToGroupWithContext(aws.Context, *cognitoidentityprovider.AdminAddUserToGroupInput, ...request.Option) (*cognitoidentityprovider.AdminAddUserToGroupOutput, error)
	AdminAddUserToGroupRequest(*cognitoidentityprovider.AdminAddUserToGroupInput) (*request.Request, *cognitoidentityprovider.AdminAddUserToGroupOutput)

	AdminConfirmSignUp(*cognitoidentityprovider.AdminConfirmSignUpInput) (*cognitoidentityprovider.AdminConfirmSignUpOutput, error)
	AdminConfirmSignUpWithContext(aws.Context, *cognitoidentityprovider.AdminConfirmSignUpInput, ...request.Option) (*cognitoidentityprovider.AdminConfirmSignUpOutput, error)
	AdminConfirmSignUpRequest(*cognitoidentityprovider.AdminConfirmSignUpInput) (*request.Request, *cognitoidentityprovider.AdminConfirmSignUpOutput)

	AdminCreateUser(*cognitoidentityprovider.AdminCreateUserInput) (*cognitoidentityprovider.AdminCreateUserOutput, error)
	AdminCreateUserWithContext(aws.Context, *cognitoidentityprovider.AdminCreateUserInput, ...request.Option) (*cognitoidentityprovider.AdminCreateUserOutput, error)
	AdminCreateUserRequest(*cognitoidentityprovider.AdminCreateUserInput) (*request.Request, *cognitoidentityprovider.AdminCreateUserOutput)

	AdminDeleteUser(*cognitoidentityprovider.AdminDeleteUserInput) (*cognitoidentityprovider.AdminDeleteUserOutput, error)
	AdminDeleteUserWithContext(aws.Context, *cognitoidentityprovider.AdminDeleteUserInput, ...request.Option) (*cognitoidentityprovider.AdminDeleteUserOutput, error)
	AdminDeleteUserRequest(*cognitoidentityprovider.AdminDeleteUserInput) (*request.Request, *cognitoidentityprovider.AdminDeleteUserOutput)

	AdminDeleteUserAttributes(*cognitoidentityprovider.AdminDeleteUserAttributesInput) (*cognitoidentityprovider.AdminDeleteUserAttributesOutput, error)
	AdminDeleteUserAttributesWithContext(aws.Context, *cognitoidentityprovider.AdminDeleteUserAttributesInput, ...request.Option) (*cognitoidentityprovider.AdminDeleteUserAttributesOutput, error)
	AdminDeleteUserAttributesRequest(*cognitoidentityprovider.AdminDeleteUserAttributesInput) (*request.Request, *cognitoidentityprovider.AdminDeleteUserAttributesOutput)

	AdminDisableUser(*cognitoidentityprovider.AdminDisableUserInput) (*cognitoidentityprovider.AdminDisableUserOutput, error)
	AdminDisableUserWithContext(aws.Context, *cognitoidentityprovider.AdminDisableUserInput, ...request.Option) (*cognitoidentityprovider.AdminDisableUserOutput, error)
	AdminDisableUserRequest(*cognitoidentityprovider.AdminDisableUserInput) (*request.Request, *cognitoidentityprovider.AdminDisableUserOutput)

	AdminEnableUser(*cognitoidentityprovider.AdminEnableUserInput) (*cognitoidentityprovider.AdminEnableUserOutput, error)
	AdminEnableUserWithContext(aws.Context, *cognitoidentityprovider.AdminEnableUserInput, ...request.Option) (*cognitoidentityprovider.AdminEnableUserOutput, error)
	AdminEnableUserRequest(*cognitoidentityprovider.AdminEnableUserInput) (*request.Request, *cognitoidentityprovider.AdminEnableUserOutput)

	AdminForgetDevice(*cognitoidentityprovider.AdminForgetDeviceInput) (*cognitoidentityprovider.AdminForgetDeviceOutput, error)
	AdminForgetDeviceWithContext(aws.Context, *cognitoidentityprovider.AdminForgetDeviceInput, ...request.Option) (*cognitoidentityprovider.AdminForgetDeviceOutput, error)
	AdminForgetDeviceRequest(*cognitoidentityprovider.AdminForgetDeviceInput) (*request.Request, *cognitoidentityprovider.AdminForgetDeviceOutput)

	AdminGetDevice(*cognitoidentityprovider.AdminGetDeviceInput) (*cognitoidentityprovider.AdminGetDeviceOutput, error)
	AdminGetDeviceWithContext(aws.Context, *cognitoidentityprovider.AdminGetDeviceInput, ...request.Option) (*cognitoidentityprovider.AdminGetDeviceOutput, error)
	AdminGetDeviceRequest(*cognitoidentityprovider.AdminGetDeviceInput) (*request.Request, *cognitoidentityprovider.AdminGetDeviceOutput)

	AdminGetUser(*cognitoidentityprovider.AdminGetUserInput) (*cognitoidentityprovider.AdminGetUserOutput, error)
	AdminGetUserWithContext(aws.Context, *cognitoidentityprovider.AdminGetUserInput, ...request.Option) (*cognitoidentityprovider.AdminGetUserOutput, error)
	AdminGetUserRequest(*cognitoidentityprovider.AdminGetUserInput) (*request.Request, *cognitoidentityprovider.AdminGetUserOutput)

	AdminInitiateAuth(*cognitoidentityprovider.AdminInitiateAuthInput) (*cognitoidentityprovider.AdminInitiateAuthOutput, error)
	AdminInitiateAuthWithContext(aws.Context, *cognitoidentityprovider.AdminInitiateAuthInput, ...request.Option) (*cognitoidentityprovider.AdminInitiateAuthOutput, error)
	AdminInitiateAuthRequest(*cognitoidentityprovider.AdminInitiateAuthInput) (*request.Request, *cognitoidentityprovider.AdminInitiateAuthOutput)

	AdminListDevices(*cognitoidentityprovider.AdminListDevicesInput) (*cognitoidentityprovider.AdminListDevicesOutput, error)
	AdminListDevicesWithContext(aws.Context, *cognitoidentityprovider.AdminListDevicesInput, ...request.Option) (*cognitoidentityprovider.AdminListDevicesOutput, error)
	AdminListDevicesRequest(*cognitoidentityprovider.AdminListDevicesInput) (*request.Request, *cognitoidentityprovider.AdminListDevicesOutput)

	AdminListGroupsForUser(*cognitoidentityprovider.AdminListGroupsForUserInput) (*cognitoidentityprovider.AdminListGroupsForUserOutput, error)
	AdminListGroupsForUserWithContext(aws.Context, *cognitoidentityprovider.AdminListGroupsForUserInput, ...request.Option) (*cognitoidentityprovider.AdminListGroupsForUserOutput, error)
	AdminListGroupsForUserRequest(*cognitoidentityprovider.AdminListGroupsForUserInput) (*request.Request, *cognitoidentityprovider.AdminListGroupsForUserOutput)

	AdminRemoveUserFromGroup(*cognitoidentityprovider.AdminRemoveUserFromGroupInput) (*cognitoidentityprovider.AdminRemoveUserFromGroupOutput, error)
	AdminRemoveUserFromGroupWithContext(aws.Context, *cognitoidentityprovider.AdminRemoveUserFromGroupInput, ...request.Option) (*cognitoidentityprovider.AdminRemoveUserFromGroupOutput, error)
	AdminRemoveUserFromGroupRequest(*cognitoidentityprovider.AdminRemoveUserFromGroupInput) (*request.Request, *cognitoidentityprovider.AdminRemoveUserFromGroupOutput)

	AdminResetUserPassword(*cognitoidentityprovider.AdminResetUserPasswordInput) (*cognitoidentityprovider.AdminResetUserPasswordOutput, error)
	AdminResetUserPasswordWithContext(aws.Context, *cognitoidentityprovider.AdminResetUserPasswordInput, ...request.Option) (*cognitoidentityprovider.AdminResetUserPasswordOutput, error)
	AdminResetUserPasswordRequest(*cognitoidentityprovider.AdminResetUserPasswordInput) (*request.Request, *cognitoidentityprovider.AdminResetUserPasswordOutput)

	AdminRespondToAuthChallenge(*cognitoidentityprovider.AdminRespondToAuthChallengeInput) (*cognitoidentityprovider.AdminRespondToAuthChallengeOutput, error)
	AdminRespondToAuthChallengeWithContext(aws.Context, *cognitoidentityprovider.AdminRespondToAuthChallengeInput, ...request.Option) (*cognitoidentityprovider.AdminRespondToAuthChallengeOutput, error)
	AdminRespondToAuthChallengeRequest(*cognitoidentityprovider.AdminRespondToAuthChallengeInput) (*request.Request, *cognitoidentityprovider.AdminRespondToAuthChallengeOutput)

	AdminSetUserSettings(*cognitoidentityprovider.AdminSetUserSettingsInput) (*cognitoidentityprovider.AdminSetUserSettingsOutput, error)
	AdminSetUserSettingsWithContext(aws.Context, *cognitoidentityprovider.AdminSetUserSettingsInput, ...request.Option) (*cognitoidentityprovider.AdminSetUserSettingsOutput, error)
	AdminSetUserSettingsRequest(*cognitoidentityprovider.AdminSetUserSettingsInput) (*request.Request, *cognitoidentityprovider.AdminSetUserSettingsOutput)

	AdminUpdateDeviceStatus(*cognitoidentityprovider.AdminUpdateDeviceStatusInput) (*cognitoidentityprovider.AdminUpdateDeviceStatusOutput, error)
	AdminUpdateDeviceStatusWithContext(aws.Context, *cognitoidentityprovider.AdminUpdateDeviceStatusInput, ...request.Option) (*cognitoidentityprovider.AdminUpdateDeviceStatusOutput, error)
	AdminUpdateDeviceStatusRequest(*cognitoidentityprovider.AdminUpdateDeviceStatusInput) (*request.Request, *cognitoidentityprovider.AdminUpdateDeviceStatusOutput)

	AdminUpdateUserAttributes(*cognitoidentityprovider.AdminUpdateUserAttributesInput) (*cognitoidentityprovider.AdminUpdateUserAttributesOutput, error)
	AdminUpdateUserAttributesWithContext(aws.Context, *cognitoidentityprovider.AdminUpdateUserAttributesInput, ...request.Option) (*cognitoidentityprovider.AdminUpdateUserAttributesOutput, error)
	AdminUpdateUserAttributesRequest(*cognitoidentityprovider.AdminUpdateUserAttributesInput) (*request.Request, *cognitoidentityprovider.AdminUpdateUserAttributesOutput)

	AdminUserGlobalSignOut(*cognitoidentityprovider.AdminUserGlobalSignOutInput) (*cognitoidentityprovider.AdminUserGlobalSignOutOutput, error)
	AdminUserGlobalSignOutWithContext(aws.Context, *cognitoidentityprovider.AdminUserGlobalSignOutInput, ...request.Option) (*cognitoidentityprovider.AdminUserGlobalSignOutOutput, error)
	AdminUserGlobalSignOutRequest(*cognitoidentityprovider.AdminUserGlobalSignOutInput) (*request.Request, *cognitoidentityprovider.AdminUserGlobalSignOutOutput)

	ChangePassword(*cognitoidentityprovider.ChangePasswordInput) (*cognitoidentityprovider.ChangePasswordOutput, error)
	ChangePasswordWithContext(aws.Context, *cognitoidentityprovider.ChangePasswordInput, ...request.Option) (*cognitoidentityprovider.ChangePasswordOutput, error)
	ChangePasswordRequest(*cognitoidentityprovider.ChangePasswordInput) (*request.Request, *cognitoidentityprovider.ChangePasswordOutput)

	ConfirmDevice(*cognitoidentityprovider.ConfirmDeviceInput) (*cognitoidentityprovider.ConfirmDeviceOutput, error)
	ConfirmDeviceWithContext(aws.Context, *cognitoidentityprovider.ConfirmDeviceInput, ...request.Option) (*cognitoidentityprovider.ConfirmDeviceOutput, error)
	ConfirmDeviceRequest(*cognitoidentityprovider.ConfirmDeviceInput) (*request.Request, *cognitoidentityprovider.ConfirmDeviceOutput)

	ConfirmForgotPassword(*cognitoidentityprovider.ConfirmForgotPasswordInput) (*cognitoidentityprovider.ConfirmForgotPasswordOutput, error)
	ConfirmForgotPasswordWithContext(aws.Context, *cognitoidentityprovider.ConfirmForgotPasswordInput, ...request.Option) (*cognitoidentityprovider.ConfirmForgotPasswordOutput, error)
	ConfirmForgotPasswordRequest(*cognitoidentityprovider.ConfirmForgotPasswordInput) (*request.Request, *cognitoidentityprovider.ConfirmForgotPasswordOutput)

	ConfirmSignUp(*cognitoidentityprovider.ConfirmSignUpInput) (*cognitoidentityprovider.ConfirmSignUpOutput, error)
	ConfirmSignUpWithContext(aws.Context, *cognitoidentityprovider.ConfirmSignUpInput, ...request.Option) (*cognitoidentityprovider.ConfirmSignUpOutput, error)
	ConfirmSignUpRequest(*cognitoidentityprovider.ConfirmSignUpInput) (*request.Request, *cognitoidentityprovider.ConfirmSignUpOutput)

	CreateGroup(*cognitoidentityprovider.CreateGroupInput) (*cognitoidentityprovider.CreateGroupOutput, error)
	CreateGroupWithContext(aws.Context, *cognitoidentityprovider.CreateGroupInput, ...request.Option) (*cognitoidentityprovider.CreateGroupOutput, error)
	CreateGroupRequest(*cognitoidentityprovider.CreateGroupInput) (*request.Request, *cognitoidentityprovider.CreateGroupOutput)

	CreateUserImportJob(*cognitoidentityprovider.CreateUserImportJobInput) (*cognitoidentityprovider.CreateUserImportJobOutput, error)
	CreateUserImportJobWithContext(aws.Context, *cognitoidentityprovider.CreateUserImportJobInput, ...request.Option) (*cognitoidentityprovider.CreateUserImportJobOutput, error)
	CreateUserImportJobRequest(*cognitoidentityprovider.CreateUserImportJobInput) (*request.Request, *cognitoidentityprovider.CreateUserImportJobOutput)

	CreateUserPool(*cognitoidentityprovider.CreateUserPoolInput) (*cognitoidentityprovider.CreateUserPoolOutput, error)
	CreateUserPoolWithContext(aws.Context, *cognitoidentityprovider.CreateUserPoolInput, ...request.Option) (*cognitoidentityprovider.CreateUserPoolOutput, error)
	CreateUserPoolRequest(*cognitoidentityprovider.CreateUserPoolInput) (*request.Request, *cognitoidentityprovider.CreateUserPoolOutput)

	CreateUserPoolClient(*cognitoidentityprovider.CreateUserPoolClientInput) (*cognitoidentityprovider.CreateUserPoolClientOutput, error)
	CreateUserPoolClientWithContext(aws.Context, *cognitoidentityprovider.CreateUserPoolClientInput, ...request.Option) (*cognitoidentityprovider.CreateUserPoolClientOutput, error)
	CreateUserPoolClientRequest(*cognitoidentityprovider.CreateUserPoolClientInput) (*request.Request, *cognitoidentityprovider.CreateUserPoolClientOutput)

	DeleteGroup(*cognitoidentityprovider.DeleteGroupInput) (*cognitoidentityprovider.DeleteGroupOutput, error)
	DeleteGroupWithContext(aws.Context, *cognitoidentityprovider.DeleteGroupInput, ...request.Option) (*cognitoidentityprovider.DeleteGroupOutput, error)
	DeleteGroupRequest(*cognitoidentityprovider.DeleteGroupInput) (*request.Request, *cognitoidentityprovider.DeleteGroupOutput)

	DeleteUser(*cognitoidentityprovider.DeleteUserInput) (*cognitoidentityprovider.DeleteUserOutput, error)
	DeleteUserWithContext(aws.Context, *cognitoidentityprovider.DeleteUserInput, ...request.Option) (*cognitoidentityprovider.DeleteUserOutput, error)
	DeleteUserRequest(*cognitoidentityprovider.DeleteUserInput) (*request.Request, *cognitoidentityprovider.DeleteUserOutput)

	DeleteUserAttributes(*cognitoidentityprovider.DeleteUserAttributesInput) (*cognitoidentityprovider.DeleteUserAttributesOutput, error)
	DeleteUserAttributesWithContext(aws.Context, *cognitoidentityprovider.DeleteUserAttributesInput, ...request.Option) (*cognitoidentityprovider.DeleteUserAttributesOutput, error)
	DeleteUserAttributesRequest(*cognitoidentityprovider.DeleteUserAttributesInput) (*request.Request, *cognitoidentityprovider.DeleteUserAttributesOutput)

	DeleteUserPool(*cognitoidentityprovider.DeleteUserPoolInput) (*cognitoidentityprovider.DeleteUserPoolOutput, error)
	DeleteUserPoolWithContext(aws.Context, *cognitoidentityprovider.DeleteUserPoolInput, ...request.Option) (*cognitoidentityprovider.DeleteUserPoolOutput, error)
	DeleteUserPoolRequest(*cognitoidentityprovider.DeleteUserPoolInput) (*request.Request, *cognitoidentityprovider.DeleteUserPoolOutput)

	DeleteUserPoolClient(*cognitoidentityprovider.DeleteUserPoolClientInput) (*cognitoidentityprovider.DeleteUserPoolClientOutput, error)
	DeleteUserPoolClientWithContext(aws.Context, *cognitoidentityprovider.DeleteUserPoolClientInput, ...request.Option) (*cognitoidentityprovider.DeleteUserPoolClientOutput, error)
	DeleteUserPoolClientRequest(*cognitoidentityprovider.DeleteUserPoolClientInput) (*request.Request, *cognitoidentityprovider.DeleteUserPoolClientOutput)

	DescribeUserImportJob(*cognitoidentityprovider.DescribeUserImportJobInput) (*cognitoidentityprovider.DescribeUserImportJobOutput, error)
	DescribeUserImportJobWithContext(aws.Context, *cognitoidentityprovider.DescribeUserImportJobInput, ...request.Option) (*cognitoidentityprovider.DescribeUserImportJobOutput, error)
	DescribeUserImportJobRequest(*cognitoidentityprovider.DescribeUserImportJobInput) (*request.Request, *cognitoidentityprovider.DescribeUserImportJobOutput)

	DescribeUserPool(*cognitoidentityprovider.DescribeUserPoolInput) (*cognitoidentityprovider.DescribeUserPoolOutput, error)
	DescribeUserPoolWithContext(aws.Context, *cognitoidentityprovider.DescribeUserPoolInput, ...request.Option) (*cognitoidentityprovider.DescribeUserPoolOutput, error)
	DescribeUserPoolRequest(*cognitoidentityprovider.DescribeUserPoolInput) (*request.Request, *cognitoidentityprovider.DescribeUserPoolOutput)

	DescribeUserPoolClient(*cognitoidentityprovider.DescribeUserPoolClientInput) (*cognitoidentityprovider.DescribeUserPoolClientOutput, error)
	DescribeUserPoolClientWithContext(aws.Context, *cognitoidentityprovider.DescribeUserPoolClientInput, ...request.Option) (*cognitoidentityprovider.DescribeUserPoolClientOutput, error)
	DescribeUserPoolClientRequest(*cognitoidentityprovider.DescribeUserPoolClientInput) (*request.Request, *cognitoidentityprovider.DescribeUserPoolClientOutput)

	ForgetDevice(*cognitoidentityprovider.ForgetDeviceInput) (*cognitoidentityprovider.ForgetDeviceOutput, error)
	ForgetDeviceWithContext(aws.Context, *cognitoidentityprovider.ForgetDeviceInput, ...request.Option) (*cognitoidentityprovider.ForgetDeviceOutput, error)
	ForgetDeviceRequest(*cognitoidentityprovider.ForgetDeviceInput) (*request.Request, *cognitoidentityprovider.ForgetDeviceOutput)

	ForgotPassword(*cognitoidentityprovider.ForgotPasswordInput) (*cognitoidentityprovider.ForgotPasswordOutput, error)
	ForgotPasswordWithContext(aws.Context, *cognitoidentityprovider.ForgotPasswordInput, ...request.Option) (*cognitoidentityprovider.ForgotPasswordOutput, error)
	ForgotPasswordRequest(*cognitoidentityprovider.ForgotPasswordInput) (*request.Request, *cognitoidentityprovider.ForgotPasswordOutput)

	GetCSVHeader(*cognitoidentityprovider.GetCSVHeaderInput) (*cognitoidentityprovider.GetCSVHeaderOutput, error)
	GetCSVHeaderWithContext(aws.Context, *cognitoidentityprovider.GetCSVHeaderInput, ...request.Option) (*cognitoidentityprovider.GetCSVHeaderOutput, error)
	GetCSVHeaderRequest(*cognitoidentityprovider.GetCSVHeaderInput) (*request.Request, *cognitoidentityprovider.GetCSVHeaderOutput)

	GetDevice(*cognitoidentityprovider.GetDeviceInput) (*cognitoidentityprovider.GetDeviceOutput, error)
	GetDeviceWithContext(aws.Context, *cognitoidentityprovider.GetDeviceInput, ...request.Option) (*cognitoidentityprovider.GetDeviceOutput, error)
	GetDeviceRequest(*cognitoidentityprovider.GetDeviceInput) (*request.Request, *cognitoidentityprovider.GetDeviceOutput)

	GetGroup(*cognitoidentityprovider.GetGroupInput) (*cognitoidentityprovider.GetGroupOutput, error)
	GetGroupWithContext(aws.Context, *cognitoidentityprovider.GetGroupInput, ...request.Option) (*cognitoidentityprovider.GetGroupOutput, error)
	GetGroupRequest(*cognitoidentityprovider.GetGroupInput) (*request.Request, *cognitoidentityprovider.GetGroupOutput)

	GetUser(*cognitoidentityprovider.GetUserInput) (*cognitoidentityprovider.GetUserOutput, error)
	GetUserWithContext(aws.Context, *cognitoidentityprovider.GetUserInput, ...request.Option) (*cognitoidentityprovider.GetUserOutput, error)
	GetUserRequest(*cognitoidentityprovider.GetUserInput) (*request.Request, *cognitoidentityprovider.GetUserOutput)

	GetUserAttributeVerificationCode(*cognitoidentityprovider.GetUserAttributeVerificationCodeInput) (*cognitoidentityprovider.GetUserAttributeVerificationCodeOutput, error)
	GetUserAttributeVerificationCodeWithContext(aws.Context, *cognitoidentityprovider.GetUserAttributeVerificationCodeInput, ...request.Option) (*cognitoidentityprovider.GetUserAttributeVerificationCodeOutput, error)
	GetUserAttributeVerificationCodeRequest(*cognitoidentityprovider.GetUserAttributeVerificationCodeInput) (*request.Request, *cognitoidentityprovider.GetUserAttributeVerificationCodeOutput)

	GlobalSignOut(*cognitoidentityprovider.GlobalSignOutInput) (*cognitoidentityprovider.GlobalSignOutOutput, error)
	GlobalSignOutWithContext(aws.Context, *cognitoidentityprovider.GlobalSignOutInput, ...request.Option) (*cognitoidentityprovider.GlobalSignOutOutput, error)
	GlobalSignOutRequest(*cognitoidentityprovider.GlobalSignOutInput) (*request.Request, *cognitoidentityprovider.GlobalSignOutOutput)

	InitiateAuth(*cognitoidentityprovider.InitiateAuthInput) (*cognitoidentityprovider.InitiateAuthOutput, error)
	InitiateAuthWithContext(aws.Context, *cognitoidentityprovider.InitiateAuthInput, ...request.Option) (*cognitoidentityprovider.InitiateAuthOutput, error)
	InitiateAuthRequest(*cognitoidentityprovider.InitiateAuthInput) (*request.Request, *cognitoidentityprovider.InitiateAuthOutput)

	ListDevices(*cognitoidentityprovider.ListDevicesInput) (*cognitoidentityprovider.ListDevicesOutput, error)
	ListDevicesWithContext(aws.Context, *cognitoidentityprovider.ListDevicesInput, ...request.Option) (*cognitoidentityprovider.ListDevicesOutput, error)
	ListDevicesRequest(*cognitoidentityprovider.ListDevicesInput) (*request.Request, *cognitoidentityprovider.ListDevicesOutput)

	ListGroups(*cognitoidentityprovider.ListGroupsInput) (*cognitoidentityprovider.ListGroupsOutput, error)
	ListGroupsWithContext(aws.Context, *cognitoidentityprovider.ListGroupsInput, ...request.Option) (*cognitoidentityprovider.ListGroupsOutput, error)
	ListGroupsRequest(*cognitoidentityprovider.ListGroupsInput) (*request.Request, *cognitoidentityprovider.ListGroupsOutput)

	ListUserImportJobs(*cognitoidentityprovider.ListUserImportJobsInput) (*cognitoidentityprovider.ListUserImportJobsOutput, error)
	ListUserImportJobsWithContext(aws.Context, *cognitoidentityprovider.ListUserImportJobsInput, ...request.Option) (*cognitoidentityprovider.ListUserImportJobsOutput, error)
	ListUserImportJobsRequest(*cognitoidentityprovider.ListUserImportJobsInput) (*request.Request, *cognitoidentityprovider.ListUserImportJobsOutput)

	ListUserPoolClients(*cognitoidentityprovider.ListUserPoolClientsInput) (*cognitoidentityprovider.ListUserPoolClientsOutput, error)
	ListUserPoolClientsWithContext(aws.Context, *cognitoidentityprovider.ListUserPoolClientsInput, ...request.Option) (*cognitoidentityprovider.ListUserPoolClientsOutput, error)
	ListUserPoolClientsRequest(*cognitoidentityprovider.ListUserPoolClientsInput) (*request.Request, *cognitoidentityprovider.ListUserPoolClientsOutput)

	ListUserPools(*cognitoidentityprovider.ListUserPoolsInput) (*cognitoidentityprovider.ListUserPoolsOutput, error)
	ListUserPoolsWithContext(aws.Context, *cognitoidentityprovider.ListUserPoolsInput, ...request.Option) (*cognitoidentityprovider.ListUserPoolsOutput, error)
	ListUserPoolsRequest(*cognitoidentityprovider.ListUserPoolsInput) (*request.Request, *cognitoidentityprovider.ListUserPoolsOutput)

	ListUsers(*cognitoidentityprovider.ListUsersInput) (*cognitoidentityprovider.ListUsersOutput, error)
	ListUsersWithContext(aws.Context, *cognitoidentityprovider.ListUsersInput, ...request.Option) (*cognitoidentityprovider.ListUsersOutput, error)
	ListUsersRequest(*cognitoidentityprovider.ListUsersInput) (*request.Request, *cognitoidentityprovider.ListUsersOutput)

	ListUsersInGroup(*cognitoidentityprovider.ListUsersInGroupInput) (*cognitoidentityprovider.ListUsersInGroupOutput, error)
	ListUsersInGroupWithContext(aws.Context, *cognitoidentityprovider.ListUsersInGroupInput, ...request.Option) (*cognitoidentityprovider.ListUsersInGroupOutput, error)
	ListUsersInGroupRequest(*cognitoidentityprovider.ListUsersInGroupInput) (*request.Request, *cognitoidentityprovider.ListUsersInGroupOutput)

	ResendConfirmationCode(*cognitoidentityprovider.ResendConfirmationCodeInput) (*cognitoidentityprovider.ResendConfirmationCodeOutput, error)
	ResendConfirmationCodeWithContext(aws.Context, *cognitoidentityprovider.ResendConfirmationCodeInput, ...request.Option) (*cognitoidentityprovider.ResendConfirmationCodeOutput, error)
	ResendConfirmationCodeRequest(*cognitoidentityprovider.ResendConfirmationCodeInput) (*request.Request, *cognitoidentityprovider.ResendConfirmationCodeOutput)

	RespondToAuthChallenge(*cognitoidentityprovider.RespondToAuthChallengeInput) (*cognitoidentityprovider.RespondToAuthChallengeOutput, error)
	RespondToAuthChallengeWithContext(aws.Context, *cognitoidentityprovider.RespondToAuthChallengeInput, ...request.Option) (*cognitoidentityprovider.RespondToAuthChallengeOutput, error)
	RespondToAuthChallengeRequest(*cognitoidentityprovider.RespondToAuthChallengeInput) (*request.Request, *cognitoidentityprovider.RespondToAuthChallengeOutput)

	SetUserSettings(*cognitoidentityprovider.SetUserSettingsInput) (*cognitoidentityprovider.SetUserSettingsOutput, error)
	SetUserSettingsWithContext(aws.Context, *cognitoidentityprovider.SetUserSettingsInput, ...request.Option) (*cognitoidentityprovider.SetUserSettingsOutput, error)
	SetUserSettingsRequest(*cognitoidentityprovider.SetUserSettingsInput) (*request.Request, *cognitoidentityprovider.SetUserSettingsOutput)

	SignUp(*cognitoidentityprovider.SignUpInput) (*cognitoidentityprovider.SignUpOutput, error)
	SignUpWithContext(aws.Context, *cognitoidentityprovider.SignUpInput, ...request.Option) (*cognitoidentityprovider.SignUpOutput, error)
	SignUpRequest(*cognitoidentityprovider.SignUpInput) (*request.Request, *cognitoidentityprovider.SignUpOutput)

	StartUserImportJob(*cognitoidentityprovider.StartUserImportJobInput) (*cognitoidentityprovider.StartUserImportJobOutput, error)
	StartUserImportJobWithContext(aws.Context, *cognitoidentityprovider.StartUserImportJobInput, ...request.Option) (*cognitoidentityprovider.StartUserImportJobOutput, error)
	StartUserImportJobRequest(*cognitoidentityprovider.StartUserImportJobInput) (*request.Request, *cognitoidentityprovider.StartUserImportJobOutput)

	StopUserImportJob(*cognitoidentityprovider.StopUserImportJobInput) (*cognitoidentityprovider.StopUserImportJobOutput, error)
	StopUserImportJobWithContext(aws.Context, *cognitoidentityprovider.StopUserImportJobInput, ...request.Option) (*cognitoidentityprovider.StopUserImportJobOutput, error)
	StopUserImportJobRequest(*cognitoidentityprovider.StopUserImportJobInput) (*request.Request, *cognitoidentityprovider.StopUserImportJobOutput)

	UpdateDeviceStatus(*cognitoidentityprovider.UpdateDeviceStatusInput) (*cognitoidentityprovider.UpdateDeviceStatusOutput, error)
	UpdateDeviceStatusWithContext(aws.Context, *cognitoidentityprovider.UpdateDeviceStatusInput, ...request.Option) (*cognitoidentityprovider.UpdateDeviceStatusOutput, error)
	UpdateDeviceStatusRequest(*cognitoidentityprovider.UpdateDeviceStatusInput) (*request.Request, *cognitoidentityprovider.UpdateDeviceStatusOutput)

	UpdateGroup(*cognitoidentityprovider.UpdateGroupInput) (*cognitoidentityprovider.UpdateGroupOutput, error)
	UpdateGroupWithContext(aws.Context, *cognitoidentityprovider.UpdateGroupInput, ...request.Option) (*cognitoidentityprovider.UpdateGroupOutput, error)
	UpdateGroupRequest(*cognitoidentityprovider.UpdateGroupInput) (*request.Request, *cognitoidentityprovider.UpdateGroupOutput)

	UpdateUserAttributes(*cognitoidentityprovider.UpdateUserAttributesInput) (*cognitoidentityprovider.UpdateUserAttributesOutput, error)
	UpdateUserAttributesWithContext(aws.Context, *cognitoidentityprovider.UpdateUserAttributesInput, ...request.Option) (*cognitoidentityprovider.UpdateUserAttributesOutput, error)
	UpdateUserAttributesRequest(*cognitoidentityprovider.UpdateUserAttributesInput) (*request.Request, *cognitoidentityprovider.UpdateUserAttributesOutput)

	UpdateUserPool(*cognitoidentityprovider.UpdateUserPoolInput) (*cognitoidentityprovider.UpdateUserPoolOutput, error)
	UpdateUserPoolWithContext(aws.Context, *cognitoidentityprovider.UpdateUserPoolInput, ...request.Option) (*cognitoidentityprovider.UpdateUserPoolOutput, error)
	UpdateUserPoolRequest(*cognitoidentityprovider.UpdateUserPoolInput) (*request.Request, *cognitoidentityprovider.UpdateUserPoolOutput)

	UpdateUserPoolClient(*cognitoidentityprovider.UpdateUserPoolClientInput) (*cognitoidentityprovider.UpdateUserPoolClientOutput, error)
	UpdateUserPoolClientWithContext(aws.Context, *cognitoidentityprovider.UpdateUserPoolClientInput, ...request.Option) (*cognitoidentityprovider.UpdateUserPoolClientOutput, error)
	UpdateUserPoolClientRequest(*cognitoidentityprovider.UpdateUserPoolClientInput) (*request.Request, *cognitoidentityprovider.UpdateUserPoolClientOutput)

	VerifyUserAttribute(*cognitoidentityprovider.VerifyUserAttributeInput) (*cognitoidentityprovider.VerifyUserAttributeOutput, error)
	VerifyUserAttributeWithContext(aws.Context, *cognitoidentityprovider.VerifyUserAttributeInput, ...request.Option) (*cognitoidentityprovider.VerifyUserAttributeOutput, error)
	VerifyUserAttributeRequest(*cognitoidentityprovider.VerifyUserAttributeInput) (*request.Request, *cognitoidentityprovider.VerifyUserAttributeOutput)
}

var _ CognitoIdentityProviderAPI = (*cognitoidentityprovider.CognitoIdentityProvider)(nil)