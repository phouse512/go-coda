package coda

import (
	"fmt"
	"log"
)

type Formula struct {
	Id    string      `json:"id"`
	Type  string      `json:"type"`
	Href  string      `json:"href"`
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
}

type ListFormulasResponse struct {
	Formulas []Formula `json:"items"`
	PaginationResponse
}

type GetFormulaResponse struct {
	Formula
}

func (c *Client) ListFormulas(docId string, paginationPayload PaginationPayload) (ListFormulasResponse, error) {
	docPath := fmt.Sprintf("docs/%s/formulas", docId)
	var formulasResp ListFormulasResponse
	err := c.apiCall("GET", docPath, paginationPayload, &formulasResp)
	if err != nil {
		log.Print("Unable to make api call with error.")
	}
	return formulasResp, err
}

func (c *Client) GetFormula(docId string, formulaIdOrName string) (GetFormulaResponse, error) {
	docPath := fmt.Sprintf("docs/%s/formulas/%s", docId, formulaIdOrName)
	var formulaResp GetFormulaResponse
	err := c.apiCall("GET", docPath, nil, &formulaResp)
	if err != nil {
		log.Print("Unable to make api call with error.")
	}
	return formulaResp, err
}
