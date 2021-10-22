package models

import (
	"encoding/json"
	"errors"
)

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
