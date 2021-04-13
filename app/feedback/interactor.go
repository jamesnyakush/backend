package feedback

type Interactor interface {

}

func NewInteractor() Interactor {
	return &interactor{

	}
}

type interactor struct {
}
