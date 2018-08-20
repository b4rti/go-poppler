package poppler

// #cgo pkg-config: --cflags-only-I poppler-glib
// #include <poppler.h>
// #include <glib.h>
import "C"

type FormField struct{}

type FormFieldType int

const (
	FormFieldTypeUnknown FormFieldType = iota
	FormFieldTypeButton
	FormFieldTypeText
	FormFieldTypeChoice
	FormFieldTypeSignature
)

func (f *FormField) GetFieldType() FormFieldType

func (f *FormField) GetID() int

func (f *FormField) IsReadOnly() bool

func (f *FormField) GetFontSize() float64

func (f *FormField) GetPartialName() string

func (f *FormField) GetMappingName() string

func (f *FormField) GetName() string

func (f *FormField) GetAction() *Action

type FormButtonType int

const (
	FormButtonTypePush FormButtonType = iota
	FormButtonTypeCheck
	FormButtonTypeRadio
)

func (f *FormField) GetButtonType() FormButtonType

func (f *FormField) GetButtonState() bool

func (f *FormField) SetButtonState(state bool)

type FormChoiceType int

const (
	FormChoiceTypeCombo FormChoiceType = iota
	FormChoiceTypeList
)

func (f *FormField) GetChoiceType() FormChoiceType

func (f *FormField) CanSelectMultiple() bool

func (f *FormField) CommitChoiceOnChange() bool

func (f *FormField) DoChoiceSpellCheck() bool

func (f *FormField) GetChoiceItem(index int) string

func (f *FormField) GetChoiceNItems() int

func (f *FormField) GetChoiceText() string

func (f *FormField) SetChoiceText(text string)

func (f *FormField) IsChoiceEditable() bool

func (f *FormField) IsChoiceItemSelected() bool

func (f *FormField) SelectChoiceItem(index int)

func (f *FormField) ToggleChoiceItem(index int)

func (f *FormField) UnselectAllItems()

type FormTextType int

const (
	FormTextTypeNormal FormTextType = iota
	FormTextTypeMultiline
	FormTextTypeFileSelect
)

func (f *FormField) GetTextType() FormTextType

func (f *FormField) GetText() string

func (f *FormField) SetText(text string)

func (f *FormField) GetTextMaxLen() int

func (f *FormField) DoTextScroll() bool

func (f *FormField) DoTextSpellCheck() bool

func (f *FormField) IsTextPassword() bool

func (f *FormField) IsRichText() bool
