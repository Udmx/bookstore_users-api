package services

var (
	ItemsService itemsServiceInterface = &itemService{}
)

type itemsServiceInterface interface {
	GetItem()
	SaveItem()
}

type itemService struct {
}

func (s *itemService) GetItem() {

}

func (s *itemService) SaveItem() {

}
