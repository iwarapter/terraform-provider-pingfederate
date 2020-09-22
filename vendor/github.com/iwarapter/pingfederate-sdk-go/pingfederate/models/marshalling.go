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
