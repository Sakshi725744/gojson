package parse

import (
	"fmt"

	"github.com/spdx/json/model"
	spdx "github.com/spdx/json/spdx"
)

func ParseAnnotation2v1(parser *model.SpdxJson) (annotations []spdx.Annotation2_1) {
	//type annotations *[]spdx.Annotation2_2
	a := spdx.Annotation2_1{}

	for i := 0; i < len(parser.Annotations); i++ {

		//Annotation Date in 2v2
		a.AnnotationDate = (*parser).Annotations[i].AnnotationDate

		//Annotation Comment in 2v2
		a.AnnotationComment = (*parser).Annotations[i].Comment

		//Annotation Type in 2v2
		setAnnotationType2v1((*parser).Annotations[i].AnnotationType, &a)

		//Annotator and Annotator Type in 2v2
		setAnnotatorFromString2v1((*parser).Annotations[i].Annotator, &a)

		(annotations) = append((annotations), a)

	}
	return annotations
}

func setAnnotatorFromString2v1(annotatorString string, ann *spdx.Annotation2_1) error {
	subkey, subvalue, err := ExtractSubs(annotatorString, ":")
	if err != nil {
		return err
	}
	if subkey == "Person" || subkey == "Organization" || subkey == "Tool" {
		ann.AnnotatorType = subkey
		ann.Annotator = subvalue
		return nil
	}
	return fmt.Errorf("unrecognized Annotator type %v while parsing annotation", subkey)
}

func setAnnotationType2v1(annType string, ann *spdx.Annotation2_1) error {
	if annType == "OTHER" || annType == "REVIEW" {
		ann.AnnotationType = annType
	} else {
		return fmt.Errorf("unknown annotation type %s", annType)
	}
	return nil
}
