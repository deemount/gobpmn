package artifacts

type Type_Text TextRepository
type Type_Text_Annotation TextAnnotationRepository

type ArtifactsRepository interface {
	TextRepository
	TextAnnotationRepository
}

type Artifacts struct {
	ArtifactsRepository
}

func NewArtifacts() ArtifactsRepository {
	return &Artifacts{}
}
