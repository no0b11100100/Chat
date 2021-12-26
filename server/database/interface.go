package database

type DBInterface interface {
	AddRecord(record Record)
	Select(request string) (Record, error)
	IsEmailUnique(string) bool
	Close()
}
