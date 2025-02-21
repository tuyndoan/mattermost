// Code generated by mockery v2.42.2. DO NOT EDIT.

// Regenerate this file using `make searchengine-mocks`.

package mocks

import (
	model "github.com/mattermost/mattermost/server/public/model"
	request "github.com/mattermost/mattermost/server/public/shared/request"
	mock "github.com/stretchr/testify/mock"

	time "time"
)

// SearchEngineInterface is an autogenerated mock type for the SearchEngineInterface type
type SearchEngineInterface struct {
	mock.Mock
}

// DataRetentionDeleteIndexes provides a mock function with given fields: rctx, cutoff
func (_m *SearchEngineInterface) DataRetentionDeleteIndexes(rctx request.CTX, cutoff time.Time) *model.AppError {
	ret := _m.Called(rctx, cutoff)

	if len(ret) == 0 {
		panic("no return value specified for DataRetentionDeleteIndexes")
	}

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func(request.CTX, time.Time) *model.AppError); ok {
		r0 = rf(rctx, cutoff)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// DeleteChannel provides a mock function with given fields: channel
func (_m *SearchEngineInterface) DeleteChannel(channel *model.Channel) *model.AppError {
	ret := _m.Called(channel)

	if len(ret) == 0 {
		panic("no return value specified for DeleteChannel")
	}

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func(*model.Channel) *model.AppError); ok {
		r0 = rf(channel)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// DeleteChannelPosts provides a mock function with given fields: rctx, channelID
func (_m *SearchEngineInterface) DeleteChannelPosts(rctx request.CTX, channelID string) *model.AppError {
	ret := _m.Called(rctx, channelID)

	if len(ret) == 0 {
		panic("no return value specified for DeleteChannelPosts")
	}

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func(request.CTX, string) *model.AppError); ok {
		r0 = rf(rctx, channelID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// DeleteFile provides a mock function with given fields: fileID
func (_m *SearchEngineInterface) DeleteFile(fileID string) *model.AppError {
	ret := _m.Called(fileID)

	if len(ret) == 0 {
		panic("no return value specified for DeleteFile")
	}

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func(string) *model.AppError); ok {
		r0 = rf(fileID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// DeleteFilesBatch provides a mock function with given fields: rctx, endTime, limit
func (_m *SearchEngineInterface) DeleteFilesBatch(rctx request.CTX, endTime int64, limit int64) *model.AppError {
	ret := _m.Called(rctx, endTime, limit)

	if len(ret) == 0 {
		panic("no return value specified for DeleteFilesBatch")
	}

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func(request.CTX, int64, int64) *model.AppError); ok {
		r0 = rf(rctx, endTime, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// DeletePost provides a mock function with given fields: post
func (_m *SearchEngineInterface) DeletePost(post *model.Post) *model.AppError {
	ret := _m.Called(post)

	if len(ret) == 0 {
		panic("no return value specified for DeletePost")
	}

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func(*model.Post) *model.AppError); ok {
		r0 = rf(post)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// DeletePostFiles provides a mock function with given fields: rctx, postID
func (_m *SearchEngineInterface) DeletePostFiles(rctx request.CTX, postID string) *model.AppError {
	ret := _m.Called(rctx, postID)

	if len(ret) == 0 {
		panic("no return value specified for DeletePostFiles")
	}

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func(request.CTX, string) *model.AppError); ok {
		r0 = rf(rctx, postID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// DeleteUser provides a mock function with given fields: user
func (_m *SearchEngineInterface) DeleteUser(user *model.User) *model.AppError {
	ret := _m.Called(user)

	if len(ret) == 0 {
		panic("no return value specified for DeleteUser")
	}

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func(*model.User) *model.AppError); ok {
		r0 = rf(user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// DeleteUserFiles provides a mock function with given fields: rctx, userID
func (_m *SearchEngineInterface) DeleteUserFiles(rctx request.CTX, userID string) *model.AppError {
	ret := _m.Called(rctx, userID)

	if len(ret) == 0 {
		panic("no return value specified for DeleteUserFiles")
	}

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func(request.CTX, string) *model.AppError); ok {
		r0 = rf(rctx, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// DeleteUserPosts provides a mock function with given fields: rctx, userID
func (_m *SearchEngineInterface) DeleteUserPosts(rctx request.CTX, userID string) *model.AppError {
	ret := _m.Called(rctx, userID)

	if len(ret) == 0 {
		panic("no return value specified for DeleteUserPosts")
	}

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func(request.CTX, string) *model.AppError); ok {
		r0 = rf(rctx, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// GetFullVersion provides a mock function with given fields:
func (_m *SearchEngineInterface) GetFullVersion() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetFullVersion")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetName provides a mock function with given fields:
func (_m *SearchEngineInterface) GetName() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetName")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetPlugins provides a mock function with given fields:
func (_m *SearchEngineInterface) GetPlugins() []string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetPlugins")
	}

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// GetVersion provides a mock function with given fields:
func (_m *SearchEngineInterface) GetVersion() int {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetVersion")
	}

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// IndexChannel provides a mock function with given fields: rctx, channel, userIDs, teamMemberIDs
func (_m *SearchEngineInterface) IndexChannel(rctx request.CTX, channel *model.Channel, userIDs []string, teamMemberIDs []string) *model.AppError {
	ret := _m.Called(rctx, channel, userIDs, teamMemberIDs)

	if len(ret) == 0 {
		panic("no return value specified for IndexChannel")
	}

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func(request.CTX, *model.Channel, []string, []string) *model.AppError); ok {
		r0 = rf(rctx, channel, userIDs, teamMemberIDs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// IndexFile provides a mock function with given fields: file, channelId
func (_m *SearchEngineInterface) IndexFile(file *model.FileInfo, channelId string) *model.AppError {
	ret := _m.Called(file, channelId)

	if len(ret) == 0 {
		panic("no return value specified for IndexFile")
	}

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func(*model.FileInfo, string) *model.AppError); ok {
		r0 = rf(file, channelId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// IndexPost provides a mock function with given fields: post, teamId
func (_m *SearchEngineInterface) IndexPost(post *model.Post, teamId string) *model.AppError {
	ret := _m.Called(post, teamId)

	if len(ret) == 0 {
		panic("no return value specified for IndexPost")
	}

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func(*model.Post, string) *model.AppError); ok {
		r0 = rf(post, teamId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// IndexUser provides a mock function with given fields: rctx, user, teamsIds, channelsIds
func (_m *SearchEngineInterface) IndexUser(rctx request.CTX, user *model.User, teamsIds []string, channelsIds []string) *model.AppError {
	ret := _m.Called(rctx, user, teamsIds, channelsIds)

	if len(ret) == 0 {
		panic("no return value specified for IndexUser")
	}

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func(request.CTX, *model.User, []string, []string) *model.AppError); ok {
		r0 = rf(rctx, user, teamsIds, channelsIds)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// IsActive provides a mock function with given fields:
func (_m *SearchEngineInterface) IsActive() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for IsActive")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// IsAutocompletionEnabled provides a mock function with given fields:
func (_m *SearchEngineInterface) IsAutocompletionEnabled() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for IsAutocompletionEnabled")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// IsEnabled provides a mock function with given fields:
func (_m *SearchEngineInterface) IsEnabled() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for IsEnabled")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// IsIndexingEnabled provides a mock function with given fields:
func (_m *SearchEngineInterface) IsIndexingEnabled() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for IsIndexingEnabled")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// IsIndexingSync provides a mock function with given fields:
func (_m *SearchEngineInterface) IsIndexingSync() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for IsIndexingSync")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// IsSearchEnabled provides a mock function with given fields:
func (_m *SearchEngineInterface) IsSearchEnabled() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for IsSearchEnabled")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// PurgeIndexList provides a mock function with given fields: rctx, indexes
func (_m *SearchEngineInterface) PurgeIndexList(rctx request.CTX, indexes []string) *model.AppError {
	ret := _m.Called(rctx, indexes)

	if len(ret) == 0 {
		panic("no return value specified for PurgeIndexList")
	}

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func(request.CTX, []string) *model.AppError); ok {
		r0 = rf(rctx, indexes)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// PurgeIndexes provides a mock function with given fields: rctx
func (_m *SearchEngineInterface) PurgeIndexes(rctx request.CTX) *model.AppError {
	ret := _m.Called(rctx)

	if len(ret) == 0 {
		panic("no return value specified for PurgeIndexes")
	}

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func(request.CTX) *model.AppError); ok {
		r0 = rf(rctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// RefreshIndexes provides a mock function with given fields: rctx
func (_m *SearchEngineInterface) RefreshIndexes(rctx request.CTX) *model.AppError {
	ret := _m.Called(rctx)

	if len(ret) == 0 {
		panic("no return value specified for RefreshIndexes")
	}

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func(request.CTX) *model.AppError); ok {
		r0 = rf(rctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// SearchChannels provides a mock function with given fields: teamId, userID, term, isGuest, includeDeleted
func (_m *SearchEngineInterface) SearchChannels(teamId string, userID string, term string, isGuest bool, includeDeleted bool) ([]string, *model.AppError) {
	ret := _m.Called(teamId, userID, term, isGuest, includeDeleted)

	if len(ret) == 0 {
		panic("no return value specified for SearchChannels")
	}

	var r0 []string
	var r1 *model.AppError
	if rf, ok := ret.Get(0).(func(string, string, string, bool, bool) ([]string, *model.AppError)); ok {
		return rf(teamId, userID, term, isGuest, includeDeleted)
	}
	if rf, ok := ret.Get(0).(func(string, string, string, bool, bool) []string); ok {
		r0 = rf(teamId, userID, term, isGuest, includeDeleted)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string, string, bool, bool) *model.AppError); ok {
		r1 = rf(teamId, userID, term, isGuest, includeDeleted)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// SearchFiles provides a mock function with given fields: channels, searchParams, page, perPage
func (_m *SearchEngineInterface) SearchFiles(channels model.ChannelList, searchParams []*model.SearchParams, page int, perPage int) ([]string, *model.AppError) {
	ret := _m.Called(channels, searchParams, page, perPage)

	if len(ret) == 0 {
		panic("no return value specified for SearchFiles")
	}

	var r0 []string
	var r1 *model.AppError
	if rf, ok := ret.Get(0).(func(model.ChannelList, []*model.SearchParams, int, int) ([]string, *model.AppError)); ok {
		return rf(channels, searchParams, page, perPage)
	}
	if rf, ok := ret.Get(0).(func(model.ChannelList, []*model.SearchParams, int, int) []string); ok {
		r0 = rf(channels, searchParams, page, perPage)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(model.ChannelList, []*model.SearchParams, int, int) *model.AppError); ok {
		r1 = rf(channels, searchParams, page, perPage)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// SearchPosts provides a mock function with given fields: channels, searchParams, page, perPage
func (_m *SearchEngineInterface) SearchPosts(channels model.ChannelList, searchParams []*model.SearchParams, page int, perPage int) ([]string, model.PostSearchMatches, *model.AppError) {
	ret := _m.Called(channels, searchParams, page, perPage)

	if len(ret) == 0 {
		panic("no return value specified for SearchPosts")
	}

	var r0 []string
	var r1 model.PostSearchMatches
	var r2 *model.AppError
	if rf, ok := ret.Get(0).(func(model.ChannelList, []*model.SearchParams, int, int) ([]string, model.PostSearchMatches, *model.AppError)); ok {
		return rf(channels, searchParams, page, perPage)
	}
	if rf, ok := ret.Get(0).(func(model.ChannelList, []*model.SearchParams, int, int) []string); ok {
		r0 = rf(channels, searchParams, page, perPage)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(model.ChannelList, []*model.SearchParams, int, int) model.PostSearchMatches); ok {
		r1 = rf(channels, searchParams, page, perPage)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(model.PostSearchMatches)
		}
	}

	if rf, ok := ret.Get(2).(func(model.ChannelList, []*model.SearchParams, int, int) *model.AppError); ok {
		r2 = rf(channels, searchParams, page, perPage)
	} else {
		if ret.Get(2) != nil {
			r2 = ret.Get(2).(*model.AppError)
		}
	}

	return r0, r1, r2
}

// SearchUsersInChannel provides a mock function with given fields: teamId, channelId, restrictedToChannels, term, options
func (_m *SearchEngineInterface) SearchUsersInChannel(teamId string, channelId string, restrictedToChannels []string, term string, options *model.UserSearchOptions) ([]string, []string, *model.AppError) {
	ret := _m.Called(teamId, channelId, restrictedToChannels, term, options)

	if len(ret) == 0 {
		panic("no return value specified for SearchUsersInChannel")
	}

	var r0 []string
	var r1 []string
	var r2 *model.AppError
	if rf, ok := ret.Get(0).(func(string, string, []string, string, *model.UserSearchOptions) ([]string, []string, *model.AppError)); ok {
		return rf(teamId, channelId, restrictedToChannels, term, options)
	}
	if rf, ok := ret.Get(0).(func(string, string, []string, string, *model.UserSearchOptions) []string); ok {
		r0 = rf(teamId, channelId, restrictedToChannels, term, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string, []string, string, *model.UserSearchOptions) []string); ok {
		r1 = rf(teamId, channelId, restrictedToChannels, term, options)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]string)
		}
	}

	if rf, ok := ret.Get(2).(func(string, string, []string, string, *model.UserSearchOptions) *model.AppError); ok {
		r2 = rf(teamId, channelId, restrictedToChannels, term, options)
	} else {
		if ret.Get(2) != nil {
			r2 = ret.Get(2).(*model.AppError)
		}
	}

	return r0, r1, r2
}

// SearchUsersInTeam provides a mock function with given fields: teamId, restrictedToChannels, term, options
func (_m *SearchEngineInterface) SearchUsersInTeam(teamId string, restrictedToChannels []string, term string, options *model.UserSearchOptions) ([]string, *model.AppError) {
	ret := _m.Called(teamId, restrictedToChannels, term, options)

	if len(ret) == 0 {
		panic("no return value specified for SearchUsersInTeam")
	}

	var r0 []string
	var r1 *model.AppError
	if rf, ok := ret.Get(0).(func(string, []string, string, *model.UserSearchOptions) ([]string, *model.AppError)); ok {
		return rf(teamId, restrictedToChannels, term, options)
	}
	if rf, ok := ret.Get(0).(func(string, []string, string, *model.UserSearchOptions) []string); ok {
		r0 = rf(teamId, restrictedToChannels, term, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(string, []string, string, *model.UserSearchOptions) *model.AppError); ok {
		r1 = rf(teamId, restrictedToChannels, term, options)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// Start provides a mock function with given fields:
func (_m *SearchEngineInterface) Start() *model.AppError {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Start")
	}

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func() *model.AppError); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// Stop provides a mock function with given fields:
func (_m *SearchEngineInterface) Stop() *model.AppError {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Stop")
	}

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func() *model.AppError); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// TestConfig provides a mock function with given fields: rctx, cfg
func (_m *SearchEngineInterface) TestConfig(rctx request.CTX, cfg *model.Config) *model.AppError {
	ret := _m.Called(rctx, cfg)

	if len(ret) == 0 {
		panic("no return value specified for TestConfig")
	}

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func(request.CTX, *model.Config) *model.AppError); ok {
		r0 = rf(rctx, cfg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// UpdateConfig provides a mock function with given fields: cfg
func (_m *SearchEngineInterface) UpdateConfig(cfg *model.Config) {
	_m.Called(cfg)
}

// NewSearchEngineInterface creates a new instance of SearchEngineInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSearchEngineInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *SearchEngineInterface {
	mock := &SearchEngineInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
