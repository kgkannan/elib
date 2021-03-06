// autogenerated: do not edit!
// generated from gentemplate [gentemplate -d Package=elib -id Uint32 -d VecType=Uint32Vec -d Type=uint32 vec.tmpl]

// Copyright 2016 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package elib

type Uint32Vec []uint32

func (p *Uint32Vec) Resize(n uint) {
	old_cap := uint(cap(*p))
	new_len := uint(len(*p)) + n
	if new_len > old_cap {
		new_cap := NextResizeCap(new_len)
		q := make([]uint32, new_len, new_cap)
		copy(q, *p)
		*p = q
	}
	*p = (*p)[:new_len]
}

func (p *Uint32Vec) validate(new_len uint, zero uint32) *uint32 {
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

func (p *Uint32Vec) validateSlowPath(zero uint32, old_cap, new_len, old_len uint) *uint32 {
	if new_len > old_cap {
		new_cap := NextResizeCap(new_len)
		q := make([]uint32, new_cap, new_cap)
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

func (p *Uint32Vec) Validate(i uint) *uint32 {
	var zero uint32
	return p.validate(i+1, zero)
}

func (p *Uint32Vec) ValidateInit(i uint, zero uint32) *uint32 {
	return p.validate(i+1, zero)
}

func (p *Uint32Vec) ValidateLen(l uint) (v *uint32) {
	if l > 0 {
		var zero uint32
		v = p.validate(l, zero)
	}
	return
}

func (p *Uint32Vec) ValidateLenInit(l uint, zero uint32) (v *uint32) {
	if l > 0 {
		v = p.validate(l, zero)
	}
	return
}

func (p *Uint32Vec) ResetLen() {
	if *p != nil {
		*p = (*p)[:0]
	}
}

func (p Uint32Vec) Len() uint { return uint(len(p)) }
