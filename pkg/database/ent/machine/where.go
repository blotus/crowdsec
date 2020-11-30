// Code generated by entc, DO NOT EDIT.

package machine

import (
	"time"

	"github.com/crowdsecurity/crowdsec/pkg/database/ent/predicate"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// MachineId applies equality check predicate on the "machineId" field. It's identical to MachineIdEQ.
func MachineId(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMachineId), v))
	})
}

// Password applies equality check predicate on the "password" field. It's identical to PasswordEQ.
func Password(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPassword), v))
	})
}

// IpAddress applies equality check predicate on the "ipAddress" field. It's identical to IpAddressEQ.
func IpAddress(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIpAddress), v))
	})
}

// Scenarios applies equality check predicate on the "scenarios" field. It's identical to ScenariosEQ.
func Scenarios(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldScenarios), v))
	})
}

// Version applies equality check predicate on the "version" field. It's identical to VersionEQ.
func Version(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldVersion), v))
	})
}

// IsValidated applies equality check predicate on the "isValidated" field. It's identical to IsValidatedEQ.
func IsValidated(v bool) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsValidated), v))
	})
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Machine {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Machine(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Machine {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Machine(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Machine {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Machine(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Machine {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Machine(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// MachineIdEQ applies the EQ predicate on the "machineId" field.
func MachineIdEQ(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMachineId), v))
	})
}

// MachineIdNEQ applies the NEQ predicate on the "machineId" field.
func MachineIdNEQ(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMachineId), v))
	})
}

// MachineIdIn applies the In predicate on the "machineId" field.
func MachineIdIn(vs ...string) predicate.Machine {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Machine(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldMachineId), v...))
	})
}

// MachineIdNotIn applies the NotIn predicate on the "machineId" field.
func MachineIdNotIn(vs ...string) predicate.Machine {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Machine(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldMachineId), v...))
	})
}

// MachineIdGT applies the GT predicate on the "machineId" field.
func MachineIdGT(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMachineId), v))
	})
}

// MachineIdGTE applies the GTE predicate on the "machineId" field.
func MachineIdGTE(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMachineId), v))
	})
}

// MachineIdLT applies the LT predicate on the "machineId" field.
func MachineIdLT(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMachineId), v))
	})
}

// MachineIdLTE applies the LTE predicate on the "machineId" field.
func MachineIdLTE(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMachineId), v))
	})
}

// MachineIdContains applies the Contains predicate on the "machineId" field.
func MachineIdContains(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldMachineId), v))
	})
}

// MachineIdHasPrefix applies the HasPrefix predicate on the "machineId" field.
func MachineIdHasPrefix(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldMachineId), v))
	})
}

// MachineIdHasSuffix applies the HasSuffix predicate on the "machineId" field.
func MachineIdHasSuffix(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldMachineId), v))
	})
}

// MachineIdEqualFold applies the EqualFold predicate on the "machineId" field.
func MachineIdEqualFold(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldMachineId), v))
	})
}

// MachineIdContainsFold applies the ContainsFold predicate on the "machineId" field.
func MachineIdContainsFold(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldMachineId), v))
	})
}

// PasswordEQ applies the EQ predicate on the "password" field.
func PasswordEQ(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPassword), v))
	})
}

// PasswordNEQ applies the NEQ predicate on the "password" field.
func PasswordNEQ(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPassword), v))
	})
}

// PasswordIn applies the In predicate on the "password" field.
func PasswordIn(vs ...string) predicate.Machine {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Machine(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPassword), v...))
	})
}

// PasswordNotIn applies the NotIn predicate on the "password" field.
func PasswordNotIn(vs ...string) predicate.Machine {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Machine(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPassword), v...))
	})
}

// PasswordGT applies the GT predicate on the "password" field.
func PasswordGT(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPassword), v))
	})
}

// PasswordGTE applies the GTE predicate on the "password" field.
func PasswordGTE(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPassword), v))
	})
}

// PasswordLT applies the LT predicate on the "password" field.
func PasswordLT(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPassword), v))
	})
}

// PasswordLTE applies the LTE predicate on the "password" field.
func PasswordLTE(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPassword), v))
	})
}

// PasswordContains applies the Contains predicate on the "password" field.
func PasswordContains(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPassword), v))
	})
}

