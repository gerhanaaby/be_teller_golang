package models

type HostInqRequest struct {
	Branch     string `json:"branch"`
	TerminalID string `json:"terminal_ID"`
	UserID     string `json:"user_ID"`
	ToAccount  string `json:"to_account"`
	MataUang   string `json:"mata_uang"`

	TrxAmount    string `json:"trx_amount"`
	SettleAmount string `json:"settle_amount"`
	TargetAmount string `json:"target_amount"`
	ReffNo       string `json:"reff_no"`
	Message      string `json:"message"`

	SupervisorID  string `json:"supervisor_ID"`
	AuditNo       string `json:"audit_No"`
	ChargesAmount string `json:"charges_amount"`
	Overbooking   string `json:"overbooking"`
	KeyApp        string `json:"key_app"`

	Rfnum            string `json:"rfnum"`
	CreditName       string `json:"credit_name"`
	CreditBranchCode string `json:"credit_branch_code"`
	DebitBranchCode  string `json:"Debit_branch_code"`
	Ftopt            string `json:"ftopt"`
}

type HostInq struct {
	AcquiringInstitutionIdentificationCode string                 `json:"acquiring_institution_identification_code"`
	CardAcceptorIdentificationCode         string                 `json:"card_acceptor_identification_code"`
	CardAcceptorName                       string                 `json:"card_acceptor_name"`
	CardAcceptorTerminalIdentification     string                 `json:"card_acceptor_terminal_identification"`
	DateSettlement                         string                 `json:"date_settlement"`
	PointOfServiceEntryMode                string                 `json:"point_of_service_entry_mode"`
	ProcessingCode                         string                 `json:"processing_code"`
	PxObjClass                             string                 `json:"pxObjClass"`
	ReceivingInstitutionIdentificationCode string                 `json:"receiving_institution_identification_code"`
	ReservedForISOUse1                     string                 `json:"reserved_for_iso_use_1"`
	ReservedForPrivateUse                  string                 `json:"reserved_for_private_use"`
	ResponseCode                           string                 `json:"response_code"`
	RetrievalReferenceNumber               string                 `json:"retrieval_reference_number"`
	SystemsTraceAuditNumber                string                 `json:"systems_trace_audit_number"`
	ToAccount                              string                 `json:"to_account"`
	TransmissionDateTime                   string                 `json:"transmission_date_time"`
	ReservedForPrivateUse1                 ReservedForPrivateUse1 `json:"reserved_for_private_use_1"`
}

type ReservedForPrivateUse1 struct {
	CreditAccountType            string `json:"CreditAccountType"`
	CreditAvailableBalance       string `json:"CreditAvailableBalance"`
	CreditBranchCode             string `json:"CreditBranchCode"`
	CreditBranchDescription      string `json:"CreditBranchDescription"`
	CreditChequeStatus           string `json:"CreditChequeStatus"`
	CreditCurrency               string `json:"CreditCurrency"`
	CreditDealingConvertedAmount string `json:"CreditDealingConvertedAmount"`
	CreditDealingCurrency        string `json:"CreditDealingCurrency"`
	CreditDealingRate            string `json:"CreditDealingRate"`
	CreditName                   string `json:"CreditName"`
	CreditProductCode            string `json:"CreditProductCode"`
	CreditProductDescription     string `json:"CreditProductDescription"`
	CreditReservedAlphaField1    string `json:"CreditReservedAlphaField1"`
	CreditReservedAlphaField2    string `json:"CreditReservedAlphaField2"`
	CreditReservedFillerField    string `json:"CreditReservedFillerField"`
	CreditReservedNumericField1  string `json:"CreditReservedNumericField1"`
	CreditReservedNumericField2  string `json:"CreditReservedNumericField2"`
	CreditReturnCode             string `json:"CreditReturnCode"`
	DebitAccountType             string `json:"DebitAccountType"`
	DebitAvailableBalance        string `json:"DebitAvailableBalance"`
	DebitBranchCode              string `json:"DebitBranchCode"`
	DebitBranchDescription       string `json:"DebitBranchDescription"`
	DebitChequeStatus            string `json:"DebitChequeStatus"`
	DebitCurrency                string `json:"DebitCurrency"`
	DebitDealingConvertedAmount  string `json:"DebitDealingConvertedAmount"`
	DebitDealingCurrency         string `json:"DebitDealingCurrency"`
	DebitDealingRate             string `json:"DebitDealingRate"`
	DebitName                    string `json:"DebitName"`
	DebitProductCode             string `json:"DebitProductCode"`
	DebitProductDescription      string `json:"DebitProductDescription"`
	DebitReservedAlphaField1     string `json:"DebitReservedAlphaField1"`
	DebitReservedAlphaField2     string `json:"DebitReservedAlphaField2"`
	DebitReservedFillerField     string `json:"DebitReservedFillerField"`
	DebitReservedNumericField1   string `json:"DebitReservedNumericField1"`
	DebitReservedNumericField2   string `json:"DebitReservedNumericField2"`
	DebitReturnCode              string `json:"DebitReturnCode"`
	PxObjClass                   string `json:"pxObjClass"`
}
