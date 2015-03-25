package cognitosync

// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

import (
	"time"

	"github.com/awslabs/aws-sdk-go/aws"
)

// DeleteDatasetRequest generates a request for the DeleteDataset operation.
func (c *CognitoSync) DeleteDatasetRequest(input *DeleteDatasetInput) (req *aws.Request, output *DeleteDatasetOutput) {
	if opDeleteDataset == nil {
		opDeleteDataset = &aws.Operation{
			Name:       "DeleteDataset",
			HTTPMethod: "DELETE",
			HTTPPath:   "/identitypools/{IdentityPoolId}/identities/{IdentityId}/datasets/{DatasetName}",
		}
	}

	req = aws.NewRequest(c.Service, opDeleteDataset, input, output)
	output = &DeleteDatasetOutput{}
	req.Data = output
	return
}

// Deletes the specific dataset. The dataset will be deleted permanently, and
// the action can't be undone. Datasets that this dataset was merged with will
// no longer report the merge. Any consequent operation on this dataset will
// result in a ResourceNotFoundException.
func (c *CognitoSync) DeleteDataset(input *DeleteDatasetInput) (output *DeleteDatasetOutput, err error) {
	req, out := c.DeleteDatasetRequest(input)
	output = out
	err = req.Send()
	return
}

var opDeleteDataset *aws.Operation

// DescribeDatasetRequest generates a request for the DescribeDataset operation.
func (c *CognitoSync) DescribeDatasetRequest(input *DescribeDatasetInput) (req *aws.Request, output *DescribeDatasetOutput) {
	if opDescribeDataset == nil {
		opDescribeDataset = &aws.Operation{
			Name:       "DescribeDataset",
			HTTPMethod: "GET",
			HTTPPath:   "/identitypools/{IdentityPoolId}/identities/{IdentityId}/datasets/{DatasetName}",
		}
	}

	req = aws.NewRequest(c.Service, opDescribeDataset, input, output)
	output = &DescribeDatasetOutput{}
	req.Data = output
	return
}

// Gets metadata about a dataset by identity and dataset name. The credentials
// used to make this API call need to have access to the identity data. With
// Amazon Cognito Sync, each identity has access only to its own data. You should
// use Amazon Cognito Identity service to retrieve the credentials necessary
// to make this API call.
func (c *CognitoSync) DescribeDataset(input *DescribeDatasetInput) (output *DescribeDatasetOutput, err error) {
	req, out := c.DescribeDatasetRequest(input)
	output = out
	err = req.Send()
	return
}

var opDescribeDataset *aws.Operation

// DescribeIdentityPoolUsageRequest generates a request for the DescribeIdentityPoolUsage operation.
func (c *CognitoSync) DescribeIdentityPoolUsageRequest(input *DescribeIdentityPoolUsageInput) (req *aws.Request, output *DescribeIdentityPoolUsageOutput) {
	if opDescribeIdentityPoolUsage == nil {
		opDescribeIdentityPoolUsage = &aws.Operation{
			Name:       "DescribeIdentityPoolUsage",
			HTTPMethod: "GET",
			HTTPPath:   "/identitypools/{IdentityPoolId}",
		}
	}

	req = aws.NewRequest(c.Service, opDescribeIdentityPoolUsage, input, output)
	output = &DescribeIdentityPoolUsageOutput{}
	req.Data = output
	return
}

// Gets usage details (for example, data storage) about a particular identity
// pool.
func (c *CognitoSync) DescribeIdentityPoolUsage(input *DescribeIdentityPoolUsageInput) (output *DescribeIdentityPoolUsageOutput, err error) {
	req, out := c.DescribeIdentityPoolUsageRequest(input)
	output = out
	err = req.Send()
	return
}

var opDescribeIdentityPoolUsage *aws.Operation

// DescribeIdentityUsageRequest generates a request for the DescribeIdentityUsage operation.
func (c *CognitoSync) DescribeIdentityUsageRequest(input *DescribeIdentityUsageInput) (req *aws.Request, output *DescribeIdentityUsageOutput) {
	if opDescribeIdentityUsage == nil {
		opDescribeIdentityUsage = &aws.Operation{
			Name:       "DescribeIdentityUsage",
			HTTPMethod: "GET",
			HTTPPath:   "/identitypools/{IdentityPoolId}/identities/{IdentityId}",
		}
	}

	req = aws.NewRequest(c.Service, opDescribeIdentityUsage, input, output)
	output = &DescribeIdentityUsageOutput{}
	req.Data = output
	return
}

