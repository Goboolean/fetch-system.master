package etcd

import (
	"context"

	"go.etcd.io/etcd/client/v3"
	"github.com/Goboolean/fetch-system.master/internal/infrastructure/etcd/util"
)


func (c *Client) InsertWorker(ctx context.Context, w *Worker) error {
	payload, err := etcdutil.Mmarshal(w)

	var ops []clientv3.Op
	for k, v := range payload {
		ops = append(ops, clientv3.OpPut(k, v))
	}

	var conditions []clientv3.Cmp
	for k := range payload {
		conditions = append(conditions, clientv3.Compare(clientv3.Version(k), "=", 0))
	}

	resp, err := c.client.Txn(ctx).
		If(conditions...).
		Then(ops...).
		Commit()

	if err != nil {
		return err
	}
	if flag := resp.Succeeded; !flag {
		return ErrObjectExists
	}
	return err
}


func (c *Client) GetWorker(ctx context.Context, id string) (*Worker, error) {

	resp, err := c.client.Get(context.Background(), etcdutil.Identifier("worker", id), clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}

	if len(resp.Kvs) == 0 {
		return nil, ErrWorkerNotExists
	}

	m := etcdutil.PayloadToMap(resp)

	var w Worker
	if err := etcdutil.Unmarshal(m, &w); err != nil {
		return nil, err
	}
	return &w, nil
}

func (c *Client) GetAllWorkers(ctx context.Context) ([]*Worker, error) {
	
	resp, err := c.client.Get(context.Background(), etcdutil.Group("worker"), clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}
	m := etcdutil.PayloadToMap(resp)

	list, err := etcdutil.GroupBy(m)
	if err != nil {
		return nil, err
	}

	var w []*Worker = make([]*Worker, len(list))
	for i, v := range list {
		var worker Worker
		if err := etcdutil.Unmarshal(v, &worker); err != nil {
			return nil, err
		}
		w[i] = &worker
	}
	return w, nil
}

func (c *Client) UpdateWorkerStatus(ctx context.Context, id string, status string) error {
	_, err := c.client.Put(context.Background(), etcdutil.Field("worker", id, "status"), status)
	return err
}

func (c *Client) DeleteWorker(ctx context.Context, id string) error {
	_, err := c.client.Delete(context.Background(), etcdutil.Identifier("worker", id), clientv3.WithPrefix())
	return err
}


func (c *Client) InsertOneProduct(ctx context.Context, p *Product) error {
	payload, err := etcdutil.Mmarshal(p)


	var conditions []clientv3.Cmp
	for k := range payload {
		conditions = append(conditions, clientv3.Compare(clientv3.Version(k), "=", 0))
	}

	var ops []clientv3.Op
	for k, v := range payload {
		ops = append(ops, clientv3.OpPut(k, v))
	}

	resp, err := c.client.Txn(ctx).
		If(conditions...).
		Then(ops...).
		Commit()

	if err != nil {
		return err
	}
	if flag := resp.Succeeded; !flag {
		return ErrObjectExists
	}
	return nil
}

func (c *Client) InsertProducts(ctx context.Context , p []*Product) error {

	var conditions []clientv3.Cmp
	var ops []clientv3.Op

	for _, v := range p {
		payload, err := etcdutil.Mmarshal(v)
		if err != nil {
			return err
		}
		for k, v := range payload {
			ops = append(ops, clientv3.OpPut(k, v))
			conditions = append(conditions, clientv3.Compare(clientv3.Version(k), "=", 0))
		}
	}

	resp, err := c.client.Txn(ctx).
		If(conditions...).
		Then(ops...).
		Commit()

	if err != nil {
		return err
	}
	if flag := resp.Succeeded; !flag {
		return ErrObjectExists
	}
	return nil
}

func (c *Client) GetProduct(ctx context.Context, id string) (*Product, error) {

	resp, err := c.client.Get(context.Background(), etcdutil.Identifier("product", id), clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}

	if len(resp.Kvs) == 0 {
		return nil, ErrProductNotExists
	}

	m := etcdutil.PayloadToMap(resp)

	var p Product
	if err := etcdutil.Unmarshal(m, &p); err != nil {
		return nil, err
	}
	return &p, nil
}

func (c *Client) GetAllProducts(ctx context.Context) ([]*Product, error) {

	resp, err := c.client.Get(context.Background(), etcdutil.Group("product"), clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}
	m := etcdutil.PayloadToMap(resp)

	list, err := etcdutil.GroupBy(m)
	if err != nil {
		return nil, err
	}

	var p []*Product = make([]*Product, len(list))
	for i, v := range list {
		var product Product
		if err := etcdutil.Unmarshal(v, &product); err != nil {
			return nil, err
		}
		p[i] = &product
	}
	return p, nil
}

func (c *Client) UpdateProductStatus(ctx context.Context, id string, status string) error {
	_, err := c.client.Put(context.Background(), etcdutil.Field("product", id, "status"), status)
	return err
}

func (c *Client) UpdateProductWorker(ctx context.Context, id string, worker string) error {
	_, err := c.client.Put(context.Background(), etcdutil.Field("product", id, "worker"), worker)
	return err
}

func (c *Client) DeleteProduct(ctx context.Context, id string) error {
	_, err := c.client.Delete(context.Background(), etcdutil.Identifier("product", id), clientv3.WithPrefix())
	return err
}