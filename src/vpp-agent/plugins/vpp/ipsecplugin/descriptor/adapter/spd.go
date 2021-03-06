// Code generated by adapter-generator. DO NOT EDIT.

package adapter

import (
	"google.golang.org/protobuf/proto"
	. "go.ligato.io/vpp-agent/v3/plugins/kvscheduler/api"
	"go.ligato.io/vpp-agent/v3/proto/ligato/vpp/ipsec"
)

////////// type-safe key-value pair with metadata //////////

type SPDKVWithMetadata struct {
	Key      string
	Value    *vpp_ipsec.SecurityPolicyDatabase
	Metadata interface{}
	Origin   ValueOrigin
}

////////// type-safe Descriptor structure //////////

type SPDDescriptor struct {
	Name                 string
	KeySelector          KeySelector
	ValueTypeName        string
	KeyLabel             func(key string) string
	ValueComparator      func(key string, oldValue, newValue *vpp_ipsec.SecurityPolicyDatabase) bool
	NBKeyPrefix          string
	WithMetadata         bool
	MetadataMapFactory   MetadataMapFactory
	Validate             func(key string, value *vpp_ipsec.SecurityPolicyDatabase) error
	Create               func(key string, value *vpp_ipsec.SecurityPolicyDatabase) (metadata interface{}, err error)
	Delete               func(key string, value *vpp_ipsec.SecurityPolicyDatabase, metadata interface{}) error
	Update               func(key string, oldValue, newValue *vpp_ipsec.SecurityPolicyDatabase, oldMetadata interface{}) (newMetadata interface{}, err error)
	UpdateWithRecreate   func(key string, oldValue, newValue *vpp_ipsec.SecurityPolicyDatabase, metadata interface{}) bool
	Retrieve             func(correlate []SPDKVWithMetadata) ([]SPDKVWithMetadata, error)
	IsRetriableFailure   func(err error) bool
	DerivedValues        func(key string, value *vpp_ipsec.SecurityPolicyDatabase) []KeyValuePair
	Dependencies         func(key string, value *vpp_ipsec.SecurityPolicyDatabase) []Dependency
	RetrieveDependencies []string /* descriptor name */
}

////////// Descriptor adapter //////////

type SPDDescriptorAdapter struct {
	descriptor *SPDDescriptor
}

func NewSPDDescriptor(typedDescriptor *SPDDescriptor) *KVDescriptor {
	adapter := &SPDDescriptorAdapter{descriptor: typedDescriptor}
	descriptor := &KVDescriptor{
		Name:                 typedDescriptor.Name,
		KeySelector:          typedDescriptor.KeySelector,
		ValueTypeName:        typedDescriptor.ValueTypeName,
		KeyLabel:             typedDescriptor.KeyLabel,
		NBKeyPrefix:          typedDescriptor.NBKeyPrefix,
		WithMetadata:         typedDescriptor.WithMetadata,
		MetadataMapFactory:   typedDescriptor.MetadataMapFactory,
		IsRetriableFailure:   typedDescriptor.IsRetriableFailure,
		RetrieveDependencies: typedDescriptor.RetrieveDependencies,
	}
	if typedDescriptor.ValueComparator != nil {
		descriptor.ValueComparator = adapter.ValueComparator
	}
	if typedDescriptor.Validate != nil {
		descriptor.Validate = adapter.Validate
	}
	if typedDescriptor.Create != nil {
		descriptor.Create = adapter.Create
	}
	if typedDescriptor.Delete != nil {
		descriptor.Delete = adapter.Delete
	}
	if typedDescriptor.Update != nil {
		descriptor.Update = adapter.Update
	}
	if typedDescriptor.UpdateWithRecreate != nil {
		descriptor.UpdateWithRecreate = adapter.UpdateWithRecreate
	}
	if typedDescriptor.Retrieve != nil {
		descriptor.Retrieve = adapter.Retrieve
	}
	if typedDescriptor.Dependencies != nil {
		descriptor.Dependencies = adapter.Dependencies
	}
	if typedDescriptor.DerivedValues != nil {
		descriptor.DerivedValues = adapter.DerivedValues
	}
	return descriptor
}

func (da *SPDDescriptorAdapter) ValueComparator(key string, oldValue, newValue proto.Message) bool {
	typedOldValue, err1 := castSPDValue(key, oldValue)
	typedNewValue, err2 := castSPDValue(key, newValue)
	if err1 != nil || err2 != nil {
		return false
	}
	return da.descriptor.ValueComparator(key, typedOldValue, typedNewValue)
}

