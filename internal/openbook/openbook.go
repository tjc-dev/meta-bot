package openbook

import (
	"context"
	"errors"
	"math/big"
	"metabot/internal/idl/openbook_v2"

	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

var (
	InnerNodeTag = uint8(1)
	LeafNodeTag  = uint8(2)
)

type OpenBook struct {
	RPC    *rpc.Client
	PubKey solana.PublicKey
}

type Side int

const (
	Bid Side = iota
	Ask
)

// Order represents an order in the book side.
type Order struct {
	Market         *openbook_v2.Market
	Side           Side
	PriceLots      *big.Int
	SizeLots       *big.Int
	SeqNum         *big.Int
	IsExpired      bool
	IsOraclePegged bool
	LeafNode       *openbook_v2.LeafNode // Add this line
}

// NewOrder creates a new Order instance
func NewOrder(market *openbook_v2.Market, leafNode *openbook_v2.LeafNode, side Side, isExpired, isOraclePegged bool) (*Order, error) {
	order := &Order{
		Market:         market,
		LeafNode:       leafNode,
		Side:           side,
		IsExpired:      isExpired,
		IsOraclePegged: isOraclePegged,
	}

	U64MaxBn := big.NewInt(0).SetUint64(^uint64(0))
	leafNodeKeyBigInt := big.NewInt(0).SetBytes(leafNode.Key.Bytes())

	// Masking the first 64 bits of the leafNode.Key for SeqNum
	mask := big.NewInt(1).Lsh(big.NewInt(1), 64).Sub(big.NewInt(1), big.NewInt(1))
	if side == Bid {
		order.SeqNum = big.NewInt(0).Sub(U64MaxBn, big.NewInt(0).And(leafNodeKeyBigInt, mask))
	} else { // Ask
		order.SeqNum = big.NewInt(0).And(leafNodeKeyBigInt, mask)
	}

	// Extracting the next 64 bits for PriceLots
	priceData := big.NewInt(0).Rsh(leafNodeKeyBigInt, 64)
	if isOraclePegged {
		// Oracle pegged logic not implemented
		return nil, errors.New("oracle pegged price logic not implemented")
	} else {
		order.PriceLots = priceData
	}

	return order, nil
}

type LocalBookSide struct {
	ClusterTime big.Int
	Market      *openbook_v2.Market
	*openbook_v2.BookSide
}

func NewLocalBookSide(clusterTime big.Int, market *openbook_v2.Market, bookSide *openbook_v2.BookSide) LocalBookSide {
	return LocalBookSide{
		ClusterTime: clusterTime,
		Market:      market,
		BookSide:    bookSide,
	}
}

func (ob *OpenBook) GetMarket(address solana.PublicKey) (market openbook_v2.Market, err error) {
	resp, err := ob.RPC.GetAccountInfo(
		context.TODO(),
		address,
	)
	if err != nil {
		return
	}

	err = bin.NewBorshDecoder(resp.GetBinary()).Decode(&market)
	if err != nil {
		return
	}
	return
}
func NewOpenBook(rpc *rpc.Client, programAddress solana.PublicKey) OpenBook {
	return OpenBook{
		RPC:    rpc,
		PubKey: programAddress,
	}
}

func (ob *OpenBook) GetSide(address solana.PublicKey) (side openbook_v2.BookSide, err error) {
	resp, err := ob.RPC.GetAccountInfo(
		context.TODO(),
		address,
	)
	if err != nil {
		return
	}

	err = bin.NewBorshDecoder(resp.GetBinary()).Decode(&side)
	if err != nil {
		return
	}
	return
}

func (bs *LocalBookSide) FixedItems() []*Order {
	var orders []*Order

	if bs.Roots[0].LeafCount == 0 {
		return orders // Return empty if no items
	}

	stack := []uint32{bs.Roots[0].MaybeNode}
	var left, right uint32
	left, right = 0, 1

	for len(stack) > 0 {
		index := stack[len(stack)-1]
		stack = stack[:len(stack)-1] // Pop

		node := bs.Nodes.Nodes[index]
		if node.Tag == InnerNodeTag {
			nodeDataWithTag := append([]byte{InnerNodeTag}, node.Data[:]...)
			innerNode, err := bs.ToInnerNode(nodeDataWithTag)
			if err != nil {
				return orders
			}
			stack = append(stack, innerNode.Children[right], innerNode.Children[left])
		} else if node.Tag == LeafNodeTag {
			nodeDataWithTag := append([]byte{LeafNodeTag}, node.Data[:]...)
			leafNode, err := bs.ToLeafNode(nodeDataWithTag)
			if err != nil {
				return orders
			}
			expiryTimestamp := big.NewInt(0)
			if leafNode.TimeInForce != 0 {
				expiryTimestamp = new(big.Int).Add(big.NewInt(int64(leafNode.Timestamp)), big.NewInt(int64(leafNode.TimeInForce)))
			}
			isExpired := bs.ClusterTime.Cmp(expiryTimestamp) > 0
			order, err := NewOrder(bs.Market, leafNode, 1, isExpired, false)
			if err != nil {
				panic(err)
			}
			orders = append(orders, order)
		}
	}

	return orders
}

func (bs *LocalBookSide) ToInnerNode(data []byte) (*openbook_v2.InnerNode, error) {
	var innerNode openbook_v2.InnerNode
	decoder := bin.NewBorshDecoder(data)
	if err := decoder.Decode(&innerNode); err != nil {
		return nil, err
	}
	return &innerNode, nil
}

func (bs *LocalBookSide) ToLeafNode(data []byte) (*openbook_v2.LeafNode, error) {
	var leafNode openbook_v2.LeafNode
	decoder := bin.NewBorshDecoder(data)
	if err := decoder.Decode(&leafNode); err != nil {
		return nil, err
	}
	return &leafNode, nil
}
