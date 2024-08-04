// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/runatlantis/atlantis/server/core/terraform (interfaces: Downloader)

package mocks

import (
	go_version "github.com/hashicorp/go-version"
	pegomock "github.com/petergtz/pegomock/v4"
	"reflect"
	"time"
)

type MockDownloader struct {
	fail func(message string, callerSkip ...int)
}

func NewMockDownloader(options ...pegomock.Option) *MockDownloader {
	mock := &MockDownloader{}
	for _, option := range options {
		option.Apply(mock)
	}
	return mock
}

func (mock *MockDownloader) SetFailHandler(fh pegomock.FailHandler) { mock.fail = fh }
func (mock *MockDownloader) FailHandler() pegomock.FailHandler      { return mock.fail }

func (mock *MockDownloader) Install(dir string, downloadURL string, v *go_version.Version) (string, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockDownloader().")
	}
	_params := []pegomock.Param{dir, downloadURL, v}
	_result := pegomock.GetGenericMockFrom(mock).Invoke("Install", _params, []reflect.Type{reflect.TypeOf((*string)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var _ret0 string
	var _ret1 error
	if len(_result) != 0 {
		if _result[0] != nil {
			_ret0 = _result[0].(string)
		}
		if _result[1] != nil {
			_ret1 = _result[1].(error)
		}
	}
	return _ret0, _ret1
}

func (mock *MockDownloader) VerifyWasCalledOnce() *VerifierMockDownloader {
	return &VerifierMockDownloader{
		mock:                   mock,
		invocationCountMatcher: pegomock.Times(1),
	}
}

func (mock *MockDownloader) VerifyWasCalled(invocationCountMatcher pegomock.InvocationCountMatcher) *VerifierMockDownloader {
	return &VerifierMockDownloader{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
	}
}

func (mock *MockDownloader) VerifyWasCalledInOrder(invocationCountMatcher pegomock.InvocationCountMatcher, inOrderContext *pegomock.InOrderContext) *VerifierMockDownloader {
	return &VerifierMockDownloader{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		inOrderContext:         inOrderContext,
	}
}

func (mock *MockDownloader) VerifyWasCalledEventually(invocationCountMatcher pegomock.InvocationCountMatcher, timeout time.Duration) *VerifierMockDownloader {
	return &VerifierMockDownloader{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		timeout:                timeout,
	}
}

type VerifierMockDownloader struct {
	mock                   *MockDownloader
	invocationCountMatcher pegomock.InvocationCountMatcher
	inOrderContext         *pegomock.InOrderContext
	timeout                time.Duration
}

func (verifier *VerifierMockDownloader) Install(dir string, downloadURL string, v *go_version.Version) *MockDownloader_Install_OngoingVerification {
	_params := []pegomock.Param{dir, downloadURL, v}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Install", _params, verifier.timeout)
	return &MockDownloader_Install_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockDownloader_Install_OngoingVerification struct {
	mock              *MockDownloader
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockDownloader_Install_OngoingVerification) GetCapturedArguments() (string, string, *go_version.Version) {
	dir, downloadURL, v := c.GetAllCapturedArguments()
	return dir[len(dir)-1], downloadURL[len(downloadURL)-1], v[len(v)-1]
}

func (c *MockDownloader_Install_OngoingVerification) GetAllCapturedArguments() (_param0 []string, _param1 []string, _param2 []*go_version.Version) {
	_params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(_params) > 0 {
		if len(_params) > 0 {
			_param0 = make([]string, len(c.methodInvocations))
			for u, param := range _params[0] {
				_param0[u] = param.(string)
			}
		}
		if len(_params) > 1 {
			_param1 = make([]string, len(c.methodInvocations))
			for u, param := range _params[1] {
				_param1[u] = param.(string)
			}
		}
		if len(_params) > 2 {
			_param2 = make([]*go_version.Version, len(c.methodInvocations))
			for u, param := range _params[2] {
				_param2[u] = param.(*go_version.Version)
			}
		}
	}
	return
}
