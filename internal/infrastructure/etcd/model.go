package etcd



type Worker struct {
	ID       string `etcd:"id"`       // uuid format
	Platform string `etcd:"platform"` // kis, polygon, buycycle, ...
	Status   string `etcd:"status"`   // active, waiting, dead
}



type Product struct {
	ID       string `etcd:"id"`       // product_type.name.region
	Platform string `etcd:"platform"` // kis, polygon, buycycle, ...
	Symbol   string `etcd:"symbol"`   // identifier inside platform
	Worker 	 string `etcd:"worker"`   // uuid format
	Status   string `etcd:"status"`   // onsubscribe, notsubscribed
}