// Gets usage information for an identity, including number of datasets and
// data usage.
func (c *CognitoSync) DescribeIdentityUsage(input *DescribeIdentityUsageInput) (output *DescribeIdentityUsageOutput, err error) {
	req, out := c.DescribeIdentityUsageRequest(input)
	output = out
	err = req.Send()
	return
}

var opDescribeIdentityUsage *aws.Operation

// GetIdentityPoolConfigurationRequest generates a request for the GetIdentityPoolConfiguration operation.
func (c *CognitoSync) GetIdentityPoolConfigurationRequest(input *GetIdentityPoolConfigurationInput) (req *aws.Request, output *GetIdentityPoolConfigurationOutput) {
	if opGetIdentityPoolConfiguration == nil {
		opGetIdentityPoolConfiguration = &aws.Operation{
			Name:       "GetIdentityPoolConfiguration",
			HTTPMethod: "GET",
			HTTPPath:   "/identitypools/{IdentityPoolId}/configuration",
		}
	}

	req = aws.NewRequest(c.Service, opGetIdentityPoolConfiguration, input, output)
	output = &GetIdentityPoolConfigurationOutput{}
	req.Data = output
	return
}

// Gets the configuration settings of an identity pool.
func (c *CognitoSync) GetIdentityPoolConfiguration(input *GetIdentityPoolConfigurationInput) (output *GetIdentityPoolConfigurationOutput, err error) {
	req, out := c.GetIdentityPoolConfigurationRequest(input)
	output = out
	err = req.Send()
	return
}

var opGetIdentityPoolConfiguration *aws.Operation

// ListDatasetsRequest generates a request for the ListDatasets operation.
func (c *CognitoSync) ListDatasetsRequest(input *ListDatasetsInput) (req *aws.Request, output *ListDatasetsOutput) {
	if opListDatasets == nil {
		opListDatasets = &aws.Operation{
			Name:       "ListDatasets",
			HTTPMethod: "GET",
			HTTPPath:   "/identitypools/{IdentityPoolId}/identities/{IdentityId}/datasets",
		}
	}

	req = aws.NewRequest(c.Service, opListDatasets, input, output)
	output = &ListDatasetsOutput{}
	req.Data = output
	return
}

// Lists datasets for an identity. The credentials used to make this API call
// need to have access to the identity data. With Amazon Cognito Sync, each
// identity has access only to its own data. You should use Amazon Cognito Identity
// service to retrieve the credentials necessary to make this API call.
func (c *CognitoSync) ListDatasets(input *ListDatasetsInput) (output *ListDatasetsOutput, err error) {
	req, out := c.ListDatasetsRequest(input)
	output = out
	err = req.Send()
	return
}

var opListDatasets *aws.Operation

// ListIdentityPoolUsageRequest generates a request for the ListIdentityPoolUsage operation.
func (c *CognitoSync) ListIdentityPoolUsageRequest(input *ListIdentityPoolUsageInput) (req *aws.Request, output *ListIdentityPoolUsageOutput) {
	if opListIdentityPoolUsage == nil {
		opListIdentityPoolUsage = &aws.Operation{
			Name:       "ListIdentityPoolUsage",
			HTTPMethod: "GET",
			HTTPPath:   "/identitypools",
		}
	}

	req = aws.NewRequest(c.Service, opListIdentityPoolUsage, input, output)
	output = &ListIdentityPoolUsageOutput{}
	req.Data = output
	return
}

// Gets a list of identity pools registered with Cognito.
func (c *CognitoSync) ListIdentityPoolUsage(input *ListIdentityPoolUsageInput) (output *ListIdentityPoolUsageOutput, err error) {
	req, out := c.ListIdentityPoolUsageRequest(input)
	output = out
	err = req.Send()
	return
}

var opListIdentityPoolUsage *aws.Operation

// ListRecordsRequest generates a request for the ListRecords operation.
func (c *CognitoSync) ListRecordsRequest(input *ListRecordsInput) (req *aws.Request, output *ListRecordsOutput) {
	if opListRecords == nil {
		opListRecords = &aws.Operation{
			Name:       "ListRecords",
			HTTPMethod: "GET",
			HTTPPath:   "/identitypools/{IdentityPoolId}/identities/{IdentityId}/datasets/{DatasetName}/records",
		}
	}

	req = aws.NewRequest(c.Service, opListRecords, input, output)
	output = &ListRecordsOutput{}
	req.Data = output
	return
}

