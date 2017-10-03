// autogenerated: do not edit!
// generated from gentemplate [gentemplate -d Package=elib -id freeElts -d VecType=freeEltsVec -d Type=freeEltVec vec.tmpl]

// Copyright 2016 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package elib

type freeEltsVec []freeEltVec

func (p *freeEltsVec) Resize(n uint) {
	old_cap := uint(cap(*p))
	new_len := uint(len(*p)) + n
	if new_len > old_cap {
		new_cap := NextResizeCap(new_len)
		q := make([]freeEltVec, new_len, new_cap)
		copy(q, *p)
		*p = q
	}
	*p = (*p)[:new_len]
}

func (p *freeEltsVec) validate(new_len uint, zero freeEltVec) *freeEltVec {
	old_cap := uint(cap(*p))
	old_len := uint(len(*p))
	if new_len <= old_cap {
		// Need to reslice to larger length?
		if new_len > old_len {
			*p = (*p)[:new_len]
			for i := old_len; i < new_len; i++ {
				(*p)[i] = zero
			}
		}
		return &(*p)[new_len-1]
	}
	return p.validateSlowPath(zero, old_cap, new_len, old_len)
}

func (p *freeEltsVec) validateSlowPath(zero freeEltVec, old_cap, new_len, old_len uint) *freeEltVec {
	if new_len > old_cap {
		new_cap := NextResizeCap(new_len)
		q := make([]freeEltVec, new_cap, new_cap)
		copy(q, *p)
		for i := old_len; i < new_cap; i++ {
			q[i] = zero
		}
		*p = q[:new_len]
	}
	if new_len > old_len {
		*p = (*p)[:new_len]
	}
	return &(*p)[new_len-1]
}

func (p *freeEltsVec) Validate(i uint) *freeEltVec {
	var zero freeEltVec
	return p.validate(i+1, zero)
}

func (p *freeEltsVec) ValidateInit(i uint, zero freeEltVec) *freeEltVec {
	return p.validate(i+1, zero)
}

func (p *freeEltsVec) ValidateLen(l uint) (v *freeEltVec) {
	if l > 0 {
		var zero freeEltVec
		v = p.validate(l, zero)
	}
	return
}

func (p *freeEltsVec) ValidateLenInit(l uint, zero freeEltVec) (v *freeEltVec) {
	if l > 0 {
		v = p.validate(l, zero)
	}
	return
}

func (p *freeEltsVec) ResetLen() {
	if *p != nil {
		*p = (*p)[:0]
	}
}

func (p freeEltsVec) Len() uint { return uint(len(p)) }
