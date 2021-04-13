package model

type SpdxJson struct {
	SpdxId string `json:"SPDXID"`

	SpdxVersion string `json:"spdxVersion"`

	CreationInfo struct {
		Comment            string   `json:"comment"`
		Created            string   `json:"created"`
		Creators           []string `json:"creators"`
		LicenseListVersion string   `json:"licenseListVersion"`
	} `json:"creationInfo"`

	Name string `json:"name"`

	DataLicense string `json:"dataLicense"`

	Comment string `json:"comment"`

	ExternalDocumentRefs []struct {
		ExternalDocumentId string `json:"externalDocumentId"`
		Checksum           struct {
			Algorithm     string `json:"algorithm"`
			ChecksumValue string `json:"checksumValue"`
		} `json:"checksum"`
		SpdxDocument string `json:"spdxDocument"`
	} `json:"externalDocumentRefs"`

	HasExtractedLicensingInfos []struct {
		ExtractedTexts string   `json:"extractedTexts"`
		LicenseId      string   `json:"licenseId"`
		Comment        string   `json:"comment,omitempty"`
		Name           string   `json:"name,omitempty"`
		SeeAlsos       []string `json:"seeAlsos,omitempty"`
	} `json:"hasExtractedLicensingInfos"`

	Annotations []struct {
		AnnotationDate string `json:"annotationDate"`
		AnnotationType string `json:"annotationType"`
		Annotator      string `json:"annotator"`
		Comment        string `json:"comment"`
	} `json:"annotations"`

	DocumentNamespace string `json:"documentNamespace"`

	DocumentDescribes []string `json:"documentDescribes"`

	Packages []struct {
		SpdxId      string `json:"SPDXID"`
		Annotations []struct {
			AnnotationDate string `json:"annotationDate"`
			AnnotationType string `json:"annotationType"`
			Annotator      string `json:"annotator"`
			Comment        string `json:"comment"`
		} `json:"annotations,omitempty"`
		Checksums []struct {
			Algorithm     string `json:"algorithm"`
			ChecksumValue string `json:"checksumValue"`
		} `json:"checksums,omitempty"`
		CopyrightText    string `json:"copyrightText,omitempty"`
		Description      string `json:"description,omitempty"`
		DownloadLocation string `json:"downloadLocation,omitempty"`
		ExternalRefs     []struct {
			Comment           string `json:"comment,omitempty"`
			ReferenceCategory string `json:"referenceCategory"`
			ReferenceLocator  string `json:"referenceLocator"`
			ReferenceType     string `json:"referenceType"`
		} `json:"externalRefs,omitempty"`
		FilesAnalyzed           bool     `json:"filesAnalyzed,omitempty"`
		HasFiles                []string `json:"hasFiles,omitempty"`
		Homepage                string   `json:"homepage,omitempty"`
		LicenseComments         string   `json:"licenseComments,omitempty"`
		LicenseConcluded        string   `json:"licenseConcluded,omitempty"`
		LicenseDeclared         string   `json:"licenseDeclared,omitempty"`
		LicenseInfoInFiles      []string `json:"licenseInfoInFiles,omitempty"`
		Name                    string   `json:"name,omitempty"`
		Originator              string   `json:"originator,omitempty"`
		PackageFileName         string   `json:"packageFileName"`
		PackageVerificationCode struct {
			PackageVerificationCodeExcludedFiles []string `json:"packageVerificationCodeExcludedFiles"`
			PackageVerificationCodeValue         string   `json:"packageVerificationCodeValue"`
		} `json:"packageVerificationCode,omitempty"`
		SourceInfo  string `json:"sourceInfo,omitempty"`
		Summary     string `json:"summary,omitempty"`
		Supplier    string `json:"supplier,omitempty"`
		VersionInfo string `json:"versionInfo,omitempty"`
	}

	Files []struct {
		SpdxId      string `json:"SPDXID"`
		Annotations []struct {
			AnnotationDate string `json:"annotationDate"`
			AnnotationType string `json:"annotationType"`
			Annotator      string `json:"annotator"`
			Comment        string `json:"comment"`
		} `json:"annotations,omitempty"`
		Checksums []struct {
			Algorithm     string `json:"algorithm"`
			ChecksumValue string `json:"checksumValue"`
		} `json:"checksums"`
		Comment            string   `json:"comment,omitempty"`
		CopyrightText      string   `json:"copyrightText "`
		FileContributors   []string `json:"fileContributors"`
		FileDependencies   []string `json:"fileDependencies"`
		FileName           string   `json:"fileName"`
		FileTypes          []string `json:"fileTypes"`
		LicenseComments    string   `json:"licenseComments,omitempty"`
		LicenseConcluded   string   `json:"licenseConcluded"`
		LicenseInfoInFiles []string `json:"licenseInfoInFiles"`
		NoticeText         string   `json:"noticeText,omitempty"`
	}

	Snippets []struct {
		SpdxId             string   `json:"SPDXID"`
		Comment            string   `json:"comment,omitempty"`
		CopyrightText      string   `json:"copyrightText "`
		LicenseComments    string   `json:"licenseComments,omitempty"`
		LicenseConcluded   string   `json:"licenseConcluded"`
		LicenseInfoInFiles []string `json:"licenseInfoInFiles"`
		Name               string   `json:"name"`
		Ranges             []struct {
			EndPointer struct {
				Offset    int    `json:"offset"`
				Reference string `json:"reference"`
			} `json:"endPointer"`
			StartPointer struct {
				Offset    int    `json:"offset"`
				Reference string `json:"reference"`
			} `json:"name"`
		} `json:"startPointer"`
		SnippetFromFile string `json:"snippetFromFile"`
	}

	Relationships []struct {
		SpdxElementId      string `json:"spdxElementId"`
		RelatedSpdxElement string `json:"relatedSpdxElement"`
		RelationshipType   string `json:"relationshipType"`
	} `json:"relationships"`
}
