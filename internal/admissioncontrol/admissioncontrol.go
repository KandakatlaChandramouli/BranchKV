package admissioncontrol

import (
	"sync/atomic"
)

type AdmissionControl struct {
	admitted atomic.Uint64
	rejected atomic.Uint64
}

func NewAdmissionControl() *AdmissionControl {

	return &AdmissionControl{}
}

func (a *AdmissionControl) Admit(
	ok bool,
) {

	if ok {
		a.admitted.Add(1)
		return
	}

	a.rejected.Add(1)
}

func (a *AdmissionControl) Admitted() uint64 {
	return a.admitted.Load()
}

func (a *AdmissionControl) Rejected() uint64 {
	return a.rejected.Load()
}