// Gets paginated records, optionally changed after a particular sync count
// for a dataset and identity. The credentials used to make this API call need
// to have access to the identity data. With Amazon Cognito Sync, each identity
// has access only to its own data. You should use Amazon Cognito Identity service
// to retrieve the credentials necessary to make this API call.
func (c *CognitoSync) ListRecords(input *ListRecordsInput) (output *ListRecordsOutput, err error) {
	req, out := c.ListRecordsRequest(input)
	output = out
	err = req.Send()
	return
}

var opListRecords *aws.Operation

// RegisterDeviceRequest generates a request for the RegisterDevice operation.
func (c *CognitoSync) RegisterDeviceRequest(input *RegisterDeviceInput) (req *aws.Request, output *RegisterDeviceOutput) {
	if opRegisterDevice == nil {
		opRegisterDevice = &aws.Operation{
			Name:       "RegisterDevice",
			HTTPMethod: "POST",
			HTTPPath:   "/identitypools/{IdentityPoolId}/identity/{IdentityId}/device",
		}
	}

	req = aws.NewRequest(c.Service, opRegisterDevice, input, output)
	output = &RegisterDeviceOutput{}
	req.Data = output
	return
}

// Registers a device to receive push sync notifications.
func (c *CognitoSync) RegisterDevice(input *RegisterDeviceInput) (output *RegisterDeviceOutput, err error) {
	req, out := c.RegisterDeviceRequest(input)
	output = out
	err = req.Send()
	return
}

var opRegisterDevice *aws.Operation

// SetIdentityPoolConfigurationRequest generates a request for the SetIdentityPoolConfiguration operation.
func (c *CognitoSync) SetIdentityPoolConfigurationRequest(input *SetIdentityPoolConfigurationInput) (req *aws.Request, output *SetIdentityPoolConfigurationOutput) {
	if opSetIdentityPoolConfiguration == nil {
		opSetIdentityPoolConfiguration = &aws.Operation{
			Name:       "SetIdentityPoolConfiguration",
			HTTPMethod: "POST",
			HTTPPath:   "/identitypools/{IdentityPoolId}/configuration",
		}
	}

	req = aws.NewRequest(c.Service, opSetIdentityPoolConfiguration, input, output)
	output = &SetIdentityPoolConfigurationOutput{}
	req.Data = output
	return
}

// Sets the necessary configuration for push sync.
func (c *CognitoSync) SetIdentityPoolConfiguration(input *SetIdentityPoolConfigurationInput) (output *SetIdentityPoolConfigurationOutput, err error) {
	req, out := c.SetIdentityPoolConfigurationRequest(input)
	output = out
	err = req.Send()
	return
}

var opSetIdentityPoolConfiguration *aws.Operation

// SubscribeToDatasetRequest generates a request for the SubscribeToDataset operation.
func (c *CognitoSync) SubscribeToDatasetRequest(input *SubscribeToDatasetInput) (req *aws.Request, output *SubscribeToDatasetOutput) {
	if opSubscribeToDataset == nil {
		opSubscribeToDataset = &aws.Operation{
			Name:       "SubscribeToDataset",
			HTTPMethod: "POST",
			HTTPPath:   "/identitypools/{IdentityPoolId}/identities/{IdentityId}/datasets/{DatasetName}/subscriptions/{DeviceId}",
		}
	}

	req = aws.NewRequest(c.Service, opSubscribeToDataset, input, output)
	output = &SubscribeToDatasetOutput{}
	req.Data = output
	return
}

// Subscribes to receive notifications when a dataset is modified by another
// device.
func (c *CognitoSync) SubscribeToDataset(input *SubscribeToDatasetInput) (output *SubscribeToDatasetOutput, err error) {
	req, out := c.SubscribeToDatasetRequest(input)
	output = out
	err = req.Send()
	return
}

var opSubscribeToDataset *aws.Operation

// UnsubscribeFromDatasetRequest generates a request for the UnsubscribeFromDataset operation.
func (c *CognitoSync) UnsubscribeFromDatasetRequest(input *UnsubscribeFromDatasetInput) (req *aws.Request, output *UnsubscribeFromDatasetOutput) {
	if opUnsubscribeFromDataset == nil {
		opUnsubscribeFromDataset = &aws.Operation{
			Name:       "UnsubscribeFromDataset",
			HTTPMethod: "DELETE",
			HTTPPath:   "/identitypools/{IdentityPoolId}/identities/{IdentityId}/datasets/{DatasetName}/subscriptions/{DeviceId}",
		}
	}

	req = aws.NewRequest(c.Service, opUnsubscribeFromDataset, input, output)
	output = &UnsubscribeFromDatasetOutput{}
	req.Data = output
	return
}

