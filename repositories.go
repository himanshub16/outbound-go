package main

type LinkRepository interface {
	InsertLink(link Link) (error)
	FindLinkByShortIdInt(id uint) (*Link, error)
	UpdateLink(link Link) (error)
	close()
}

type CounterRepository interface {
	FindCounterById(id string) (*Counter, error)
	UpsertCounter(counter Counter) (error)
	close()
}
