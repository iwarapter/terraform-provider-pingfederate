package framework

import (
	"context"
	"net/http"

	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
	"github.com/iwarapter/pingfederate-sdk-go/services/oauthAuthenticationPolicyContractMappings"
)

type oauthAuthenticationPolicyContractMappingsMock struct {
	oauthAuthenticationPolicyContractMappings.OauthAuthenticationPolicyContractMappingsAPI
}

func (t oauthAuthenticationPolicyContractMappingsMock) CreateApcMappingWithContext(ctx context.Context, input *oauthAuthenticationPolicyContractMappings.CreateApcMappingInput) (output *pf.ApcToPersistentGrantMapping, resp *http.Response, err error) {
	return &pf.ApcToPersistentGrantMapping{
		Id: String("foo"),
		AuthenticationPolicyContractRef: &pf.ResourceLink{
			Id: String("foo"),
		},
		AttributeContractFulfillment: map[string]*pf.AttributeFulfillmentValue{},
		AttributeSources:             &[]*pf.AttributeSource{},
		IssuanceCriteria: &pf.IssuanceCriteria{
			ConditionalCriteria: &[]*pf.ConditionalIssuanceCriteriaEntry{
				{
					AttributeName: String("foo"),
					Condition:     String("foo"),
					ErrorResult:   String("foo"),
					Source: &pf.SourceTypeIdKey{
						Id:   String("foo"),
						Type: String("foo"),
					},
					Value: String("foo"),
				},
			},
		},
	}, nil, nil
}
