// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package moka

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson11d1a9baDecodeGithubComEmirmuminogluMokaposGo(in *jlexer.Lexer, out *Request) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "PaymentDealrAuthentication":
			easyjson11d1a9baDecodeGithubComEmirmuminogluMokaposGo1(in, &out.PaymentDealerAuthentication)
		case "PaymentDealerRequest":
			easyjson11d1a9baDecodeGithubComEmirmuminogluMokaposGo2(in, &out.PaymentDealerRequest)
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson11d1a9baEncodeGithubComEmirmuminogluMokaposGo(out *jwriter.Writer, in Request) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"PaymentDealrAuthentication\":"
		out.RawString(prefix[1:])
		easyjson11d1a9baEncodeGithubComEmirmuminogluMokaposGo1(out, in.PaymentDealerAuthentication)
	}
	{
		const prefix string = ",\"PaymentDealerRequest\":"
		out.RawString(prefix)
		easyjson11d1a9baEncodeGithubComEmirmuminogluMokaposGo2(out, in.PaymentDealerRequest)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Request) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson11d1a9baEncodeGithubComEmirmuminogluMokaposGo(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Request) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson11d1a9baEncodeGithubComEmirmuminogluMokaposGo(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Request) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson11d1a9baDecodeGithubComEmirmuminogluMokaposGo(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Request) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson11d1a9baDecodeGithubComEmirmuminogluMokaposGo(l, v)
}
func easyjson11d1a9baDecodeGithubComEmirmuminogluMokaposGo2(in *jlexer.Lexer, out *PaymentDealerRequest) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "VirtualPosOrderId":
			out.VirtualPosOrderID = string(in.String())
		case "CardHolderFullName":
			out.CardHolderFullName = string(in.String())
		case "CardNumber":
			out.CardNumber = string(in.String())
		case "ExpMonth":
			out.ExpMonth = string(in.String())
		case "ExpYear":
			out.ExpYear = string(in.String())
		case "CvcNumber":
			out.CvcNumber = string(in.String())
		case "Amount":
			out.Amount = uint(in.Uint())
		case "Currency":
			out.Currency = string(in.String())
		case "InstallmentNumber":
			out.InstallmentNumber = int8(in.Int8())
		case "ClientIP":
			out.ClientIP = string(in.String())
		case "RedirectUrl":
			out.RedirectURL = string(in.String())
		case "RedirectType":
			out.RedirectType = int8(in.Int8())
		case "OtherTrxCode":
			out.OtherTrxCode = string(in.String())
		case "IsPreAuth":
			out.IsPreAuth = int8(in.Int8())
		case "IsPoolPayment":
			out.IsPoolPayment = int8(in.Int8())
		case "IsTokenized":
			out.IsTokenized = int8(in.Int8())
		case "IntegratorId":
			out.IntegratorID = int8(in.Int8())
		case "Software":
			out.Software = string(in.String())
		case "SubMerchantName":
			out.SubMerchantName = string(in.String())
		case "Description":
			out.Description = string(in.String())
		case "BuyerInformation":
			easyjson11d1a9baDecodeGithubComEmirmuminogluMokaposGo3(in, &out.BuyerInformation)
		case "CustomerInformation":
			easyjson11d1a9baDecodeGithubComEmirmuminogluMokaposGo4(in, &out.CustomerInformation)
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson11d1a9baEncodeGithubComEmirmuminogluMokaposGo2(out *jwriter.Writer, in PaymentDealerRequest) {
	out.RawByte('{')
	first := true
	_ = first
	if in.VirtualPosOrderID != "" {
		const prefix string = ",\"VirtualPosOrderId\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.VirtualPosOrderID))
	}
	if in.CardHolderFullName != "" {
		const prefix string = ",\"CardHolderFullName\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.CardHolderFullName))
	}
	if in.CardNumber != "" {
		const prefix string = ",\"CardNumber\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.CardNumber))
	}
	if in.ExpMonth != "" {
		const prefix string = ",\"ExpMonth\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ExpMonth))
	}
	if in.ExpYear != "" {
		const prefix string = ",\"ExpYear\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ExpYear))
	}
	if in.CvcNumber != "" {
		const prefix string = ",\"CvcNumber\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.CvcNumber))
	}
	if in.Amount != 0 {
		const prefix string = ",\"Amount\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Uint(uint(in.Amount))
	}
	if in.Currency != "" {
		const prefix string = ",\"Currency\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Currency))
	}
	if in.InstallmentNumber != 0 {
		const prefix string = ",\"InstallmentNumber\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int8(int8(in.InstallmentNumber))
	}
	if in.ClientIP != "" {
		const prefix string = ",\"ClientIP\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ClientIP))
	}
	if in.RedirectURL != "" {
		const prefix string = ",\"RedirectUrl\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.RedirectURL))
	}
	if in.RedirectType != 0 {
		const prefix string = ",\"RedirectType\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int8(int8(in.RedirectType))
	}
	if in.OtherTrxCode != "" {
		const prefix string = ",\"OtherTrxCode\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.OtherTrxCode))
	}
	if in.IsPreAuth != 0 {
		const prefix string = ",\"IsPreAuth\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int8(int8(in.IsPreAuth))
	}
	if in.IsPoolPayment != 0 {
		const prefix string = ",\"IsPoolPayment\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int8(int8(in.IsPoolPayment))
	}
	if in.IsTokenized != 0 {
		const prefix string = ",\"IsTokenized\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int8(int8(in.IsTokenized))
	}
	if in.IntegratorID != 0 {
		const prefix string = ",\"IntegratorId\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int8(int8(in.IntegratorID))
	}
	if in.Software != "" {
		const prefix string = ",\"Software\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Software))
	}
	if in.SubMerchantName != "" {
		const prefix string = ",\"SubMerchantName\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.SubMerchantName))
	}
	if in.Description != "" {
		const prefix string = ",\"Description\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Description))
	}
	if true{
		const prefix string = ",\"BuyerInformation\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjson11d1a9baEncodeGithubComEmirmuminogluMokaposGo3(out, in.BuyerInformation)
	}
	if true{
		const prefix string = ",\"CustomerInformation\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjson11d1a9baEncodeGithubComEmirmuminogluMokaposGo4(out, in.CustomerInformation)
	}
	out.RawByte('}')
}
func easyjson11d1a9baDecodeGithubComEmirmuminogluMokaposGo4(in *jlexer.Lexer, out *CustomerInformation) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "DealerCustomerId":
			out.DealerCustomerID = int(in.Int())
		case "CustomerCode":
			out.CustomerCode = string(in.String())
		case "FirstName":
			out.FirstName = string(in.String())
		case "LastName":
			out.LastName = string(in.String())
		case "Gender":
			out.Gender = int8(in.Int8())
		case "BirthDate":
			out.BirthDate = string(in.String())
		case "GsmNumber":
			out.GsmNumber = string(in.String())
		case "Email":
			out.EMail = string(in.String())
		case "Address":
			out.Address = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson11d1a9baEncodeGithubComEmirmuminogluMokaposGo4(out *jwriter.Writer, in CustomerInformation) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"DealerCustomerId\":"
		out.RawString(prefix[1:])
		out.Int(int(in.DealerCustomerID))
	}
	{
		const prefix string = ",\"CustomerCode\":"
		out.RawString(prefix)
		out.String(string(in.CustomerCode))
	}
	{
		const prefix string = ",\"FirstName\":"
		out.RawString(prefix)
		out.String(string(in.FirstName))
	}
	{
		const prefix string = ",\"LastName\":"
		out.RawString(prefix)
		out.String(string(in.LastName))
	}
	if in.Gender != 0 {
		const prefix string = ",\"Gender\":"
		out.RawString(prefix)
		out.Int8(int8(in.Gender))
	}
	if in.BirthDate != "" {
		const prefix string = ",\"BirthDate\":"
		out.RawString(prefix)
		out.String(string(in.BirthDate))
	}
	if in.GsmNumber != "" {
		const prefix string = ",\"GsmNumber\":"
		out.RawString(prefix)
		out.String(string(in.GsmNumber))
	}
	if in.EMail != "" {
		const prefix string = ",\"Email\":"
		out.RawString(prefix)
		out.String(string(in.EMail))
	}
	if in.Address != "" {
		const prefix string = ",\"Address\":"
		out.RawString(prefix)
		out.String(string(in.Address))
	}
	out.RawByte('}')
}
func easyjson11d1a9baDecodeGithubComEmirmuminogluMokaposGo3(in *jlexer.Lexer, out *BuyerInformation) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "BuyerFullName":
			out.FullName = string(in.String())
		case "BuyerEmail":
			out.EMail = string(in.String())
		case "BuyerGsmNumber":
			out.GSMNumber = string(in.String())
		case "BuyerAddress":
			out.Address = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson11d1a9baEncodeGithubComEmirmuminogluMokaposGo3(out *jwriter.Writer, in BuyerInformation) {
	out.RawByte('{')
	first := true
	_ = first
	if in.FullName != "" {
		const prefix string = ",\"BuyerFullName\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.FullName))
	}
	if in.EMail != "" {
		const prefix string = ",\"BuyerEmail\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.EMail))
	}
	if in.GSMNumber != "" {
		const prefix string = ",\"BuyerGsmNumber\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.GSMNumber))
	}
	if in.Address != "" {
		const prefix string = ",\"BuyerAddress\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Address))
	}
	out.RawByte('}')
}
func easyjson11d1a9baDecodeGithubComEmirmuminogluMokaposGo1(in *jlexer.Lexer, out *PaymentDealerAuthentication) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "DealerCode":
			out.Code = string(in.String())
		case "Username":
			out.Username = string(in.String())
		case "Password":
			out.Password = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson11d1a9baEncodeGithubComEmirmuminogluMokaposGo1(out *jwriter.Writer, in PaymentDealerAuthentication) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"DealerCode\":"
		out.RawString(prefix[1:])
		out.String(string(in.Code))
	}
	{
		const prefix string = ",\"Username\":"
		out.RawString(prefix)
		out.String(string(in.Username))
	}
	{
		const prefix string = ",\"Password\":"
		out.RawString(prefix)
		out.String(string(in.Password))
	}
	out.RawByte('}')
}
