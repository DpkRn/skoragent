// SC_AGENT_CUSTOMER_MAPPING
package models

import "gorm.io/gorm"

type SC_AGENT_CUSTOMER_MAPPING struct {
	gorm.Model
	Agent_ID       string
	Customer_phone string
	Customer_name  string
}