// Unsubscribe from receiving notifications when a dataset is modified by another
// device.
func (c *CognitoSync) UnsubscribeFromDataset(input *UnsubscribeFromDatasetInput) (output *UnsubscribeFromDatasetOutput, err error) {
	req, out := c.UnsubscribeFromDatasetRequest(input)
	output = out
	err = req.Send()
	return
}

var opUnsubscribeFromDataset *aws.Operation

// UpdateRecordsRequest generates a request for the UpdateRecords operation.
func (c *CognitoSync) UpdateRecordsRequest(input *UpdateRecordsInput) (req *aws.Request, output *UpdateRecordsOutput) {
	if opUpdateRecords == nil {
		opUpdateRecords = &aws.Operation{
			Name:       "UpdateRecords",
			HTTPMethod: "POST",
			HTTPPath:   "/identitypools/{IdentityPoolId}/identities/{IdentityId}/datasets/{DatasetName}",
		}
	}

	req = aws.NewRequest(c.Service, opUpdateRecords, input, output)
	output = &UpdateRecordsOutput{}
	req.Data = output
	return
}

// Posts updates to records and add and delete records for a dataset and user.
// The credentials used to make this API call need to have access to the identity
// data. With Amazon Cognito Sync, each identity has access only to its own
// data. You should use Amazon Cognito Identity service to retrieve the credentials
// necessary to make this API call.
func (c *CognitoSync) UpdateRecords(input *UpdateRecordsInput) (output *UpdateRecordsOutput, err error) {
	req, out := c.UpdateRecordsRequest(input)
	output = out
	err = req.Send()
	return
}

var opUpdateRecords *aws.Operation

// A collection of data for an identity pool. An identity pool can have multiple
// datasets. A dataset is per identity and can be general or associated with
// a particular entity in an application (like a saved game). Datasets are automatically
// created if they don't exist. Data is synced by dataset, and a dataset can
// hold up to 1MB of key-value pairs.
type Dataset struct {
	// Date on which the dataset was created.
	CreationDate *time.Time `type:"timestamp" timestampFormat:"unix"`

	// Total size in bytes of the records in this dataset.
	DataStorage *int64 `type:"long"`

	// A string of up to 128 characters. Allowed characters are a-z, A-Z, 0-9, '_'
	// (underscore), '-' (dash), and '.' (dot).
	DatasetName *string `type:"string"`

	// A name-spaced GUID (for example, us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE)
	// created by Amazon Cognito. GUID generation is unique within a region.
	IdentityID *string `locationName:"IdentityId" type:"string"`

	// The device that made the last change to this dataset.
	LastModifiedBy *string `type:"string"`

	// Date when the dataset was last modified.
	LastModifiedDate *time.Time `type:"timestamp" timestampFormat:"unix"`

	// Number of records in this dataset.
	NumRecords *int64 `type:"long"`

	metadataDataset `json:"-", xml:"-"`
}

type metadataDataset struct {
	SDKShapeTraits bool `type:"structure"`
}

// A request to delete the specific dataset.
type DeleteDatasetInput struct {
	// A string of up to 128 characters. Allowed characters are a-z, A-Z, 0-9, '_'
	// (underscore), '-' (dash), and '.' (dot).
	DatasetName *string `location:"uri" locationName:"DatasetName" type:"string" required:"true"`

	// A name-spaced GUID (for example, us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE)
	// created by Amazon Cognito. GUID generation is unique within a region.
	IdentityID *string `location:"uri" locationName:"IdentityId" type:"string" required:"true"`

	// A name-spaced GUID (for example, us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE)
	// created by Amazon Cognito. GUID generation is unique within a region.
	IdentityPoolID *string `location:"uri" locationName:"IdentityPoolId" type:"string" required:"true"`

	metadataDeleteDatasetInput `json:"-", xml:"-"`
}

type metadataDeleteDatasetInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// Response to a successful DeleteDataset request.
type DeleteDatasetOutput struct {
	// A collection of data for an identity pool. An identity pool can have multiple
	// datasets. A dataset is per identity and can be general or associated with
	// a particular entity in an application (like a saved game). Datasets are automatically
	// created if they don't exist. Data is synced by dataset, and a dataset can
	// hold up to 1MB of key-value pairs.
	Dataset *Dataset `type:"structure"`

	metadataDeleteDatasetOutput `json:"-", xml:"-"`
}

type metadataDeleteDatasetOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// A request for metadata about a dataset (creation date, number of records,
// size) by owner and dataset name.
type DescribeDatasetInput struct {
	// A string of up to 128 characters. Allowed characters are a-z, A-Z, 0-9, '_'
	// (underscore), '-' (dash), and '.' (dot).
	DatasetName *string `location:"uri" locationName:"DatasetName" type:"string" required:"true"`

	// A name-spaced GUID (for example, us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE)
	// created by Amazon Cognito. GUID generation is unique within a region.
	IdentityID *string `location:"uri" locationName:"IdentityId" type:"string" required:"true"`

	// A name-spaced GUID (for example, us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE)
	// created by Amazon Cognito. GUID generation is unique within a region.
	IdentityPoolID *string `location:"uri" locationName:"IdentityPoolId" type:"string" required:"true"`

	metadataDescribeDatasetInput `json:"-", xml:"-"`
}

type metadataDescribeDatasetInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// Response to a successful DescribeDataset request.
type DescribeDatasetOutput struct {
	// Metadata for a collection of data for an identity. An identity can have multiple
	// datasets. A dataset can be general or associated with a particular entity
	// in an application (like a saved game). Datasets are automatically created
	// if they don't exist. Data is synced by dataset, and a dataset can hold up
	// to 1MB of key-value pairs.
	Dataset *Dataset `type:"structure"`

	metadataDescribeDatasetOutput `json:"-", xml:"-"`
}

type metadataDescribeDatasetOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// A request for usage information about the identity pool.
type DescribeIdentityPoolUsageInput struct {
	// A name-spaced GUID (for example, us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE)
	// created by Amazon Cognito. GUID generation is unique within a region.
	IdentityPoolID *string `location:"uri" locationName:"IdentityPoolId" type:"string" required:"true"`

	metadataDescribeIdentityPoolUsageInput `json:"-", xml:"-"`
}

type metadataDescribeIdentityPoolUsageInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// Response to a successful DescribeIdentityPoolUsage request.
type DescribeIdentityPoolUsageOutput struct {
	// Information about the usage of the identity pool.
	IdentityPoolUsage *IdentityPoolUsage `type:"structure"`

	metadataDescribeIdentityPoolUsageOutput `json:"-", xml:"-"`
}

type metadataDescribeIdentityPoolUsageOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// A request for information about the usage of an identity pool.
type DescribeIdentityUsageInput struct {
	// A name-spaced GUID (for example, us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE)
	// created by Amazon Cognito. GUID generation is unique within a region.
	IdentityID *string `location:"uri" locationName:"IdentityId" type:"string" required:"true"`

	// A name-spaced GUID (for example, us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE)
	// created by Amazon Cognito. GUID generation is unique within a region.
	IdentityPoolID *string `location:"uri" locationName:"IdentityPoolId" type:"string" required:"true"`

	metadataDescribeIdentityUsageInput `json:"-", xml:"-"`
}

type metadataDescribeIdentityUsageInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// The response to a successful DescribeIdentityUsage request.
type DescribeIdentityUsageOutput struct {
	// Usage information for the identity.
	IdentityUsage *IdentityUsage `type:"structure"`

	metadataDescribeIdentityUsageOutput `json:"-", xml:"-"`
}

type metadataDescribeIdentityUsageOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// A request to GetIdentityPoolConfigurationRequest.
type GetIdentityPoolConfigurationInput struct {
	// A name-spaced GUID (for example, us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE)
	// created by Amazon Cognito. This is the ID of the pool for which to return
	// a configuration.
	IdentityPoolID *string `location:"uri" locationName:"IdentityPoolId" type:"string" required:"true"`

	metadataGetIdentityPoolConfigurationInput `json:"-", xml:"-"`
}

type metadataGetIdentityPoolConfigurationInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// The response from GetIdentityPoolConfigurationResponse.
type GetIdentityPoolConfigurationOutput struct {
	// A name-spaced GUID (for example, us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE)
	// created by Amazon Cognito.
	IdentityPoolID *string `locationName:"IdentityPoolId" type:"string"`

	// Configuration options applied to the identity pool.
	PushSync *PushSync `type:"structure"`

	metadataGetIdentityPoolConfigurationOutput `json:"-", xml:"-"`
}

type metadataGetIdentityPoolConfigurationOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// Usage information for the identity pool.
type IdentityPoolUsage struct {
	// Data storage information for the identity pool.
	DataStorage *int64 `type:"long"`

	// A name-spaced GUID (for example, us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE)
	// created by Amazon Cognito. GUID generation is unique within a region.
	IdentityPoolID *string `locationName:"IdentityPoolId" type:"string"`

	// Date on which the identity pool was last modified.
	LastModifiedDate *time.Time `type:"timestamp" timestampFormat:"unix"`

	// Number of sync sessions for the identity pool.
	SyncSessionsCount *int64 `type:"long"`

	metadataIdentityPoolUsage `json:"-", xml:"-"`
}