func (da *SPDDescriptorAdapter) Validate(key string, value proto.Message) (err error) {
	typedValue, err := castSPDValue(key, value)
	if err != nil {
		return err
	}
	return da.descriptor.Validate(key, typedValue)
}

func (da *SPDDescriptorAdapter) Create(key string, value proto.Message) (metadata Metadata, err error) {
	typedValue, err := castSPDValue(key, value)
	if err != nil {
		return nil, err
	}
	return da.descriptor.Create(key, typedValue)
}

func (da *SPDDescriptorAdapter) Update(key string, oldValue, newValue proto.Message, oldMetadata Metadata) (newMetadata Metadata, err error) {
	oldTypedValue, err := castSPDValue(key, oldValue)
	if err != nil {
		return nil, err
	}
	newTypedValue, err := castSPDValue(key, newValue)
	if err != nil {
		return nil, err
	}
	typedOldMetadata, err := castSPDMetadata(key, oldMetadata)
	if err != nil {
		return nil, err
	}
	return da.descriptor.Update(key, oldTypedValue, newTypedValue, typedOldMetadata)
}

func (da *SPDDescriptorAdapter) Delete(key string, value proto.Message, metadata Metadata) error {
	typedValue, err := castSPDValue(key, value)
	if err != nil {
		return err
	}
	typedMetadata, err := castSPDMetadata(key, metadata)
	if err != nil {
		return err
	}
	return da.descriptor.Delete(key, typedValue, typedMetadata)
}

func (da *SPDDescriptorAdapter) UpdateWithRecreate(key string, oldValue, newValue proto.Message, metadata Metadata) bool {
	oldTypedValue, err := castSPDValue(key, oldValue)
	if err != nil {
		return true
	}
	newTypedValue, err := castSPDValue(key, newValue)
	if err != nil {
		return true
	}
	typedMetadata, err := castSPDMetadata(key, metadata)
	if err != nil {
		return true
	}
	return da.descriptor.UpdateWithRecreate(key, oldTypedValue, newTypedValue, typedMetadata)
}

func (da *SPDDescriptorAdapter) Retrieve(correlate []KVWithMetadata) ([]KVWithMetadata, error) {
	var correlateWithType []SPDKVWithMetadata
	for _, kvpair := range correlate {
		typedValue, err := castSPDValue(kvpair.Key, kvpair.Value)
		if err != nil {
			continue
		}
		typedMetadata, err := castSPDMetadata(kvpair.Key, kvpair.Metadata)
		if err != nil {
			continue
		}
		correlateWithType = append(correlateWithType,
			SPDKVWithMetadata{
				Key:      kvpair.Key,
				Value:    typedValue,
				Metadata: typedMetadata,
				Origin:   kvpair.Origin,
			})
	}

	typedValues, err := da.descriptor.Retrieve(correlateWithType)
	if err != nil {
		return nil, err
	}
	var values []KVWithMetadata
	for _, typedKVWithMetadata := range typedValues {
		kvWithMetadata := KVWithMetadata{
			Key:      typedKVWithMetadata.Key,
			Metadata: typedKVWithMetadata.Metadata,
			Origin:   typedKVWithMetadata.Origin,
		}
		kvWithMetadata.Value = typedKVWithMetadata.Value
		values = append(values, kvWithMetadata)
	}
	return values, err
}

func (da *SPDDescriptorAdapter) DerivedValues(key string, value proto.Message) []KeyValuePair {
	typedValue, err := castSPDValue(key, value)
	if err != nil {
		return nil
	}
	return da.descriptor.DerivedValues(key, typedValue)
}

func (da *SPDDescriptorAdapter) Dependencies(key string, value proto.Message) []Dependency {
	typedValue, err := castSPDValue(key, value)
	if err != nil {
		return nil
	}
	return da.descriptor.Dependencies(key, typedValue)
}

////////// Helper methods //////////

func castSPDValue(key string, value proto.Message) (*vpp_ipsec.SecurityPolicyDatabase, error) {
	typedValue, ok := value.(*vpp_ipsec.SecurityPolicyDatabase)
	if !ok {
		return nil, ErrInvalidValueType(key, value)
	}
	return typedValue, nil
}

func castSPDMetadata(key string, metadata Metadata) (interface{}, error) {
	if metadata == nil {
		return nil, nil
	}
	typedMetadata, ok := metadata.(interface{})
	if !ok {
		return nil, ErrInvalidMetadataType(key)
	}
	return typedMetadata, nil
}
