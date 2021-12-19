package repository

import (
	"github.com/lib/pq"
	"testing"
)

func TestRepo_Init(t *testing.T) {
	r := NewPostgresRepository()
	err := r.Init("", "", "", "", "", 5)
	if err == nil {
		t.Errorf("expected error actual:%v", err)
	}
}

func TestValidatePSQLError(t *testing.T) {
	psqlerr := &pq.Error{
		Code: "",
	}

	equal := ValidatePSQLError(psqlerr, UniqueViolation)
	if equal {
		t.Errorf("expected false actual %v", equal)
	}
}
