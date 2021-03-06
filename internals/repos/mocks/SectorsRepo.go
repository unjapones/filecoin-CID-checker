// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	bsontypes "github.com/protofire/filecoin-CID-checker/internals/bsontypes"
	mock "github.com/stretchr/testify/mock"
)

// SectorsRepo is an autogenerated mock type for the SectorsRepo type
type SectorsRepo struct {
	mock.Mock
}

// BulkWrite provides a mock function with given fields: sectors
func (_m *SectorsRepo) BulkWrite(sectors []*bsontypes.SectorInfo) error {
	ret := _m.Called(sectors)

	var r0 error
	if rf, ok := ret.Get(0).(func([]*bsontypes.SectorInfo) error); ok {
		r0 = rf(sectors)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BulkWriteInfo provides a mock function with given fields: sectors
func (_m *SectorsRepo) BulkWriteInfo(sectors []*bsontypes.SectorInfo) error {
	ret := _m.Called(sectors)

	var r0 error
	if rf, ok := ret.Get(0).(func([]*bsontypes.SectorInfo) error); ok {
		r0 = rf(sectors)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateIndexes provides a mock function with given fields:
func (_m *SectorsRepo) CreateIndexes() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetSector provides a mock function with given fields: sectorID
func (_m *SectorsRepo) GetSector(sectorID uint64) (bsontypes.SectorInfo, error) {
	ret := _m.Called(sectorID)

	var r0 bsontypes.SectorInfo
	if rf, ok := ret.Get(0).(func(uint64) bsontypes.SectorInfo); ok {
		r0 = rf(sectorID)
	} else {
		r0 = ret.Get(0).(bsontypes.SectorInfo)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(sectorID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SectorWithDeal provides a mock function with given fields: dealID
func (_m *SectorsRepo) SectorWithDeal(dealID uint64) (*bsontypes.SectorInfo, error) {
	ret := _m.Called(dealID)

	var r0 *bsontypes.SectorInfo
	if rf, ok := ret.Get(0).(func(uint64) *bsontypes.SectorInfo); ok {
		r0 = rf(dealID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bsontypes.SectorInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(dealID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetFaultSectors provides a mock function with given fields: sectors
func (_m *SectorsRepo) SetFaultSectors(sectors []uint64) error {
	ret := _m.Called(sectors)

	var r0 error
	if rf, ok := ret.Get(0).(func([]uint64) error); ok {
		r0 = rf(sectors)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetRecoveriesSectors provides a mock function with given fields: sectors
func (_m *SectorsRepo) SetRecoveriesSectors(sectors []uint64) error {
	ret := _m.Called(sectors)

	var r0 error
	if rf, ok := ret.Get(0).(func([]uint64) error); ok {
		r0 = rf(sectors)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