// PasswordHasPrefix applies the HasPrefix predicate on the "password" field.
func PasswordHasPrefix(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPassword), v))
	})
}

// PasswordHasSuffix applies the HasSuffix predicate on the "password" field.
func PasswordHasSuffix(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPassword), v))
	})
}

// PasswordEqualFold applies the EqualFold predicate on the "password" field.
func PasswordEqualFold(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPassword), v))
	})
}

// PasswordContainsFold applies the ContainsFold predicate on the "password" field.
func PasswordContainsFold(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPassword), v))
	})
}

// IpAddressEQ applies the EQ predicate on the "ipAddress" field.
func IpAddressEQ(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIpAddress), v))
	})
}

// IpAddressNEQ applies the NEQ predicate on the "ipAddress" field.
func IpAddressNEQ(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIpAddress), v))
	})
}

// IpAddressIn applies the In predicate on the "ipAddress" field.
func IpAddressIn(vs ...string) predicate.Machine {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Machine(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldIpAddress), v...))
	})
}

// IpAddressNotIn applies the NotIn predicate on the "ipAddress" field.
func IpAddressNotIn(vs ...string) predicate.Machine {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Machine(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldIpAddress), v...))
	})
}

// IpAddressGT applies the GT predicate on the "ipAddress" field.
func IpAddressGT(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldIpAddress), v))
	})
}

// IpAddressGTE applies the GTE predicate on the "ipAddress" field.
func IpAddressGTE(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldIpAddress), v))
	})
}

// IpAddressLT applies the LT predicate on the "ipAddress" field.
func IpAddressLT(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldIpAddress), v))
	})
}

// IpAddressLTE applies the LTE predicate on the "ipAddress" field.
func IpAddressLTE(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldIpAddress), v))
	})
}

// IpAddressContains applies the Contains predicate on the "ipAddress" field.
func IpAddressContains(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldIpAddress), v))
	})
}

// IpAddressHasPrefix applies the HasPrefix predicate on the "ipAddress" field.
func IpAddressHasPrefix(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldIpAddress), v))
	})
}

// IpAddressHasSuffix applies the HasSuffix predicate on the "ipAddress" field.
func IpAddressHasSuffix(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldIpAddress), v))
	})
}

// IpAddressEqualFold applies the EqualFold predicate on the "ipAddress" field.
func IpAddressEqualFold(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldIpAddress), v))
	})
}

// IpAddressContainsFold applies the ContainsFold predicate on the "ipAddress" field.
func IpAddressContainsFold(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldIpAddress), v))
	})
}

// ScenariosEQ applies the EQ predicate on the "scenarios" field.
func ScenariosEQ(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldScenarios), v))
	})
}

// ScenariosNEQ applies the NEQ predicate on the "scenarios" field.
func ScenariosNEQ(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldScenarios), v))
	})
}

// ScenariosIn applies the In predicate on the "scenarios" field.
func ScenariosIn(vs ...string) predicate.Machine {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Machine(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldScenarios), v...))
	})
}

// ScenariosNotIn applies the NotIn predicate on the "scenarios" field.
func ScenariosNotIn(vs ...string) predicate.Machine {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Machine(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldScenarios), v...))
	})
}

// ScenariosGT applies the GT predicate on the "scenarios" field.
func ScenariosGT(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldScenarios), v))
	})
}

// ScenariosGTE applies the GTE predicate on the "scenarios" field.
func ScenariosGTE(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldScenarios), v))
	})
}

// ScenariosLT applies the LT predicate on the "scenarios" field.
func ScenariosLT(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldScenarios), v))
	})
}

// ScenariosLTE applies the LTE predicate on the "scenarios" field.
func ScenariosLTE(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldScenarios), v))
	})
}

// ScenariosContains applies the Contains predicate on the "scenarios" field.
func ScenariosContains(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldScenarios), v))
	})
}

// ScenariosHasPrefix applies the HasPrefix predicate on the "scenarios" field.
func ScenariosHasPrefix(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldScenarios), v))
	})
}

// ScenariosHasSuffix applies the HasSuffix predicate on the "scenarios" field.
func ScenariosHasSuffix(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldScenarios), v))
	})
}

// ScenariosIsNil applies the IsNil predicate on the "scenarios" field.
func ScenariosIsNil() predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldScenarios)))
	})
}

