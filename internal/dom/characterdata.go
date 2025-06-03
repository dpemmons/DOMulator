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
	cd.nodeValue = data
}

func (cd *characterDataImpl) Length() int {
	return len(cd.nodeValue)
}

func (cd *characterDataImpl) SubstringData(offset, count int) string {
	runes := []rune(cd.nodeValue)
	if offset < 0 || offset > len(runes) {
		// TODO: Throw DOMException
		return ""
	}
	if count < 0 {
		// TODO: Throw DOMException
		return ""
	}

	end := offset + count
	if end > len(runes) {
		end = len(runes)
	}
	return string(runes[offset:end])
}

func (cd *characterDataImpl) AppendData(arg string) {
	cd.nodeValue += arg
}

func (cd *characterDataImpl) InsertData(offset int, arg string) {
	runes := []rune(cd.nodeValue)
	if offset < 0 || offset > len(runes) {
		// TODO: Throw DOMException
		return
	}
	cd.nodeValue = string(runes[:offset]) + arg + string(runes[offset:])
}

func (cd *characterDataImpl) DeleteData(offset, count int) {
	runes := []rune(cd.nodeValue)
	if offset < 0 || offset > len(runes) {
		// TODO: Throw DOMException
		return
	}
	if count < 0 {
		// TODO: Throw DOMException
		return
	}

	end := offset + count
	if end > len(runes) {
		end = len(runes)
	}
	cd.nodeValue = string(runes[:offset]) + string(runes[end:])
}

func (cd *characterDataImpl) ReplaceData(offset, count int, arg string) {
	runes := []rune(cd.nodeValue)
	if offset < 0 || offset > len(runes) {
		// TODO: Throw DOMException
		return
	}
	if count < 0 {
		// TODO: Throw DOMException
		return
	}

	end := offset + count
	if end > len(runes) {
		end = len(runes)
	}
	cd.nodeValue = string(runes[:offset]) + arg + string(runes[end:])
}

// toJS is a placeholder for JavaScript binding.
func (cd *characterDataImpl) toJS(vm *goja.Runtime) goja.Value {
	// This will be implemented later when integrating with Goja
	return vm.ToValue(nil)
}
