// Copyright 2023, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package mpegts

import (
	"github.com/q191201771/naza/pkg/nazabits"
)

// PsiId
const (
	TsPsiIdPas            = 0x00 // program_association_section
	TsPsiIdCas            = 0x01 // conditional_access_section (CA_section)
	TsPsiIdPms            = 0x02 // TS_program_map_section
	TsPsiIdDs             = 0x03 // TS_description_section
	TsPsiIdSds            = 0x04 // ISO_IEC_14496_scene_description_section
	TsPsiIdOds            = 0x05 // ISO_IEC_14496_object_descriptor_section
	TsPsiIdIso138181Start = 0x06 // ITU-T Rec. H.222.0 | ISO/IEC 13818-1 reserved
	TsPsiIdIso138181End   = 0x37
	TsPsiIdIso138186Start = 0x38 // Defined in ISO/IEC 13818-6
	TsPsiIdIso138186End   = 0x3F
	TsPsiIdUserStart      = 0x40 // User private
	TsPsiIdUserEnd        = 0xFE
	TsPsiIdForbidden      = 0xFF // forbidden
)

const (
	DescriptorTagAC3                        = 0x6a
	DescriptorTagAVCVideo                   = 0x28
	DescriptorTagComponent                  = 0x50
	DescriptorTagContent                    = 0x54
	DescriptorTagDataStreamAlignment        = 0x6
	DescriptorTagEnhancedAC3                = 0x7a
	DescriptorTagExtendedEvent              = 0x4e
	DescriptorTagExtension                  = 0x7f
	DescriptorTagISO639LanguageAndAudioType = 0xa
	DescriptorTagLocalTimeOffset            = 0x58
	DescriptorTagMaximumBitrate             = 0xe
	DescriptorTagNetworkName                = 0x40
	DescriptorTagParentalRating             = 0x55
	DescriptorTagPrivateDataIndicator       = 0xf
	DescriptorTagPrivateDataSpecifier       = 0x5f
	DescriptorTagRegistration               = 0x5
	DescriptorTagService                    = 0x48
	DescriptorTagShortEvent                 = 0x4d
	DescriptorTagStreamIdentifier           = 0x52
	DescriptorTagSubtitling                 = 0x59
	DescriptorTagTeletext                   = 0x56
	DescriptorTagVBIData                    = 0x45
	DescriptorTagVBITeletext                = 0x46
)

const (
	opusIdentifier = 0x4f707573 // Opus
)

type PsiSection struct {
	pointerFileld uint8
	sectionData   PsiSectionData
}

type PsiSectionData struct {
	header  PsiTableHeader
	section PsiTableSyntaxSection
	patData PatSpecificData
	pmtData PmtSpecificData
}

type PsiTableHeader struct {
	tableId                uint8
	sectionSyntaxIndicator uint8
	sectionLength          uint16
}

type PsiTableSyntaxSection struct {
	tableIdExtension     uint16
	versionNumber        uint8
	currentNextIndicator uint8
	sectionNumber        uint8
	lastSectionNumber    uint8
	tableData            []byte
	crc32                uint32
}

type PatSpecificData struct {
	pes []PatProgramElement
}

type PmtSpecificData struct {
	pcrPid            uint16
	programInfoLength uint16
	pes               []PmtProgramElement
}

func NewPsi() *PsiSection { _ = "STUB: not implemented"; return nil }

func (psi *PsiSection) Pack() (int, []byte) { _ = "STUB: not implemented"; return 0, nil }

func (psi *PsiSection) writePsiTableHeader(bw *nazabits.BitWriter) {
	_ = "STUB: not implemented"
	return
}

func (psi *PsiSection) writePsiTableSyntaxSection(bw *nazabits.BitWriter) {
	_ = "STUB: not implemented"
	return
}

func (psi *PsiSection) writePsiTableSyntaxSectionHeader(bw *nazabits.BitWriter) {
	_ = "STUB: not implemented"
	return
}

func (psi *PsiSection) writePsiTableSyntaxSectionData(bw *nazabits.BitWriter) {
	_ = "STUB: not implemented"
	return
}

func (psi *PsiSection) calcPsiSectionLength() (length uint16) { _ = "STUB: not implemented"; return 0 }

// Table ID extension(16 bits)+Reserved bits(2 bits)+Version number(5 bits)+Current next Indicator(1 bit)+Section number(8 bits)+Last section number(8 bits)

//crc32

func (psi *PsiSection) calaPatSectionLength() (length uint16) { _ = "STUB: not implemented"; return 0 }

func (psi *PsiSection) calaPmtSectionLength() (length uint16) {
	_ = "STUB: not implemented"
	// Reserved bits(3 bits)+PCR PID(13 bits)+Reserved bits(4 bits)+Program info length(12 bits)
	return 0
}

func (psi *PsiSection) calcDescriptorsLength(ds []Descriptor) uint16 {
	_ = "STUB: not implemented"
	return 0
}

// tag and length

func (psi *PsiSection) calcDescriptorLength(d Descriptor) uint8 {
	_ = "STUB: not implemented"
	return 0
}

func (psi *PsiSection) calcDescriptorRegistrationLength(d DescriptorRegistration) uint8 {
	_ = "STUB: not implemented"
	return 0
}

func (psi *PsiSection) calcDescriptorExtensionLength(d DescriptorExtension) uint8 {
	_ = "STUB: not implemented"
	// tag
	return 0
}

func (psi *PsiSection) writePatSection(bw *nazabits.BitWriter) { _ = "STUB: not implemented"; return }

func (psi *PsiSection) writePmtSection(bw *nazabits.BitWriter) { _ = "STUB: not implemented"; return }

func (psi *PsiSection) writeDescriptorsWithLength(bw *nazabits.BitWriter, dps []Descriptor) {
	_ = "STUB: not implemented"
	return
}

func (psi *PsiSection) writeDescriptor(bw *nazabits.BitWriter, d Descriptor) {
	_ = "STUB: not implemented"
	return
}

func (psi *PsiSection) writeDescriptorRegistration(bw *nazabits.BitWriter, d DescriptorRegistration) {
	_ = "STUB: not implemented"
	return
}

func (psi *PsiSection) writeDescriptorExtension(bw *nazabits.BitWriter, d DescriptorExtension) {
	_ = "STUB: not implemented"
	return
}

type Descriptor struct {
	Length       uint8
	Tag          uint8
	Registration DescriptorRegistration
	Extension    DescriptorExtension
}

type DescriptorRegistration struct {
	AdditionalIdentificationInfo []byte
	FormatIdentifier             uint32
}

type DescriptorExtension struct {
	SupplementaryAudio DescriptorExtensionSupplementaryAudio
	Tag                uint8
	Unknown            []byte
}

type DescriptorExtensionSupplementaryAudio struct {
	EditorialClassification uint8
	HasLanguageCode         bool
	LanguageCode            []byte
	MixType                 bool
	PrivateData             []byte
}
