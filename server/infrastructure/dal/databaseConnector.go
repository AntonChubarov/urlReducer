package dal

type DatabaseConnector struct {

}

func NewDatabaseConnector() *DatabaseConnector {
	return &DatabaseConnector{}
}

func (d *DatabaseConnector) GetInitialLinkFromStorage(id string) (url string, err error) {
	return "http://this.should.be/an?initial=link", err
}

func (d *DatabaseConnector) SaveInitialLinkToStorage(url string, id string) error {
	return nil
}

