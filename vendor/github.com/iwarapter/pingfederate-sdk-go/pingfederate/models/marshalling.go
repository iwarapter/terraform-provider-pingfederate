package models

import (
	"encoding/json"
	"errors"
	"fmt"
)

func (f *AttributeSource) UnmarshalJSON(b []byte) error {
	type resp AttributeSource
	err := json.Unmarshal(b, (*resp)(f))
	if err != nil {
		return err
	}
	switch *f.Type {
	case "LDAP":
		f.LdapAttributeSource = LdapAttributeSource{
			AttributeContractFulfillment: f.AttributeContractFulfillment,
			BaseDn:                       f.BaseDn,
			BinaryAttributeSettings:      f.BinaryAttributeSettings,
			DataStoreRef:                 f.DataStoreRef,
			Description:                  f.Description,
			Id:                           f.Id,
			MemberOfNestedGroup:          f.MemberOfNestedGroup,
			SearchAttributes:             f.SearchAttributes,
			SearchFilter:                 f.SearchFilter,
			SearchScope:                  f.SearchScope,
			Type:                         f.Type,
		}
	case "JDBC":
		f.JdbcAttributeSource = JdbcAttributeSource{
			AttributeContractFulfillment: f.AttributeContractFulfillment,
			ColumnNames:                  f.ColumnNames,
			DataStoreRef:                 f.DataStoreRef,
			Description:                  f.Description,
			Filter:                       f.Filter,
			Id:                           f.Id,
			Schema:                       f.Schema,
			Table:                        f.Table,
			Type:                         f.Type,
		}
	case "CUSTOM":
		f.CustomAttributeSource = CustomAttributeSource{
			AttributeContractFulfillment: f.AttributeContractFulfillment,
			DataStoreRef:                 f.DataStoreRef,
			Description:                  f.Description,
			FilterFields:                 f.FilterFields,
			Id:                           f.Id,
			Type:                         f.Type,
		}
	default:
		return errors.New("unknown data store type")
	}
	return nil
}

func (f *DataStores) UnmarshalJSON(b []byte) error {
	type resp DataStores
	err := json.Unmarshal(b, (*resp)(f))
	if err != nil {
		return err
	}
	if f.Items == nil {
		f.Items = &[]interface{}{}
	}

	for _, raw := range *f.RawItems {
		var v DataStore
		err = json.Unmarshal(raw, &v)
		if err != nil {
			return err
		}
		var i interface{}
		switch *v.Type {
		case "LDAP":
			i = &LdapDataStore{}
		case "JDBC":
			i = &JdbcDataStore{}
		case "CUSTOM":
			i = &CustomDataStore{}
		default:
			return errors.New("unknown data store type")
		}
		err = json.Unmarshal(raw, i)
		if err != nil {
			return err
		}
		*f.Items = append(*f.Items, i)
	}
	return nil
}

func (f *FieldDescriptor) UnmarshalJSON(b []byte) error {
	type resp FieldDescriptor
	err := json.Unmarshal(b, (*resp)(f))
	if err != nil {
		return err
	}
	switch *f.Type {
	case "RADIO_GROUP":
		err := json.Unmarshal(b, &f.RadioGroupFieldDescriptor)
		if err != nil {
			return err
		}
	case "SELECT_FIELD":
		err := json.Unmarshal(b, &f.SelectFieldDescriptor)
		if err != nil {
			return err
		}
	case "CHECK_BOX":
		err := json.Unmarshal(b, &f.CheckBoxFieldDescriptor)
		if err != nil {
			return err
		}
	case "UPLOAD_FILE":
		err := json.Unmarshal(b, &f.UploadFileFieldDescriptor)
		if err != nil {
			return err
		}
	case "TEXT_AREA":
		err := json.Unmarshal(b, &f.TextAreaFieldDescriptor)
		if err != nil {
			return err
		}
	case "FILTERABLE_SELECT":
		err := json.Unmarshal(b, &f.TextFieldDescriptor)
		if err != nil {
			return err
		}
	case "HASHED_TEXT":
		err := json.Unmarshal(b, &f.HashedTextFieldDescriptor)
		if err != nil {
			return err
		}
	}
	return nil
}

func (f *PolicyAction) UnmarshalJSON(b []byte) error {
	type resp PolicyAction
	err := json.Unmarshal(b, (*resp)(f))
	if err != nil {
		return err
	}
	switch *f.Type {
	case "AUTHN_SOURCE":
		err := json.Unmarshal(b, &f.AuthnSourcePolicyAction)
		if err != nil {
			return err
		}
	case "APC_MAPPING":
		err := json.Unmarshal(b, &f.ApcMappingPolicyAction)
		if err != nil {
			return err
		}
	case "LOCAL_IDENTITY_MAPPING":
		err := json.Unmarshal(b, &f.LocalIdentityMappingPolicyAction)
		if err != nil {
			return err
		}
	case "AUTHN_SELECTOR":
		err := json.Unmarshal(b, &f.AuthnSelectorPolicyAction)
		if err != nil {
			return err
		}
	case "FRAGMENT":
		err := json.Unmarshal(b, &f.FragmentPolicyAction)
		if err != nil {
			return err
		}
	}
	return nil
}

func (f *PolicyAction) MarshalJSON() ([]byte, error) {
	if f.Type == nil {
		return nil, fmt.Errorf("unable to marshal Type is a required field")
	}
	switch *f.Type {
	case "AUTHN_SOURCE":
		f.AuthnSourcePolicyAction.Context = f.Context
		f.AuthnSourcePolicyAction.Type = f.Type
		return json.Marshal(&f.AuthnSourcePolicyAction)
	case "APC_MAPPING":
		f.ApcMappingPolicyAction.Context = f.Context
		f.ApcMappingPolicyAction.Type = f.Type
		return json.Marshal(&f.ApcMappingPolicyAction)
	case "LOCAL_IDENTITY_MAPPING":
		f.LocalIdentityMappingPolicyAction.Context = f.Context
		f.LocalIdentityMappingPolicyAction.Type = f.Type
		return json.Marshal(&f.LocalIdentityMappingPolicyAction)
	case "AUTHN_SELECTOR":
		f.AuthnSelectorPolicyAction.Context = f.Context
		f.AuthnSelectorPolicyAction.Type = f.Type
		return json.Marshal(&f.AuthnSelectorPolicyAction)
	case "FRAGMENT":
		f.FragmentPolicyAction.Context = f.Context
		f.FragmentPolicyAction.Type = f.Type
		return json.Marshal(&f.FragmentPolicyAction)
	case "RESTART":
		f.RestartPolicyAction.Context = f.Context
		f.RestartPolicyAction.Type = f.Type
		return json.Marshal(&f.RestartPolicyAction)
	case "DONE":
		f.DonePolicyAction.Context = f.Context
		f.DonePolicyAction.Type = f.Type
		return json.Marshal(&f.DonePolicyAction)
	case "CONTINUE":
		f.ContinuePolicyAction.Context = f.Context
		f.ContinuePolicyAction.Type = f.Type
		return json.Marshal(&f.ContinuePolicyAction)
	default:
		return nil, fmt.Errorf("no known type to marshal: %s", *f.Type)
	}
}
