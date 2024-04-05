package eBay

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func (api *API) GetOrders(client *http.Client) error {
	url := "https://api.sandbox.ebay.com/sell/fulfillment/v1/order?filter=creationdate%3A[2024-01-21T08%3A25%3A43.511Z..2024-02-21T08%3A25%3A43.511Z]%2Clastmodifieddate%3A[2024-02-15T08%3A25%3A43.511Z..]%2Corderfulfillmentstatus%3A{NOT_STARTED%7CIN_PROGRESS}"
	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// A successful call returns an HTTP status code of 204 and has no response payload.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%s", string(body))
	}

	return nil
}

// Doc: https://developer.ebay.com/api-docs/sell/inventory/resources/inventory_item/methods/createOrReplaceInventoryItem
func (api *API) CreateOrReplaceInventoryItem(client *http.Client) error {
	url := "https://api.sandbox.ebay.com/sell/inventory/v1/inventory_item/Iphone12"
	payload := strings.NewReader(`{
		"product": {
			"title": "Test listing - do not bid or buy - awesome Apple watch test 2",
			"aspects": {
				"Feature":[
				  "Water resistance", "GPS"
				],
				"CPU":[
				  "Dual-Core Processor"
				]
			},
			"description": "Test listing - do not bid or buy \n Built-in GPS. Water resistance to 50 meters.1 A new lightning-fast dual-core processor. And a display that\u2019s two times brighter than before. Full of features that help you stay active, motivated, and connected, Apple Watch Series 2 is designed for all the ways you move ",
			"upc": ["888462079525"],
			"imageUrls": [
				"http://store.storeimages.cdn-apple.com/4973/as-images.apple.com/is/image/AppleInc/aos/published/images/S/1/S1/42/S1-42-alu-silver-sport-white-grid?wid=332&hei=392&fmt=jpeg&qlt=95&op_sharpen=0&resMode=bicub&op_usm=0.5,0.5,0,0&iccEmbed=0&layer=comp&.v=1472247758975",
				"http://store.storeimages.cdn-apple.com/4973/as-images.apple.com/is/image/AppleInc/aos/published/images/4/2/42/stainless/42-stainless-sport-white-grid?wid=332&hei=392&fmt=jpeg&qlt=95&op_sharpen=0&resMode=bicub&op_usm=0.5,0.5,0,0&iccEmbed=0&layer=comp&.v=1472247760390",
				"http://store.storeimages.cdn-apple.com/4973/as-images.apple.com/is/image/AppleInc/aos/published/images/4/2/42/ceramic/42-ceramic-sport-cloud-grid?wid=332&hei=392&fmt=jpeg&qlt=95&op_sharpen=0&resMode=bicub&op_usm=0.5,0.5,0,0&iccEmbed=0&layer=comp&.v=1472247758007"
			]
		},
		"condition": "NEW",
		"packageWeightAndSize": {
			"dimensions": {
				"height": 5,
				"length": 10,
				"width": 15,
				"unit": "INCH"
			},
			"packageType": "MAILING_BOX",
			"weight": {
				"value": 2,
				"unit": "POUND"
			}
		},
		"availability": {
			"shipToLocationAvailability": {
				"quantity": 10
			}
		}
	}`)

	// Create a new request with PUT method and payload
	req, err := http.NewRequest("PUT", url, payload)
	if err != nil {
		return err
	}

	// Set Header
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Content-Language", "en-US")

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// A successful call returns an HTTP status code of 204 and has no response payload.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(body))

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("%s", string(body))
	}

	return nil
}
