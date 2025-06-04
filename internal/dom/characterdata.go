package dom

import (
	"github.com/dop251/goja"
)

// CharacterData represents a Node that contains characters.
// This is an abstract interface implemented by Text, Comment, and ProcessingInstruction.
type CharacterData interface {
	Node
	Data() string
	SetData(data string)
	Length() int
	SubstringData(offset, count int) string
	AppendData(arg string)
	InsertData(offset int, arg string)
	DeleteData(offset, count int)
	ReplaceData(offset, count int, arg string)
}

// characterDataImpl provides a common base for CharacterData nodes.
type characterDataImpl struct {
	nodeImpl
}

func (cd *characterDataImpl) Data() string {
	return cd.nodeValue
}

func (cd *characterDataImpl) SetData(data string) {
	// Use the replace data algorithm to maintain consistency
	cd.replaceDataInternal(0, cd.Length(), data)
}

func (cd *characterDataImpl) Length() int {
	// Return length in Unicode code units (UTF-16)
	return len([]rune(cd.nodeValue))
}

// replaceDataInternal implements the "replace data" algorithm from the WHATWG DOM specification
func (cd *characterDataImpl) replaceDataInternal(offset, count int, data string) {
	// Step 1: Let length be node's length
	length := cd.Length()

	// Step 2: If offset is greater than length, then throw an "IndexSizeError" DOMException
	if offset > length {
		panic(NewIndexSizeError("offset is greater than length"))
	}

	// Step 3: If offset plus count is greater than length, then set count to length minus offset
	if offset+count > length {
		count = length - offset
	}

	// Step 4: Queue a mutation record of "characterData" for node
	// TODO: Implement mutation record queuing when mutation observer integration is needed

	// Steps 5-7: Perform the data replacement
	runes := []rune(cd.nodeValue)

	// Insert data at offset
	newRunes := make([]rune, 0, len(runes)+len([]rune(data))-count)
	newRunes = append(newRunes, runes[:offset]...)
	newRunes = append(newRunes, []rune(data)...)
	newRunes = append(newRunes, runes[offset+count:]...)

	cd.nodeValue = string(newRunes)

	// Steps 8-11: Update live ranges
	// Note: Live ranges are not implemented in this DOM implementation yet
	// The specification requires updating range start/end offsets here

	// Step 12: If node's parent is non-null, then run the children changed steps
	// Note: This would trigger parent element updates if implemented
}

// substringDataInternal implements the "substring data" algorithm from the WHATWG DOM specification
func (cd *characterDataImpl) substringDataInternal(offset, count int) string {
	// Step 1: Let length be node's length
	length := cd.Length()

	// Step 2: If offset is greater than length, then throw an "IndexSizeError" DOMException
	if offset > length {
		panic(NewIndexSizeError("offset is greater than length"))
	}

	// Step 3: If offset plus count is greater than length, return substring to end
	runes := []rune(cd.nodeValue)
	if offset+count > length {
		return string(runes[offset:])
	}

	// Step 4: Return the substring
	return string(runes[offset : offset+count])
}

func (cd *characterDataImpl) SubstringData(offset, count int) string {
	return cd.substringDataInternal(offset, count)
}

func (cd *characterDataImpl) AppendData(arg string) {
	// Use replace data algorithm: offset = length, count = 0, data = arg
	cd.replaceDataInternal(cd.Length(), 0, arg)
}

func (cd *characterDataImpl) InsertData(offset int, arg string) {
	// Use replace data algorithm: count = 0, data = arg
	cd.replaceDataInternal(offset, 0, arg)
}

func (cd *characterDataImpl) DeleteData(offset, count int) {
	// Use replace data algorithm: data = empty string
	cd.replaceDataInternal(offset, count, "")
}

func (cd *characterDataImpl) ReplaceData(offset, count int, arg string) {
	// Use replace data algorithm directly
	cd.replaceDataInternal(offset, count, arg)
}

// toJS is a placeholder for JavaScript binding.
func (cd *characterDataImpl) toJS(vm *goja.Runtime) goja.Value {
	// This will be implemented later when integrating with Goja
	return vm.ToValue(nil)
}
