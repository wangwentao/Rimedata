package models

type RimeReponse struct {
	Data struct {
		RealmSheet struct {
			TotalCount int         `json:"totalCount"`
			Aggregate  interface{} `json:"aggregate"`
			Nodes      []Node      `json:"nodes"`
			PageInfo   struct {
				HasPreviousPage bool   `json:"hasPreviousPage"`
				HasNextPage     bool   `json:"hasNextPage"`
				StartCursor     string `json:"startCursor"`
				EndCursor       string `json:"endCursor"`
				Typename        string `json:"__typename"`
			} `json:"pageInfo"`
			Permissions []struct {
				Typename  string `json:"__typename"`
				Used      int    `json:"used"`
				Total     int    `json:"total"`
				Remaining int    `json:"remaining"`
			} `json:"permissions"`
			Typename string `json:"__typename"`
		} `json:"realmSheet"`
	} `json:"data"`
}

type Node struct {
	EntityID   string      `json:"entityID"`
	EntityType string      `json:"entityType"`
	Data       string      `json:"data"`
	Id         interface{} `json:"id"`
	Typename   string      `json:"__typename"`
}

type ResponseNodeData struct {
	ColumnSerial []struct {
		Pattern string `json:"pattern"`
		Value   string `json:"value"`
	} `json:"column_serial"`
	Features []struct {
		Pattern string `json:"pattern"`
		Value   string `json:"value"`
	} `json:"features"`
	HardTechIndustries []struct {
		Pattern string `json:"pattern"`
		Value   string `json:"value,omitempty"`
		Modal   struct {
			Type string `json:"type"`
			Args struct {
				EntityId   string `json:"entity_id"`
				EntityType string `json:"entity_type"`
			} `json:"args"`
		} `json:"modal,omitempty"`
		Action struct {
			Type string `json:"type"`
			Args struct {
				EntityId   string `json:"entity_id"`
				EntityType string `json:"entity_type"`
			} `json:"args"`
		} `json:"action,omitempty"`
	} `json:"hard_tech_industries"`
	HardTechVerticals []struct {
		Pattern string `json:"pattern"`
		Value   string `json:"value,omitempty"`
		Modal   struct {
			Type string `json:"type"`
			Args struct {
				EntityId   string `json:"entity_id"`
				EntityType string `json:"entity_type"`
			} `json:"args"`
		} `json:"modal,omitempty"`
		Action struct {
			Type string `json:"type"`
			Args struct {
				EntityId   string `json:"entity_id"`
				EntityType string `json:"entity_type"`
			} `json:"args"`
		} `json:"action,omitempty"`
	} `json:"hard_tech_verticals"`
	Industries []struct {
		Pattern string `json:"pattern"`
		Value   string `json:"value,omitempty"`
		Modal   struct {
			Type string `json:"type"`
			Args struct {
				EntityId   string `json:"entity_id"`
				EntityType string `json:"entity_type"`
			} `json:"args"`
		} `json:"modal,omitempty"`
		Action struct {
			Type string `json:"type"`
			Args struct {
				EntityId   string `json:"entity_id"`
				EntityType string `json:"entity_type"`
			} `json:"args"`
		} `json:"action,omitempty"`
	} `json:"industries"`
	LegalName []struct {
		Pattern string `json:"pattern"`
		Value   string `json:"value"`
		Modal   struct {
			Type string `json:"type"`
			Args struct {
				EntityId   string `json:"entity_id"`
				EntityType string `json:"entity_type"`
			} `json:"args"`
		} `json:"modal"`
		Action struct {
			Type string `json:"type"`
			Args struct {
				EntityId   string `json:"entity_id"`
				EntityType string `json:"entity_type"`
			} `json:"args"`
		} `json:"action"`
	} `json:"legal_name"`
	PrimaryName []struct {
		Pattern string `json:"pattern"`
		Value   string `json:"value"`
		Modal   struct {
			Type string `json:"type"`
			Args struct {
				EntityId   string `json:"entity_id"`
				EntityType string `json:"entity_type"`
			} `json:"args"`
		} `json:"modal,omitempty"`
		Action struct {
			Type string `json:"type"`
			Args struct {
				EntityId   string `json:"entity_id"`
				EntityType string `json:"entity_type"`
			} `json:"args"`
		} `json:"action,omitempty"`
	} `json:"primary_name"`
	Qualifications []struct {
		Pattern string `json:"pattern"`
		Value   string `json:"value"`
	} `json:"qualifications"`
	Verticals []struct {
		Pattern string `json:"pattern"`
		Value   string `json:"value,omitempty"`
		Modal   struct {
			Type string `json:"type"`
			Args struct {
				EntityId   string `json:"entity_id"`
				EntityType string `json:"entity_type"`
			} `json:"args"`
		} `json:"modal,omitempty"`
		Action struct {
			Type string `json:"type"`
			Args struct {
				EntityId   string `json:"entity_id"`
				EntityType string `json:"entity_type"`
			} `json:"args"`
		} `json:"action,omitempty"`
	} `json:"verticals"`
}

