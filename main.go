package main

import (
	"Rimedata/models"
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly/v2"
	"github.com/xuri/excelize/v2"
	"strconv"
	"strings"
	"time"
)

func main() {

	fIndex := 1
	SheetName := "Sheet1"

	f := initExcelFIle(fIndex, SheetName)

	fIndex++

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		// Header
		r.Headers.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySUQiOiIyYTMzOGY1Ny0yY2JhLTQyZDQtYTc5My0yNzBiOTlhYTU0NDkiLCJ0ZWFtSUQiOiIiLCJleHAiOjE2NTM5NjM4Mjl9.CC61lJgsJ4htqTAcVabtO35AJPhRgTujS2TjBrpUZg4")
		r.Headers.Set("content-type", "application/json")
		r.Headers.Set("Origin", "https://rimedata.com")
		r.Headers.Set("Referer", "https://rimedata.com/")
		r.Headers.Set("sec-ch-ua", "\"Chromium\";v=\"88\", \"Google Chrome\";v=\"88\", \";Not A Brand\";v=\"99\"")
		r.Headers.Set("sec-ch-ua-mobile", "?0")
		r.Headers.Set("Sec-Fetch-Dest", "empty")
		r.Headers.Set("Sec-Fetch-Mode", "cors")
		r.Headers.Set("Sec-Fetch-Site", "same-site")
		r.Headers.Set("TeamID", "a8f3b8d1-fa83-4fce-b717-94b5140084e5")
		r.Headers.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.182 Safari/537.36")
		//r.Headers.Set("", "")
		//r.Headers.Set("", "")

		r.Method = "POST"
		//payload
		// r.Body = payload
	})

	c.OnResponse(func(r *colly.Response) {
		jsonStr := []byte(r.Body)

		resp := new(models.RimeReponse)
		err := json.Unmarshal(jsonStr, resp)
		if err != nil {
			fmt.Println(err)
		}

		nodes := resp.Data.RealmSheet.Nodes

		for i := 0; i < len(nodes); i++ {

			//fmt.Println(nodes[i].Data)
			//fmt.Println(json.Valid([]byte(nodes[i].Data)))
			fmt.Println(fIndex)
			nodeData := new(models.NodeData)

			err := json.Unmarshal([]byte(nodes[i].Data), nodeData)
			if err != nil {
				fmt.Println(err)
			}

			// column_serial
			var cs []string
			for i := 0; i < len(nodeData.ColumnSerial); i++ {
				cs = append(cs, nodeData.ColumnSerial[i].Value)
			}
			css := strings.Join(cs, ";")
			fmt.Println(css)
			err = f.SetCellValue(SheetName, "A"+strconv.Itoa(fIndex), css)
			if err != nil {
				fmt.Println(err)
			}

			// primary_name
			var pn []string
			for i := 0; i < len(nodeData.PrimaryName); i++ {
				pn = append(pn, nodeData.PrimaryName[i].Value)
			}
			//pns := strings.Join(pn, ";")
			//fmt.Println(pns)

			// legal_name
			var ln []string
			for i := 0; i < len(nodeData.LegalName); i++ {
				ln = append(ln, nodeData.LegalName[i].Value)
			}
			lns := strings.Join(ln, ";")
			//fmt.Println(lns)
			err = f.SetCellValue(SheetName, "B"+strconv.Itoa(fIndex), lns)
			if err != nil {
				fmt.Println(err)
			}

			// industries
			var ind []string
			for i := 0; i < len(nodeData.Industries); i++ {
				ind = append(ind, nodeData.Industries[i].Value)
			}
			inds := strings.Join(ind, ";")
			//fmt.Println(inds)
			err = f.SetCellValue(SheetName, "C"+strconv.Itoa(fIndex), inds)
			if err != nil {
				fmt.Println(err)
			}

			// qualifications
			var qct []string
			for i := 0; i < len(nodeData.Qualifications); i++ {
				qct = append(qct, nodeData.Qualifications[i].Value)
			}
			qcts := strings.Join(qct, ";")
			//fmt.Println(qcts)
			err = f.SetCellValue(SheetName, "D"+strconv.Itoa(fIndex), qcts)
			if err != nil {
				fmt.Println(err)
			}

			// hard_tech_verticals
			var htv []string
			for i := 0; i < len(nodeData.HardTechVerticals); i++ {
				htv = append(htv, nodeData.HardTechVerticals[i].Value)
			}
			htvs := strings.Join(htv, ";")
			//fmt.Println(htvs)
			err = f.SetCellValue(SheetName, "E"+strconv.Itoa(fIndex), htvs)
			if err != nil {
				fmt.Println(err)
			}

			// verticals
			var vtc []string
			for i := 0; i < len(nodeData.Verticals); i++ {
				vtc = append(vtc, nodeData.Verticals[i].Value)
			}
			vtcs := strings.Join(vtc, ";")
			//fmt.Println(vtcs)
			err = f.SetCellValue(SheetName, "F"+strconv.Itoa(fIndex), vtcs)
			if err != nil {
				fmt.Println(err)
			}

			// hard_tech_industries
			var hti []string
			for i := 0; i < len(nodeData.HardTechIndustries); i++ {
				hti = append(hti, nodeData.HardTechIndustries[i].Value)
			}
			htis := strings.Join(hti, ";")
			//fmt.Println(htis)
			err = f.SetCellValue(SheetName, "G"+strconv.Itoa(fIndex), htis)
			if err != nil {
				fmt.Println(err)
			}

			// features
			var fte []string
			for i := 0; i < len(nodeData.Features); i++ {
				fte = append(fte, nodeData.Features[i].Value)
			}
			//ftes := strings.Join(fte, ";")
			//fmt.Println(ftes)

			//financing_stage
			var fst []string
			for i := 0; i < len(nodeData.FinancingStage); i++ {
				fst = append(fst, nodeData.FinancingStage[i].Value)
			}
			fsts := strings.Join(fst, ";")
			//fmt.Println(htis)
			err = f.SetCellValue(SheetName, "H"+strconv.Itoa(fIndex), fsts)
			if err != nil {
				fmt.Println(err)
			}

			//latest_deal_type
			var ldt []string
			for i := 0; i < len(nodeData.LatestDealType); i++ {
				ldt = append(ldt, nodeData.LatestDealType[i].Value)
			}
			ldts := strings.Join(ldt, ";")
			//fmt.Println(htis)
			err = f.SetCellValue(SheetName, "I"+strconv.Itoa(fIndex), ldts)
			if err != nil {
				fmt.Println(err)
			}

			fIndex++
		}

		fmt.Println(resp.Data.RealmSheet.PageInfo.EndCursor)
		time.Sleep(time.Duration(50) * time.Second)
		postRimeData(c, resp.Data.RealmSheet.PageInfo.EndCursor)
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
		if err := f.SaveAs("./docs/Book1.xlsx"); err != nil {
			fmt.Println(err)
		}
		err := f.Close()
		if err != nil {
			fmt.Println(err)
		}
	})

	postRimeData(c, "")

}

