package powervs

import (
	"context"
	"github.com/openshift/installer/pkg/types"
	"sync"
)

//go:generate mockgen -source=./metadata.go -destination=./mock/powervsmetadata_generated.go -package=mock

// MetadataAPI represents functions that eventually call out to the API
type MetadataAPI interface {
	AccountID(ctx context.Context) (string, error)
	APIKey(ctx context.Context) (string, error)
	CISInstanceCRN(ctx context.Context) (string, error)
	DNSInstanceCRN(ctx context.Context) (string, error)
}

// Metadata holds additional metadata for InstallConfig resources that
// do not need to be user-supplied (e.g. because it can be retrieved
// from external APIs).
type Metadata struct {
	BaseDomain string

	accountID      string
	apiKey         string
	cisInstanceCRN string
	dnsInstanceCRN string
	client         *Client

	mutex sync.Mutex
}

// NewMetadata initializes a new Metadata object.
func NewMetadata(baseDomain string) *Metadata {
	return &Metadata{BaseDomain: baseDomain}
}

// AccountID returns the IBM Cloud account ID associated with the authentication
// credentials.
func (m *Metadata) AccountID(ctx context.Context) (string, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if m.client == nil {
		client, err := NewClient()
		if err != nil {
			return "", err
		}

		m.client = client
	}

	if m.accountID == "" {
		apiKeyDetails, err := m.client.GetAuthenticatorAPIKeyDetails(ctx)
		if err != nil {
			return "", err
		}

		m.accountID = *apiKeyDetails.AccountID
	}

	return m.accountID, nil
}

// APIKey returns the IBM Cloud account API Key associated with the authentication
// credentials.
func (m *Metadata) APIKey(ctx context.Context) (string, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if m.client == nil {
		client, err := NewClient()
		if err != nil {
			return "", err
		}

		m.client = client
	}

	if m.apiKey == "" {
		m.apiKey = m.client.GetAPIKey()
	}

	return m.apiKey, nil
}

// CISInstanceCRN returns the Cloud Internet Services instance CRN that is
// managing the DNS zone for the base domain.
func (m *Metadata) CISInstanceCRN(ctx context.Context) (string, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	var err error
	if m.client == nil {
		client, err := NewClient()
		if err != nil {
			return "", err
		}

		m.client = client
	}

	if m.cisInstanceCRN == "" {
		m.cisInstanceCRN, err = m.client.GetInstanceCRNByName(ctx, m.BaseDomain, types.ExternalPublishingStrategy)
		if err != nil {
			return "", err
		}
	}
	return m.cisInstanceCRN, nil
}

// SetCISInstanceCRN sets Cloud Internet Services instance CRN to a string value.
func (m *Metadata) SetCISInstanceCRN(crn string) {
	m.cisInstanceCRN = crn
}

// DNSInstanceCRN returns the IBM DNS Service instance CRN that is
// managing the DNS zone for the base domain.
func (m *Metadata) DNSInstanceCRN(ctx context.Context) (string, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	var err error
	if m.client == nil {
		client, err := NewClient()
		if err != nil {
			return "", err
		}

		m.client = client
	}

	if m.dnsInstanceCRN == "" {
		m.dnsInstanceCRN, err = m.client.GetInstanceCRNByName(ctx, m.BaseDomain, types.InternalPublishingStrategy)
		if err != nil {
			return "", err
		}
	}

	return m.dnsInstanceCRN, nil
}

// SetDNSInstanceCRN sets IBM DNS Service instance CRN to a string value.
func (m *Metadata) SetDNSInstanceCRN(crn string) {
	m.dnsInstanceCRN = crn
}
