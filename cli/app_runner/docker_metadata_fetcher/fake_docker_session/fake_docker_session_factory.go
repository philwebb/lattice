// This file was generated by counterfeiter
package fake_docker_session

import (
	"sync"

	"github.com/cloudfoundry-incubator/lattice/cli/app_runner/docker_metadata_fetcher"
)

type FakeDockerSessionFactory struct {
	MakeSessionStub        func(repoName string) (docker_metadata_fetcher.DockerSession, error)
	makeSessionMutex       sync.RWMutex
	makeSessionArgsForCall []struct {
		repoName string
	}
	makeSessionReturns struct {
		result1 docker_metadata_fetcher.DockerSession
		result2 error
	}
}

func (fake *FakeDockerSessionFactory) MakeSession(repoName string) (docker_metadata_fetcher.DockerSession, error) {
	fake.makeSessionMutex.Lock()
	fake.makeSessionArgsForCall = append(fake.makeSessionArgsForCall, struct {
		repoName string
	}{repoName})
	fake.makeSessionMutex.Unlock()
	if fake.MakeSessionStub != nil {
		return fake.MakeSessionStub(repoName)
	} else {
		return fake.makeSessionReturns.result1, fake.makeSessionReturns.result2
	}
}

func (fake *FakeDockerSessionFactory) MakeSessionCallCount() int {
	fake.makeSessionMutex.RLock()
	defer fake.makeSessionMutex.RUnlock()
	return len(fake.makeSessionArgsForCall)
}

func (fake *FakeDockerSessionFactory) MakeSessionArgsForCall(i int) string {
	fake.makeSessionMutex.RLock()
	defer fake.makeSessionMutex.RUnlock()
	return fake.makeSessionArgsForCall[i].repoName
}

func (fake *FakeDockerSessionFactory) MakeSessionReturns(result1 docker_metadata_fetcher.DockerSession, result2 error) {
	fake.MakeSessionStub = nil
	fake.makeSessionReturns = struct {
		result1 docker_metadata_fetcher.DockerSession
		result2 error
	}{result1, result2}
}

var _ docker_metadata_fetcher.DockerSessionFactory = new(FakeDockerSessionFactory)