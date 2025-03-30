package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/hyperledger/fabric-chaincode-go/pkg/cid"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type ProductContract struct {
	contractapi.Contract
}

type Product struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	State  string `json:"state"`
	Owner  string `json:"owner"`
}

// CreateProduct permite a Org1 registrar un producto en la red con estado "fabricado"
func (c *ProductContract) CreateProduct(ctx contractapi.TransactionContextInterface, id string, name string) error {
	mspID, err := cid.GetMSPID(ctx.GetStub())
	if err != nil {
		return fmt.Errorf("error getting MSP ID: %v", err)
	}
	if mspID != "Org1MSP" {
		return errors.New("only Org1 can create products")
	}

	product := Product{
		ID:    id,
		Name:  name,
		State: "fabricado",
		Owner: "Org1",
	}

	productJSON, err := json.Marshal(product)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, productJSON)
}

// UpdateProductState permite a Org2 y Org3 actualizar el estado del producto
func (c *ProductContract) UpdateProductState(ctx contractapi.TransactionContextInterface, id string, newState string) error {
	productJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return fmt.Errorf("failed to read product: %v", err)
	}
	if productJSON == nil {
		return errors.New("product does not exist")
	}

	var product Product
	if err := json.Unmarshal(productJSON, &product); err != nil {
		return err
	}

	mspID, err := cid.GetMSPID(ctx.GetStub())
	if err != nil {
		return fmt.Errorf("error getting MSP ID: %v", err)
	}

	if (mspID == "Org2MSP" && product.State == "fabricado" && newState == "en transito") ||
		(mspID == "Org3MSP" && product.State == "en transito" && newState == "Entregado al minorista") {
		product.State = newState
		product.Owner = mspID
	} else {
		return errors.New("invalid state transition or unauthorized organization")
	}

	updatedProductJSON, err := json.Marshal(product)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, updatedProductJSON)
}
