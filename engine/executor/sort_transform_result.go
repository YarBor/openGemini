// Generated by tmpl
// https://github.com/benbjohnson/tmpl
//
// DO NOT EDIT!
// Source: sort_transform_result.go.tmpl

/*
Copyright 2023 Huawei Cloud Computing Technologies Co., Ltd.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package executor

const (
	less = iota
	eq
	greater
)

type sortEleMsg interface {
	LessThan(ele sortEleMsg) int
	SetVal(col Column, startLoc int)
	AppendToCol(col Column)
}

type sortRowMsg struct {
	sortEle []sortEleMsg // tags+fields+time
}

func NewSortRowMsg(eles []sortEleMsg) *sortRowMsg {
	return &sortRowMsg{
		sortEle: eles,
	}
}

func (sr *sortRowMsg) SetVals(chunk Chunk, startLoc int, tagVals []string) {
	colLoc := 0
	for ; tagVals != nil && colLoc < len(tagVals); colLoc++ {
		sr.sortEle[colLoc].(*stringSortEle).val = tagVals[colLoc]
		sr.sortEle[colLoc].(*stringSortEle).validVal = true
	}
	for _, col := range chunk.Columns() {
		sr.sortEle[colLoc].SetVal(col, startLoc)
		colLoc++
	}
	sr.sortEle[colLoc].(*integerSortEle).val = chunk.TimeByIndex(startLoc)
	sr.sortEle[colLoc].(*integerSortEle).validVal = true
}

func (sr *sortRowMsg) AppendToChunk(chunk Chunk, startColLoc int) {
	for colLoc := startColLoc; colLoc < len(sr.sortEle)-1; colLoc++ {
		sr.sortEle[colLoc].AppendToCol(chunk.Column(colLoc - startColLoc))
	}
	chunk.AppendTime(sr.sortEle[len(sr.sortEle)-1].(*integerSortEle).val)
}

func (sr *sortRowMsg) LessThan(osr *sortRowMsg, sortKeysIdxs []int, ascending []bool) bool {
	for i, idx := range sortKeysIdxs {
		subResult := sr.sortEle[idx].LessThan(osr.sortEle[idx])
		if ascending[i] {
			if subResult == less {
				return true
			} else if subResult == greater {
				return false
			}
		} else {
			if subResult == greater {
				return true
			} else if subResult == less {
				return false
			}
		}
	}
	// two row equ return true, not exchange val
	return true
}

type floatSortEle struct {
	val      float64
	validVal bool
}

func NewFloatSortEle() sortEleMsg {
	return &floatSortEle{
		val:      0,
		validVal: false,
	}
}

func (ele *floatSortEle) LessThan(oele sortEleMsg) int {
	if ele.validVal && oele.(*floatSortEle).validVal {
		if ele.val < oele.(*floatSortEle).val {
			return less
		} else if ele.val == oele.(*floatSortEle).val {
			return eq
		} else {
			return greater
		}
	} else {
		if !ele.validVal && oele.(*floatSortEle).validVal {
			return less
		} else if !ele.validVal && !oele.(*floatSortEle).validVal {
			return eq
		} else {
			return greater
		}
	}
}

func (ele *floatSortEle) SetVal(col Column, startLoc int) {
	if col.IsNilV2(startLoc) {
		return
	}
	ele.validVal = true
	if col.NilCount() == 0 {
		ele.val = col.FloatValue(startLoc)
		return
	}
	startLoc = col.GetValueIndexV2(startLoc)
	ele.val = col.FloatValue(startLoc)
}

func (ele *floatSortEle) AppendToCol(col Column) {
	if !ele.validVal {
		col.AppendNilsV2(ele.validVal)
	} else {
		col.AppendFloatValue(ele.val)
		col.AppendNilsV2(ele.validVal)
	}
}

type integerSortEle struct {
	val      int64
	validVal bool
}

func NewIntegerSortEle() sortEleMsg {
	return &integerSortEle{
		val:      0,
		validVal: false,
	}
}

func (ele *integerSortEle) LessThan(oele sortEleMsg) int {
	if ele.validVal && oele.(*integerSortEle).validVal {
		if ele.val < oele.(*integerSortEle).val {
			return less
		} else if ele.val == oele.(*integerSortEle).val {
			return eq
		} else {
			return greater
		}
	} else {
		if !ele.validVal && oele.(*integerSortEle).validVal {
			return less
		} else if !ele.validVal && !oele.(*integerSortEle).validVal {
			return eq
		} else {
			return greater
		}
	}
}

func (ele *integerSortEle) SetVal(col Column, startLoc int) {
	if col.IsNilV2(startLoc) {
		return
	}
	ele.validVal = true
	if col.NilCount() == 0 {
		ele.val = col.IntegerValue(startLoc)
		return
	}
	startLoc = col.GetValueIndexV2(startLoc)
	ele.val = col.IntegerValue(startLoc)
}

func (ele *integerSortEle) AppendToCol(col Column) {
	if !ele.validVal {
		col.AppendNilsV2(ele.validVal)
	} else {
		col.AppendIntegerValue(ele.val)
		col.AppendNilsV2(ele.validVal)
	}
}

type stringSortEle struct {
	val      string
	validVal bool
}

func NewStringSortEle() sortEleMsg {
	return &stringSortEle{
		val:      "",
		validVal: false,
	}
}

func (ele *stringSortEle) LessThan(oele sortEleMsg) int {
	if ele.validVal && oele.(*stringSortEle).validVal {
		if ele.val < oele.(*stringSortEle).val {
			return less
		} else if ele.val == oele.(*stringSortEle).val {
			return eq
		} else {
			return greater
		}
	} else {
		if !ele.validVal && oele.(*stringSortEle).validVal {
			return less
		} else if !ele.validVal && !oele.(*stringSortEle).validVal {
			return eq
		} else {
			return greater
		}
	}
}

func (ele *stringSortEle) SetVal(col Column, startLoc int) {
	if col.IsNilV2(startLoc) {
		return
	}
	ele.validVal = true
	if col.NilCount() == 0 {
		ele.val = col.StringValue(startLoc)
		return
	}
	startLoc = col.GetValueIndexV2(startLoc)
	ele.val = col.StringValue(startLoc)
}

func (ele *stringSortEle) AppendToCol(col Column) {
	if !ele.validVal {
		col.AppendNilsV2(ele.validVal)
	} else {
		col.AppendStringValue(ele.val)
		col.AppendNilsV2(ele.validVal)
	}
}

type boolSortEle struct {
	val      bool
	validVal bool
}

func NewBoolSortEle() sortEleMsg {
	return &boolSortEle{
		val:      false,
		validVal: false,
	}
}

func (ele *boolSortEle) LessThan(oele sortEleMsg) int {
	if ele.validVal && oele.(*boolSortEle).validVal {
		if !ele.val && oele.(*boolSortEle).val {
			return less
		} else if ele.val == oele.(*boolSortEle).val {
			return eq
		} else {
			return greater
		}
	} else {
		if !ele.validVal && oele.(*boolSortEle).validVal {
			return less
		} else if !ele.validVal && !oele.(*boolSortEle).validVal {
			return eq
		} else {
			return greater
		}
	}
}

func (ele *boolSortEle) SetVal(col Column, startLoc int) {
	if col.IsNilV2(startLoc) {
		return
	}
	ele.validVal = true
	if col.NilCount() == 0 {
		ele.val = col.BooleanValue(startLoc)
		return
	}
	startLoc = col.GetValueIndexV2(startLoc)
	ele.val = col.BooleanValue(startLoc)
}

func (ele *boolSortEle) AppendToCol(col Column) {
	if !ele.validVal {
		col.AppendNilsV2(ele.validVal)
	} else {
		col.AppendBooleanValue(ele.val)
		col.AppendNilsV2(ele.validVal)
	}
}