func postRimeData(c *colly.Collector, pageStart string) {
	var payloadStr = ""
	if pageStart == "" {
		payloadStr = `{"query":"query RealmSheet($realmID: String!, $sheetID: String!, $first: MaybeInt32, $after: MaybeString, $last: MaybeInt32, $before: MaybeString, $conditions: [NextSearchConditionInput]!, $columnIDs: [String!]!, $templateID: MaybeString, $orderColumns: [OrderColumn]) {\n  realmSheet(realmID: $realmID, sheetID: $sheetID, first: $first, after: $after, last: $last, before: $before, conditions: $conditions, columnIDs: $columnIDs, templateID: $templateID, orderColumns: $orderColumns) {\n    totalCount\n    aggregate {\n      column\n      average\n      median\n      summation\n      max\n      min\n      __typename\n    }\n    nodes {\n      entityID\n      entityType\n      data\n      id\n      __typename\n    }\n    pageInfo {\n      ...PageInfoFields\n      __typename\n    }\n    permissions {\n      ...PermissionFields\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment PageInfoFields on PageInfo {\n  hasPreviousPage\n  hasNextPage\n  startCursor\n  endCursor\n}\n\nfragment PermissionFields on Permission {\n  __typename\n  ... on FrequencyPermission {\n    used\n    total\n    remaining\n    __typename\n  }\n  ... on LockPermission {\n    locked\n    __typename\n  }\n  ... on PaginationPermission {\n    rowLimit\n    lockNextPage\n    lockPreviousPage\n    __typename\n  }\n}\n","operationName":"RealmSheet","variables":{"realmID":"pevc.enterprise","sheetID":"pevc.enterprise_all","columnIDs":["column_serial","primary_name","legal_name","industries","qualifications","hard_tech_verticals","verticals","hard_tech_industries","features","financing_stage","latest_deal_type"],"conditions":[{"id":"enterprise.region","name":"总部地区","type":"SEARCH_FIELD","operator":{"operator":"INCLUDES_ANY"},"options":[{"id":"0301010100","name":"上海市","values":["0301010100"],"operator":"INCLUDES_ANY"}]}],"orderColumns":[],"first":1000}}`

		//payloadStr = `{"query":"query RealmSheet($realmID: String!, $sheetID: String!, $first: MaybeInt32, $after: MaybeString, $last: MaybeInt32, $before: MaybeString, $conditions: [NextSearchConditionInput]!, $columnIDs: [String!]!, $templateID: MaybeString, $orderColumns: [OrderColumn]) {\n  realmSheet(realmID: $realmID, sheetID: $sheetID, first: $first, after: $after, last: $last, before: $before, conditions: $conditions, columnIDs: $columnIDs, templateID: $templateID, orderColumns: $orderColumns) {\n    totalCount\n    aggregate {\n      column\n      average\n      median\n      summation\n      max\n      min\n      __typename\n    }\n    nodes {\n      entityID\n      entityType\n      data\n      id\n      __typename\n    }\n    pageInfo {\n      ...PageInfoFields\n      __typename\n    }\n    permissions {\n      ...PermissionFields\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment PageInfoFields on PageInfo {\n  hasPreviousPage\n  hasNextPage\n  startCursor\n  endCursor\n}\n\nfragment PermissionFields on Permission {\n  __typename\n  ... on FrequencyPermission {\n    used\n    total\n    remaining\n    __typename\n  }\n  ... on LockPermission {\n    locked\n    __typename\n  }\n  ... on PaginationPermission {\n    rowLimit\n    lockNextPage\n    lockPreviousPage\n    __typename\n  }\n}\n","operationName":"RealmSheet","variables":{"realmID":"pevc.enterprise","sheetID":"pevc.enterprise_all","columnIDs":["column_serial","primary_name","legal_name","industries","qualifications","hard_tech_verticals","verticals","hard_tech_industries","features"],"conditions":[{"id":"enterprise.region","name":"总部地区","type":"SEARCH_FIELD","operator":{"operator":"INCLUDES_ANY"},"options":[{"id":"0301010100","name":"上海市","values":["0301010100"],"operator":"INCLUDES_ANY"}]},{"id":"enterprise.qualification","name":"企业资质","type":"CHECKBOX","operator":{"name":"并集","operator":"INCLUDES_ANY"},"options":[{"id":"specialized_and_new_enterprise","name":"专精特新企业","values":["108020116"],"operator":"INCLUDES_ANY"}]}],"orderColumns":[],"first":100}}`
	} else {
		payloadStr = `{"query":"query RealmSheet($realmID: String!, $sheetID: String!, $first: MaybeInt32, $after: MaybeString, $last: MaybeInt32, $before: MaybeString, $conditions: [NextSearchConditionInput]!, $columnIDs: [String!]!, $templateID: MaybeString, $orderColumns: [OrderColumn]) {\n  realmSheet(realmID: $realmID, sheetID: $sheetID, first: $first, after: $after, last: $last, before: $before, conditions: $conditions, columnIDs: $columnIDs, templateID: $templateID, orderColumns: $orderColumns) {\n    totalCount\n    aggregate {\n      column\n      average\n      median\n      summation\n      max\n      min\n      __typename\n    }\n    nodes {\n      entityID\n      entityType\n      data\n      id\n      __typename\n    }\n    pageInfo {\n      ...PageInfoFields\n      __typename\n    }\n    permissions {\n      ...PermissionFields\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment PageInfoFields on PageInfo {\n  hasPreviousPage\n  hasNextPage\n  startCursor\n  endCursor\n}\n\nfragment PermissionFields on Permission {\n  __typename\n  ... on FrequencyPermission {\n    used\n    total\n    remaining\n    __typename\n  }\n  ... on LockPermission {\n    locked\n    __typename\n  }\n  ... on PaginationPermission {\n    rowLimit\n    lockNextPage\n    lockPreviousPage\n    __typename\n  }\n}\n","operationName":"RealmSheet","variables":{"realmID":"pevc.enterprise","sheetID":"pevc.enterprise_all","columnIDs":["column_serial","primary_name","legal_name","industries","qualifications","hard_tech_verticals","verticals","hard_tech_industries","features","financing_stage","latest_deal_type"],"conditions":[{"id":"enterprise.region","name":"总部地区","type":"SEARCH_FIELD","operator":{"operator":"INCLUDES_ANY"},"options":[{"id":"0301010100","name":"上海市","values":["0301010100"],"operator":"INCLUDES_ANY"}]}],"orderColumns":[],"first":1000,"after":"` + pageStart + `"}}`

		//payloadStr = `{"query":"query RealmSheet($realmID: String!, $sheetID: String!, $first: MaybeInt32, $after: MaybeString, $last: MaybeInt32, $before: MaybeString, $conditions: [NextSearchConditionInput]!, $columnIDs: [String!]!, $templateID: MaybeString, $orderColumns: [OrderColumn]) {\n  realmSheet(realmID: $realmID, sheetID: $sheetID, first: $first, after: $after, last: $last, before: $before, conditions: $conditions, columnIDs: $columnIDs, templateID: $templateID, orderColumns: $orderColumns) {\n    totalCount\n    aggregate {\n      column\n      average\n      median\n      summation\n      max\n      min\n      __typename\n    }\n    nodes {\n      entityID\n      entityType\n      data\n      id\n      __typename\n    }\n    pageInfo {\n      ...PageInfoFields\n      __typename\n    }\n    permissions {\n      ...PermissionFields\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment PageInfoFields on PageInfo {\n  hasPreviousPage\n  hasNextPage\n  startCursor\n  endCursor\n}\n\nfragment PermissionFields on Permission {\n  __typename\n  ... on FrequencyPermission {\n    used\n    total\n    remaining\n    __typename\n  }\n  ... on LockPermission {\n    locked\n    __typename\n  }\n  ... on PaginationPermission {\n    rowLimit\n    lockNextPage\n    lockPreviousPage\n    __typename\n  }\n}\n","operationName":"RealmSheet","variables":{"realmID":"pevc.enterprise","sheetID":"pevc.enterprise_all","columnIDs":["column_serial","primary_name","legal_name","industries","qualifications","hard_tech_verticals","verticals","hard_tech_industries","features"],"conditions":[{"id":"enterprise.region","name":"总部地区","type":"SEARCH_FIELD","operator":{"operator":"INCLUDES_ANY"},"options":[{"id":"0301010100","name":"上海市","values":["0301010100"],"operator":"INCLUDES_ANY"}]},{"id":"enterprise.qualification","name":"企业资质","type":"CHECKBOX","operator":{"name":"并集","operator":"INCLUDES_ANY"},"options":[{"id":"specialized_and_new_enterprise","name":"专精特新企业","values":["108020116"],"operator":"INCLUDES_ANY"}]}],"orderColumns":[],"first":100,"after":"`+pageStart+`"}}`
	}

	payload := []byte(payloadStr)

	err := c.PostRaw("https://rimedata.com:9998/v2/graphql", payload)
	if err != nil {
		fmt.Println(err)
	}
}

func initExcelFIle(fIndex int, SheetName string) *excelize.File {
	f := excelize.NewFile()
	s1Index := f.NewSheet(SheetName)

	f.SetCellValue(SheetName, "A"+strconv.Itoa(fIndex), "行号")
	f.SetCellValue(SheetName, "B"+strconv.Itoa(fIndex), "企业全称")
	f.SetCellValue(SheetName, "C"+strconv.Itoa(fIndex), "行业领域")
	f.SetCellValue(SheetName, "D"+strconv.Itoa(fIndex), "企业资质")
	f.SetCellValue(SheetName, "E"+strconv.Itoa(fIndex), "硬科技赛道")
	f.SetCellValue(SheetName, "F"+strconv.Itoa(fIndex), "细分赛道")
	f.SetCellValue(SheetName, "G"+strconv.Itoa(fIndex), "硬科技领域")

	f.SetCellValue(SheetName, "H"+strconv.Itoa(fIndex), "融资阶段")
	f.SetCellValue(SheetName, "I"+strconv.Itoa(fIndex), "最近融资类型")

	f.SetActiveSheet(s1Index)

	return f
}
