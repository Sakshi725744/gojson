package main

import (
	"os"
	"fmt"
	"io/ioutil"
	"log"
    "encoding/json"
	"github.com/spdx/json/model"
	parse "github.com/spdx/json/parse_Annotation"
	
)

func readConfJson(filename string) (*model.SpdxJson, error) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	c := &model.SpdxJson{}
	err = json.Unmarshal(buf, c)
	if err != nil {
		return nil, fmt.Errorf("in file %q: %v", filename, err)
	}

	return c, nil
}


func main() {
	c, err := readConfJson("./examples/spdxJson.json")
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.OpenFile("./examples/jsonOutput.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal(err)
    }

	if _, err := f.Write([]byte("Annotations2v2:\n\n")); err != nil {
		log.Fatal(err)
	}
	annotations2v2 := parse.ParseAnnotation2v2(c)

	for i := 0; i < len(annotations2v2); i++ {
		fmt.Println((annotations2v2)[i].AnnotationComment)
		fmt.Println((annotations2v2)[i].AnnotationType)
		fmt.Println((annotations2v2)[i].AnnotatorType)
		fmt.Println((annotations2v2)[i].Annotator)
		fmt.Println((annotations2v2)[i].AnnotationDate)

        //Writing annotations2v2 to ./examples/yamlOutput.txt
		if _, err := f.Write([]byte((annotations2v2)[i].AnnotationComment+"\n")); err != nil {
			log.Fatal(err)
		}

		if _, err := f.Write([]byte((annotations2v2)[i].AnnotationType+"\n")); err != nil {
			log.Fatal(err)
		}

		if _, err := f.Write([]byte((annotations2v2)[i].AnnotatorType+"\n")); err != nil {
			log.Fatal(err)
		}

		if _, err := f.Write([]byte((annotations2v2)[i].AnnotationDate+"\n\n")); err != nil {
			log.Fatal(err)
		}
	}
	annotations2v1 := parse.ParseAnnotation2v1(c)
	if _, err := f.Write([]byte("Annotations2v1:\n\n")); err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(annotations2v1); i++ {
		fmt.Println((annotations2v1)[i].AnnotationComment)
		fmt.Println((annotations2v1)[i].AnnotationType)
		fmt.Println((annotations2v1)[i].AnnotatorType)
		fmt.Println((annotations2v1)[i].Annotator)
		fmt.Println((annotations2v1)[i].AnnotationDate)

		//Writing annotations2v1 to ./examples/yamlOutput.txt
		if _, err := f.Write([]byte((annotations2v1)[i].AnnotationComment+"\n")); err != nil {
			log.Fatal(err)
		}
		if _, err := f.Write([]byte((annotations2v1)[i].AnnotationType+"\n")); err != nil {
			log.Fatal(err)
		}
		if _, err := f.Write([]byte((annotations2v1)[i].AnnotatorType+"\n")); err != nil {
			log.Fatal(err)
		}
		if _, err := f.Write([]byte((annotations2v1)[i].AnnotationDate+"\n\n")); err != nil {
			log.Fatal(err)
		}
	}
	   
    if err := f.Close(); err != nil {
        log.Fatal(err)
    }
}