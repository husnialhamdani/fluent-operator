package filter

import "kubesphere.io/fluentbit-operator/api/v1alpha2/plugins"

// +kubebuilder:object:generate:=true

// The Record Modifier Filter plugin allows to append fields or to exclude specific fields.
// RemoveKeys and WhitelistKeys are exclusive.
type RecordModifier struct {
	// Append fields. This parameter needs key and value pair.
	Records []string `json:"records,omitempty"`
	// If the key is matched, that field is removed.
	RemoveKeys []string `json:"removeKeys,omitempty"`
	// If the key is not matched, that field is removed.
	WhitelistKeys []string `json:"whitelistKeys,omitempty"`
}

func (_ *RecordModifier) Name() string {
	return "record_modifier"
}

func (rm *RecordModifier) Params(_ plugins.SecretLoader) (*plugins.KVs, error) {
	kvs := plugins.NewKVs()
	for _, record := range rm.Records {
		kvs.Insert("Record", record)
	}
	for _, key := range rm.RemoveKeys {
		kvs.Insert("Remove_key", key)
	}
	for _, key := range rm.WhitelistKeys {
		kvs.Insert("Whitelist_key", key)
	}
	return kvs, nil
}
