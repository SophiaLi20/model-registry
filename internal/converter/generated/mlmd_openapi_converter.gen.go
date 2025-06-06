// Code generated by github.com/jmattheis/goverter, DO NOT EDIT.
//go:build !goverter

package generated

import (
	"fmt"
	converter "github.com/kubeflow/model-registry/internal/converter"
	proto "github.com/kubeflow/model-registry/internal/ml_metadata/proto"
	openapi "github.com/kubeflow/model-registry/pkg/openapi"
)

type MLMDToOpenAPIConverterImpl struct{}

func (c *MLMDToOpenAPIConverterImpl) ConvertDocArtifact(source *proto.Artifact) (*openapi.DocArtifact, error) {
	var pOpenapiDocArtifact *openapi.DocArtifact
	if source != nil {
		var openapiDocArtifact openapi.DocArtifact
		mapStringOpenapiMetadataValue, err := converter.MapMLMDCustomProperties((*source).CustomProperties)
		if err != nil {
			return nil, fmt.Errorf("error setting field CustomProperties: %w", err)
		}
		openapiDocArtifact.CustomProperties = &mapStringOpenapiMetadataValue
		openapiDocArtifact.Description = converter.MapDescription((*source).Properties)
		if (*source).ExternalId != nil {
			xstring := *(*source).ExternalId
			openapiDocArtifact.ExternalId = &xstring
		}
		openapiDocArtifact.Name = converter.MapNameFromOwned((*source).Name)
		openapiDocArtifact.Id = converter.Int64ToString((*source).Id)
		openapiDocArtifact.CreateTimeSinceEpoch = converter.Int64ToString((*source).CreateTimeSinceEpoch)
		openapiDocArtifact.LastUpdateTimeSinceEpoch = converter.Int64ToString((*source).LastUpdateTimeSinceEpoch)
		pString, err := converter.MapArtifactType(source)
		if err != nil {
			return nil, fmt.Errorf("error setting field ArtifactType: %w", err)
		}
		openapiDocArtifact.ArtifactType = pString
		if (*source).Uri != nil {
			xstring2 := *(*source).Uri
			openapiDocArtifact.Uri = &xstring2
		}
		openapiDocArtifact.State = converter.MapMLMDArtifactState((*source).State)
		pOpenapiDocArtifact = &openapiDocArtifact
	}
	return pOpenapiDocArtifact, nil
}
func (c *MLMDToOpenAPIConverterImpl) ConvertInferenceService(source *proto.Context) (*openapi.InferenceService, error) {
	var pOpenapiInferenceService *openapi.InferenceService
	if source != nil {
		var openapiInferenceService openapi.InferenceService
		mapStringOpenapiMetadataValue, err := converter.MapMLMDCustomProperties((*source).CustomProperties)
		if err != nil {
			return nil, fmt.Errorf("error setting field CustomProperties: %w", err)
		}
		openapiInferenceService.CustomProperties = &mapStringOpenapiMetadataValue
		openapiInferenceService.Description = converter.MapDescription((*source).Properties)
		if (*source).ExternalId != nil {
			xstring := *(*source).ExternalId
			openapiInferenceService.ExternalId = &xstring
		}
		openapiInferenceService.Name = converter.MapNameFromOwned((*source).Name)
		openapiInferenceService.Id = converter.Int64ToString((*source).Id)
		openapiInferenceService.CreateTimeSinceEpoch = converter.Int64ToString((*source).CreateTimeSinceEpoch)
		openapiInferenceService.LastUpdateTimeSinceEpoch = converter.Int64ToString((*source).LastUpdateTimeSinceEpoch)
		openapiInferenceService.ModelVersionId = converter.MapPropertyModelVersionId((*source).Properties)
		openapiInferenceService.Runtime = converter.MapPropertyRuntime((*source).Properties)
		openapiInferenceService.DesiredState = converter.MapInferenceServiceDesiredState((*source).Properties)
		openapiInferenceService.RegisteredModelId = converter.MapPropertyRegisteredModelId((*source).Properties)
		openapiInferenceService.ServingEnvironmentId = converter.MapPropertyServingEnvironmentId((*source).Properties)
		pOpenapiInferenceService = &openapiInferenceService
	}
	return pOpenapiInferenceService, nil
}
func (c *MLMDToOpenAPIConverterImpl) ConvertModelArtifact(source *proto.Artifact) (*openapi.ModelArtifact, error) {
	var pOpenapiModelArtifact *openapi.ModelArtifact
	if source != nil {
		var openapiModelArtifact openapi.ModelArtifact
		mapStringOpenapiMetadataValue, err := converter.MapMLMDCustomProperties((*source).CustomProperties)
		if err != nil {
			return nil, fmt.Errorf("error setting field CustomProperties: %w", err)
		}
		openapiModelArtifact.CustomProperties = &mapStringOpenapiMetadataValue
		openapiModelArtifact.Description = converter.MapDescription((*source).Properties)
		if (*source).ExternalId != nil {
			xstring := *(*source).ExternalId
			openapiModelArtifact.ExternalId = &xstring
		}
		openapiModelArtifact.Name = converter.MapNameFromOwned((*source).Name)
		openapiModelArtifact.Id = converter.Int64ToString((*source).Id)
		openapiModelArtifact.CreateTimeSinceEpoch = converter.Int64ToString((*source).CreateTimeSinceEpoch)
		openapiModelArtifact.LastUpdateTimeSinceEpoch = converter.Int64ToString((*source).LastUpdateTimeSinceEpoch)
		pString, err := converter.MapArtifactType(source)
		if err != nil {
			return nil, fmt.Errorf("error setting field ArtifactType: %w", err)
		}
		openapiModelArtifact.ArtifactType = pString
		openapiModelArtifact.ModelFormatName = converter.MapModelArtifactFormatName((*source).Properties)
		openapiModelArtifact.StorageKey = converter.MapModelArtifactStorageKey((*source).Properties)
		openapiModelArtifact.StoragePath = converter.MapModelArtifactStoragePath((*source).Properties)
		openapiModelArtifact.ModelFormatVersion = converter.MapModelArtifactFormatVersion((*source).Properties)
		openapiModelArtifact.ServiceAccountName = converter.MapModelArtifactServiceAccountName((*source).Properties)
		openapiModelArtifact.ModelSourceKind = converter.MapModelArtifactModelSourceKind((*source).Properties)
		openapiModelArtifact.ModelSourceClass = converter.MapModelArtifactModelSourceClass((*source).Properties)
		openapiModelArtifact.ModelSourceGroup = converter.MapModelArtifactModelSourceGroup((*source).Properties)
		openapiModelArtifact.ModelSourceId = converter.MapModelArtifactModelSourceId((*source).Properties)
		openapiModelArtifact.ModelSourceName = converter.MapModelArtifactModelSourceName((*source).Properties)
		if (*source).Uri != nil {
			xstring2 := *(*source).Uri
			openapiModelArtifact.Uri = &xstring2
		}
		openapiModelArtifact.State = converter.MapMLMDArtifactState((*source).State)
		pOpenapiModelArtifact = &openapiModelArtifact
	}
	return pOpenapiModelArtifact, nil
}
func (c *MLMDToOpenAPIConverterImpl) ConvertModelVersion(source *proto.Context) (*openapi.ModelVersion, error) {
	var pOpenapiModelVersion *openapi.ModelVersion
	if source != nil {
		var openapiModelVersion openapi.ModelVersion
		mapStringOpenapiMetadataValue, err := converter.MapMLMDCustomProperties((*source).CustomProperties)
		if err != nil {
			return nil, fmt.Errorf("error setting field CustomProperties: %w", err)
		}
		openapiModelVersion.CustomProperties = &mapStringOpenapiMetadataValue
		openapiModelVersion.Description = converter.MapDescription((*source).Properties)
		if (*source).ExternalId != nil {
			xstring := *(*source).ExternalId
			openapiModelVersion.ExternalId = &xstring
		}
		openapiModelVersion.Name = converter.MapName((*source).Name)
		openapiModelVersion.State = converter.MapModelVersionState((*source).Properties)
		openapiModelVersion.Author = converter.MapPropertyAuthor((*source).Properties)
		xstring2, err := converter.MapRegisteredModelIdFromOwned((*source).Name)
		if err != nil {
			return nil, fmt.Errorf("error setting field RegisteredModelId: %w", err)
		}
		openapiModelVersion.RegisteredModelId = xstring2
		openapiModelVersion.Id = converter.Int64ToString((*source).Id)
		openapiModelVersion.CreateTimeSinceEpoch = converter.Int64ToString((*source).CreateTimeSinceEpoch)
		openapiModelVersion.LastUpdateTimeSinceEpoch = converter.Int64ToString((*source).LastUpdateTimeSinceEpoch)
		pOpenapiModelVersion = &openapiModelVersion
	}
	return pOpenapiModelVersion, nil
}
func (c *MLMDToOpenAPIConverterImpl) ConvertRegisteredModel(source *proto.Context) (*openapi.RegisteredModel, error) {
	var pOpenapiRegisteredModel *openapi.RegisteredModel
	if source != nil {
		var openapiRegisteredModel openapi.RegisteredModel
		mapStringOpenapiMetadataValue, err := converter.MapMLMDCustomProperties((*source).CustomProperties)
		if err != nil {
			return nil, fmt.Errorf("error setting field CustomProperties: %w", err)
		}
		openapiRegisteredModel.CustomProperties = &mapStringOpenapiMetadataValue
		openapiRegisteredModel.Description = converter.MapDescription((*source).Properties)
		if (*source).ExternalId != nil {
			xstring := *(*source).ExternalId
			openapiRegisteredModel.ExternalId = &xstring
		}
		if (*source).Name != nil {
			openapiRegisteredModel.Name = *(*source).Name
		}
		openapiRegisteredModel.Id = converter.Int64ToString((*source).Id)
		openapiRegisteredModel.CreateTimeSinceEpoch = converter.Int64ToString((*source).CreateTimeSinceEpoch)
		openapiRegisteredModel.LastUpdateTimeSinceEpoch = converter.Int64ToString((*source).LastUpdateTimeSinceEpoch)
		openapiRegisteredModel.Owner = converter.MapOwner((*source).Properties)
		openapiRegisteredModel.State = converter.MapRegisteredModelState((*source).Properties)
		pOpenapiRegisteredModel = &openapiRegisteredModel
	}
	return pOpenapiRegisteredModel, nil
}
func (c *MLMDToOpenAPIConverterImpl) ConvertServeModel(source *proto.Execution) (*openapi.ServeModel, error) {
	var pOpenapiServeModel *openapi.ServeModel
	if source != nil {
		var openapiServeModel openapi.ServeModel
		mapStringOpenapiMetadataValue, err := converter.MapMLMDCustomProperties((*source).CustomProperties)
		if err != nil {
			return nil, fmt.Errorf("error setting field CustomProperties: %w", err)
		}
		openapiServeModel.CustomProperties = &mapStringOpenapiMetadataValue
		openapiServeModel.Description = converter.MapDescription((*source).Properties)
		if (*source).ExternalId != nil {
			xstring := *(*source).ExternalId
			openapiServeModel.ExternalId = &xstring
		}
		openapiServeModel.Name = converter.MapNameFromOwned((*source).Name)
		openapiServeModel.Id = converter.Int64ToString((*source).Id)
		openapiServeModel.CreateTimeSinceEpoch = converter.Int64ToString((*source).CreateTimeSinceEpoch)
		openapiServeModel.LastUpdateTimeSinceEpoch = converter.Int64ToString((*source).LastUpdateTimeSinceEpoch)
		openapiServeModel.LastKnownState = converter.MapMLMDServeModelLastKnownState((*source).LastKnownState)
		openapiServeModel.ModelVersionId = converter.MapPropertyModelVersionIdAsValue((*source).Properties)
		pOpenapiServeModel = &openapiServeModel
	}
	return pOpenapiServeModel, nil
}
func (c *MLMDToOpenAPIConverterImpl) ConvertServingEnvironment(source *proto.Context) (*openapi.ServingEnvironment, error) {
	var pOpenapiServingEnvironment *openapi.ServingEnvironment
	if source != nil {
		var openapiServingEnvironment openapi.ServingEnvironment
		mapStringOpenapiMetadataValue, err := converter.MapMLMDCustomProperties((*source).CustomProperties)
		if err != nil {
			return nil, fmt.Errorf("error setting field CustomProperties: %w", err)
		}
		openapiServingEnvironment.CustomProperties = &mapStringOpenapiMetadataValue
		openapiServingEnvironment.Description = converter.MapDescription((*source).Properties)
		if (*source).ExternalId != nil {
			xstring := *(*source).ExternalId
			openapiServingEnvironment.ExternalId = &xstring
		}
		openapiServingEnvironment.Name = converter.MapName((*source).Name)
		openapiServingEnvironment.Id = converter.Int64ToString((*source).Id)
		openapiServingEnvironment.CreateTimeSinceEpoch = converter.Int64ToString((*source).CreateTimeSinceEpoch)
		openapiServingEnvironment.LastUpdateTimeSinceEpoch = converter.Int64ToString((*source).LastUpdateTimeSinceEpoch)
		pOpenapiServingEnvironment = &openapiServingEnvironment
	}
	return pOpenapiServingEnvironment, nil
}
