package containerregistry

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// BuildTasksClient is the client for the BuildTasks methods of the Containerregistry service.
type BuildTasksClient struct {
	BaseClient
}

// NewBuildTasksClient creates an instance of the BuildTasksClient client.
func NewBuildTasksClient(subscriptionID string) BuildTasksClient {
	return NewBuildTasksClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewBuildTasksClientWithBaseURI creates an instance of the BuildTasksClient client.
func NewBuildTasksClientWithBaseURI(baseURI string, subscriptionID string) BuildTasksClient {
	return BuildTasksClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create creates a build task for a container registry with the specified parameters.
// Parameters:
// resourceGroupName - the name of the resource group to which the container registry belongs.
// registryName - the name of the container registry.
// buildTaskName - the name of the container registry build task.
// buildTaskCreateParameters - the parameters for creating a build task.
func (client BuildTasksClient) Create(ctx context.Context, resourceGroupName string, registryName string, buildTaskName string, buildTaskCreateParameters BuildTask) (result BuildTasksCreateFuture, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: registryName,
			Constraints: []validation.Constraint{{Target: "registryName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "registryName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "registryName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}},
		{TargetValue: buildTaskName,
			Constraints: []validation.Constraint{{Target: "buildTaskName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "buildTaskName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "buildTaskName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}},
		{TargetValue: buildTaskCreateParameters,
			Constraints: []validation.Constraint{{Target: "buildTaskCreateParameters.BuildTaskProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "buildTaskCreateParameters.BuildTaskProperties.Alias", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "buildTaskCreateParameters.BuildTaskProperties.SourceRepository", Name: validation.Null, Rule: true,
						Chain: []validation.Constraint{{Target: "buildTaskCreateParameters.BuildTaskProperties.SourceRepository.RepositoryURL", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "buildTaskCreateParameters.BuildTaskProperties.SourceRepository.SourceControlAuthProperties", Name: validation.Null, Rule: false,
								Chain: []validation.Constraint{{Target: "buildTaskCreateParameters.BuildTaskProperties.SourceRepository.SourceControlAuthProperties.Token", Name: validation.Null, Rule: true, Chain: nil}}},
						}},
					{Target: "buildTaskCreateParameters.BuildTaskProperties.Platform", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "buildTaskCreateParameters.BuildTaskProperties.Timeout", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "buildTaskCreateParameters.BuildTaskProperties.Timeout", Name: validation.InclusiveMaximum, Rule: 28800, Chain: nil},
							{Target: "buildTaskCreateParameters.BuildTaskProperties.Timeout", Name: validation.InclusiveMinimum, Rule: 300, Chain: nil},
						}},
				}}}}}); err != nil {
		return result, validation.NewError("containerregistry.BuildTasksClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, resourceGroupName, registryName, buildTaskName, buildTaskCreateParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.BuildTasksClient", "Create", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.BuildTasksClient", "Create", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client BuildTasksClient) CreatePreparer(ctx context.Context, resourceGroupName string, registryName string, buildTaskName string, buildTaskCreateParameters BuildTask) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"buildTaskName":     autorest.Encode("path", buildTaskName),
		"registryName":      autorest.Encode("path", registryName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-02-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/buildTasks/{buildTaskName}", pathParameters),
		autorest.WithJSON(buildTaskCreateParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client BuildTasksClient) CreateSender(req *http.Request) (future BuildTasksCreateFuture, err error) {
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	err = autorest.Respond(resp, azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client BuildTasksClient) CreateResponder(resp *http.Response) (result BuildTask, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a specified build task.
// Parameters:
// resourceGroupName - the name of the resource group to which the container registry belongs.
// registryName - the name of the container registry.
// buildTaskName - the name of the container registry build task.
func (client BuildTasksClient) Delete(ctx context.Context, resourceGroupName string, registryName string, buildTaskName string) (result BuildTasksDeleteFuture, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: registryName,
			Constraints: []validation.Constraint{{Target: "registryName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "registryName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "registryName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}},
		{TargetValue: buildTaskName,
			Constraints: []validation.Constraint{{Target: "buildTaskName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "buildTaskName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "buildTaskName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("containerregistry.BuildTasksClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, registryName, buildTaskName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.BuildTasksClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.BuildTasksClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client BuildTasksClient) DeletePreparer(ctx context.Context, resourceGroupName string, registryName string, buildTaskName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"buildTaskName":     autorest.Encode("path", buildTaskName),
		"registryName":      autorest.Encode("path", registryName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-02-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/buildTasks/{buildTaskName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client BuildTasksClient) DeleteSender(req *http.Request) (future BuildTasksDeleteFuture, err error) {
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	err = autorest.Respond(resp, azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client BuildTasksClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get the properties of a specified build task.
// Parameters:
// resourceGroupName - the name of the resource group to which the container registry belongs.
// registryName - the name of the container registry.
// buildTaskName - the name of the container registry build task.
func (client BuildTasksClient) Get(ctx context.Context, resourceGroupName string, registryName string, buildTaskName string) (result BuildTask, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: registryName,
			Constraints: []validation.Constraint{{Target: "registryName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "registryName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "registryName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}},
		{TargetValue: buildTaskName,
			Constraints: []validation.Constraint{{Target: "buildTaskName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "buildTaskName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "buildTaskName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("containerregistry.BuildTasksClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, registryName, buildTaskName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.BuildTasksClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "containerregistry.BuildTasksClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.BuildTasksClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client BuildTasksClient) GetPreparer(ctx context.Context, resourceGroupName string, registryName string, buildTaskName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"buildTaskName":     autorest.Encode("path", buildTaskName),
		"registryName":      autorest.Encode("path", registryName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-02-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/buildTasks/{buildTaskName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client BuildTasksClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client BuildTasksClient) GetResponder(resp *http.Response) (result BuildTask, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists all the build tasks for a specified container registry.
// Parameters:
// resourceGroupName - the name of the resource group to which the container registry belongs.
// registryName - the name of the container registry.
// filter - the build task filter to apply on the operation.
// skipToken - $skipToken is supported on get list of build tasks, which provides the next page in the list of
// tasks.
func (client BuildTasksClient) List(ctx context.Context, resourceGroupName string, registryName string, filter string, skipToken string) (result BuildTaskListResultPage, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: registryName,
			Constraints: []validation.Constraint{{Target: "registryName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "registryName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "registryName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("containerregistry.BuildTasksClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, registryName, filter, skipToken)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.BuildTasksClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.btlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "containerregistry.BuildTasksClient", "List", resp, "Failure sending request")
		return
	}

	result.btlr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.BuildTasksClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client BuildTasksClient) ListPreparer(ctx context.Context, resourceGroupName string, registryName string, filter string, skipToken string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"registryName":      autorest.Encode("path", registryName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-02-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if len(skipToken) > 0 {
		queryParameters["$skipToken"] = autorest.Encode("query", skipToken)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/buildTasks", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client BuildTasksClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client BuildTasksClient) ListResponder(resp *http.Response) (result BuildTaskListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client BuildTasksClient) listNextResults(lastResults BuildTaskListResult) (result BuildTaskListResult, err error) {
	req, err := lastResults.buildTaskListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "containerregistry.BuildTasksClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "containerregistry.BuildTasksClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.BuildTasksClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client BuildTasksClient) ListComplete(ctx context.Context, resourceGroupName string, registryName string, filter string, skipToken string) (result BuildTaskListResultIterator, err error) {
	result.page, err = client.List(ctx, resourceGroupName, registryName, filter, skipToken)
	return
}

// ListSourceRepositoryProperties get the source control properties for a build task.
// Parameters:
// resourceGroupName - the name of the resource group to which the container registry belongs.
// registryName - the name of the container registry.
// buildTaskName - the name of the container registry build task.
func (client BuildTasksClient) ListSourceRepositoryProperties(ctx context.Context, resourceGroupName string, registryName string, buildTaskName string) (result SourceRepositoryProperties, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: registryName,
			Constraints: []validation.Constraint{{Target: "registryName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "registryName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "registryName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}},
		{TargetValue: buildTaskName,
			Constraints: []validation.Constraint{{Target: "buildTaskName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "buildTaskName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "buildTaskName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("containerregistry.BuildTasksClient", "ListSourceRepositoryProperties", err.Error())
	}

	req, err := client.ListSourceRepositoryPropertiesPreparer(ctx, resourceGroupName, registryName, buildTaskName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.BuildTasksClient", "ListSourceRepositoryProperties", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSourceRepositoryPropertiesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "containerregistry.BuildTasksClient", "ListSourceRepositoryProperties", resp, "Failure sending request")
		return
	}

	result, err = client.ListSourceRepositoryPropertiesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.BuildTasksClient", "ListSourceRepositoryProperties", resp, "Failure responding to request")
	}

	return
}

// ListSourceRepositoryPropertiesPreparer prepares the ListSourceRepositoryProperties request.
func (client BuildTasksClient) ListSourceRepositoryPropertiesPreparer(ctx context.Context, resourceGroupName string, registryName string, buildTaskName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"buildTaskName":     autorest.Encode("path", buildTaskName),
		"registryName":      autorest.Encode("path", registryName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-02-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/buildTasks/{buildTaskName}/listSourceRepositoryProperties", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSourceRepositoryPropertiesSender sends the ListSourceRepositoryProperties request. The method will close the
// http.Response Body if it receives an error.
func (client BuildTasksClient) ListSourceRepositoryPropertiesSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListSourceRepositoryPropertiesResponder handles the response to the ListSourceRepositoryProperties request. The method always
// closes the http.Response Body.
func (client BuildTasksClient) ListSourceRepositoryPropertiesResponder(resp *http.Response) (result SourceRepositoryProperties, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Update updates a build task with the specified parameters.
// Parameters:
// resourceGroupName - the name of the resource group to which the container registry belongs.
// registryName - the name of the container registry.
// buildTaskName - the name of the container registry build task.
// buildTaskUpdateParameters - the parameters for updating a build task.
func (client BuildTasksClient) Update(ctx context.Context, resourceGroupName string, registryName string, buildTaskName string, buildTaskUpdateParameters BuildTaskUpdateParameters) (result BuildTasksUpdateFuture, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: registryName,
			Constraints: []validation.Constraint{{Target: "registryName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "registryName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "registryName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}},
		{TargetValue: buildTaskName,
			Constraints: []validation.Constraint{{Target: "buildTaskName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "buildTaskName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "buildTaskName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("containerregistry.BuildTasksClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, resourceGroupName, registryName, buildTaskName, buildTaskUpdateParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.BuildTasksClient", "Update", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.BuildTasksClient", "Update", result.Response(), "Failure sending request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client BuildTasksClient) UpdatePreparer(ctx context.Context, resourceGroupName string, registryName string, buildTaskName string, buildTaskUpdateParameters BuildTaskUpdateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"buildTaskName":     autorest.Encode("path", buildTaskName),
		"registryName":      autorest.Encode("path", registryName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-02-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/buildTasks/{buildTaskName}", pathParameters),
		autorest.WithJSON(buildTaskUpdateParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client BuildTasksClient) UpdateSender(req *http.Request) (future BuildTasksUpdateFuture, err error) {
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	err = autorest.Respond(resp, azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client BuildTasksClient) UpdateResponder(resp *http.Response) (result BuildTask, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
