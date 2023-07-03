package resource

import (
	"testing"

	"github.com/tanveerprottoy/basic-go-server/pkg/data/sqlxpkg"
)

type TestService struct {
	service *Service
}

func NewTestService() *Service {
	dbClient := sqlxpkg.GetInstance()
	s := new(Service)
	s.repository = NewRepository(dbClient.DB)
	return s
}

func TestAll(t *testing.T) {
	// test service
	s := NewTestService()
	e, err := s.readOneInternal("1")
	if err != nil {
		t.Fatalf("DoSomething() returned error: %s", err)
	}
	if &e == nil {
		t.Logf("expected val, got nil: %d", &e)
		return
	}
	t.Logf("val returned, val: %d", e)
}