type metadataIdentityPoolUsage struct {
	SDKShapeTraits bool `type:"structure"`
}

// Usage information for the identity.
type IdentityUsage struct {
	// Total data storage for this identity.
	DataStorage *int64 `type:"long"`

	// Number of datasets for the identity.
	DatasetCount *int64 `type:"integer"`

	// A name-spaced GUID (for example, us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE)
	// created by Amazon Cognito. GUID generation is unique within a region.
	IdentityID *string `locationName:"IdentityId" type:"string"`

	// A name-spaced GUID (for example, us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE)
	// created by Amazon Cognito. GUID generation is unique within a region.
	IdentityPoolID *string `locationName:"IdentityPoolId" type:"string"`

	// Date on which the identity was last modified.
	LastModifiedDate *time.Time `type:"timestamp" timestampFormat:"unix"`

	metadataIdentityUsage `json:"-", xml:"-"`
}

type metadataIdentityUsage struct {
	SDKShapeTraits bool `type:"structure"`
}

// Request for a list of datasets for an identity.
type ListDatasetsInput struct {
	// A name-spaced GUID (for example, us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE)
	// created by Amazon Cognito. GUID generation is unique within a region.
	IdentityID *string `location:"uri" locationName:"IdentityId" type:"string" required:"true"`

	// A name-spaced GUID (for example, us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE)
	// created by Amazon Cognito. GUID generation is unique within a region.
	IdentityPoolID *string `location:"uri" locationName:"IdentityPoolId" type:"string" required:"true"`

	// The maximum number of results to be returned.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" type:"integer"`

	// A pagination token for obtaining the next page of results.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`

	metadataListDatasetsInput `json:"-", xml:"-"`
}

type metadataListDatasetsInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// Returned for a successful ListDatasets request.
type ListDatasetsOutput struct {
	// Number of datasets returned.
	Count *int64 `type:"integer"`

	// A set of datasets.
	Datasets []*Dataset `type:"list"`

	// A pagination token for obtaining the next page of results.
	NextToken *string `type:"string"`

	metadataListDatasetsOutput `json:"-", xml:"-"`
}

type metadataListDatasetsOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// A request for usage information on an identity pool.
type ListIdentityPoolUsageInput struct {
	// The maximum number of results to be returned.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" type:"integer"`

	// A pagination token for obtaining the next page of results.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`

	metadataListIdentityPoolUsageInput `json:"-", xml:"-"`
}

type metadataListIdentityPoolUsageInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// Returned for a successful ListIdentityPoolUsage request.
type ListIdentityPoolUsageOutput struct {
	// Total number of identities for the identity pool.
	Count *int64 `type:"integer"`

	// Usage information for the identity pools.
	IdentityPoolUsages []*IdentityPoolUsage `type:"list"`

	// The maximum number of results to be returned.
	MaxResults *int64 `type:"integer"`

	// A pagination token for obtaining the next page of results.
	NextToken *string `type:"string"`

	metadataListIdentityPoolUsageOutput `json:"-", xml:"-"`
}

type metadataListIdentityPoolUsageOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// A request for a list of records.
type ListRecordsInput struct {
	// A string of up to 128 characters. Allowed characters are a-z, A-Z, 0-9, '_'
	// (underscore), '-' (dash), and '.' (dot).
	DatasetName *string `location:"uri" locationName:"DatasetName" type:"string" required:"true"`

	// A name-spaced GUID (for example, us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE)
	// created by Amazon Cognito. GUID generation is unique within a region.
	IdentityID *string `location:"uri" locationName:"IdentityId" type:"string" required:"true"`

	// A name-spaced GUID (for example, us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE)
	// created by Amazon Cognito. GUID generation is unique within a region.
	IdentityPoolID *string `location:"uri" locationName:"IdentityPoolId" type:"string" required:"true"`

	// The last server sync count for this record.
	LastSyncCount *int64 `location:"querystring" locationName:"lastSyncCount" type:"long"`

	// The maximum number of results to be returned.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" type:"integer"`

	// A pagination token for obtaining the next page of results.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`

	// A token containing a session ID, identity ID, and expiration.
	SyncSessionToken *string `location:"querystring" locationName:"syncSessionToken" type:"string"`

	metadataListRecordsInput `json:"-", xml:"-"`
}

type metadataListRecordsInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// Returned for a successful ListRecordsRequest.
type ListRecordsOutput struct {
	// Total number of records.
	Count *int64 `type:"integer"`

	// A boolean value specifying whether to delete the dataset locally.
	DatasetDeletedAfterRequestedSyncCount *bool `type:"boolean"`

	// Indicates whether the dataset exists.
	DatasetExists *bool `type:"boolean"`

	// Server sync count for this dataset.
	DatasetSyncCount *int64 `type:"long"`

	// The user/device that made the last change to this record.
	LastModifiedBy *string `type:"string"`

	// Names of merged datasets.
	MergedDatasetNames []*string `type:"list"`

	// A pagination token for obtaining the next page of results.
	NextToken *string `type:"string"`

	// A list of all records.
	Records []*Record `type:"list"`

	// A token containing a session ID, identity ID, and expiration.
	SyncSessionToken *string `type:"string"`

	metadataListRecordsOutput `json:"-", xml:"-"`
}

type metadataListRecordsOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// Configuration options to be applied to the identity pool.
type PushSync struct {
	// List of SNS platform application ARNs that could be used by clients.
	ApplicationARNs []*string `locationName:"ApplicationArns" type:"list"`

	// A role configured to allow Cognito to call SNS on behalf of the developer.
	RoleARN *string `locationName:"RoleArn" type:"string"`

	metadataPushSync `json:"-", xml:"-"`
}

type metadataPushSync struct {
	SDKShapeTraits bool `type:"structure"`
}

// The basic data structure of a dataset.
type Record struct {
	// The last modified date of the client device.
	DeviceLastModifiedDate *time.Time `type:"timestamp" timestampFormat:"unix"`

	// The key for the record.
	Key *string `type:"string"`

	// The user/device that made the last change to this record.
	LastModifiedBy *string `type:"string"`

	// The date on which the record was last modified.
	LastModifiedDate *time.Time `type:"timestamp" timestampFormat:"unix"`

	// The server sync count for this record.
	SyncCount *int64 `type:"long"`

	// The value for the record.
	Value *string `type:"string"`

	metadataRecord `json:"-", xml:"-"`
}

type metadataRecord struct {
	SDKShapeTraits bool `type:"structure"`
}

// An update operation for a record.
type RecordPatch struct {
	// The last modified date of the client device.
	DeviceLastModifiedDate *time.Time `type:"timestamp" timestampFormat:"unix"`

	// The key associated with the record patch.
	Key *string `type:"string" required:"true"`

	// An operation, either replace or remove.
	Op *string `type:"string" required:"true"`

	// Last known server sync count for this record. Set to 0 if unknown.
	SyncCount *int64 `type:"long" required:"true"`

	// The value associated with the record patch.
	Value *string `type:"string"`

	metadataRecordPatch `json:"-", xml:"-"`
}

type metadataRecordPatch struct {
	SDKShapeTraits bool `type:"structure"`
}

// A request to RegisterDevice.
type RegisterDeviceInput struct {
	// The unique ID for this identity.
	IdentityID *string `location:"uri" locationName:"IdentityId" type:"string" required:"true"`

	// A name-spaced GUID (for example, us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE)
	// created by Amazon Cognito. Here, the ID of the pool that the identity belongs
	// to.
	IdentityPoolID *string `location:"uri" locationName:"IdentityPoolId" type:"string" required:"true"`

	// The SNS platform type (e.g. GCM, SDM, APNS, APNS_SANDBOX).
	Platform *string `type:"string" required:"true"`

	// The push token.
	Token *string `type:"string" required:"true"`

	metadataRegisterDeviceInput `json:"-", xml:"-"`
}

type metadataRegisterDeviceInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// Response to a RegisterDevice request.
type RegisterDeviceOutput struct {
	// The unique ID generated for this device by Cognito.
	DeviceID *string `locationName:"DeviceId" type:"string"`

	metadataRegisterDeviceOutput `json:"-", xml:"-"`
}

type metadataRegisterDeviceOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// A request to SetIdentityPoolConfiguration.
type SetIdentityPoolConfigurationInput struct {
	// A name-spaced GUID (for example, us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE)
	// created by Amazon Cognito. This is the ID of the pool to modify.
	IdentityPoolID *string `location:"uri" locationName:"IdentityPoolId" type:"string" required:"true"`

	// Configuration options to be applied to the identity pool.
	PushSync *PushSync `type:"structure"`

	metadataSetIdentityPoolConfigurationInput `json:"-", xml:"-"`
}

type metadataSetIdentityPoolConfigurationInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// Response to a SetIdentityPoolConfiguration request.
type SetIdentityPoolConfigurationOutput struct {
	// A name-spaced GUID (for example, us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE)
	// created by Amazon Cognito.
	IdentityPoolID *string `locationName:"IdentityPoolId" type:"string"`

	// Configuration options applied to the identity pool.
	PushSync *PushSync `type:"structure"`

	metadataSetIdentityPoolConfigurationOutput `json:"-", xml:"-"`
}

type metadataSetIdentityPoolConfigurationOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// A request to SubscribeToDatasetRequest.
type SubscribeToDatasetInput struct {
	// The name of the dataset to subcribe to.
	DatasetName *string `location:"uri" locationName:"DatasetName" type:"string" required:"true"`

	// The unique ID generated for this device by Cognito.
	DeviceID *string `location:"uri" locationName:"DeviceId" type:"string" required:"true"`

	// Unique ID for this identity.
	IdentityID *string `location:"uri" locationName:"IdentityId" type:"string" required:"true"`

	// A name-spaced GUID (for example, us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE)
	// created by Amazon Cognito. The ID of the pool to which the identity belongs.
	IdentityPoolID *string `location:"uri" locationName:"IdentityPoolId" type:"string" required:"true"`

	metadataSubscribeToDatasetInput `json:"-", xml:"-"`
}

type metadataSubscribeToDatasetInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// Response to a SubscribeToDataset request.
type SubscribeToDatasetOutput struct {
	metadataSubscribeToDatasetOutput `json:"-", xml:"-"`
}

type metadataSubscribeToDatasetOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// A request to UnsubscribeFromDataset.
type UnsubscribeFromDatasetInput struct {
	// The name of the dataset from which to unsubcribe.
	DatasetName *string `location:"uri" locationName:"DatasetName" type:"string" required:"true"`

	// The unique ID generated for this device by Cognito.
	DeviceID *string `location:"uri" locationName:"DeviceId" type:"string" required:"true"`

	// Unique ID for this identity.
	IdentityID *string `location:"uri" locationName:"IdentityId" type:"string" required:"true"`

	// A name-spaced GUID (for example, us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE)
	// created by Amazon Cognito. The ID of the pool to which this identity belongs.
	IdentityPoolID *string `location:"uri" locationName:"IdentityPoolId" type:"string" required:"true"`

	metadataUnsubscribeFromDatasetInput `json:"-", xml:"-"`
}

type metadataUnsubscribeFromDatasetInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// Response to an UnsubscribeFromDataset request.
type UnsubscribeFromDatasetOutput struct {
	metadataUnsubscribeFromDatasetOutput `json:"-", xml:"-"`
}

type metadataUnsubscribeFromDatasetOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// A request to post updates to records or add and delete records for a dataset
// and user.
type UpdateRecordsInput struct {
	// Intended to supply a device ID that will populate the lastModifiedBy field
	// referenced in other methods. The ClientContext field is not yet implemented.
	ClientContext *string `location:"header" locationName:"x-amz-Client-Context" type:"string"`

	// A string of up to 128 characters. Allowed characters are a-z, A-Z, 0-9, '_'
	// (underscore), '-' (dash), and '.' (dot).
	DatasetName *string `location:"uri" locationName:"DatasetName" type:"string" required:"true"`

	// The unique ID generated for this device by Cognito.
	DeviceID *string `locationName:"DeviceId" type:"string"`

	// A name-spaced GUID (for example, us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE)
	// created by Amazon Cognito. GUID generation is unique within a region.
	IdentityID *string `location:"uri" locationName:"IdentityId" type:"string" required:"true"`

	// A name-spaced GUID (for example, us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE)
	// created by Amazon Cognito. GUID generation is unique within a region.
	IdentityPoolID *string `location:"uri" locationName:"IdentityPoolId" type:"string" required:"true"`

	// A list of patch operations.
	RecordPatches []*RecordPatch `type:"list"`

	// The SyncSessionToken returned by a previous call to ListRecords for this
	// dataset and identity.
	SyncSessionToken *string `type:"string" required:"true"`

	metadataUpdateRecordsInput `json:"-", xml:"-"`
}

type metadataUpdateRecordsInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// Returned for a successful UpdateRecordsRequest.
type UpdateRecordsOutput struct {
	// A list of records that have been updated.
	Records []*Record `type:"list"`

	metadataUpdateRecordsOutput `json:"-", xml:"-"`
}

type metadataUpdateRecordsOutput struct {
	SDKShapeTraits bool `type:"structure"`
}