type NodeData struct {
	ColumnSerial []struct {
		Pattern string `json:"pattern"`
		Value   string `json:"value"`
	} `json:"column_serial"`
	Features []struct {
		Pattern string `json:"pattern"`
		Value   string `json:"value"`
	} `json:"features"`
	FinancingStage []struct {
		Pattern string `json:"pattern"`
		Value   string `json:"value"`
	} `json:"financing_stage"`
	HardTechIndustries []struct {
		Pattern string `json:"pattern"`
		Value   string `json:"value,omitempty"`
		Modal   struct {
			Type string `json:"type"`
			Args struct {
				EntityId   string `json:"entity_id"`
				EntityType string `json:"entity_type"`
			} `json:"args"`
		} `json:"modal,omitempty"`
		Action struct {
			Type string `json:"type"`
			Args struct {
				EntityId   string `json:"entity_id"`
				EntityType string `json:"entity_type"`
			} `json:"args"`
		} `json:"action,omitempty"`
	} `json:"hard_tech_industries"`
	HardTechVerticals []struct {
		Pattern string `json:"pattern"`
		Value   string `json:"value,omitempty"`
		Modal   struct {
			Type string `json:"type"`
			Args struct {
				EntityId   string `json:"entity_id"`
				EntityType string `json:"entity_type"`
			} `json:"args"`
		} `json:"modal,omitempty"`
		Action struct {
			Type string `json:"type"`
			Args struct {
				EntityId   string `json:"entity_id"`
				EntityType string `json:"entity_type"`
			} `json:"args"`
		} `json:"action,omitempty"`
	} `json:"hard_tech_verticals"`
	Industries []struct {
		Pattern string `json:"pattern"`
		Value   string `json:"value,omitempty"`
		Modal   struct {
			Type string `json:"type"`
			Args struct {
				EntityId   string `json:"entity_id"`
				EntityType string `json:"entity_type"`
			} `json:"args"`
		} `json:"modal,omitempty"`
		Action struct {
			Type string `json:"type"`
			Args struct {
				EntityId   string `json:"entity_id"`
				EntityType string `json:"entity_type"`
			} `json:"args"`
		} `json:"action,omitempty"`
	} `json:"industries"`
	LatestDealType []struct {
		Pattern string `json:"pattern"`
		Value   string `json:"value"`
		Modal   struct {
			Type string `json:"type"`
			Args struct {
				EntityId   string `json:"entity_id"`
				EntityType string `json:"entity_type"`
			} `json:"args"`
		} `json:"modal"`
		Action struct {
			Type string `json:"type"`
			Args struct {
				EntityId   string `json:"entity_id"`
				EntityType string `json:"entity_type"`
			} `json:"args"`
		} `json:"action"`
	} `json:"latest_deal_type"`
	LegalName []struct {
		Pattern string `json:"pattern"`
		Value   string `json:"value"`
		Modal   struct {
			Type string `json:"type"`
			Args struct {
				EntityId   string `json:"entity_id"`
				EntityType string `json:"entity_type"`
			} `json:"args"`
		} `json:"modal"`
		Action struct {
			Type string `json:"type"`
			Args struct {
				EntityId   string `json:"entity_id"`
				EntityType string `json:"entity_type"`
			} `json:"args"`
		} `json:"action"`
	} `json:"legal_name"`
	PrimaryName []struct {
		Pattern string `json:"pattern"`
		Value   string `json:"value"`
		Modal   struct {
			Type string `json:"type"`
			Args struct {
				EntityId   string `json:"entity_id"`
				EntityType string `json:"entity_type"`
			} `json:"args"`
		} `json:"modal,omitempty"`
		Action struct {
			Type string `json:"type"`
			Args struct {
				EntityId   string `json:"entity_id"`
				EntityType string `json:"entity_type"`
			} `json:"args"`
		} `json:"action,omitempty"`
	} `json:"primary_name"`
	Qualifications []struct {
		Pattern string `json:"pattern"`
		Value   string `json:"value"`
	} `json:"qualifications"`
	Verticals []struct {
		Pattern string `json:"pattern"`
		Value   string `json:"value,omitempty"`
		Modal   struct {
			Type string `json:"type"`
			Args struct {
				EntityId   string `json:"entity_id"`
				EntityType string `json:"entity_type"`
			} `json:"args"`
		} `json:"modal,omitempty"`
		Action struct {
			Type string `json:"type"`
			Args struct {
				EntityId   string `json:"entity_id"`
				EntityType string `json:"entity_type"`
			} `json:"args"`
		} `json:"action,omitempty"`
	} `json:"verticals"`
}
