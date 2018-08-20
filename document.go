package poppler

// #cgo pkg-config: --cflags-only-I poppler-glib
// #include <poppler.h>
// #include <stdlib.h>
// #include <glib.h>
// #include <unistd.h>
import "C"
import (
	"errors"
	"io"
	"path/filepath"
)

type Document struct {
	doc poppDoc
}

type poppDoc *C.struct__PopplerDocument

func NewDocumentFromPath(path string) (doc *Document, err error) {
	path, err = filepath.Abs(path)
	if err != nil {
		return
	}
	var e *C.GError
	fn := C.g_filename_to_uri((*C.gchar)(C.CString(path)), nil, nil)
	var d poppDoc
	d = C.poppler_document_new_from_file((*C.char)(fn), nil, &e)
	if e != nil {
		err = errors.New(C.GoString((*C.char)(e.message)))
	}
	doc = &Document{
		doc: d,
	}
	return
}

func NewDocumentFromData(data []byte) (*Document, error)

func NewDocumentFromReader(reader io.Reader) (*Document, error)

func (d *Document) SaveToPath(path string) error

func (d *Document) SaveToData() ([]byte, error)

func (d *Document) SaveToWriter(writer io.Writer) error

func (d *Document) SaveCopyToPath(path string) error

func (d *Document) SaveCopyToData() ([]byte, error)

func (d *Document) SaveCopyToWriter(writer io.Writer) error

type DocumentID struct {
	PermanentID string
	UpdateID    string
}

func (d *Document) GetID() DocumentID

func (d *Document) GetVersionString() string

type DocumentVersion struct {
	Major uint
	Minor uint
}

func (d *Document) GetVersion() *DocumentVersion

func (d *Document) GetTitle() string

func (d *Document) SetTitle(title string)

func (d *Document) GetAuthor() string

func (d *Document) SetAuthor(author string)

func (d *Document) GetSubject() string

func (d *Document) SetSubject(subject string)

func (d *Document) GetKeywords() string

func (d *Document) SetKeywords(keywords string)

func (d *Document) GetCreator() string

func (d *Document) SetCreator(creator string)

func (d *Document) GetProducer() string

func (d *Document) SetProducer(producer string)

func (d *Document) GetCreationDate() int

func (d *Document) SetCreationDate(creationDate int)

func (d *Document) GetModificationDate() int

func (d *Document) SetModificationDate(modificationDate int)

type PageLayout int

const (
	PageLayoutUnset PageLayout = iota
	PageLayoutSignlePage
	PageLayoutOneColumn
	PageLayoutTwoColumnLeft
	PageLayoutTwoColumnRight
	PageLayoutTwoPageLeft
	PageLayoutTwoPageRigh
)

func (d *Document) GetPageLayout() PageLayout

type PageMode int

const (
	PageModeUnset PageMode = 1 << iota
	PageModeNone
	PageModeUseOutlines
	PageModeUseThumbs
	PageModeFullScreen
	PageModeUseOC
	PageModeUseAttachments
)

func (d *Document) GetPageMode() PageMode

type Permissions int

const (
	PermissionsOKToPrint Permissions = 1 << iota
	PermissionsOKToModify
	PermissionsOKToCopy
	PermissionsOKToAddNotes
	PermissionsOKToFillForm
	PermissionsOKToExtractContent
	PermissionsOKToAssemble
	PermissionsOKToPrintHighResolution
	PermissionsFull
)

func (d *Document) GetPermissions() Permissions

func (d *Document) GetMetadata() string

func (d *Document) IsLineazied() bool

func (d *Document) GetNPages() int {
	return int(C.poppler_document_get_n_pages(d.doc))
}

func (d *Document) GetPage(index int) *Page {
	p := C.poppler_document_get_page(d.doc, C.int(index))
	return &Page{p: p}
}

func (d *Document) GetPageByLabel(label string) *Page

type Destination struct {
	destType DestinationType

	pageNum    int
	left       float64
	bottom     float64
	right      float64
	top        float64
	zoom       float64
	namedDest  string
	changeLeft uint
	changeTop  uint
	changeZoom uint
}

type DestinationType int

const (
	DestinationTypeUnknown DestinationType = iota
	DestinationTypeXYZ
	DestinationTypeFit
	DestinationTypeFitH
	DestinationTypeFitV
	DestinationTypeFitR
	DestinationTypeFitB
	DestinationTypeFitBH
	DestinationTypeFitBV
	DestinationTypeNamed
)

func (d *Document) FindDest(linkName string) (destination Destination)

func (d *Document) GetNAttachments() int {
	return int(C.poppler_document_get_n_attachments(d.doc))
}

func (d *Document) HasAttachments() bool {
	return toBool(C.poppler_document_has_attachments(d.doc))
}

func (d *Document) GetAttachments() []*Attachment

//
//
//
//
//

type DocumentInfo struct {
	PdfVersion, Title, Author, Subject, KeyWords, Creator, Producer, Metadata string
	CreationDate, ModificationDate, Pages                                     int
	IsLinearized                                                              bool
}

func (d *Document) Info() DocumentInfo {
	return DocumentInfo{
		PdfVersion:       toString(C.poppler_document_get_pdf_version_string(d.doc)),
		Title:            toString(C.poppler_document_get_title(d.doc)),
		Author:           toString(C.poppler_document_get_author(d.doc)),
		Subject:          toString(C.poppler_document_get_subject(d.doc)),
		KeyWords:         toString(C.poppler_document_get_keywords(d.doc)),
		Creator:          toString(C.poppler_document_get_creator(d.doc)),
		Producer:         toString(C.poppler_document_get_producer(d.doc)),
		Metadata:         toString(C.poppler_document_get_metadata(d.doc)),
		CreationDate:     int(C.poppler_document_get_creation_date(d.doc)),
		ModificationDate: int(C.poppler_document_get_modification_date(d.doc)),
		Pages:            int(C.poppler_document_get_n_pages(d.doc)),
		IsLinearized:     toBool(C.poppler_document_is_linearized(d.doc)),
	}
}
