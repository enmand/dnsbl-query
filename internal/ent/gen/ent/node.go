// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/enmand/dnsbl-query/internal/ent/gen/ent/dnsblquery"
	"github.com/enmand/dnsbl-query/internal/ent/gen/ent/dnsblresponse"
	"github.com/enmand/dnsbl-query/internal/ent/gen/ent/ip"
	"github.com/facebookincubator/ent-contrib/entgql"
	"github.com/google/uuid"
	"github.com/hashicorp/go-multierror"
)

// Noder wraps the basic Node method.
type Noder interface {
	Node(context.Context) (*Node, error)
}

// Node in the graph.
type Node struct {
	ID     uuid.UUID `json:"id,omitemty"`      // node id.
	Type   string    `json:"type,omitempty"`   // node type.
	Fields []*Field  `json:"fields,omitempty"` // node fields.
	Edges  []*Edge   `json:"edges,omitempty"`  // node edges.
}

// Field of a node.
type Field struct {
	Type  string `json:"type,omitempty"`  // field type.
	Name  string `json:"name,omitempty"`  // field name (as in struct).
	Value string `json:"value,omitempty"` // stringified value.
}

// Edges between two nodes.
type Edge struct {
	Type string      `json:"type,omitempty"` // edge type.
	Name string      `json:"name,omitempty"` // edge name.
	IDs  []uuid.UUID `json:"ids,omitempty"`  // node ids (where this edge point to).
}

func (dq *DNSBLQuery) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     dq.ID,
		Type:   "DNSBLQuery",
		Fields: make([]*Field, 0),
		Edges:  make([]*Edge, 2),
	}
	node.Edges[0] = &Edge{
		Type: "DNSBLResponse",
		Name: "responses",
	}
	err = dq.QueryResponses().
		Select(dnsblresponse.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[1] = &Edge{
		Type: "IP",
		Name: "ip_address",
	}
	err = dq.QueryIPAddress().
		Select(ip.FieldID).
		Scan(ctx, &node.Edges[1].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (dr *DNSBLResponse) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     dr.ID,
		Type:   "DNSBLResponse",
		Fields: make([]*Field, 2),
		Edges:  make([]*Edge, 1),
	}
	var buf []byte
	if buf, err = json.Marshal(dr.Code); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "string",
		Name:  "code",
		Value: string(buf),
	}
	if buf, err = json.Marshal(dr.Description); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "string",
		Name:  "description",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "DNSBLQuery",
		Name: "query",
	}
	err = dr.QueryQuery().
		Select(dnsblquery.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (i *IP) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     i.ID,
		Type:   "IP",
		Fields: make([]*Field, 1),
		Edges:  make([]*Edge, 1),
	}
	var buf []byte
	if buf, err = json.Marshal(i.IPAddress); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "string",
		Name:  "ip_address",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "DNSBLQuery",
		Name: "queries",
	}
	err = i.QueryQueries().
		Select(dnsblquery.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (c *Client) Node(ctx context.Context, id uuid.UUID) (*Node, error) {
	n, err := c.Noder(ctx, id)
	if err != nil {
		return nil, err
	}
	return n.Node(ctx)
}

var errNodeInvalidID = &NotFoundError{"node"}

// NodeOption allows configuring the Noder execution using functional options.
type NodeOption func(*NodeOptions)

// WithNodeType sets the Type of the node (i.e. the table to query).
// If was not provided, the table will be derived from the universal-id
// configuration as described in: https://entgo.io/docs/migrate/#universal-ids.
func WithNodeType(t string) NodeOption {
	return func(o *NodeOptions) {
		o.Type = t
	}
}

// NodeOptions holds the configuration for Noder execution.
type NodeOptions struct {
	// Type of the node (schema table).
	Type string
}

// Noder returns a Node by its id. If the NodeType was not provided, it will
// be derived from the id value according to the universal-id configuration.
//
//		c.Noder(ctx, id)
//		c.Noder(ctx, id, ent.WithNodeType(pet.Table))
//
func (c *Client) Noder(ctx context.Context, id uuid.UUID, opts ...NodeOption) (_ Noder, err error) {
	defer func() {
		if IsNotFound(err) {
			err = multierror.Append(err, entgql.ErrNodeNotFound(id))
		}
	}()
	options := &NodeOptions{}
	for _, opt := range opts {
		opt(options)
	}
	if options.Type == "" {
		return nil, fmt.Errorf("cannot resolve Noder (%v) without its type", id)
	}
	return c.noder(ctx, options.Type, id)
}

func (c *Client) noder(ctx context.Context, tbl string, id uuid.UUID) (Noder, error) {
	switch tbl {
	case dnsblquery.Table:
		n, err := c.DNSBLQuery.Query().
			Where(dnsblquery.ID(id)).
			CollectFields(ctx, "DNSBLQuery").
			Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case dnsblresponse.Table:
		n, err := c.DNSBLResponse.Query().
			Where(dnsblresponse.ID(id)).
			CollectFields(ctx, "DNSBLResponse").
			Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case ip.Table:
		n, err := c.IP.Query().
			Where(ip.ID(id)).
			CollectFields(ctx, "IP").
			Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	default:
		return nil, fmt.Errorf("cannot resolve Noder from table %q: %w", tbl, errNodeInvalidID)
	}
}