// ScenariosNotNil applies the NotNil predicate on the "scenarios" field.
func ScenariosNotNil() predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldScenarios)))
	})
}

// ScenariosEqualFold applies the EqualFold predicate on the "scenarios" field.
func ScenariosEqualFold(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldScenarios), v))
	})
}

// ScenariosContainsFold applies the ContainsFold predicate on the "scenarios" field.
func ScenariosContainsFold(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldScenarios), v))
	})
}

// VersionEQ applies the EQ predicate on the "version" field.
func VersionEQ(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldVersion), v))
	})
}

// VersionNEQ applies the NEQ predicate on the "version" field.
func VersionNEQ(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldVersion), v))
	})
}

// VersionIn applies the In predicate on the "version" field.
func VersionIn(vs ...string) predicate.Machine {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Machine(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldVersion), v...))
	})
}

// VersionNotIn applies the NotIn predicate on the "version" field.
func VersionNotIn(vs ...string) predicate.Machine {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Machine(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldVersion), v...))
	})
}

// VersionGT applies the GT predicate on the "version" field.
func VersionGT(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldVersion), v))
	})
}

// VersionGTE applies the GTE predicate on the "version" field.
func VersionGTE(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldVersion), v))
	})
}

// VersionLT applies the LT predicate on the "version" field.
func VersionLT(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldVersion), v))
	})
}

// VersionLTE applies the LTE predicate on the "version" field.
func VersionLTE(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldVersion), v))
	})
}

// VersionContains applies the Contains predicate on the "version" field.
func VersionContains(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldVersion), v))
	})
}

// VersionHasPrefix applies the HasPrefix predicate on the "version" field.
func VersionHasPrefix(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldVersion), v))
	})
}

// VersionHasSuffix applies the HasSuffix predicate on the "version" field.
func VersionHasSuffix(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldVersion), v))
	})
}

// VersionIsNil applies the IsNil predicate on the "version" field.
func VersionIsNil() predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldVersion)))
	})
}

// VersionNotNil applies the NotNil predicate on the "version" field.
func VersionNotNil() predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldVersion)))
	})
}

// VersionEqualFold applies the EqualFold predicate on the "version" field.
func VersionEqualFold(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldVersion), v))
	})
}

// VersionContainsFold applies the ContainsFold predicate on the "version" field.
func VersionContainsFold(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldVersion), v))
	})
}

// IsValidatedEQ applies the EQ predicate on the "isValidated" field.
func IsValidatedEQ(v bool) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsValidated), v))
	})
}

// IsValidatedNEQ applies the NEQ predicate on the "isValidated" field.
func IsValidatedNEQ(v bool) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIsValidated), v))
	})
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStatus), v))
	})
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...string) predicate.Machine {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Machine(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStatus), v...))
	})
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...string) predicate.Machine {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Machine(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStatus), v...))
	})
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStatus), v))
	})
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStatus), v))
	})
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStatus), v))
	})
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStatus), v))
	})
}

// StatusContains applies the Contains predicate on the "status" field.
func StatusContains(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldStatus), v))
	})
}

// StatusHasPrefix applies the HasPrefix predicate on the "status" field.
func StatusHasPrefix(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldStatus), v))
	})
}

// StatusHasSuffix applies the HasSuffix predicate on the "status" field.
func StatusHasSuffix(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldStatus), v))
	})
}

// StatusIsNil applies the IsNil predicate on the "status" field.
func StatusIsNil() predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldStatus)))
	})
}

// StatusNotNil applies the NotNil predicate on the "status" field.
func StatusNotNil() predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldStatus)))
	})
}

// StatusEqualFold applies the EqualFold predicate on the "status" field.
func StatusEqualFold(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldStatus), v))
	})
}

// StatusContainsFold applies the ContainsFold predicate on the "status" field.
func StatusContainsFold(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldStatus), v))
	})
}

// HasAlerts applies the HasEdge predicate on the "alerts" edge.
func HasAlerts() predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AlertsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, AlertsTable, AlertsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAlertsWith applies the HasEdge predicate on the "alerts" edge with a given conditions (other predicates).
func HasAlertsWith(preds ...predicate.Alert) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AlertsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, AlertsTable, AlertsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Machine) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Machine) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Machine) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		p(s.Not())
	})
}