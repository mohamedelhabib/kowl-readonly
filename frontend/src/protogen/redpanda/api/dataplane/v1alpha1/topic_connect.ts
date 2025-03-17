// @generated by protoc-gen-connect-es v1.2.0 with parameter "target=ts,import_extension="
// @generated from file redpanda/api/dataplane/v1alpha1/topic.proto (package redpanda.api.dataplane.v1alpha1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { CreateTopicRequest, CreateTopicResponse, DeleteTopicRequest, DeleteTopicResponse, GetTopicConfigurationsRequest, GetTopicConfigurationsResponse, ListTopicsRequest, ListTopicsResponse, SetTopicConfigurationsRequest, SetTopicConfigurationsResponse, UpdateTopicConfigurationsRequest, UpdateTopicConfigurationsResponse } from "./topic_pb";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service redpanda.api.dataplane.v1alpha1.TopicService
 * @deprecated
 */
export const TopicService = {
  typeName: "redpanda.api.dataplane.v1alpha1.TopicService",
  methods: {
    /**
     * @generated from rpc redpanda.api.dataplane.v1alpha1.TopicService.CreateTopic
     * @deprecated
     */
    createTopic: {
      name: "CreateTopic",
      I: CreateTopicRequest,
      O: CreateTopicResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc redpanda.api.dataplane.v1alpha1.TopicService.ListTopics
     * @deprecated
     */
    listTopics: {
      name: "ListTopics",
      I: ListTopicsRequest,
      O: ListTopicsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc redpanda.api.dataplane.v1alpha1.TopicService.DeleteTopic
     * @deprecated
     */
    deleteTopic: {
      name: "DeleteTopic",
      I: DeleteTopicRequest,
      O: DeleteTopicResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc redpanda.api.dataplane.v1alpha1.TopicService.GetTopicConfigurations
     * @deprecated
     */
    getTopicConfigurations: {
      name: "GetTopicConfigurations",
      I: GetTopicConfigurationsRequest,
      O: GetTopicConfigurationsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc redpanda.api.dataplane.v1alpha1.TopicService.UpdateTopicConfigurations
     * @deprecated
     */
    updateTopicConfigurations: {
      name: "UpdateTopicConfigurations",
      I: UpdateTopicConfigurationsRequest,
      O: UpdateTopicConfigurationsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc redpanda.api.dataplane.v1alpha1.TopicService.SetTopicConfigurations
     * @deprecated
     */
    setTopicConfigurations: {
      name: "SetTopicConfigurations",
      I: SetTopicConfigurationsRequest,
      O: SetTopicConfigurationsResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

